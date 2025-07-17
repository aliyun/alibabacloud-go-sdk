// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMemoryNodeResponseBody
	GetRequestId() *string
}

type DeleteMemoryNodeResponseBody struct {
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoryNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryNodeResponseBody) SetRequestId(v string) *DeleteMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
