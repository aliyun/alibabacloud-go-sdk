// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ComponentUpdateCmd) *UpdateComponentRequest
	GetBody() *ComponentUpdateCmd
}

type UpdateComponentRequest struct {
	// This parameter is required.
	Body *ComponentUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComponentRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentRequest) GetBody() *ComponentUpdateCmd {
	return s.Body
}

func (s *UpdateComponentRequest) SetBody(v *ComponentUpdateCmd) *UpdateComponentRequest {
	s.Body = v
	return s
}

func (s *UpdateComponentRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
