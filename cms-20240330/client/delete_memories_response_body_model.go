// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMemoriesResponseBody
	GetRequestId() *string
}

type DeleteMemoriesResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoriesResponseBody) SetRequestId(v string) *DeleteMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
