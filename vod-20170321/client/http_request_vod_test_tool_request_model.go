// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpRequestVodTestToolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v string) *HttpRequestVodTestToolRequest
	GetArgs() *string
	SetBody(v string) *HttpRequestVodTestToolRequest
	GetBody() *string
	SetHeader(v map[string]interface{}) *HttpRequestVodTestToolRequest
	GetHeader() map[string]interface{}
	SetHost(v string) *HttpRequestVodTestToolRequest
	GetHost() *string
	SetMethod(v string) *HttpRequestVodTestToolRequest
	GetMethod() *string
	SetOwnerId(v int64) *HttpRequestVodTestToolRequest
	GetOwnerId() *int64
	SetProxyIp(v string) *HttpRequestVodTestToolRequest
	GetProxyIp() *string
	SetScheme(v string) *HttpRequestVodTestToolRequest
	GetScheme() *string
	SetUri(v string) *HttpRequestVodTestToolRequest
	GetUri() *string
}

type HttpRequestVodTestToolRequest struct {
	Args   *string                `json:"Args,omitempty" xml:"Args,omitempty"`
	Body   *string                `json:"Body,omitempty" xml:"Body,omitempty"`
	Header map[string]interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
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

func (s HttpRequestVodTestToolRequest) String() string {
	return dara.Prettify(s)
}

func (s HttpRequestVodTestToolRequest) GoString() string {
	return s.String()
}

func (s *HttpRequestVodTestToolRequest) GetArgs() *string {
	return s.Args
}

func (s *HttpRequestVodTestToolRequest) GetBody() *string {
	return s.Body
}

func (s *HttpRequestVodTestToolRequest) GetHeader() map[string]interface{} {
	return s.Header
}

func (s *HttpRequestVodTestToolRequest) GetHost() *string {
	return s.Host
}

func (s *HttpRequestVodTestToolRequest) GetMethod() *string {
	return s.Method
}

func (s *HttpRequestVodTestToolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *HttpRequestVodTestToolRequest) GetProxyIp() *string {
	return s.ProxyIp
}

func (s *HttpRequestVodTestToolRequest) GetScheme() *string {
	return s.Scheme
}

func (s *HttpRequestVodTestToolRequest) GetUri() *string {
	return s.Uri
}

func (s *HttpRequestVodTestToolRequest) SetArgs(v string) *HttpRequestVodTestToolRequest {
	s.Args = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetBody(v string) *HttpRequestVodTestToolRequest {
	s.Body = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetHeader(v map[string]interface{}) *HttpRequestVodTestToolRequest {
	s.Header = v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetHost(v string) *HttpRequestVodTestToolRequest {
	s.Host = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetMethod(v string) *HttpRequestVodTestToolRequest {
	s.Method = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetOwnerId(v int64) *HttpRequestVodTestToolRequest {
	s.OwnerId = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetProxyIp(v string) *HttpRequestVodTestToolRequest {
	s.ProxyIp = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetScheme(v string) *HttpRequestVodTestToolRequest {
	s.Scheme = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) SetUri(v string) *HttpRequestVodTestToolRequest {
	s.Uri = &v
	return s
}

func (s *HttpRequestVodTestToolRequest) Validate() error {
	return dara.Validate(s)
}
