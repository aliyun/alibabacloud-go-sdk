// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplatePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSend(v bool) *ModifyChatappTemplatePropertiesRequest
	GetAllowSend() *bool
	SetCategoryChangePaused(v bool) *ModifyChatappTemplatePropertiesRequest
	GetCategoryChangePaused() *bool
	SetCustSpaceId(v string) *ModifyChatappTemplatePropertiesRequest
	GetCustSpaceId() *string
	SetLanguage(v string) *ModifyChatappTemplatePropertiesRequest
	GetLanguage() *string
	SetOwnerId(v int64) *ModifyChatappTemplatePropertiesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyChatappTemplatePropertiesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyChatappTemplatePropertiesRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *ModifyChatappTemplatePropertiesRequest
	GetTemplateCode() *string
	SetTemplateType(v string) *ModifyChatappTemplatePropertiesRequest
	GetTemplateType() *string
}

type ModifyChatappTemplatePropertiesRequest struct {
	// example:
	//
	// true
	AllowSend *bool `json:"AllowSend,omitempty" xml:"AllowSend,omitempty"`
	// example:
	//
	// false
	CategoryChangePaused *bool `json:"CategoryChangePaused,omitempty" xml:"CategoryChangePaused,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-idk***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 929938***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifyChatappTemplatePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplatePropertiesRequest) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplatePropertiesRequest) GetAllowSend() *bool {
	return s.AllowSend
}

func (s *ModifyChatappTemplatePropertiesRequest) GetCategoryChangePaused() *bool {
	return s.CategoryChangePaused
}

func (s *ModifyChatappTemplatePropertiesRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyChatappTemplatePropertiesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ModifyChatappTemplatePropertiesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyChatappTemplatePropertiesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyChatappTemplatePropertiesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyChatappTemplatePropertiesRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ModifyChatappTemplatePropertiesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ModifyChatappTemplatePropertiesRequest) SetAllowSend(v bool) *ModifyChatappTemplatePropertiesRequest {
	s.AllowSend = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetCategoryChangePaused(v bool) *ModifyChatappTemplatePropertiesRequest {
	s.CategoryChangePaused = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetCustSpaceId(v string) *ModifyChatappTemplatePropertiesRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetLanguage(v string) *ModifyChatappTemplatePropertiesRequest {
	s.Language = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetOwnerId(v int64) *ModifyChatappTemplatePropertiesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetResourceOwnerAccount(v string) *ModifyChatappTemplatePropertiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetResourceOwnerId(v int64) *ModifyChatappTemplatePropertiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetTemplateCode(v string) *ModifyChatappTemplatePropertiesRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) SetTemplateType(v string) *ModifyChatappTemplatePropertiesRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
