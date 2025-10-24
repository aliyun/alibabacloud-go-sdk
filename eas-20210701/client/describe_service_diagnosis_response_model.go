// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceDiagnosisResponseBody) *DescribeServiceDiagnosisResponse
	GetBody() *DescribeServiceDiagnosisResponseBody
}

type DescribeServiceDiagnosisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceDiagnosisResponse) GetBody() *DescribeServiceDiagnosisResponseBody {
	return s.Body
}

func (s *DescribeServiceDiagnosisResponse) SetHeaders(v map[string]*string) *DescribeServiceDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceDiagnosisResponse) SetStatusCode(v int32) *DescribeServiceDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceDiagnosisResponse) SetBody(v *DescribeServiceDiagnosisResponseBody) *DescribeServiceDiagnosisResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
