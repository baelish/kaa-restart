package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
    "os"
)

var region = os.Getenv("AWS_DEFAULT_REGION")

type kubeInstanceInfo struct {
    dns string
    imageId string
    instanceId string
    old bool
    state string
}

type kubeInstances struct {
    info []kubeInstanceInfo
}

func main() {
    sess, err := session.NewSession()

    if err != nil {
        fmt.Println("Error creating session ", err)
        return
    }

    ec2Svc := ec2.New(sess, aws.NewConfig().WithRegion(region))
    result, err := ec2Svc.DescribeInstances(nil)

    if err != nil {
        fmt.Println("Error", err)
    } else {
        for _, reservation := range result.Reservations {
            for _, instance := range reservation.Instances {
                instanceDetails := kubeInstanceInfo{
                    dns: *instance.PrivateDnsName,
                    imageId: *instance.ImageId,
                    instanceId: *instance.InstanceId,
                    old: true,
                    state: "",
                }
                fmt.Println(instanceDetails)
            }
        }

    }


}
