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
	EndTime                *string                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event                  *string                                        `json:"Event,omitempty" xml:"Event,omitempty"`
	GmtCreateTime          *string                                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	ItemConditionArray     *string                                        `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	ItemConditionExpress   *string                                        `json:"ItemConditionExpress,omitempty" xml:"ItemConditionExpress,omitempty"`
	ItemConditionType      *string                                        `json:"ItemConditionType,omitempty" xml:"ItemConditionType,omitempty"`
	Name                   *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	NewProductRegulation   *bool                                          `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	RecallName             *string                                        `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	RequestId              *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SplitParts             *GetTrafficControlTargetResponseBodySplitParts `json:"SplitParts,omitempty" xml:"SplitParts,omitempty" type:"Struct"`
	StartTime              *string                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatisPeriod           *string                                        `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	Status                 *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	ToleranceValue         *int64                                         `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	TrafficControlTargetId *string                                        `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	TrafficControlTaskId   *string                                        `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	Value                  *float32                                       `json:"Value,omitempty" xml:"Value,omitempty"`
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
	return dara.Validate(s)
}

type GetTrafficControlTargetResponseBodySplitParts struct {
	SetPoints  []*int64 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
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
