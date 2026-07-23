// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateTrafficControlTargetRequest
	GetEndTime() *string
	SetEvent(v string) *CreateTrafficControlTargetRequest
	GetEvent() *string
	SetItemConditionArray(v string) *CreateTrafficControlTargetRequest
	GetItemConditionArray() *string
	SetItemConditionExpress(v string) *CreateTrafficControlTargetRequest
	GetItemConditionExpress() *string
	SetItemConditionType(v string) *CreateTrafficControlTargetRequest
	GetItemConditionType() *string
	SetName(v string) *CreateTrafficControlTargetRequest
	GetName() *string
	SetNewProductRegulation(v bool) *CreateTrafficControlTargetRequest
	GetNewProductRegulation() *bool
	SetRecallName(v string) *CreateTrafficControlTargetRequest
	GetRecallName() *string
	SetStartTime(v string) *CreateTrafficControlTargetRequest
	GetStartTime() *string
	SetStatisPeriod(v string) *CreateTrafficControlTargetRequest
	GetStatisPeriod() *string
	SetStatus(v string) *CreateTrafficControlTargetRequest
	GetStatus() *string
	SetToleranceValue(v int64) *CreateTrafficControlTargetRequest
	GetToleranceValue() *int64
	SetTrafficControlTaskId(v string) *CreateTrafficControlTargetRequest
	GetTrafficControlTaskId() *string
	SetValue(v float32) *CreateTrafficControlTargetRequest
	GetValue() *float32
}

type CreateTrafficControlTargetRequest struct {
	// The end time of the traffic control target.
	//
	// example:
	//
	// 2024-04-25
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event for the traffic control target.
	//
	// example:
	//
	// click
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The item condition in array format.
	//
	// example:
	//
	// [{"field":"status","option":"=","value":"1"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// The item condition in expression format.
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
	// Specifies whether to enable new product regulation.
	//
	// example:
	//
	// false
	NewProductRegulation *bool `json:"NewProductRegulation,omitempty" xml:"NewProductRegulation,omitempty"`
	// The recall strategy name.
	//
	// example:
	//
	// recall-1
	RecallName *string `json:"RecallName,omitempty" xml:"RecallName,omitempty"`
	// The start time of the traffic control target.
	//
	// example:
	//
	// 2024-03-25
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The statistical period.
	//
	// example:
	//
	// Daily
	StatisPeriod *string `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	// The status of the traffic control target.
	//
	// example:
	//
	// Opened
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tolerance value for the traffic control target.
	//
	// example:
	//
	// 10
	ToleranceValue *int64 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	// The traffic control task ID.
	//
	// example:
	//
	// 1
	TrafficControlTaskId *string `json:"TrafficControlTaskId,omitempty" xml:"TrafficControlTaskId,omitempty"`
	// The value of the traffic control target.
	//
	// example:
	//
	// 10
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTargetRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateTrafficControlTargetRequest) GetEvent() *string {
	return s.Event
}

func (s *CreateTrafficControlTargetRequest) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *CreateTrafficControlTargetRequest) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *CreateTrafficControlTargetRequest) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *CreateTrafficControlTargetRequest) GetName() *string {
	return s.Name
}

func (s *CreateTrafficControlTargetRequest) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *CreateTrafficControlTargetRequest) GetRecallName() *string {
	return s.RecallName
}

func (s *CreateTrafficControlTargetRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateTrafficControlTargetRequest) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *CreateTrafficControlTargetRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateTrafficControlTargetRequest) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *CreateTrafficControlTargetRequest) GetTrafficControlTaskId() *string {
	return s.TrafficControlTaskId
}

func (s *CreateTrafficControlTargetRequest) GetValue() *float32 {
	return s.Value
}

func (s *CreateTrafficControlTargetRequest) SetEndTime(v string) *CreateTrafficControlTargetRequest {
	s.EndTime = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetEvent(v string) *CreateTrafficControlTargetRequest {
	s.Event = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetItemConditionArray(v string) *CreateTrafficControlTargetRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetItemConditionExpress(v string) *CreateTrafficControlTargetRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetItemConditionType(v string) *CreateTrafficControlTargetRequest {
	s.ItemConditionType = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetName(v string) *CreateTrafficControlTargetRequest {
	s.Name = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetNewProductRegulation(v bool) *CreateTrafficControlTargetRequest {
	s.NewProductRegulation = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetRecallName(v string) *CreateTrafficControlTargetRequest {
	s.RecallName = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetStartTime(v string) *CreateTrafficControlTargetRequest {
	s.StartTime = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetStatisPeriod(v string) *CreateTrafficControlTargetRequest {
	s.StatisPeriod = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetStatus(v string) *CreateTrafficControlTargetRequest {
	s.Status = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetToleranceValue(v int64) *CreateTrafficControlTargetRequest {
	s.ToleranceValue = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetTrafficControlTaskId(v string) *CreateTrafficControlTargetRequest {
	s.TrafficControlTaskId = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) SetValue(v float32) *CreateTrafficControlTargetRequest {
	s.Value = &v
	return s
}

func (s *CreateTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
