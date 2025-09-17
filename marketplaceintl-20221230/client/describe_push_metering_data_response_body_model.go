// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePushMeteringDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePushMeteringDataResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *DescribePushMeteringDataResponseBody
	GetDynamicMessage() *string
	SetForceFatal(v bool) *DescribePushMeteringDataResponseBody
	GetForceFatal() *bool
	SetMessage(v string) *DescribePushMeteringDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePushMeteringDataResponseBody
	GetRequestId() *string
	SetResult(v *DescribePushMeteringDataResponseBodyResult) *DescribePushMeteringDataResponseBody
	GetResult() *DescribePushMeteringDataResponseBodyResult
	SetSuccess(v bool) *DescribePushMeteringDataResponseBody
	GetSuccess() *bool
}

type DescribePushMeteringDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// parameter \\"service\\" can not be blank.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// false
	ForceFatal *bool `json:"ForceFatal,omitempty" xml:"ForceFatal,omitempty"`
	// example:
	//
	// Instance 5723f7ee-952d-411f-94f4-b942a550d9b8 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A6A33748-D573-593C-A3BC-593E33D68311
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribePushMeteringDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePushMeteringDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePushMeteringDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePushMeteringDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePushMeteringDataResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribePushMeteringDataResponseBody) GetForceFatal() *bool {
	return s.ForceFatal
}

func (s *DescribePushMeteringDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePushMeteringDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePushMeteringDataResponseBody) GetResult() *DescribePushMeteringDataResponseBodyResult {
	return s.Result
}

func (s *DescribePushMeteringDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePushMeteringDataResponseBody) SetCode(v string) *DescribePushMeteringDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePushMeteringDataResponseBody) SetDynamicMessage(v string) *DescribePushMeteringDataResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribePushMeteringDataResponseBody) SetForceFatal(v bool) *DescribePushMeteringDataResponseBody {
	s.ForceFatal = &v
	return s
}

func (s *DescribePushMeteringDataResponseBody) SetMessage(v string) *DescribePushMeteringDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePushMeteringDataResponseBody) SetRequestId(v string) *DescribePushMeteringDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePushMeteringDataResponseBody) SetResult(v *DescribePushMeteringDataResponseBodyResult) *DescribePushMeteringDataResponseBody {
	s.Result = v
	return s
}

func (s *DescribePushMeteringDataResponseBody) SetSuccess(v bool) *DescribePushMeteringDataResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePushMeteringDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePushMeteringDataResponseBodyResult struct {
	// example:
	//
	// 2024-09-18T03:15:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 5000002763123
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsPushBilling *bool   `json:"IsPushBilling,omitempty" xml:"IsPushBilling,omitempty"`
	// example:
	//
	// sgcmgj0003XXXX-Period-1
	MeteringAssist *string `json:"MeteringAssist,omitempty" xml:"MeteringAssist,omitempty"`
	// example:
	//
	// {"Frequency":1}
	MeteringEntity *string `json:"MeteringEntity,omitempty" xml:"MeteringEntity,omitempty"`
	// example:
	//
	// pushData123456
	PushOrderBizId *string `json:"PushOrderBizId,omitempty" xml:"PushOrderBizId,omitempty"`
	// example:
	//
	// 2025-01-09T02:04:58Z
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePushMeteringDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribePushMeteringDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribePushMeteringDataResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePushMeteringDataResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePushMeteringDataResponseBodyResult) GetIsPushBilling() *bool {
	return s.IsPushBilling
}

func (s *DescribePushMeteringDataResponseBodyResult) GetMeteringAssist() *string {
	return s.MeteringAssist
}

func (s *DescribePushMeteringDataResponseBodyResult) GetMeteringEntity() *string {
	return s.MeteringEntity
}

func (s *DescribePushMeteringDataResponseBodyResult) GetPushOrderBizId() *string {
	return s.PushOrderBizId
}

func (s *DescribePushMeteringDataResponseBodyResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePushMeteringDataResponseBodyResult) SetEndTime(v int64) *DescribePushMeteringDataResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) SetInstanceId(v string) *DescribePushMeteringDataResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) SetIsPushBilling(v bool) *DescribePushMeteringDataResponseBodyResult {
	s.IsPushBilling = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) SetMeteringAssist(v string) *DescribePushMeteringDataResponseBodyResult {
	s.MeteringAssist = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) SetMeteringEntity(v string) *DescribePushMeteringDataResponseBodyResult {
	s.MeteringEntity = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) SetPushOrderBizId(v string) *DescribePushMeteringDataResponseBodyResult {
	s.PushOrderBizId = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) SetStartTime(v int64) *DescribePushMeteringDataResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *DescribePushMeteringDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
