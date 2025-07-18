// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisMonitorPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosisMonitorPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosisMonitorPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosisMonitorPerformanceResponseBody) *DescribeDiagnosisMonitorPerformanceResponse
	GetBody() *DescribeDiagnosisMonitorPerformanceResponseBody
}

type DescribeDiagnosisMonitorPerformanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosisMonitorPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosisMonitorPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisMonitorPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) GetBody() *DescribeDiagnosisMonitorPerformanceResponseBody {
	return s.Body
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetStatusCode(v int32) *DescribeDiagnosisMonitorPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) SetBody(v *DescribeDiagnosisMonitorPerformanceResponseBody) *DescribeDiagnosisMonitorPerformanceResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosisMonitorPerformanceResponse) Validate() error {
	return dara.Validate(s)
}
