// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetWhatsappHealthStatusRequest
	GetCustSpaceId() *string
	SetLanguage(v string) *GetWhatsappHealthStatusRequest
	GetLanguage() *string
	SetNodeType(v string) *GetWhatsappHealthStatusRequest
	GetNodeType() *string
	SetOwnerId(v int64) *GetWhatsappHealthStatusRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetWhatsappHealthStatusRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetWhatsappHealthStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetWhatsappHealthStatusRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *GetWhatsappHealthStatusRequest
	GetTemplateCode() *string
	SetWabaId(v string) *GetWhatsappHealthStatusRequest
	GetWabaId() *string
}

type GetWhatsappHealthStatusRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account or the instance ID of the customer of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2993****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The template language.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The node type.
	//
	// Valid values:
	//
	// 	- template: message template
	//
	// 	- phone: phone number
	//
	// 	- waba: WhatsApp Business Account (WABA)
	//
	// This parameter is required.
	//
	// example:
	//
	// waba
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number of the enterprise.
	//
	// example:
	//
	// 86138***
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template code.
	//
	// example:
	//
	// 399299***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// WabaId
	//
	// example:
	//
	// 299399****
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s GetWhatsappHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetWhatsappHealthStatusRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetWhatsappHealthStatusRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetWhatsappHealthStatusRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *GetWhatsappHealthStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetWhatsappHealthStatusRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetWhatsappHealthStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetWhatsappHealthStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetWhatsappHealthStatusRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetWhatsappHealthStatusRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *GetWhatsappHealthStatusRequest) SetCustSpaceId(v string) *GetWhatsappHealthStatusRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetLanguage(v string) *GetWhatsappHealthStatusRequest {
	s.Language = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetNodeType(v string) *GetWhatsappHealthStatusRequest {
	s.NodeType = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetOwnerId(v int64) *GetWhatsappHealthStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetPhoneNumber(v string) *GetWhatsappHealthStatusRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetResourceOwnerAccount(v string) *GetWhatsappHealthStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetResourceOwnerId(v int64) *GetWhatsappHealthStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetTemplateCode(v string) *GetWhatsappHealthStatusRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) SetWabaId(v string) *GetWhatsappHealthStatusRequest {
	s.WabaId = &v
	return s
}

func (s *GetWhatsappHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}
