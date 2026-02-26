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
	// This parameter is required.
	AddressDetail *string `json:"addressDetail,omitempty" xml:"addressDetail,omitempty"`
	// example:
	//
	// 0
	AddressId *int64 `json:"addressId,omitempty" xml:"addressId,omitempty"`
	// example:
	//
	// 610102
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	Receiver *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 182***5674
	ReceiverPhone *string `json:"receiverPhone,omitempty" xml:"receiverPhone,omitempty"`
	// example:
	//
	// 61010212
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
