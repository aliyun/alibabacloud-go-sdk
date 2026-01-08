// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *AddContactsRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *AddContactsRequest
	GetBizExtend() map[string]interface{}
	SetContactDetails(v string) *AddContactsRequest
	GetContactDetails() *string
	SetContactName(v string) *AddContactsRequest
	GetContactName() *string
	SetCountry(v string) *AddContactsRequest
	GetCountry() *string
	SetEmail(v string) *AddContactsRequest
	GetEmail() *string
	SetFilePath(v string) *AddContactsRequest
	GetFilePath() *string
	SetGroups(v string) *AddContactsRequest
	GetGroups() *string
	SetNeedUpdate(v bool) *AddContactsRequest
	GetNeedUpdate() *bool
	SetOwnerId(v int64) *AddContactsRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddContactsRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddContactsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddContactsRequest
	GetResourceOwnerId() *int64
}

type AddContactsRequest struct {
	// example:
	//
	// 示例值示例值示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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

func (s AddContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContactsRequest) GoString() string {
	return s.String()
}

func (s *AddContactsRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *AddContactsRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *AddContactsRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *AddContactsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *AddContactsRequest) GetCountry() *string {
	return s.Country
}

func (s *AddContactsRequest) GetEmail() *string {
	return s.Email
}

func (s *AddContactsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *AddContactsRequest) GetGroups() *string {
	return s.Groups
}

func (s *AddContactsRequest) GetNeedUpdate() *bool {
	return s.NeedUpdate
}

func (s *AddContactsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddContactsRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddContactsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddContactsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddContactsRequest) SetBizCode(v string) *AddContactsRequest {
	s.BizCode = &v
	return s
}

func (s *AddContactsRequest) SetBizExtend(v map[string]interface{}) *AddContactsRequest {
	s.BizExtend = v
	return s
}

func (s *AddContactsRequest) SetContactDetails(v string) *AddContactsRequest {
	s.ContactDetails = &v
	return s
}

func (s *AddContactsRequest) SetContactName(v string) *AddContactsRequest {
	s.ContactName = &v
	return s
}

func (s *AddContactsRequest) SetCountry(v string) *AddContactsRequest {
	s.Country = &v
	return s
}

func (s *AddContactsRequest) SetEmail(v string) *AddContactsRequest {
	s.Email = &v
	return s
}

func (s *AddContactsRequest) SetFilePath(v string) *AddContactsRequest {
	s.FilePath = &v
	return s
}

func (s *AddContactsRequest) SetGroups(v string) *AddContactsRequest {
	s.Groups = &v
	return s
}

func (s *AddContactsRequest) SetNeedUpdate(v bool) *AddContactsRequest {
	s.NeedUpdate = &v
	return s
}

func (s *AddContactsRequest) SetOwnerId(v int64) *AddContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddContactsRequest) SetRemark(v string) *AddContactsRequest {
	s.Remark = &v
	return s
}

func (s *AddContactsRequest) SetResourceOwnerAccount(v string) *AddContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddContactsRequest) SetResourceOwnerId(v int64) *AddContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddContactsRequest) Validate() error {
	return dara.Validate(s)
}
