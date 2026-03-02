// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSettledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SettledUpdateCmd) *UpdateSettledRequest
	GetBody() *SettledUpdateCmd
}

type UpdateSettledRequest struct {
	Body *SettledUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSettledRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSettledRequest) GoString() string {
	return s.String()
}

func (s *UpdateSettledRequest) GetBody() *SettledUpdateCmd {
	return s.Body
}

func (s *UpdateSettledRequest) SetBody(v *SettledUpdateCmd) *UpdateSettledRequest {
	s.Body = v
	return s
}

func (s *UpdateSettledRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
