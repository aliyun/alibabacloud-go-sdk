// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEvaluatorResponseBody
	GetRequestId() *string
}

type UpdateEvaluatorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateEvaluatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEvaluatorResponseBody) SetRequestId(v string) *UpdateEvaluatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEvaluatorResponseBody) Validate() error {
	return dara.Validate(s)
}
