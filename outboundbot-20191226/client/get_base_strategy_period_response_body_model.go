// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaseStrategyPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBaseStrategyPeriodResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetBaseStrategyPeriodResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBaseStrategyPeriodResponseBody
	GetMessage() *string
	SetOnlyWeekdays(v bool) *GetBaseStrategyPeriodResponseBody
	GetOnlyWeekdays() *bool
	SetOnlyWorkdays(v bool) *GetBaseStrategyPeriodResponseBody
	GetOnlyWorkdays() *bool
	SetRequestId(v string) *GetBaseStrategyPeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBaseStrategyPeriodResponseBody
	GetSuccess() *bool
	SetWorkingTime(v []*GetBaseStrategyPeriodResponseBodyWorkingTime) *GetBaseStrategyPeriodResponseBody
	GetWorkingTime() []*GetBaseStrategyPeriodResponseBodyWorkingTime
}

type GetBaseStrategyPeriodResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	OnlyWeekdays *bool `json:"OnlyWeekdays,omitempty" xml:"OnlyWeekdays,omitempty"`
	OnlyWorkdays *bool `json:"OnlyWorkdays,omitempty" xml:"OnlyWorkdays,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// []
	WorkingTime []*GetBaseStrategyPeriodResponseBodyWorkingTime `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
}

func (s GetBaseStrategyPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBaseStrategyPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaseStrategyPeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBaseStrategyPeriodResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBaseStrategyPeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBaseStrategyPeriodResponseBody) GetOnlyWeekdays() *bool {
	return s.OnlyWeekdays
}

func (s *GetBaseStrategyPeriodResponseBody) GetOnlyWorkdays() *bool {
	return s.OnlyWorkdays
}

func (s *GetBaseStrategyPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBaseStrategyPeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBaseStrategyPeriodResponseBody) GetWorkingTime() []*GetBaseStrategyPeriodResponseBodyWorkingTime {
	return s.WorkingTime
}

func (s *GetBaseStrategyPeriodResponseBody) SetCode(v string) *GetBaseStrategyPeriodResponseBody {
	s.Code = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetHttpStatusCode(v int32) *GetBaseStrategyPeriodResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetMessage(v string) *GetBaseStrategyPeriodResponseBody {
	s.Message = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetOnlyWeekdays(v bool) *GetBaseStrategyPeriodResponseBody {
	s.OnlyWeekdays = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetOnlyWorkdays(v bool) *GetBaseStrategyPeriodResponseBody {
	s.OnlyWorkdays = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetRequestId(v string) *GetBaseStrategyPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetSuccess(v bool) *GetBaseStrategyPeriodResponseBody {
	s.Success = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) SetWorkingTime(v []*GetBaseStrategyPeriodResponseBodyWorkingTime) *GetBaseStrategyPeriodResponseBody {
	s.WorkingTime = v
	return s
}

func (s *GetBaseStrategyPeriodResponseBody) Validate() error {
	if s.WorkingTime != nil {
		for _, item := range s.WorkingTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBaseStrategyPeriodResponseBodyWorkingTime struct {
	// example:
	//
	// 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 0
	BeginTimeMillis *int64 `json:"BeginTimeMillis,omitempty" xml:"BeginTimeMillis,omitempty"`
	// example:
	//
	// 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0
	EndTimeMillis *int64 `json:"EndTimeMillis,omitempty" xml:"EndTimeMillis,omitempty"`
}

func (s GetBaseStrategyPeriodResponseBodyWorkingTime) String() string {
	return dara.Prettify(s)
}

func (s GetBaseStrategyPeriodResponseBodyWorkingTime) GoString() string {
	return s.String()
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) GetBeginTimeMillis() *int64 {
	return s.BeginTimeMillis
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) GetEndTime() *string {
	return s.EndTime
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) GetEndTimeMillis() *int64 {
	return s.EndTimeMillis
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) SetBeginTime(v string) *GetBaseStrategyPeriodResponseBodyWorkingTime {
	s.BeginTime = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) SetBeginTimeMillis(v int64) *GetBaseStrategyPeriodResponseBodyWorkingTime {
	s.BeginTimeMillis = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) SetEndTime(v string) *GetBaseStrategyPeriodResponseBodyWorkingTime {
	s.EndTime = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) SetEndTimeMillis(v int64) *GetBaseStrategyPeriodResponseBodyWorkingTime {
	s.EndTimeMillis = &v
	return s
}

func (s *GetBaseStrategyPeriodResponseBodyWorkingTime) Validate() error {
	return dara.Validate(s)
}
