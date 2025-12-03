// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIosPayload interface {
	dara.Model
	String() string
	GoString() string
	SetAps(v *Aps) *IosPayload
	GetAps() *Aps
	SetExtra(v map[string]interface{}) *IosPayload
	GetExtra() map[string]interface{}
}

type IosPayload struct {
	Aps   *Aps                   `json:"aps,omitempty" xml:"aps,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s IosPayload) String() string {
	return dara.Prettify(s)
}

func (s IosPayload) GoString() string {
	return s.String()
}

func (s *IosPayload) GetAps() *Aps {
	return s.Aps
}

func (s *IosPayload) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *IosPayload) SetAps(v *Aps) *IosPayload {
	s.Aps = v
	return s
}

func (s *IosPayload) SetExtra(v map[string]interface{}) *IosPayload {
	s.Extra = v
	return s
}

func (s *IosPayload) Validate() error {
	if s.Aps != nil {
		if err := s.Aps.Validate(); err != nil {
			return err
		}
	}
	return nil
}
