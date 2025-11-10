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
	SetGrantedPermission(v *UpdateServiceInstanceAttributesRequestGrantedPermission) *UpdateServiceInstanceAttributesRequest
	GetGrantedPermission() *UpdateServiceInstanceAttributesRequestGrantedPermission
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
	EnableOperation   *bool                                                    `json:"EnableOperation,omitempty" xml:"EnableOperation,omitempty"`
	GrantedPermission *UpdateServiceInstanceAttributesRequestGrantedPermission `json:"GrantedPermission,omitempty" xml:"GrantedPermission,omitempty" type:"Struct"`
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

func (s *UpdateServiceInstanceAttributesRequest) GetGrantedPermission() *UpdateServiceInstanceAttributesRequestGrantedPermission {
	return s.GrantedPermission
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

func (s *UpdateServiceInstanceAttributesRequest) SetGrantedPermission(v *UpdateServiceInstanceAttributesRequestGrantedPermission) *UpdateServiceInstanceAttributesRequest {
	s.GrantedPermission = v
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
	if s.GrantedPermission != nil {
		if err := s.GrantedPermission.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceInstanceAttributesRequestGrantedPermission struct {
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	PolicyNames      *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
}

func (s UpdateServiceInstanceAttributesRequestGrantedPermission) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributesRequestGrantedPermission) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesRequestGrantedPermission) GetOperationEndTime() *string {
	return s.OperationEndTime
}

func (s *UpdateServiceInstanceAttributesRequestGrantedPermission) GetPolicyNames() *string {
	return s.PolicyNames
}

func (s *UpdateServiceInstanceAttributesRequestGrantedPermission) SetOperationEndTime(v string) *UpdateServiceInstanceAttributesRequestGrantedPermission {
	s.OperationEndTime = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequestGrantedPermission) SetPolicyNames(v string) *UpdateServiceInstanceAttributesRequestGrantedPermission {
	s.PolicyNames = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequestGrantedPermission) Validate() error {
	return dara.Validate(s)
}
