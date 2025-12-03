// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAndroidShortPayload interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *AndroidShortPayloadBody) *AndroidShortPayload
	GetBody() *AndroidShortPayloadBody
	SetExtra(v map[string]interface{}) *AndroidShortPayload
	GetExtra() map[string]interface{}
}

type AndroidShortPayload struct {
	Body  *AndroidShortPayloadBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	Extra map[string]interface{}   `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AndroidShortPayload) String() string {
	return dara.Prettify(s)
}

func (s AndroidShortPayload) GoString() string {
	return s.String()
}

func (s *AndroidShortPayload) GetBody() *AndroidShortPayloadBody {
	return s.Body
}

func (s *AndroidShortPayload) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *AndroidShortPayload) SetBody(v *AndroidShortPayloadBody) *AndroidShortPayload {
	s.Body = v
	return s
}

func (s *AndroidShortPayload) SetExtra(v map[string]interface{}) *AndroidShortPayload {
	s.Extra = v
	return s
}

func (s *AndroidShortPayload) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AndroidShortPayloadBody struct {
	Custom *string `json:"custom,omitempty" xml:"custom,omitempty"`
}

func (s AndroidShortPayloadBody) String() string {
	return dara.Prettify(s)
}

func (s AndroidShortPayloadBody) GoString() string {
	return s.String()
}

func (s *AndroidShortPayloadBody) GetCustom() *string {
	return s.Custom
}

func (s *AndroidShortPayloadBody) SetCustom(v string) *AndroidShortPayloadBody {
	s.Custom = &v
	return s
}

func (s *AndroidShortPayloadBody) Validate() error {
	return dara.Validate(s)
}
