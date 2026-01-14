// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiredDay(v string) *AddBlacklistRequest
	GetExpiredDay() *string
	SetNumbers(v []*string) *AddBlacklistRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *AddBlacklistRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddBlacklistRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddBlacklistRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddBlacklistRequest
	GetResourceOwnerId() *int64
}

type AddBlacklistRequest struct {
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
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s AddBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddBlacklistRequest) GetExpiredDay() *string {
	return s.ExpiredDay
}

func (s *AddBlacklistRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *AddBlacklistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddBlacklistRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddBlacklistRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddBlacklistRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddBlacklistRequest) SetExpiredDay(v string) *AddBlacklistRequest {
	s.ExpiredDay = &v
	return s
}

func (s *AddBlacklistRequest) SetNumbers(v []*string) *AddBlacklistRequest {
	s.Numbers = v
	return s
}

func (s *AddBlacklistRequest) SetOwnerId(v int64) *AddBlacklistRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBlacklistRequest) SetRemark(v string) *AddBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *AddBlacklistRequest) SetResourceOwnerAccount(v string) *AddBlacklistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBlacklistRequest) SetResourceOwnerId(v int64) *AddBlacklistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
