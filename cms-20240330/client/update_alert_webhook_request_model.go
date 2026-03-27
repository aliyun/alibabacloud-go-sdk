// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *UpdateAlertWebhookRequest
	GetContentType() *string
	SetHeaders(v map[string]*string) *UpdateAlertWebhookRequest
	GetHeaders() map[string]*string
	SetLang(v string) *UpdateAlertWebhookRequest
	GetLang() *string
	SetMethod(v string) *UpdateAlertWebhookRequest
	GetMethod() *string
	SetName(v string) *UpdateAlertWebhookRequest
	GetName() *string
	SetUrl(v string) *UpdateAlertWebhookRequest
	GetUrl() *string
}

type UpdateAlertWebhookRequest struct {
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
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// http://aliyun.com/test
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UpdateAlertWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertWebhookRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateAlertWebhookRequest) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertWebhookRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAlertWebhookRequest) GetMethod() *string {
	return s.Method
}

func (s *UpdateAlertWebhookRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertWebhookRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateAlertWebhookRequest) SetContentType(v string) *UpdateAlertWebhookRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateAlertWebhookRequest) SetHeaders(v map[string]*string) *UpdateAlertWebhookRequest {
	s.Headers = v
	return s
}

func (s *UpdateAlertWebhookRequest) SetLang(v string) *UpdateAlertWebhookRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAlertWebhookRequest) SetMethod(v string) *UpdateAlertWebhookRequest {
	s.Method = &v
	return s
}

func (s *UpdateAlertWebhookRequest) SetName(v string) *UpdateAlertWebhookRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertWebhookRequest) SetUrl(v string) *UpdateAlertWebhookRequest {
	s.Url = &v
	return s
}

func (s *UpdateAlertWebhookRequest) Validate() error {
	return dara.Validate(s)
}
