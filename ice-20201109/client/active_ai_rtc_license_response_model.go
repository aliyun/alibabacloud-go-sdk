// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveAiRtcLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActiveAiRtcLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActiveAiRtcLicenseResponse
	GetStatusCode() *int32
	SetBody(v *ActiveAiRtcLicenseResponseBody) *ActiveAiRtcLicenseResponse
	GetBody() *ActiveAiRtcLicenseResponseBody
}

type ActiveAiRtcLicenseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveAiRtcLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActiveAiRtcLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s ActiveAiRtcLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActiveAiRtcLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActiveAiRtcLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActiveAiRtcLicenseResponse) GetBody() *ActiveAiRtcLicenseResponseBody {
	return s.Body
}

func (s *ActiveAiRtcLicenseResponse) SetHeaders(v map[string]*string) *ActiveAiRtcLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActiveAiRtcLicenseResponse) SetStatusCode(v int32) *ActiveAiRtcLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveAiRtcLicenseResponse) SetBody(v *ActiveAiRtcLicenseResponseBody) *ActiveAiRtcLicenseResponse {
	s.Body = v
	return s
}

func (s *ActiveAiRtcLicenseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
