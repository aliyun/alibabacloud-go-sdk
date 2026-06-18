// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMemoryId(v string) *CreateMemoryResponseBody
	GetMemoryId() *string
	SetRequestId(v string) *CreateMemoryResponseBody
	GetRequestId() *string
}

type CreateMemoryResponseBody struct {
	// The long-term memory ID.
	//
	// > Store this value properly. It is required for all subsequent API operations related to this long-term memory.
	//
	// >.
	//
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbxxxx
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A--2446A84821CA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBody) GetMemoryId() *string {
	return s.MemoryId
}

func (s *CreateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryResponseBody) SetMemoryId(v string) *CreateMemoryResponseBody {
	s.MemoryId = &v
	return s
}

func (s *CreateMemoryResponseBody) SetRequestId(v string) *CreateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
