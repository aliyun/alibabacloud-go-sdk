// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *UpdateWebhookRequest
	GetBody() *string
	SetContactId(v int64) *UpdateWebhookRequest
	GetContactId() *int64
	SetContactName(v string) *UpdateWebhookRequest
	GetContactName() *string
	SetHttpHeaders(v string) *UpdateWebhookRequest
	GetHttpHeaders() *string
	SetHttpParams(v string) *UpdateWebhookRequest
	GetHttpParams() *string
	SetMethod(v string) *UpdateWebhookRequest
	GetMethod() *string
	SetRecoverBody(v string) *UpdateWebhookRequest
	GetRecoverBody() *string
	SetRegionId(v string) *UpdateWebhookRequest
	GetRegionId() *string
	SetUrl(v string) *UpdateWebhookRequest
	GetUrl() *string
}

type UpdateWebhookRequest struct {
	// The notification template that is sent when an alert is triggered. This parameter is required if the **Method*	- parameter is set to **Post**. You can use the $content placeholder to specify the notification content. The content cannot exceed 500 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "Alert name":"{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster name":"{{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application name":"{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification policy":"{{ .dispatchRuleName }}", "Alert time":"{{ .startTime }}", "Alert content":"{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The ID of the webhook alert contact. You can call the **SearchAlertContact*	- operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 48716
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the webhook alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// Webhook alert
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The HTTP request headers.
	//
	// example:
	//
	// [{"Content-Type":"application/json"}]
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	// The parameters in the HTTP request.
	//
	// example:
	//
	// [{"name":"mike"}]
	HttpParams *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- `Get`
	//
	// 	- `Post`
	//
	// This parameter is required.
	//
	// example:
	//
	// Post
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The notification template that is sent when an alert is resolved. This parameter is required if the **Method*	- parameter is set to **Post**. You can use the $content placeholder to specify the notification content. The content cannot exceed 500 characters in length.
	//
	// example:
	//
	// { "Alert name":"{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster name":"{{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application name":"{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification policy":"{{ .dispatchRuleName }}", "Recovery time":"{{ .endTime }}", "Alert content":"{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The URL of the HTTP request method.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121ddbfce1ca963d115ef88cc7219583c4fb79fe6e398fbfb688******
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebhookRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateWebhookRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateWebhookRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateWebhookRequest) GetHttpHeaders() *string {
	return s.HttpHeaders
}

func (s *UpdateWebhookRequest) GetHttpParams() *string {
	return s.HttpParams
}

func (s *UpdateWebhookRequest) GetMethod() *string {
	return s.Method
}

func (s *UpdateWebhookRequest) GetRecoverBody() *string {
	return s.RecoverBody
}

func (s *UpdateWebhookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWebhookRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateWebhookRequest) SetBody(v string) *UpdateWebhookRequest {
	s.Body = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactId(v int64) *UpdateWebhookRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactName(v string) *UpdateWebhookRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpHeaders(v string) *UpdateWebhookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpParams(v string) *UpdateWebhookRequest {
	s.HttpParams = &v
	return s
}

func (s *UpdateWebhookRequest) SetMethod(v string) *UpdateWebhookRequest {
	s.Method = &v
	return s
}

func (s *UpdateWebhookRequest) SetRecoverBody(v string) *UpdateWebhookRequest {
	s.RecoverBody = &v
	return s
}

func (s *UpdateWebhookRequest) SetRegionId(v string) *UpdateWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWebhookRequest) SetUrl(v string) *UpdateWebhookRequest {
	s.Url = &v
	return s
}

func (s *UpdateWebhookRequest) Validate() error {
	return dara.Validate(s)
}
