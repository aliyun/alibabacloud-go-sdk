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
	SetRegionId(v string) *UpdateWebhookRequest
	GetRegionId() *string
	SetUrl(v string) *UpdateWebhookRequest
	GetUrl() *string
}

type UpdateWebhookRequest struct {
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// This parameter is required.
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
