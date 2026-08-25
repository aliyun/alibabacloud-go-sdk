// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DebugModelRequestBody) *DebugModelRequest
	GetBody() *DebugModelRequestBody
}

type DebugModelRequest struct {
	Body *DebugModelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s DebugModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugModelRequest) GoString() string {
	return s.String()
}

func (s *DebugModelRequest) GetBody() *DebugModelRequestBody {
	return s.Body
}

func (s *DebugModelRequest) SetBody(v *DebugModelRequestBody) *DebugModelRequest {
	s.Body = v
	return s
}

func (s *DebugModelRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DebugModelRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// hello
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
}

func (s DebugModelRequestBody) String() string {
	return dara.Prettify(s)
}

func (s DebugModelRequestBody) GoString() string {
	return s.String()
}

func (s *DebugModelRequestBody) GetPrompt() *string {
	return s.Prompt
}

func (s *DebugModelRequestBody) SetPrompt(v string) *DebugModelRequestBody {
	s.Prompt = &v
	return s
}

func (s *DebugModelRequestBody) Validate() error {
	return dara.Validate(s)
}
