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
	// 必填的业务网关 AK，由服务接入方签发，用于识别租户和凭据；与阿里云 API 鉴权凭据用途不同，不发送业务网关 SK。
	//
	// This parameter is required.
	//
	// example:
	//
	// example-gateway-access
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// 必填的一次性随机串，原值参与业务网关签名；每次请求和重试重新生成，不能重放原签名。
	//
	// This parameter is required.
	//
	// example:
	//
	// c093adf9-287a-4684-982a-1aea5a248c2e
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// 必填的业务网关 HMAC-SHA256 签名，64 位小写十六进制。签名原文为方法、路径、空查询串、实际请求体字节的 SHA-256 小写摘要、时间戳、nonce 六段，以换行符连接。签名后不得改变请求体字节；阿里云 API 签名另行生成。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0000000000000000000000000000000000000000000000000000000000000000
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// 必填的当前 Unix 毫秒时间戳，用数字字符串表示，原值参与业务网关签名；重试时更新并重新签名。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794787200000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// 必填的客户请求关联号，恰好一个，匹配 [A-Za-z0-9._:-]{1,128}；不能使用 gw-ingress- 保留前缀。仅用于关联本次调用，不是订单号或幂等令牌。
	//
	// This parameter is required.
	//
	// example:
	//
	// req-example-001
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
