// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiredDay(v string) *AddBlacklistShrinkRequest
	GetExpiredDay() *string
	SetNumbersShrink(v string) *AddBlacklistShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *AddBlacklistShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddBlacklistShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddBlacklistShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddBlacklistShrinkRequest
	GetResourceOwnerId() *int64
}

type AddBlacklistShrinkRequest struct {
	// 有效天数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// 号码列表
	//
	// This parameter is required.
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 备注
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddBlacklistShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddBlacklistShrinkRequest) GetExpiredDay() *string {
	return s.ExpiredDay
}

func (s *AddBlacklistShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *AddBlacklistShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddBlacklistShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddBlacklistShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddBlacklistShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddBlacklistShrinkRequest) SetExpiredDay(v string) *AddBlacklistShrinkRequest {
	s.ExpiredDay = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetNumbersShrink(v string) *AddBlacklistShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetOwnerId(v int64) *AddBlacklistShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetRemark(v string) *AddBlacklistShrinkRequest {
	s.Remark = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetResourceOwnerAccount(v string) *AddBlacklistShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetResourceOwnerId(v int64) *AddBlacklistShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddBlacklistShrinkRequest) Validate() error {
	return dara.Validate(s)
}
