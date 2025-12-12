// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAutoScalingRecordsShrinkRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ListAutoScalingRecordsShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAutoScalingRecordsShrinkRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *ListAutoScalingRecordsShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAutoScalingRecordsShrinkRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListAutoScalingRecordsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAutoScalingRecordsShrinkRequest
	GetResourceOwnerId() *int64
	SetScaleTypesShrink(v string) *ListAutoScalingRecordsShrinkRequest
	GetScaleTypesShrink() *string
	SetSecurityToken(v string) *ListAutoScalingRecordsShrinkRequest
	GetSecurityToken() *string
}

type ListAutoScalingRecordsShrinkRequest struct {
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
	ScaleTypesShrink     *string `json:"ScaleTypes,omitempty" xml:"ScaleTypes,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingRecordsShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAutoScalingRecordsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAutoScalingRecordsShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAutoScalingRecordsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutoScalingRecordsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAutoScalingRecordsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAutoScalingRecordsShrinkRequest) GetScaleTypesShrink() *string {
	return s.ScaleTypesShrink
}

func (s *ListAutoScalingRecordsShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListAutoScalingRecordsShrinkRequest) SetInstanceId(v string) *ListAutoScalingRecordsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetOwnerAccount(v string) *ListAutoScalingRecordsShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetOwnerId(v int64) *ListAutoScalingRecordsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetPageNum(v int32) *ListAutoScalingRecordsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetPageSize(v int32) *ListAutoScalingRecordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetResourceOwnerAccount(v string) *ListAutoScalingRecordsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetResourceOwnerId(v int64) *ListAutoScalingRecordsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetScaleTypesShrink(v string) *ListAutoScalingRecordsShrinkRequest {
	s.ScaleTypesShrink = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) SetSecurityToken(v string) *ListAutoScalingRecordsShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListAutoScalingRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
