// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAutomaticFrequencyControlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *SetAutomaticFrequencyControlConfigResponseBody
	GetActionType() *string
	SetEnable(v string) *SetAutomaticFrequencyControlConfigResponseBody
	GetEnable() *string
	SetLevel(v string) *SetAutomaticFrequencyControlConfigResponseBody
	GetLevel() *string
	SetRequestId(v string) *SetAutomaticFrequencyControlConfigResponseBody
	GetRequestId() *string
}

type SetAutomaticFrequencyControlConfigResponseBody struct {
	// example:
	//
	// js
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAutomaticFrequencyControlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAutomaticFrequencyControlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) GetActionType() *string {
	return s.ActionType
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) GetLevel() *string {
	return s.Level
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) SetActionType(v string) *SetAutomaticFrequencyControlConfigResponseBody {
	s.ActionType = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) SetEnable(v string) *SetAutomaticFrequencyControlConfigResponseBody {
	s.Enable = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) SetLevel(v string) *SetAutomaticFrequencyControlConfigResponseBody {
	s.Level = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) SetRequestId(v string) *SetAutomaticFrequencyControlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAutomaticFrequencyControlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
