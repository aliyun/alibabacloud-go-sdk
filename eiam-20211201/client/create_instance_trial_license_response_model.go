// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceTrialLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceTrialLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceTrialLicenseResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceTrialLicenseResponseBody) *CreateInstanceTrialLicenseResponse
	GetBody() *CreateInstanceTrialLicenseResponseBody
}

type CreateInstanceTrialLicenseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceTrialLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceTrialLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceTrialLicenseResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceTrialLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceTrialLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceTrialLicenseResponse) GetBody() *CreateInstanceTrialLicenseResponseBody {
	return s.Body
}

func (s *CreateInstanceTrialLicenseResponse) SetHeaders(v map[string]*string) *CreateInstanceTrialLicenseResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceTrialLicenseResponse) SetStatusCode(v int32) *CreateInstanceTrialLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceTrialLicenseResponse) SetBody(v *CreateInstanceTrialLicenseResponseBody) *CreateInstanceTrialLicenseResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceTrialLicenseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
