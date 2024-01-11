# code-generator

Golang code-generators used to implement [Kubernetes-style API types](https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md).

## Purpose

These code-generators can be used

- in the context of [CustomResourceDefinition](https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/) to build native, versioned clients,
  informers and other helpers
- in the context of [User-provider API Servers](https://github.com/kubernetes/apiserver) to build conversions between internal and versioned types, defaulters, protobuf codecs,
  internal and versioned clients and informers.

## Resources

- The example [sample controller](https://github.com/kubernetes/sample-controller) shows a code example of a controller that uses the clients, listers and informers generated by this library.
- The article [Kubernetes Deep Dive: Code Generation for CustomResources](https://cloud.redhat.com/blog/kubernetes-deep-dive-code-generation-customresources/) gives a step by step instruction on how to use this library.

## Compatibility

HEAD of this repo will match HEAD of k8s.io/apiserver, k8s.io/apimachinery, and k8s.io/client-go.

## Where does it come from?

`code-generator` is synced from <https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/code-generator>.
Code changes are made in that location, merged into `k8s.io/kubernetes` and later synced here.
