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
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbd5fc3
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
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
