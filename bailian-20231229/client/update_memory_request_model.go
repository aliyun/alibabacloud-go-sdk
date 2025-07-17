// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateMemoryRequest
	GetDescription() *string
}

type UpdateMemoryRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryRequest) SetDescription(v string) *UpdateMemoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
