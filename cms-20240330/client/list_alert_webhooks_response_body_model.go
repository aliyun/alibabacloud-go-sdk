// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertWebhooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAlertWebhooksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertWebhooksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAlertWebhooksResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAlertWebhooksResponseBody
	GetTotal() *int64
	SetWebhooks(v []*ListAlertWebhooksResponseBodyWebhooks) *ListAlertWebhooksResponseBody
	GetWebhooks() []*ListAlertWebhooksResponseBodyWebhooks
}

type ListAlertWebhooksResponseBody struct {
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8A33DBEA-*****-*****-*****-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 8
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// The webhooks.
	Webhooks []*ListAlertWebhooksResponseBodyWebhooks `json:"webhooks,omitempty" xml:"webhooks,omitempty" type:"Repeated"`
}

func (s ListAlertWebhooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertWebhooksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertWebhooksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertWebhooksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertWebhooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertWebhooksResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlertWebhooksResponseBody) GetWebhooks() []*ListAlertWebhooksResponseBodyWebhooks {
	return s.Webhooks
}

func (s *ListAlertWebhooksResponseBody) SetPageNumber(v int64) *ListAlertWebhooksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertWebhooksResponseBody) SetPageSize(v int64) *ListAlertWebhooksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertWebhooksResponseBody) SetRequestId(v string) *ListAlertWebhooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertWebhooksResponseBody) SetTotal(v int64) *ListAlertWebhooksResponseBody {
	s.Total = &v
	return s
}

func (s *ListAlertWebhooksResponseBody) SetWebhooks(v []*ListAlertWebhooksResponseBodyWebhooks) *ListAlertWebhooksResponseBody {
	s.Webhooks = v
	return s
}

func (s *ListAlertWebhooksResponseBody) Validate() error {
	if s.Webhooks != nil {
		for _, item := range s.Webhooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertWebhooksResponseBodyWebhooks struct {
	// The content type of the data. Valid values:
	//
	// - JSON
	//
	// - FORM
	//
	// example:
	//
	// JSON
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The headers.
	//
	// example:
	//
	// key
	Headers map[string]interface{} `json:"headers,omitempty" xml:"headers,omitempty"`
	// The language. Valid values:
	//
	// - zh_CN
	//
	// - en_US
	//
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// The request method. Valid values:
	//
	// - GET
	//
	// - POST
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The name of the webhook.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The URL of the alert callback.
	//
	// example:
	//
	// http://aliyun.com/test
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// The unique ID of the webhook.
	//
	// example:
	//
	// test
	WebhookId *string `json:"webhookId,omitempty" xml:"webhookId,omitempty"`
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAlertWebhooksResponseBodyWebhooks) String() string {
	return dara.Prettify(s)
}

func (s ListAlertWebhooksResponseBodyWebhooks) GoString() string {
	return s.String()
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetContentType() *string {
	return s.ContentType
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetHeaders() map[string]interface{} {
	return s.Headers
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetLang() *string {
	return s.Lang
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetMethod() *string {
	return s.Method
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetName() *string {
	return s.Name
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetUrl() *string {
	return s.Url
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetWebhookId() *string {
	return s.WebhookId
}

func (s *ListAlertWebhooksResponseBodyWebhooks) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetContentType(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.ContentType = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetHeaders(v map[string]interface{}) *ListAlertWebhooksResponseBodyWebhooks {
	s.Headers = v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetLang(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.Lang = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetMethod(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.Method = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetName(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.Name = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetUrl(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.Url = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetWebhookId(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.WebhookId = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) SetWorkspace(v string) *ListAlertWebhooksResponseBodyWebhooks {
	s.Workspace = &v
	return s
}

func (s *ListAlertWebhooksResponseBodyWebhooks) Validate() error {
	return dara.Validate(s)
}
