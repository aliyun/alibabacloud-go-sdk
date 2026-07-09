// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluationRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEvaluationRunResponseBody
	GetRequestId() *string
}

type DeleteEvaluationRunResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEvaluationRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluationRunResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEvaluationRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEvaluationRunResponseBody) SetRequestId(v string) *DeleteEvaluationRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEvaluationRunResponseBody) Validate() error {
	return dara.Validate(s)
}
