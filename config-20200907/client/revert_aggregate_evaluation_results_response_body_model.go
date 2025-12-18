// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAggregateEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevertAggregateEvaluationResultsResponseBody
	GetRequestId() *string
}

type RevertAggregateEvaluationResultsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BB11CBF2-0B0B-59F2-9E84-07B38267AD12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertAggregateEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevertAggregateEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevertAggregateEvaluationResultsResponseBody) SetRequestId(v string) *RevertAggregateEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
