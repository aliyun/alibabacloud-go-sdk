// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLicenseResponse
	GetStatusCode() *int32
	SetBody(v *GetLicenseResponseBody) *GetLicenseResponse
	GetBody() *GetLicenseResponseBody
}

type GetLicenseResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLicenseResponse) GetBody() *GetLicenseResponseBody {
	return s.Body
}

func (s *GetLicenseResponse) SetHeaders(v map[string]*string) *GetLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseResponse) SetStatusCode(v int32) *GetLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicenseResponse) SetBody(v *GetLicenseResponseBody) *GetLicenseResponse {
	s.Body = v
	return s
}

func (s *GetLicenseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
