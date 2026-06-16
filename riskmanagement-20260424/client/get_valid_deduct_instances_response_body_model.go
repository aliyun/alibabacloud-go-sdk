// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidDeductInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetValidDeductInstancesResponseBody
	GetCode() *string
	SetData(v *GetValidDeductInstancesResponseBodyData) *GetValidDeductInstancesResponseBody
	GetData() *GetValidDeductInstancesResponseBodyData
	SetMessage(v string) *GetValidDeductInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetValidDeductInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetValidDeductInstancesResponseBody
	GetSuccess() *bool
}

type GetValidDeductInstancesResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetValidDeductInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetValidDeductInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetValidDeductInstancesResponseBody) GetData() *GetValidDeductInstancesResponseBodyData {
	return s.Data
}

func (s *GetValidDeductInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetValidDeductInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetValidDeductInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetValidDeductInstancesResponseBody) SetCode(v string) *GetValidDeductInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *GetValidDeductInstancesResponseBody) SetData(v *GetValidDeductInstancesResponseBodyData) *GetValidDeductInstancesResponseBody {
	s.Data = v
	return s
}

func (s *GetValidDeductInstancesResponseBody) SetMessage(v string) *GetValidDeductInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *GetValidDeductInstancesResponseBody) SetRequestId(v string) *GetValidDeductInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetValidDeductInstancesResponseBody) SetSuccess(v bool) *GetValidDeductInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *GetValidDeductInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetValidDeductInstancesResponseBodyData struct {
	Body *GetValidDeductInstancesResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s GetValidDeductInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesResponseBodyData) GetBody() *GetValidDeductInstancesResponseBodyDataBody {
	return s.Body
}

func (s *GetValidDeductInstancesResponseBodyData) SetBody(v *GetValidDeductInstancesResponseBodyDataBody) *GetValidDeductInstancesResponseBodyData {
	s.Body = v
	return s
}

func (s *GetValidDeductInstancesResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetValidDeductInstancesResponseBodyDataBody struct {
	Data      *GetValidDeductInstancesResponseBodyDataBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetValidDeductInstancesResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesResponseBodyDataBody) GetData() *GetValidDeductInstancesResponseBodyDataBodyData {
	return s.Data
}

func (s *GetValidDeductInstancesResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetValidDeductInstancesResponseBodyDataBody) SetData(v *GetValidDeductInstancesResponseBodyDataBodyData) *GetValidDeductInstancesResponseBodyDataBody {
	s.Data = v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBody) SetRequestId(v string) *GetValidDeductInstancesResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetValidDeductInstancesResponseBodyDataBodyData struct {
	CanTry            *bool                                                               `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	DeductPackageList []*GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList `json:"DeductPackageList,omitempty" xml:"DeductPackageList,omitempty" type:"Repeated"`
}

func (s GetValidDeductInstancesResponseBodyDataBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesResponseBodyDataBodyData) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesResponseBodyDataBodyData) GetCanTry() *bool {
	return s.CanTry
}

func (s *GetValidDeductInstancesResponseBodyDataBodyData) GetDeductPackageList() []*GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	return s.DeductPackageList
}

func (s *GetValidDeductInstancesResponseBodyDataBodyData) SetCanTry(v bool) *GetValidDeductInstancesResponseBodyDataBodyData {
	s.CanTry = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyData) SetDeductPackageList(v []*GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) *GetValidDeductInstancesResponseBodyDataBodyData {
	s.DeductPackageList = v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyData) Validate() error {
	if s.DeductPackageList != nil {
		for _, item := range s.DeductPackageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList struct {
	CurrentPeriodUsed *int64   `json:"CurrentPeriodUsed,omitempty" xml:"CurrentPeriodUsed,omitempty"`
	EndTime           *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InitCapacity      *float64 `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	InstanceId        *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Module            *string  `json:"Module,omitempty" xml:"Module,omitempty"`
	PeriodCapacity    *float64 `json:"PeriodCapacity,omitempty" xml:"PeriodCapacity,omitempty"`
	StartTime         *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetCurrentPeriodUsed() *int64 {
	return s.CurrentPeriodUsed
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetInitCapacity() *float64 {
	return s.InitCapacity
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetModule() *string {
	return s.Module
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetPeriodCapacity() *float64 {
	return s.PeriodCapacity
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) GetStatus() *string {
	return s.Status
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetCurrentPeriodUsed(v int64) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.CurrentPeriodUsed = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetEndTime(v int64) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.EndTime = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetInitCapacity(v float64) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.InitCapacity = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetInstanceId(v string) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.InstanceId = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetModule(v string) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.Module = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetPeriodCapacity(v float64) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.PeriodCapacity = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetStartTime(v int64) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.StartTime = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) SetStatus(v string) *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList {
	s.Status = &v
	return s
}

func (s *GetValidDeductInstancesResponseBodyDataBodyDataDeductPackageList) Validate() error {
	return dara.Validate(s)
}
