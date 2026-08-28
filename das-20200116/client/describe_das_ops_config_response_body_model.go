// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDasOpsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDasOpsConfigResponseBody
	GetCode() *string
	SetData(v *DescribeDasOpsConfigResponseBodyData) *DescribeDasOpsConfigResponseBody
	GetData() *DescribeDasOpsConfigResponseBodyData
	SetMessage(v string) *DescribeDasOpsConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDasOpsConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDasOpsConfigResponseBody
	GetSuccess() *string
}

type DescribeDasOpsConfigResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// SqlLogConfig
	Data *DescribeDasOpsConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D00DB161-FEF6-5428-B37A-8D29A4C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDasOpsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasOpsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDasOpsConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDasOpsConfigResponseBody) GetData() *DescribeDasOpsConfigResponseBodyData {
	return s.Data
}

func (s *DescribeDasOpsConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDasOpsConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDasOpsConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDasOpsConfigResponseBody) SetCode(v string) *DescribeDasOpsConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBody) SetData(v *DescribeDasOpsConfigResponseBodyData) *DescribeDasOpsConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDasOpsConfigResponseBody) SetMessage(v string) *DescribeDasOpsConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBody) SetRequestId(v string) *DescribeDasOpsConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBody) SetSuccess(v string) *DescribeDasOpsConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDasOpsConfigResponseBodyData struct {
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// 	- **true**: Auto-renewal is enabled.
	//
	// 	- **false*	- (default): Auto-renewal is disabled.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The payment method.
	//
	// example:
	//
	// http://prodpopscanGGfYbZif.302prod.xaliyun.com
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
	// false
	EcoEnable *bool `json:"EcoEnable,omitempty" xml:"EcoEnable,omitempty"`
	// Indicates whether the Alibaba Cloud Managed Services is enabled.
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
	// Indicates whether the Alibaba Cloud Managed Services is enabled.
	OpsEnable *bool `json:"OpsEnable,omitempty" xml:"OpsEnable,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 275772887390786
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The start time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1672531200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDasOpsConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasOpsConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDasOpsConfigResponseBodyData) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeDasOpsConfigResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDasOpsConfigResponseBodyData) GetCommodityInstanceId() *string {
	return s.CommodityInstanceId
}

func (s *DescribeDasOpsConfigResponseBodyData) GetEcoEnable() *bool {
	return s.EcoEnable
}

func (s *DescribeDasOpsConfigResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeDasOpsConfigResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDasOpsConfigResponseBodyData) GetOpsEnable() *bool {
	return s.OpsEnable
}

func (s *DescribeDasOpsConfigResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeDasOpsConfigResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDasOpsConfigResponseBodyData) SetAutoRenew(v bool) *DescribeDasOpsConfigResponseBodyData {
	s.AutoRenew = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetChargeType(v string) *DescribeDasOpsConfigResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetCommodityInstanceId(v string) *DescribeDasOpsConfigResponseBodyData {
	s.CommodityInstanceId = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetEcoEnable(v bool) *DescribeDasOpsConfigResponseBodyData {
	s.EcoEnable = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetEnable(v bool) *DescribeDasOpsConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetEndTime(v int64) *DescribeDasOpsConfigResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetOpsEnable(v bool) *DescribeDasOpsConfigResponseBodyData {
	s.OpsEnable = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetOrderId(v int64) *DescribeDasOpsConfigResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) SetStartTime(v int64) *DescribeDasOpsConfigResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeDasOpsConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
