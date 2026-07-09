// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEvaluatorResponseBody
	GetRequestId() *string
}

type DeleteEvaluatorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteEvaluatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluatorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEvaluatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEvaluatorResponseBody) SetRequestId(v string) *DeleteEvaluatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEvaluatorResponseBody) Validate() error {
	return dara.Validate(s)
}
