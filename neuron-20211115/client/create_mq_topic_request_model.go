// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMqTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MqTopic) *CreateMqTopicRequest
	GetBody() *MqTopic
}

type CreateMqTopicRequest struct {
	// This parameter is required.
	Body *MqTopic `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMqTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMqTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateMqTopicRequest) GetBody() *MqTopic {
	return s.Body
}

func (s *CreateMqTopicRequest) SetBody(v *MqTopic) *CreateMqTopicRequest {
	s.Body = v
	return s
}

func (s *CreateMqTopicRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
