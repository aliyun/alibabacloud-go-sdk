// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *CreateAlertWebhookRequest
	GetContentType() *string
	SetHeaders(v map[string]*string) *CreateAlertWebhookRequest
	GetHeaders() map[string]*string
	SetLang(v string) *CreateAlertWebhookRequest
	GetLang() *string
	SetMethod(v string) *CreateAlertWebhookRequest
	GetMethod() *string
	SetName(v string) *CreateAlertWebhookRequest
	GetName() *string
	SetUrl(v string) *CreateAlertWebhookRequest
	GetUrl() *string
	SetWebhookId(v string) *CreateAlertWebhookRequest
	GetWebhookId() *string
}

type CreateAlertWebhookRequest struct {
	// example:
	//
	// JSON
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// headers
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/test
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// test
	WebhookId *string `json:"webhookId,omitempty" xml:"webhookId,omitempty"`
}

func (s CreateAlertWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertWebhookRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertWebhookRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateAlertWebhookRequest) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertWebhookRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateAlertWebhookRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateAlertWebhookRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertWebhookRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateAlertWebhookRequest) GetWebhookId() *string {
	return s.WebhookId
}

func (s *CreateAlertWebhookRequest) SetContentType(v string) *CreateAlertWebhookRequest {
	s.ContentType = &v
	return s
}

func (s *CreateAlertWebhookRequest) SetHeaders(v map[string]*string) *CreateAlertWebhookRequest {
	s.Headers = v
	return s
}

func (s *CreateAlertWebhookRequest) SetLang(v string) *CreateAlertWebhookRequest {
	s.Lang = &v
	return s
}

func (s *CreateAlertWebhookRequest) SetMethod(v string) *CreateAlertWebhookRequest {
	s.Method = &v
	return s
}

func (s *CreateAlertWebhookRequest) SetName(v string) *CreateAlertWebhookRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertWebhookRequest) SetUrl(v string) *CreateAlertWebhookRequest {
	s.Url = &v
	return s
}

func (s *CreateAlertWebhookRequest) SetWebhookId(v string) *CreateAlertWebhookRequest {
	s.WebhookId = &v
	return s
}

func (s *CreateAlertWebhookRequest) Validate() error {
	return dara.Validate(s)
}
