// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpRequestVodTestToolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v string) *HttpRequestVodTestToolShrinkRequest
	GetArgs() *string
	SetBody(v string) *HttpRequestVodTestToolShrinkRequest
	GetBody() *string
	SetHeaderShrink(v string) *HttpRequestVodTestToolShrinkRequest
	GetHeaderShrink() *string
	SetHost(v string) *HttpRequestVodTestToolShrinkRequest
	GetHost() *string
	SetMethod(v string) *HttpRequestVodTestToolShrinkRequest
	GetMethod() *string
	SetOwnerId(v int64) *HttpRequestVodTestToolShrinkRequest
	GetOwnerId() *int64
	SetProxyIp(v string) *HttpRequestVodTestToolShrinkRequest
	GetProxyIp() *string
	SetScheme(v string) *HttpRequestVodTestToolShrinkRequest
	GetScheme() *string
	SetUri(v string) *HttpRequestVodTestToolShrinkRequest
	GetUri() *string
}

type HttpRequestVodTestToolShrinkRequest struct {
	Args         *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Body         *string `json:"Body,omitempty" xml:"Body,omitempty"`
	HeaderShrink *string `json:"Header,omitempty" xml:"Header,omitempty"`
	// This parameter is required.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// This parameter is required.
	Method  *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ProxyIp *string `json:"ProxyIp,omitempty" xml:"ProxyIp,omitempty"`
	// This parameter is required.
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
	Uri    *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s HttpRequestVodTestToolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HttpRequestVodTestToolShrinkRequest) GoString() string {
	return s.String()
}

func (s *HttpRequestVodTestToolShrinkRequest) GetArgs() *string {
	return s.Args
}

func (s *HttpRequestVodTestToolShrinkRequest) GetBody() *string {
	return s.Body
}

func (s *HttpRequestVodTestToolShrinkRequest) GetHeaderShrink() *string {
	return s.HeaderShrink
}

func (s *HttpRequestVodTestToolShrinkRequest) GetHost() *string {
	return s.Host
}

func (s *HttpRequestVodTestToolShrinkRequest) GetMethod() *string {
	return s.Method
}

func (s *HttpRequestVodTestToolShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *HttpRequestVodTestToolShrinkRequest) GetProxyIp() *string {
	return s.ProxyIp
}

func (s *HttpRequestVodTestToolShrinkRequest) GetScheme() *string {
	return s.Scheme
}

func (s *HttpRequestVodTestToolShrinkRequest) GetUri() *string {
	return s.Uri
}

func (s *HttpRequestVodTestToolShrinkRequest) SetArgs(v string) *HttpRequestVodTestToolShrinkRequest {
	s.Args = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetBody(v string) *HttpRequestVodTestToolShrinkRequest {
	s.Body = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetHeaderShrink(v string) *HttpRequestVodTestToolShrinkRequest {
	s.HeaderShrink = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetHost(v string) *HttpRequestVodTestToolShrinkRequest {
	s.Host = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetMethod(v string) *HttpRequestVodTestToolShrinkRequest {
	s.Method = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetOwnerId(v int64) *HttpRequestVodTestToolShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetProxyIp(v string) *HttpRequestVodTestToolShrinkRequest {
	s.ProxyIp = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetScheme(v string) *HttpRequestVodTestToolShrinkRequest {
	s.Scheme = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) SetUri(v string) *HttpRequestVodTestToolShrinkRequest {
	s.Uri = &v
	return s
}

func (s *HttpRequestVodTestToolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
