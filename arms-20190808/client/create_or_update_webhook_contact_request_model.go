// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizHeaders(v string) *CreateOrUpdateWebhookContactRequest
	GetBizHeaders() *string
	SetBizParams(v string) *CreateOrUpdateWebhookContactRequest
	GetBizParams() *string
	SetBody(v string) *CreateOrUpdateWebhookContactRequest
	GetBody() *string
	SetMethod(v string) *CreateOrUpdateWebhookContactRequest
	GetMethod() *string
	SetRecoverBody(v string) *CreateOrUpdateWebhookContactRequest
	GetRecoverBody() *string
	SetUrl(v string) *CreateOrUpdateWebhookContactRequest
	GetUrl() *string
	SetWebhookId(v int64) *CreateOrUpdateWebhookContactRequest
	GetWebhookId() *int64
	SetWebhookName(v string) *CreateOrUpdateWebhookContactRequest
	GetWebhookName() *string
}

type CreateOrUpdateWebhookContactRequest struct {
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
	// The notification template that is sent when an alert is triggered. This parameter is required if the **Method*	- parameter is set to **Post**. You can use the `$content` placeholder to specify the notification content. The content cannot exceed 500 characters in length. For more information, see [Variable description of a notification template](https://help.aliyun.com/document_detail/251834.html).\\\\
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
	// This parameter is required.
	//
	// example:
	//
	// Post
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The notification template that is sent when an alert is resolved. This parameter is required if the **Method*	- parameter is set to **Post**. You can use the `$content` placeholder to specify the notification content. The content cannot exceed 500 characters in length. For more information, see [Variable description of a notification template](https://help.aliyun.com/document_detail/251834.html).
	//
	// example:
	//
	// { "Alert name":"{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster name":"{{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application name":"{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification policy":"{{ .dispatchRuleName }}", "Recovery time":"{{ .endTime }}", "Alert content":"{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	// The URL of the HTTP request **method**.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121******
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The ID of the webhook alert contact.
	//
	// 	- If you do not specify this parameter, a new webhook alert contact is created.
	//
	// 	- If you specify this parameter, the specified webhook alert contact is modified.
	//
	// example:
	//
	// 123
	WebhookId *int64 `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	// The name of the webhook alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// Webhook alert
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s CreateOrUpdateWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactRequest) GetBizHeaders() *string {
	return s.BizHeaders
}

func (s *CreateOrUpdateWebhookContactRequest) GetBizParams() *string {
	return s.BizParams
}

func (s *CreateOrUpdateWebhookContactRequest) GetBody() *string {
	return s.Body
}

func (s *CreateOrUpdateWebhookContactRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateOrUpdateWebhookContactRequest) GetRecoverBody() *string {
	return s.RecoverBody
}

func (s *CreateOrUpdateWebhookContactRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateOrUpdateWebhookContactRequest) GetWebhookId() *int64 {
	return s.WebhookId
}

func (s *CreateOrUpdateWebhookContactRequest) GetWebhookName() *string {
	return s.WebhookName
}

func (s *CreateOrUpdateWebhookContactRequest) SetBizHeaders(v string) *CreateOrUpdateWebhookContactRequest {
	s.BizHeaders = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetBizParams(v string) *CreateOrUpdateWebhookContactRequest {
	s.BizParams = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetBody(v string) *CreateOrUpdateWebhookContactRequest {
	s.Body = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetMethod(v string) *CreateOrUpdateWebhookContactRequest {
	s.Method = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetRecoverBody(v string) *CreateOrUpdateWebhookContactRequest {
	s.RecoverBody = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetUrl(v string) *CreateOrUpdateWebhookContactRequest {
	s.Url = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetWebhookId(v int64) *CreateOrUpdateWebhookContactRequest {
	s.WebhookId = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetWebhookName(v string) *CreateOrUpdateWebhookContactRequest {
	s.WebhookName = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
