// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v string) *ListChatappTemplateShrinkRequest
	GetAuditStatus() *string
	SetCategory(v string) *ListChatappTemplateShrinkRequest
	GetCategory() *string
	SetCode(v string) *ListChatappTemplateShrinkRequest
	GetCode() *string
	SetCustSpaceId(v string) *ListChatappTemplateShrinkRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *ListChatappTemplateShrinkRequest
	GetCustWabaId() *string
	SetIsvCode(v string) *ListChatappTemplateShrinkRequest
	GetIsvCode() *string
	SetLanguage(v string) *ListChatappTemplateShrinkRequest
	GetLanguage() *string
	SetName(v string) *ListChatappTemplateShrinkRequest
	GetName() *string
	SetOwnerId(v int64) *ListChatappTemplateShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListChatappTemplateShrinkRequest
	GetPageShrink() *string
	SetResourceOwnerAccount(v string) *ListChatappTemplateShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatappTemplateShrinkRequest
	GetResourceOwnerId() *int64
	SetTemplateType(v string) *ListChatappTemplateShrinkRequest
	GetTemplateType() *string
}

type ListChatappTemplateShrinkRequest struct {
	// The review state of the template. Valid values:
	//
	// 	- **pass**: The template is approved.
	//
	// 	- **fail**: The template is rejected.
	//
	// 	- **auditing**: The template is being reviewed.
	//
	// 	- **unaudit**: The review is suspended.
	//
	// example:
	//
	// pass
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The category of the message template.
	//
	// example:
	//
	// AUTHENTICATION
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The code of the message template.
	//
	// example:
	//
	// 838888822*****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The space ID of the RAM user within the ISV account.
	//
	// example:
	//
	// 28251486512358****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// Deprecated
	//
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
	// The language that is used in the message template. For more information, see [Language codes](https://help.aliyun.com/document_detail/463420.html).
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// hello_whatsapp
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pagination settings.
	//
	// example:
	//
	// "page": "{\\"index\\": 1,\\"size\\": 20}
	PageShrink           *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the message template.
	//
	// 	- **WHATSAPP**
	//
	// 	- **VIBER**
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListChatappTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatappTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateShrinkRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *ListChatappTemplateShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ListChatappTemplateShrinkRequest) GetCode() *string {
	return s.Code
}

func (s *ListChatappTemplateShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatappTemplateShrinkRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *ListChatappTemplateShrinkRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *ListChatappTemplateShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListChatappTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListChatappTemplateShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatappTemplateShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListChatappTemplateShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatappTemplateShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatappTemplateShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListChatappTemplateShrinkRequest) SetAuditStatus(v string) *ListChatappTemplateShrinkRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCategory(v string) *ListChatappTemplateShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCode(v string) *ListChatappTemplateShrinkRequest {
	s.Code = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCustSpaceId(v string) *ListChatappTemplateShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetCustWabaId(v string) *ListChatappTemplateShrinkRequest {
	s.CustWabaId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetIsvCode(v string) *ListChatappTemplateShrinkRequest {
	s.IsvCode = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetLanguage(v string) *ListChatappTemplateShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetName(v string) *ListChatappTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetOwnerId(v int64) *ListChatappTemplateShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetPageShrink(v string) *ListChatappTemplateShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetResourceOwnerAccount(v string) *ListChatappTemplateShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetResourceOwnerId(v int64) *ListChatappTemplateShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) SetTemplateType(v string) *ListChatappTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *ListChatappTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
