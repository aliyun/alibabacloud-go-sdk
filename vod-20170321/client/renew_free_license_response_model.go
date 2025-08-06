// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewFreeLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewFreeLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewFreeLicenseResponse
	GetStatusCode() *int32
	SetBody(v *RenewFreeLicenseResponseBody) *RenewFreeLicenseResponse
	GetBody() *RenewFreeLicenseResponseBody
}

type RenewFreeLicenseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewFreeLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewFreeLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewFreeLicenseResponse) GoString() string {
	return s.String()
}

func (s *RenewFreeLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewFreeLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewFreeLicenseResponse) GetBody() *RenewFreeLicenseResponseBody {
	return s.Body
}

func (s *RenewFreeLicenseResponse) SetHeaders(v map[string]*string) *RenewFreeLicenseResponse {
	s.Headers = v
	return s
}

func (s *RenewFreeLicenseResponse) SetStatusCode(v int32) *RenewFreeLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewFreeLicenseResponse) SetBody(v *RenewFreeLicenseResponseBody) *RenewFreeLicenseResponse {
	s.Body = v
	return s
}

func (s *RenewFreeLicenseResponse) Validate() error {
	return dara.Validate(s)
}
