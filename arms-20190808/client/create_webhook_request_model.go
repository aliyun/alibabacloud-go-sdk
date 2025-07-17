// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateWebhookRequest
	GetBody() *string
	SetContactName(v string) *CreateWebhookRequest
	GetContactName() *string
	SetHttpHeaders(v string) *CreateWebhookRequest
	GetHttpHeaders() *string
	SetHttpParams(v string) *CreateWebhookRequest
	GetHttpParams() *string
	SetMethod(v string) *CreateWebhookRequest
	GetMethod() *string
	SetRecoverBody(v string) *CreateWebhookRequest
	GetRecoverBody() *string
	SetRegionId(v string) *CreateWebhookRequest
	GetRegionId() *string
	SetUrl(v string) *CreateWebhookRequest
	GetUrl() *string
}

type CreateWebhookRequest struct {
	// The notification template that is sent when an alert is triggered. This parameter is required if the **Method*	- parameter is set to **Post**. You can use the $content placeholder to specify the notification content. The content cannot exceed 500 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "Alert Name": "{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster Name": "{{ .commonLabels.clustername }} {{ end }}{{if eq " app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application Name": "{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification Policy": "{{ .dispatchRuleName }}", "Alarm Time": "{{ .startTime }}", "Alert Content": "{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The name of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebhookAlert
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The HTTP request header.
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
	// The HTTP request method.
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
	// { "Alert Name": "{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster Name": "{{ .commonLabels.clustername }} {{ end }}{{if eq " app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application Name": "{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification Policy": "{{ .dispatchRuleName }}", "Alarm Time": "{{ .startTime }}", "Alert Content": "{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The URL of the request **method**.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=e1a049121ddbfce1ca963d115ef88cc7219583c4fb79fe6e398fbfb688******
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebhookRequest) GoString() string {
	return s.String()
}

func (s *CreateWebhookRequest) GetBody() *string {
	return s.Body
}

func (s *CreateWebhookRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateWebhookRequest) GetHttpHeaders() *string {
	return s.HttpHeaders
}

func (s *CreateWebhookRequest) GetHttpParams() *string {
	return s.HttpParams
}

func (s *CreateWebhookRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateWebhookRequest) GetRecoverBody() *string {
	return s.RecoverBody
}

func (s *CreateWebhookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWebhookRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateWebhookRequest) SetBody(v string) *CreateWebhookRequest {
	s.Body = &v
	return s
}

func (s *CreateWebhookRequest) SetContactName(v string) *CreateWebhookRequest {
	s.ContactName = &v
	return s
}

func (s *CreateWebhookRequest) SetHttpHeaders(v string) *CreateWebhookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *CreateWebhookRequest) SetHttpParams(v string) *CreateWebhookRequest {
	s.HttpParams = &v
	return s
}

func (s *CreateWebhookRequest) SetMethod(v string) *CreateWebhookRequest {
	s.Method = &v
	return s
}

func (s *CreateWebhookRequest) SetRecoverBody(v string) *CreateWebhookRequest {
	s.RecoverBody = &v
	return s
}

func (s *CreateWebhookRequest) SetRegionId(v string) *CreateWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWebhookRequest) SetUrl(v string) *CreateWebhookRequest {
	s.Url = &v
	return s
}

func (s *CreateWebhookRequest) Validate() error {
	return dara.Validate(s)
}
