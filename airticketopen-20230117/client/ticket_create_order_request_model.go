// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketCreateOrderRequest
	GetAccountNo() *int64
	SetContact(v *TicketCreateOrderRequestContact) *TicketCreateOrderRequest
	GetContact() *TicketCreateOrderRequestContact
	SetDistributorOrderId(v string) *TicketCreateOrderRequest
	GetDistributorOrderId() *string
	SetOrderProduct(v *TicketCreateOrderRequestOrderProduct) *TicketCreateOrderRequest
	GetOrderProduct() *TicketCreateOrderRequestOrderProduct
	SetQuantity(v int32) *TicketCreateOrderRequest
	GetQuantity() *int32
	SetTotalDistributionPrice(v *TicketCreateOrderRequestTotalDistributionPrice) *TicketCreateOrderRequest
	GetTotalDistributionPrice() *TicketCreateOrderRequestTotalDistributionPrice
	SetTravelers(v []*TicketCreateOrderRequestTravelers) *TicketCreateOrderRequest
	GetTravelers() []*TicketCreateOrderRequestTravelers
}

type TicketCreateOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	Contact *TicketCreateOrderRequestContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	DistributorOrderId *string `json:"DistributorOrderId,omitempty" xml:"DistributorOrderId,omitempty"`
	// This parameter is required.
	OrderProduct *TicketCreateOrderRequestOrderProduct `json:"OrderProduct,omitempty" xml:"OrderProduct,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// This parameter is required.
	TotalDistributionPrice *TicketCreateOrderRequestTotalDistributionPrice `json:"TotalDistributionPrice,omitempty" xml:"TotalDistributionPrice,omitempty" type:"Struct"`
	Travelers              []*TicketCreateOrderRequestTravelers            `json:"Travelers,omitempty" xml:"Travelers,omitempty" type:"Repeated"`
}

func (s TicketCreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderRequest) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketCreateOrderRequest) GetContact() *TicketCreateOrderRequestContact {
	return s.Contact
}

func (s *TicketCreateOrderRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketCreateOrderRequest) GetOrderProduct() *TicketCreateOrderRequestOrderProduct {
	return s.OrderProduct
}

func (s *TicketCreateOrderRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *TicketCreateOrderRequest) GetTotalDistributionPrice() *TicketCreateOrderRequestTotalDistributionPrice {
	return s.TotalDistributionPrice
}

func (s *TicketCreateOrderRequest) GetTravelers() []*TicketCreateOrderRequestTravelers {
	return s.Travelers
}

func (s *TicketCreateOrderRequest) SetAccountNo(v int64) *TicketCreateOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketCreateOrderRequest) SetContact(v *TicketCreateOrderRequestContact) *TicketCreateOrderRequest {
	s.Contact = v
	return s
}

func (s *TicketCreateOrderRequest) SetDistributorOrderId(v string) *TicketCreateOrderRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketCreateOrderRequest) SetOrderProduct(v *TicketCreateOrderRequestOrderProduct) *TicketCreateOrderRequest {
	s.OrderProduct = v
	return s
}

func (s *TicketCreateOrderRequest) SetQuantity(v int32) *TicketCreateOrderRequest {
	s.Quantity = &v
	return s
}

func (s *TicketCreateOrderRequest) SetTotalDistributionPrice(v *TicketCreateOrderRequestTotalDistributionPrice) *TicketCreateOrderRequest {
	s.TotalDistributionPrice = v
	return s
}

func (s *TicketCreateOrderRequest) SetTravelers(v []*TicketCreateOrderRequestTravelers) *TicketCreateOrderRequest {
	s.Travelers = v
	return s
}

func (s *TicketCreateOrderRequest) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	if s.OrderProduct != nil {
		if err := s.OrderProduct.Validate(); err != nil {
			return err
		}
	}
	if s.TotalDistributionPrice != nil {
		if err := s.TotalDistributionPrice.Validate(); err != nil {
			return err
		}
	}
	if s.Travelers != nil {
		for _, item := range s.Travelers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketCreateOrderRequestContact struct {
	// example:
	//
	// 11010519900301001X
	CertificateNo *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	// example:
	//
	// 1
	CertificateType *int32 `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// 86
	DialingCode *string `json:"DialingCode,omitempty" xml:"DialingCode,omitempty"`
	// example:
	//
	// test@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// ZHANG
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// SAN
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// 13800000000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TicketCreateOrderRequestContact) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderRequestContact) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderRequestContact) GetCertificateNo() *string {
	return s.CertificateNo
}

func (s *TicketCreateOrderRequestContact) GetCertificateType() *int32 {
	return s.CertificateType
}

func (s *TicketCreateOrderRequestContact) GetDialingCode() *string {
	return s.DialingCode
}

func (s *TicketCreateOrderRequestContact) GetEmail() *string {
	return s.Email
}

func (s *TicketCreateOrderRequestContact) GetFirstName() *string {
	return s.FirstName
}

func (s *TicketCreateOrderRequestContact) GetLastName() *string {
	return s.LastName
}

func (s *TicketCreateOrderRequestContact) GetMobile() *string {
	return s.Mobile
}

func (s *TicketCreateOrderRequestContact) GetName() *string {
	return s.Name
}

func (s *TicketCreateOrderRequestContact) SetCertificateNo(v string) *TicketCreateOrderRequestContact {
	s.CertificateNo = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetCertificateType(v int32) *TicketCreateOrderRequestContact {
	s.CertificateType = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetDialingCode(v string) *TicketCreateOrderRequestContact {
	s.DialingCode = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetEmail(v string) *TicketCreateOrderRequestContact {
	s.Email = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetFirstName(v string) *TicketCreateOrderRequestContact {
	s.FirstName = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetLastName(v string) *TicketCreateOrderRequestContact {
	s.LastName = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetMobile(v string) *TicketCreateOrderRequestContact {
	s.Mobile = &v
	return s
}

func (s *TicketCreateOrderRequestContact) SetName(v string) *TicketCreateOrderRequestContact {
	s.Name = &v
	return s
}

func (s *TicketCreateOrderRequestContact) Validate() error {
	return dara.Validate(s)
}

type TicketCreateOrderRequestOrderProduct struct {
	// This parameter is required.
	DistributionPrice *TicketCreateOrderRequestOrderProductDistributionPrice `json:"DistributionPrice,omitempty" xml:"DistributionPrice,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-10-01
	TravelDate *string `json:"TravelDate,omitempty" xml:"TravelDate,omitempty"`
}

func (s TicketCreateOrderRequestOrderProduct) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderRequestOrderProduct) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderRequestOrderProduct) GetDistributionPrice() *TicketCreateOrderRequestOrderProductDistributionPrice {
	return s.DistributionPrice
}

func (s *TicketCreateOrderRequestOrderProduct) GetProductId() *string {
	return s.ProductId
}

func (s *TicketCreateOrderRequestOrderProduct) GetTravelDate() *string {
	return s.TravelDate
}

func (s *TicketCreateOrderRequestOrderProduct) SetDistributionPrice(v *TicketCreateOrderRequestOrderProductDistributionPrice) *TicketCreateOrderRequestOrderProduct {
	s.DistributionPrice = v
	return s
}

func (s *TicketCreateOrderRequestOrderProduct) SetProductId(v string) *TicketCreateOrderRequestOrderProduct {
	s.ProductId = &v
	return s
}

func (s *TicketCreateOrderRequestOrderProduct) SetTravelDate(v string) *TicketCreateOrderRequestOrderProduct {
	s.TravelDate = &v
	return s
}

func (s *TicketCreateOrderRequestOrderProduct) Validate() error {
	if s.DistributionPrice != nil {
		if err := s.DistributionPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketCreateOrderRequestOrderProductDistributionPrice struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketCreateOrderRequestOrderProductDistributionPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderRequestOrderProductDistributionPrice) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderRequestOrderProductDistributionPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketCreateOrderRequestOrderProductDistributionPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketCreateOrderRequestOrderProductDistributionPrice) SetAmount(v int64) *TicketCreateOrderRequestOrderProductDistributionPrice {
	s.Amount = &v
	return s
}

func (s *TicketCreateOrderRequestOrderProductDistributionPrice) SetCurrencyCode(v string) *TicketCreateOrderRequestOrderProductDistributionPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketCreateOrderRequestOrderProductDistributionPrice) Validate() error {
	return dara.Validate(s)
}

type TicketCreateOrderRequestTotalDistributionPrice struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketCreateOrderRequestTotalDistributionPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderRequestTotalDistributionPrice) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderRequestTotalDistributionPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketCreateOrderRequestTotalDistributionPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketCreateOrderRequestTotalDistributionPrice) SetAmount(v int64) *TicketCreateOrderRequestTotalDistributionPrice {
	s.Amount = &v
	return s
}

func (s *TicketCreateOrderRequestTotalDistributionPrice) SetCurrencyCode(v string) *TicketCreateOrderRequestTotalDistributionPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketCreateOrderRequestTotalDistributionPrice) Validate() error {
	return dara.Validate(s)
}

type TicketCreateOrderRequestTravelers struct {
	// example:
	//
	// 1990-01-01
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// example:
	//
	// 1234567890
	CertificateNo *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	// example:
	//
	// 1
	CertificateType *int32 `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// 86
	DialingCode *string `json:"DialingCode,omitempty" xml:"DialingCode,omitempty"`
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// Zhang
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// 1
	Gender *int32 `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// San
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// 13800000000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 张三
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// CN
	Nationality *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
}

func (s TicketCreateOrderRequestTravelers) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderRequestTravelers) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderRequestTravelers) GetBirthday() *string {
	return s.Birthday
}

func (s *TicketCreateOrderRequestTravelers) GetCertificateNo() *string {
	return s.CertificateNo
}

func (s *TicketCreateOrderRequestTravelers) GetCertificateType() *int32 {
	return s.CertificateType
}

func (s *TicketCreateOrderRequestTravelers) GetDialingCode() *string {
	return s.DialingCode
}

func (s *TicketCreateOrderRequestTravelers) GetEmail() *string {
	return s.Email
}

func (s *TicketCreateOrderRequestTravelers) GetFirstName() *string {
	return s.FirstName
}

func (s *TicketCreateOrderRequestTravelers) GetGender() *int32 {
	return s.Gender
}

func (s *TicketCreateOrderRequestTravelers) GetLastName() *string {
	return s.LastName
}

func (s *TicketCreateOrderRequestTravelers) GetMobile() *string {
	return s.Mobile
}

func (s *TicketCreateOrderRequestTravelers) GetName() *string {
	return s.Name
}

func (s *TicketCreateOrderRequestTravelers) GetNationality() *string {
	return s.Nationality
}

func (s *TicketCreateOrderRequestTravelers) SetBirthday(v string) *TicketCreateOrderRequestTravelers {
	s.Birthday = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetCertificateNo(v string) *TicketCreateOrderRequestTravelers {
	s.CertificateNo = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetCertificateType(v int32) *TicketCreateOrderRequestTravelers {
	s.CertificateType = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetDialingCode(v string) *TicketCreateOrderRequestTravelers {
	s.DialingCode = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetEmail(v string) *TicketCreateOrderRequestTravelers {
	s.Email = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetFirstName(v string) *TicketCreateOrderRequestTravelers {
	s.FirstName = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetGender(v int32) *TicketCreateOrderRequestTravelers {
	s.Gender = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetLastName(v string) *TicketCreateOrderRequestTravelers {
	s.LastName = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetMobile(v string) *TicketCreateOrderRequestTravelers {
	s.Mobile = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetName(v string) *TicketCreateOrderRequestTravelers {
	s.Name = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) SetNationality(v string) *TicketCreateOrderRequestTravelers {
	s.Nationality = &v
	return s
}

func (s *TicketCreateOrderRequestTravelers) Validate() error {
	return dara.Validate(s)
}
