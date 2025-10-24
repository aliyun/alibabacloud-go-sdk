// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceInstanceDiagnosisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceInstanceDiagnosisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceInstanceDiagnosisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceInstanceDiagnosisResponseBody) *DescribeServiceInstanceDiagnosisResponse
	GetBody() *DescribeServiceInstanceDiagnosisResponseBody
}

type DescribeServiceInstanceDiagnosisResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceInstanceDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceInstanceDiagnosisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceInstanceDiagnosisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceInstanceDiagnosisResponse) GetBody() *DescribeServiceInstanceDiagnosisResponseBody {
	return s.Body
}

func (s *DescribeServiceInstanceDiagnosisResponse) SetHeaders(v map[string]*string) *DescribeServiceInstanceDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponse) SetStatusCode(v int32) *DescribeServiceInstanceDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponse) SetBody(v *DescribeServiceInstanceDiagnosisResponseBody) *DescribeServiceInstanceDiagnosisResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
