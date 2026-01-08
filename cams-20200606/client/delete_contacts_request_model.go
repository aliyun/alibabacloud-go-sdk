// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DeleteContactsRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *DeleteContactsRequest
	GetBizExtend() map[string]interface{}
	SetContactDetails(v string) *DeleteContactsRequest
	GetContactDetails() *string
	SetContactName(v string) *DeleteContactsRequest
	GetContactName() *string
	SetCountry(v string) *DeleteContactsRequest
	GetCountry() *string
	SetFilePath(v string) *DeleteContactsRequest
	GetFilePath() *string
	SetOwnerId(v int64) *DeleteContactsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteContactsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteContactsRequest
	GetResourceOwnerId() *int64
}

type DeleteContactsRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// 1507486****
	ContactDetails *string `json:"ContactDetails,omitempty" xml:"ContactDetails,omitempty"`
	// example:
	//
	// hahaha
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// China
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// http://****
	FilePath             *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DeleteContactsRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *DeleteContactsRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *DeleteContactsRequest) GetContactName() *string {
	return s.ContactName
}

func (s *DeleteContactsRequest) GetCountry() *string {
	return s.Country
}

func (s *DeleteContactsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteContactsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteContactsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteContactsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteContactsRequest) SetBizCode(v string) *DeleteContactsRequest {
	s.BizCode = &v
	return s
}

func (s *DeleteContactsRequest) SetBizExtend(v map[string]interface{}) *DeleteContactsRequest {
	s.BizExtend = v
	return s
}

func (s *DeleteContactsRequest) SetContactDetails(v string) *DeleteContactsRequest {
	s.ContactDetails = &v
	return s
}

func (s *DeleteContactsRequest) SetContactName(v string) *DeleteContactsRequest {
	s.ContactName = &v
	return s
}

func (s *DeleteContactsRequest) SetCountry(v string) *DeleteContactsRequest {
	s.Country = &v
	return s
}

func (s *DeleteContactsRequest) SetFilePath(v string) *DeleteContactsRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteContactsRequest) SetOwnerId(v int64) *DeleteContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContactsRequest) SetResourceOwnerAccount(v string) *DeleteContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContactsRequest) SetResourceOwnerId(v int64) *DeleteContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteContactsRequest) Validate() error {
	return dara.Validate(s)
}
