// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTextTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *TextTaskCreateCmd) *CreateTextTaskRequest
	GetBody() *TextTaskCreateCmd
}

type CreateTextTaskRequest struct {
	Body *TextTaskCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTextTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTextTaskRequest) GetBody() *TextTaskCreateCmd {
	return s.Body
}

func (s *CreateTextTaskRequest) SetBody(v *TextTaskCreateCmd) *CreateTextTaskRequest {
	s.Body = v
	return s
}

func (s *CreateTextTaskRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
