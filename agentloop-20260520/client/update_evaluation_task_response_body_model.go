// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEvaluationTaskResponseBody
	GetRequestId() *string
}

type UpdateEvaluationTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEvaluationTaskResponseBody) SetRequestId(v string) *UpdateEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
