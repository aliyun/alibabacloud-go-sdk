// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertEvaluationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevertEvaluationResultsResponseBody
	GetRequestId() *string
}

type RevertEvaluationResultsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 01ACCBF2-0B0B-59F2-9E84-07B38267BCA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertEvaluationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevertEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevertEvaluationResultsResponseBody) SetRequestId(v string) *RevertEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertEvaluationResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
