// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMemoryStoreResponseBody
	GetRequestId() *string
}

type CreateMemoryStoreResponseBody struct {
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryStoreResponseBody) SetRequestId(v string) *CreateMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
