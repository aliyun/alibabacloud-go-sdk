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
	// The business AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-ak-87019a070b49c9fc
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// The random string for this request.
	//
	// This parameter is required.
	//
	// example:
	//
	// b46a9f08-fe98-4dd3-8885-405c4d032e01
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// The format of request and response fields.
	//
	// example:
	//
	// aliyun-roa
	XGWProtocol *string `json:"X-GW-Protocol,omitempty" xml:"X-GW-Protocol,omitempty"`
	// The business request signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// df9a3e428a1faa8fd4b0add7173e1a554cd76fb8044176407777bcbbc7b94ba4
	XGWSignature *string `json:"X-GW-Signature,omitempty" xml:"X-GW-Signature,omitempty"`
	// The request timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1794996000000
	XGWTimestamp *string `json:"X-GW-Timestamp,omitempty" xml:"X-GW-Timestamp,omitempty"`
	// The business request ID.
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
