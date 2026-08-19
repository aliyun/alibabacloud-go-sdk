// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceSiteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *TraceSiteShrinkRequest
	GetBodyShrink() *string
	SetContextShrink(v string) *TraceSiteShrinkRequest
	GetContextShrink() *string
	SetCookiesShrink(v string) *TraceSiteShrinkRequest
	GetCookiesShrink() *string
	SetHeadersShrink(v string) *TraceSiteShrinkRequest
	GetHeadersShrink() *string
	SetMethod(v string) *TraceSiteShrinkRequest
	GetMethod() *string
	SetProtocol(v string) *TraceSiteShrinkRequest
	GetProtocol() *string
	SetUrl(v string) *TraceSiteShrinkRequest
	GetUrl() *string
}

type TraceSiteShrinkRequest struct {
	// The HTTP request body.
	//
	// example:
	//
	// {"PlainText":"bc58c54211db"}
	BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The environment context. This parameter is optional.
	ContextShrink *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The cookie parameters.
	//
	// example:
	//
	// []
	CookiesShrink *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The request headers.
	//
	// example:
	//
	// []
	HeadersShrink *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The HTTP method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The HTTP protocol.
	//
	// example:
	//
	// HTTP/1.1
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The URL of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/test
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s TraceSiteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteShrinkRequest) GoString() string {
	return s.String()
}

func (s *TraceSiteShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *TraceSiteShrinkRequest) GetContextShrink() *string {
	return s.ContextShrink
}

func (s *TraceSiteShrinkRequest) GetCookiesShrink() *string {
	return s.CookiesShrink
}

func (s *TraceSiteShrinkRequest) GetHeadersShrink() *string {
	return s.HeadersShrink
}

func (s *TraceSiteShrinkRequest) GetMethod() *string {
	return s.Method
}

func (s *TraceSiteShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *TraceSiteShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *TraceSiteShrinkRequest) SetBodyShrink(v string) *TraceSiteShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *TraceSiteShrinkRequest) SetContextShrink(v string) *TraceSiteShrinkRequest {
	s.ContextShrink = &v
	return s
}

func (s *TraceSiteShrinkRequest) SetCookiesShrink(v string) *TraceSiteShrinkRequest {
	s.CookiesShrink = &v
	return s
}

func (s *TraceSiteShrinkRequest) SetHeadersShrink(v string) *TraceSiteShrinkRequest {
	s.HeadersShrink = &v
	return s
}

func (s *TraceSiteShrinkRequest) SetMethod(v string) *TraceSiteShrinkRequest {
	s.Method = &v
	return s
}

func (s *TraceSiteShrinkRequest) SetProtocol(v string) *TraceSiteShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *TraceSiteShrinkRequest) SetUrl(v string) *TraceSiteShrinkRequest {
	s.Url = &v
	return s
}

func (s *TraceSiteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
