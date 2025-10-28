// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateK8sResourceRequest
	GetClusterId() *string
	SetNamespace(v string) *UpdateK8sResourceRequest
	GetNamespace() *string
	SetResourceContent(v string) *UpdateK8sResourceRequest
	GetResourceContent() *string
}

type UpdateK8sResourceRequest struct {
	// The ID of the cluster.
	//
	// example:
	//
	// 2e7059e9-2d********5e8ecac64ff
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the namespace to which the Kubernetes resource belongs.
	//
	// example:
	//
	// app-namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The description of the resource in the YAML format.
	//
	// example:
	//
	// apiVersion: apps/v1 kind: Deployment metadata:   name: demo-app   namespace: app-namespace spec:   replicas: 3   selector:     matchLabels:       cluster: abc   template: # create pods using pod definition in this template     metadata:       labels:         cluster: abc     spec:       containers:       - image: registry-vpc.cn-hangzhou.aliyuncs.com/edas-demo-image/consumer:1.0         imagePullPolicy: Always         name: test-container         ports:         - containerPort: 80         env:         - name: foo           value: bar
	ResourceContent *string `json:"ResourceContent,omitempty" xml:"ResourceContent,omitempty"`
}

func (s UpdateK8sResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateK8sResourceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateK8sResourceRequest) GetResourceContent() *string {
	return s.ResourceContent
}

func (s *UpdateK8sResourceRequest) SetClusterId(v string) *UpdateK8sResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sResourceRequest) SetNamespace(v string) *UpdateK8sResourceRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateK8sResourceRequest) SetResourceContent(v string) *UpdateK8sResourceRequest {
	s.ResourceContent = &v
	return s
}

func (s *UpdateK8sResourceRequest) Validate() error {
	return dara.Validate(s)
}
