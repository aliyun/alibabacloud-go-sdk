// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamVoiceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StreamVoiceHeaders
	GetCommonHeaders() map[string]*string
	SetXGWAccess(v string) *StreamVoiceHeaders
	GetXGWAccess() *string
	SetXGWNonce(v string) *StreamVoiceHeaders
	GetXGWNonce() *string
	SetXGWProtocol(v string) *StreamVoiceHeaders
	GetXGWProtocol() *string
	SetXGWSignature(v string) *StreamVoiceHeaders
	GetXGWSignature() *string
	SetXGWTimestamp(v string) *StreamVoiceHeaders
	GetXGWTimestamp() *string
	SetXRequestId(v string) *StreamVoiceHeaders
	GetXRequestId() *string
}

type StreamVoiceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 业务网关 Access Key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// example-gateway-access
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// 每次调用使用的随机串
	//
	// This parameter is required.
	//
	// example:
	//
	// c093adf9-287a-4684-982a-1aea5a248c2e
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// 小驼峰回源协议固定传 aliyun-roa，省略时为旧协议；不参与 HMAC 签名
	//
	// example:
	//
	// aliyun-roa
	XGWProtocol *string `json:"X-GW-Protocol,omitempty" xml:"X-GW-Protocol,omitempty"`
	// 按实际请求字节计算的 HMAC-SHA256 小写十六进制签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// 当前 Unix 毫秒时间戳
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794787200000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// 可选的 WebSocket 握手关联号，每轮业务 ID 使用 run-task 帧中的 communication.requestId
	//
	// example:
	//
	// ws-example-001
	XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id,omitempty"`
}

func (s StreamVoiceHeaders) String() string {
	return dara.Prettify(s)
}

func (s StreamVoiceHeaders) GoString() string {
	return s.String()
}

func (s *StreamVoiceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StreamVoiceHeaders) GetXGWAccess() *string {
	return s.XGWAccess
}

func (s *StreamVoiceHeaders) GetXGWNonce() *string {
	return s.XGWNonce
}

func (s *StreamVoiceHeaders) GetXGWProtocol() *string {
	return s.XGWProtocol
}

func (s *StreamVoiceHeaders) GetXGWSignature() *string {
	return s.XGWSignature
}

func (s *StreamVoiceHeaders) GetXGWTimestamp() *string {
	return s.XGWTimestamp
}

func (s *StreamVoiceHeaders) GetXRequestId() *string {
	return s.XRequestId
}

func (s *StreamVoiceHeaders) SetCommonHeaders(v map[string]*string) *StreamVoiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StreamVoiceHeaders) SetXGWAccess(v string) *StreamVoiceHeaders {
	s.XGWAccess = &v
	return s
}

func (s *StreamVoiceHeaders) SetXGWNonce(v string) *StreamVoiceHeaders {
	s.XGWNonce = &v
	return s
}

func (s *StreamVoiceHeaders) SetXGWProtocol(v string) *StreamVoiceHeaders {
	s.XGWProtocol = &v
	return s
}

func (s *StreamVoiceHeaders) SetXGWSignature(v string) *StreamVoiceHeaders {
	s.XGWSignature = &v
	return s
}

func (s *StreamVoiceHeaders) SetXGWTimestamp(v string) *StreamVoiceHeaders {
	s.XGWTimestamp = &v
	return s
}

func (s *StreamVoiceHeaders) SetXRequestId(v string) *StreamVoiceHeaders {
	s.XRequestId = &v
	return s
}

func (s *StreamVoiceHeaders) Validate() error {
	return dara.Validate(s)
}
