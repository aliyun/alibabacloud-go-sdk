// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMemoryResponseBody
	GetRequestId() *string
}

type DeleteMemoryResponseBody struct {
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryResponseBody) SetRequestId(v string) *DeleteMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
