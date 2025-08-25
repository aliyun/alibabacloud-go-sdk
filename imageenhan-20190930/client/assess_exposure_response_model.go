// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessExposureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssessExposureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssessExposureResponse
	GetStatusCode() *int32
	SetBody(v *AssessExposureResponseBody) *AssessExposureResponse
	GetBody() *AssessExposureResponseBody
}

type AssessExposureResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssessExposureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssessExposureResponse) String() string {
	return dara.Prettify(s)
}

func (s AssessExposureResponse) GoString() string {
	return s.String()
}

func (s *AssessExposureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssessExposureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssessExposureResponse) GetBody() *AssessExposureResponseBody {
	return s.Body
}

func (s *AssessExposureResponse) SetHeaders(v map[string]*string) *AssessExposureResponse {
	s.Headers = v
	return s
}

func (s *AssessExposureResponse) SetStatusCode(v int32) *AssessExposureResponse {
	s.StatusCode = &v
	return s
}

func (s *AssessExposureResponse) SetBody(v *AssessExposureResponseBody) *AssessExposureResponse {
	s.Body = v
	return s
}

func (s *AssessExposureResponse) Validate() error {
	return dara.Validate(s)
}
