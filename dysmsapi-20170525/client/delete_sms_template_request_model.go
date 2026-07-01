// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteSmsTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *DeleteSmsTemplateRequest
	GetTemplateCode() *string
}

type DeleteSmsTemplateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The SMS template code. You can delete SMS templates that are recalled, rejected, or approved. You cannot delete SMS templates that are being reviewed.
	//
	// - You can call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) operation to obtain the SMS template code.
	//
	// - You can also obtain the SMS template code on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page of the Short Message Service (SMS) console.
	//
	// 	Notice: Deleted SMS templates cannot be recovered and cannot be used to send messages. Proceed with caution..
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DeleteSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSmsTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DeleteSmsTemplateRequest) SetOwnerId(v int64) *DeleteSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetResourceOwnerAccount(v string) *DeleteSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetResourceOwnerId(v int64) *DeleteSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmsTemplateRequest) SetTemplateCode(v string) *DeleteSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *DeleteSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
