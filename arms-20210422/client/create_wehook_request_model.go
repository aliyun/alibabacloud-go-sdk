// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWehookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateWehookRequest
	GetBody() *string
	SetContactName(v string) *CreateWehookRequest
	GetContactName() *string
	SetHttpHeaders(v string) *CreateWehookRequest
	GetHttpHeaders() *string
	SetHttpParams(v string) *CreateWehookRequest
	GetHttpParams() *string
	SetMethod(v string) *CreateWehookRequest
	GetMethod() *string
	SetRegionId(v string) *CreateWehookRequest
	GetRegionId() *string
	SetUrl(v string) *CreateWehookRequest
	GetUrl() *string
}

type CreateWehookRequest struct {
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// This parameter is required.
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateWehookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWehookRequest) GoString() string {
	return s.String()
}

func (s *CreateWehookRequest) GetBody() *string {
	return s.Body
}

func (s *CreateWehookRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateWehookRequest) GetHttpHeaders() *string {
	return s.HttpHeaders
}

func (s *CreateWehookRequest) GetHttpParams() *string {
	return s.HttpParams
}

func (s *CreateWehookRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateWehookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWehookRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateWehookRequest) SetBody(v string) *CreateWehookRequest {
	s.Body = &v
	return s
}

func (s *CreateWehookRequest) SetContactName(v string) *CreateWehookRequest {
	s.ContactName = &v
	return s
}

func (s *CreateWehookRequest) SetHttpHeaders(v string) *CreateWehookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *CreateWehookRequest) SetHttpParams(v string) *CreateWehookRequest {
	s.HttpParams = &v
	return s
}

func (s *CreateWehookRequest) SetMethod(v string) *CreateWehookRequest {
	s.Method = &v
	return s
}

func (s *CreateWehookRequest) SetRegionId(v string) *CreateWehookRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWehookRequest) SetUrl(v string) *CreateWehookRequest {
	s.Url = &v
	return s
}

func (s *CreateWehookRequest) Validate() error {
	return dara.Validate(s)
}
