// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v string) *DebugModelRequest
	GetInput() *string
	SetIsOnline(v string) *DebugModelRequest
	GetIsOnline() *string
}

type DebugModelRequest struct {
	Input *string `json:"input,omitempty" xml:"input,omitempty"`
	// example:
	//
	// true
	IsOnline *string `json:"isOnline,omitempty" xml:"isOnline,omitempty"`
}

func (s DebugModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugModelRequest) GoString() string {
	return s.String()
}

func (s *DebugModelRequest) GetInput() *string {
	return s.Input
}

func (s *DebugModelRequest) GetIsOnline() *string {
	return s.IsOnline
}

func (s *DebugModelRequest) SetInput(v string) *DebugModelRequest {
	s.Input = &v
	return s
}

func (s *DebugModelRequest) SetIsOnline(v string) *DebugModelRequest {
	s.IsOnline = &v
	return s
}

func (s *DebugModelRequest) Validate() error {
	return dara.Validate(s)
}
