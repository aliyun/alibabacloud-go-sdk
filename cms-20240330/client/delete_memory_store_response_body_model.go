// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMemoryStoreResponseBody
	GetRequestId() *string
}

type DeleteMemoryStoreResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryStoreResponseBody) SetRequestId(v string) *DeleteMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
