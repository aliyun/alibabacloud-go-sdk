// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainMqConsoleLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ObtainMqConsoleLinkCmd) *ObtainMqConsoleLinkRequest
	GetBody() *ObtainMqConsoleLinkCmd
}

type ObtainMqConsoleLinkRequest struct {
	Body *ObtainMqConsoleLinkCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainMqConsoleLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainMqConsoleLinkRequest) GoString() string {
	return s.String()
}

func (s *ObtainMqConsoleLinkRequest) GetBody() *ObtainMqConsoleLinkCmd {
	return s.Body
}

func (s *ObtainMqConsoleLinkRequest) SetBody(v *ObtainMqConsoleLinkCmd) *ObtainMqConsoleLinkRequest {
	s.Body = v
	return s
}

func (s *ObtainMqConsoleLinkRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
