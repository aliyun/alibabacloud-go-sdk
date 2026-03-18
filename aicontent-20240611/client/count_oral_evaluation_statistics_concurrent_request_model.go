// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsConcurrentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OralEvaluationStatisticsConcurrentCountRequest) *CountOralEvaluationStatisticsConcurrentRequest
	GetBody() *OralEvaluationStatisticsConcurrentCountRequest
}

type CountOralEvaluationStatisticsConcurrentRequest struct {
	Body *OralEvaluationStatisticsConcurrentCountRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsConcurrentRequest) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsConcurrentRequest) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsConcurrentRequest) GetBody() *OralEvaluationStatisticsConcurrentCountRequest {
	return s.Body
}

func (s *CountOralEvaluationStatisticsConcurrentRequest) SetBody(v *OralEvaluationStatisticsConcurrentCountRequest) *CountOralEvaluationStatisticsConcurrentRequest {
	s.Body = v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
