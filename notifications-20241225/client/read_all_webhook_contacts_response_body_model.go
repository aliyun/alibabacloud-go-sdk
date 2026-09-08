// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllWebhookContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadAllWebhookContactsResponseBody
	GetCode() *string
	SetData(v []*ReadAllWebhookContactsResponseBodyData) *ReadAllWebhookContactsResponseBody
	GetData() []*ReadAllWebhookContactsResponseBodyData
	SetMessage(v string) *ReadAllWebhookContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadAllWebhookContactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadAllWebhookContactsResponseBody
	GetSuccess() *bool
}

type ReadAllWebhookContactsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data []*ReadAllWebhookContactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The business message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// /
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadAllWebhookContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadAllWebhookContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ReadAllWebhookContactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadAllWebhookContactsResponseBody) GetData() []*ReadAllWebhookContactsResponseBodyData {
	return s.Data
}

func (s *ReadAllWebhookContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadAllWebhookContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadAllWebhookContactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadAllWebhookContactsResponseBody) SetCode(v string) *ReadAllWebhookContactsResponseBody {
	s.Code = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBody) SetData(v []*ReadAllWebhookContactsResponseBodyData) *ReadAllWebhookContactsResponseBody {
	s.Data = v
	return s
}

func (s *ReadAllWebhookContactsResponseBody) SetMessage(v string) *ReadAllWebhookContactsResponseBody {
	s.Message = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBody) SetRequestId(v string) *ReadAllWebhookContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBody) SetSuccess(v bool) *ReadAllWebhookContactsResponseBody {
	s.Success = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadAllWebhookContactsResponseBodyData struct {
	// The security token.
	//
	// example:
	//
	// /
	BotSecurityToken *string `json:"BotSecurityToken,omitempty" xml:"BotSecurityToken,omitempty"`
	// webhook id
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the webhook contact.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The security token (deprecated).
	//
	// example:
	//
	// /
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The bot URL.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxxx
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The template code.
	//
	// example:
	//
	// lark
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s ReadAllWebhookContactsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadAllWebhookContactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadAllWebhookContactsResponseBodyData) GetBotSecurityToken() *string {
	return s.BotSecurityToken
}

func (s *ReadAllWebhookContactsResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadAllWebhookContactsResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ReadAllWebhookContactsResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReadAllWebhookContactsResponseBodyData) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *ReadAllWebhookContactsResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ReadAllWebhookContactsResponseBodyData) GetWebhookType() *string {
	return s.WebhookType
}

func (s *ReadAllWebhookContactsResponseBodyData) SetBotSecurityToken(v string) *ReadAllWebhookContactsResponseBodyData {
	s.BotSecurityToken = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) SetContactId(v int64) *ReadAllWebhookContactsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) SetContactName(v string) *ReadAllWebhookContactsResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) SetSecurityToken(v string) *ReadAllWebhookContactsResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) SetServerUrl(v string) *ReadAllWebhookContactsResponseBodyData {
	s.ServerUrl = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) SetTemplateCode(v string) *ReadAllWebhookContactsResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) SetWebhookType(v string) *ReadAllWebhookContactsResponseBodyData {
	s.WebhookType = &v
	return s
}

func (s *ReadAllWebhookContactsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
