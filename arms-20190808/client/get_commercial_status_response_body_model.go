// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommercialStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetCommercialStatusResponseBody
	GetRequestId() *string
	SetUserAndCommodityStatus(v *GetCommercialStatusResponseBodyUserAndCommodityStatus) *GetCommercialStatusResponseBody
	GetUserAndCommodityStatus() *GetCommercialStatusResponseBodyUserAndCommodityStatus
}

type GetCommercialStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 49C82193-E991-5F6A-AF3E-1664D8D05CA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The commercialization status of the service.
	UserAndCommodityStatus *GetCommercialStatusResponseBodyUserAndCommodityStatus `json:"UserAndCommodityStatus,omitempty" xml:"UserAndCommodityStatus,omitempty" type:"Struct"`
}

func (s GetCommercialStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommercialStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommercialStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommercialStatusResponseBody) GetUserAndCommodityStatus() *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	return s.UserAndCommodityStatus
}

func (s *GetCommercialStatusResponseBody) SetRequestId(v string) *GetCommercialStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommercialStatusResponseBody) SetUserAndCommodityStatus(v *GetCommercialStatusResponseBodyUserAndCommodityStatus) *GetCommercialStatusResponseBody {
	s.UserAndCommodityStatus = v
	return s
}

func (s *GetCommercialStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCommercialStatusResponseBodyUserAndCommodityStatus struct {
	// Indicates whether you are using the Basic Edition.
	//
	// example:
	//
	// false
	Basic *bool `json:"Basic,omitempty" xml:"Basic,omitempty"`
	// The billing method.
	//
	// example:
	//
	// usage
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Indicates whether the service is activated.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The additional information.
	//
	// example:
	//
	// info
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The number of days during which the service is free of charge.
	//
	// example:
	//
	// 10
	FreeDays *int64 `json:"FreeDays,omitempty" xml:"FreeDays,omitempty"`
	// The tags.
	//
	// example:
	//
	// NEW
	Lable *string `json:"Lable,omitempty" xml:"Lable,omitempty"`
	// The commercialization status.
	//
	// Valid values:
	//
	// 	- Normal: The service is activated.
	//
	// 	- Abnormal: An exception occurs during activation.
	//
	// 	- Free: The service is not activated.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCommercialStatusResponseBodyUserAndCommodityStatus) String() string {
	return dara.Prettify(s)
}

func (s GetCommercialStatusResponseBodyUserAndCommodityStatus) GoString() string {
	return s.String()
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetBasic() *bool {
	return s.Basic
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetEnable() *bool {
	return s.Enable
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetFreeDays() *int64 {
	return s.FreeDays
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetLable() *string {
	return s.Lable
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) GetStatus() *string {
	return s.Status
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetBasic(v bool) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.Basic = &v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetChargeType(v string) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.ChargeType = &v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetEnable(v bool) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.Enable = &v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetExtraInfo(v map[string]interface{}) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.ExtraInfo = v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetFreeDays(v int64) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.FreeDays = &v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetLable(v string) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.Lable = &v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) SetStatus(v string) *GetCommercialStatusResponseBodyUserAndCommodityStatus {
	s.Status = &v
	return s
}

func (s *GetCommercialStatusResponseBodyUserAndCommodityStatus) Validate() error {
	return dara.Validate(s)
}
