// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAutoScalingRecordsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ListAutoScalingRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAutoScalingRecordsRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *ListAutoScalingRecordsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAutoScalingRecordsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListAutoScalingRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAutoScalingRecordsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ListAutoScalingRecordsRequest
	GetSecurityToken() *string
}

type ListAutoScalingRecordsRequest struct {
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAutoScalingRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAutoScalingRecordsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAutoScalingRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutoScalingRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAutoScalingRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAutoScalingRecordsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListAutoScalingRecordsRequest) SetInstanceId(v string) *ListAutoScalingRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetOwnerAccount(v string) *ListAutoScalingRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetOwnerId(v int64) *ListAutoScalingRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetPageNum(v int32) *ListAutoScalingRecordsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetPageSize(v int32) *ListAutoScalingRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetResourceOwnerAccount(v string) *ListAutoScalingRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetResourceOwnerId(v int64) *ListAutoScalingRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetSecurityToken(v string) *ListAutoScalingRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) Validate() error {
	return dara.Validate(s)
}
