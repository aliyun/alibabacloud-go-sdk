// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMqGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MqGroupCreateCmd) *CreateMqGroupRequest
	GetBody() *MqGroupCreateCmd
}

type CreateMqGroupRequest struct {
	Body *MqGroupCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMqGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMqGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMqGroupRequest) GetBody() *MqGroupCreateCmd {
	return s.Body
}

func (s *CreateMqGroupRequest) SetBody(v *MqGroupCreateCmd) *CreateMqGroupRequest {
	s.Body = v
	return s
}

func (s *CreateMqGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
