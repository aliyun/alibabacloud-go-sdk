// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySmsTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *QuerySmsTemplateRequest
	GetTemplateCode() *string
}

type QuerySmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the template code on the **Templates*	- tab. You can also call the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation to obtain the template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_1525***
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QuerySmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *QuerySmsTemplateRequest) SetOwnerId(v int64) *QuerySmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsTemplateRequest) SetResourceOwnerAccount(v string) *QuerySmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsTemplateRequest) SetResourceOwnerId(v int64) *QuerySmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsTemplateRequest) SetTemplateCode(v string) *QuerySmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *QuerySmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
