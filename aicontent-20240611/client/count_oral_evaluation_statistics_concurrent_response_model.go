// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsConcurrentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsConcurrentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CountOralEvaluationStatisticsConcurrentResponse
	GetStatusCode() *int32
	SetBody(v *CountOralEvaluationStatisticsConcurrentResponseBody) *CountOralEvaluationStatisticsConcurrentResponse
	GetBody() *CountOralEvaluationStatisticsConcurrentResponseBody
}

type CountOralEvaluationStatisticsConcurrentResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOralEvaluationStatisticsConcurrentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsConcurrentResponse) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsConcurrentResponse) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) GetBody() *CountOralEvaluationStatisticsConcurrentResponseBody {
	return s.Body
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsConcurrentResponse {
	s.Headers = v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) SetStatusCode(v int32) *CountOralEvaluationStatisticsConcurrentResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) SetBody(v *CountOralEvaluationStatisticsConcurrentResponseBody) *CountOralEvaluationStatisticsConcurrentResponse {
	s.Body = v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
