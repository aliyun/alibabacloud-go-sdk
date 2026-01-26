// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebhookContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *DescribeWebhookContactsResponseBodyPageBean) *DescribeWebhookContactsResponseBody
	GetPageBean() *DescribeWebhookContactsResponseBodyPageBean
	SetRequestId(v string) *DescribeWebhookContactsResponseBody
	GetRequestId() *string
}

type DescribeWebhookContactsResponseBody struct {
	// The returned objects.
	PageBean *DescribeWebhookContactsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4D6C358A-A58B-4F4B-94CE-F5AAF023****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebhookContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebhookContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBody) GetPageBean() *DescribeWebhookContactsResponseBodyPageBean {
	return s.PageBean
}

func (s *DescribeWebhookContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebhookContactsResponseBody) SetPageBean(v *DescribeWebhookContactsResponseBodyPageBean) *DescribeWebhookContactsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeWebhookContactsResponseBody) SetRequestId(v string) *DescribeWebhookContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebhookContactsResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWebhookContactsResponseBodyPageBean struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert contacts displayed on each page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of alert contacts.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The list of webhook alert contacts.
	WebhookContacts []*DescribeWebhookContactsResponseBodyPageBeanWebhookContacts `json:"WebhookContacts,omitempty" xml:"WebhookContacts,omitempty" type:"Repeated"`
}

func (s DescribeWebhookContactsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebhookContactsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *DescribeWebhookContactsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *DescribeWebhookContactsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeWebhookContactsResponseBodyPageBean) GetWebhookContacts() []*DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	return s.WebhookContacts
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetPage(v int64) *DescribeWebhookContactsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetSize(v int64) *DescribeWebhookContactsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetTotal(v int64) *DescribeWebhookContactsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetWebhookContacts(v []*DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) *DescribeWebhookContactsResponseBodyPageBean {
	s.WebhookContacts = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) Validate() error {
	if s.WebhookContacts != nil {
		for _, item := range s.WebhookContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebhookContactsResponseBodyPageBeanWebhookContacts struct {
	// The details of the webhook alert contact.
	Webhook *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Struct"`
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
	// Webhook name
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) GetWebhook() *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	return s.Webhook
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) GetWebhookId() *float32 {
	return s.WebhookId
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) GetWebhookName() *string {
	return s.WebhookName
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) SetWebhook(v *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	s.Webhook = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) SetWebhookId(v float32) *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	s.WebhookId = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) SetWebhookName(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	s.WebhookName = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) Validate() error {
	if s.Webhook != nil {
		if err := s.Webhook.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook struct {
	// The header of the HTTP request.
	//
	// example:
	//
	// [{"Content-Type":"application/json"}]
	BizHeaders map[string]interface{} `json:"BizHeaders,omitempty" xml:"BizHeaders,omitempty"`
	// The parameters in the HTTP request.
	//
	// example:
	//
	// [{"content:"mike"}]
	BizParams map[string]interface{} `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	// The alert notification template.
	//
	// example:
	//
	// { "Alert name":"{{ .commonLabels.alertname }}{{if .commonLabels.clustername }}", "Cluster name":"{{ .commonLabels.clustername }} {{ end }}{{if eq "app" .commonLabels._aliyun_arms_involvedObject_kind }}", "Application name":"{{ .commonLabels._aliyun_arms_involvedObject_name }} {{ end }}", "Notification policy":"{{ .dispatchRuleName }}", "Alert time":"{{ .startTime }}", "Alert content":"{{ for .alerts }} {{ .annotations.message }} {{ end }}" }
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The HTTP request method.
	//
	// 	- Get
	//
	// 	- Post
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
	// https://oapi.dingtalk.com/robot/send?access_token=91f2f6****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GetBizHeaders() map[string]interface{} {
	return s.BizHeaders
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GetBizParams() map[string]interface{} {
	return s.BizParams
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GetBody() *string {
	return s.Body
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GetMethod() *string {
	return s.Method
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GetRecoverBody() *string {
	return s.RecoverBody
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GetUrl() *string {
	return s.Url
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetBizHeaders(v map[string]interface{}) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.BizHeaders = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetBizParams(v map[string]interface{}) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.BizParams = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetBody(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.Body = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetMethod(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.Method = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetRecoverBody(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.RecoverBody = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetUrl(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.Url = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) Validate() error {
	return dara.Validate(s)
}
