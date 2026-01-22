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
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 8A33DBEA-*****-*****-*****-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 8
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// webhooks
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
	// example:
	//
	// JSON
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// headers
	//
	// example:
	//
	// key
	Headers map[string]interface{} `json:"headers,omitempty" xml:"headers,omitempty"`
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// http://aliyun.com/test
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// test
	WebhookId *string `json:"webhookId,omitempty" xml:"webhookId,omitempty"`
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

func (s *ListAlertWebhooksResponseBodyWebhooks) Validate() error {
	return dara.Validate(s)
}
