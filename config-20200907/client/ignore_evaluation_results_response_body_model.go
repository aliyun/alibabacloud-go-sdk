// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *IgnoreEvaluationResultsResponseBody
	GetRequestId() *string
}

type IgnoreEvaluationResultsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1840CBF2-0B0B-59F2-9E84-07B38267A279
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IgnoreEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IgnoreEvaluationResultsResponseBody) SetRequestId(v string) *IgnoreEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *IgnoreEvaluationResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
