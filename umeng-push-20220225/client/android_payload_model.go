// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAndroidPayload interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Body) *AndroidPayload
	GetBody() *Body
	SetDisplayType(v string) *AndroidPayload
	GetDisplayType() *string
	SetExtra(v map[string]interface{}) *AndroidPayload
	GetExtra() map[string]interface{}
	SetMessage2ThirdChannel(v *Message2ThirdChannel) *AndroidPayload
	GetMessage2ThirdChannel() *Message2ThirdChannel
}

type AndroidPayload struct {
	Body                 *Body                  `json:"body,omitempty" xml:"body,omitempty"`
	DisplayType          *string                `json:"displayType,omitempty" xml:"displayType,omitempty"`
	Extra                map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	Message2ThirdChannel *Message2ThirdChannel  `json:"message2ThirdChannel,omitempty" xml:"message2ThirdChannel,omitempty"`
}

func (s AndroidPayload) String() string {
	return dara.Prettify(s)
}

func (s AndroidPayload) GoString() string {
	return s.String()
}

func (s *AndroidPayload) GetBody() *Body {
	return s.Body
}

func (s *AndroidPayload) GetDisplayType() *string {
	return s.DisplayType
}

func (s *AndroidPayload) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *AndroidPayload) GetMessage2ThirdChannel() *Message2ThirdChannel {
	return s.Message2ThirdChannel
}

func (s *AndroidPayload) SetBody(v *Body) *AndroidPayload {
	s.Body = v
	return s
}

func (s *AndroidPayload) SetDisplayType(v string) *AndroidPayload {
	s.DisplayType = &v
	return s
}

func (s *AndroidPayload) SetExtra(v map[string]interface{}) *AndroidPayload {
	s.Extra = v
	return s
}

func (s *AndroidPayload) SetMessage2ThirdChannel(v *Message2ThirdChannel) *AndroidPayload {
	s.Message2ThirdChannel = v
	return s
}

func (s *AndroidPayload) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Message2ThirdChannel != nil {
		if err := s.Message2ThirdChannel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
