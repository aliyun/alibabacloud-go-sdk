// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAutoScalingConfigsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ListAutoScalingConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAutoScalingConfigsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAutoScalingConfigsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAutoScalingConfigsRequest
	GetResourceOwnerId() *int64
	SetScaleTypes(v []*string) *ListAutoScalingConfigsRequest
	GetScaleTypes() []*string
	SetSecurityToken(v string) *ListAutoScalingConfigsRequest
	GetSecurityToken() *string
}

type ListAutoScalingConfigsRequest struct {
	// This parameter is required.
	InstanceId           *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleTypes           []*string `json:"ScaleTypes,omitempty" xml:"ScaleTypes,omitempty" type:"Repeated"`
	SecurityToken        *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAutoScalingConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAutoScalingConfigsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAutoScalingConfigsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAutoScalingConfigsRequest) GetScaleTypes() []*string {
	return s.ScaleTypes
}

func (s *ListAutoScalingConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListAutoScalingConfigsRequest) SetInstanceId(v string) *ListAutoScalingConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetOwnerAccount(v string) *ListAutoScalingConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetOwnerId(v int64) *ListAutoScalingConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetResourceOwnerAccount(v string) *ListAutoScalingConfigsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetResourceOwnerId(v int64) *ListAutoScalingConfigsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetScaleTypes(v []*string) *ListAutoScalingConfigsRequest {
	s.ScaleTypes = v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetSecurityToken(v string) *ListAutoScalingConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) Validate() error {
	return dara.Validate(s)
}
