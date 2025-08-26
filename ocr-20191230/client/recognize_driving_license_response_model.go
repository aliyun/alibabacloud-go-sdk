// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeDrivingLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeDrivingLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeDrivingLicenseResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse
	GetBody() *RecognizeDrivingLicenseResponseBody
}

type RecognizeDrivingLicenseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeDrivingLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDrivingLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeDrivingLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeDrivingLicenseResponse) GetBody() *RecognizeDrivingLicenseResponseBody {
	return s.Body
}

func (s *RecognizeDrivingLicenseResponse) SetHeaders(v map[string]*string) *RecognizeDrivingLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetStatusCode(v int32) *RecognizeDrivingLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse {
	s.Body = v
	return s
}

func (s *RecognizeDrivingLicenseResponse) Validate() error {
	return dara.Validate(s)
}
