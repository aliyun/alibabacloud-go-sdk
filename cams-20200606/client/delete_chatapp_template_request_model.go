// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatappTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *DeleteChatappTemplateRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *DeleteChatappTemplateRequest
	GetCustWabaId() *string
	SetIsvCode(v string) *DeleteChatappTemplateRequest
	GetIsvCode() *string
	SetLanguage(v string) *DeleteChatappTemplateRequest
	GetLanguage() *string
	SetOwnerId(v int64) *DeleteChatappTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteChatappTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteChatappTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *DeleteChatappTemplateRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *DeleteChatappTemplateRequest
	GetTemplateName() *string
	SetTemplateType(v string) *DeleteChatappTemplateRequest
	GetTemplateType() *string
}

type DeleteChatappTemplateRequest struct {
	// The space ID of the RAM user within the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The WhatsApp Business Account (WABA) ID of the RAM user within the independent software vendor (ISV) account.
	//
	// >  CustWabaId is an obsolete parameter. Use CustSpaceId instead.
	//
	// example:
	//
	// 65921621816****
	CustWabaId *string `json:"CustWabaId,omitempty" xml:"CustWabaId,omitempty"`
	// The ISV verification code. This parameter is used to verify whether the RAM user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The template language.
	//
	// example:
	//
	// zh_CN
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template code.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// test_name
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The template type. This parameter is required if you delete a template in a language.
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DeleteChatappTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatappTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteChatappTemplateRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *DeleteChatappTemplateRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *DeleteChatappTemplateRequest) GetLanguage() *string {
	return s.Language
}

func (s *DeleteChatappTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteChatappTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteChatappTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteChatappTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DeleteChatappTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DeleteChatappTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DeleteChatappTemplateRequest) SetCustSpaceId(v string) *DeleteChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetCustWabaId(v string) *DeleteChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetIsvCode(v string) *DeleteChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetLanguage(v string) *DeleteChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetOwnerId(v int64) *DeleteChatappTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetResourceOwnerAccount(v string) *DeleteChatappTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetResourceOwnerId(v int64) *DeleteChatappTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateCode(v string) *DeleteChatappTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateName(v string) *DeleteChatappTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteChatappTemplateRequest) SetTemplateType(v string) *DeleteChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *DeleteChatappTemplateRequest) Validate() error {
	return dara.Validate(s)
}
