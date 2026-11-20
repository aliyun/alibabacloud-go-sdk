// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePickUpWaybillPreQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsigneeInfo(v *CreatePickUpWaybillPreQueryRequestConsigneeInfo) *CreatePickUpWaybillPreQueryRequest
	GetConsigneeInfo() *CreatePickUpWaybillPreQueryRequestConsigneeInfo
	SetCpCode(v string) *CreatePickUpWaybillPreQueryRequest
	GetCpCode() *string
	SetOrderChannels(v string) *CreatePickUpWaybillPreQueryRequest
	GetOrderChannels() *string
	SetOuterOrderCode(v string) *CreatePickUpWaybillPreQueryRequest
	GetOuterOrderCode() *string
	SetPreWeight(v string) *CreatePickUpWaybillPreQueryRequest
	GetPreWeight() *string
	SetSenderInfo(v *CreatePickUpWaybillPreQueryRequestSenderInfo) *CreatePickUpWaybillPreQueryRequest
	GetSenderInfo() *CreatePickUpWaybillPreQueryRequestSenderInfo
}

type CreatePickUpWaybillPreQueryRequest struct {
	// The recipient information.
	//
	// This parameter is required.
	ConsigneeInfo *CreatePickUpWaybillPreQueryRequestConsigneeInfo `json:"ConsigneeInfo,omitempty" xml:"ConsigneeInfo,omitempty" type:"Struct"`
	// The courier company code. If not specified, the system automatically assigns a courier company.
	//
	// example:
	//
	// YTO
	CpCode *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	// The external channel source identifier. The value cannot contain underscores. The value is agreed upon with the system during integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	OrderChannels *string `json:"OrderChannels,omitempty" xml:"OrderChannels,omitempty"`
	// The order number in the caller\\"s system.
	//
	// example:
	//
	// 787DFHHDS989****
	OuterOrderCode *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	// The estimated weight, in grams.
	//
	// > If you want to query the estimated price, this parameter is required.
	//
	// example:
	//
	// 2000
	PreWeight *string `json:"PreWeight,omitempty" xml:"PreWeight,omitempty"`
	// The sender information.
	//
	// This parameter is required.
	SenderInfo *CreatePickUpWaybillPreQueryRequestSenderInfo `json:"SenderInfo,omitempty" xml:"SenderInfo,omitempty" type:"Struct"`
}

func (s CreatePickUpWaybillPreQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequest) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequest) GetConsigneeInfo() *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	return s.ConsigneeInfo
}

func (s *CreatePickUpWaybillPreQueryRequest) GetCpCode() *string {
	return s.CpCode
}

func (s *CreatePickUpWaybillPreQueryRequest) GetOrderChannels() *string {
	return s.OrderChannels
}

func (s *CreatePickUpWaybillPreQueryRequest) GetOuterOrderCode() *string {
	return s.OuterOrderCode
}

func (s *CreatePickUpWaybillPreQueryRequest) GetPreWeight() *string {
	return s.PreWeight
}

func (s *CreatePickUpWaybillPreQueryRequest) GetSenderInfo() *CreatePickUpWaybillPreQueryRequestSenderInfo {
	return s.SenderInfo
}

func (s *CreatePickUpWaybillPreQueryRequest) SetConsigneeInfo(v *CreatePickUpWaybillPreQueryRequestConsigneeInfo) *CreatePickUpWaybillPreQueryRequest {
	s.ConsigneeInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetCpCode(v string) *CreatePickUpWaybillPreQueryRequest {
	s.CpCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetOrderChannels(v string) *CreatePickUpWaybillPreQueryRequest {
	s.OrderChannels = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetOuterOrderCode(v string) *CreatePickUpWaybillPreQueryRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetPreWeight(v string) *CreatePickUpWaybillPreQueryRequest {
	s.PreWeight = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) SetSenderInfo(v *CreatePickUpWaybillPreQueryRequestSenderInfo) *CreatePickUpWaybillPreQueryRequest {
	s.SenderInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequest) Validate() error {
	if s.ConsigneeInfo != nil {
		if err := s.ConsigneeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SenderInfo != nil {
		if err := s.SenderInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryRequestConsigneeInfo struct {
	// The address information of the recipient.
	AddressInfo *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty" type:"Struct"`
	// The mobile phone number of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// Li XX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) GetAddressInfo() *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	return s.AddressInfo
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) GetMobile() *string {
	return s.Mobile
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) GetName() *string {
	return s.Name
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) SetAddressInfo(v *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	s.AddressInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) SetMobile(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	s.Mobile = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) SetName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfo {
	s.Name = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfo) Validate() error {
	if s.AddressInfo != nil {
		if err := s.AddressInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo struct {
	// The detailed address of the recipient.
	//
	// example:
	//
	// XX Community, No. XX
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the recipient is located.
	//
	// example:
	//
	// Chang\\"an District
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the recipient is located.
	//
	// example:
	//
	// Xi\\"an
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the recipient is located.
	//
	// example:
	//
	// Shaanxi Province
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the recipient is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GetAddressDetail() *string {
	return s.AddressDetail
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GetAreaName() *string {
	return s.AreaName
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GetCityName() *string {
	return s.CityName
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) GetTownName() *string {
	return s.TownName
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetAddressDetail(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetAreaName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetCityName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetProvinceName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) SetTownName(v string) *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo {
	s.TownName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestConsigneeInfoAddressInfo) Validate() error {
	return dara.Validate(s)
}

type CreatePickUpWaybillPreQueryRequestSenderInfo struct {
	// The address information of the sender.
	AddressInfo *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty" type:"Struct"`
	// The mobile phone number of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the sender.
	//
	// This parameter is required.
	//
	// example:
	//
	// Wang XX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) GetAddressInfo() *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	return s.AddressInfo
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) GetMobile() *string {
	return s.Mobile
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) GetName() *string {
	return s.Name
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) SetAddressInfo(v *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) *CreatePickUpWaybillPreQueryRequestSenderInfo {
	s.AddressInfo = v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) SetMobile(v string) *CreatePickUpWaybillPreQueryRequestSenderInfo {
	s.Mobile = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) SetName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfo {
	s.Name = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfo) Validate() error {
	if s.AddressInfo != nil {
		if err := s.AddressInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo struct {
	// The detailed address of the sender.
	//
	// example:
	//
	// XX Community, No. XX
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	// The district where the sender is located.
	//
	// example:
	//
	// Xihu District
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// The city where the sender is located.
	//
	// example:
	//
	// Hangzhou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The province where the sender is located.
	//
	// example:
	//
	// Zhejiang Province
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	// The street where the sender is located.
	//
	// example:
	//
	// XX Street
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) String() string {
	return dara.Prettify(s)
}

func (s CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GoString() string {
	return s.String()
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GetAddressDetail() *string {
	return s.AddressDetail
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GetAreaName() *string {
	return s.AreaName
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GetCityName() *string {
	return s.CityName
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) GetTownName() *string {
	return s.TownName
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetAddressDetail(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.AddressDetail = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetAreaName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.AreaName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetCityName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.CityName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetProvinceName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.ProvinceName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) SetTownName(v string) *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo {
	s.TownName = &v
	return s
}

func (s *CreatePickUpWaybillPreQueryRequestSenderInfoAddressInfo) Validate() error {
	return dara.Validate(s)
}
