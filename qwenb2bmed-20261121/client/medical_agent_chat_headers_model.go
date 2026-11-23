// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalAgentChatHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MedicalAgentChatHeaders
	GetCommonHeaders() map[string]*string
	SetXDashScopeSSE(v string) *MedicalAgentChatHeaders
	GetXDashScopeSSE() *string
	SetXGWAccess(v string) *MedicalAgentChatHeaders
	GetXGWAccess() *string
	SetXGWNonce(v string) *MedicalAgentChatHeaders
	GetXGWNonce() *string
	SetXGWProtocol(v string) *MedicalAgentChatHeaders
	GetXGWProtocol() *string
	SetXGWSignature(v string) *MedicalAgentChatHeaders
	GetXGWSignature() *string
	SetXGWTimestamp(v string) *MedicalAgentChatHeaders
	GetXGWTimestamp() *string
	SetXRequestId(v string) *MedicalAgentChatHeaders
	GetXRequestId() *string
}

type MedicalAgentChatHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 可选。请求 SSE 时可填 enable。正文中必须显式填写 JSON 布尔值 params.responseStreaming；true 返回 SSE，false 返回完整 JSON。
	//
	// example:
	//
	// enable
	XDashScopeSSE *string `json:"X-DashScope-SSE,omitempty" xml:"X-DashScope-SSE,omitempty"`
	// 接入时领取的业务 Access Key ID，用于识别调用方。阿里云调用凭证按 SDK 的要求单独配置。
	//
	// This parameter is required.
	//
	// example:
	//
	// example-ak-87019a070b49c9fc
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// 每次发送时重新生成的随机串，例如 UUID。重试时也要重新生成 nonce、时间戳和签名。
	//
	// This parameter is required.
	//
	// example:
	//
	// b46a9f08-fe98-4dd3-8885-405c4d032e01
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// 本文示例固定传 aliyun-roa，使用小驼峰格式的请求和响应。此请求头不参与业务签名。
	//
	// example:
	//
	// aliyun-roa
	XGWProtocol *string `json:"X-GW-Protocol,omitempty" xml:"X-GW-Protocol,omitempty"`
	// 使用接入时领取的业务签名密钥（SK）对六段内容计算 HMAC-SHA256。依次为 HTTP 方法（大写）、请求路径、规范化查询串、实际发送正文的 SHA-256 小写十六进制、X-GW-Timestamp、X-GW-Nonce。以换行连接，末尾不加换行；无查询参数时第三段保留为空。查询串按 & 拆分并去除空项，按原始参数项字典序排序后以 & 拼接，不解码。签名传 64 个小写十六进制字符。正文用 UTF-8 序列化一次，用同一份字节签名和发送，SK 仅用于本地计算。
	//
	// This parameter is required.
	//
	// example:
	//
	// df9a3e428a1faa8fd4b0add7173e1a554cd76fb8044176407777bcbbc7b94ba4
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// 发送请求时的 Unix 毫秒时间戳，使用十进制字符串。签名和请求头使用同一个值。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794996000000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// 本次业务请求的关联号，单值，与 communication.requestId 完全相同。允许 1 至 128 个字母、数字或 . _ : - 字符，gw-ingress- 前缀保留给服务诊断。
	//
	// This parameter is required.
	//
	// example:
	//
	// med-req-001
	XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id,omitempty"`
}

func (s MedicalAgentChatHeaders) String() string {
	return dara.Prettify(s)
}

func (s MedicalAgentChatHeaders) GoString() string {
	return s.String()
}

func (s *MedicalAgentChatHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MedicalAgentChatHeaders) GetXDashScopeSSE() *string {
	return s.XDashScopeSSE
}

func (s *MedicalAgentChatHeaders) GetXGWAccess() *string {
	return s.XGWAccess
}

func (s *MedicalAgentChatHeaders) GetXGWNonce() *string {
	return s.XGWNonce
}

func (s *MedicalAgentChatHeaders) GetXGWProtocol() *string {
	return s.XGWProtocol
}

func (s *MedicalAgentChatHeaders) GetXGWSignature() *string {
	return s.XGWSignature
}

func (s *MedicalAgentChatHeaders) GetXGWTimestamp() *string {
	return s.XGWTimestamp
}

func (s *MedicalAgentChatHeaders) GetXRequestId() *string {
	return s.XRequestId
}

func (s *MedicalAgentChatHeaders) SetCommonHeaders(v map[string]*string) *MedicalAgentChatHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MedicalAgentChatHeaders) SetXDashScopeSSE(v string) *MedicalAgentChatHeaders {
	s.XDashScopeSSE = &v
	return s
}

func (s *MedicalAgentChatHeaders) SetXGWAccess(v string) *MedicalAgentChatHeaders {
	s.XGWAccess = &v
	return s
}

func (s *MedicalAgentChatHeaders) SetXGWNonce(v string) *MedicalAgentChatHeaders {
	s.XGWNonce = &v
	return s
}

func (s *MedicalAgentChatHeaders) SetXGWProtocol(v string) *MedicalAgentChatHeaders {
	s.XGWProtocol = &v
	return s
}

func (s *MedicalAgentChatHeaders) SetXGWSignature(v string) *MedicalAgentChatHeaders {
	s.XGWSignature = &v
	return s
}

func (s *MedicalAgentChatHeaders) SetXGWTimestamp(v string) *MedicalAgentChatHeaders {
	s.XGWTimestamp = &v
	return s
}

func (s *MedicalAgentChatHeaders) SetXRequestId(v string) *MedicalAgentChatHeaders {
	s.XRequestId = &v
	return s
}

func (s *MedicalAgentChatHeaders) Validate() error {
	return dara.Validate(s)
}
