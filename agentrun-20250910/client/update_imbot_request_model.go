// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIMBotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateIMBotInput) *UpdateIMBotRequest
	GetBody() *UpdateIMBotInput
}

type UpdateIMBotRequest struct {
	// This parameter is required.
	Body *UpdateIMBotInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIMBotRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIMBotRequest) GoString() string {
	return s.String()
}

func (s *UpdateIMBotRequest) GetBody() *UpdateIMBotInput {
	return s.Body
}

func (s *UpdateIMBotRequest) SetBody(v *UpdateIMBotInput) *UpdateIMBotRequest {
	s.Body = v
	return s
}

func (s *UpdateIMBotRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
