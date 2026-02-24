// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEvaluationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutEvaluationsResponseBody
	GetRequestId() *string
	SetResult(v bool) *PutEvaluationsResponseBody
	GetResult() *bool
}

type PutEvaluationsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PutEvaluationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutEvaluationsResponseBody) GoString() string {
	return s.String()
}

func (s *PutEvaluationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutEvaluationsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PutEvaluationsResponseBody) SetRequestId(v string) *PutEvaluationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEvaluationsResponseBody) SetResult(v bool) *PutEvaluationsResponseBody {
	s.Result = &v
	return s
}

func (s *PutEvaluationsResponseBody) Validate() error {
	return dara.Validate(s)
}
