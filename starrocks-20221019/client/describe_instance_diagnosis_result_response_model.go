// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDiagnosisResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDiagnosisResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDiagnosisResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDiagnosisResultResponseBody) *DescribeInstanceDiagnosisResultResponse
	GetBody() *DescribeInstanceDiagnosisResultResponseBody
}

type DescribeInstanceDiagnosisResultResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDiagnosisResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDiagnosisResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDiagnosisResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDiagnosisResultResponse) GetBody() *DescribeInstanceDiagnosisResultResponseBody {
	return s.Body
}

func (s *DescribeInstanceDiagnosisResultResponse) SetHeaders(v map[string]*string) *DescribeInstanceDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponse) SetStatusCode(v int32) *DescribeInstanceDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponse) SetBody(v *DescribeInstanceDiagnosisResultResponseBody) *DescribeInstanceDiagnosisResultResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDiagnosisResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
