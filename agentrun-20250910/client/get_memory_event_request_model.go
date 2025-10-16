// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v int64) *GetMemoryEventRequest
	GetFrom() *int64
	SetTo(v int64) *GetMemoryEventRequest
	GetTo() *int64
}

type GetMemoryEventRequest struct {
	// example:
	//
	// 1758273080
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// example:
	//
	// 1758273680
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetMemoryEventRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryEventRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryEventRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetMemoryEventRequest) GetTo() *int64 {
	return s.To
}

func (s *GetMemoryEventRequest) SetFrom(v int64) *GetMemoryEventRequest {
	s.From = &v
	return s
}

func (s *GetMemoryEventRequest) SetTo(v int64) *GetMemoryEventRequest {
	s.To = &v
	return s
}

func (s *GetMemoryEventRequest) Validate() error {
	return dara.Validate(s)
}
