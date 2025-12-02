// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApsWebhookResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApsWebhookResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListApsWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApsWebhookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApsWebhookResponseBody
	GetSuccess() *bool
	SetWebhook(v []*ListApsWebhookResponseBodyWebhook) *ListApsWebhookResponseBody
	GetWebhook() []*ListApsWebhookResponseBodyWebhook
}

type ListApsWebhookResponseBody struct {
	// API status or POP error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// exampleRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The array of webhooks.
	Webhook []*ListApsWebhookResponseBodyWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Repeated"`
}

func (s ListApsWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApsWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *ListApsWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApsWebhookResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApsWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApsWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApsWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApsWebhookResponseBody) GetWebhook() []*ListApsWebhookResponseBodyWebhook {
	return s.Webhook
}

func (s *ListApsWebhookResponseBody) SetCode(v string) *ListApsWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *ListApsWebhookResponseBody) SetHttpStatusCode(v int32) *ListApsWebhookResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApsWebhookResponseBody) SetMessage(v string) *ListApsWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *ListApsWebhookResponseBody) SetRequestId(v string) *ListApsWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApsWebhookResponseBody) SetSuccess(v bool) *ListApsWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *ListApsWebhookResponseBody) SetWebhook(v []*ListApsWebhookResponseBodyWebhook) *ListApsWebhookResponseBody {
	s.Webhook = v
	return s
}

func (s *ListApsWebhookResponseBody) Validate() error {
	if s.Webhook != nil {
		for _, item := range s.Webhook {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApsWebhookResponseBodyWebhook struct {
	// Signing key
	//
	// example:
	//
	// your_secret_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the webhook.
	//
	// example:
	//
	// webhook_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request URL.
	//
	// example:
	//
	// https://example.com/webhook
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The ID of the webhook that you want to delete.
	//
	// example:
	//
	// **35***
	WebhookId *string `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	// Webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s ListApsWebhookResponseBodyWebhook) String() string {
	return dara.Prettify(s)
}

func (s ListApsWebhookResponseBodyWebhook) GoString() string {
	return s.String()
}

func (s *ListApsWebhookResponseBodyWebhook) GetKey() *string {
	return s.Key
}

func (s *ListApsWebhookResponseBodyWebhook) GetName() *string {
	return s.Name
}

func (s *ListApsWebhookResponseBodyWebhook) GetUrl() *string {
	return s.Url
}

func (s *ListApsWebhookResponseBodyWebhook) GetWebhookId() *string {
	return s.WebhookId
}

func (s *ListApsWebhookResponseBodyWebhook) GetWebhookType() *string {
	return s.WebhookType
}

func (s *ListApsWebhookResponseBodyWebhook) SetKey(v string) *ListApsWebhookResponseBodyWebhook {
	s.Key = &v
	return s
}

func (s *ListApsWebhookResponseBodyWebhook) SetName(v string) *ListApsWebhookResponseBodyWebhook {
	s.Name = &v
	return s
}

func (s *ListApsWebhookResponseBodyWebhook) SetUrl(v string) *ListApsWebhookResponseBodyWebhook {
	s.Url = &v
	return s
}

func (s *ListApsWebhookResponseBodyWebhook) SetWebhookId(v string) *ListApsWebhookResponseBodyWebhook {
	s.WebhookId = &v
	return s
}

func (s *ListApsWebhookResponseBodyWebhook) SetWebhookType(v string) *ListApsWebhookResponseBodyWebhook {
	s.WebhookType = &v
	return s
}

func (s *ListApsWebhookResponseBodyWebhook) Validate() error {
	return dara.Validate(s)
}
