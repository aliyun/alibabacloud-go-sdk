// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMemoryStoreResponseBody
	GetRequestId() *string
}

type UpdateMemoryStoreResponseBody struct {
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryStoreResponseBody) SetRequestId(v string) *UpdateMemoryStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryStoreResponseBody) Validate() error {
	return dara.Validate(s)
}
