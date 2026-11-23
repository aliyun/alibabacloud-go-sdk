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
	// 握手必填的业务网关 AK，用于识别租户和凭据；与阿里云 API AK 用途不同。业务网关 SK 仅在客户端用于计算签名，不发送。
	//
	// This parameter is required.
	//
	// example:
	//
	// example-gateway-access
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// 握手必填，一次性随机串，原值参与业务签名；每次重新连接重新生成，不能重放原握手签名。
	//
	// This parameter is required.
	//
	// example:
	//
	// c093adf9-287a-4684-982a-1aea5a248c2e
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// 可选，当前公开小驼峰字段协议应传 aliyun-roa；省略时采用兼容旧协议，其字段命名不同。握手时选择后作用于本连接，不能在后续帧中切换。本头不参与 X-GW 签名。
	//
	// example:
	//
	// aliyun-roa
	XGWProtocol *string `json:"X-GW-Protocol,omitempty" xml:"X-GW-Protocol,omitempty"`
	// 握手必填，64 位小写十六进制 HMAC-SHA256。签名原文为 GET、实际路径、规范化实际 query、空请求体的 SHA-256 小写摘要、时间戳和 nonce，以换行符连接。对原文 UTF-8 字节使用业务网关 SK 签名；阿里云 API 签名另行生成。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// 握手必填，当前 Unix 毫秒时间戳的数字字符串，原值参与业务签名；每次重新连接更新后重新签名。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794787200000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// 可选连接级关联号。省略时生成服务端诊断号。显式提供时只允许一个 1–128 位的 [A-Za-z0-9._:-] 字符串，gw-ingress- 前缀保留给服务。格式非法或提供多个值时返回 HTTP 400。每轮业务请求号使用 run-task.communication.requestId。
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
