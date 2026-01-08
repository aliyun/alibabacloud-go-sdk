// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *UpdateContactByIdRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *UpdateContactByIdRequest
	GetBizExtend() map[string]interface{}
	SetContactDetails(v string) *UpdateContactByIdRequest
	GetContactDetails() *string
	SetContactId(v string) *UpdateContactByIdRequest
	GetContactId() *string
	SetContactName(v string) *UpdateContactByIdRequest
	GetContactName() *string
	SetCountry(v string) *UpdateContactByIdRequest
	GetCountry() *string
	SetEmail(v string) *UpdateContactByIdRequest
	GetEmail() *string
	SetOwnerId(v int64) *UpdateContactByIdRequest
	GetOwnerId() *int64
	SetRemark(v string) *UpdateContactByIdRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateContactByIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateContactByIdRequest
	GetResourceOwnerId() *int64
}

type UpdateContactByIdRequest struct {
	// example:
	//
	// 示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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

func (s UpdateContactByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactByIdRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactByIdRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *UpdateContactByIdRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *UpdateContactByIdRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *UpdateContactByIdRequest) GetContactId() *string {
	return s.ContactId
}

func (s *UpdateContactByIdRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateContactByIdRequest) GetCountry() *string {
	return s.Country
}

func (s *UpdateContactByIdRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateContactByIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateContactByIdRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateContactByIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateContactByIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateContactByIdRequest) SetBizCode(v string) *UpdateContactByIdRequest {
	s.BizCode = &v
	return s
}

func (s *UpdateContactByIdRequest) SetBizExtend(v map[string]interface{}) *UpdateContactByIdRequest {
	s.BizExtend = v
	return s
}

func (s *UpdateContactByIdRequest) SetContactDetails(v string) *UpdateContactByIdRequest {
	s.ContactDetails = &v
	return s
}

func (s *UpdateContactByIdRequest) SetContactId(v string) *UpdateContactByIdRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateContactByIdRequest) SetContactName(v string) *UpdateContactByIdRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateContactByIdRequest) SetCountry(v string) *UpdateContactByIdRequest {
	s.Country = &v
	return s
}

func (s *UpdateContactByIdRequest) SetEmail(v string) *UpdateContactByIdRequest {
	s.Email = &v
	return s
}

func (s *UpdateContactByIdRequest) SetOwnerId(v int64) *UpdateContactByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContactByIdRequest) SetRemark(v string) *UpdateContactByIdRequest {
	s.Remark = &v
	return s
}

func (s *UpdateContactByIdRequest) SetResourceOwnerAccount(v string) *UpdateContactByIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContactByIdRequest) SetResourceOwnerId(v int64) *UpdateContactByIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateContactByIdRequest) Validate() error {
	return dara.Validate(s)
}
