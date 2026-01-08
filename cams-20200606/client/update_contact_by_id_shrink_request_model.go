// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactByIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *UpdateContactByIdShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *UpdateContactByIdShrinkRequest
	GetBizExtendShrink() *string
	SetContactDetails(v string) *UpdateContactByIdShrinkRequest
	GetContactDetails() *string
	SetContactId(v string) *UpdateContactByIdShrinkRequest
	GetContactId() *string
	SetContactName(v string) *UpdateContactByIdShrinkRequest
	GetContactName() *string
	SetCountry(v string) *UpdateContactByIdShrinkRequest
	GetCountry() *string
	SetEmail(v string) *UpdateContactByIdShrinkRequest
	GetEmail() *string
	SetOwnerId(v int64) *UpdateContactByIdShrinkRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateContactByIdShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateContactByIdShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateContactByIdShrinkRequest
	GetResourceOwnerId() *int64
}

type UpdateContactByIdShrinkRequest struct {
	// example:
	//
	// 示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ContactDetails *string `json:"ContactDetails,omitempty" xml:"ContactDetails,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 示例值
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// test
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateContactByIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactByIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactByIdShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateContactByIdShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *UpdateContactByIdShrinkRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *UpdateContactByIdShrinkRequest) GetContactId() *string {
	return s.ContactId
}

func (s *UpdateContactByIdShrinkRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateContactByIdShrinkRequest) GetCountry() *string {
	return s.Country
}

func (s *UpdateContactByIdShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateContactByIdShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateContactByIdShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateContactByIdShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateContactByIdShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateContactByIdShrinkRequest) SetBizCode(v string) *UpdateContactByIdShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetBizExtendShrink(v string) *UpdateContactByIdShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetContactDetails(v string) *UpdateContactByIdShrinkRequest {
	s.ContactDetails = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetContactId(v string) *UpdateContactByIdShrinkRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetContactName(v string) *UpdateContactByIdShrinkRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetCountry(v string) *UpdateContactByIdShrinkRequest {
	s.Country = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetEmail(v string) *UpdateContactByIdShrinkRequest {
	s.Email = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetOwnerId(v int64) *UpdateContactByIdShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetRemark(v string) *UpdateContactByIdShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetResourceOwnerAccount(v string) *UpdateContactByIdShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) SetResourceOwnerId(v int64) *UpdateContactByIdShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateContactByIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
