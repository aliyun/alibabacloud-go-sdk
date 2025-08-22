// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTaskResponseBody
	GetRequestId() *string
}

type DeleteTaskResponseBody struct {
	// example:
	//
	// 73B38F77-62BA-5878-8952-530DFE21C93B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskResponseBody) SetRequestId(v string) *DeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
