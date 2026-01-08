// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *AddContactsShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *AddContactsShrinkRequest
	GetBizExtendShrink() *string
	SetContactDetails(v string) *AddContactsShrinkRequest
	GetContactDetails() *string
	SetContactName(v string) *AddContactsShrinkRequest
	GetContactName() *string
	SetCountry(v string) *AddContactsShrinkRequest
	GetCountry() *string
	SetEmail(v string) *AddContactsShrinkRequest
	GetEmail() *string
	SetFilePath(v string) *AddContactsShrinkRequest
	GetFilePath() *string
	SetGroups(v string) *AddContactsShrinkRequest
	GetGroups() *string
	SetNeedUpdate(v bool) *AddContactsShrinkRequest
	GetNeedUpdate() *bool
	SetOwnerId(v int64) *AddContactsShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddContactsShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddContactsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddContactsShrinkRequest
	GetResourceOwnerId() *int64
}

type AddContactsShrinkRequest struct {
	// example:
	//
	// 示例值示例值示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// 15111111111
	ContactDetails *string `json:"ContactDetails,omitempty" xml:"ContactDetails,omitempty"`
	// example:
	//
	// mary
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 示例值示例值
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Groups *string `json:"Groups,omitempty" xml:"Groups,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedUpdate *bool  `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// test
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddContactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddContactsShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *AddContactsShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *AddContactsShrinkRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *AddContactsShrinkRequest) GetContactName() *string {
	return s.ContactName
}

func (s *AddContactsShrinkRequest) GetCountry() *string {
	return s.Country
}

func (s *AddContactsShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *AddContactsShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *AddContactsShrinkRequest) GetGroups() *string {
	return s.Groups
}

func (s *AddContactsShrinkRequest) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *AddContactsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddContactsShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddContactsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddContactsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddContactsShrinkRequest) SetBizCode(v string) *AddContactsShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *AddContactsShrinkRequest) SetBizExtendShrink(v string) *AddContactsShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *AddContactsShrinkRequest) SetContactDetails(v string) *AddContactsShrinkRequest {
	s.ContactDetails = &v
	return s
}

func (s *AddContactsShrinkRequest) SetContactName(v string) *AddContactsShrinkRequest {
	s.ContactName = &v
	return s
}

func (s *AddContactsShrinkRequest) SetCountry(v string) *AddContactsShrinkRequest {
	s.Country = &v
	return s
}

func (s *AddContactsShrinkRequest) SetEmail(v string) *AddContactsShrinkRequest {
	s.Email = &v
	return s
}

func (s *AddContactsShrinkRequest) SetFilePath(v string) *AddContactsShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *AddContactsShrinkRequest) SetGroups(v string) *AddContactsShrinkRequest {
	s.Groups = &v
	return s
}

func (s *AddContactsShrinkRequest) SetNeedUpdate(v bool) *AddContactsShrinkRequest {
	s.NeedUpdate = &v
	return s
}

func (s *AddContactsShrinkRequest) SetOwnerId(v int64) *AddContactsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddContactsShrinkRequest) SetRemark(v string) *AddContactsShrinkRequest {
	s.Remark = &v
	return s
}

func (s *AddContactsShrinkRequest) SetResourceOwnerAccount(v string) *AddContactsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddContactsShrinkRequest) SetResourceOwnerId(v int64) *AddContactsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddContactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
