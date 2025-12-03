// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHarmonyPayload interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayType(v string) *HarmonyPayload
	GetDisplayType() *string
	SetExtra(v map[string]interface{}) *HarmonyPayload
	GetExtra() map[string]interface{}
	SetHarmonyBody(v *HarmonyBody) *HarmonyPayload
	GetHarmonyBody() *HarmonyBody
}

type HarmonyPayload struct {
	// This parameter is required.
	DisplayType *string                `json:"displayType,omitempty" xml:"displayType,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	HarmonyBody *HarmonyBody           `json:"harmonyBody,omitempty" xml:"harmonyBody,omitempty"`
}

func (s HarmonyPayload) String() string {
	return dara.Prettify(s)
}

func (s HarmonyPayload) GoString() string {
	return s.String()
}

func (s *HarmonyPayload) GetDisplayType() *string {
	return s.DisplayType
}

func (s *HarmonyPayload) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *HarmonyPayload) GetHarmonyBody() *HarmonyBody {
	return s.HarmonyBody
}

func (s *HarmonyPayload) SetDisplayType(v string) *HarmonyPayload {
	s.DisplayType = &v
	return s
}

func (s *HarmonyPayload) SetExtra(v map[string]interface{}) *HarmonyPayload {
	s.Extra = v
	return s
}

func (s *HarmonyPayload) SetHarmonyBody(v *HarmonyBody) *HarmonyPayload {
	s.HarmonyBody = v
	return s
}

func (s *HarmonyPayload) Validate() error {
	if s.HarmonyBody != nil {
		if err := s.HarmonyBody.Validate(); err != nil {
			return err
		}
	}
	return nil
}
