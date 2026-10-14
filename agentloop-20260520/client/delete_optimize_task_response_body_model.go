// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOptimizeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOptimizeTaskResponseBody
	GetRequestId() *string
}

type DeleteOptimizeTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteOptimizeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOptimizeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOptimizeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOptimizeTaskResponseBody) SetRequestId(v string) *DeleteOptimizeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOptimizeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
