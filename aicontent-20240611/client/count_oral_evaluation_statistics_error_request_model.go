// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountOralEvaluationStatisticsErrorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OralEvaluationStatisticsErrorCountRequest) *CountOralEvaluationStatisticsErrorRequest
	GetBody() *OralEvaluationStatisticsErrorCountRequest
}

type CountOralEvaluationStatisticsErrorRequest struct {
	Body *OralEvaluationStatisticsErrorCountRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsErrorRequest) String() string {
	return dara.Prettify(s)
}

func (s CountOralEvaluationStatisticsErrorRequest) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsErrorRequest) GetBody() *OralEvaluationStatisticsErrorCountRequest {
	return s.Body
}

func (s *CountOralEvaluationStatisticsErrorRequest) SetBody(v *OralEvaluationStatisticsErrorCountRequest) *CountOralEvaluationStatisticsErrorRequest {
	s.Body = v
	return s
}

func (s *CountOralEvaluationStatisticsErrorRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
