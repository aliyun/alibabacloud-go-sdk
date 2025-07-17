// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSilencePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveTimeType(v string) *CreateOrUpdateSilencePolicyRequest
	GetEffectiveTimeType() *string
	SetId(v int64) *CreateOrUpdateSilencePolicyRequest
	GetId() *int64
	SetMatchingRules(v string) *CreateOrUpdateSilencePolicyRequest
	GetMatchingRules() *string
	SetName(v string) *CreateOrUpdateSilencePolicyRequest
	GetName() *string
	SetRegionId(v string) *CreateOrUpdateSilencePolicyRequest
	GetRegionId() *string
	SetState(v string) *CreateOrUpdateSilencePolicyRequest
	GetState() *string
	SetTimePeriod(v string) *CreateOrUpdateSilencePolicyRequest
	GetTimePeriod() *string
	SetTimeSlots(v string) *CreateOrUpdateSilencePolicyRequest
	GetTimeSlots() *string
}

type CreateOrUpdateSilencePolicyRequest struct {
	// The effective duration of the silence policy. Valid values: PERMANENT, CUSTOM_TIME, and CYCLE_EFFECT.
	//
	// example:
	//
	// PERMANENT
	EffectiveTimeType *string `json:"EffectiveTimeType,omitempty" xml:"EffectiveTimeType,omitempty"`
	// The ID of the silence policy.
	//
	// 	- If you do not configure this parameter, a new silence policy is created.
	//
	// 	- If you configure this parameter, the specified silence policy is modified.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The matching rules. The following code shows the format of matching rules:
	//
	//     [
	//
	//          {
	//
	//     	 "matchingConditions": [
	//
	//     	 {
	//
	//     	 "value": "test", // The value of the matching condition.
	//
	//     	 "key": "altertname", // The key of the matching condition.
	//
	//     	 "operator": "eq" // The logical operator of the matching condition, including eq (equal to), neq (not equal to), in (contains), nin (does not contain), re (regular expression match), and nre (regular expression mismatch).
	//
	//     	 }
	//
	//     	 ]
	//
	//          }
	//
	//     	 ]
	//
	// example:
	//
	// [ 	 { 	 "matchingConditions": [ 	 { 	 "value": "test", 	 "key": "altertname", 	 "operator": "eq" 	 } 	 ]      } 	 ]
	MatchingRules *string `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty"`
	// The name of the silence policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// silencepolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the silence policy. Valid values: enable and disable.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The recurring period. This parameter is required when EffectiveTimeType is set to CYCLE_EFFECT. DAY: The silence policy is effective by day. WEEK: The silence policy is effective by week.
	//
	// example:
	//
	// DAY
	TimePeriod *string `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty"`
	// The time period during which the silence policy is effective. If you set EffectiveTimeType to CUSTOM_TIME, specify a custom time period in the following format: [{"startTime":"2024-08-04 22:13","endTime":"2024-08-04 22:21"}] If you set EffectiveTimeType to CYCLE_EFFECT and TimePeriod to DAY, specify a custom time period in the following format: [{"startTime":"22:13","endTime":"22:21"}]. The start time cannot be later than the end time. If you set EffectiveTimeType to CYCLE_EFFECT and TimePeriod to WEEK, specify a custom time period in the following format: [{"startWeek":"1", "endWeek":"2" "startTime":"22:13","endTime":"22:21"}]. Valid values of startWeek and endWeek: 1 to 7. The start time cannot be later than the end time.
	//
	// example:
	//
	// [{"startTime":"2024-08-04 22:13","endTime":"2024-08-04 22:21"}]
	TimeSlots *string `json:"TimeSlots,omitempty" xml:"TimeSlots,omitempty"`
}

func (s CreateOrUpdateSilencePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyRequest) GetEffectiveTimeType() *string {
	return s.EffectiveTimeType
}

func (s *CreateOrUpdateSilencePolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSilencePolicyRequest) GetMatchingRules() *string {
	return s.MatchingRules
}

func (s *CreateOrUpdateSilencePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSilencePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateSilencePolicyRequest) GetState() *string {
	return s.State
}

func (s *CreateOrUpdateSilencePolicyRequest) GetTimePeriod() *string {
	return s.TimePeriod
}

func (s *CreateOrUpdateSilencePolicyRequest) GetTimeSlots() *string {
	return s.TimeSlots
}

func (s *CreateOrUpdateSilencePolicyRequest) SetEffectiveTimeType(v string) *CreateOrUpdateSilencePolicyRequest {
	s.EffectiveTimeType = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetId(v int64) *CreateOrUpdateSilencePolicyRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetMatchingRules(v string) *CreateOrUpdateSilencePolicyRequest {
	s.MatchingRules = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetName(v string) *CreateOrUpdateSilencePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetRegionId(v string) *CreateOrUpdateSilencePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetState(v string) *CreateOrUpdateSilencePolicyRequest {
	s.State = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetTimePeriod(v string) *CreateOrUpdateSilencePolicyRequest {
	s.TimePeriod = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetTimeSlots(v string) *CreateOrUpdateSilencePolicyRequest {
	s.TimeSlots = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) Validate() error {
	return dara.Validate(s)
}
