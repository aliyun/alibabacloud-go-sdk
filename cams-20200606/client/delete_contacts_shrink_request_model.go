// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *DeleteContactsShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *DeleteContactsShrinkRequest
	GetBizExtendShrink() *string
	SetContactDetails(v string) *DeleteContactsShrinkRequest
	GetContactDetails() *string
	SetContactName(v string) *DeleteContactsShrinkRequest
	GetContactName() *string
	SetCountry(v string) *DeleteContactsShrinkRequest
	GetCountry() *string
	SetFilePath(v string) *DeleteContactsShrinkRequest
	GetFilePath() *string
	SetOwnerId(v int64) *DeleteContactsShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteContactsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteContactsShrinkRequest
	GetResourceOwnerId() *int64
}

type DeleteContactsShrinkRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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

func (s DeleteContactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *DeleteContactsShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *DeleteContactsShrinkRequest) GetContactDetails() *string {
	return s.ContactDetails
}

func (s *DeleteContactsShrinkRequest) GetContactName() *string {
	return s.ContactName
}

func (s *DeleteContactsShrinkRequest) GetCountry() *string {
	return s.Country
}

func (s *DeleteContactsShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteContactsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteContactsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteContactsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteContactsShrinkRequest) SetBizCode(v string) *DeleteContactsShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetBizExtendShrink(v string) *DeleteContactsShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetContactDetails(v string) *DeleteContactsShrinkRequest {
	s.ContactDetails = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetContactName(v string) *DeleteContactsShrinkRequest {
	s.ContactName = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetCountry(v string) *DeleteContactsShrinkRequest {
	s.Country = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetFilePath(v string) *DeleteContactsShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetOwnerId(v int64) *DeleteContactsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetResourceOwnerAccount(v string) *DeleteContactsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContactsShrinkRequest) SetResourceOwnerId(v int64) *DeleteContactsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteContactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
