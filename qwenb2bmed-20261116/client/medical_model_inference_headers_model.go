// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMedicalModelInferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MedicalModelInferenceHeaders
	GetCommonHeaders() map[string]*string
	SetXGWAccess(v string) *MedicalModelInferenceHeaders
	GetXGWAccess() *string
	SetXGWNonce(v string) *MedicalModelInferenceHeaders
	GetXGWNonce() *string
	SetXGWProtocol(v string) *MedicalModelInferenceHeaders
	GetXGWProtocol() *string
	SetXGWSignature(v string) *MedicalModelInferenceHeaders
	GetXGWSignature() *string
	SetXGWTimestamp(v string) *MedicalModelInferenceHeaders
	GetXGWTimestamp() *string
	SetXRequestId(v string) *MedicalModelInferenceHeaders
	GetXRequestId() *string
}

type MedicalModelInferenceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The business AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-ak-87019a070b49c9fc
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// The random string for the current request.
	//
	// This parameter is required.
	//
	// example:
	//
	// b46a9f08-fe98-4dd3-8885-405c4d032e02
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// The format of the request and response fields.
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
	// 234ad5ad03406bc34b605b3d59746563420d58cb4cdcd2979b677790551a7278
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

func (s MedicalModelInferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s MedicalModelInferenceHeaders) GoString() string {
	return s.String()
}

func (s *MedicalModelInferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MedicalModelInferenceHeaders) GetXGWAccess() *string {
	return s.XGWAccess
}

func (s *MedicalModelInferenceHeaders) GetXGWNonce() *string {
	return s.XGWNonce
}

func (s *MedicalModelInferenceHeaders) GetXGWProtocol() *string {
	return s.XGWProtocol
}

func (s *MedicalModelInferenceHeaders) GetXGWSignature() *string {
	return s.XGWSignature
}

func (s *MedicalModelInferenceHeaders) GetXGWTimestamp() *string {
	return s.XGWTimestamp
}

func (s *MedicalModelInferenceHeaders) GetXRequestId() *string {
	return s.XRequestId
}

func (s *MedicalModelInferenceHeaders) SetCommonHeaders(v map[string]*string) *MedicalModelInferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MedicalModelInferenceHeaders) SetXGWAccess(v string) *MedicalModelInferenceHeaders {
	s.XGWAccess = &v
	return s
}

func (s *MedicalModelInferenceHeaders) SetXGWNonce(v string) *MedicalModelInferenceHeaders {
	s.XGWNonce = &v
	return s
}

func (s *MedicalModelInferenceHeaders) SetXGWProtocol(v string) *MedicalModelInferenceHeaders {
	s.XGWProtocol = &v
	return s
}

func (s *MedicalModelInferenceHeaders) SetXGWSignature(v string) *MedicalModelInferenceHeaders {
	s.XGWSignature = &v
	return s
}

func (s *MedicalModelInferenceHeaders) SetXGWTimestamp(v string) *MedicalModelInferenceHeaders {
	s.XGWTimestamp = &v
	return s
}

func (s *MedicalModelInferenceHeaders) SetXRequestId(v string) *MedicalModelInferenceHeaders {
	s.XRequestId = &v
	return s
}

func (s *MedicalModelInferenceHeaders) Validate() error {
	return dara.Validate(s)
}
