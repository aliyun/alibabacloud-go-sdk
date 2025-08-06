// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFreeLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFreeLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFreeLicenseResponse
	GetStatusCode() *int32
	SetBody(v *AddFreeLicenseResponseBody) *AddFreeLicenseResponse
	GetBody() *AddFreeLicenseResponseBody
}

type AddFreeLicenseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFreeLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFreeLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFreeLicenseResponse) GoString() string {
	return s.String()
}

func (s *AddFreeLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFreeLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFreeLicenseResponse) GetBody() *AddFreeLicenseResponseBody {
	return s.Body
}

func (s *AddFreeLicenseResponse) SetHeaders(v map[string]*string) *AddFreeLicenseResponse {
	s.Headers = v
	return s
}

func (s *AddFreeLicenseResponse) SetStatusCode(v int32) *AddFreeLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFreeLicenseResponse) SetBody(v *AddFreeLicenseResponseBody) *AddFreeLicenseResponse {
	s.Body = v
	return s
}

func (s *AddFreeLicenseResponse) Validate() error {
	return dara.Validate(s)
}
