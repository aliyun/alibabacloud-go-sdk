// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumComputeMetricsByUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumComputeMetricsByUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumComputeMetricsByUsageResponse
	GetStatusCode() *int32
	SetBody(v *SumComputeMetricsByUsageResponseBody) *SumComputeMetricsByUsageResponse
	GetBody() *SumComputeMetricsByUsageResponseBody
}

type SumComputeMetricsByUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumComputeMetricsByUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumComputeMetricsByUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByUsageResponse) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumComputeMetricsByUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumComputeMetricsByUsageResponse) GetBody() *SumComputeMetricsByUsageResponseBody {
	return s.Body
}

func (s *SumComputeMetricsByUsageResponse) SetHeaders(v map[string]*string) *SumComputeMetricsByUsageResponse {
	s.Headers = v
	return s
}

func (s *SumComputeMetricsByUsageResponse) SetStatusCode(v int32) *SumComputeMetricsByUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *SumComputeMetricsByUsageResponse) SetBody(v *SumComputeMetricsByUsageResponseBody) *SumComputeMetricsByUsageResponse {
	s.Body = v
	return s
}

func (s *SumComputeMetricsByUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
