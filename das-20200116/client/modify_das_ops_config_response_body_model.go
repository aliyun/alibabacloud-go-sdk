// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDasOpsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDasOpsConfigResponseBody
	GetCode() *string
	SetData(v *ModifyDasOpsConfigResponseBodyData) *ModifyDasOpsConfigResponseBody
	GetData() *ModifyDasOpsConfigResponseBodyData
	SetMessage(v string) *ModifyDasOpsConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDasOpsConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyDasOpsConfigResponseBody
	GetSuccess() *string
}

type ModifyDasOpsConfigResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// SqlLogConfig
	Data *ModifyDasOpsConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDasOpsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasOpsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDasOpsConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDasOpsConfigResponseBody) GetData() *ModifyDasOpsConfigResponseBodyData {
	return s.Data
}

func (s *ModifyDasOpsConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDasOpsConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDasOpsConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyDasOpsConfigResponseBody) SetCode(v string) *ModifyDasOpsConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBody) SetData(v *ModifyDasOpsConfigResponseBodyData) *ModifyDasOpsConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDasOpsConfigResponseBody) SetMessage(v string) *ModifyDasOpsConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBody) SetRequestId(v string) *ModifyDasOpsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBody) SetSuccess(v string) *ModifyDasOpsConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDasOpsConfigResponseBodyData struct {
	// The payment method.
	//
	// example:
	//
	// prepay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The Alibaba Cloud Managed Services instance ID.
	//
	// example:
	//
	// pc-2zelo5v5u0678jx04
	CommodityInstanceId *string `json:"CommodityInstanceId,omitempty" xml:"CommodityInstanceId,omitempty"`
	// Indicates whether DAS Economy Edition is enabled.
	//
	// example:
	//
	// true
	EcoEnable *bool `json:"EcoEnable,omitempty" xml:"EcoEnable,omitempty"`
	// Indicates whether the Alibaba Cloud Managed Services feature is enabled (including DAS Economy Edition).
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The end time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1608888296000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether Alibaba Cloud Managed Services is enabled.
	OpsEnable *bool `json:"OpsEnable,omitempty" xml:"OpsEnable,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 285412912420536
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1672531200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - **INIT**: Pending scheduling.
	//
	// - **RUNNING**: Running.
	//
	// - **FAILED**: Failed.
	//
	// - **CANCELED**: Canceled.
	//
	// - **COMPLETED**: Completed.
	//
	// > When the task is in the **COMPLETED*	- state, you can view the task result.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyDasOpsConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasOpsConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDasOpsConfigResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyDasOpsConfigResponseBodyData) GetCommodityInstanceId() *string {
	return s.CommodityInstanceId
}

func (s *ModifyDasOpsConfigResponseBodyData) GetEcoEnable() *bool {
	return s.EcoEnable
}

func (s *ModifyDasOpsConfigResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyDasOpsConfigResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyDasOpsConfigResponseBodyData) GetOpsEnable() *bool {
	return s.OpsEnable
}

func (s *ModifyDasOpsConfigResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyDasOpsConfigResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifyDasOpsConfigResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ModifyDasOpsConfigResponseBodyData) SetChargeType(v string) *ModifyDasOpsConfigResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetCommodityInstanceId(v string) *ModifyDasOpsConfigResponseBodyData {
	s.CommodityInstanceId = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetEcoEnable(v bool) *ModifyDasOpsConfigResponseBodyData {
	s.EcoEnable = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetEnable(v bool) *ModifyDasOpsConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetEndTime(v int64) *ModifyDasOpsConfigResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetOpsEnable(v bool) *ModifyDasOpsConfigResponseBodyData {
	s.OpsEnable = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetOrderId(v int64) *ModifyDasOpsConfigResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetStartTime(v int64) *ModifyDasOpsConfigResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) SetStatus(v string) *ModifyDasOpsConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModifyDasOpsConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
