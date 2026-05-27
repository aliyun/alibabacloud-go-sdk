// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumComputeMetricsByRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumComputeMetricsByRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumComputeMetricsByRecordResponse
	GetStatusCode() *int32
	SetBody(v *SumComputeMetricsByRecordResponseBody) *SumComputeMetricsByRecordResponse
	GetBody() *SumComputeMetricsByRecordResponseBody
}

type SumComputeMetricsByRecordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumComputeMetricsByRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumComputeMetricsByRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByRecordResponse) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumComputeMetricsByRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumComputeMetricsByRecordResponse) GetBody() *SumComputeMetricsByRecordResponseBody {
	return s.Body
}

func (s *SumComputeMetricsByRecordResponse) SetHeaders(v map[string]*string) *SumComputeMetricsByRecordResponse {
	s.Headers = v
	return s
}

func (s *SumComputeMetricsByRecordResponse) SetStatusCode(v int32) *SumComputeMetricsByRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SumComputeMetricsByRecordResponse) SetBody(v *SumComputeMetricsByRecordResponseBody) *SumComputeMetricsByRecordResponse {
	s.Body = v
	return s
}

func (s *SumComputeMetricsByRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
