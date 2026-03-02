// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ComponentCreateCmd) *CreateComponentRequest
	GetBody() *ComponentCreateCmd
}

type CreateComponentRequest struct {
	// This parameter is required.
	Body *ComponentCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComponentRequest) GoString() string {
	return s.String()
}

func (s *CreateComponentRequest) GetBody() *ComponentCreateCmd {
	return s.Body
}

func (s *CreateComponentRequest) SetBody(v *ComponentCreateCmd) *CreateComponentRequest {
	s.Body = v
	return s
}

func (s *CreateComponentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
