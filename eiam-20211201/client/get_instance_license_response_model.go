// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceLicenseResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceLicenseResponseBody) *GetInstanceLicenseResponse
	GetBody() *GetInstanceLicenseResponseBody
}

type GetInstanceLicenseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceLicenseResponse) GetBody() *GetInstanceLicenseResponseBody {
	return s.Body
}

func (s *GetInstanceLicenseResponse) SetHeaders(v map[string]*string) *GetInstanceLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceLicenseResponse) SetStatusCode(v int32) *GetInstanceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceLicenseResponse) SetBody(v *GetInstanceLicenseResponseBody) *GetInstanceLicenseResponse {
	s.Body = v
	return s
}

func (s *GetInstanceLicenseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
