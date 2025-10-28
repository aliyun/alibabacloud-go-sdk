// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertK8sResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConvertK8sResourceRequest
	GetClusterId() *string
	SetNamespace(v string) *ConvertK8sResourceRequest
	GetNamespace() *string
	SetResourceName(v string) *ConvertK8sResourceRequest
	GetResourceName() *string
	SetResourceType(v string) *ConvertK8sResourceRequest
	GetResourceType() *string
}

type ConvertK8sResourceRequest struct {
	// The ID of the cluster. You can call the ListCluster operation to query the cluster ID. For more information, see [ListCluster](https://help.aliyun.com/document_detail/154995.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// b07c8192-****-adf4f7447720
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// deployment-to-convert
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource that is used. Set the value to deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// deployment
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ConvertK8sResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertK8sResourceRequest) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConvertK8sResourceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ConvertK8sResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ConvertK8sResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ConvertK8sResourceRequest) SetClusterId(v string) *ConvertK8sResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *ConvertK8sResourceRequest) SetNamespace(v string) *ConvertK8sResourceRequest {
	s.Namespace = &v
	return s
}

func (s *ConvertK8sResourceRequest) SetResourceName(v string) *ConvertK8sResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *ConvertK8sResourceRequest) SetResourceType(v string) *ConvertK8sResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *ConvertK8sResourceRequest) Validate() error {
	return dara.Validate(s)
}
