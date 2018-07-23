# kaa-restart
Restart AWS Instances in an Autoscaling group that form part of a Kubernetes cluster in a controlled manner.

plan:
- find current nodes in Autoscaling group and mark them old
- check the state of each node in Kubernetes
- if any are not ready handle them first.
- spin up new node in asg in same az as node to be removed
- when new node is ready in Kubernetes
  - drain the old node
  - remove and terminate it from Autoscaling group
- Repeat until no old nodes are left
