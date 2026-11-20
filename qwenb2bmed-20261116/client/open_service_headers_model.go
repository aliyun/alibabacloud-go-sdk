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
	// The business AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-ak-87019a070b49c9fc
	XGWAccess *string `json:"X-GW-Access,omitempty" xml:"X-GW-Access,omitempty"`
	// The random string for the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// b46a9f08-fe98-4dd3-8885-405c4d032eff
	XGWNonce *string `json:"X-GW-Nonce,omitempty" xml:"X-GW-Nonce,omitempty"`
	// The business request signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe3241f17796f523d9df27138f9a22abaead6cc9d1b82bae4b45eedc0e1858d7
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
