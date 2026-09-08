// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadWebhookContactResponseBody
	GetCode() *string
	SetData(v *ReadWebhookContactResponseBodyData) *ReadWebhookContactResponseBody
	GetData() *ReadWebhookContactResponseBodyData
	SetMessage(v string) *ReadWebhookContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadWebhookContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadWebhookContactResponseBody
	GetSuccess() *bool
}

type ReadWebhookContactResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *ReadWebhookContactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The business message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
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

func (s ReadWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadWebhookContactResponseBody) GetData() *ReadWebhookContactResponseBodyData {
	return s.Data
}

func (s *ReadWebhookContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadWebhookContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadWebhookContactResponseBody) SetCode(v string) *ReadWebhookContactResponseBody {
	s.Code = &v
	return s
}

func (s *ReadWebhookContactResponseBody) SetData(v *ReadWebhookContactResponseBodyData) *ReadWebhookContactResponseBody {
	s.Data = v
	return s
}

func (s *ReadWebhookContactResponseBody) SetMessage(v string) *ReadWebhookContactResponseBody {
	s.Message = &v
	return s
}

func (s *ReadWebhookContactResponseBody) SetRequestId(v string) *ReadWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadWebhookContactResponseBody) SetSuccess(v bool) *ReadWebhookContactResponseBody {
	s.Success = &v
	return s
}

func (s *ReadWebhookContactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadWebhookContactResponseBodyData struct {
	// The security token.
	//
	// example:
	//
	// ***
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
	// Deprecated
	//
	// The security token (deprecated).
	//
	// example:
	//
	// ***
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The bot server URL.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxxxxx
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s ReadWebhookContactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactResponseBodyData) GetBotSecurityToken() *string {
	return s.BotSecurityToken
}

func (s *ReadWebhookContactResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadWebhookContactResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ReadWebhookContactResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReadWebhookContactResponseBodyData) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *ReadWebhookContactResponseBodyData) GetWebhookType() *string {
	return s.WebhookType
}

func (s *ReadWebhookContactResponseBodyData) SetBotSecurityToken(v string) *ReadWebhookContactResponseBodyData {
	s.BotSecurityToken = &v
	return s
}

func (s *ReadWebhookContactResponseBodyData) SetContactId(v int64) *ReadWebhookContactResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ReadWebhookContactResponseBodyData) SetContactName(v string) *ReadWebhookContactResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ReadWebhookContactResponseBodyData) SetSecurityToken(v string) *ReadWebhookContactResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *ReadWebhookContactResponseBodyData) SetServerUrl(v string) *ReadWebhookContactResponseBodyData {
	s.ServerUrl = &v
	return s
}

func (s *ReadWebhookContactResponseBodyData) SetWebhookType(v string) *ReadWebhookContactResponseBodyData {
	s.WebhookType = &v
	return s
}

func (s *ReadWebhookContactResponseBodyData) Validate() error {
	return dara.Validate(s)
}
