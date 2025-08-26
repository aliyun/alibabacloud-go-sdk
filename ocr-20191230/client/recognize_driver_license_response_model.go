// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeDriverLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeDriverLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeDriverLicenseResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeDriverLicenseResponseBody) *RecognizeDriverLicenseResponse
	GetBody() *RecognizeDriverLicenseResponseBody
}

type RecognizeDriverLicenseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeDriverLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDriverLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeDriverLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeDriverLicenseResponse) GetBody() *RecognizeDriverLicenseResponseBody {
	return s.Body
}

func (s *RecognizeDriverLicenseResponse) SetHeaders(v map[string]*string) *RecognizeDriverLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDriverLicenseResponse) SetStatusCode(v int32) *RecognizeDriverLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDriverLicenseResponse) SetBody(v *RecognizeDriverLicenseResponseBody) *RecognizeDriverLicenseResponse {
	s.Body = v
	return s
}

func (s *RecognizeDriverLicenseResponse) Validate() error {
	return dara.Validate(s)
}
