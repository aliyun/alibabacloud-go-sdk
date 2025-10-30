// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppointGotEndTime(v string) *CreatePickUpWaybillRequest
	GetAppointGotEndTime() *string
	SetAppointGotStartTime(v string) *CreatePickUpWaybillRequest
	GetAppointGotStartTime() *string
	SetBizType(v int32) *CreatePickUpWaybillRequest
	GetBizType() *int32
	SetConsigneeAddress(v *CreatePickUpWaybillRequestConsigneeAddress) *CreatePickUpWaybillRequest
	GetConsigneeAddress() *CreatePickUpWaybillRequestConsigneeAddress
	SetConsigneeMobile(v string) *CreatePickUpWaybillRequest
	GetConsigneeMobile() *string
	SetConsigneeName(v string) *CreatePickUpWaybillRequest
	GetConsigneeName() *string
	SetConsigneePhone(v string) *CreatePickUpWaybillRequest
	GetConsigneePhone() *string
	SetCpCode(v string) *CreatePickUpWaybillRequest
	GetCpCode() *string
	SetGoodsInfos(v []*CreatePickUpWaybillRequestGoodsInfos) *CreatePickUpWaybillRequest
	GetGoodsInfos() []*CreatePickUpWaybillRequestGoodsInfos
	SetOrderChannels(v string) *CreatePickUpWaybillRequest
	GetOrderChannels() *string
	SetOuterOrderCode(v string) *CreatePickUpWaybillRequest
	GetOuterOrderCode() *string
	SetRemark(v string) *CreatePickUpWaybillRequest
	GetRemark() *string
	SetSendAddress(v *CreatePickUpWaybillRequestSendAddress) *CreatePickUpWaybillRequest
	GetSendAddress() *CreatePickUpWaybillRequestSendAddress
	SetSendMobile(v string) *CreatePickUpWaybillRequest
	GetSendMobile() *string
	SetSendName(v string) *CreatePickUpWaybillRequest
	GetSendName() *string
	SetSendPhone(v string) *CreatePickUpWaybillRequest
	GetSendPhone() *string
}

type CreatePickUpWaybillRequest struct {
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
	ConsigneeAddress *CreatePickUpWaybillRequestConsigneeAddress `json:"ConsigneeAddress,omitempty" xml:"ConsigneeAddress,omitempty" type:"Struct"`
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
	GoodsInfos []*CreatePickUpWaybillRequestGoodsInfos `json:"GoodsInfos,omitempty" xml:"GoodsInfos,omitempty" type:"Repeated"`
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
	SendAddress *CreatePickUpWaybillRequestSendAddress `json:"SendAddress,omitempty" xml:"SendAddress,omitempty" type:"Struct"`
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

func (s CreatePickUpWaybillRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequest) GetAppointGotEndTime() *string {
	return s.AppointGotEndTime
}

func (s *CreatePickUpWaybillRequest) GetAppointGotStartTime() *string {
	return s.AppointGotStartTime
}

func (s *CreatePickUpWaybillRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *CreatePickUpWaybillRequest) GetConsigneeAddress() *CreatePickUpWaybillRequestConsigneeAddress {
	return s.ConsigneeAddress
}

func (s *CreatePickUpWaybillRequest) GetConsigneeMobile() *string {
	return s.ConsigneeMobile
}

func (s *CreatePickUpWaybillRequest) GetConsigneeName() *string {
	return s.ConsigneeName
}

func (s *CreatePickUpWaybillRequest) GetConsigneePhone() *string {
	return s.ConsigneePhone
}

func (s *CreatePickUpWaybillRequest) GetCpCode() *string {
	return s.CpCode
}

func (s *CreatePickUpWaybillRequest) GetGoodsInfos() []*CreatePickUpWaybillRequestGoodsInfos {
	return s.GoodsInfos
}

func (s *CreatePickUpWaybillRequest) GetOrderChannels() *string {
	return s.OrderChannels
}

func (s *CreatePickUpWaybillRequest) GetOuterOrderCode() *string {
	return s.OuterOrderCode
}

func (s *CreatePickUpWaybillRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreatePickUpWaybillRequest) GetSendAddress() *CreatePickUpWaybillRequestSendAddress {
	return s.SendAddress
}

func (s *CreatePickUpWaybillRequest) GetSendMobile() *string {
	return s.SendMobile
}

func (s *CreatePickUpWaybillRequest) GetSendName() *string {
	return s.SendName
}

func (s *CreatePickUpWaybillRequest) GetSendPhone() *string {
	return s.SendPhone
}

func (s *CreatePickUpWaybillRequest) SetAppointGotEndTime(v string) *CreatePickUpWaybillRequest {
	s.AppointGotEndTime = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetAppointGotStartTime(v string) *CreatePickUpWaybillRequest {
	s.AppointGotStartTime = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetBizType(v int32) *CreatePickUpWaybillRequest {
	s.BizType = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneeAddress(v *CreatePickUpWaybillRequestConsigneeAddress) *CreatePickUpWaybillRequest {
	s.ConsigneeAddress = v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneeMobile(v string) *CreatePickUpWaybillRequest {
	s.ConsigneeMobile = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneeName(v string) *CreatePickUpWaybillRequest {
	s.ConsigneeName = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetConsigneePhone(v string) *CreatePickUpWaybillRequest {
	s.ConsigneePhone = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetCpCode(v string) *CreatePickUpWaybillRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetGoodsInfos(v []*CreatePickUpWaybillRequestGoodsInfos) *CreatePickUpWaybillRequest {
	s.GoodsInfos = v
	return s
}

func (s *CreatePickUpWaybillRequest) SetOrderChannels(v string) *CreatePickUpWaybillRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetRemark(v string) *CreatePickUpWaybillRequest {
	s.Remark = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendAddress(v *CreatePickUpWaybillRequestSendAddress) *CreatePickUpWaybillRequest {
	s.SendAddress = v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendMobile(v string) *CreatePickUpWaybillRequest {
	s.SendMobile = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendName(v string) *CreatePickUpWaybillRequest {
	s.SendName = &v
	return s
}

func (s *CreatePickUpWaybillRequest) SetSendPhone(v string) *CreatePickUpWaybillRequest {
	s.SendPhone = &v
	return s
}

func (s *CreatePickUpWaybillRequest) Validate() error {
	if s.ConsigneeAddress != nil {
		if err := s.ConsigneeAddress.Validate(); err != nil {
			return err
		}
	}
	if s.GoodsInfos != nil {
		for _, item := range s.GoodsInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SendAddress != nil {
		if err := s.SendAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillRequestConsigneeAddress struct {
	// The detailed address of the consignee.
	//
	// This parameter is required.
	//
	// example:
	//
	// XX community
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the consignee is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// xihu
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the consignee is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the consignee is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// zhejiang
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the consignee is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillRequestConsigneeAddress) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillRequestConsigneeAddress) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) GetAddressDetail() *string {
	return s.AddressDetail
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) GetAreaName() *string {
	return s.AreaName
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) GetCityName() *string {
	return s.CityName
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) GetTownName() *string {
	return s.TownName
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetAddressDetail(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetAreaName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetCityName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetProvinceName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) SetTownName(v string) *CreatePickUpWaybillRequestConsigneeAddress {
	s.TownName = &v
	return s
}

func (s *CreatePickUpWaybillRequestConsigneeAddress) Validate() error {
	return dara.Validate(s)
}

type CreatePickUpWaybillRequestGoodsInfos struct {
	// The item name.
	//
	// example:
	//
	// zhang
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The item quantity.
	//
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The item weight. Unit: gram.
	//
	// example:
	//
	// 1000
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreatePickUpWaybillRequestGoodsInfos) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillRequestGoodsInfos) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequestGoodsInfos) GetName() *string {
	return s.Name
}

func (s *CreatePickUpWaybillRequestGoodsInfos) GetQuantity() *string {
	return s.Quantity
}

func (s *CreatePickUpWaybillRequestGoodsInfos) GetWeight() *string {
	return s.Weight
}

func (s *CreatePickUpWaybillRequestGoodsInfos) SetName(v string) *CreatePickUpWaybillRequestGoodsInfos {
	s.Name = &v
	return s
}

func (s *CreatePickUpWaybillRequestGoodsInfos) SetQuantity(v string) *CreatePickUpWaybillRequestGoodsInfos {
	s.Quantity = &v
	return s
}

func (s *CreatePickUpWaybillRequestGoodsInfos) SetWeight(v string) *CreatePickUpWaybillRequestGoodsInfos {
	s.Weight = &v
	return s
}

func (s *CreatePickUpWaybillRequestGoodsInfos) Validate() error {
	return dara.Validate(s)
}

type CreatePickUpWaybillRequestSendAddress struct {
	// The detailed address of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// XX community
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the sender is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// wenjiang
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the sender is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// chengdu
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the sender is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sichuan
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the sender is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillRequestSendAddress) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillRequestSendAddress) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillRequestSendAddress) GetAddressDetail() *string {
	return s.AddressDetail
}

func (s *CreatePickUpWaybillRequestSendAddress) GetAreaName() *string {
	return s.AreaName
}

func (s *CreatePickUpWaybillRequestSendAddress) GetCityName() *string {
	return s.CityName
}

func (s *CreatePickUpWaybillRequestSendAddress) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *CreatePickUpWaybillRequestSendAddress) GetTownName() *string {
	return s.TownName
}

func (s *CreatePickUpWaybillRequestSendAddress) SetAddressDetail(v string) *CreatePickUpWaybillRequestSendAddress {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetAreaName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetCityName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetProvinceName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) SetTownName(v string) *CreatePickUpWaybillRequestSendAddress {
	s.TownName = &v
	return s
}

func (s *CreatePickUpWaybillRequestSendAddress) Validate() error {
	return dara.Validate(s)
}
