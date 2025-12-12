// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingConfigsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAutoScalingConfigsShrinkRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ListAutoScalingConfigsShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAutoScalingConfigsShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAutoScalingConfigsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAutoScalingConfigsShrinkRequest
	GetResourceOwnerId() *int64
	SetScaleTypesShrink(v string) *ListAutoScalingConfigsShrinkRequest
	GetScaleTypesShrink() *string
	SetSecurityToken(v string) *ListAutoScalingConfigsShrinkRequest
	GetSecurityToken() *string
}

type ListAutoScalingConfigsShrinkRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleTypesShrink     *string `json:"ScaleTypes,omitempty" xml:"ScaleTypes,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingConfigsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingConfigsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAutoScalingConfigsShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAutoScalingConfigsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAutoScalingConfigsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAutoScalingConfigsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAutoScalingConfigsShrinkRequest) GetScaleTypesShrink() *string {
	return s.ScaleTypesShrink
}

func (s *ListAutoScalingConfigsShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListAutoScalingConfigsShrinkRequest) SetInstanceId(v string) *ListAutoScalingConfigsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) SetOwnerAccount(v string) *ListAutoScalingConfigsShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) SetOwnerId(v int64) *ListAutoScalingConfigsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) SetResourceOwnerAccount(v string) *ListAutoScalingConfigsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) SetResourceOwnerId(v int64) *ListAutoScalingConfigsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) SetScaleTypesShrink(v string) *ListAutoScalingConfigsShrinkRequest {
	s.ScaleTypesShrink = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) SetSecurityToken(v string) *ListAutoScalingConfigsShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListAutoScalingConfigsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
