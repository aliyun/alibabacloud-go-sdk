// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSettledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SettledCreateCmd) *CreateSettledRequest
	GetBody() *SettledCreateCmd
}

type CreateSettledRequest struct {
	Body *SettledCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSettledRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSettledRequest) GoString() string {
	return s.String()
}

func (s *CreateSettledRequest) GetBody() *SettledCreateCmd {
	return s.Body
}

func (s *CreateSettledRequest) SetBody(v *SettledCreateCmd) *CreateSettledRequest {
	s.Body = v
	return s
}

func (s *CreateSettledRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
