// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreAggregateEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *IgnoreAggregateEvaluationResultsResponseBody
	GetRequestId() *string
}

type IgnoreAggregateEvaluationResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IgnoreAggregateEvaluationResultsResponseBody) SetRequestId(v string) *IgnoreAggregateEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
