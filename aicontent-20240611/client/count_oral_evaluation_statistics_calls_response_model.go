// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsCallsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsCallsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CountOralEvaluationStatisticsCallsResponse
	GetStatusCode() *int32
	SetBody(v *CountOralEvaluationStatisticsCallsResponseBody) *CountOralEvaluationStatisticsCallsResponse
	GetBody() *CountOralEvaluationStatisticsCallsResponseBody
}

type CountOralEvaluationStatisticsCallsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOralEvaluationStatisticsCallsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsCallsResponse) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsCallsResponse) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsCallsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CountOralEvaluationStatisticsCallsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CountOralEvaluationStatisticsCallsResponse) GetBody() *CountOralEvaluationStatisticsCallsResponseBody {
	return s.Body
}

func (s *CountOralEvaluationStatisticsCallsResponse) SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsCallsResponse {
	s.Headers = v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponse) SetStatusCode(v int32) *CountOralEvaluationStatisticsCallsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponse) SetBody(v *CountOralEvaluationStatisticsCallsResponseBody) *CountOralEvaluationStatisticsCallsResponse {
	s.Body = v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
