// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutomaticFrequencyControlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *GetAutomaticFrequencyControlConfigResponseBody
	GetActionType() *string
	SetEnable(v string) *GetAutomaticFrequencyControlConfigResponseBody
	GetEnable() *string
	SetInterval(v int32) *GetAutomaticFrequencyControlConfigResponseBody
	GetInterval() *int32
	SetLevel(v string) *GetAutomaticFrequencyControlConfigResponseBody
	GetLevel() *string
	SetPunishTime(v int32) *GetAutomaticFrequencyControlConfigResponseBody
	GetPunishTime() *int32
	SetRequestId(v string) *GetAutomaticFrequencyControlConfigResponseBody
	GetRequestId() *string
	SetRuleId(v int64) *GetAutomaticFrequencyControlConfigResponseBody
	GetRuleId() *int64
	SetThreshold(v int32) *GetAutomaticFrequencyControlConfigResponseBody
	GetThreshold() *int32
}

type GetAutomaticFrequencyControlConfigResponseBody struct {
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
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// normal
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 20000000
	PunishTime *int32 `json:"PunishTime,omitempty" xml:"PunishTime,omitempty"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 11957665
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 100
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s GetAutomaticFrequencyControlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutomaticFrequencyControlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetActionType() *string {
	return s.ActionType
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetLevel() *string {
	return s.Level
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetPunishTime() *int32 {
	return s.PunishTime
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) GetThreshold() *int32 {
	return s.Threshold
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetActionType(v string) *GetAutomaticFrequencyControlConfigResponseBody {
	s.ActionType = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetEnable(v string) *GetAutomaticFrequencyControlConfigResponseBody {
	s.Enable = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetInterval(v int32) *GetAutomaticFrequencyControlConfigResponseBody {
	s.Interval = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetLevel(v string) *GetAutomaticFrequencyControlConfigResponseBody {
	s.Level = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetPunishTime(v int32) *GetAutomaticFrequencyControlConfigResponseBody {
	s.PunishTime = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetRequestId(v string) *GetAutomaticFrequencyControlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetRuleId(v int64) *GetAutomaticFrequencyControlConfigResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) SetThreshold(v int32) *GetAutomaticFrequencyControlConfigResponseBody {
	s.Threshold = &v
	return s
}

func (s *GetAutomaticFrequencyControlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
