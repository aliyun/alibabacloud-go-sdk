// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetTrafficControlTargetResponseBody
	GetEndTime() *string
	SetEvent(v string) *GetTrafficControlTargetResponseBody
	GetEvent() *string
	SetGmtCreateTime(v string) *GetTrafficControlTargetResponseBody
	GetGmtCreateTime() *string
	SetItemConditionArray(v string) *GetTrafficControlTargetResponseBody
	GetItemConditionArray() *string
	SetItemConditionExpress(v string) *GetTrafficControlTargetResponseBody
	GetItemConditionExpress() *string
	SetItemConditionType(v string) *GetTrafficControlTargetResponseBody
	GetItemConditionType() *string
	SetName(v string) *GetTrafficControlTargetResponseBody
	GetName() *string
	SetNewProductRegulation(v bool) *GetTrafficControlTargetResponseBody
	GetNewProductRegulation() *bool
	SetRecallName(v string) *GetTrafficControlTargetResponseBody
	GetRecallName() *string
	SetRequestId(v string) *GetTrafficControlTargetResponseBody
	GetRequestId() *string
	SetSplitParts(v *GetTrafficControlTargetResponseBodySplitParts) *GetTrafficControlTargetResponseBody
	GetSplitParts() *GetTrafficControlTargetResponseBodySplitParts
	SetStartTime(v string) *GetTrafficControlTargetResponseBody
	GetStartTime() *string
	SetStatisPeriod(v string) *GetTrafficControlTargetResponseBody
	GetStatisPeriod() *string
	SetStatus(v string) *GetTrafficControlTargetResponseBody
	GetStatus() *string
	SetToleranceValue(v int64) *GetTrafficControlTargetResponseBody
	GetToleranceValue() *int64
	SetTrafficControlTargetId(v string) *GetTrafficControlTargetResponseBody
	GetTrafficControlTargetId() *string
	SetTrafficControlTaskId(v string) *GetTrafficControlTargetResponseBody
	GetTrafficControlTaskId() *string
	SetValue(v float32) *GetTrafficControlTargetResponseBody
	GetValue() *float32
}

type GetTrafficControlTargetResponseBody struct {
	// The end time.
	//
	// example:
	//
	// 2024-04-25
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event of the control target.
	//
	// example:
	//
	// click
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The time when the traffic control target was created.
	//
	// example:
	//
	// 2024-01-03T02:28:00.000Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The item condition, specified in an array format.
	//
	// example:
	//
	// [{"field":"status","option":"=","value":"1"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// The item condition, specified in an expression format.
	//
	// example:
	//
	// status=1
	ItemConditionExpress *string `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	// The item condition type.
	//
	// example:
	//
	// Array
	ItemConditionType *string `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	// The name of the traffic control target.
	//
	// example:
	//
	// target-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether new product regulation is enabled.
	//
	// example:
	//
	// false
	NewProductRegulation *bool `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	// The name of the recall strategy.
	//
	// example:
	//
	// recall-1
	RecallName *string `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The piecewise control settings.
	SplitParts *GetTrafficControlTargetResponseBodySplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	// The start time.
	//
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The statistics period.
	//
	// example:
	//
	// Daily
	StatisPeriod *string `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	// The status of the traffic control target.
	//
	// example:
	//
	// 枚举值：开启：Opened关闭：Closed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tolerance value.
	//
	// example:
	//
	// 10
	ToleranceValue *int64 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	// The ID of the traffic control target.
	//
	// example:
	//
	// 1
	TrafficControlTargetId *string `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	// The ID of the traffic control task.
	//
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	// The value of the control target.
	//
	// example:
	//
	// 30
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrafficControlTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTargetResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTrafficControlTargetResponseBody) GetEvent() *string {
	return s.Event
}

func (s *GetTrafficControlTargetResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTrafficControlTargetResponseBody) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *GetTrafficControlTargetResponseBody) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *GetTrafficControlTargetResponseBody) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *GetTrafficControlTargetResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTrafficControlTargetResponseBody) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *GetTrafficControlTargetResponseBody) GetRecallName() *string {
	return s.RecallName
}

func (s *GetTrafficControlTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrafficControlTargetResponseBody) GetSplitParts() *GetTrafficControlTargetResponseBodySplitParts {
	return s.SplitParts
}

func (s *GetTrafficControlTargetResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTrafficControlTargetResponseBody) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *GetTrafficControlTargetResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTrafficControlTargetResponseBody) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *GetTrafficControlTargetResponseBody) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *GetTrafficControlTargetResponseBody) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *GetTrafficControlTargetResponseBody) GetValue() *float32 {
	return s.Value
}

func (s *GetTrafficControlTargetResponseBody) SetEndTime(v string) *GetTrafficControlTargetResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetEvent(v string) *GetTrafficControlTargetResponseBody {
	s.Event = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetGmtCreateTime(v string) *GetTrafficControlTargetResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetItemConditionArray(v string) *GetTrafficControlTargetResponseBody {
	s.ItemConditionArray = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetItemConditionExpress(v string) *GetTrafficControlTargetResponseBody {
	s.ItemConditionExpress = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetItemConditionType(v string) *GetTrafficControlTargetResponseBody {
	s.ItemConditionType = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetName(v string) *GetTrafficControlTargetResponseBody {
	s.Name = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetNewProductRegulation(v bool) *GetTrafficControlTargetResponseBody {
	s.NewProductRegulation = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetRecallName(v string) *GetTrafficControlTargetResponseBody {
	s.RecallName = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetRequestId(v string) *GetTrafficControlTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetSplitParts(v *GetTrafficControlTargetResponseBodySplitParts) *GetTrafficControlTargetResponseBody {
	s.SplitParts = v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetStartTime(v string) *GetTrafficControlTargetResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetStatisPeriod(v string) *GetTrafficControlTargetResponseBody {
	s.StatisPeriod = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetStatus(v string) *GetTrafficControlTargetResponseBody {
	s.Status = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetToleranceValue(v int64) *GetTrafficControlTargetResponseBody {
	s.ToleranceValue = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetTrafficControlTargetId(v string) *GetTrafficControlTargetResponseBody {
	s.TrafficControlTargetId = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetTrafficControlTaskId(v string) *GetTrafficControlTargetResponseBody {
	s.TrafficControlTaskId = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) SetValue(v float32) *GetTrafficControlTargetResponseBody {
	s.Value = &v
	return s
}

func (s *GetTrafficControlTargetResponseBody) Validate() error {
	if s.SplitParts != nil {
		if err := s.SplitParts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrafficControlTargetResponseBodySplitParts struct {
	// The set points.
	SetPoints []*int64 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	// The set values.
	SetValues []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	// The time points.
	TimePoints []*int64 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s GetTrafficControlTargetResponseBodySplitParts) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTargetResponseBodySplitParts) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetResponseBodySplitParts) GetSetPoints() []*int64 {
	return s.SetPoints
}

func (s *GetTrafficControlTargetResponseBodySplitParts) GetSetValues() []*int64 {
	return s.SetValues
}

func (s *GetTrafficControlTargetResponseBodySplitParts) GetTimePoints() []*int64 {
	return s.TimePoints
}

func (s *GetTrafficControlTargetResponseBodySplitParts) SetSetPoints(v []*int64) *GetTrafficControlTargetResponseBodySplitParts {
	s.SetPoints = v
	return s
}

func (s *GetTrafficControlTargetResponseBodySplitParts) SetSetValues(v []*int64) *GetTrafficControlTargetResponseBodySplitParts {
	s.SetValues = v
	return s
}

func (s *GetTrafficControlTargetResponseBodySplitParts) SetTimePoints(v []*int64) *GetTrafficControlTargetResponseBodySplitParts {
	s.TimePoints = v
	return s
}

func (s *GetTrafficControlTargetResponseBodySplitParts) Validate() error {
	return dara.Validate(s)
}
