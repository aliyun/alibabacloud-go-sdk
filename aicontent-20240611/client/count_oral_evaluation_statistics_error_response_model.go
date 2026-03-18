// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsErrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsErrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CountOralEvaluationStatisticsErrorResponse
	GetStatusCode() *int32
	SetBody(v *CountOralEvaluationStatisticsErrorResponseBody) *CountOralEvaluationStatisticsErrorResponse
	GetBody() *CountOralEvaluationStatisticsErrorResponseBody
}

type CountOralEvaluationStatisticsErrorResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOralEvaluationStatisticsErrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsErrorResponse) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsErrorResponse) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsErrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CountOralEvaluationStatisticsErrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CountOralEvaluationStatisticsErrorResponse) GetBody() *CountOralEvaluationStatisticsErrorResponseBody {
	return s.Body
}

func (s *CountOralEvaluationStatisticsErrorResponse) SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsErrorResponse {
	s.Headers = v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponse) SetStatusCode(v int32) *CountOralEvaluationStatisticsErrorResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponse) SetBody(v *CountOralEvaluationStatisticsErrorResponseBody) *CountOralEvaluationStatisticsErrorResponse {
	s.Body = v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
