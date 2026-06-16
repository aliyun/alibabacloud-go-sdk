// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIMBotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateIMBotInput) *CreateIMBotRequest
	GetBody() *CreateIMBotInput
}

type CreateIMBotRequest struct {
	// The request body.
	//
	// This parameter is required.
	Body *CreateIMBotInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIMBotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIMBotRequest) GoString() string {
	return s.String()
}

func (s *CreateIMBotRequest) GetBody() *CreateIMBotInput {
	return s.Body
}

func (s *CreateIMBotRequest) SetBody(v *CreateIMBotInput) *CreateIMBotRequest {
	s.Body = v
	return s
}

func (s *CreateIMBotRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
