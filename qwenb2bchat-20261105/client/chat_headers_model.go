// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ChatHeaders
	GetCommonHeaders() map[string]*string
	SetAccept(v string) *ChatHeaders
	GetAccept() *string
	SetXDashScopeSSE(v string) *ChatHeaders
	GetXDashScopeSSE() *string
	SetXGWAccess(v string) *ChatHeaders
	GetXGWAccess() *string
	SetXGWNonce(v string) *ChatHeaders
	GetXGWNonce() *string
	SetXGWProtocol(v string) *ChatHeaders
	GetXGWProtocol() *string
	SetXGWSignature(v string) *ChatHeaders
	GetXGWSignature() *string
	SetXGWTimestamp(v string) *ChatHeaders
	GetXGWTimestamp() *string
	SetXRequestId(v string) *ChatHeaders
	GetXRequestId() *string
}

type ChatHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The response media types accepted by the client.
	//
	// example:
	//
	// application/json
	Accept *string `json:"Accept,omitempty" xml:"Accept,omitempty"`
	// The compatible SSE selector header.
	//
	// example:
	//
	// enable
	XDashScopeSSE *string `json:"X-DashScope-SSE,omitempty" xml:"X-DashScope-SSE,omitempty"`
	// The AccessKey ID for the business gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-gateway-access
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// The random string used for each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// c093adf9-287a-4684-982a-1aea5a248c2e
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// The field naming protocol. Set to aliyun-roa to use lower camel case fields, or omit to use the legacy protocol. This header does not participate in HMAC signing.
	//
	// example:
	//
	// aliyun-roa
	XGWProtocol *string `json:"X-GW-Protocol,omitempty" xml:"X-GW-Protocol,omitempty"`
	// The lowercase hexadecimal HMAC-SHA256 signature calculated based on the actual request bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// The current Unix millisecond timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794787200000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// The business request ID, which must be exactly the same as communication.requestId.
	//
	// This parameter is required.
	//
	// example:
	//
	// req-example-001
	XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id,omitempty"`
}

func (s ChatHeaders) String() string {
	return dara.Prettify(s)
}

func (s ChatHeaders) GoString() string {
	return s.String()
}

func (s *ChatHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ChatHeaders) GetAccept() *string {
	return s.Accept
}

func (s *ChatHeaders) GetXDashScopeSSE() *string {
	return s.XDashScopeSSE
}

func (s *ChatHeaders) GetXGWAccess() *string {
	return s.XGWAccess
}

func (s *ChatHeaders) GetXGWNonce() *string {
	return s.XGWNonce
}

func (s *ChatHeaders) GetXGWProtocol() *string {
	return s.XGWProtocol
}

func (s *ChatHeaders) GetXGWSignature() *string {
	return s.XGWSignature
}

func (s *ChatHeaders) GetXGWTimestamp() *string {
	return s.XGWTimestamp
}

func (s *ChatHeaders) GetXRequestId() *string {
	return s.XRequestId
}

func (s *ChatHeaders) SetCommonHeaders(v map[string]*string) *ChatHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChatHeaders) SetAccept(v string) *ChatHeaders {
	s.Accept = &v
	return s
}

func (s *ChatHeaders) SetXDashScopeSSE(v string) *ChatHeaders {
	s.XDashScopeSSE = &v
	return s
}

func (s *ChatHeaders) SetXGWAccess(v string) *ChatHeaders {
	s.XGWAccess = &v
	return s
}

func (s *ChatHeaders) SetXGWNonce(v string) *ChatHeaders {
	s.XGWNonce = &v
	return s
}

func (s *ChatHeaders) SetXGWProtocol(v string) *ChatHeaders {
	s.XGWProtocol = &v
	return s
}

func (s *ChatHeaders) SetXGWSignature(v string) *ChatHeaders {
	s.XGWSignature = &v
	return s
}

func (s *ChatHeaders) SetXGWTimestamp(v string) *ChatHeaders {
	s.XGWTimestamp = &v
	return s
}

func (s *ChatHeaders) SetXRequestId(v string) *ChatHeaders {
	s.XRequestId = &v
	return s
}

func (s *ChatHeaders) Validate() error {
	return dara.Validate(s)
}
