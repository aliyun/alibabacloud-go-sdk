// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *OpenServiceHeaders
	GetCommonHeaders() map[string]*string
	SetXGWAccess(v string) *OpenServiceHeaders
	GetXGWAccess() *string
	SetXGWNonce(v string) *OpenServiceHeaders
	GetXGWNonce() *string
	SetXGWSignature(v string) *OpenServiceHeaders
	GetXGWSignature() *string
	SetXGWTimestamp(v string) *OpenServiceHeaders
	GetXGWTimestamp() *string
	SetXRequestId(v string) *OpenServiceHeaders
	GetXRequestId() *string
}

type OpenServiceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
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
	// b46a9f08-fe98-4dd3-8885-405c4d032eff
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// 使用接入时领取的业务签名密钥（SK）对六段内容计算 HMAC-SHA256。依次为 HTTP 方法（大写）、请求路径、规范化查询串、实际发送正文的 SHA-256 小写十六进制、X-GW-Timestamp、X-GW-Nonce。以换行连接，末尾不加换行；无查询参数时第三段保留为空。查询串按 & 拆分并去除空项，按原始参数项字典序排序后以 & 拼接，不解码。签名传 64 个小写十六进制字符。正文用 UTF-8 序列化一次，用同一份字节签名和发送，SK 仅用于本地计算。
	//
	// This parameter is required.
	//
	// example:
	//
	// fe3241f17796f523d9df27138f9a22abaead6cc9d1b82bae4b45eedc0e1858d7
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// 发送请求时的 Unix 毫秒时间戳，使用十进制字符串。签名和请求头使用同一个值。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794996000000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// 每次业务调用生成一个新的请求 ID，用于关联请求和响应。允许 1 至 128 个字母、数字或 . _ : - 字符。
	//
	// This parameter is required.
	//
	// example:
	//
	// med-req-001
	XRequestId *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id,omitempty"`
}

func (s OpenServiceHeaders) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceHeaders) GoString() string {
	return s.String()
}

func (s *OpenServiceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *OpenServiceHeaders) GetXGWAccess() *string {
	return s.XGWAccess
}

func (s *OpenServiceHeaders) GetXGWNonce() *string {
	return s.XGWNonce
}

func (s *OpenServiceHeaders) GetXGWSignature() *string {
	return s.XGWSignature
}

func (s *OpenServiceHeaders) GetXGWTimestamp() *string {
	return s.XGWTimestamp
}

func (s *OpenServiceHeaders) GetXRequestId() *string {
	return s.XRequestId
}

func (s *OpenServiceHeaders) SetCommonHeaders(v map[string]*string) *OpenServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenServiceHeaders) SetXGWAccess(v string) *OpenServiceHeaders {
	s.XGWAccess = &v
	return s
}

func (s *OpenServiceHeaders) SetXGWNonce(v string) *OpenServiceHeaders {
	s.XGWNonce = &v
	return s
}

func (s *OpenServiceHeaders) SetXGWSignature(v string) *OpenServiceHeaders {
	s.XGWSignature = &v
	return s
}

func (s *OpenServiceHeaders) SetXGWTimestamp(v string) *OpenServiceHeaders {
	s.XGWTimestamp = &v
	return s
}

func (s *OpenServiceHeaders) SetXRequestId(v string) *OpenServiceHeaders {
	s.XRequestId = &v
	return s
}

func (s *OpenServiceHeaders) Validate() error {
	return dara.Validate(s)
}
