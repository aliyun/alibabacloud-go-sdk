// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableOperation(v bool) *UpdateServiceInstanceAttributesRequest
	GetEnableOperation() *bool
	SetRegionId(v string) *UpdateServiceInstanceAttributesRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceAttributesRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceAttributesRequest struct {
	// Specifies whether to authorize the service provider to perform O\\&M operations on the service instance.
	//
	// example:
	//
	// true
	EnableOperation *bool `json:"EnableOperation,omitempty" xml:"EnableOperation,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17xxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesRequest) GetEnableOperation() *bool {
	return s.EnableOperation
}

func (s *UpdateServiceInstanceAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceInstanceAttributesRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceAttributesRequest) SetEnableOperation(v bool) *UpdateServiceInstanceAttributesRequest {
	s.EnableOperation = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequest) SetRegionId(v string) *UpdateServiceInstanceAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequest) Validate() error {
	return dara.Validate(s)
}
