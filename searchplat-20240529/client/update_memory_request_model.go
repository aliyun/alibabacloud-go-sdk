// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemory(v string) *UpdateMemoryRequest
	GetMemory() *string
}

type UpdateMemoryRequest struct {
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetMemory() *string {
	return s.Memory
}

func (s *UpdateMemoryRequest) SetMemory(v string) *UpdateMemoryRequest {
	s.Memory = &v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
