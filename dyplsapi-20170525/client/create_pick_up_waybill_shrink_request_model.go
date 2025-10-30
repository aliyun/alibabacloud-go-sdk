// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppointGotEndTime(v string) *CreatePickUpWaybillShrinkRequest
	GetAppointGotEndTime() *string
	SetAppointGotStartTime(v string) *CreatePickUpWaybillShrinkRequest
	GetAppointGotStartTime() *string
	SetBizType(v int32) *CreatePickUpWaybillShrinkRequest
	GetBizType() *int32
	SetConsigneeAddressShrink(v string) *CreatePickUpWaybillShrinkRequest
	GetConsigneeAddressShrink() *string
	SetConsigneeMobile(v string) *CreatePickUpWaybillShrinkRequest
	GetConsigneeMobile() *string
	SetConsigneeName(v string) *CreatePickUpWaybillShrinkRequest
	GetConsigneeName() *string
	SetConsigneePhone(v string) *CreatePickUpWaybillShrinkRequest
	GetConsigneePhone() *string
	SetCpCode(v string) *CreatePickUpWaybillShrinkRequest
	GetCpCode() *string
	SetGoodsInfosShrink(v string) *CreatePickUpWaybillShrinkRequest
	GetGoodsInfosShrink() *string
	SetOrderChannels(v string) *CreatePickUpWaybillShrinkRequest
	GetOrderChannels() *string
	SetOuterOrderCode(v string) *CreatePickUpWaybillShrinkRequest
	GetOuterOrderCode() *string
	SetRemark(v string) *CreatePickUpWaybillShrinkRequest
	GetRemark() *string
	SetSendAddressShrink(v string) *CreatePickUpWaybillShrinkRequest
	GetSendAddressShrink() *string
	SetSendMobile(v string) *CreatePickUpWaybillShrinkRequest
	GetSendMobile() *string
	SetSendName(v string) *CreatePickUpWaybillShrinkRequest
	GetSendName() *string
	SetSendPhone(v string) *CreatePickUpWaybillShrinkRequest
	GetSendPhone() *string
}

type CreatePickUpWaybillShrinkRequest struct {
	// The end time of the door-to-door pickup in the appointment. The value of **AppointGotEndTime*	- is the value of **EndTime*	- of **AppointTimes*	- in **CpTimeSelectList*	- returned by the [CreatePickUpWaybillPreQuery](~~CreatePickUpWaybillPreQuery~~#resultMapping) operation.
	//
	// >  This parameter is required when **BizType*	- is set to **1**.
	//
	// example:
	//
	// 2021-01-01 12:00:00
	AppointGotEndTime *string `json:"AppointGotEndTime,omitempty" xml:"AppointGotEndTime,omitempty"`
	// The start time of the door-to-door pickup in the appointment. The value of **AppointGotStartTime*	- is the value of **StartTime*	- of **AppointTimes*	- in **CpTimeSelectList*	- returned by the [CreatePickUpWaybillPreQuery](~~CreatePickUpWaybillPreQuery~~#resultMapping) operation.
	//
	// >  This parameter is required when **BizType*	- is set to **1**.
	//
	// example:
	//
	// 2021-01-01 10:00:00
	AppointGotStartTime *string `json:"AppointGotStartTime,omitempty" xml:"AppointGotStartTime,omitempty"`
	// The pickup mode. Valid values:
	//
	// 	- **0*	- (default): real-time order.
	//
	// 	- **1**: appointment order.
	//
	// example:
	//
	// 0
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The address of the consignee.
	//
	// This parameter is required.
	ConsigneeAddressShrink *string `json:"ConsigneeAddress,omitempty" xml:"ConsigneeAddress,omitempty"`
	// The mobile phone number of the consignee.
	//
	// >  Either ConsigneeMobile or ConsigneePhone must be set.
	//
	// example:
	//
	// 1580000****
	ConsigneeMobile *string `json:"ConsigneeMobile,omitempty" xml:"ConsigneeMobile,omitempty"`
	// The name of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// Li
	ConsigneeName *string `json:"ConsigneeName,omitempty" xml:"ConsigneeName,omitempty"`
	// The landline phone number of the consignee.
	//
	// >  Either ConsigneeMobile or ConsigneePhone must be set.
	//
	// example:
	//
	// 0570000****
	ConsigneePhone *string `json:"ConsigneePhone,omitempty" xml:"ConsigneePhone,omitempty"`
	// The code of the courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The items.
	GoodsInfosShrink *string `json:"GoodsInfos,omitempty" xml:"GoodsInfos,omitempty"`
	// The external channel sources.
	//
	// This parameter is required.
	//
	// example:
	//
	// YUN_DIAN_SHANG
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The ID of the external order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 143234234266****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The additional information about the order. The additional information will be printed on the order.
	//
	// example:
	//
	// fragile
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The address of the sender.
	//
	// This parameter is required.
	SendAddressShrink *string `json:"SendAddress,omitempty" xml:"SendAddress,omitempty"`
	// The mobile phone number of the sender.
	//
	// >  Either SendMobile or SendPhone must be set.
	//
	// example:
	//
	// 1596714****
	SendMobile *string `json:"SendMobile,omitempty" xml:"SendMobile,omitempty"`
	// The name of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// Wang
	SendName *string `json:"SendName,omitempty" xml:"SendName,omitempty"`
	// The landline phone number of the sender.
	//
	// >  Either SendMobile or SendPhone must be set.
	//
	// example:
	//
	// 05718845****
	SendPhone *string `json:"SendPhone,omitempty" xml:"SendPhone,omitempty"`
}

func (s CreatePickUpWaybillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillShrinkRequest) GetAppointGotEndTime() *string {
	return s.AppointGotEndTime
}

func (s *CreatePickUpWaybillShrinkRequest) GetAppointGotStartTime() *string {
	return s.AppointGotStartTime
}

func (s *CreatePickUpWaybillShrinkRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *CreatePickUpWaybillShrinkRequest) GetConsigneeAddressShrink() *string {
	return s.ConsigneeAddressShrink
}

func (s *CreatePickUpWaybillShrinkRequest) GetConsigneeMobile() *string {
	return s.ConsigneeMobile
}

func (s *CreatePickUpWaybillShrinkRequest) GetConsigneeName() *string {
	return s.ConsigneeName
}

func (s *CreatePickUpWaybillShrinkRequest) GetConsigneePhone() *string {
	return s.ConsigneePhone
}

func (s *CreatePickUpWaybillShrinkRequest) GetCpCode() *string {
	return s.CpCode
}

func (s *CreatePickUpWaybillShrinkRequest) GetGoodsInfosShrink() *string {
	return s.GoodsInfosShrink
}

func (s *CreatePickUpWaybillShrinkRequest) GetOrderChannels() *string {
	return s.OrderChannels
}

func (s *CreatePickUpWaybillShrinkRequest) GetOuterOrderCode() *string {
	return s.OuterOrderCode
}

func (s *CreatePickUpWaybillShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreatePickUpWaybillShrinkRequest) GetSendAddressShrink() *string {
	return s.SendAddressShrink
}

func (s *CreatePickUpWaybillShrinkRequest) GetSendMobile() *string {
	return s.SendMobile
}

func (s *CreatePickUpWaybillShrinkRequest) GetSendName() *string {
	return s.SendName
}

func (s *CreatePickUpWaybillShrinkRequest) GetSendPhone() *string {
	return s.SendPhone
}

func (s *CreatePickUpWaybillShrinkRequest) SetAppointGotEndTime(v string) *CreatePickUpWaybillShrinkRequest {
	s.AppointGotEndTime = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetAppointGotStartTime(v string) *CreatePickUpWaybillShrinkRequest {
	s.AppointGotStartTime = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetBizType(v int32) *CreatePickUpWaybillShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneeAddressShrink(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneeAddressShrink = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneeMobile(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneeMobile = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneeName(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneeName = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetConsigneePhone(v string) *CreatePickUpWaybillShrinkRequest {
	s.ConsigneePhone = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetCpCode(v string) *CreatePickUpWaybillShrinkRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetGoodsInfosShrink(v string) *CreatePickUpWaybillShrinkRequest {
	s.GoodsInfosShrink = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetOrderChannels(v string) *CreatePickUpWaybillShrinkRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillShrinkRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetRemark(v string) *CreatePickUpWaybillShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendAddressShrink(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendAddressShrink = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendMobile(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendMobile = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendName(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendName = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) SetSendPhone(v string) *CreatePickUpWaybillShrinkRequest {
	s.SendPhone = &v
	return s
}

func (s *CreatePickUpWaybillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
