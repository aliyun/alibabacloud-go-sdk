// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetSmsTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *GetSmsTemplateRequest
	GetTemplateCode() *string
}

type GetSmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// SMS template code.
	//
	// - Obtain the SMS template code from the return parameters of the [CreateSmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-createsmstemplate?spm) API.
	//
	// - View the SMS template code on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetSmsTemplateRequest) SetOwnerId(v int64) *GetSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsTemplateRequest) SetResourceOwnerAccount(v string) *GetSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsTemplateRequest) SetResourceOwnerId(v int64) *GetSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsTemplateRequest) SetTemplateCode(v string) *GetSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
