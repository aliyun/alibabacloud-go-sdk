// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCustomLogShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookiesShrink(v string) *UpdateSiteCustomLogShrinkRequest
	GetCookiesShrink() *string
	SetRequestHeadersShrink(v string) *UpdateSiteCustomLogShrinkRequest
	GetRequestHeadersShrink() *string
	SetResponseHeadersShrink(v string) *UpdateSiteCustomLogShrinkRequest
	GetResponseHeadersShrink() *string
	SetSiteId(v int64) *UpdateSiteCustomLogShrinkRequest
	GetSiteId() *int64
}

type UpdateSiteCustomLogShrinkRequest struct {
	// The cookie fields.
	CookiesShrink *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The request header fields.
	RequestHeadersShrink *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	// The response header fields.
	ResponseHeadersShrink *string `json:"ResponseHeaders,omitempty" xml:"ResponseHeaders,omitempty"`
	// site id
	//
	// example:
	//
	// 11223****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateSiteCustomLogShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCustomLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogShrinkRequest) GetCookiesShrink() *string {
	return s.CookiesShrink
}

func (s *UpdateSiteCustomLogShrinkRequest) GetRequestHeadersShrink() *string {
	return s.RequestHeadersShrink
}

func (s *UpdateSiteCustomLogShrinkRequest) GetResponseHeadersShrink() *string {
	return s.ResponseHeadersShrink
}

func (s *UpdateSiteCustomLogShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateSiteCustomLogShrinkRequest) SetCookiesShrink(v string) *UpdateSiteCustomLogShrinkRequest {
	s.CookiesShrink = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) SetRequestHeadersShrink(v string) *UpdateSiteCustomLogShrinkRequest {
	s.RequestHeadersShrink = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) SetResponseHeadersShrink(v string) *UpdateSiteCustomLogShrinkRequest {
	s.ResponseHeadersShrink = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) SetSiteId(v int64) *UpdateSiteCustomLogShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateSiteCustomLogShrinkRequest) Validate() error {
	return dara.Validate(s)
}
