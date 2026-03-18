// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsCallsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OralEvaluationStatisticsCallsCountRequest) *CountOralEvaluationStatisticsCallsRequest
	GetBody() *OralEvaluationStatisticsCallsCountRequest
}

type CountOralEvaluationStatisticsCallsRequest struct {
	Body *OralEvaluationStatisticsCallsCountRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsCallsRequest) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsCallsRequest) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsCallsRequest) GetBody() *OralEvaluationStatisticsCallsCountRequest {
	return s.Body
}

func (s *CountOralEvaluationStatisticsCallsRequest) SetBody(v *OralEvaluationStatisticsCallsCountRequest) *CountOralEvaluationStatisticsCallsRequest {
	s.Body = v
	return s
}

func (s *CountOralEvaluationStatisticsCallsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
