// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluationRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEvaluationRunResponseBody
	GetRequestId() *string
}

type UpdateEvaluationRunResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEvaluationRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluationRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEvaluationRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEvaluationRunResponseBody) SetRequestId(v string) *UpdateEvaluationRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEvaluationRunResponseBody) Validate() error {
	return dara.Validate(s)
}
