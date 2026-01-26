// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOrUpdateWebhookContactResponseBody
	GetRequestId() *string
	SetWebhookContact(v *CreateOrUpdateWebhookContactResponseBodyWebhookContact) *CreateOrUpdateWebhookContactResponseBody
	GetWebhookContact() *CreateOrUpdateWebhookContactResponseBodyWebhookContact
}

type CreateOrUpdateWebhookContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned webhook alert contact.
	WebhookContact *CreateOrUpdateWebhookContactResponseBodyWebhookContact `json:"WebhookContact,omitempty" xml:"WebhookContact,omitempty" type:"Struct"`
}

func (s CreateOrUpdateWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateWebhookContactResponseBody) GetWebhookContact() *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	return s.WebhookContact
}

func (s *CreateOrUpdateWebhookContactResponseBody) SetRequestId(v string) *CreateOrUpdateWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBody) SetWebhookContact(v *CreateOrUpdateWebhookContactResponseBodyWebhookContact) *CreateOrUpdateWebhookContactResponseBody {
	s.WebhookContact = v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBody) Validate() error {
	if s.WebhookContact != nil {
		if err := s.WebhookContact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateWebhookContactResponseBodyWebhookContact struct {
	// The information about the webhook alert contact.
	Webhook *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Struct"`
	// The ID of the webhook alert contact.
	//
	// example:
	//
	// 123
	WebhookId *float32 `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	// The name of the webhook alert contact.
	//
	// example:
	//
	// Webhook alert
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContact) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContact) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) GetWebhook() *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	return s.Webhook
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) GetWebhookId() *float32 {
	return s.WebhookId
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) GetWebhookName() *string {
	return s.WebhookName
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) SetWebhook(v *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	s.Webhook = v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) SetWebhookId(v float32) *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	s.WebhookId = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) SetWebhookName(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	s.WebhookName = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) Validate() error {
	if s.Webhook != nil {
		if err := s.Webhook.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook struct {
	// The HTTP request headers.
	//
	// example:
	//
	// [{"Content-Type":"application/json;charset=utf-8"}]
	BizHeaders *string `json:"BizHeaders,omitempty" xml:"BizHeaders,omitempty"`
	// The parameters in the HTTP request.
	//
	// example:
	//
	// [{"content":"mike"}]
	BizParams *string `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	// The alert notification template.
	//
	// example:
	//
	// { "Alert name":"{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster name":"{{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application name":"{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification policy":"{{ .dispatchRuleName }}", "Alert time":"{{ .startTime }}", "Alert content":"{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The HTTP request method.
	//
	// 	- Post
	//
	// 	- Get
	//
	// example:
	//
	// Post
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The notification template for clearing alerts.
	//
	// example:
	//
	// { "Alert name":"{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster name":"{{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application name":"{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification policy":"{{ .dispatchRuleName }}", "Recovery time":"{{ .endTime }}", "Alert content":"{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	// The URL of the request method.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121******
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GetBizHeaders() *string {
	return s.BizHeaders
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GetBizParams() *string {
	return s.BizParams
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GetBody() *string {
	return s.Body
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GetMethod() *string {
	return s.Method
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GetRecoverBody() *string {
	return s.RecoverBody
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GetUrl() *string {
	return s.Url
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetBizHeaders(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.BizHeaders = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetBizParams(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.BizParams = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetBody(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.Body = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetMethod(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.Method = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetRecoverBody(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.RecoverBody = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetUrl(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.Url = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) Validate() error {
	return dara.Validate(s)
}
