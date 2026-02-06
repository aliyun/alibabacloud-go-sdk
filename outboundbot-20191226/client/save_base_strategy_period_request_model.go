// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBaseStrategyPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *SaveBaseStrategyPeriodRequest
	GetEntryId() *string
	SetOnlyWeekdays(v bool) *SaveBaseStrategyPeriodRequest
	GetOnlyWeekdays() *bool
	SetOnlyWorkdays(v bool) *SaveBaseStrategyPeriodRequest
	GetOnlyWorkdays() *bool
	SetStrategyLevel(v int32) *SaveBaseStrategyPeriodRequest
	GetStrategyLevel() *int32
	SetWorkingTime(v []*string) *SaveBaseStrategyPeriodRequest
	GetWorkingTime() []*string
	SetWorkingTimeFramesJson(v string) *SaveBaseStrategyPeriodRequest
	GetWorkingTimeFramesJson() *string
}

type SaveBaseStrategyPeriodRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fe51eca2-a904-4b4d-b3ff-31be334b9500
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// example:
	//
	// false
	OnlyWeekdays *bool `json:"OnlyWeekdays,omitempty" xml:"OnlyWeekdays,omitempty"`
	OnlyWorkdays *bool `json:"OnlyWorkdays,omitempty" xml:"OnlyWorkdays,omitempty"`
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
	// example:
	//
	// []
	WorkingTime []*string `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
	// example:
	//
	// [{\\"beginTime\\":\\"09:00:00\\",\\"beginTimeMillis\\":324000000000,\\"endTime\\":\\"21:00:00\\",\\"endTimeMillis\\":756000000000}]
	WorkingTimeFramesJson *string `json:"WorkingTimeFramesJson,omitempty" xml:"WorkingTimeFramesJson,omitempty"`
}

func (s SaveBaseStrategyPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBaseStrategyPeriodRequest) GoString() string {
	return s.String()
}

func (s *SaveBaseStrategyPeriodRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *SaveBaseStrategyPeriodRequest) GetOnlyWeekdays() *bool {
	return s.OnlyWeekdays
}

func (s *SaveBaseStrategyPeriodRequest) GetOnlyWorkdays() *bool {
	return s.OnlyWorkdays
}

func (s *SaveBaseStrategyPeriodRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *SaveBaseStrategyPeriodRequest) GetWorkingTime() []*string {
	return s.WorkingTime
}

func (s *SaveBaseStrategyPeriodRequest) GetWorkingTimeFramesJson() *string {
	return s.WorkingTimeFramesJson
}

func (s *SaveBaseStrategyPeriodRequest) SetEntryId(v string) *SaveBaseStrategyPeriodRequest {
	s.EntryId = &v
	return s
}

func (s *SaveBaseStrategyPeriodRequest) SetOnlyWeekdays(v bool) *SaveBaseStrategyPeriodRequest {
	s.OnlyWeekdays = &v
	return s
}

func (s *SaveBaseStrategyPeriodRequest) SetOnlyWorkdays(v bool) *SaveBaseStrategyPeriodRequest {
	s.OnlyWorkdays = &v
	return s
}

func (s *SaveBaseStrategyPeriodRequest) SetStrategyLevel(v int32) *SaveBaseStrategyPeriodRequest {
	s.StrategyLevel = &v
	return s
}

func (s *SaveBaseStrategyPeriodRequest) SetWorkingTime(v []*string) *SaveBaseStrategyPeriodRequest {
	s.WorkingTime = v
	return s
}

func (s *SaveBaseStrategyPeriodRequest) SetWorkingTimeFramesJson(v string) *SaveBaseStrategyPeriodRequest {
	s.WorkingTimeFramesJson = &v
	return s
}

func (s *SaveBaseStrategyPeriodRequest) Validate() error {
	return dara.Validate(s)
}
