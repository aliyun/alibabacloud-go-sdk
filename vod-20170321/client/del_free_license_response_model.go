// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelFreeLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelFreeLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelFreeLicenseResponse
	GetStatusCode() *int32
	SetBody(v *DelFreeLicenseResponseBody) *DelFreeLicenseResponse
	GetBody() *DelFreeLicenseResponseBody
}

type DelFreeLicenseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelFreeLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelFreeLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s DelFreeLicenseResponse) GoString() string {
	return s.String()
}

func (s *DelFreeLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelFreeLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelFreeLicenseResponse) GetBody() *DelFreeLicenseResponseBody {
	return s.Body
}

func (s *DelFreeLicenseResponse) SetHeaders(v map[string]*string) *DelFreeLicenseResponse {
	s.Headers = v
	return s
}

func (s *DelFreeLicenseResponse) SetStatusCode(v int32) *DelFreeLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *DelFreeLicenseResponse) SetBody(v *DelFreeLicenseResponseBody) *DelFreeLicenseResponse {
	s.Body = v
	return s
}

func (s *DelFreeLicenseResponse) Validate() error {
	return dara.Validate(s)
}
