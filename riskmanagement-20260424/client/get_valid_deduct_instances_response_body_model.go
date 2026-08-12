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
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data list.
	Data *GetValidDeductInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FBDD713-00A5-5C98-B661-3FD31A349B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current API call itself is successful. This does not indicate the success of subsequent business operations.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The message body.
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
	// The data body.
	Data *GetValidDeductInstancesResponseBodyDataBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The Security Center request ID.
	//
	// example:
	//
	// A6FB9AC3-4431-538F-BA8A-2A13AEA208A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Indicates whether the account is eligible for a trial. Valid values:
	//
	// - **true**: Eligible.
	//
	// - **false**: Not eligible.
	//
	// example:
	//
	// true
	CanTry *bool `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	// The resource plan usage details.
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
	// The number of authorizations consumed in the current metering cycle.
	//
	// example:
	//
	// 10
	CurrentPeriodUsed *int64 `json:"CurrentPeriodUsed,omitempty" xml:"CurrentPeriodUsed,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1737734400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The total capacity of the resource plan.
	//
	// example:
	//
	// 1000
	InitCapacity *float64 `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	// The resource plan instance ID.
	//
	// example:
	//
	// apigateway-hz-96f6659a1490
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module code. Valid values:
	//
	// - **POST_HOST**: Host and container protection.
	//
	// - **CSPM**: Cloud product configuration check.
	//
	// - **VUL**: Vulnerability scanning.
	//
	// example:
	//
	// POST_HOST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of authorizations consumed up to the previous billing cycle.
	//
	// example:
	//
	// 1
	PeriodCapacity *float64 `json:"PeriodCapacity,omitempty" xml:"PeriodCapacity,omitempty"`
	// The start timestamp, in milliseconds.
	//
	// example:
	//
	// 1737734400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The resource plan status. Valid values:
	//
	// - **valid**: Valid.
	//
	// - **invalid**: Invalid.
	//
	// example:
	//
	// CREATE_FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
