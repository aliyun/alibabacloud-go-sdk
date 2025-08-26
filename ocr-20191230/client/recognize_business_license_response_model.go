// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeBusinessLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeBusinessLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeBusinessLicenseResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse
	GetBody() *RecognizeBusinessLicenseResponseBody
}

type RecognizeBusinessLicenseResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusinessLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeBusinessLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeBusinessLicenseResponse) GetBody() *RecognizeBusinessLicenseResponseBody {
	return s.Body
}

func (s *RecognizeBusinessLicenseResponse) SetHeaders(v map[string]*string) *RecognizeBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetStatusCode(v int32) *RecognizeBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse {
	s.Body = v
	return s
}

func (s *RecognizeBusinessLicenseResponse) Validate() error {
	return dara.Validate(s)
}
