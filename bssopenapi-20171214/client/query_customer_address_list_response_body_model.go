// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerAddressListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCustomerAddressListResponseBody
	GetCode() *string
	SetData(v *QueryCustomerAddressListResponseBodyData) *QueryCustomerAddressListResponseBody
	GetData() *QueryCustomerAddressListResponseBodyData
	SetMessage(v string) *QueryCustomerAddressListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomerAddressListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCustomerAddressListResponseBody
	GetSuccess() *bool
}

type QueryCustomerAddressListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCustomerAddressListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BBEF51A3-E933-4F40-A534-C673CBDB9C80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomerAddressListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerAddressListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCustomerAddressListResponseBody) GetData() *QueryCustomerAddressListResponseBodyData {
	return s.Data
}

func (s *QueryCustomerAddressListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomerAddressListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomerAddressListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomerAddressListResponseBody) SetCode(v string) *QueryCustomerAddressListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetData(v *QueryCustomerAddressListResponseBodyData) *QueryCustomerAddressListResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetMessage(v string) *QueryCustomerAddressListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetRequestId(v string) *QueryCustomerAddressListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetSuccess(v bool) *QueryCustomerAddressListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCustomerAddressListResponseBodyData struct {
	// The details of addresses to which invoices are mailed.
	CustomerInvoiceAddressList *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList `json:"CustomerInvoiceAddressList,omitempty" xml:"CustomerInvoiceAddressList,omitempty" type:"Struct"`
}

func (s QueryCustomerAddressListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerAddressListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBodyData) GetCustomerInvoiceAddressList() *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList {
	return s.CustomerInvoiceAddressList
}

func (s *QueryCustomerAddressListResponseBodyData) SetCustomerInvoiceAddressList(v *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) *QueryCustomerAddressListResponseBodyData {
	s.CustomerInvoiceAddressList = v
	return s
}

func (s *QueryCustomerAddressListResponseBodyData) Validate() error {
	if s.CustomerInvoiceAddressList != nil {
		if err := s.CustomerInvoiceAddressList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList struct {
	CustomerInvoiceAddress []*QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress `json:"CustomerInvoiceAddress,omitempty" xml:"CustomerInvoiceAddress,omitempty" type:"Repeated"`
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) GetCustomerInvoiceAddress() []*QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	return s.CustomerInvoiceAddress
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) SetCustomerInvoiceAddress(v []*QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList {
	s.CustomerInvoiceAddress = v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) Validate() error {
	if s.CustomerInvoiceAddress != nil {
		for _, item := range s.CustomerInvoiceAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress struct {
	// The addressee.
	//
	// example:
	//
	// test
	Addressee *string `json:"Addressee,omitempty" xml:"Addressee,omitempty"`
	// The business type.
	//
	// example:
	//
	// test
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The city to which the invoice is mailed.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The name of the district to which the invoice is mailed.
	//
	// example:
	//
	// Test District
	County *string `json:"County,omitempty" xml:"County,omitempty"`
	// The detailed address to which the invoice is mailed. This parameter is returned after fields are concatenated.
	//
	// example:
	//
	// Test Address
	DeliveryAddress *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	// The ID.
	//
	// example:
	//
	// 311601051
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number of the addressee.
	//
	// example:
	//
	// 138xxxxxxxx
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The postcode.
	//
	// example:
	//
	// 000000
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// The province to which the invoice is mailed.
	//
	// example:
	//
	// Zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The name of the street to which the invoice is mailed.
	//
	// example:
	//
	// Test Street
	Street *string `json:"Street,omitempty" xml:"Street,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 4382956342857
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// testNick
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetAddressee() *string {
	return s.Addressee
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetBizType() *string {
	return s.BizType
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetCity() *string {
	return s.City
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetCounty() *string {
	return s.County
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetDeliveryAddress() *string {
	return s.DeliveryAddress
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetId() *int64 {
	return s.Id
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetPhone() *string {
	return s.Phone
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetProvince() *string {
	return s.Province
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetStreet() *string {
	return s.Street
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GetUserNick() *string {
	return s.UserNick
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetAddressee(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Addressee = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetBizType(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.BizType = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetCity(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.City = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetCounty(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.County = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetDeliveryAddress(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.DeliveryAddress = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetId(v int64) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Id = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetPhone(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Phone = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetPostalCode(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.PostalCode = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetProvince(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Province = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetStreet(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Street = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetUserId(v int64) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.UserId = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetUserNick(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.UserNick = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) Validate() error {
	return dara.Validate(s)
}
