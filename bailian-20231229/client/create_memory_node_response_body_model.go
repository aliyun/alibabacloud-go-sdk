// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMemoryNodeId(v string) *CreateMemoryNodeResponseBody
	GetMemoryNodeId() *string
	SetRequestId(v string) *CreateMemoryNodeResponseBody
	GetRequestId() *string
}

type CreateMemoryNodeResponseBody struct {
	// example:
	//
	// 68de06c95368463a8be4a84efc872cc5
	MemoryNodeId *string `json:"memoryNodeId,omitempty" xml:"memoryNodeId,omitempty"`
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryNodeResponseBody) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *CreateMemoryNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryNodeResponseBody) SetMemoryNodeId(v string) *CreateMemoryNodeResponseBody {
	s.MemoryNodeId = &v
	return s
}

func (s *CreateMemoryNodeResponseBody) SetRequestId(v string) *CreateMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
