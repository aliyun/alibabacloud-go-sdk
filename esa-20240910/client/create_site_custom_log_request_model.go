// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteCustomLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookies(v []*string) *CreateSiteCustomLogRequest
	GetCookies() []*string
	SetRequestHeaders(v []*string) *CreateSiteCustomLogRequest
	GetRequestHeaders() []*string
	SetResponseHeaders(v []*string) *CreateSiteCustomLogRequest
	GetResponseHeaders() []*string
	SetSiteId(v int64) *CreateSiteCustomLogRequest
	GetSiteId() *int64
}

type CreateSiteCustomLogRequest struct {
	// The cookie fields.
	Cookies []*string `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	// The request header fields.
	RequestHeaders []*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// The response header fields.
	ResponseHeaders []*string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 11223
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateSiteCustomLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteCustomLogRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogRequest) GetCookies() []*string {
	return s.Cookies
}

func (s *CreateSiteCustomLogRequest) GetRequestHeaders() []*string {
	return s.RequestHeaders
}

func (s *CreateSiteCustomLogRequest) GetResponseHeaders() []*string {
	return s.ResponseHeaders
}

func (s *CreateSiteCustomLogRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteCustomLogRequest) SetCookies(v []*string) *CreateSiteCustomLogRequest {
	s.Cookies = v
	return s
}

func (s *CreateSiteCustomLogRequest) SetRequestHeaders(v []*string) *CreateSiteCustomLogRequest {
	s.RequestHeaders = v
	return s
}

func (s *CreateSiteCustomLogRequest) SetResponseHeaders(v []*string) *CreateSiteCustomLogRequest {
	s.ResponseHeaders = v
	return s
}

func (s *CreateSiteCustomLogRequest) SetSiteId(v int64) *CreateSiteCustomLogRequest {
	s.SiteId = &v
	return s
}

func (s *CreateSiteCustomLogRequest) Validate() error {
	return dara.Validate(s)
}
