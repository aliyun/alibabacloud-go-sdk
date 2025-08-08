// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateLicenseResponse
	GetStatusCode() *int32
	SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse
	GetBody() *ActivateLicenseResponseBody
}

type ActivateLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateLicenseResponse) GetBody() *ActivateLicenseResponseBody {
	return s.Body
}

func (s *ActivateLicenseResponse) SetHeaders(v map[string]*string) *ActivateLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActivateLicenseResponse) SetStatusCode(v int32) *ActivateLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateLicenseResponse) SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse {
	s.Body = v
	return s
}

func (s *ActivateLicenseResponse) Validate() error {
	return dara.Validate(s)
}
