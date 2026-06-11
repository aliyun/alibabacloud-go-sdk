// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFormatted(v bool) *GetContextRequest
	GetFormatted() *bool
}

type GetContextRequest struct {
	// Whether to return the context in a formatted structure. Valid values: `true` and `false`. Default value: `false`.
	//
	// example:
	//
	// true
	Formatted *bool `json:"formatted,omitempty" xml:"formatted,omitempty"`
}

func (s GetContextRequest) String() string {
	return dara.Prettify(s)
}

func (s GetContextRequest) GoString() string {
	return s.String()
}

func (s *GetContextRequest) GetFormatted() *bool {
	return s.Formatted
}

func (s *GetContextRequest) SetFormatted(v bool) *GetContextRequest {
	s.Formatted = &v
	return s
}

func (s *GetContextRequest) Validate() error {
	return dara.Validate(s)
}
