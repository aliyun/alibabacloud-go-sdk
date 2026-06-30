// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeliveryAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*DescribeDeliveryAddressResponseBodyAddresses) *DescribeDeliveryAddressResponseBody
	GetAddresses() []*DescribeDeliveryAddressResponseBodyAddresses
	SetRequestId(v string) *DescribeDeliveryAddressResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDeliveryAddressResponseBody
	GetTotalCount() *int32
}

type DescribeDeliveryAddressResponseBody struct {
	// The address list.
	Addresses []*DescribeDeliveryAddressResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The request ID. Alibaba Cloud generates this unique identifier for each request.
	//
	// example:
	//
	// 72481C12-69AB-5CE1-8A35-A8EFA921****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total count of addresses.
	//
	// > Up to 20 addresses are supported.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeliveryAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBody) GetAddresses() []*DescribeDeliveryAddressResponseBodyAddresses {
	return s.Addresses
}

func (s *DescribeDeliveryAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeliveryAddressResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDeliveryAddressResponseBody) SetAddresses(v []*DescribeDeliveryAddressResponseBodyAddresses) *DescribeDeliveryAddressResponseBody {
	s.Addresses = v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) SetRequestId(v string) *DescribeDeliveryAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) SetTotalCount(v int32) *DescribeDeliveryAddressResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBody) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeliveryAddressResponseBodyAddresses struct {
	// The area.
	Area *DescribeDeliveryAddressResponseBodyAddressesArea `json:"Area,omitempty" xml:"Area,omitempty" type:"Struct"`
	// The city.
	City *DescribeDeliveryAddressResponseBodyAddressesCity `json:"City,omitempty" xml:"City,omitempty" type:"Struct"`
	// The contact name.
	//
	// example:
	//
	// Alice
	Contacts *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	// Indicates whether this is the default address.
	//
	// example:
	//
	// true
	DefaultAddress *bool `json:"DefaultAddress,omitempty" xml:"DefaultAddress,omitempty"`
	// The detailed address, which should include the street name and house number.
	//
	// example:
	//
	// 灯彩街1008号云园区2号楼测试中心
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 1381111****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The postal code.
	//
	// example:
	//
	// 03****
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// The province.
	Province *DescribeDeliveryAddressResponseBodyAddressesProvince `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	// The town.
	Town *DescribeDeliveryAddressResponseBodyAddressesTown `json:"Town,omitempty" xml:"Town,omitempty" type:"Struct"`
}

func (s DescribeDeliveryAddressResponseBodyAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetArea() *DescribeDeliveryAddressResponseBodyAddressesArea {
	return s.Area
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetCity() *DescribeDeliveryAddressResponseBodyAddressesCity {
	return s.City
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetContacts() *string {
	return s.Contacts
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetDefaultAddress() *bool {
	return s.DefaultAddress
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetDetail() *string {
	return s.Detail
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetMobile() *string {
	return s.Mobile
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetPostalCode() *string {
	return s.PostalCode
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetProvince() *DescribeDeliveryAddressResponseBodyAddressesProvince {
	return s.Province
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) GetTown() *DescribeDeliveryAddressResponseBodyAddressesTown {
	return s.Town
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetArea(v *DescribeDeliveryAddressResponseBodyAddressesArea) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Area = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetCity(v *DescribeDeliveryAddressResponseBodyAddressesCity) *DescribeDeliveryAddressResponseBodyAddresses {
	s.City = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetContacts(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Contacts = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetDefaultAddress(v bool) *DescribeDeliveryAddressResponseBodyAddresses {
	s.DefaultAddress = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetDetail(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Detail = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetMobile(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Mobile = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetPostalCode(v string) *DescribeDeliveryAddressResponseBodyAddresses {
	s.PostalCode = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetProvince(v *DescribeDeliveryAddressResponseBodyAddressesProvince) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Province = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) SetTown(v *DescribeDeliveryAddressResponseBodyAddressesTown) *DescribeDeliveryAddressResponseBodyAddresses {
	s.Town = v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddresses) Validate() error {
	if s.Area != nil {
		if err := s.Area.Validate(); err != nil {
			return err
		}
	}
	if s.City != nil {
		if err := s.City.Validate(); err != nil {
			return err
		}
	}
	if s.Province != nil {
		if err := s.Province.Validate(); err != nil {
			return err
		}
	}
	if s.Town != nil {
		if err := s.Town.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeliveryAddressResponseBodyAddressesArea struct {
	// The area ID.
	//
	// example:
	//
	// 33****
	AreaId *int64 `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The area name.
	//
	// example:
	//
	// 西湖区
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesArea) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesArea) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) GetAreaId() *int64 {
	return s.AreaId
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) GetAreaName() *string {
	return s.AreaName
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) SetAreaId(v int64) *DescribeDeliveryAddressResponseBodyAddressesArea {
	s.AreaId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) SetAreaName(v string) *DescribeDeliveryAddressResponseBodyAddressesArea {
	s.AreaName = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesArea) Validate() error {
	return dara.Validate(s)
}

type DescribeDeliveryAddressResponseBodyAddressesCity struct {
	// The city ID.
	//
	// example:
	//
	// 33****
	CityId *int64 `json:"CityId,omitempty" xml:"CityId,omitempty"`
	// The city name.
	//
	// example:
	//
	// 杭州市
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesCity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesCity) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) GetCityId() *int64 {
	return s.CityId
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) GetCityName() *string {
	return s.CityName
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) SetCityId(v int64) *DescribeDeliveryAddressResponseBodyAddressesCity {
	s.CityId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) SetCityName(v string) *DescribeDeliveryAddressResponseBodyAddressesCity {
	s.CityName = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesCity) Validate() error {
	return dara.Validate(s)
}

type DescribeDeliveryAddressResponseBodyAddressesProvince struct {
	// The province ID. For more information, see [China regions, countries and territories codes](https://help.aliyun.com/document_detail/167926.html).
	//
	// For example, **110000*	- represents Beijing and **120000*	- represents Tianjin.
	//
	// example:
	//
	// 330000
	ProvinceId *int64 `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	// The province name.
	//
	// example:
	//
	// 浙江省
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesProvince) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesProvince) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) GetProvinceId() *int64 {
	return s.ProvinceId
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) SetProvinceId(v int64) *DescribeDeliveryAddressResponseBodyAddressesProvince {
	s.ProvinceId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) SetProvinceName(v string) *DescribeDeliveryAddressResponseBodyAddressesProvince {
	s.ProvinceName = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesProvince) Validate() error {
	return dara.Validate(s)
}

type DescribeDeliveryAddressResponseBodyAddressesTown struct {
	// The town ID.
	//
	// example:
	//
	// 3001****
	TownId *int64 `json:"TownId,omitempty" xml:"TownId,omitempty"`
	// The town name.
	//
	// example:
	//
	// 三墩镇
	TownName *string `json:"TownName,omitempty" xml:"TownName,omitempty"`
}

func (s DescribeDeliveryAddressResponseBodyAddressesTown) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponseBodyAddressesTown) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) GetTownId() *int64 {
	return s.TownId
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) GetTownName() *string {
	return s.TownName
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) SetTownId(v int64) *DescribeDeliveryAddressResponseBodyAddressesTown {
	s.TownId = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) SetTownName(v string) *DescribeDeliveryAddressResponseBodyAddressesTown {
	s.TownName = &v
	return s
}

func (s *DescribeDeliveryAddressResponseBodyAddressesTown) Validate() error {
	return dara.Validate(s)
}
