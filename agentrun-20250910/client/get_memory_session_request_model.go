// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemorySessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetMemorySessionRequest
	GetFrom() *int64
	SetSize(v int32) *GetMemorySessionRequest
	GetSize() *int32
	SetTo(v int64) *GetMemorySessionRequest
	GetTo() *int64
}

type GetMemorySessionRequest struct {
	// example:
	//
	// 1736558346
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 1736561898
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetMemorySessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySessionRequest) GoString() string {
	return s.String()
}

func (s *GetMemorySessionRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetMemorySessionRequest) GetSize() *int32 {
	return s.Size
}

func (s *GetMemorySessionRequest) GetTo() *int64 {
	return s.To
}

func (s *GetMemorySessionRequest) SetFrom(v int64) *GetMemorySessionRequest {
	s.From = &v
	return s
}

func (s *GetMemorySessionRequest) SetSize(v int32) *GetMemorySessionRequest {
	s.Size = &v
	return s
}

func (s *GetMemorySessionRequest) SetTo(v int64) *GetMemorySessionRequest {
	s.To = &v
	return s
}

func (s *GetMemorySessionRequest) Validate() error {
	return dara.Validate(s)
}
