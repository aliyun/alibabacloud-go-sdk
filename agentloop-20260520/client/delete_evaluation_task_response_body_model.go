// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEvaluationTaskResponseBody
	GetRequestId() *string
}

type DeleteEvaluationTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEvaluationTaskResponseBody) SetRequestId(v string) *DeleteEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
