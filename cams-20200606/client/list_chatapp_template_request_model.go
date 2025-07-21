// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatappTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v string) *ListChatappTemplateRequest
	GetAuditStatus() *string
	SetCategory(v string) *ListChatappTemplateRequest
	GetCategory() *string
	SetCode(v string) *ListChatappTemplateRequest
	GetCode() *string
	SetCustSpaceId(v string) *ListChatappTemplateRequest
	GetCustSpaceId() *string
	SetCustWabaId(v string) *ListChatappTemplateRequest
	GetCustWabaId() *string
	SetIsvCode(v string) *ListChatappTemplateRequest
	GetIsvCode() *string
	SetLanguage(v string) *ListChatappTemplateRequest
	GetLanguage() *string
	SetName(v string) *ListChatappTemplateRequest
	GetName() *string
	SetOwnerId(v int64) *ListChatappTemplateRequest
	GetOwnerId() *int64
	SetPage(v *ListChatappTemplateRequestPage) *ListChatappTemplateRequest
	GetPage() *ListChatappTemplateRequestPage
	SetResourceOwnerAccount(v string) *ListChatappTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListChatappTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateType(v string) *ListChatappTemplateRequest
	GetTemplateType() *string
}

type ListChatappTemplateRequest struct {
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
	Page                 *ListChatappTemplateRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s ListChatappTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatappTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *ListChatappTemplateRequest) GetCategory() *string {
	return s.Category
}

func (s *ListChatappTemplateRequest) GetCode() *string {
	return s.Code
}

func (s *ListChatappTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListChatappTemplateRequest) GetCustWabaId() *string {
	return s.CustWabaId
}

func (s *ListChatappTemplateRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *ListChatappTemplateRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListChatappTemplateRequest) GetName() *string {
	return s.Name
}

func (s *ListChatappTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListChatappTemplateRequest) GetPage() *ListChatappTemplateRequestPage {
	return s.Page
}

func (s *ListChatappTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListChatappTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListChatappTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListChatappTemplateRequest) SetAuditStatus(v string) *ListChatappTemplateRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListChatappTemplateRequest) SetCategory(v string) *ListChatappTemplateRequest {
	s.Category = &v
	return s
}

func (s *ListChatappTemplateRequest) SetCode(v string) *ListChatappTemplateRequest {
	s.Code = &v
	return s
}

func (s *ListChatappTemplateRequest) SetCustSpaceId(v string) *ListChatappTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetCustWabaId(v string) *ListChatappTemplateRequest {
	s.CustWabaId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetIsvCode(v string) *ListChatappTemplateRequest {
	s.IsvCode = &v
	return s
}

func (s *ListChatappTemplateRequest) SetLanguage(v string) *ListChatappTemplateRequest {
	s.Language = &v
	return s
}

func (s *ListChatappTemplateRequest) SetName(v string) *ListChatappTemplateRequest {
	s.Name = &v
	return s
}

func (s *ListChatappTemplateRequest) SetOwnerId(v int64) *ListChatappTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetPage(v *ListChatappTemplateRequestPage) *ListChatappTemplateRequest {
	s.Page = v
	return s
}

func (s *ListChatappTemplateRequest) SetResourceOwnerAccount(v string) *ListChatappTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListChatappTemplateRequest) SetResourceOwnerId(v int64) *ListChatappTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListChatappTemplateRequest) SetTemplateType(v string) *ListChatappTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ListChatappTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type ListChatappTemplateRequestPage struct {
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListChatappTemplateRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListChatappTemplateRequestPage) GoString() string {
	return s.String()
}

func (s *ListChatappTemplateRequestPage) GetIndex() *int32 {
	return s.Index
}

func (s *ListChatappTemplateRequestPage) GetSize() *int32 {
	return s.Size
}

func (s *ListChatappTemplateRequestPage) SetIndex(v int32) *ListChatappTemplateRequestPage {
	s.Index = &v
	return s
}

func (s *ListChatappTemplateRequestPage) SetSize(v int32) *ListChatappTemplateRequestPage {
	s.Size = &v
	return s
}

func (s *ListChatappTemplateRequestPage) Validate() error {
	return dara.Validate(s)
}
