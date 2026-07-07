// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAddressDetail(v string) *AddressInfo
	GetAddressDetail() *string
	SetAddressId(v int64) *AddressInfo
	GetAddressId() *int64
	SetDivisionCode(v string) *AddressInfo
	GetDivisionCode() *string
	SetReceiver(v string) *AddressInfo
	GetReceiver() *string
	SetReceiverPhone(v string) *AddressInfo
	GetReceiverPhone() *string
	SetTownDivisionCode(v string) *AddressInfo
	GetTownDivisionCode() *string
}

type AddressInfo struct {
	// The detailed shipping address. Enter the full address in the format of province, city, district/county, street, and community.
	//
	// This parameter is required.
	//
	// example:
	//
	// 陕西省西安市新城区xx街道xxx大厦xx室
	AddressDetail *string `json:"addressDetail,omitempty" xml:"addressDetail,omitempty"`
	// The address ID.
	//
	// example:
	//
	// 0
	AddressId *int64 `json:"addressId,omitempty" xml:"addressId,omitempty"`
	// The level-4 address code for the district or county. This parameter is recommended.
	//
	// example:
	//
	// 330106
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// The recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// 任先生
	Receiver *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// The phone number of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// 182***5674
	ReceiverPhone *string `json:"receiverPhone,omitempty" xml:"receiverPhone,omitempty"`
	// The level-5 address code for the town or street. This parameter is required.
	//
	// example:
	//
	// 330106109
	TownDivisionCode *string `json:"townDivisionCode,omitempty" xml:"townDivisionCode,omitempty"`
}

func (s AddressInfo) String() string {
	return dara.Prettify(s)
}

func (s AddressInfo) GoString() string {
	return s.String()
}

func (s *AddressInfo) GetAddressDetail() *string {
	return s.AddressDetail
}

func (s *AddressInfo) GetAddressId() *int64 {
	return s.AddressId
}

func (s *AddressInfo) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *AddressInfo) GetReceiver() *string {
	return s.Receiver
}

func (s *AddressInfo) GetReceiverPhone() *string {
	return s.ReceiverPhone
}

func (s *AddressInfo) GetTownDivisionCode() *string {
	return s.TownDivisionCode
}

func (s *AddressInfo) SetAddressDetail(v string) *AddressInfo {
	s.AddressDetail = &v
	return s
}

func (s *AddressInfo) SetAddressId(v int64) *AddressInfo {
	s.AddressId = &v
	return s
}

func (s *AddressInfo) SetDivisionCode(v string) *AddressInfo {
	s.DivisionCode = &v
	return s
}

func (s *AddressInfo) SetReceiver(v string) *AddressInfo {
	s.Receiver = &v
	return s
}

func (s *AddressInfo) SetReceiverPhone(v string) *AddressInfo {
	s.ReceiverPhone = &v
	return s
}

func (s *AddressInfo) SetTownDivisionCode(v string) *AddressInfo {
	s.TownDivisionCode = &v
	return s
}

func (s *AddressInfo) Validate() error {
	return dara.Validate(s)
}
