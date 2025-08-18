// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCustomLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookies(v []*string) *UpdateSiteCustomLogRequest
	GetCookies() []*string
	SetRequestHeaders(v []*string) *UpdateSiteCustomLogRequest
	GetRequestHeaders() []*string
	SetResponseHeaders(v []*string) *UpdateSiteCustomLogRequest
	GetResponseHeaders() []*string
	SetSiteId(v int64) *UpdateSiteCustomLogRequest
	GetSiteId() *int64
}

type UpdateSiteCustomLogRequest struct {
	// The cookie fields.
	Cookies []*string `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	// The request header fields.
	RequestHeaders []*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// The response header fields.
	ResponseHeaders []*string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty" type:"Repeated"`
	// site id
	//
	// example:
	//
	// 11223****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteCustomLogRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCustomLogRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogRequest) GetCookies() []*string {
	return s.Cookies
}

func (s *UpdateSiteCustomLogRequest) GetRequestHeaders() []*string {
	return s.RequestHeaders
}

func (s *UpdateSiteCustomLogRequest) GetResponseHeaders() []*string {
	return s.ResponseHeaders
}

func (s *UpdateSiteCustomLogRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteCustomLogRequest) SetCookies(v []*string) *UpdateSiteCustomLogRequest {
	s.Cookies = v
	return s
}

func (s *UpdateSiteCustomLogRequest) SetRequestHeaders(v []*string) *UpdateSiteCustomLogRequest {
	s.RequestHeaders = v
	return s
}

func (s *UpdateSiteCustomLogRequest) SetResponseHeaders(v []*string) *UpdateSiteCustomLogRequest {
	s.ResponseHeaders = v
	return s
}

func (s *UpdateSiteCustomLogRequest) SetSiteId(v int64) *UpdateSiteCustomLogRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteCustomLogRequest) Validate() error {
	return dara.Validate(s)
}
