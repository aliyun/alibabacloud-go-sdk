// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAutoScalingConfigurationResponseBody
	GetCode() *int64
	SetData(v *GetAutoScalingConfigurationResponseBodyData) *GetAutoScalingConfigurationResponseBody
	GetData() *GetAutoScalingConfigurationResponseBodyData
	SetMessage(v string) *GetAutoScalingConfigurationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoScalingConfigurationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoScalingConfigurationResponseBody
	GetSuccess() *bool
}

type GetAutoScalingConfigurationResponseBody struct {
	// The response code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAutoScalingConfigurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B7A39AE5-0B36-4442-A304-E0885265***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAutoScalingConfigurationResponseBody) GetData() *GetAutoScalingConfigurationResponseBodyData {
	return s.Data
}

func (s *GetAutoScalingConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoScalingConfigurationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoScalingConfigurationResponseBody) SetCode(v int64) *GetAutoScalingConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetData(v *GetAutoScalingConfigurationResponseBodyData) *GetAutoScalingConfigurationResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetMessage(v string) *GetAutoScalingConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetRequestId(v string) *GetAutoScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetSuccess(v bool) *GetAutoScalingConfigurationResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingConfigurationResponseBodyData struct {
	// The scheduled scaling rules.
	ScheduledScalingRules *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules `json:"ScheduledScalingRules,omitempty" xml:"ScheduledScalingRules,omitempty" type:"Struct"`
}

func (s GetAutoScalingConfigurationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyData) GetScheduledScalingRules() *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules {
	return s.ScheduledScalingRules
}

func (s *GetAutoScalingConfigurationResponseBodyData) SetScheduledScalingRules(v *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) *GetAutoScalingConfigurationResponseBodyData {
	s.ScheduledScalingRules = v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyData) Validate() error {
	if s.ScheduledScalingRules != nil {
		if err := s.ScheduledScalingRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules struct {
	ScheduledScalingRules []*GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules `json:"ScheduledScalingRules,omitempty" xml:"ScheduledScalingRules,omitempty" type:"Repeated"`
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) GetScheduledScalingRules() []*GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	return s.ScheduledScalingRules
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) SetScheduledScalingRules(v []*GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules {
	s.ScheduledScalingRules = v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) Validate() error {
	if s.ScheduledScalingRules != nil {
		for _, item := range s.ScheduledScalingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules struct {
	// The duration of a scheduled scaling task. Unit: minutes.
	//
	// example:
	//
	// 60
	DurationMinutes *int64 `json:"DurationMinutes,omitempty" xml:"DurationMinutes,omitempty"`
	// Indicates whether the scheduled scaling rule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The estimated scale-in duration. Unit: seconds.
	//
	// example:
	//
	// 780
	EstimatedElasticScalingDownTimeSecs *int64 `json:"EstimatedElasticScalingDownTimeSecs,omitempty" xml:"EstimatedElasticScalingDownTimeSecs,omitempty"`
	// The estimated scale-out duration. Unit: seconds.
	//
	// example:
	//
	// 780
	EstimatedElasticScalingUpTimeSecs *int64 `json:"EstimatedElasticScalingUpTimeSecs,omitempty" xml:"EstimatedElasticScalingUpTimeSecs,omitempty"`
	// The timestamp that indicates the start time of the scheduled scaling task.
	//
	// example:
	//
	// 1714467540000
	FirstScheduledTime *int64 `json:"FirstScheduledTime,omitempty" xml:"FirstScheduledTime,omitempty"`
	// The frequency at which the scheduled scaling task is executed. This parameter is returned only if ScheduleType is set to repeat. Valid values:
	//
	// 	- Daily: The scheduled scaling task is executed every day.
	//
	// 	- Weekly: The scheduled scaling task is executed every week.
	//
	// example:
	//
	// Weekly
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The reserved production capacity for scheduled scaling. Unit: MB/s.
	//
	// example:
	//
	// 120
	ReservedPubFlow *int64 `json:"ReservedPubFlow,omitempty" xml:"ReservedPubFlow,omitempty"`
	// The reserved consumption capacity for scheduled scaling. Unit: MB/s.
	//
	// example:
	//
	// 120
	ReservedSubFlow *int64 `json:"ReservedSubFlow,omitempty" xml:"ReservedSubFlow,omitempty"`
	// The ID of the scheduled scaling rule.
	//
	// example:
	//
	// 64
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the scheduled scaling task. Valid values:
	//
	// 	- at: The scheduled scaling task is executed only once.
	//
	// 	- repeat: The scheduled scaling task is repeatedly executed.
	//
	// example:
	//
	// at
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The time zone in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// GMT+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The day on which the scheduled scaling task is repeatedly executed. You can specify multiple days for this parameter.
	WeeklyTypes *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty" type:"Struct"`
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetDurationMinutes() *int64 {
	return s.DurationMinutes
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetEnable() *bool {
	return s.Enable
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetEstimatedElasticScalingDownTimeSecs() *int64 {
	return s.EstimatedElasticScalingDownTimeSecs
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetEstimatedElasticScalingUpTimeSecs() *int64 {
	return s.EstimatedElasticScalingUpTimeSecs
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetFirstScheduledTime() *int64 {
	return s.FirstScheduledTime
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetRepeatType() *string {
	return s.RepeatType
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetReservedPubFlow() *int64 {
	return s.ReservedPubFlow
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetReservedSubFlow() *int64 {
	return s.ReservedSubFlow
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GetWeeklyTypes() *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes {
	return s.WeeklyTypes
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetDurationMinutes(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.DurationMinutes = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetEnable(v bool) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.Enable = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetEstimatedElasticScalingDownTimeSecs(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.EstimatedElasticScalingDownTimeSecs = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetEstimatedElasticScalingUpTimeSecs(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.EstimatedElasticScalingUpTimeSecs = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetFirstScheduledTime(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.FirstScheduledTime = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetRepeatType(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.RepeatType = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetReservedPubFlow(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.ReservedPubFlow = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetReservedSubFlow(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.ReservedSubFlow = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetRuleId(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetRuleName(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetScheduleType(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.ScheduleType = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetTimeZone(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.TimeZone = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetWeeklyTypes(v *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.WeeklyTypes = v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) Validate() error {
	if s.WeeklyTypes != nil {
		if err := s.WeeklyTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes struct {
	WeeklyTypes []*string `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty" type:"Repeated"`
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) GetWeeklyTypes() []*string {
	return s.WeeklyTypes
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) SetWeeklyTypes(v []*string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes {
	s.WeeklyTypes = v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) Validate() error {
	return dara.Validate(s)
}
