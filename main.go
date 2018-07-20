package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

//    "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"

    "fmt"
    "os"
//    "path/filepath"
)

var region = os.Getenv("AWS_DEFAULT_REGION")

type kubeInstanceInfo struct {
    dns string
    imageId string
    instanceId string
    old bool
    region string
    state string
}

func getInstances() ([]kubeInstanceInfo, error) {
    sess, err := session.NewSession()

    if err != nil {
        return nil, err
    }

    ec2Svc := ec2.New(sess, aws.NewConfig().WithRegion(region))
    var instances []kubeInstanceInfo

    result, err := ec2Svc.DescribeInstances(nil)
    if err != nil {
        return nil, err
    } else {
        for _, reservation := range result.Reservations {
            for _, instance := range reservation.Instances {
                instanceDetails := kubeInstanceInfo{
                    dns: *instance.PrivateDnsName,
                    imageId: *instance.ImageId,
                    instanceId: *instance.InstanceId,
                    old: true,
                    region: *instance.Placement.AvailabilityZone,
                    state: "unknown",
                }
                instances = append (instances, instanceDetails)
            }
        }

    }
    return instances, nil
}

func main() {
/*    result, err := getInstances()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)

        os.Exit(1)
    }
    fmt.Println(result)
*/
    kubeconfig := "/Users/drosth/.kube/alpaca-dev.conf"

    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(nodes)
}
