// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteCustomLogShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookiesShrink(v string) *CreateSiteCustomLogShrinkRequest
	GetCookiesShrink() *string
	SetRequestHeadersShrink(v string) *CreateSiteCustomLogShrinkRequest
	GetRequestHeadersShrink() *string
	SetResponseHeadersShrink(v string) *CreateSiteCustomLogShrinkRequest
	GetResponseHeadersShrink() *string
	SetSiteId(v int64) *CreateSiteCustomLogShrinkRequest
	GetSiteId() *int64
}

type CreateSiteCustomLogShrinkRequest struct {
	// The cookie fields.
	CookiesShrink *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The request header fields.
	RequestHeadersShrink *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	// The response header fields.
	ResponseHeadersShrink *string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 11223
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateSiteCustomLogShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteCustomLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogShrinkRequest) GetCookiesShrink() *string {
	return s.CookiesShrink
}

func (s *CreateSiteCustomLogShrinkRequest) GetRequestHeadersShrink() *string {
	return s.RequestHeadersShrink
}

func (s *CreateSiteCustomLogShrinkRequest) GetResponseHeadersShrink() *string {
	return s.ResponseHeadersShrink
}

func (s *CreateSiteCustomLogShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteCustomLogShrinkRequest) SetCookiesShrink(v string) *CreateSiteCustomLogShrinkRequest {
	s.CookiesShrink = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) SetRequestHeadersShrink(v string) *CreateSiteCustomLogShrinkRequest {
	s.RequestHeadersShrink = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) SetResponseHeadersShrink(v string) *CreateSiteCustomLogShrinkRequest {
	s.ResponseHeadersShrink = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) SetSiteId(v int64) *CreateSiteCustomLogShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateSiteCustomLogShrinkRequest) Validate() error {
	return dara.Validate(s)
}
