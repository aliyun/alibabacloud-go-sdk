// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateTrafficControlTargetRequest
	GetEndTime() *string
	SetEvent(v string) *UpdateTrafficControlTargetRequest
	GetEvent() *string
	SetItemConditionArray(v string) *UpdateTrafficControlTargetRequest
	GetItemConditionArray() *string
	SetItemConditionExpress(v string) *UpdateTrafficControlTargetRequest
	GetItemConditionExpress() *string
	SetItemConditionType(v string) *UpdateTrafficControlTargetRequest
	GetItemConditionType() *string
	SetName(v string) *UpdateTrafficControlTargetRequest
	GetName() *string
	SetNewProductRegulation(v bool) *UpdateTrafficControlTargetRequest
	GetNewProductRegulation() *bool
	SetRecallName(v string) *UpdateTrafficControlTargetRequest
	GetRecallName() *string
	SetStartTime(v string) *UpdateTrafficControlTargetRequest
	GetStartTime() *string
	SetStatisPeriod(v string) *UpdateTrafficControlTargetRequest
	GetStatisPeriod() *string
	SetStatus(v string) *UpdateTrafficControlTargetRequest
	GetStatus() *string
	SetToleranceValue(v int64) *UpdateTrafficControlTargetRequest
	GetToleranceValue() *int64
	SetValue(v float32) *UpdateTrafficControlTargetRequest
	GetValue() *float32
	SetNewParam3(v string) *UpdateTrafficControlTargetRequest
	GetNewParam3() *string
}

type UpdateTrafficControlTargetRequest struct {
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
	// The traffic control target name.
	//
	// example:
	//
	// target-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the control rule applies to a new product.
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
	// The statistics period.
	//
	// example:
	//
	// Daily
	StatisPeriod *string `json:"StatisPeriod,omitempty" xml:"StatisPeriod,omitempty"`
	// The traffic control target status.
	//
	// example:
	//
	// Opened
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tolerance range for the traffic control target.
	//
	// example:
	//
	// 10
	ToleranceValue *int64 `json:"ToleranceValue,omitempty" xml:"ToleranceValue,omitempty"`
	// The traffic control target value.
	//
	// example:
	//
	// 30
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
	// This parameter is invalid.
	//
	// example:
	//
	// null
	NewParam3 *string `json:"new-param-3,omitempty" xml:"new-param-3,omitempty"`
}

func (s UpdateTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTargetRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateTrafficControlTargetRequest) GetEvent() *string {
	return s.Event
}

func (s *UpdateTrafficControlTargetRequest) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *UpdateTrafficControlTargetRequest) GetItemConditionExpress() *string {
	return s.ItemConditionExpress
}

func (s *UpdateTrafficControlTargetRequest) GetItemConditionType() *string {
	return s.ItemConditionType
}

func (s *UpdateTrafficControlTargetRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTrafficControlTargetRequest) GetNewProductRegulation() *bool {
	return s.NewProductRegulation
}

func (s *UpdateTrafficControlTargetRequest) GetRecallName() *string {
	return s.RecallName
}

func (s *UpdateTrafficControlTargetRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateTrafficControlTargetRequest) GetStatisPeriod() *string {
	return s.StatisPeriod
}

func (s *UpdateTrafficControlTargetRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateTrafficControlTargetRequest) GetToleranceValue() *int64 {
	return s.ToleranceValue
}

func (s *UpdateTrafficControlTargetRequest) GetValue() *float32 {
	return s.Value
}

func (s *UpdateTrafficControlTargetRequest) GetNewParam3() *string {
	return s.NewParam3
}

func (s *UpdateTrafficControlTargetRequest) SetEndTime(v string) *UpdateTrafficControlTargetRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetEvent(v string) *UpdateTrafficControlTargetRequest {
	s.Event = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetItemConditionArray(v string) *UpdateTrafficControlTargetRequest {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetItemConditionExpress(v string) *UpdateTrafficControlTargetRequest {
	s.ItemConditionExpress = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetItemConditionType(v string) *UpdateTrafficControlTargetRequest {
	s.ItemConditionType = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetName(v string) *UpdateTrafficControlTargetRequest {
	s.Name = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetNewProductRegulation(v bool) *UpdateTrafficControlTargetRequest {
	s.NewProductRegulation = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetRecallName(v string) *UpdateTrafficControlTargetRequest {
	s.RecallName = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetStartTime(v string) *UpdateTrafficControlTargetRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetStatisPeriod(v string) *UpdateTrafficControlTargetRequest {
	s.StatisPeriod = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetStatus(v string) *UpdateTrafficControlTargetRequest {
	s.Status = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetToleranceValue(v int64) *UpdateTrafficControlTargetRequest {
	s.ToleranceValue = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetValue(v float32) *UpdateTrafficControlTargetRequest {
	s.Value = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) SetNewParam3(v string) *UpdateTrafficControlTargetRequest {
	s.NewParam3 = &v
	return s
}

func (s *UpdateTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
