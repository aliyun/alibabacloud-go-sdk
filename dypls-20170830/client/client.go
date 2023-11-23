// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyArInvoiceWithSourceRequest struct {
	Address              *ApplyArInvoiceWithSourceRequestAddress      `json:"Address,omitempty" xml:"Address,omitempty" type:"Struct"`
	Amount               *float64                                     `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCode              *string                                      `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	Applier              *string                                      `json:"Applier,omitempty" xml:"Applier,omitempty"`
	ApplyDate            *string                                      `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
	CurrencyCode         *string                                      `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	Customer             *ApplyArInvoiceWithSourceRequestCustomer     `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	EncryptProps         map[string]*string                           `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	ExcludingTaxAmount   *float64                                     `json:"ExcludingTaxAmount,omitempty" xml:"ExcludingTaxAmount,omitempty"`
	InputType            *string                                      `json:"InputType,omitempty" xml:"InputType,omitempty"`
	InvoiceType          *string                                      `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	IsMerged             *bool                                        `json:"IsMerged,omitempty" xml:"IsMerged,omitempty"`
	Language             *string                                      `json:"Language,omitempty" xml:"Language,omitempty"`
	MaterialType         *string                                      `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	Memo                 *string                                      `json:"Memo,omitempty" xml:"Memo,omitempty"`
	OuCode               *string                                      `json:"OuCode,omitempty" xml:"OuCode,omitempty"`
	PurchaserBankInfo    *string                                      `json:"PurchaserBankInfo,omitempty" xml:"PurchaserBankInfo,omitempty"`
	PurchaserContactInfo *string                                      `json:"PurchaserContactInfo,omitempty" xml:"PurchaserContactInfo,omitempty"`
	PurchaserName        *string                                      `json:"PurchaserName,omitempty" xml:"PurchaserName,omitempty"`
	PurchaserTaxNo       *string                                      `json:"PurchaserTaxNo,omitempty" xml:"PurchaserTaxNo,omitempty"`
	RequestNo            *string                                      `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	Sign                 *string                                      `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SiteId               *string                                      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SourceList           []*ApplyArInvoiceWithSourceRequestSourceList `json:"SourceList,omitempty" xml:"SourceList,omitempty" type:"Repeated"`
	TaxAmount            *float64                                     `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	Uuid                 *string                                      `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyArInvoiceWithSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceRequest) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceRequest) SetAddress(v *ApplyArInvoiceWithSourceRequestAddress) *ApplyArInvoiceWithSourceRequest {
	s.Address = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetAmount(v float64) *ApplyArInvoiceWithSourceRequest {
	s.Amount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetAppCode(v string) *ApplyArInvoiceWithSourceRequest {
	s.AppCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetApplier(v string) *ApplyArInvoiceWithSourceRequest {
	s.Applier = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetApplyDate(v string) *ApplyArInvoiceWithSourceRequest {
	s.ApplyDate = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetCurrencyCode(v string) *ApplyArInvoiceWithSourceRequest {
	s.CurrencyCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetCustomer(v *ApplyArInvoiceWithSourceRequestCustomer) *ApplyArInvoiceWithSourceRequest {
	s.Customer = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceRequest {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetExcludingTaxAmount(v float64) *ApplyArInvoiceWithSourceRequest {
	s.ExcludingTaxAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetInputType(v string) *ApplyArInvoiceWithSourceRequest {
	s.InputType = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetInvoiceType(v string) *ApplyArInvoiceWithSourceRequest {
	s.InvoiceType = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetIsMerged(v bool) *ApplyArInvoiceWithSourceRequest {
	s.IsMerged = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetLanguage(v string) *ApplyArInvoiceWithSourceRequest {
	s.Language = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetMaterialType(v string) *ApplyArInvoiceWithSourceRequest {
	s.MaterialType = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetMemo(v string) *ApplyArInvoiceWithSourceRequest {
	s.Memo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetOuCode(v string) *ApplyArInvoiceWithSourceRequest {
	s.OuCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetPurchaserBankInfo(v string) *ApplyArInvoiceWithSourceRequest {
	s.PurchaserBankInfo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetPurchaserContactInfo(v string) *ApplyArInvoiceWithSourceRequest {
	s.PurchaserContactInfo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetPurchaserName(v string) *ApplyArInvoiceWithSourceRequest {
	s.PurchaserName = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetPurchaserTaxNo(v string) *ApplyArInvoiceWithSourceRequest {
	s.PurchaserTaxNo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetRequestNo(v string) *ApplyArInvoiceWithSourceRequest {
	s.RequestNo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetSign(v string) *ApplyArInvoiceWithSourceRequest {
	s.Sign = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetSiteId(v string) *ApplyArInvoiceWithSourceRequest {
	s.SiteId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetSourceList(v []*ApplyArInvoiceWithSourceRequestSourceList) *ApplyArInvoiceWithSourceRequest {
	s.SourceList = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetTaxAmount(v float64) *ApplyArInvoiceWithSourceRequest {
	s.TaxAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequest) SetUuid(v string) *ApplyArInvoiceWithSourceRequest {
	s.Uuid = &v
	return s
}

type ApplyArInvoiceWithSourceRequestAddress struct {
	AddressId       *string                                         `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	AppCode         *string                                         `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	City            *string                                         `json:"City,omitempty" xml:"City,omitempty"`
	Country         *string                                         `json:"Country,omitempty" xml:"Country,omitempty"`
	Creator         *string                                         `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Customer        *ApplyArInvoiceWithSourceRequestAddressCustomer `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	DetailedAddress *string                                         `json:"DetailedAddress,omitempty" xml:"DetailedAddress,omitempty"`
	District        *string                                         `json:"District,omitempty" xml:"District,omitempty"`
	EncryptProps    map[string]*string                              `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	FixedNumber     *string                                         `json:"FixedNumber,omitempty" xml:"FixedNumber,omitempty"`
	FullAddress     *string                                         `json:"FullAddress,omitempty" xml:"FullAddress,omitempty"`
	IsDefault       *bool                                           `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IsEffective     *bool                                           `json:"IsEffective,omitempty" xml:"IsEffective,omitempty"`
	Language        *string                                         `json:"Language,omitempty" xml:"Language,omitempty"`
	MobileNumber    *string                                         `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	Province        *string                                         `json:"Province,omitempty" xml:"Province,omitempty"`
	RecipientName   *string                                         `json:"RecipientName,omitempty" xml:"RecipientName,omitempty"`
	RelatedId       *string                                         `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	RelatedSystem   *string                                         `json:"RelatedSystem,omitempty" xml:"RelatedSystem,omitempty"`
	Sign            *string                                         `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Street          *string                                         `json:"Street,omitempty" xml:"Street,omitempty"`
	Uuid            *string                                         `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	ZipCode         *string                                         `json:"ZipCode,omitempty" xml:"ZipCode,omitempty"`
}

func (s ApplyArInvoiceWithSourceRequestAddress) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceRequestAddress) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetAddressId(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.AddressId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetAppCode(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.AppCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetCity(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.City = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetCountry(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Country = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetCreator(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Creator = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetCustomer(v *ApplyArInvoiceWithSourceRequestAddressCustomer) *ApplyArInvoiceWithSourceRequestAddress {
	s.Customer = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetDetailedAddress(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.DetailedAddress = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetDistrict(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.District = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceRequestAddress {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetFixedNumber(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.FixedNumber = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetFullAddress(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.FullAddress = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetIsDefault(v bool) *ApplyArInvoiceWithSourceRequestAddress {
	s.IsDefault = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetIsEffective(v bool) *ApplyArInvoiceWithSourceRequestAddress {
	s.IsEffective = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetLanguage(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Language = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetMobileNumber(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.MobileNumber = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetProvince(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Province = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetRecipientName(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.RecipientName = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetRelatedId(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.RelatedId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetRelatedSystem(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.RelatedSystem = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetSign(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Sign = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetStreet(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Street = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetUuid(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.Uuid = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddress) SetZipCode(v string) *ApplyArInvoiceWithSourceRequestAddress {
	s.ZipCode = &v
	return s
}

type ApplyArInvoiceWithSourceRequestAddressCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyArInvoiceWithSourceRequestAddressCustomer) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceRequestAddressCustomer) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetAppCode(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.AppCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetCustomerId(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.CustomerId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetCustomerSite(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.CustomerSite = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetCustomerSystem(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetLanguage(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.Language = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetSign(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.Sign = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestAddressCustomer) SetUuid(v string) *ApplyArInvoiceWithSourceRequestAddressCustomer {
	s.Uuid = &v
	return s
}

type ApplyArInvoiceWithSourceRequestCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyArInvoiceWithSourceRequestCustomer) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceRequestCustomer) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetAppCode(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.AppCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetCustomerId(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.CustomerId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetCustomerSite(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.CustomerSite = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetCustomerSystem(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetLanguage(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.Language = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetSign(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.Sign = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestCustomer) SetUuid(v string) *ApplyArInvoiceWithSourceRequestCustomer {
	s.Uuid = &v
	return s
}

type ApplyArInvoiceWithSourceRequestSourceList struct {
	Amount                     *float64                                           `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCode                    *string                                            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	BillAmount                 *float64                                           `json:"BillAmount,omitempty" xml:"BillAmount,omitempty"`
	BillDomain                 *string                                            `json:"BillDomain,omitempty" xml:"BillDomain,omitempty"`
	BillNo                     *string                                            `json:"BillNo,omitempty" xml:"BillNo,omitempty"`
	BillType                   *string                                            `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BlueSourceId               *int64                                             `json:"BlueSourceId,omitempty" xml:"BlueSourceId,omitempty"`
	CanMerge                   *bool                                              `json:"CanMerge,omitempty" xml:"CanMerge,omitempty"`
	CargoName                  *string                                            `json:"CargoName,omitempty" xml:"CargoName,omitempty"`
	Category                   *string                                            `json:"Category,omitempty" xml:"Category,omitempty"`
	CompanyName                *string                                            `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	CurrencyCode               *string                                            `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	Customer                   *ApplyArInvoiceWithSourceRequestSourceListCustomer `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	DiscountAmount             *float64                                           `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	DiscountTaxAmount          *float64                                           `json:"DiscountTaxAmount,omitempty" xml:"DiscountTaxAmount,omitempty"`
	EncryptProps               map[string]*string                                 `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	ExcludingTaxAmount         *float64                                           `json:"ExcludingTaxAmount,omitempty" xml:"ExcludingTaxAmount,omitempty"`
	ExcludingTaxDiscountAmount *float64                                           `json:"ExcludingTaxDiscountAmount,omitempty" xml:"ExcludingTaxDiscountAmount,omitempty"`
	ExcludingTaxRedAmount      *float64                                           `json:"ExcludingTaxRedAmount,omitempty" xml:"ExcludingTaxRedAmount,omitempty"`
	ExcludingTaxRemainAmount   *float64                                           `json:"ExcludingTaxRemainAmount,omitempty" xml:"ExcludingTaxRemainAmount,omitempty"`
	GmtBill                    *string                                            `json:"GmtBill,omitempty" xml:"GmtBill,omitempty"`
	GmtBillEnd                 *string                                            `json:"GmtBillEnd,omitempty" xml:"GmtBillEnd,omitempty"`
	GmtBillStart               *string                                            `json:"GmtBillStart,omitempty" xml:"GmtBillStart,omitempty"`
	GmtBuild                   *string                                            `json:"GmtBuild,omitempty" xml:"GmtBuild,omitempty"`
	IsApply                    *string                                            `json:"IsApply,omitempty" xml:"IsApply,omitempty"`
	Language                   *string                                            `json:"Language,omitempty" xml:"Language,omitempty"`
	MajorBillNo                *string                                            `json:"MajorBillNo,omitempty" xml:"MajorBillNo,omitempty"`
	Model                      *string                                            `json:"Model,omitempty" xml:"Model,omitempty"`
	OuCode                     *string                                            `json:"OuCode,omitempty" xml:"OuCode,omitempty"`
	ParentCategory             *string                                            `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ProductDomain              *string                                            `json:"ProductDomain,omitempty" xml:"ProductDomain,omitempty"`
	ProductId                  *string                                            `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName                *string                                            `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Quantity                   *float64                                           `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	QuantityUnit               *string                                            `json:"QuantityUnit,omitempty" xml:"QuantityUnit,omitempty"`
	RedAmount                  *float64                                           `json:"RedAmount,omitempty" xml:"RedAmount,omitempty"`
	RelatedId                  *string                                            `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	RemainAmount               *float64                                           `json:"RemainAmount,omitempty" xml:"RemainAmount,omitempty"`
	RevenueType                *string                                            `json:"RevenueType,omitempty" xml:"RevenueType,omitempty"`
	ServiceName                *string                                            `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Sign                       *string                                            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SiteId                     *string                                            `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SourceId                   *int64                                             `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	TaxAmount                  *float64                                           `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	TaxRate                    *float64                                           `json:"TaxRate,omitempty" xml:"TaxRate,omitempty"`
	UnitPrice                  *float64                                           `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
	Uuid                       *string                                            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyArInvoiceWithSourceRequestSourceList) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceRequestSourceList) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Amount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetAppCode(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.AppCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetBillAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.BillAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetBillDomain(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.BillDomain = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetBillNo(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.BillNo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetBillType(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.BillType = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetBlueSourceId(v int64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.BlueSourceId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetCanMerge(v bool) *ApplyArInvoiceWithSourceRequestSourceList {
	s.CanMerge = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetCargoName(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.CargoName = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetCategory(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Category = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetCompanyName(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.CompanyName = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetCurrencyCode(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.CurrencyCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetCustomer(v *ApplyArInvoiceWithSourceRequestSourceListCustomer) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Customer = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetDiscountAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.DiscountAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetDiscountTaxAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.DiscountTaxAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetExcludingTaxAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ExcludingTaxAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetExcludingTaxDiscountAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ExcludingTaxDiscountAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetExcludingTaxRedAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ExcludingTaxRedAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetExcludingTaxRemainAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ExcludingTaxRemainAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetGmtBill(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.GmtBill = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetGmtBillEnd(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.GmtBillEnd = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetGmtBillStart(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.GmtBillStart = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetGmtBuild(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.GmtBuild = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetIsApply(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.IsApply = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetLanguage(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Language = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetMajorBillNo(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.MajorBillNo = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetModel(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Model = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetOuCode(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.OuCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetParentCategory(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ParentCategory = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetProductDomain(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ProductDomain = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetProductId(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ProductId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetProductName(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ProductName = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetQuantity(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Quantity = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetQuantityUnit(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.QuantityUnit = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetRedAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.RedAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetRelatedId(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.RelatedId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetRemainAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.RemainAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetRevenueType(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.RevenueType = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetServiceName(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.ServiceName = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetSign(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Sign = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetSiteId(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.SiteId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetSourceId(v int64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.SourceId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetTaxAmount(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.TaxAmount = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetTaxRate(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.TaxRate = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetUnitPrice(v float64) *ApplyArInvoiceWithSourceRequestSourceList {
	s.UnitPrice = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceList) SetUuid(v string) *ApplyArInvoiceWithSourceRequestSourceList {
	s.Uuid = &v
	return s
}

type ApplyArInvoiceWithSourceRequestSourceListCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApplyArInvoiceWithSourceRequestSourceListCustomer) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceRequestSourceListCustomer) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetAppCode(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.AppCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetCustomerId(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.CustomerId = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetCustomerSite(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.CustomerSite = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetCustomerSystem(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetLanguage(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.Language = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetSign(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.Sign = &v
	return s
}

func (s *ApplyArInvoiceWithSourceRequestSourceListCustomer) SetUuid(v string) *ApplyArInvoiceWithSourceRequestSourceListCustomer {
	s.Uuid = &v
	return s
}

type ApplyArInvoiceWithSourceResponseBody struct {
	ErrorCode    *string                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                            `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReturnValue  *ApplyArInvoiceWithSourceResponseBodyReturnValue `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Struct"`
}

func (s ApplyArInvoiceWithSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceResponseBody) SetErrorCode(v string) *ApplyArInvoiceWithSourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceResponseBody) SetErrorMessage(v string) *ApplyArInvoiceWithSourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApplyArInvoiceWithSourceResponseBody) SetIsSuccess(v bool) *ApplyArInvoiceWithSourceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ApplyArInvoiceWithSourceResponseBody) SetReturnValue(v *ApplyArInvoiceWithSourceResponseBodyReturnValue) *ApplyArInvoiceWithSourceResponseBody {
	s.ReturnValue = v
	return s
}

type ApplyArInvoiceWithSourceResponseBodyReturnValue struct {
	EncryptProps          map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	OuterSystemEncryptStr *string            `json:"OuterSystemEncryptStr,omitempty" xml:"OuterSystemEncryptStr,omitempty"`
	OuterSystemSignStr    *string            `json:"OuterSystemSignStr,omitempty" xml:"OuterSystemSignStr,omitempty"`
	Sign                  *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s ApplyArInvoiceWithSourceResponseBodyReturnValue) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceResponseBodyReturnValue) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceResponseBodyReturnValue) SetEncryptProps(v map[string]*string) *ApplyArInvoiceWithSourceResponseBodyReturnValue {
	s.EncryptProps = v
	return s
}

func (s *ApplyArInvoiceWithSourceResponseBodyReturnValue) SetOuterSystemEncryptStr(v string) *ApplyArInvoiceWithSourceResponseBodyReturnValue {
	s.OuterSystemEncryptStr = &v
	return s
}

func (s *ApplyArInvoiceWithSourceResponseBodyReturnValue) SetOuterSystemSignStr(v string) *ApplyArInvoiceWithSourceResponseBodyReturnValue {
	s.OuterSystemSignStr = &v
	return s
}

func (s *ApplyArInvoiceWithSourceResponseBodyReturnValue) SetSign(v string) *ApplyArInvoiceWithSourceResponseBodyReturnValue {
	s.Sign = &v
	return s
}

type ApplyArInvoiceWithSourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyArInvoiceWithSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyArInvoiceWithSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyArInvoiceWithSourceResponse) GoString() string {
	return s.String()
}

func (s *ApplyArInvoiceWithSourceResponse) SetHeaders(v map[string]*string) *ApplyArInvoiceWithSourceResponse {
	s.Headers = v
	return s
}

func (s *ApplyArInvoiceWithSourceResponse) SetStatusCode(v int32) *ApplyArInvoiceWithSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyArInvoiceWithSourceResponse) SetBody(v *ApplyArInvoiceWithSourceResponseBody) *ApplyArInvoiceWithSourceResponse {
	s.Body = v
	return s
}

type ApplyBlackInfoExportRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	BlackType            *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ApplyBlackInfoExportRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyBlackInfoExportRequest) GoString() string {
	return s.String()
}

func (s *ApplyBlackInfoExportRequest) SetBillId(v string) *ApplyBlackInfoExportRequest {
	s.BillId = &v
	return s
}

func (s *ApplyBlackInfoExportRequest) SetBlackType(v string) *ApplyBlackInfoExportRequest {
	s.BlackType = &v
	return s
}

func (s *ApplyBlackInfoExportRequest) SetOwnerId(v int64) *ApplyBlackInfoExportRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyBlackInfoExportRequest) SetProdCode(v string) *ApplyBlackInfoExportRequest {
	s.ProdCode = &v
	return s
}

func (s *ApplyBlackInfoExportRequest) SetResourceOwnerAccount(v string) *ApplyBlackInfoExportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApplyBlackInfoExportRequest) SetResourceOwnerId(v int64) *ApplyBlackInfoExportRequest {
	s.ResourceOwnerId = &v
	return s
}

type ApplyBlackInfoExportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ApplyBlackInfoExportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyBlackInfoExportResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyBlackInfoExportResponseBody) SetRequestId(v string) *ApplyBlackInfoExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyBlackInfoExportResponseBody) SetCode(v string) *ApplyBlackInfoExportResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyBlackInfoExportResponseBody) SetData(v string) *ApplyBlackInfoExportResponseBody {
	s.Data = &v
	return s
}

type ApplyBlackInfoExportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyBlackInfoExportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyBlackInfoExportResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyBlackInfoExportResponse) GoString() string {
	return s.String()
}

func (s *ApplyBlackInfoExportResponse) SetHeaders(v map[string]*string) *ApplyBlackInfoExportResponse {
	s.Headers = v
	return s
}

func (s *ApplyBlackInfoExportResponse) SetStatusCode(v int32) *ApplyBlackInfoExportResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyBlackInfoExportResponse) SetBody(v *ApplyBlackInfoExportResponseBody) *ApplyBlackInfoExportResponse {
	s.Body = v
	return s
}

type ApplyCallRecordExportRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	CallDate             *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s ApplyCallRecordExportRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCallRecordExportRequest) GoString() string {
	return s.String()
}

func (s *ApplyCallRecordExportRequest) SetBillId(v string) *ApplyCallRecordExportRequest {
	s.BillId = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetCallDate(v string) *ApplyCallRecordExportRequest {
	s.CallDate = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetCallId(v string) *ApplyCallRecordExportRequest {
	s.CallId = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetOwnerId(v int64) *ApplyCallRecordExportRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetPhoneNoA(v string) *ApplyCallRecordExportRequest {
	s.PhoneNoA = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetPhoneNoB(v string) *ApplyCallRecordExportRequest {
	s.PhoneNoB = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetProdCode(v string) *ApplyCallRecordExportRequest {
	s.ProdCode = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetResType(v int32) *ApplyCallRecordExportRequest {
	s.ResType = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetResourceOwnerAccount(v string) *ApplyCallRecordExportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetResourceOwnerId(v int64) *ApplyCallRecordExportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApplyCallRecordExportRequest) SetSecretNo(v string) *ApplyCallRecordExportRequest {
	s.SecretNo = &v
	return s
}

type ApplyCallRecordExportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ApplyCallRecordExportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyCallRecordExportResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCallRecordExportResponseBody) SetRequestId(v string) *ApplyCallRecordExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCallRecordExportResponseBody) SetCode(v string) *ApplyCallRecordExportResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyCallRecordExportResponseBody) SetData(v string) *ApplyCallRecordExportResponseBody {
	s.Data = &v
	return s
}

type ApplyCallRecordExportResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyCallRecordExportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyCallRecordExportResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyCallRecordExportResponse) GoString() string {
	return s.String()
}

func (s *ApplyCallRecordExportResponse) SetHeaders(v map[string]*string) *ApplyCallRecordExportResponse {
	s.Headers = v
	return s
}

func (s *ApplyCallRecordExportResponse) SetStatusCode(v int32) *ApplyCallRecordExportResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCallRecordExportResponse) SetBody(v *ApplyCallRecordExportResponseBody) *ApplyCallRecordExportResponse {
	s.Body = v
	return s
}

type ApplyGroupNumberExportRequest struct {
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ApplyGroupNumberExportRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyGroupNumberExportRequest) GoString() string {
	return s.String()
}

func (s *ApplyGroupNumberExportRequest) SetGroupId(v string) *ApplyGroupNumberExportRequest {
	s.GroupId = &v
	return s
}

func (s *ApplyGroupNumberExportRequest) SetOwnerId(v int64) *ApplyGroupNumberExportRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyGroupNumberExportRequest) SetPoolKey(v string) *ApplyGroupNumberExportRequest {
	s.PoolKey = &v
	return s
}

func (s *ApplyGroupNumberExportRequest) SetProdCode(v string) *ApplyGroupNumberExportRequest {
	s.ProdCode = &v
	return s
}

func (s *ApplyGroupNumberExportRequest) SetResourceOwnerAccount(v string) *ApplyGroupNumberExportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApplyGroupNumberExportRequest) SetResourceOwnerId(v int64) *ApplyGroupNumberExportRequest {
	s.ResourceOwnerId = &v
	return s
}

type ApplyGroupNumberExportResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ApplyGroupNumberExportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyGroupNumberExportResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyGroupNumberExportResponseBody) SetRequestId(v string) *ApplyGroupNumberExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyGroupNumberExportResponseBody) SetCode(v string) *ApplyGroupNumberExportResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyGroupNumberExportResponseBody) SetData(v string) *ApplyGroupNumberExportResponseBody {
	s.Data = &v
	return s
}

type ApplyGroupNumberExportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyGroupNumberExportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyGroupNumberExportResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyGroupNumberExportResponse) GoString() string {
	return s.String()
}

func (s *ApplyGroupNumberExportResponse) SetHeaders(v map[string]*string) *ApplyGroupNumberExportResponse {
	s.Headers = v
	return s
}

func (s *ApplyGroupNumberExportResponse) SetStatusCode(v int32) *ApplyGroupNumberExportResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyGroupNumberExportResponse) SetBody(v *ApplyGroupNumberExportResponseBody) *ApplyGroupNumberExportResponse {
	s.Body = v
	return s
}

type ApplyRingToneRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayType             *string `json:"PlayType,omitempty" xml:"PlayType,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ApplyRingToneRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyRingToneRequest) GoString() string {
	return s.String()
}

func (s *ApplyRingToneRequest) SetBillId(v string) *ApplyRingToneRequest {
	s.BillId = &v
	return s
}

func (s *ApplyRingToneRequest) SetId(v string) *ApplyRingToneRequest {
	s.Id = &v
	return s
}

func (s *ApplyRingToneRequest) SetOwnerId(v int64) *ApplyRingToneRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyRingToneRequest) SetPlayType(v string) *ApplyRingToneRequest {
	s.PlayType = &v
	return s
}

func (s *ApplyRingToneRequest) SetProdCode(v string) *ApplyRingToneRequest {
	s.ProdCode = &v
	return s
}

func (s *ApplyRingToneRequest) SetResourceOwnerAccount(v string) *ApplyRingToneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApplyRingToneRequest) SetResourceOwnerId(v int64) *ApplyRingToneRequest {
	s.ResourceOwnerId = &v
	return s
}

type ApplyRingToneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ApplyRingToneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyRingToneResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyRingToneResponseBody) SetRequestId(v string) *ApplyRingToneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyRingToneResponseBody) SetCode(v string) *ApplyRingToneResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyRingToneResponseBody) SetData(v string) *ApplyRingToneResponseBody {
	s.Data = &v
	return s
}

type ApplyRingToneResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyRingToneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyRingToneResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyRingToneResponse) GoString() string {
	return s.String()
}

func (s *ApplyRingToneResponse) SetHeaders(v map[string]*string) *ApplyRingToneResponse {
	s.Headers = v
	return s
}

func (s *ApplyRingToneResponse) SetStatusCode(v int32) *ApplyRingToneResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyRingToneResponse) SetBody(v *ApplyRingToneResponseBody) *ApplyRingToneResponse {
	s.Body = v
	return s
}

type BatchOccupySecretResRequest struct {
	BatchOccupyList      []*BatchOccupySecretResRequestBatchOccupyList `json:"BatchOccupyList,omitempty" xml:"BatchOccupyList,omitempty" type:"Repeated"`
	OwnerId              *int64                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string                                       `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string                                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BatchOccupySecretResRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchOccupySecretResRequest) GoString() string {
	return s.String()
}

func (s *BatchOccupySecretResRequest) SetBatchOccupyList(v []*BatchOccupySecretResRequestBatchOccupyList) *BatchOccupySecretResRequest {
	s.BatchOccupyList = v
	return s
}

func (s *BatchOccupySecretResRequest) SetOwnerId(v int64) *BatchOccupySecretResRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchOccupySecretResRequest) SetProdCode(v string) *BatchOccupySecretResRequest {
	s.ProdCode = &v
	return s
}

func (s *BatchOccupySecretResRequest) SetResourceOwnerAccount(v string) *BatchOccupySecretResRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchOccupySecretResRequest) SetResourceOwnerId(v int64) *BatchOccupySecretResRequest {
	s.ResourceOwnerId = &v
	return s
}

type BatchOccupySecretResRequestBatchOccupyList struct {
	Count         *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	OrderDetailId *int64  `json:"OrderDetailId,omitempty" xml:"OrderDetailId,omitempty"`
	OrderId       *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PartnerKey    *string `json:"PartnerKey,omitempty" xml:"PartnerKey,omitempty"`
	ResType       *int64  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	SecretNoType  *int64  `json:"SecretNoType,omitempty" xml:"SecretNoType,omitempty"`
}

func (s BatchOccupySecretResRequestBatchOccupyList) String() string {
	return tea.Prettify(s)
}

func (s BatchOccupySecretResRequestBatchOccupyList) GoString() string {
	return s.String()
}

func (s *BatchOccupySecretResRequestBatchOccupyList) SetCount(v int32) *BatchOccupySecretResRequestBatchOccupyList {
	s.Count = &v
	return s
}

func (s *BatchOccupySecretResRequestBatchOccupyList) SetOrderDetailId(v int64) *BatchOccupySecretResRequestBatchOccupyList {
	s.OrderDetailId = &v
	return s
}

func (s *BatchOccupySecretResRequestBatchOccupyList) SetOrderId(v int64) *BatchOccupySecretResRequestBatchOccupyList {
	s.OrderId = &v
	return s
}

func (s *BatchOccupySecretResRequestBatchOccupyList) SetPartnerKey(v string) *BatchOccupySecretResRequestBatchOccupyList {
	s.PartnerKey = &v
	return s
}

func (s *BatchOccupySecretResRequestBatchOccupyList) SetResType(v int64) *BatchOccupySecretResRequestBatchOccupyList {
	s.ResType = &v
	return s
}

func (s *BatchOccupySecretResRequestBatchOccupyList) SetSecretNoType(v int64) *BatchOccupySecretResRequestBatchOccupyList {
	s.SecretNoType = &v
	return s
}

type BatchOccupySecretResShrinkRequest struct {
	BatchOccupyListShrink *string `json:"BatchOccupyList,omitempty" xml:"BatchOccupyList,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode              *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BatchOccupySecretResShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchOccupySecretResShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchOccupySecretResShrinkRequest) SetBatchOccupyListShrink(v string) *BatchOccupySecretResShrinkRequest {
	s.BatchOccupyListShrink = &v
	return s
}

func (s *BatchOccupySecretResShrinkRequest) SetOwnerId(v int64) *BatchOccupySecretResShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchOccupySecretResShrinkRequest) SetProdCode(v string) *BatchOccupySecretResShrinkRequest {
	s.ProdCode = &v
	return s
}

func (s *BatchOccupySecretResShrinkRequest) SetResourceOwnerAccount(v string) *BatchOccupySecretResShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchOccupySecretResShrinkRequest) SetResourceOwnerId(v int64) *BatchOccupySecretResShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

type BatchOccupySecretResResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchOccupySecretResResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchOccupySecretResResponseBody) GoString() string {
	return s.String()
}

func (s *BatchOccupySecretResResponseBody) SetCode(v string) *BatchOccupySecretResResponseBody {
	s.Code = &v
	return s
}

func (s *BatchOccupySecretResResponseBody) SetData(v bool) *BatchOccupySecretResResponseBody {
	s.Data = &v
	return s
}

func (s *BatchOccupySecretResResponseBody) SetMessage(v string) *BatchOccupySecretResResponseBody {
	s.Message = &v
	return s
}

func (s *BatchOccupySecretResResponseBody) SetRequestId(v string) *BatchOccupySecretResResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchOccupySecretResResponseBody) SetSuccess(v bool) *BatchOccupySecretResResponseBody {
	s.Success = &v
	return s
}

type BatchOccupySecretResResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchOccupySecretResResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchOccupySecretResResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchOccupySecretResResponse) GoString() string {
	return s.String()
}

func (s *BatchOccupySecretResResponse) SetHeaders(v map[string]*string) *BatchOccupySecretResResponse {
	s.Headers = v
	return s
}

func (s *BatchOccupySecretResResponse) SetStatusCode(v int32) *BatchOccupySecretResResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchOccupySecretResResponse) SetBody(v *BatchOccupySecretResResponseBody) *BatchOccupySecretResResponse {
	s.Body = v
	return s
}

type BindResourceRequest struct {
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	AsrStatus            *bool   `json:"AsrStatus,omitempty" xml:"AsrStatus,omitempty"`
	AxnExtensionB        *string `json:"AxnExtensionB,omitempty" xml:"AxnExtensionB,omitempty"`
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	ExpTime              *string `json:"ExpTime,omitempty" xml:"ExpTime,omitempty"`
	IsRecord             *bool   `json:"IsRecord,omitempty" xml:"IsRecord,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s BindResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindResourceRequest) GoString() string {
	return s.String()
}

func (s *BindResourceRequest) SetAsrModelId(v string) *BindResourceRequest {
	s.AsrModelId = &v
	return s
}

func (s *BindResourceRequest) SetAsrStatus(v bool) *BindResourceRequest {
	s.AsrStatus = &v
	return s
}

func (s *BindResourceRequest) SetAxnExtensionB(v string) *BindResourceRequest {
	s.AxnExtensionB = &v
	return s
}

func (s *BindResourceRequest) SetBillId(v string) *BindResourceRequest {
	s.BillId = &v
	return s
}

func (s *BindResourceRequest) SetExpTime(v string) *BindResourceRequest {
	s.ExpTime = &v
	return s
}

func (s *BindResourceRequest) SetIsRecord(v bool) *BindResourceRequest {
	s.IsRecord = &v
	return s
}

func (s *BindResourceRequest) SetOwnerId(v int64) *BindResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *BindResourceRequest) SetPhoneNoA(v string) *BindResourceRequest {
	s.PhoneNoA = &v
	return s
}

func (s *BindResourceRequest) SetPhoneNoB(v string) *BindResourceRequest {
	s.PhoneNoB = &v
	return s
}

func (s *BindResourceRequest) SetProdCode(v string) *BindResourceRequest {
	s.ProdCode = &v
	return s
}

func (s *BindResourceRequest) SetResType(v int32) *BindResourceRequest {
	s.ResType = &v
	return s
}

func (s *BindResourceRequest) SetResourceOwnerAccount(v string) *BindResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindResourceRequest) SetResourceOwnerId(v int64) *BindResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindResourceRequest) SetSecretNo(v string) *BindResourceRequest {
	s.SecretNo = &v
	return s
}

type BindResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BindResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindResourceResponseBody) GoString() string {
	return s.String()
}

func (s *BindResourceResponseBody) SetRequestId(v string) *BindResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindResourceResponseBody) SetCode(v string) *BindResourceResponseBody {
	s.Code = &v
	return s
}

func (s *BindResourceResponseBody) SetData(v string) *BindResourceResponseBody {
	s.Data = &v
	return s
}

type BindResourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindResourceResponse) GoString() string {
	return s.String()
}

func (s *BindResourceResponse) SetHeaders(v map[string]*string) *BindResourceResponse {
	s.Headers = v
	return s
}

func (s *BindResourceResponse) SetStatusCode(v int32) *BindResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindResourceResponse) SetBody(v *BindResourceResponseBody) *BindResourceResponse {
	s.Body = v
	return s
}

type BlackOperateRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	BlackMap             *string `json:"BlackMap,omitempty" xml:"BlackMap,omitempty"`
	BlackType            *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	OperateType          *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BlackOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s BlackOperateRequest) GoString() string {
	return s.String()
}

func (s *BlackOperateRequest) SetBillId(v string) *BlackOperateRequest {
	s.BillId = &v
	return s
}

func (s *BlackOperateRequest) SetBlackMap(v string) *BlackOperateRequest {
	s.BlackMap = &v
	return s
}

func (s *BlackOperateRequest) SetBlackType(v string) *BlackOperateRequest {
	s.BlackType = &v
	return s
}

func (s *BlackOperateRequest) SetOperateType(v string) *BlackOperateRequest {
	s.OperateType = &v
	return s
}

func (s *BlackOperateRequest) SetOwnerId(v int64) *BlackOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *BlackOperateRequest) SetProdCode(v string) *BlackOperateRequest {
	s.ProdCode = &v
	return s
}

func (s *BlackOperateRequest) SetResourceOwnerAccount(v string) *BlackOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BlackOperateRequest) SetResourceOwnerId(v int64) *BlackOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

type BlackOperateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s BlackOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BlackOperateResponseBody) GoString() string {
	return s.String()
}

func (s *BlackOperateResponseBody) SetRequestId(v string) *BlackOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlackOperateResponseBody) SetCode(v string) *BlackOperateResponseBody {
	s.Code = &v
	return s
}

func (s *BlackOperateResponseBody) SetData(v string) *BlackOperateResponseBody {
	s.Data = &v
	return s
}

func (s *BlackOperateResponseBody) SetMessage(v string) *BlackOperateResponseBody {
	s.Message = &v
	return s
}

type BlackOperateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BlackOperateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BlackOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s BlackOperateResponse) GoString() string {
	return s.String()
}

func (s *BlackOperateResponse) SetHeaders(v map[string]*string) *BlackOperateResponse {
	s.Headers = v
	return s
}

func (s *BlackOperateResponse) SetStatusCode(v int32) *BlackOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *BlackOperateResponse) SetBody(v *BlackOperateResponseBody) *BlackOperateResponse {
	s.Body = v
	return s
}

type CreateCertifyInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateCertifyInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertifyInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateCertifyInfoRequest) SetOwnerId(v int64) *CreateCertifyInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCertifyInfoRequest) SetPhoneNo(v string) *CreateCertifyInfoRequest {
	s.PhoneNo = &v
	return s
}

func (s *CreateCertifyInfoRequest) SetProdCode(v string) *CreateCertifyInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateCertifyInfoRequest) SetResourceOwnerAccount(v string) *CreateCertifyInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCertifyInfoRequest) SetResourceOwnerId(v int64) *CreateCertifyInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateCertifyInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateCertifyInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertifyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertifyInfoResponseBody) SetRequestId(v string) *CreateCertifyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertifyInfoResponseBody) SetCode(v string) *CreateCertifyInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCertifyInfoResponseBody) SetData(v string) *CreateCertifyInfoResponseBody {
	s.Data = &v
	return s
}

type CreateCertifyInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCertifyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertifyInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertifyInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateCertifyInfoResponse) SetHeaders(v map[string]*string) *CreateCertifyInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateCertifyInfoResponse) SetStatusCode(v int32) *CreateCertifyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCertifyInfoResponse) SetBody(v *CreateCertifyInfoResponseBody) *CreateCertifyInfoResponse {
	s.Body = v
	return s
}

type CreateContactsRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContactsRequest) GoString() string {
	return s.String()
}

func (s *CreateContactsRequest) SetBillId(v string) *CreateContactsRequest {
	s.BillId = &v
	return s
}

func (s *CreateContactsRequest) SetName(v string) *CreateContactsRequest {
	s.Name = &v
	return s
}

func (s *CreateContactsRequest) SetOwnerId(v int64) *CreateContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateContactsRequest) SetPhoneNumber(v string) *CreateContactsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateContactsRequest) SetProdCode(v string) *CreateContactsRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateContactsRequest) SetResourceOwnerAccount(v string) *CreateContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateContactsRequest) SetResourceOwnerId(v int64) *CreateContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateContactsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContactsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactsResponseBody) SetRequestId(v string) *CreateContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContactsResponseBody) SetCode(v string) *CreateContactsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateContactsResponseBody) SetData(v string) *CreateContactsResponseBody {
	s.Data = &v
	return s
}

type CreateContactsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContactsResponse) GoString() string {
	return s.String()
}

func (s *CreateContactsResponse) SetHeaders(v map[string]*string) *CreateContactsResponse {
	s.Headers = v
	return s
}

func (s *CreateContactsResponse) SetStatusCode(v int32) *CreateContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContactsResponse) SetBody(v *CreateContactsResponseBody) *CreateContactsResponse {
	s.Body = v
	return s
}

type CreateGroupDetailRequest struct {
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NumberList           *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupDetailRequest) SetGroupId(v string) *CreateGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupDetailRequest) SetNumberList(v string) *CreateGroupDetailRequest {
	s.NumberList = &v
	return s
}

func (s *CreateGroupDetailRequest) SetOwnerId(v int64) *CreateGroupDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGroupDetailRequest) SetPoolKey(v string) *CreateGroupDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *CreateGroupDetailRequest) SetProdCode(v string) *CreateGroupDetailRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateGroupDetailRequest) SetResourceOwnerAccount(v string) *CreateGroupDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGroupDetailRequest) SetResourceOwnerId(v int64) *CreateGroupDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateGroupDetailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupDetailResponseBody) SetRequestId(v string) *CreateGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupDetailResponseBody) SetCode(v string) *CreateGroupDetailResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupDetailResponseBody) SetData(v bool) *CreateGroupDetailResponseBody {
	s.Data = &v
	return s
}

type CreateGroupDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupDetailResponse) SetHeaders(v map[string]*string) *CreateGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupDetailResponse) SetStatusCode(v int32) *CreateGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupDetailResponse) SetBody(v *CreateGroupDetailResponseBody) *CreateGroupDetailResponse {
	s.Body = v
	return s
}

type CreateGroupInfoRequest struct {
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NumberList           *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupInfoRequest) SetName(v string) *CreateGroupInfoRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupInfoRequest) SetNumberList(v string) *CreateGroupInfoRequest {
	s.NumberList = &v
	return s
}

func (s *CreateGroupInfoRequest) SetOwnerId(v int64) *CreateGroupInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGroupInfoRequest) SetPoolKey(v string) *CreateGroupInfoRequest {
	s.PoolKey = &v
	return s
}

func (s *CreateGroupInfoRequest) SetProdCode(v string) *CreateGroupInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateGroupInfoRequest) SetResourceOwnerAccount(v string) *CreateGroupInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGroupInfoRequest) SetResourceOwnerId(v int64) *CreateGroupInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateGroupInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupInfoResponseBody) SetRequestId(v string) *CreateGroupInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupInfoResponseBody) SetCode(v string) *CreateGroupInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupInfoResponseBody) SetData(v bool) *CreateGroupInfoResponseBody {
	s.Data = &v
	return s
}

type CreateGroupInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupInfoResponse) SetHeaders(v map[string]*string) *CreateGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupInfoResponse) SetStatusCode(v int32) *CreateGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupInfoResponse) SetBody(v *CreateGroupInfoResponseBody) *CreateGroupInfoResponse {
	s.Body = v
	return s
}

type CreateLogicalDeleteRequest struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ProdCode       *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
}

func (s CreateLogicalDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogicalDeleteRequest) GoString() string {
	return s.String()
}

func (s *CreateLogicalDeleteRequest) SetBid(v string) *CreateLogicalDeleteRequest {
	s.Bid = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetCountry(v string) *CreateLogicalDeleteRequest {
	s.Country = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetGmtWakeup(v string) *CreateLogicalDeleteRequest {
	s.GmtWakeup = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetHid(v int64) *CreateLogicalDeleteRequest {
	s.Hid = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetInterrupt(v bool) *CreateLogicalDeleteRequest {
	s.Interrupt = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetInvoker(v string) *CreateLogicalDeleteRequest {
	s.Invoker = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetMessage(v string) *CreateLogicalDeleteRequest {
	s.Message = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetPk(v string) *CreateLogicalDeleteRequest {
	s.Pk = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetProdCode(v string) *CreateLogicalDeleteRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetSuccess(v bool) *CreateLogicalDeleteRequest {
	s.Success = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetTaskExtraData(v string) *CreateLogicalDeleteRequest {
	s.TaskExtraData = &v
	return s
}

func (s *CreateLogicalDeleteRequest) SetTaskIdentifier(v string) *CreateLogicalDeleteRequest {
	s.TaskIdentifier = &v
	return s
}

type CreateLogicalDeleteResponseBody struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
}

func (s CreateLogicalDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLogicalDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogicalDeleteResponseBody) SetBid(v string) *CreateLogicalDeleteResponseBody {
	s.Bid = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetCountry(v string) *CreateLogicalDeleteResponseBody {
	s.Country = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetGmtWakeup(v string) *CreateLogicalDeleteResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetHid(v int64) *CreateLogicalDeleteResponseBody {
	s.Hid = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetInterrupt(v bool) *CreateLogicalDeleteResponseBody {
	s.Interrupt = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetInvoker(v string) *CreateLogicalDeleteResponseBody {
	s.Invoker = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetMessage(v string) *CreateLogicalDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetPk(v string) *CreateLogicalDeleteResponseBody {
	s.Pk = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetSuccess(v bool) *CreateLogicalDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetTaskExtraData(v string) *CreateLogicalDeleteResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *CreateLogicalDeleteResponseBody) SetTaskIdentifier(v string) *CreateLogicalDeleteResponseBody {
	s.TaskIdentifier = &v
	return s
}

type CreateLogicalDeleteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLogicalDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLogicalDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogicalDeleteResponse) GoString() string {
	return s.String()
}

func (s *CreateLogicalDeleteResponse) SetHeaders(v map[string]*string) *CreateLogicalDeleteResponse {
	s.Headers = v
	return s
}

func (s *CreateLogicalDeleteResponse) SetStatusCode(v int32) *CreateLogicalDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogicalDeleteResponse) SetBody(v *CreateLogicalDeleteResponseBody) *CreateLogicalDeleteResponse {
	s.Body = v
	return s
}

type CreateMessageCallbackRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMessageCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageCallbackRequest) SetBizType(v string) *CreateMessageCallbackRequest {
	s.BizType = &v
	return s
}

func (s *CreateMessageCallbackRequest) SetCallbackUrl(v string) *CreateMessageCallbackRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateMessageCallbackRequest) SetOwnerId(v int64) *CreateMessageCallbackRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMessageCallbackRequest) SetProdCode(v string) *CreateMessageCallbackRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateMessageCallbackRequest) SetResourceOwnerAccount(v string) *CreateMessageCallbackRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMessageCallbackRequest) SetResourceOwnerId(v int64) *CreateMessageCallbackRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateMessageCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateMessageCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageCallbackResponseBody) SetRequestId(v string) *CreateMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessageCallbackResponseBody) SetCode(v string) *CreateMessageCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMessageCallbackResponseBody) SetData(v string) *CreateMessageCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *CreateMessageCallbackResponseBody) SetMessage(v string) *CreateMessageCallbackResponseBody {
	s.Message = &v
	return s
}

type CreateMessageCallbackResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMessageCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageCallbackResponse) SetHeaders(v map[string]*string) *CreateMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageCallbackResponse) SetStatusCode(v int32) *CreateMessageCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageCallbackResponse) SetBody(v *CreateMessageCallbackResponseBody) *CreateMessageCallbackResponse {
	s.Body = v
	return s
}

type CreateMessageQueueRequest struct {
	BillIds              *string `json:"BillIds,omitempty" xml:"BillIds,omitempty"`
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	QueueTitle           *string `json:"QueueTitle,omitempty" xml:"QueueTitle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMessageQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageQueueRequest) SetBillIds(v string) *CreateMessageQueueRequest {
	s.BillIds = &v
	return s
}

func (s *CreateMessageQueueRequest) SetBizType(v string) *CreateMessageQueueRequest {
	s.BizType = &v
	return s
}

func (s *CreateMessageQueueRequest) SetOwnerId(v int64) *CreateMessageQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMessageQueueRequest) SetProdCode(v string) *CreateMessageQueueRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateMessageQueueRequest) SetQueueName(v string) *CreateMessageQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateMessageQueueRequest) SetQueueTitle(v string) *CreateMessageQueueRequest {
	s.QueueTitle = &v
	return s
}

func (s *CreateMessageQueueRequest) SetResourceOwnerAccount(v string) *CreateMessageQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMessageQueueRequest) SetResourceOwnerId(v int64) *CreateMessageQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateMessageQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateMessageQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageQueueResponseBody) SetRequestId(v string) *CreateMessageQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMessageQueueResponseBody) SetCode(v string) *CreateMessageQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMessageQueueResponseBody) SetData(v string) *CreateMessageQueueResponseBody {
	s.Data = &v
	return s
}

type CreateMessageQueueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMessageQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMessageQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageQueueResponse) SetHeaders(v map[string]*string) *CreateMessageQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageQueueResponse) SetStatusCode(v int32) *CreateMessageQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageQueueResponse) SetBody(v *CreateMessageQueueResponseBody) *CreateMessageQueueResponse {
	s.Body = v
	return s
}

type CreatePhysicalDeleteRequest struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ProdCode       *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
}

func (s CreatePhysicalDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalDeleteRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalDeleteRequest) SetBid(v string) *CreatePhysicalDeleteRequest {
	s.Bid = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetCountry(v string) *CreatePhysicalDeleteRequest {
	s.Country = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetGmtWakeup(v string) *CreatePhysicalDeleteRequest {
	s.GmtWakeup = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetHid(v int64) *CreatePhysicalDeleteRequest {
	s.Hid = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetInterrupt(v bool) *CreatePhysicalDeleteRequest {
	s.Interrupt = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetInvoker(v string) *CreatePhysicalDeleteRequest {
	s.Invoker = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetMessage(v string) *CreatePhysicalDeleteRequest {
	s.Message = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetPk(v string) *CreatePhysicalDeleteRequest {
	s.Pk = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetProdCode(v string) *CreatePhysicalDeleteRequest {
	s.ProdCode = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetSuccess(v bool) *CreatePhysicalDeleteRequest {
	s.Success = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetTaskExtraData(v string) *CreatePhysicalDeleteRequest {
	s.TaskExtraData = &v
	return s
}

func (s *CreatePhysicalDeleteRequest) SetTaskIdentifier(v string) *CreatePhysicalDeleteRequest {
	s.TaskIdentifier = &v
	return s
}

type CreatePhysicalDeleteResponseBody struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
}

func (s CreatePhysicalDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhysicalDeleteResponseBody) SetBid(v string) *CreatePhysicalDeleteResponseBody {
	s.Bid = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetCountry(v string) *CreatePhysicalDeleteResponseBody {
	s.Country = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetGmtWakeup(v string) *CreatePhysicalDeleteResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetHid(v int64) *CreatePhysicalDeleteResponseBody {
	s.Hid = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetInterrupt(v bool) *CreatePhysicalDeleteResponseBody {
	s.Interrupt = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetInvoker(v string) *CreatePhysicalDeleteResponseBody {
	s.Invoker = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetMessage(v string) *CreatePhysicalDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetPk(v string) *CreatePhysicalDeleteResponseBody {
	s.Pk = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetSuccess(v bool) *CreatePhysicalDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetTaskExtraData(v string) *CreatePhysicalDeleteResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *CreatePhysicalDeleteResponseBody) SetTaskIdentifier(v string) *CreatePhysicalDeleteResponseBody {
	s.TaskIdentifier = &v
	return s
}

type CreatePhysicalDeleteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePhysicalDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePhysicalDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhysicalDeleteResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalDeleteResponse) SetHeaders(v map[string]*string) *CreatePhysicalDeleteResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalDeleteResponse) SetStatusCode(v int32) *CreatePhysicalDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhysicalDeleteResponse) SetBody(v *CreatePhysicalDeleteResponseBody) *CreatePhysicalDeleteResponse {
	s.Body = v
	return s
}

type CreatePoolInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePoolInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolInfoRequest) GoString() string {
	return s.String()
}

func (s *CreatePoolInfoRequest) SetOwnerId(v int64) *CreatePoolInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePoolInfoRequest) SetPoolName(v string) *CreatePoolInfoRequest {
	s.PoolName = &v
	return s
}

func (s *CreatePoolInfoRequest) SetProdCode(v string) *CreatePoolInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *CreatePoolInfoRequest) SetResType(v int32) *CreatePoolInfoRequest {
	s.ResType = &v
	return s
}

func (s *CreatePoolInfoRequest) SetResourceOwnerAccount(v string) *CreatePoolInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePoolInfoRequest) SetResourceOwnerId(v int64) *CreatePoolInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreatePoolInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreatePoolInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePoolInfoResponseBody) SetRequestId(v string) *CreatePoolInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePoolInfoResponseBody) SetCode(v string) *CreatePoolInfoResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePoolInfoResponseBody) SetData(v string) *CreatePoolInfoResponseBody {
	s.Data = &v
	return s
}

type CreatePoolInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePoolInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePoolInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePoolInfoResponse) GoString() string {
	return s.String()
}

func (s *CreatePoolInfoResponse) SetHeaders(v map[string]*string) *CreatePoolInfoResponse {
	s.Headers = v
	return s
}

func (s *CreatePoolInfoResponse) SetStatusCode(v int32) *CreatePoolInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePoolInfoResponse) SetBody(v *CreatePoolInfoResponseBody) *CreatePoolInfoResponse {
	s.Body = v
	return s
}

type CreateProductRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ProdId               *string `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) SetOwnerId(v int64) *CreateProductRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateProductRequest) SetProdCode(v string) *CreateProductRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateProductRequest) SetProdId(v string) *CreateProductRequest {
	s.ProdId = &v
	return s
}

func (s *CreateProductRequest) SetResourceOwnerAccount(v string) *CreateProductRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateProductRequest) SetResourceOwnerId(v int64) *CreateProductRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateProductResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProductResponseBody) SetCode(v string) *CreateProductResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProductResponseBody) SetData(v bool) *CreateProductResponseBody {
	s.Data = &v
	return s
}

type CreateProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetStatusCode(v int32) *CreateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductResponse) SetBody(v *CreateProductResponseBody) *CreateProductResponse {
	s.Body = v
	return s
}

type CreateRingToneRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	FileKey              *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayType             *string `json:"PlayType,omitempty" xml:"PlayType,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RingName             *string `json:"RingName,omitempty" xml:"RingName,omitempty"`
	Tts                  *string `json:"Tts,omitempty" xml:"Tts,omitempty"`
}

func (s CreateRingToneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRingToneRequest) GoString() string {
	return s.String()
}

func (s *CreateRingToneRequest) SetBillId(v string) *CreateRingToneRequest {
	s.BillId = &v
	return s
}

func (s *CreateRingToneRequest) SetFileKey(v string) *CreateRingToneRequest {
	s.FileKey = &v
	return s
}

func (s *CreateRingToneRequest) SetOwnerId(v int64) *CreateRingToneRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRingToneRequest) SetPlayType(v string) *CreateRingToneRequest {
	s.PlayType = &v
	return s
}

func (s *CreateRingToneRequest) SetProdCode(v string) *CreateRingToneRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateRingToneRequest) SetResourceOwnerAccount(v string) *CreateRingToneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRingToneRequest) SetResourceOwnerId(v int64) *CreateRingToneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRingToneRequest) SetRingName(v string) *CreateRingToneRequest {
	s.RingName = &v
	return s
}

func (s *CreateRingToneRequest) SetTts(v string) *CreateRingToneRequest {
	s.Tts = &v
	return s
}

type CreateRingToneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateRingToneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRingToneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRingToneResponseBody) SetRequestId(v string) *CreateRingToneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRingToneResponseBody) SetCode(v string) *CreateRingToneResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRingToneResponseBody) SetData(v string) *CreateRingToneResponseBody {
	s.Data = &v
	return s
}

type CreateRingToneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRingToneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRingToneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRingToneResponse) GoString() string {
	return s.String()
}

func (s *CreateRingToneResponse) SetHeaders(v map[string]*string) *CreateRingToneResponse {
	s.Headers = v
	return s
}

func (s *CreateRingToneResponse) SetStatusCode(v int32) *CreateRingToneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRingToneResponse) SetBody(v *CreateRingToneResponseBody) *CreateRingToneResponse {
	s.Body = v
	return s
}

type CreateSubsTrialRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneA               *string `json:"PhoneA,omitempty" xml:"PhoneA,omitempty"`
	PhoneB               *string `json:"PhoneB,omitempty" xml:"PhoneB,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *string `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSubsTrialRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubsTrialRequest) GoString() string {
	return s.String()
}

func (s *CreateSubsTrialRequest) SetOwnerId(v int64) *CreateSubsTrialRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubsTrialRequest) SetPhoneA(v string) *CreateSubsTrialRequest {
	s.PhoneA = &v
	return s
}

func (s *CreateSubsTrialRequest) SetPhoneB(v string) *CreateSubsTrialRequest {
	s.PhoneB = &v
	return s
}

func (s *CreateSubsTrialRequest) SetProdCode(v string) *CreateSubsTrialRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateSubsTrialRequest) SetResType(v string) *CreateSubsTrialRequest {
	s.ResType = &v
	return s
}

func (s *CreateSubsTrialRequest) SetResourceOwnerAccount(v string) *CreateSubsTrialRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSubsTrialRequest) SetResourceOwnerId(v int64) *CreateSubsTrialRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateSubsTrialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateSubsTrialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubsTrialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubsTrialResponseBody) SetRequestId(v string) *CreateSubsTrialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubsTrialResponseBody) SetCode(v string) *CreateSubsTrialResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubsTrialResponseBody) SetData(v bool) *CreateSubsTrialResponseBody {
	s.Data = &v
	return s
}

type CreateSubsTrialResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubsTrialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubsTrialResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubsTrialResponse) GoString() string {
	return s.String()
}

func (s *CreateSubsTrialResponse) SetHeaders(v map[string]*string) *CreateSubsTrialResponse {
	s.Headers = v
	return s
}

func (s *CreateSubsTrialResponse) SetStatusCode(v int32) *CreateSubsTrialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubsTrialResponse) SetBody(v *CreateSubsTrialResponseBody) *CreateSubsTrialResponse {
	s.Body = v
	return s
}

type CreateTransferRecordRequest struct {
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	NumberList           *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OriginBillId         *string `json:"OriginBillId,omitempty" xml:"OriginBillId,omitempty"`
	OriginName           *string `json:"OriginName,omitempty" xml:"OriginName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TargetBillId         *string `json:"TargetBillId,omitempty" xml:"TargetBillId,omitempty"`
	TargetName           *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	Total                *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
	TransferType         *int32  `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s CreateTransferRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransferRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateTransferRecordRequest) SetCity(v string) *CreateTransferRecordRequest {
	s.City = &v
	return s
}

func (s *CreateTransferRecordRequest) SetNumberList(v string) *CreateTransferRecordRequest {
	s.NumberList = &v
	return s
}

func (s *CreateTransferRecordRequest) SetOriginBillId(v string) *CreateTransferRecordRequest {
	s.OriginBillId = &v
	return s
}

func (s *CreateTransferRecordRequest) SetOriginName(v string) *CreateTransferRecordRequest {
	s.OriginName = &v
	return s
}

func (s *CreateTransferRecordRequest) SetOwnerId(v int64) *CreateTransferRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransferRecordRequest) SetProdCode(v string) *CreateTransferRecordRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateTransferRecordRequest) SetResourceOwnerAccount(v string) *CreateTransferRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransferRecordRequest) SetResourceOwnerId(v int64) *CreateTransferRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransferRecordRequest) SetTargetBillId(v string) *CreateTransferRecordRequest {
	s.TargetBillId = &v
	return s
}

func (s *CreateTransferRecordRequest) SetTargetName(v string) *CreateTransferRecordRequest {
	s.TargetName = &v
	return s
}

func (s *CreateTransferRecordRequest) SetTotal(v int32) *CreateTransferRecordRequest {
	s.Total = &v
	return s
}

func (s *CreateTransferRecordRequest) SetTransferType(v int32) *CreateTransferRecordRequest {
	s.TransferType = &v
	return s
}

type CreateTransferRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CreateTransferRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransferRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransferRecordResponseBody) SetRequestId(v string) *CreateTransferRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransferRecordResponseBody) SetCode(v string) *CreateTransferRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTransferRecordResponseBody) SetData(v string) *CreateTransferRecordResponseBody {
	s.Data = &v
	return s
}

type CreateTransferRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTransferRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransferRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransferRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateTransferRecordResponse) SetHeaders(v map[string]*string) *CreateTransferRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateTransferRecordResponse) SetStatusCode(v int32) *CreateTransferRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransferRecordResponse) SetBody(v *CreateTransferRecordResponseBody) *CreateTransferRecordResponse {
	s.Body = v
	return s
}

type DeleteCertifyInfoRequest struct {
	CertifyId            *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCertifyInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertifyInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertifyInfoRequest) SetCertifyId(v string) *DeleteCertifyInfoRequest {
	s.CertifyId = &v
	return s
}

func (s *DeleteCertifyInfoRequest) SetOwnerId(v int64) *DeleteCertifyInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCertifyInfoRequest) SetProdCode(v string) *DeleteCertifyInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *DeleteCertifyInfoRequest) SetResourceOwnerAccount(v string) *DeleteCertifyInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCertifyInfoRequest) SetResourceOwnerId(v int64) *DeleteCertifyInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteCertifyInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteCertifyInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertifyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertifyInfoResponseBody) SetRequestId(v string) *DeleteCertifyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCertifyInfoResponseBody) SetCode(v string) *DeleteCertifyInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCertifyInfoResponseBody) SetData(v string) *DeleteCertifyInfoResponseBody {
	s.Data = &v
	return s
}

type DeleteCertifyInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCertifyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCertifyInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertifyInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteCertifyInfoResponse) SetHeaders(v map[string]*string) *DeleteCertifyInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteCertifyInfoResponse) SetStatusCode(v int32) *DeleteCertifyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCertifyInfoResponse) SetBody(v *DeleteCertifyInfoResponseBody) *DeleteCertifyInfoResponse {
	s.Body = v
	return s
}

type DeleteContactsRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsRequest) SetBillId(v string) *DeleteContactsRequest {
	s.BillId = &v
	return s
}

func (s *DeleteContactsRequest) SetId(v int64) *DeleteContactsRequest {
	s.Id = &v
	return s
}

func (s *DeleteContactsRequest) SetOwnerId(v int64) *DeleteContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContactsRequest) SetProdCode(v string) *DeleteContactsRequest {
	s.ProdCode = &v
	return s
}

func (s *DeleteContactsRequest) SetResourceOwnerAccount(v string) *DeleteContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContactsRequest) SetResourceOwnerId(v int64) *DeleteContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteContactsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactsResponseBody) SetRequestId(v string) *DeleteContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactsResponseBody) SetCode(v string) *DeleteContactsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactsResponseBody) SetData(v string) *DeleteContactsResponseBody {
	s.Data = &v
	return s
}

type DeleteContactsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactsResponse) SetHeaders(v map[string]*string) *DeleteContactsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactsResponse) SetStatusCode(v int32) *DeleteContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactsResponse) SetBody(v *DeleteContactsResponseBody) *DeleteContactsResponse {
	s.Body = v
	return s
}

type DeleteGroupDetailRequest struct {
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IdList               *string `json:"IdList,omitempty" xml:"IdList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupDetailRequest) SetGroupId(v string) *DeleteGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupDetailRequest) SetIdList(v string) *DeleteGroupDetailRequest {
	s.IdList = &v
	return s
}

func (s *DeleteGroupDetailRequest) SetOwnerId(v int64) *DeleteGroupDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGroupDetailRequest) SetPoolKey(v string) *DeleteGroupDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *DeleteGroupDetailRequest) SetProdCode(v string) *DeleteGroupDetailRequest {
	s.ProdCode = &v
	return s
}

func (s *DeleteGroupDetailRequest) SetResourceOwnerAccount(v string) *DeleteGroupDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGroupDetailRequest) SetResourceOwnerId(v int64) *DeleteGroupDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteGroupDetailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupDetailResponseBody) SetRequestId(v string) *DeleteGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupDetailResponseBody) SetCode(v string) *DeleteGroupDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGroupDetailResponseBody) SetData(v bool) *DeleteGroupDetailResponseBody {
	s.Data = &v
	return s
}

type DeleteGroupDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupDetailResponse) SetHeaders(v map[string]*string) *DeleteGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupDetailResponse) SetStatusCode(v int32) *DeleteGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupDetailResponse) SetBody(v *DeleteGroupDetailResponseBody) *DeleteGroupDetailResponse {
	s.Body = v
	return s
}

type DeleteMessageCallbackRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMessageCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackRequest) SetBizType(v string) *DeleteMessageCallbackRequest {
	s.BizType = &v
	return s
}

func (s *DeleteMessageCallbackRequest) SetOwnerId(v int64) *DeleteMessageCallbackRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMessageCallbackRequest) SetProdCode(v string) *DeleteMessageCallbackRequest {
	s.ProdCode = &v
	return s
}

func (s *DeleteMessageCallbackRequest) SetResourceOwnerAccount(v string) *DeleteMessageCallbackRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMessageCallbackRequest) SetResourceOwnerId(v int64) *DeleteMessageCallbackRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteMessageCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteMessageCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackResponseBody) SetRequestId(v string) *DeleteMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageCallbackResponseBody) SetCode(v string) *DeleteMessageCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMessageCallbackResponseBody) SetData(v bool) *DeleteMessageCallbackResponseBody {
	s.Data = &v
	return s
}

type DeleteMessageCallbackResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMessageCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackResponse) SetHeaders(v map[string]*string) *DeleteMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageCallbackResponse) SetStatusCode(v int32) *DeleteMessageCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageCallbackResponse) SetBody(v *DeleteMessageCallbackResponseBody) *DeleteMessageCallbackResponse {
	s.Body = v
	return s
}

type DeleteRingToneRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteRingToneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRingToneRequest) GoString() string {
	return s.String()
}

func (s *DeleteRingToneRequest) SetBillId(v string) *DeleteRingToneRequest {
	s.BillId = &v
	return s
}

func (s *DeleteRingToneRequest) SetId(v string) *DeleteRingToneRequest {
	s.Id = &v
	return s
}

func (s *DeleteRingToneRequest) SetOwnerId(v int64) *DeleteRingToneRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRingToneRequest) SetProdCode(v string) *DeleteRingToneRequest {
	s.ProdCode = &v
	return s
}

func (s *DeleteRingToneRequest) SetResourceOwnerAccount(v string) *DeleteRingToneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRingToneRequest) SetResourceOwnerId(v int64) *DeleteRingToneRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteRingToneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteRingToneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRingToneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRingToneResponseBody) SetRequestId(v string) *DeleteRingToneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRingToneResponseBody) SetCode(v string) *DeleteRingToneResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRingToneResponseBody) SetData(v string) *DeleteRingToneResponseBody {
	s.Data = &v
	return s
}

type DeleteRingToneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRingToneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRingToneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRingToneResponse) GoString() string {
	return s.String()
}

func (s *DeleteRingToneResponse) SetHeaders(v map[string]*string) *DeleteRingToneResponse {
	s.Headers = v
	return s
}

func (s *DeleteRingToneResponse) SetStatusCode(v int32) *DeleteRingToneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRingToneResponse) SetBody(v *DeleteRingToneResponseBody) *DeleteRingToneResponse {
	s.Body = v
	return s
}

type DownloadCompleteRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DownloadCompleteRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadCompleteRequest) GoString() string {
	return s.String()
}

func (s *DownloadCompleteRequest) SetBizType(v string) *DownloadCompleteRequest {
	s.BizType = &v
	return s
}

func (s *DownloadCompleteRequest) SetOwnerId(v int64) *DownloadCompleteRequest {
	s.OwnerId = &v
	return s
}

func (s *DownloadCompleteRequest) SetProdCode(v string) *DownloadCompleteRequest {
	s.ProdCode = &v
	return s
}

func (s *DownloadCompleteRequest) SetResourceOwnerAccount(v string) *DownloadCompleteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DownloadCompleteRequest) SetResourceOwnerId(v int64) *DownloadCompleteRequest {
	s.ResourceOwnerId = &v
	return s
}

type DownloadCompleteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DownloadCompleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadCompleteResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadCompleteResponseBody) SetRequestId(v string) *DownloadCompleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadCompleteResponseBody) SetCode(v string) *DownloadCompleteResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadCompleteResponseBody) SetData(v string) *DownloadCompleteResponseBody {
	s.Data = &v
	return s
}

type DownloadCompleteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DownloadCompleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadCompleteResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadCompleteResponse) GoString() string {
	return s.String()
}

func (s *DownloadCompleteResponse) SetHeaders(v map[string]*string) *DownloadCompleteResponse {
	s.Headers = v
	return s
}

func (s *DownloadCompleteResponse) SetStatusCode(v int32) *DownloadCompleteResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadCompleteResponse) SetBody(v *DownloadCompleteResponseBody) *DownloadCompleteResponse {
	s.Body = v
	return s
}

type ExportResRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResBindStatus        *int32  `json:"ResBindStatus,omitempty" xml:"ResBindStatus,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s ExportResRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportResRequest) GoString() string {
	return s.String()
}

func (s *ExportResRequest) SetBillId(v string) *ExportResRequest {
	s.BillId = &v
	return s
}

func (s *ExportResRequest) SetOwnerId(v int64) *ExportResRequest {
	s.OwnerId = &v
	return s
}

func (s *ExportResRequest) SetProdCode(v string) *ExportResRequest {
	s.ProdCode = &v
	return s
}

func (s *ExportResRequest) SetResBindStatus(v int32) *ExportResRequest {
	s.ResBindStatus = &v
	return s
}

func (s *ExportResRequest) SetResType(v int32) *ExportResRequest {
	s.ResType = &v
	return s
}

func (s *ExportResRequest) SetResourceOwnerAccount(v string) *ExportResRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExportResRequest) SetResourceOwnerId(v int64) *ExportResRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExportResRequest) SetSecretNo(v string) *ExportResRequest {
	s.SecretNo = &v
	return s
}

type ExportResResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ExportResResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportResResponseBody) GoString() string {
	return s.String()
}

func (s *ExportResResponseBody) SetRequestId(v string) *ExportResResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportResResponseBody) SetCode(v string) *ExportResResponseBody {
	s.Code = &v
	return s
}

func (s *ExportResResponseBody) SetData(v string) *ExportResResponseBody {
	s.Data = &v
	return s
}

type ExportResResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExportResResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportResResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportResResponse) GoString() string {
	return s.String()
}

func (s *ExportResResponse) SetHeaders(v map[string]*string) *ExportResResponse {
	s.Headers = v
	return s
}

func (s *ExportResResponse) SetStatusCode(v int32) *ExportResResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportResResponse) SetBody(v *ExportResResponseBody) *ExportResResponse {
	s.Body = v
	return s
}

type GetEinvoicePdfDataRequest struct {
	AppCode      *string                            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	Customer     *GetEinvoicePdfDataRequestCustomer `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	EncryptProps map[string]*string                 `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	InvoiceCode  *string                            `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceNo    *string                            `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	Language     *string                            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign         *string                            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid         *string                            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetEinvoicePdfDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEinvoicePdfDataRequest) GoString() string {
	return s.String()
}

func (s *GetEinvoicePdfDataRequest) SetAppCode(v string) *GetEinvoicePdfDataRequest {
	s.AppCode = &v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetCustomer(v *GetEinvoicePdfDataRequestCustomer) *GetEinvoicePdfDataRequest {
	s.Customer = v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetEncryptProps(v map[string]*string) *GetEinvoicePdfDataRequest {
	s.EncryptProps = v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetInvoiceCode(v string) *GetEinvoicePdfDataRequest {
	s.InvoiceCode = &v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetInvoiceNo(v string) *GetEinvoicePdfDataRequest {
	s.InvoiceNo = &v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetLanguage(v string) *GetEinvoicePdfDataRequest {
	s.Language = &v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetSign(v string) *GetEinvoicePdfDataRequest {
	s.Sign = &v
	return s
}

func (s *GetEinvoicePdfDataRequest) SetUuid(v string) *GetEinvoicePdfDataRequest {
	s.Uuid = &v
	return s
}

type GetEinvoicePdfDataRequestCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetEinvoicePdfDataRequestCustomer) String() string {
	return tea.Prettify(s)
}

func (s GetEinvoicePdfDataRequestCustomer) GoString() string {
	return s.String()
}

func (s *GetEinvoicePdfDataRequestCustomer) SetAppCode(v string) *GetEinvoicePdfDataRequestCustomer {
	s.AppCode = &v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetCustomerId(v string) *GetEinvoicePdfDataRequestCustomer {
	s.CustomerId = &v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetCustomerSite(v string) *GetEinvoicePdfDataRequestCustomer {
	s.CustomerSite = &v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetCustomerSystem(v string) *GetEinvoicePdfDataRequestCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetEncryptProps(v map[string]*string) *GetEinvoicePdfDataRequestCustomer {
	s.EncryptProps = v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetLanguage(v string) *GetEinvoicePdfDataRequestCustomer {
	s.Language = &v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetSign(v string) *GetEinvoicePdfDataRequestCustomer {
	s.Sign = &v
	return s
}

func (s *GetEinvoicePdfDataRequestCustomer) SetUuid(v string) *GetEinvoicePdfDataRequestCustomer {
	s.Uuid = &v
	return s
}

type GetEinvoicePdfDataResponseBody struct {
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                      `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReturnValue  *GetEinvoicePdfDataResponseBodyReturnValue `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Struct"`
}

func (s GetEinvoicePdfDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEinvoicePdfDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEinvoicePdfDataResponseBody) SetErrorCode(v string) *GetEinvoicePdfDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBody) SetErrorMessage(v string) *GetEinvoicePdfDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBody) SetIsSuccess(v bool) *GetEinvoicePdfDataResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBody) SetReturnValue(v *GetEinvoicePdfDataResponseBodyReturnValue) *GetEinvoicePdfDataResponseBody {
	s.ReturnValue = v
	return s
}

type GetEinvoicePdfDataResponseBodyReturnValue struct {
	AppCode      *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	EInvoiceData []*int32           `json:"EInvoiceData,omitempty" xml:"EInvoiceData,omitempty" type:"Repeated"`
	EncryptProps map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	InvoiceCode  *string            `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceNo    *string            `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	Language     *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign         *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid         *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetEinvoicePdfDataResponseBodyReturnValue) String() string {
	return tea.Prettify(s)
}

func (s GetEinvoicePdfDataResponseBodyReturnValue) GoString() string {
	return s.String()
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetAppCode(v string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.AppCode = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetEInvoiceData(v []*int32) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.EInvoiceData = v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetEncryptProps(v map[string]*string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.EncryptProps = v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetInvoiceCode(v string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.InvoiceCode = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetInvoiceNo(v string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.InvoiceNo = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetLanguage(v string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.Language = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetSign(v string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.Sign = &v
	return s
}

func (s *GetEinvoicePdfDataResponseBodyReturnValue) SetUuid(v string) *GetEinvoicePdfDataResponseBodyReturnValue {
	s.Uuid = &v
	return s
}

type GetEinvoicePdfDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEinvoicePdfDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEinvoicePdfDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEinvoicePdfDataResponse) GoString() string {
	return s.String()
}

func (s *GetEinvoicePdfDataResponse) SetHeaders(v map[string]*string) *GetEinvoicePdfDataResponse {
	s.Headers = v
	return s
}

func (s *GetEinvoicePdfDataResponse) SetStatusCode(v int32) *GetEinvoicePdfDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEinvoicePdfDataResponse) SetBody(v *GetEinvoicePdfDataResponseBody) *GetEinvoicePdfDataResponse {
	s.Body = v
	return s
}

type GetSecretAsrInfoRequest struct {
	CallId   *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	PoolKey  *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
}

func (s GetSecretAsrInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSecretAsrInfoRequest) SetCallId(v string) *GetSecretAsrInfoRequest {
	s.CallId = &v
	return s
}

func (s *GetSecretAsrInfoRequest) SetCallTime(v string) *GetSecretAsrInfoRequest {
	s.CallTime = &v
	return s
}

func (s *GetSecretAsrInfoRequest) SetPoolKey(v string) *GetSecretAsrInfoRequest {
	s.PoolKey = &v
	return s
}

func (s *GetSecretAsrInfoRequest) SetProdCode(v string) *GetSecretAsrInfoRequest {
	s.ProdCode = &v
	return s
}

type GetSecretAsrInfoResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetSecretAsrInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretAsrInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretAsrInfoResponseBody) SetCode(v string) *GetSecretAsrInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretAsrInfoResponseBody) SetData(v []*GetSecretAsrInfoResponseBodyData) *GetSecretAsrInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSecretAsrInfoResponseBody) SetMessage(v string) *GetSecretAsrInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretAsrInfoResponseBody) SetRequestId(v string) *GetSecretAsrInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetSecretAsrInfoResponseBodyData struct {
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Text    *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetSecretAsrInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecretAsrInfoResponseBodyData) SetChannel(v string) *GetSecretAsrInfoResponseBodyData {
	s.Channel = &v
	return s
}

func (s *GetSecretAsrInfoResponseBodyData) SetText(v string) *GetSecretAsrInfoResponseBodyData {
	s.Text = &v
	return s
}

type GetSecretAsrInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSecretAsrInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecretAsrInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretAsrInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSecretAsrInfoResponse) SetHeaders(v map[string]*string) *GetSecretAsrInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSecretAsrInfoResponse) SetStatusCode(v int32) *GetSecretAsrInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretAsrInfoResponse) SetBody(v *GetSecretAsrInfoResponseBody) *GetSecretAsrInfoResponse {
	s.Body = v
	return s
}

type GetUserResourceTagStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetUserResourceTagStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceTagStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUserResourceTagStatusRequest) SetOwnerId(v int64) *GetUserResourceTagStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetUserResourceTagStatusRequest) SetProdCode(v string) *GetUserResourceTagStatusRequest {
	s.ProdCode = &v
	return s
}

func (s *GetUserResourceTagStatusRequest) SetResourceOwnerAccount(v string) *GetUserResourceTagStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetUserResourceTagStatusRequest) SetResourceOwnerId(v int64) *GetUserResourceTagStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetUserResourceTagStatusRequest) SetResourceType(v int32) *GetUserResourceTagStatusRequest {
	s.ResourceType = &v
	return s
}

type GetUserResourceTagStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResourceTagStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceTagStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResourceTagStatusResponseBody) SetCode(v string) *GetUserResourceTagStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResourceTagStatusResponseBody) SetData(v string) *GetUserResourceTagStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetUserResourceTagStatusResponseBody) SetRequestId(v string) *GetUserResourceTagStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResourceTagStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResourceTagStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResourceTagStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResourceTagStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserResourceTagStatusResponse) SetHeaders(v map[string]*string) *GetUserResourceTagStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserResourceTagStatusResponse) SetStatusCode(v int32) *GetUserResourceTagStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResourceTagStatusResponse) SetBody(v *GetUserResourceTagStatusResponseBody) *GetUserResourceTagStatusResponse {
	s.Body = v
	return s
}

type ListAsrLanguageModelsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAsrLanguageModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsrLanguageModelsRequest) GoString() string {
	return s.String()
}

func (s *ListAsrLanguageModelsRequest) SetOwnerId(v int64) *ListAsrLanguageModelsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAsrLanguageModelsRequest) SetProdCode(v string) *ListAsrLanguageModelsRequest {
	s.ProdCode = &v
	return s
}

func (s *ListAsrLanguageModelsRequest) SetResourceOwnerAccount(v string) *ListAsrLanguageModelsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAsrLanguageModelsRequest) SetResourceOwnerId(v int64) *ListAsrLanguageModelsRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListAsrLanguageModelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAsrLanguageModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsrLanguageModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsrLanguageModelsResponseBody) SetCode(v string) *ListAsrLanguageModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsrLanguageModelsResponseBody) SetData(v string) *ListAsrLanguageModelsResponseBody {
	s.Data = &v
	return s
}

func (s *ListAsrLanguageModelsResponseBody) SetMessage(v string) *ListAsrLanguageModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsrLanguageModelsResponseBody) SetRequestId(v string) *ListAsrLanguageModelsResponseBody {
	s.RequestId = &v
	return s
}

type ListAsrLanguageModelsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAsrLanguageModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAsrLanguageModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsrLanguageModelsResponse) GoString() string {
	return s.String()
}

func (s *ListAsrLanguageModelsResponse) SetHeaders(v map[string]*string) *ListAsrLanguageModelsResponse {
	s.Headers = v
	return s
}

func (s *ListAsrLanguageModelsResponse) SetStatusCode(v int32) *ListAsrLanguageModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsrLanguageModelsResponse) SetBody(v *ListAsrLanguageModelsResponseBody) *ListAsrLanguageModelsResponse {
	s.Body = v
	return s
}

type LockResourceRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s LockResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockResourceRequest) GoString() string {
	return s.String()
}

func (s *LockResourceRequest) SetBillId(v string) *LockResourceRequest {
	s.BillId = &v
	return s
}

func (s *LockResourceRequest) SetOwnerId(v int64) *LockResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *LockResourceRequest) SetProdCode(v string) *LockResourceRequest {
	s.ProdCode = &v
	return s
}

func (s *LockResourceRequest) SetResourceOwnerAccount(v string) *LockResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockResourceRequest) SetResourceOwnerId(v int64) *LockResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockResourceRequest) SetSecretNo(v string) *LockResourceRequest {
	s.SecretNo = &v
	return s
}

type LockResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s LockResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockResourceResponseBody) GoString() string {
	return s.String()
}

func (s *LockResourceResponseBody) SetRequestId(v string) *LockResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockResourceResponseBody) SetCode(v string) *LockResourceResponseBody {
	s.Code = &v
	return s
}

func (s *LockResourceResponseBody) SetData(v string) *LockResourceResponseBody {
	s.Data = &v
	return s
}

type LockResourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LockResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockResourceResponse) GoString() string {
	return s.String()
}

func (s *LockResourceResponse) SetHeaders(v map[string]*string) *LockResourceResponse {
	s.Headers = v
	return s
}

func (s *LockResourceResponse) SetStatusCode(v int32) *LockResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *LockResourceResponse) SetBody(v *LockResourceResponseBody) *LockResourceResponse {
	s.Body = v
	return s
}

type OccupySecretResRequest struct {
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	IsDisplayPool        *bool   `json:"IsDisplayPool,omitempty" xml:"IsDisplayPool,omitempty"`
	NoLike               *string `json:"NoLike,omitempty" xml:"NoLike,omitempty"`
	OrderDetailId        *int64  `json:"OrderDetailId,omitempty" xml:"OrderDetailId,omitempty"`
	OrderId              *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PartnerKey           *string `json:"PartnerKey,omitempty" xml:"PartnerKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int64  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNoType         *int32  `json:"SecretNoType,omitempty" xml:"SecretNoType,omitempty"`
	TotalCount           *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SecretNo             *string `json:"secretNo,omitempty" xml:"secretNo,omitempty"`
}

func (s OccupySecretResRequest) String() string {
	return tea.Prettify(s)
}

func (s OccupySecretResRequest) GoString() string {
	return s.String()
}

func (s *OccupySecretResRequest) SetCity(v string) *OccupySecretResRequest {
	s.City = &v
	return s
}

func (s *OccupySecretResRequest) SetIsDisplayPool(v bool) *OccupySecretResRequest {
	s.IsDisplayPool = &v
	return s
}

func (s *OccupySecretResRequest) SetNoLike(v string) *OccupySecretResRequest {
	s.NoLike = &v
	return s
}

func (s *OccupySecretResRequest) SetOrderDetailId(v int64) *OccupySecretResRequest {
	s.OrderDetailId = &v
	return s
}

func (s *OccupySecretResRequest) SetOrderId(v int64) *OccupySecretResRequest {
	s.OrderId = &v
	return s
}

func (s *OccupySecretResRequest) SetOwnerId(v int64) *OccupySecretResRequest {
	s.OwnerId = &v
	return s
}

func (s *OccupySecretResRequest) SetPartnerKey(v string) *OccupySecretResRequest {
	s.PartnerKey = &v
	return s
}

func (s *OccupySecretResRequest) SetProdCode(v string) *OccupySecretResRequest {
	s.ProdCode = &v
	return s
}

func (s *OccupySecretResRequest) SetResType(v int64) *OccupySecretResRequest {
	s.ResType = &v
	return s
}

func (s *OccupySecretResRequest) SetResourceOwnerAccount(v string) *OccupySecretResRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OccupySecretResRequest) SetResourceOwnerId(v int64) *OccupySecretResRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OccupySecretResRequest) SetSecretNoType(v int32) *OccupySecretResRequest {
	s.SecretNoType = &v
	return s
}

func (s *OccupySecretResRequest) SetTotalCount(v int32) *OccupySecretResRequest {
	s.TotalCount = &v
	return s
}

func (s *OccupySecretResRequest) SetSecretNo(v string) *OccupySecretResRequest {
	s.SecretNo = &v
	return s
}

type OccupySecretResResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OccupySecretResResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OccupySecretResResponseBody) GoString() string {
	return s.String()
}

func (s *OccupySecretResResponseBody) SetCode(v string) *OccupySecretResResponseBody {
	s.Code = &v
	return s
}

func (s *OccupySecretResResponseBody) SetData(v bool) *OccupySecretResResponseBody {
	s.Data = &v
	return s
}

func (s *OccupySecretResResponseBody) SetMessage(v string) *OccupySecretResResponseBody {
	s.Message = &v
	return s
}

func (s *OccupySecretResResponseBody) SetRequestId(v string) *OccupySecretResResponseBody {
	s.RequestId = &v
	return s
}

func (s *OccupySecretResResponseBody) SetSuccess(v bool) *OccupySecretResResponseBody {
	s.Success = &v
	return s
}

type OccupySecretResResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OccupySecretResResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OccupySecretResResponse) String() string {
	return tea.Prettify(s)
}

func (s OccupySecretResResponse) GoString() string {
	return s.String()
}

func (s *OccupySecretResResponse) SetHeaders(v map[string]*string) *OccupySecretResResponse {
	s.Headers = v
	return s
}

func (s *OccupySecretResResponse) SetStatusCode(v int32) *OccupySecretResResponse {
	s.StatusCode = &v
	return s
}

func (s *OccupySecretResResponse) SetBody(v *OccupySecretResResponseBody) *OccupySecretResResponse {
	s.Body = v
	return s
}

type OrderSucceededCallbackRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Data                 *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OrderSucceededCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderSucceededCallbackRequest) GoString() string {
	return s.String()
}

func (s *OrderSucceededCallbackRequest) SetOwnerId(v int64) *OrderSucceededCallbackRequest {
	s.OwnerId = &v
	return s
}

func (s *OrderSucceededCallbackRequest) SetProdCode(v string) *OrderSucceededCallbackRequest {
	s.ProdCode = &v
	return s
}

func (s *OrderSucceededCallbackRequest) SetResourceOwnerAccount(v string) *OrderSucceededCallbackRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OrderSucceededCallbackRequest) SetResourceOwnerId(v int64) *OrderSucceededCallbackRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OrderSucceededCallbackRequest) SetData(v string) *OrderSucceededCallbackRequest {
	s.Data = &v
	return s
}

type OrderSucceededCallbackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s OrderSucceededCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderSucceededCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *OrderSucceededCallbackResponseBody) SetCode(v string) *OrderSucceededCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *OrderSucceededCallbackResponseBody) SetData(v string) *OrderSucceededCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *OrderSucceededCallbackResponseBody) SetMessage(v string) *OrderSucceededCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *OrderSucceededCallbackResponseBody) SetRequestId(v string) *OrderSucceededCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderSucceededCallbackResponseBody) SetSuccess(v string) *OrderSucceededCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *OrderSucceededCallbackResponseBody) SetSynchro(v string) *OrderSucceededCallbackResponseBody {
	s.Synchro = &v
	return s
}

type OrderSucceededCallbackResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OrderSucceededCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OrderSucceededCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderSucceededCallbackResponse) GoString() string {
	return s.String()
}

func (s *OrderSucceededCallbackResponse) SetHeaders(v map[string]*string) *OrderSucceededCallbackResponse {
	s.Headers = v
	return s
}

func (s *OrderSucceededCallbackResponse) SetStatusCode(v int32) *OrderSucceededCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderSucceededCallbackResponse) SetBody(v *OrderSucceededCallbackResponseBody) *OrderSucceededCallbackResponse {
	s.Body = v
	return s
}

type PoolConfigRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	CallbackType         *int32  `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	FrozenDay            *int32  `json:"FrozenDay,omitempty" xml:"FrozenDay,omitempty"`
	NeedAllCallRecords   *bool   `json:"NeedAllCallRecords,omitempty" xml:"NeedAllCallRecords,omitempty"`
	OpenSmsWhite         *bool   `json:"OpenSmsWhite,omitempty" xml:"OpenSmsWhite,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolWarningLimit     *int32  `json:"PoolWarningLimit,omitempty" xml:"PoolWarningLimit,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SelectXMode          *string `json:"SelectXMode,omitempty" xml:"SelectXMode,omitempty"`
	SmartSmsWhitelist    *string `json:"SmartSmsWhitelist,omitempty" xml:"SmartSmsWhitelist,omitempty"`
	SmsChannel           *string `json:"SmsChannel,omitempty" xml:"SmsChannel,omitempty"`
}

func (s PoolConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PoolConfigRequest) GoString() string {
	return s.String()
}

func (s *PoolConfigRequest) SetBillId(v string) *PoolConfigRequest {
	s.BillId = &v
	return s
}

func (s *PoolConfigRequest) SetCallbackType(v int32) *PoolConfigRequest {
	s.CallbackType = &v
	return s
}

func (s *PoolConfigRequest) SetFrozenDay(v int32) *PoolConfigRequest {
	s.FrozenDay = &v
	return s
}

func (s *PoolConfigRequest) SetNeedAllCallRecords(v bool) *PoolConfigRequest {
	s.NeedAllCallRecords = &v
	return s
}

func (s *PoolConfigRequest) SetOpenSmsWhite(v bool) *PoolConfigRequest {
	s.OpenSmsWhite = &v
	return s
}

func (s *PoolConfigRequest) SetOwnerId(v int64) *PoolConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *PoolConfigRequest) SetPoolWarningLimit(v int32) *PoolConfigRequest {
	s.PoolWarningLimit = &v
	return s
}

func (s *PoolConfigRequest) SetProdCode(v string) *PoolConfigRequest {
	s.ProdCode = &v
	return s
}

func (s *PoolConfigRequest) SetResourceOwnerAccount(v string) *PoolConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PoolConfigRequest) SetResourceOwnerId(v int64) *PoolConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PoolConfigRequest) SetSelectXMode(v string) *PoolConfigRequest {
	s.SelectXMode = &v
	return s
}

func (s *PoolConfigRequest) SetSmartSmsWhitelist(v string) *PoolConfigRequest {
	s.SmartSmsWhitelist = &v
	return s
}

func (s *PoolConfigRequest) SetSmsChannel(v string) *PoolConfigRequest {
	s.SmsChannel = &v
	return s
}

type PoolConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PoolConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PoolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PoolConfigResponseBody) SetRequestId(v string) *PoolConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PoolConfigResponseBody) SetCode(v string) *PoolConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PoolConfigResponseBody) SetData(v string) *PoolConfigResponseBody {
	s.Data = &v
	return s
}

type PoolConfigResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PoolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PoolConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PoolConfigResponse) GoString() string {
	return s.String()
}

func (s *PoolConfigResponse) SetHeaders(v map[string]*string) *PoolConfigResponse {
	s.Headers = v
	return s
}

func (s *PoolConfigResponse) SetStatusCode(v int32) *PoolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PoolConfigResponse) SetBody(v *PoolConfigResponseBody) *PoolConfigResponse {
	s.Body = v
	return s
}

type PurchaseResourcesRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	BuyNumber            *int32  `json:"BuyNumber,omitempty" xml:"BuyNumber,omitempty"`
	IsDisplayPool        *bool   `json:"IsDisplayPool,omitempty" xml:"IsDisplayPool,omitempty"`
	NoLike               *string `json:"NoLike,omitempty" xml:"NoLike,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RegionName           *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	UsageScenarios       *string `json:"UsageScenarios,omitempty" xml:"UsageScenarios,omitempty"`
}

func (s PurchaseResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s PurchaseResourcesRequest) GoString() string {
	return s.String()
}

func (s *PurchaseResourcesRequest) SetBillId(v string) *PurchaseResourcesRequest {
	s.BillId = &v
	return s
}

func (s *PurchaseResourcesRequest) SetBuyNumber(v int32) *PurchaseResourcesRequest {
	s.BuyNumber = &v
	return s
}

func (s *PurchaseResourcesRequest) SetIsDisplayPool(v bool) *PurchaseResourcesRequest {
	s.IsDisplayPool = &v
	return s
}

func (s *PurchaseResourcesRequest) SetNoLike(v string) *PurchaseResourcesRequest {
	s.NoLike = &v
	return s
}

func (s *PurchaseResourcesRequest) SetOwnerId(v int64) *PurchaseResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *PurchaseResourcesRequest) SetProdCode(v string) *PurchaseResourcesRequest {
	s.ProdCode = &v
	return s
}

func (s *PurchaseResourcesRequest) SetRegionName(v string) *PurchaseResourcesRequest {
	s.RegionName = &v
	return s
}

func (s *PurchaseResourcesRequest) SetResType(v int32) *PurchaseResourcesRequest {
	s.ResType = &v
	return s
}

func (s *PurchaseResourcesRequest) SetResourceOwnerAccount(v string) *PurchaseResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PurchaseResourcesRequest) SetResourceOwnerId(v int64) *PurchaseResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PurchaseResourcesRequest) SetSpecId(v int64) *PurchaseResourcesRequest {
	s.SpecId = &v
	return s
}

func (s *PurchaseResourcesRequest) SetUsageScenarios(v string) *PurchaseResourcesRequest {
	s.UsageScenarios = &v
	return s
}

type PurchaseResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PurchaseResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurchaseResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseResourcesResponseBody) SetRequestId(v string) *PurchaseResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseResourcesResponseBody) SetCode(v string) *PurchaseResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *PurchaseResourcesResponseBody) SetData(v string) *PurchaseResourcesResponseBody {
	s.Data = &v
	return s
}

type PurchaseResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PurchaseResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PurchaseResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s PurchaseResourcesResponse) GoString() string {
	return s.String()
}

func (s *PurchaseResourcesResponse) SetHeaders(v map[string]*string) *PurchaseResourcesResponse {
	s.Headers = v
	return s
}

func (s *PurchaseResourcesResponse) SetStatusCode(v int32) *PurchaseResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseResourcesResponse) SetBody(v *PurchaseResourcesResponseBody) *PurchaseResourcesResponse {
	s.Body = v
	return s
}

type QueryBindingDetailsRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	SubId                *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
}

func (s QueryBindingDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBindingDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryBindingDetailsRequest) SetBillId(v string) *QueryBindingDetailsRequest {
	s.BillId = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetOwnerId(v int64) *QueryBindingDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetPageNo(v int32) *QueryBindingDetailsRequest {
	s.PageNo = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetPageSize(v int32) *QueryBindingDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetPhoneNoA(v string) *QueryBindingDetailsRequest {
	s.PhoneNoA = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetPhoneNoB(v string) *QueryBindingDetailsRequest {
	s.PhoneNoB = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetProdCode(v string) *QueryBindingDetailsRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetResType(v int32) *QueryBindingDetailsRequest {
	s.ResType = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetResourceOwnerAccount(v string) *QueryBindingDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetResourceOwnerId(v int64) *QueryBindingDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetSecretNo(v string) *QueryBindingDetailsRequest {
	s.SecretNo = &v
	return s
}

func (s *QueryBindingDetailsRequest) SetSubId(v string) *QueryBindingDetailsRequest {
	s.SubId = &v
	return s
}

type QueryBindingDetailsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryBindingDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBindingDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBindingDetailsResponseBody) SetRequestId(v string) *QueryBindingDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBindingDetailsResponseBody) SetCode(v string) *QueryBindingDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBindingDetailsResponseBody) SetData(v string) *QueryBindingDetailsResponseBody {
	s.Data = &v
	return s
}

type QueryBindingDetailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBindingDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBindingDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBindingDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryBindingDetailsResponse) SetHeaders(v map[string]*string) *QueryBindingDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryBindingDetailsResponse) SetStatusCode(v int32) *QueryBindingDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBindingDetailsResponse) SetBody(v *QueryBindingDetailsResponseBody) *QueryBindingDetailsResponse {
	s.Body = v
	return s
}

type QueryBlackListRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	BlackPrefix          *string `json:"BlackPrefix,omitempty" xml:"BlackPrefix,omitempty"`
	BlackType            *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackListRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackListRequest) SetBillId(v string) *QueryBlackListRequest {
	s.BillId = &v
	return s
}

func (s *QueryBlackListRequest) SetBlackPrefix(v string) *QueryBlackListRequest {
	s.BlackPrefix = &v
	return s
}

func (s *QueryBlackListRequest) SetBlackType(v string) *QueryBlackListRequest {
	s.BlackType = &v
	return s
}

func (s *QueryBlackListRequest) SetOwnerId(v int64) *QueryBlackListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBlackListRequest) SetPageNo(v int32) *QueryBlackListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryBlackListRequest) SetPageSize(v int32) *QueryBlackListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBlackListRequest) SetProdCode(v string) *QueryBlackListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryBlackListRequest) SetResourceOwnerAccount(v string) *QueryBlackListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryBlackListRequest) SetResourceOwnerId(v int64) *QueryBlackListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s QueryBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackListResponseBody) SetRequestId(v string) *QueryBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBlackListResponseBody) SetCode(v string) *QueryBlackListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBlackListResponseBody) SetData(v string) *QueryBlackListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryBlackListResponseBody) SetMessage(v string) *QueryBlackListResponseBody {
	s.Message = &v
	return s
}

type QueryBlackListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackListResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackListResponse) SetHeaders(v map[string]*string) *QueryBlackListResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackListResponse) SetStatusCode(v int32) *QueryBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackListResponse) SetBody(v *QueryBlackListResponseBody) *QueryBlackListResponse {
	s.Body = v
	return s
}

type QueryBuyPageInitDataRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryBuyPageInitDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageInitDataRequest) GoString() string {
	return s.String()
}

func (s *QueryBuyPageInitDataRequest) SetOwnerId(v int64) *QueryBuyPageInitDataRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBuyPageInitDataRequest) SetProdCode(v string) *QueryBuyPageInitDataRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryBuyPageInitDataRequest) SetResourceOwnerAccount(v string) *QueryBuyPageInitDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryBuyPageInitDataRequest) SetResourceOwnerId(v int64) *QueryBuyPageInitDataRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryBuyPageInitDataResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryBuyPageInitDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageInitDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBuyPageInitDataResponseBody) SetRequestId(v string) *QueryBuyPageInitDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBuyPageInitDataResponseBody) SetCode(v string) *QueryBuyPageInitDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBuyPageInitDataResponseBody) SetData(v string) *QueryBuyPageInitDataResponseBody {
	s.Data = &v
	return s
}

type QueryBuyPageInitDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBuyPageInitDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBuyPageInitDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageInitDataResponse) GoString() string {
	return s.String()
}

func (s *QueryBuyPageInitDataResponse) SetHeaders(v map[string]*string) *QueryBuyPageInitDataResponse {
	s.Headers = v
	return s
}

func (s *QueryBuyPageInitDataResponse) SetStatusCode(v int32) *QueryBuyPageInitDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBuyPageInitDataResponse) SetBody(v *QueryBuyPageInitDataResponseBody) *QueryBuyPageInitDataResponse {
	s.Body = v
	return s
}

type QueryBuyPageResCountRequest struct {
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	Like                 *string `json:"Like,omitempty" xml:"Like,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QueryBuyPageResCountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageResCountRequest) GoString() string {
	return s.String()
}

func (s *QueryBuyPageResCountRequest) SetCity(v string) *QueryBuyPageResCountRequest {
	s.City = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetLike(v string) *QueryBuyPageResCountRequest {
	s.Like = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetOwnerId(v int64) *QueryBuyPageResCountRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetProdCode(v string) *QueryBuyPageResCountRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetResType(v int32) *QueryBuyPageResCountRequest {
	s.ResType = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetResourceOwnerAccount(v string) *QueryBuyPageResCountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetResourceOwnerId(v int64) *QueryBuyPageResCountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryBuyPageResCountRequest) SetSpecId(v int64) *QueryBuyPageResCountRequest {
	s.SpecId = &v
	return s
}

type QueryBuyPageResCountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryBuyPageResCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageResCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBuyPageResCountResponseBody) SetRequestId(v string) *QueryBuyPageResCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBuyPageResCountResponseBody) SetCode(v string) *QueryBuyPageResCountResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBuyPageResCountResponseBody) SetData(v string) *QueryBuyPageResCountResponseBody {
	s.Data = &v
	return s
}

type QueryBuyPageResCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBuyPageResCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBuyPageResCountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageResCountResponse) GoString() string {
	return s.String()
}

func (s *QueryBuyPageResCountResponse) SetHeaders(v map[string]*string) *QueryBuyPageResCountResponse {
	s.Headers = v
	return s
}

func (s *QueryBuyPageResCountResponse) SetStatusCode(v int32) *QueryBuyPageResCountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBuyPageResCountResponse) SetBody(v *QueryBuyPageResCountResponseBody) *QueryBuyPageResCountResponse {
	s.Body = v
	return s
}

type QueryBuyPageResInfoRequest struct {
	Like                 *string `json:"Like,omitempty" xml:"Like,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QueryBuyPageResInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageResInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryBuyPageResInfoRequest) SetLike(v string) *QueryBuyPageResInfoRequest {
	s.Like = &v
	return s
}

func (s *QueryBuyPageResInfoRequest) SetOwnerId(v int64) *QueryBuyPageResInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBuyPageResInfoRequest) SetProdCode(v string) *QueryBuyPageResInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryBuyPageResInfoRequest) SetResType(v int32) *QueryBuyPageResInfoRequest {
	s.ResType = &v
	return s
}

func (s *QueryBuyPageResInfoRequest) SetResourceOwnerAccount(v string) *QueryBuyPageResInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryBuyPageResInfoRequest) SetResourceOwnerId(v int64) *QueryBuyPageResInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryBuyPageResInfoRequest) SetSpecId(v int64) *QueryBuyPageResInfoRequest {
	s.SpecId = &v
	return s
}

type QueryBuyPageResInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryBuyPageResInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageResInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBuyPageResInfoResponseBody) SetRequestId(v string) *QueryBuyPageResInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBuyPageResInfoResponseBody) SetCode(v string) *QueryBuyPageResInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBuyPageResInfoResponseBody) SetData(v string) *QueryBuyPageResInfoResponseBody {
	s.Data = &v
	return s
}

type QueryBuyPageResInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBuyPageResInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBuyPageResInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyPageResInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryBuyPageResInfoResponse) SetHeaders(v map[string]*string) *QueryBuyPageResInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryBuyPageResInfoResponse) SetStatusCode(v int32) *QueryBuyPageResInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBuyPageResInfoResponse) SetBody(v *QueryBuyPageResInfoResponseBody) *QueryBuyPageResInfoResponse {
	s.Body = v
	return s
}

type QueryBuyResInfoRequest struct {
	Like                 *string `json:"Like,omitempty" xml:"Like,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QueryBuyResInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyResInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryBuyResInfoRequest) SetLike(v string) *QueryBuyResInfoRequest {
	s.Like = &v
	return s
}

func (s *QueryBuyResInfoRequest) SetOwnerId(v int64) *QueryBuyResInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBuyResInfoRequest) SetProdCode(v string) *QueryBuyResInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryBuyResInfoRequest) SetResType(v int32) *QueryBuyResInfoRequest {
	s.ResType = &v
	return s
}

func (s *QueryBuyResInfoRequest) SetResourceOwnerAccount(v string) *QueryBuyResInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryBuyResInfoRequest) SetResourceOwnerId(v int64) *QueryBuyResInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryBuyResInfoRequest) SetSpecId(v int64) *QueryBuyResInfoRequest {
	s.SpecId = &v
	return s
}

type QueryBuyResInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryBuyResInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyResInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBuyResInfoResponseBody) SetRequestId(v string) *QueryBuyResInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBuyResInfoResponseBody) SetCode(v string) *QueryBuyResInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBuyResInfoResponseBody) SetData(v string) *QueryBuyResInfoResponseBody {
	s.Data = &v
	return s
}

type QueryBuyResInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBuyResInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBuyResInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBuyResInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryBuyResInfoResponse) SetHeaders(v map[string]*string) *QueryBuyResInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryBuyResInfoResponse) SetStatusCode(v int32) *QueryBuyResInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBuyResInfoResponse) SetBody(v *QueryBuyResInfoResponseBody) *QueryBuyResInfoResponse {
	s.Body = v
	return s
}

type QueryCallRecordingListRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	CallDate             *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	PhoneNoB             *string `json:"PhoneNoB,omitempty" xml:"PhoneNoB,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s QueryCallRecordingListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallRecordingListRequest) GoString() string {
	return s.String()
}

func (s *QueryCallRecordingListRequest) SetBillId(v string) *QueryCallRecordingListRequest {
	s.BillId = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetCallDate(v string) *QueryCallRecordingListRequest {
	s.CallDate = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetCallId(v string) *QueryCallRecordingListRequest {
	s.CallId = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetOwnerId(v int64) *QueryCallRecordingListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetPageNo(v int32) *QueryCallRecordingListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetPageSize(v int32) *QueryCallRecordingListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetPhoneNoA(v string) *QueryCallRecordingListRequest {
	s.PhoneNoA = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetPhoneNoB(v string) *QueryCallRecordingListRequest {
	s.PhoneNoB = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetProdCode(v string) *QueryCallRecordingListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetResType(v int32) *QueryCallRecordingListRequest {
	s.ResType = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetResourceOwnerAccount(v string) *QueryCallRecordingListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetResourceOwnerId(v int64) *QueryCallRecordingListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallRecordingListRequest) SetSecretNo(v string) *QueryCallRecordingListRequest {
	s.SecretNo = &v
	return s
}

type QueryCallRecordingListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCallRecordingListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallRecordingListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallRecordingListResponseBody) SetRequestId(v string) *QueryCallRecordingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallRecordingListResponseBody) SetCode(v string) *QueryCallRecordingListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallRecordingListResponseBody) SetData(v string) *QueryCallRecordingListResponseBody {
	s.Data = &v
	return s
}

type QueryCallRecordingListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCallRecordingListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCallRecordingListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallRecordingListResponse) GoString() string {
	return s.String()
}

func (s *QueryCallRecordingListResponse) SetHeaders(v map[string]*string) *QueryCallRecordingListResponse {
	s.Headers = v
	return s
}

func (s *QueryCallRecordingListResponse) SetStatusCode(v int32) *QueryCallRecordingListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallRecordingListResponse) SetBody(v *QueryCallRecordingListResponseBody) *QueryCallRecordingListResponse {
	s.Body = v
	return s
}

type QueryCertifyInfoListRequest struct {
	CertifyStatus        *string `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCertifyInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCertifyInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryCertifyInfoListRequest) SetCertifyStatus(v string) *QueryCertifyInfoListRequest {
	s.CertifyStatus = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetOwnerId(v int64) *QueryCertifyInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetPageNo(v int32) *QueryCertifyInfoListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetPageSize(v int32) *QueryCertifyInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetPhoneNo(v string) *QueryCertifyInfoListRequest {
	s.PhoneNo = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetProdCode(v string) *QueryCertifyInfoListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetResourceOwnerAccount(v string) *QueryCertifyInfoListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCertifyInfoListRequest) SetResourceOwnerId(v int64) *QueryCertifyInfoListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCertifyInfoListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCertifyInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCertifyInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCertifyInfoListResponseBody) SetRequestId(v string) *QueryCertifyInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCertifyInfoListResponseBody) SetCode(v string) *QueryCertifyInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCertifyInfoListResponseBody) SetData(v string) *QueryCertifyInfoListResponseBody {
	s.Data = &v
	return s
}

type QueryCertifyInfoListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCertifyInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCertifyInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCertifyInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryCertifyInfoListResponse) SetHeaders(v map[string]*string) *QueryCertifyInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryCertifyInfoListResponse) SetStatusCode(v int32) *QueryCertifyInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCertifyInfoListResponse) SetBody(v *QueryCertifyInfoListResponseBody) *QueryCertifyInfoListResponse {
	s.Body = v
	return s
}

type QueryCertifyOverviewInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCertifyOverviewInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCertifyOverviewInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCertifyOverviewInfoRequest) SetOwnerId(v int64) *QueryCertifyOverviewInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCertifyOverviewInfoRequest) SetProdCode(v string) *QueryCertifyOverviewInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCertifyOverviewInfoRequest) SetResourceOwnerAccount(v string) *QueryCertifyOverviewInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCertifyOverviewInfoRequest) SetResourceOwnerId(v int64) *QueryCertifyOverviewInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCertifyOverviewInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryCertifyOverviewInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCertifyOverviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCertifyOverviewInfoResponseBody) SetRequestId(v string) *QueryCertifyOverviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCertifyOverviewInfoResponseBody) SetCode(v string) *QueryCertifyOverviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCertifyOverviewInfoResponseBody) SetData(v string) *QueryCertifyOverviewInfoResponseBody {
	s.Data = &v
	return s
}

type QueryCertifyOverviewInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCertifyOverviewInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCertifyOverviewInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCertifyOverviewInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCertifyOverviewInfoResponse) SetHeaders(v map[string]*string) *QueryCertifyOverviewInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCertifyOverviewInfoResponse) SetStatusCode(v int32) *QueryCertifyOverviewInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCertifyOverviewInfoResponse) SetBody(v *QueryCertifyOverviewInfoResponseBody) *QueryCertifyOverviewInfoResponse {
	s.Body = v
	return s
}

type QueryContactsListRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryContactsListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContactsListRequest) GoString() string {
	return s.String()
}

func (s *QueryContactsListRequest) SetBillId(v string) *QueryContactsListRequest {
	s.BillId = &v
	return s
}

func (s *QueryContactsListRequest) SetOwnerId(v int64) *QueryContactsListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryContactsListRequest) SetProdCode(v string) *QueryContactsListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryContactsListRequest) SetResourceOwnerAccount(v string) *QueryContactsListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryContactsListRequest) SetResourceOwnerId(v int64) *QueryContactsListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryContactsListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryContactsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryContactsListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContactsListResponseBody) SetRequestId(v string) *QueryContactsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContactsListResponseBody) SetCode(v string) *QueryContactsListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryContactsListResponseBody) SetData(v string) *QueryContactsListResponseBody {
	s.Data = &v
	return s
}

type QueryContactsListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryContactsListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryContactsListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContactsListResponse) GoString() string {
	return s.String()
}

func (s *QueryContactsListResponse) SetHeaders(v map[string]*string) *QueryContactsListResponse {
	s.Headers = v
	return s
}

func (s *QueryContactsListResponse) SetStatusCode(v int32) *QueryContactsListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContactsListResponse) SetBody(v *QueryContactsListResponseBody) *QueryContactsListResponse {
	s.Body = v
	return s
}

type QueryCustInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCustInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCustInfoRequest) SetOwnerId(v int64) *QueryCustInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCustInfoRequest) SetProdCode(v string) *QueryCustInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCustInfoRequest) SetResourceOwnerAccount(v string) *QueryCustInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCustInfoRequest) SetResourceOwnerId(v int64) *QueryCustInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCustInfoResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data      *QueryCustInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryCustInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustInfoResponseBody) SetRequestId(v string) *QueryCustInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustInfoResponseBody) SetCode(v string) *QueryCustInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustInfoResponseBody) SetData(v *QueryCustInfoResponseBodyData) *QueryCustInfoResponseBody {
	s.Data = v
	return s
}

type QueryCustInfoResponseBodyData struct {
	CertifyType    *int32  `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	ContactPhone   *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	CustId         *int64  `json:"CustId,omitempty" xml:"CustId,omitempty"`
	CustName       *string `json:"CustName,omitempty" xml:"CustName,omitempty"`
	IsDayuCustomer *bool   `json:"IsDayuCustomer,omitempty" xml:"IsDayuCustomer,omitempty"`
	OsStatus       *int32  `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	PartnerId      *int64  `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	UserTag        *int64  `json:"UserTag,omitempty" xml:"UserTag,omitempty"`
	UserTag2       *int64  `json:"UserTag2,omitempty" xml:"UserTag2,omitempty"`
}

func (s QueryCustInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCustInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustInfoResponseBodyData) SetCertifyType(v int32) *QueryCustInfoResponseBodyData {
	s.CertifyType = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetContactPhone(v string) *QueryCustInfoResponseBodyData {
	s.ContactPhone = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetCustId(v int64) *QueryCustInfoResponseBodyData {
	s.CustId = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetCustName(v string) *QueryCustInfoResponseBodyData {
	s.CustName = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetIsDayuCustomer(v bool) *QueryCustInfoResponseBodyData {
	s.IsDayuCustomer = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetOsStatus(v int32) *QueryCustInfoResponseBodyData {
	s.OsStatus = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetPartnerId(v int64) *QueryCustInfoResponseBodyData {
	s.PartnerId = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetUserTag(v int64) *QueryCustInfoResponseBodyData {
	s.UserTag = &v
	return s
}

func (s *QueryCustInfoResponseBodyData) SetUserTag2(v int64) *QueryCustInfoResponseBodyData {
	s.UserTag2 = &v
	return s
}

type QueryCustInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryCustInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCustInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCustInfoResponse) SetHeaders(v map[string]*string) *QueryCustInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCustInfoResponse) SetStatusCode(v int32) *QueryCustInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustInfoResponse) SetBody(v *QueryCustInfoResponseBody) *QueryCustInfoResponse {
	s.Body = v
	return s
}

type QueryDownloadUrlRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryDownloadUrlRequest) SetBizType(v string) *QueryDownloadUrlRequest {
	s.BizType = &v
	return s
}

func (s *QueryDownloadUrlRequest) SetOwnerId(v int64) *QueryDownloadUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryDownloadUrlRequest) SetProdCode(v string) *QueryDownloadUrlRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryDownloadUrlRequest) SetResourceOwnerAccount(v string) *QueryDownloadUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryDownloadUrlRequest) SetResourceOwnerId(v int64) *QueryDownloadUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryDownloadUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDownloadUrlResponseBody) SetRequestId(v string) *QueryDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDownloadUrlResponseBody) SetCode(v string) *QueryDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDownloadUrlResponseBody) SetData(v string) *QueryDownloadUrlResponseBody {
	s.Data = &v
	return s
}

type QueryDownloadUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryDownloadUrlResponse) SetStatusCode(v int32) *QueryDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDownloadUrlResponse) SetBody(v *QueryDownloadUrlResponseBody) *QueryDownloadUrlResponse {
	s.Body = v
	return s
}

type QueryEffectiveInvoiceListByBillNosRequest struct {
	AppCode       *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	BillNo        *string            `json:"BillNo,omitempty" xml:"BillNo,omitempty"`
	EncryptProps  map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language      *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	MajorBillNo   *string            `json:"MajorBillNo,omitempty" xml:"MajorBillNo,omitempty"`
	OuCode        *string            `json:"OuCode,omitempty" xml:"OuCode,omitempty"`
	RelatedSystem *string            `json:"RelatedSystem,omitempty" xml:"RelatedSystem,omitempty"`
	Sign          *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid          *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryEffectiveInvoiceListByBillNosRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEffectiveInvoiceListByBillNosRequest) GoString() string {
	return s.String()
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetAppCode(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.AppCode = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetBillNo(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.BillNo = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetEncryptProps(v map[string]*string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.EncryptProps = v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetLanguage(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.Language = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetMajorBillNo(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.MajorBillNo = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetOuCode(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.OuCode = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetRelatedSystem(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.RelatedSystem = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetSign(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.Sign = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosRequest) SetUuid(v string) *QueryEffectiveInvoiceListByBillNosRequest {
	s.Uuid = &v
	return s
}

type QueryEffectiveInvoiceListByBillNosResponseBody struct {
	ErrorCode    *string                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                                      `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReturnValue  *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Struct"`
}

func (s QueryEffectiveInvoiceListByBillNosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEffectiveInvoiceListByBillNosResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBody) SetErrorCode(v string) *QueryEffectiveInvoiceListByBillNosResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBody) SetErrorMessage(v string) *QueryEffectiveInvoiceListByBillNosResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBody) SetIsSuccess(v bool) *QueryEffectiveInvoiceListByBillNosResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBody) SetReturnValue(v *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue) *QueryEffectiveInvoiceListByBillNosResponseBody {
	s.ReturnValue = v
	return s
}

type QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue struct {
	EncryptProps map[string]*string                                               `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	List         []*QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Sign         *string                                                          `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue) String() string {
	return tea.Prettify(s)
}

func (s QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue) GoString() string {
	return s.String()
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue) SetEncryptProps(v map[string]*string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue {
	s.EncryptProps = v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue) SetList(v []*QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue {
	s.List = v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue) SetSign(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValue {
	s.Sign = &v
	return s
}

type QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList struct {
	AliCompany       *string            `json:"AliCompany,omitempty" xml:"AliCompany,omitempty"`
	AliId            *int64             `json:"AliId,omitempty" xml:"AliId,omitempty"`
	Amount           *float64           `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCode          *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	BuildAmount      *float64           `json:"BuildAmount,omitempty" xml:"BuildAmount,omitempty"`
	Category         *string            `json:"Category,omitempty" xml:"Category,omitempty"`
	EncryptProps     map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	InvoiceNo        *string            `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	InvoiceStatus    *string            `json:"InvoiceStatus,omitempty" xml:"InvoiceStatus,omitempty"`
	InvoiceTitle     *string            `json:"InvoiceTitle,omitempty" xml:"InvoiceTitle,omitempty"`
	Language         *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	OrderItemNo      *string            `json:"OrderItemNo,omitempty" xml:"OrderItemNo,omitempty"`
	ParentContractNo *string            `json:"ParentContractNo,omitempty" xml:"ParentContractNo,omitempty"`
	Sign             *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Site             *string            `json:"Site,omitempty" xml:"Site,omitempty"`
	TaxRegisgerNo    *string            `json:"TaxRegisgerNo,omitempty" xml:"TaxRegisgerNo,omitempty"`
	Uuid             *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) String() string {
	return tea.Prettify(s)
}

func (s QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) GoString() string {
	return s.String()
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetAliCompany(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.AliCompany = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetAliId(v int64) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.AliId = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetAmount(v float64) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.Amount = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetAppCode(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.AppCode = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetBuildAmount(v float64) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.BuildAmount = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetCategory(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.Category = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetEncryptProps(v map[string]*string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.EncryptProps = v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetInvoiceNo(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.InvoiceNo = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetInvoiceStatus(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.InvoiceStatus = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetInvoiceTitle(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.InvoiceTitle = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetLanguage(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.Language = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetOrderItemNo(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.OrderItemNo = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetParentContractNo(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.ParentContractNo = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetSign(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.Sign = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetSite(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.Site = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetTaxRegisgerNo(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.TaxRegisgerNo = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList) SetUuid(v string) *QueryEffectiveInvoiceListByBillNosResponseBodyReturnValueList {
	s.Uuid = &v
	return s
}

type QueryEffectiveInvoiceListByBillNosResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEffectiveInvoiceListByBillNosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEffectiveInvoiceListByBillNosResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEffectiveInvoiceListByBillNosResponse) GoString() string {
	return s.String()
}

func (s *QueryEffectiveInvoiceListByBillNosResponse) SetHeaders(v map[string]*string) *QueryEffectiveInvoiceListByBillNosResponse {
	s.Headers = v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponse) SetStatusCode(v int32) *QueryEffectiveInvoiceListByBillNosResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEffectiveInvoiceListByBillNosResponse) SetBody(v *QueryEffectiveInvoiceListByBillNosResponseBody) *QueryEffectiveInvoiceListByBillNosResponse {
	s.Body = v
	return s
}

type QueryExportResUrlRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryExportResUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryExportResUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryExportResUrlRequest) SetBillId(v string) *QueryExportResUrlRequest {
	s.BillId = &v
	return s
}

func (s *QueryExportResUrlRequest) SetOwnerId(v int64) *QueryExportResUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryExportResUrlRequest) SetProdCode(v string) *QueryExportResUrlRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryExportResUrlRequest) SetResType(v int32) *QueryExportResUrlRequest {
	s.ResType = &v
	return s
}

func (s *QueryExportResUrlRequest) SetResourceOwnerAccount(v string) *QueryExportResUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryExportResUrlRequest) SetResourceOwnerId(v int64) *QueryExportResUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryExportResUrlRequest) SetTaskId(v string) *QueryExportResUrlRequest {
	s.TaskId = &v
	return s
}

type QueryExportResUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryExportResUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExportResUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExportResUrlResponseBody) SetRequestId(v string) *QueryExportResUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryExportResUrlResponseBody) SetCode(v string) *QueryExportResUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryExportResUrlResponseBody) SetData(v string) *QueryExportResUrlResponseBody {
	s.Data = &v
	return s
}

type QueryExportResUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryExportResUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryExportResUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExportResUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryExportResUrlResponse) SetHeaders(v map[string]*string) *QueryExportResUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryExportResUrlResponse) SetStatusCode(v int32) *QueryExportResUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExportResUrlResponse) SetBody(v *QueryExportResUrlResponseBody) *QueryExportResUrlResponse {
	s.Body = v
	return s
}

type QueryGroupDetailListRequest struct {
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryGroupDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupDetailListRequest) SetGroupId(v string) *QueryGroupDetailListRequest {
	s.GroupId = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetNumber(v string) *QueryGroupDetailListRequest {
	s.Number = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetOwnerId(v int64) *QueryGroupDetailListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetPageNo(v string) *QueryGroupDetailListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetPageSize(v string) *QueryGroupDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetPoolKey(v string) *QueryGroupDetailListRequest {
	s.PoolKey = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetProdCode(v string) *QueryGroupDetailListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetResourceOwnerAccount(v string) *QueryGroupDetailListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGroupDetailListRequest) SetResourceOwnerId(v int64) *QueryGroupDetailListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryGroupDetailListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryGroupDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupDetailListResponseBody) SetRequestId(v string) *QueryGroupDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGroupDetailListResponseBody) SetCode(v string) *QueryGroupDetailListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGroupDetailListResponseBody) SetData(v bool) *QueryGroupDetailListResponseBody {
	s.Data = &v
	return s
}

type QueryGroupDetailListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupDetailListResponse) SetHeaders(v map[string]*string) *QueryGroupDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupDetailListResponse) SetStatusCode(v int32) *QueryGroupDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupDetailListResponse) SetBody(v *QueryGroupDetailListResponseBody) *QueryGroupDetailListResponse {
	s.Body = v
	return s
}

type QueryGroupInfoListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	QueryKey             *string `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryGroupInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoListRequest) SetOwnerId(v int64) *QueryGroupInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetPageNo(v string) *QueryGroupInfoListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetPageSize(v string) *QueryGroupInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetPoolKey(v string) *QueryGroupInfoListRequest {
	s.PoolKey = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetProdCode(v string) *QueryGroupInfoListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetQueryKey(v string) *QueryGroupInfoListRequest {
	s.QueryKey = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetResourceOwnerAccount(v string) *QueryGroupInfoListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGroupInfoListRequest) SetResourceOwnerId(v int64) *QueryGroupInfoListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryGroupInfoListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryGroupInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoListResponseBody) SetRequestId(v string) *QueryGroupInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGroupInfoListResponseBody) SetCode(v string) *QueryGroupInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGroupInfoListResponseBody) SetData(v bool) *QueryGroupInfoListResponseBody {
	s.Data = &v
	return s
}

type QueryGroupInfoListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoListResponse) SetHeaders(v map[string]*string) *QueryGroupInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupInfoListResponse) SetStatusCode(v int32) *QueryGroupInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupInfoListResponse) SetBody(v *QueryGroupInfoListResponseBody) *QueryGroupInfoListResponse {
	s.Body = v
	return s
}

type QueryInvoiceInfoByRequestNoRequest struct {
	AppCode       *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	EncryptProps  map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language      *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	RelatedSystem *string            `json:"RelatedSystem,omitempty" xml:"RelatedSystem,omitempty"`
	RequestNo     *string            `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	Sign          *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid          *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoRequest) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetAppCode(v string) *QueryInvoiceInfoByRequestNoRequest {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoRequest {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetLanguage(v string) *QueryInvoiceInfoByRequestNoRequest {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetRelatedSystem(v string) *QueryInvoiceInfoByRequestNoRequest {
	s.RelatedSystem = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetRequestNo(v string) *QueryInvoiceInfoByRequestNoRequest {
	s.RequestNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetSign(v string) *QueryInvoiceInfoByRequestNoRequest {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoRequest) SetUuid(v string) *QueryInvoiceInfoByRequestNoRequest {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBody struct {
	ErrorCode    *string                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                               `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ReturnValue  *QueryInvoiceInfoByRequestNoResponseBodyReturnValue `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Struct"`
}

func (s QueryInvoiceInfoByRequestNoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBody) SetErrorCode(v string) *QueryInvoiceInfoByRequestNoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBody) SetErrorMessage(v string) *QueryInvoiceInfoByRequestNoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBody) SetIsSuccess(v bool) *QueryInvoiceInfoByRequestNoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBody) SetReturnValue(v *QueryInvoiceInfoByRequestNoResponseBodyReturnValue) *QueryInvoiceInfoByRequestNoResponseBody {
	s.ReturnValue = v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValue struct {
	EncryptProps map[string]*string                                        `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	List         []*QueryInvoiceInfoByRequestNoResponseBodyReturnValueList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Sign         *string                                                   `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValue) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValue) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValue) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValue {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValue) SetList(v []*QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) *QueryInvoiceInfoByRequestNoResponseBodyReturnValue {
	s.List = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValue) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValue {
	s.Sign = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueList struct {
	Amount                   *float64                                                             `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCode                  *string                                                              `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CurrencyCode             *string                                                              `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	Customer                 *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer      `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	DetailList               []*QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList  `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
	EncryptProps             map[string]*string                                                   `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	ExcludingTaxAmount       *float64                                                             `json:"ExcludingTaxAmount,omitempty" xml:"ExcludingTaxAmount,omitempty"`
	ExcludingTaxRedAmount    *float64                                                             `json:"ExcludingTaxRedAmount,omitempty" xml:"ExcludingTaxRedAmount,omitempty"`
	ExcludingTaxRemainAmount *float64                                                             `json:"ExcludingTaxRemainAmount,omitempty" xml:"ExcludingTaxRemainAmount,omitempty"`
	InvoiceCode              *string                                                              `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDate              *string                                                              `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	InvoiceId                *int64                                                               `json:"InvoiceId,omitempty" xml:"InvoiceId,omitempty"`
	InvoiceNo                *string                                                              `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	InvoiceStatus            *string                                                              `json:"InvoiceStatus,omitempty" xml:"InvoiceStatus,omitempty"`
	InvoiceType              *string                                                              `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	IsRed                    *bool                                                                `json:"IsRed,omitempty" xml:"IsRed,omitempty"`
	IsReissue                *bool                                                                `json:"IsReissue,omitempty" xml:"IsReissue,omitempty"`
	Language                 *string                                                              `json:"Language,omitempty" xml:"Language,omitempty"`
	LinkInvoiceCode          *string                                                              `json:"LinkInvoiceCode,omitempty" xml:"LinkInvoiceCode,omitempty"`
	LinkInvoiceNo            *string                                                              `json:"LinkInvoiceNo,omitempty" xml:"LinkInvoiceNo,omitempty"`
	LogisticsInfo            *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo `json:"LogisticsInfo,omitempty" xml:"LogisticsInfo,omitempty" type:"Struct"`
	MaterialType             *string                                                              `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	Memo                     *string                                                              `json:"Memo,omitempty" xml:"Memo,omitempty"`
	OuCode                   *string                                                              `json:"OuCode,omitempty" xml:"OuCode,omitempty"`
	PurchaserBankInfo        *string                                                              `json:"PurchaserBankInfo,omitempty" xml:"PurchaserBankInfo,omitempty"`
	PurchaserContactInfo     *string                                                              `json:"PurchaserContactInfo,omitempty" xml:"PurchaserContactInfo,omitempty"`
	PurchaserName            *string                                                              `json:"PurchaserName,omitempty" xml:"PurchaserName,omitempty"`
	PurchaserTaxNo           *string                                                              `json:"PurchaserTaxNo,omitempty" xml:"PurchaserTaxNo,omitempty"`
	RedAmount                *float64                                                             `json:"RedAmount,omitempty" xml:"RedAmount,omitempty"`
	RelatedId                *string                                                              `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	RemainAmount             *float64                                                             `json:"RemainAmount,omitempty" xml:"RemainAmount,omitempty"`
	Sign                     *string                                                              `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SiteId                   *string                                                              `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TaxAmount                *float64                                                             `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	Uuid                     *string                                                              `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.Amount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetCurrencyCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.CurrencyCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetCustomer(v *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.Customer = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetDetailList(v []*QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.DetailList = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetExcludingTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.ExcludingTaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetExcludingTaxRedAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.ExcludingTaxRedAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetExcludingTaxRemainAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.ExcludingTaxRemainAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetInvoiceCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.InvoiceCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetInvoiceDate(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.InvoiceDate = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetInvoiceId(v int64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.InvoiceId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetInvoiceNo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.InvoiceNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetInvoiceStatus(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.InvoiceStatus = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetInvoiceType(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.InvoiceType = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetIsRed(v bool) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.IsRed = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetIsReissue(v bool) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.IsReissue = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetLinkInvoiceCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.LinkInvoiceCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetLinkInvoiceNo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.LinkInvoiceNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetLogisticsInfo(v *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.LogisticsInfo = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetMaterialType(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.MaterialType = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetMemo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.Memo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetOuCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.OuCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetPurchaserBankInfo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.PurchaserBankInfo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetPurchaserContactInfo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.PurchaserContactInfo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetPurchaserName(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.PurchaserName = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetPurchaserTaxNo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.PurchaserTaxNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetRedAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.RedAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetRelatedId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.RelatedId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetRemainAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.RemainAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetSiteId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.SiteId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.TaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueList {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetCustomerId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.CustomerId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetCustomerSite(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.CustomerSite = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetCustomerSystem(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListCustomer {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList struct {
	Amount                     *float64                                                                      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCode                    *string                                                                       `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CargoName                  *string                                                                       `json:"CargoName,omitempty" xml:"CargoName,omitempty"`
	DiscountAmount             *float64                                                                      `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	DiscountTaxAmount          *float64                                                                      `json:"DiscountTaxAmount,omitempty" xml:"DiscountTaxAmount,omitempty"`
	EncryptProps               map[string]*string                                                            `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	ExcludingTaxAmount         *float64                                                                      `json:"ExcludingTaxAmount,omitempty" xml:"ExcludingTaxAmount,omitempty"`
	ExcludingTaxDiscountAmount *float64                                                                      `json:"ExcludingTaxDiscountAmount,omitempty" xml:"ExcludingTaxDiscountAmount,omitempty"`
	ExcludingTaxRedAmount      *float64                                                                      `json:"ExcludingTaxRedAmount,omitempty" xml:"ExcludingTaxRedAmount,omitempty"`
	ExcludingTaxRemainAmount   *float64                                                                      `json:"ExcludingTaxRemainAmount,omitempty" xml:"ExcludingTaxRemainAmount,omitempty"`
	ExcludingTaxUnitPrice      *float64                                                                      `json:"ExcludingTaxUnitPrice,omitempty" xml:"ExcludingTaxUnitPrice,omitempty"`
	InvoiceDetailId            *int64                                                                        `json:"InvoiceDetailId,omitempty" xml:"InvoiceDetailId,omitempty"`
	Language                   *string                                                                       `json:"Language,omitempty" xml:"Language,omitempty"`
	Model                      *string                                                                       `json:"Model,omitempty" xml:"Model,omitempty"`
	Quantity                   *float64                                                                      `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	QuantityUnit               *string                                                                       `json:"QuantityUnit,omitempty" xml:"QuantityUnit,omitempty"`
	RedAmount                  *float64                                                                      `json:"RedAmount,omitempty" xml:"RedAmount,omitempty"`
	RelatedId                  *string                                                                       `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	RemainAmount               *float64                                                                      `json:"RemainAmount,omitempty" xml:"RemainAmount,omitempty"`
	Sign                       *string                                                                       `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SourceList                 []*QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList `json:"SourceList,omitempty" xml:"SourceList,omitempty" type:"Repeated"`
	TaxAmount                  *float64                                                                      `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	TaxRate                    *float64                                                                      `json:"TaxRate,omitempty" xml:"TaxRate,omitempty"`
	UnitPrice                  *float64                                                                      `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
	Uuid                       *string                                                                       `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.Amount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetCargoName(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.CargoName = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetDiscountAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.DiscountAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetDiscountTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.DiscountTaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetExcludingTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.ExcludingTaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetExcludingTaxDiscountAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.ExcludingTaxDiscountAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetExcludingTaxRedAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.ExcludingTaxRedAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetExcludingTaxRemainAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.ExcludingTaxRemainAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetExcludingTaxUnitPrice(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.ExcludingTaxUnitPrice = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetInvoiceDetailId(v int64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.InvoiceDetailId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetModel(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.Model = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetQuantity(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.Quantity = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetQuantityUnit(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.QuantityUnit = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetRedAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.RedAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetRelatedId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.RelatedId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetRemainAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.RemainAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetSourceList(v []*QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.SourceList = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.TaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetTaxRate(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.TaxRate = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetUnitPrice(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.UnitPrice = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailList {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList struct {
	Amount                     *float64                                                                            `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCode                    *string                                                                             `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	BillAmount                 *float64                                                                            `json:"BillAmount,omitempty" xml:"BillAmount,omitempty"`
	BillDomain                 *string                                                                             `json:"BillDomain,omitempty" xml:"BillDomain,omitempty"`
	BillNo                     *string                                                                             `json:"BillNo,omitempty" xml:"BillNo,omitempty"`
	BillType                   *string                                                                             `json:"BillType,omitempty" xml:"BillType,omitempty"`
	BlueSourceId               *int64                                                                              `json:"BlueSourceId,omitempty" xml:"BlueSourceId,omitempty"`
	CanMerge                   *bool                                                                               `json:"CanMerge,omitempty" xml:"CanMerge,omitempty"`
	CargoName                  *string                                                                             `json:"CargoName,omitempty" xml:"CargoName,omitempty"`
	Category                   *string                                                                             `json:"Category,omitempty" xml:"Category,omitempty"`
	CompanyName                *string                                                                             `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	CurrencyCode               *string                                                                             `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	Customer                   *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	DiscountAmount             *float64                                                                            `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	DiscountTaxAmount          *float64                                                                            `json:"DiscountTaxAmount,omitempty" xml:"DiscountTaxAmount,omitempty"`
	EncryptProps               map[string]*string                                                                  `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	ExcludingTaxAmount         *float64                                                                            `json:"ExcludingTaxAmount,omitempty" xml:"ExcludingTaxAmount,omitempty"`
	ExcludingTaxDiscountAmount *float64                                                                            `json:"ExcludingTaxDiscountAmount,omitempty" xml:"ExcludingTaxDiscountAmount,omitempty"`
	ExcludingTaxRedAmount      *float64                                                                            `json:"ExcludingTaxRedAmount,omitempty" xml:"ExcludingTaxRedAmount,omitempty"`
	ExcludingTaxRemainAmount   *float64                                                                            `json:"ExcludingTaxRemainAmount,omitempty" xml:"ExcludingTaxRemainAmount,omitempty"`
	GmtBill                    *string                                                                             `json:"GmtBill,omitempty" xml:"GmtBill,omitempty"`
	GmtBillEnd                 *string                                                                             `json:"GmtBillEnd,omitempty" xml:"GmtBillEnd,omitempty"`
	GmtBillStart               *string                                                                             `json:"GmtBillStart,omitempty" xml:"GmtBillStart,omitempty"`
	GmtBuild                   *string                                                                             `json:"GmtBuild,omitempty" xml:"GmtBuild,omitempty"`
	IsApply                    *string                                                                             `json:"IsApply,omitempty" xml:"IsApply,omitempty"`
	Language                   *string                                                                             `json:"Language,omitempty" xml:"Language,omitempty"`
	MajorBillNo                *string                                                                             `json:"MajorBillNo,omitempty" xml:"MajorBillNo,omitempty"`
	Model                      *string                                                                             `json:"Model,omitempty" xml:"Model,omitempty"`
	OuCode                     *string                                                                             `json:"OuCode,omitempty" xml:"OuCode,omitempty"`
	ParentCategory             *string                                                                             `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ProductDomain              *string                                                                             `json:"ProductDomain,omitempty" xml:"ProductDomain,omitempty"`
	ProductId                  *string                                                                             `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName                *string                                                                             `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Quantity                   *float64                                                                            `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	QuantityUnit               *string                                                                             `json:"QuantityUnit,omitempty" xml:"QuantityUnit,omitempty"`
	RedAmount                  *float64                                                                            `json:"RedAmount,omitempty" xml:"RedAmount,omitempty"`
	RelatedId                  *string                                                                             `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	RemainAmount               *float64                                                                            `json:"RemainAmount,omitempty" xml:"RemainAmount,omitempty"`
	RevenueType                *string                                                                             `json:"RevenueType,omitempty" xml:"RevenueType,omitempty"`
	ServiceName                *string                                                                             `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Sign                       *string                                                                             `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SiteId                     *string                                                                             `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SourceId                   *int64                                                                              `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	TaxAmount                  *float64                                                                            `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	TaxRate                    *float64                                                                            `json:"TaxRate,omitempty" xml:"TaxRate,omitempty"`
	UnitPrice                  *float64                                                                            `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
	Uuid                       *string                                                                             `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Amount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetBillAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.BillAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetBillDomain(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.BillDomain = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetBillNo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.BillNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetBillType(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.BillType = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetBlueSourceId(v int64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.BlueSourceId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetCanMerge(v bool) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.CanMerge = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetCargoName(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.CargoName = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetCategory(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Category = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetCompanyName(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.CompanyName = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetCurrencyCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.CurrencyCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetCustomer(v *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Customer = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetDiscountAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.DiscountAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetDiscountTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.DiscountTaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetExcludingTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ExcludingTaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetExcludingTaxDiscountAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ExcludingTaxDiscountAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetExcludingTaxRedAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ExcludingTaxRedAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetExcludingTaxRemainAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ExcludingTaxRemainAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetGmtBill(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.GmtBill = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetGmtBillEnd(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.GmtBillEnd = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetGmtBillStart(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.GmtBillStart = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetGmtBuild(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.GmtBuild = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetIsApply(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.IsApply = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetMajorBillNo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.MajorBillNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetModel(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Model = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetOuCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.OuCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetParentCategory(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ParentCategory = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetProductDomain(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ProductDomain = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetProductId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ProductId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetProductName(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ProductName = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetQuantity(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Quantity = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetQuantityUnit(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.QuantityUnit = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetRedAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.RedAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetRelatedId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.RelatedId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetRemainAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.RemainAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetRevenueType(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.RevenueType = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetServiceName(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.ServiceName = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetSiteId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.SiteId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetSourceId(v int64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.SourceId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetTaxAmount(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.TaxAmount = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetTaxRate(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.TaxRate = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetUnitPrice(v float64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.UnitPrice = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceList {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetCustomerId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.CustomerId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetCustomerSite(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.CustomerSite = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetCustomerSystem(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListDetailListSourceListCustomer {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo struct {
	AppCode            *string                                                                      `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	Customer           *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer `json:"Customer,omitempty" xml:"Customer,omitempty" type:"Struct"`
	EncryptProps       map[string]*string                                                           `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	GmtSend            *string                                                                      `json:"GmtSend,omitempty" xml:"GmtSend,omitempty"`
	InvoiceCode        *string                                                                      `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDate        *string                                                                      `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	InvoiceId          *int64                                                                       `json:"InvoiceId,omitempty" xml:"InvoiceId,omitempty"`
	InvoiceNo          *string                                                                      `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	InvoiceNos         *string                                                                      `json:"InvoiceNos,omitempty" xml:"InvoiceNos,omitempty"`
	Language           *string                                                                      `json:"Language,omitempty" xml:"Language,omitempty"`
	LogisticsCompanies *string                                                                      `json:"LogisticsCompanies,omitempty" xml:"LogisticsCompanies,omitempty"`
	RelatedId          *string                                                                      `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	Sender             *string                                                                      `json:"Sender,omitempty" xml:"Sender,omitempty"`
	Sign               *string                                                                      `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Timestamp          *int64                                                                       `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TrackingNumber     *string                                                                      `json:"TrackingNumber,omitempty" xml:"TrackingNumber,omitempty"`
	Uuid               *string                                                                      `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetCustomer(v *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.Customer = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetGmtSend(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.GmtSend = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetInvoiceCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.InvoiceCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetInvoiceDate(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.InvoiceDate = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetInvoiceId(v int64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.InvoiceId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetInvoiceNo(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.InvoiceNo = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetInvoiceNos(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.InvoiceNos = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetLogisticsCompanies(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.LogisticsCompanies = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetRelatedId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.RelatedId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetSender(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.Sender = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetTimestamp(v int64) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.Timestamp = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetTrackingNumber(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.TrackingNumber = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfo {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer struct {
	AppCode        *string            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	CustomerId     *string            `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerSite   *string            `json:"CustomerSite,omitempty" xml:"CustomerSite,omitempty"`
	CustomerSystem *string            `json:"CustomerSystem,omitempty" xml:"CustomerSystem,omitempty"`
	EncryptProps   map[string]*string `json:"EncryptProps,omitempty" xml:"EncryptProps,omitempty"`
	Language       *string            `json:"Language,omitempty" xml:"Language,omitempty"`
	Sign           *string            `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Uuid           *string            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetAppCode(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.AppCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetCustomerId(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.CustomerId = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetCustomerSite(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.CustomerSite = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetCustomerSystem(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.CustomerSystem = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetEncryptProps(v map[string]*string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.EncryptProps = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetLanguage(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.Language = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetSign(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.Sign = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer) SetUuid(v string) *QueryInvoiceInfoByRequestNoResponseBodyReturnValueListLogisticsInfoCustomer {
	s.Uuid = &v
	return s
}

type QueryInvoiceInfoByRequestNoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryInvoiceInfoByRequestNoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInvoiceInfoByRequestNoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoiceInfoByRequestNoResponse) GoString() string {
	return s.String()
}

func (s *QueryInvoiceInfoByRequestNoResponse) SetHeaders(v map[string]*string) *QueryInvoiceInfoByRequestNoResponse {
	s.Headers = v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponse) SetStatusCode(v int32) *QueryInvoiceInfoByRequestNoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInvoiceInfoByRequestNoResponse) SetBody(v *QueryInvoiceInfoByRequestNoResponseBody) *QueryInvoiceInfoByRequestNoResponse {
	s.Body = v
	return s
}

type QueryMessageCallbackInfoRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMessageCallbackInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageCallbackInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageCallbackInfoRequest) SetBizType(v string) *QueryMessageCallbackInfoRequest {
	s.BizType = &v
	return s
}

func (s *QueryMessageCallbackInfoRequest) SetOwnerId(v int64) *QueryMessageCallbackInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMessageCallbackInfoRequest) SetProdCode(v string) *QueryMessageCallbackInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryMessageCallbackInfoRequest) SetResourceOwnerAccount(v string) *QueryMessageCallbackInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMessageCallbackInfoRequest) SetResourceOwnerId(v int64) *QueryMessageCallbackInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryMessageCallbackInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryMessageCallbackInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageCallbackInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageCallbackInfoResponseBody) SetRequestId(v string) *QueryMessageCallbackInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageCallbackInfoResponseBody) SetCode(v string) *QueryMessageCallbackInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessageCallbackInfoResponseBody) SetData(v bool) *QueryMessageCallbackInfoResponseBody {
	s.Data = &v
	return s
}

type QueryMessageCallbackInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMessageCallbackInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMessageCallbackInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageCallbackInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageCallbackInfoResponse) SetHeaders(v map[string]*string) *QueryMessageCallbackInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageCallbackInfoResponse) SetStatusCode(v int32) *QueryMessageCallbackInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageCallbackInfoResponse) SetBody(v *QueryMessageCallbackInfoResponseBody) *QueryMessageCallbackInfoResponse {
	s.Body = v
	return s
}

type QueryMessageQueueListRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMessageQueueListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageQueueListRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageQueueListRequest) SetBizType(v string) *QueryMessageQueueListRequest {
	s.BizType = &v
	return s
}

func (s *QueryMessageQueueListRequest) SetOwnerId(v int64) *QueryMessageQueueListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMessageQueueListRequest) SetProdCode(v string) *QueryMessageQueueListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryMessageQueueListRequest) SetResourceOwnerAccount(v string) *QueryMessageQueueListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMessageQueueListRequest) SetResourceOwnerId(v int64) *QueryMessageQueueListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryMessageQueueListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryMessageQueueListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageQueueListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageQueueListResponseBody) SetRequestId(v string) *QueryMessageQueueListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageQueueListResponseBody) SetCode(v string) *QueryMessageQueueListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessageQueueListResponseBody) SetData(v string) *QueryMessageQueueListResponseBody {
	s.Data = &v
	return s
}

type QueryMessageQueueListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMessageQueueListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMessageQueueListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageQueueListResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageQueueListResponse) SetHeaders(v map[string]*string) *QueryMessageQueueListResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageQueueListResponse) SetStatusCode(v int32) *QueryMessageQueueListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageQueueListResponse) SetBody(v *QueryMessageQueueListResponseBody) *QueryMessageQueueListResponse {
	s.Body = v
	return s
}

type QueryMonthlyBillInfoRequest struct {
	BillCycle            *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	ItemId               *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName             *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubjectItemId        *string `json:"SubjectItemId,omitempty" xml:"SubjectItemId,omitempty"`
}

func (s QueryMonthlyBillInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillInfoRequest) SetBillCycle(v string) *QueryMonthlyBillInfoRequest {
	s.BillCycle = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetItemId(v string) *QueryMonthlyBillInfoRequest {
	s.ItemId = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetItemName(v string) *QueryMonthlyBillInfoRequest {
	s.ItemName = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetOwnerId(v int64) *QueryMonthlyBillInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetProdCode(v string) *QueryMonthlyBillInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetResourceOwnerAccount(v string) *QueryMonthlyBillInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetResourceOwnerId(v int64) *QueryMonthlyBillInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMonthlyBillInfoRequest) SetSubjectItemId(v string) *QueryMonthlyBillInfoRequest {
	s.SubjectItemId = &v
	return s
}

type QueryMonthlyBillInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryMonthlyBillInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillInfoResponseBody) SetRequestId(v string) *QueryMonthlyBillInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonthlyBillInfoResponseBody) SetCode(v string) *QueryMonthlyBillInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonthlyBillInfoResponseBody) SetData(v string) *QueryMonthlyBillInfoResponseBody {
	s.Data = &v
	return s
}

type QueryMonthlyBillInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMonthlyBillInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonthlyBillInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillInfoResponse) SetHeaders(v map[string]*string) *QueryMonthlyBillInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthlyBillInfoResponse) SetStatusCode(v int32) *QueryMonthlyBillInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMonthlyBillInfoResponse) SetBody(v *QueryMonthlyBillInfoResponseBody) *QueryMonthlyBillInfoResponse {
	s.Body = v
	return s
}

type QueryMonthlyStatisticsInfoRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryMonthlyStatisticsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyStatisticsInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthlyStatisticsInfoRequest) SetBillId(v string) *QueryMonthlyStatisticsInfoRequest {
	s.BillId = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetEndDate(v string) *QueryMonthlyStatisticsInfoRequest {
	s.EndDate = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetOwnerId(v int64) *QueryMonthlyStatisticsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetProdCode(v string) *QueryMonthlyStatisticsInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetResType(v int32) *QueryMonthlyStatisticsInfoRequest {
	s.ResType = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetResourceOwnerAccount(v string) *QueryMonthlyStatisticsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetResourceOwnerId(v int64) *QueryMonthlyStatisticsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoRequest) SetStartDate(v string) *QueryMonthlyStatisticsInfoRequest {
	s.StartDate = &v
	return s
}

type QueryMonthlyStatisticsInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryMonthlyStatisticsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyStatisticsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthlyStatisticsInfoResponseBody) SetRequestId(v string) *QueryMonthlyStatisticsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoResponseBody) SetCode(v string) *QueryMonthlyStatisticsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoResponseBody) SetData(v string) *QueryMonthlyStatisticsInfoResponseBody {
	s.Data = &v
	return s
}

type QueryMonthlyStatisticsInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMonthlyStatisticsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonthlyStatisticsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyStatisticsInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthlyStatisticsInfoResponse) SetHeaders(v map[string]*string) *QueryMonthlyStatisticsInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthlyStatisticsInfoResponse) SetStatusCode(v int32) *QueryMonthlyStatisticsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMonthlyStatisticsInfoResponse) SetBody(v *QueryMonthlyStatisticsInfoResponseBody) *QueryMonthlyStatisticsInfoResponse {
	s.Body = v
	return s
}

type QueryNoBuyTasksRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryNoBuyTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryNoBuyTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryNoBuyTasksRequest) SetBillId(v string) *QueryNoBuyTasksRequest {
	s.BillId = &v
	return s
}

func (s *QueryNoBuyTasksRequest) SetOwnerId(v int64) *QueryNoBuyTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryNoBuyTasksRequest) SetPageNo(v int32) *QueryNoBuyTasksRequest {
	s.PageNo = &v
	return s
}

func (s *QueryNoBuyTasksRequest) SetPageSize(v int32) *QueryNoBuyTasksRequest {
	s.PageSize = &v
	return s
}

func (s *QueryNoBuyTasksRequest) SetProdCode(v string) *QueryNoBuyTasksRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryNoBuyTasksRequest) SetResourceOwnerAccount(v string) *QueryNoBuyTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryNoBuyTasksRequest) SetResourceOwnerId(v int64) *QueryNoBuyTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryNoBuyTasksResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s QueryNoBuyTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryNoBuyTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNoBuyTasksResponseBody) SetRequestId(v string) *QueryNoBuyTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNoBuyTasksResponseBody) SetCode(v string) *QueryNoBuyTasksResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNoBuyTasksResponseBody) SetData(v string) *QueryNoBuyTasksResponseBody {
	s.Data = &v
	return s
}

func (s *QueryNoBuyTasksResponseBody) SetMessage(v string) *QueryNoBuyTasksResponseBody {
	s.Message = &v
	return s
}

type QueryNoBuyTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryNoBuyTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryNoBuyTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryNoBuyTasksResponse) GoString() string {
	return s.String()
}

func (s *QueryNoBuyTasksResponse) SetHeaders(v map[string]*string) *QueryNoBuyTasksResponse {
	s.Headers = v
	return s
}

func (s *QueryNoBuyTasksResponse) SetStatusCode(v int32) *QueryNoBuyTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNoBuyTasksResponse) SetBody(v *QueryNoBuyTasksResponseBody) *QueryNoBuyTasksResponse {
	s.Body = v
	return s
}

type QueryNoDistributeRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryNoDistributeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryNoDistributeRequest) GoString() string {
	return s.String()
}

func (s *QueryNoDistributeRequest) SetBillId(v string) *QueryNoDistributeRequest {
	s.BillId = &v
	return s
}

func (s *QueryNoDistributeRequest) SetOwnerId(v int64) *QueryNoDistributeRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryNoDistributeRequest) SetProdCode(v string) *QueryNoDistributeRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryNoDistributeRequest) SetResourceOwnerAccount(v string) *QueryNoDistributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryNoDistributeRequest) SetResourceOwnerId(v int64) *QueryNoDistributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryNoDistributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryNoDistributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryNoDistributeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNoDistributeResponseBody) SetRequestId(v string) *QueryNoDistributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNoDistributeResponseBody) SetCode(v string) *QueryNoDistributeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNoDistributeResponseBody) SetData(v string) *QueryNoDistributeResponseBody {
	s.Data = &v
	return s
}

type QueryNoDistributeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryNoDistributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryNoDistributeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryNoDistributeResponse) GoString() string {
	return s.String()
}

func (s *QueryNoDistributeResponse) SetHeaders(v map[string]*string) *QueryNoDistributeResponse {
	s.Headers = v
	return s
}

func (s *QueryNoDistributeResponse) SetStatusCode(v int32) *QueryNoDistributeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNoDistributeResponse) SetBody(v *QueryNoDistributeResponseBody) *QueryNoDistributeResponse {
	s.Body = v
	return s
}

type QueryOpenStatusRequest struct {
	BusOffer             *int64  `json:"BusOffer,omitempty" xml:"BusOffer,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ProdId               *string `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryOpenStatusRequest) SetBusOffer(v int64) *QueryOpenStatusRequest {
	s.BusOffer = &v
	return s
}

func (s *QueryOpenStatusRequest) SetOwnerId(v int64) *QueryOpenStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryOpenStatusRequest) SetProdCode(v string) *QueryOpenStatusRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryOpenStatusRequest) SetProdId(v string) *QueryOpenStatusRequest {
	s.ProdId = &v
	return s
}

func (s *QueryOpenStatusRequest) SetResourceOwnerAccount(v string) *QueryOpenStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryOpenStatusRequest) SetResourceOwnerId(v int64) *QueryOpenStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryOpenStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOpenStatusResponseBody) SetRequestId(v string) *QueryOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOpenStatusResponseBody) SetCode(v string) *QueryOpenStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOpenStatusResponseBody) SetData(v bool) *QueryOpenStatusResponseBody {
	s.Data = &v
	return s
}

type QueryOpenStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryOpenStatusResponse) SetHeaders(v map[string]*string) *QueryOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryOpenStatusResponse) SetStatusCode(v int32) *QueryOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOpenStatusResponse) SetBody(v *QueryOpenStatusResponseBody) *QueryOpenStatusResponse {
	s.Body = v
	return s
}

type QueryPackageDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPackageDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryPackageDetailRequest) SetOwnerId(v int64) *QueryPackageDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPackageDetailRequest) SetPageNo(v int32) *QueryPackageDetailRequest {
	s.PageNo = &v
	return s
}

func (s *QueryPackageDetailRequest) SetPageSize(v int32) *QueryPackageDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPackageDetailRequest) SetProdCode(v string) *QueryPackageDetailRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPackageDetailRequest) SetResourceOwnerAccount(v string) *QueryPackageDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPackageDetailRequest) SetResourceOwnerId(v int64) *QueryPackageDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPackageDetailRequest) SetStatus(v string) *QueryPackageDetailRequest {
	s.Status = &v
	return s
}

type QueryPackageDetailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPackageDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPackageDetailResponseBody) SetRequestId(v string) *QueryPackageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPackageDetailResponseBody) SetCode(v string) *QueryPackageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPackageDetailResponseBody) SetData(v string) *QueryPackageDetailResponseBody {
	s.Data = &v
	return s
}

type QueryPackageDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPackageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPackageDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryPackageDetailResponse) SetHeaders(v map[string]*string) *QueryPackageDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryPackageDetailResponse) SetStatusCode(v int32) *QueryPackageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPackageDetailResponse) SetBody(v *QueryPackageDetailResponseBody) *QueryPackageDetailResponse {
	s.Body = v
	return s
}

type QueryPackageListRequest struct {
	BillCycle            *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPackageListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageListRequest) GoString() string {
	return s.String()
}

func (s *QueryPackageListRequest) SetBillCycle(v string) *QueryPackageListRequest {
	s.BillCycle = &v
	return s
}

func (s *QueryPackageListRequest) SetOwnerId(v int64) *QueryPackageListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPackageListRequest) SetProdCode(v string) *QueryPackageListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPackageListRequest) SetResourceOwnerAccount(v string) *QueryPackageListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPackageListRequest) SetResourceOwnerId(v int64) *QueryPackageListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPackageListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPackageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPackageListResponseBody) SetRequestId(v string) *QueryPackageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPackageListResponseBody) SetCode(v string) *QueryPackageListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPackageListResponseBody) SetData(v string) *QueryPackageListResponseBody {
	s.Data = &v
	return s
}

type QueryPackageListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPackageListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPackageListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageListResponse) GoString() string {
	return s.String()
}

func (s *QueryPackageListResponse) SetHeaders(v map[string]*string) *QueryPackageListResponse {
	s.Headers = v
	return s
}

func (s *QueryPackageListResponse) SetStatusCode(v int32) *QueryPackageListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPackageListResponse) SetBody(v *QueryPackageListResponseBody) *QueryPackageListResponse {
	s.Body = v
	return s
}

type QueryPackageStatisticsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPackageStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QueryPackageStatisticsRequest) SetOwnerId(v int64) *QueryPackageStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPackageStatisticsRequest) SetProdCode(v string) *QueryPackageStatisticsRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPackageStatisticsRequest) SetResourceOwnerAccount(v string) *QueryPackageStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPackageStatisticsRequest) SetResourceOwnerId(v int64) *QueryPackageStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPackageStatisticsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPackageStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPackageStatisticsResponseBody) SetRequestId(v string) *QueryPackageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPackageStatisticsResponseBody) SetCode(v string) *QueryPackageStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPackageStatisticsResponseBody) SetData(v string) *QueryPackageStatisticsResponseBody {
	s.Data = &v
	return s
}

type QueryPackageStatisticsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPackageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPackageStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPackageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QueryPackageStatisticsResponse) SetHeaders(v map[string]*string) *QueryPackageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QueryPackageStatisticsResponse) SetStatusCode(v int32) *QueryPackageStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPackageStatisticsResponse) SetBody(v *QueryPackageStatisticsResponseBody) *QueryPackageStatisticsResponse {
	s.Body = v
	return s
}

type QueryPoolCityListRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPoolCityListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolCityListRequest) GoString() string {
	return s.String()
}

func (s *QueryPoolCityListRequest) SetBillId(v string) *QueryPoolCityListRequest {
	s.BillId = &v
	return s
}

func (s *QueryPoolCityListRequest) SetOwnerId(v int64) *QueryPoolCityListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPoolCityListRequest) SetProdCode(v string) *QueryPoolCityListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPoolCityListRequest) SetResourceOwnerAccount(v string) *QueryPoolCityListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPoolCityListRequest) SetResourceOwnerId(v int64) *QueryPoolCityListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPoolCityListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPoolCityListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolCityListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPoolCityListResponseBody) SetRequestId(v string) *QueryPoolCityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPoolCityListResponseBody) SetCode(v string) *QueryPoolCityListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPoolCityListResponseBody) SetData(v string) *QueryPoolCityListResponseBody {
	s.Data = &v
	return s
}

type QueryPoolCityListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPoolCityListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPoolCityListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolCityListResponse) GoString() string {
	return s.String()
}

func (s *QueryPoolCityListResponse) SetHeaders(v map[string]*string) *QueryPoolCityListResponse {
	s.Headers = v
	return s
}

func (s *QueryPoolCityListResponse) SetStatusCode(v int32) *QueryPoolCityListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPoolCityListResponse) SetBody(v *QueryPoolCityListResponseBody) *QueryPoolCityListResponse {
	s.Body = v
	return s
}

type QueryPoolInfoListRequest struct {
	IsFuzzyQuery         *bool   `json:"IsFuzzyQuery,omitempty" xml:"IsFuzzyQuery,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SearchParam          *string `json:"SearchParam,omitempty" xml:"SearchParam,omitempty"`
}

func (s QueryPoolInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryPoolInfoListRequest) SetIsFuzzyQuery(v bool) *QueryPoolInfoListRequest {
	s.IsFuzzyQuery = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetOwnerId(v int64) *QueryPoolInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetPageNo(v string) *QueryPoolInfoListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetPageSize(v string) *QueryPoolInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetPoolName(v string) *QueryPoolInfoListRequest {
	s.PoolName = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetProdCode(v string) *QueryPoolInfoListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetResType(v int32) *QueryPoolInfoListRequest {
	s.ResType = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetResourceOwnerAccount(v string) *QueryPoolInfoListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetResourceOwnerId(v int64) *QueryPoolInfoListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPoolInfoListRequest) SetSearchParam(v string) *QueryPoolInfoListRequest {
	s.SearchParam = &v
	return s
}

type QueryPoolInfoListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPoolInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPoolInfoListResponseBody) SetRequestId(v string) *QueryPoolInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPoolInfoListResponseBody) SetCode(v string) *QueryPoolInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPoolInfoListResponseBody) SetData(v string) *QueryPoolInfoListResponseBody {
	s.Data = &v
	return s
}

type QueryPoolInfoListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPoolInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPoolInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryPoolInfoListResponse) SetHeaders(v map[string]*string) *QueryPoolInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryPoolInfoListResponse) SetStatusCode(v int32) *QueryPoolInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPoolInfoListResponse) SetBody(v *QueryPoolInfoListResponseBody) *QueryPoolInfoListResponse {
	s.Body = v
	return s
}

type QueryPoolMonthlyBillInfoRequest struct {
	BillCycle            *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPoolMonthlyBillInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolMonthlyBillInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPoolMonthlyBillInfoRequest) SetBillCycle(v string) *QueryPoolMonthlyBillInfoRequest {
	s.BillCycle = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoRequest) SetBillId(v string) *QueryPoolMonthlyBillInfoRequest {
	s.BillId = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoRequest) SetOwnerId(v int64) *QueryPoolMonthlyBillInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoRequest) SetResourceOwnerAccount(v string) *QueryPoolMonthlyBillInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoRequest) SetResourceOwnerId(v int64) *QueryPoolMonthlyBillInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPoolMonthlyBillInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPoolMonthlyBillInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolMonthlyBillInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPoolMonthlyBillInfoResponseBody) SetCode(v string) *QueryPoolMonthlyBillInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoResponseBody) SetData(v string) *QueryPoolMonthlyBillInfoResponseBody {
	s.Data = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoResponseBody) SetRequestId(v string) *QueryPoolMonthlyBillInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryPoolMonthlyBillInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPoolMonthlyBillInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPoolMonthlyBillInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolMonthlyBillInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPoolMonthlyBillInfoResponse) SetHeaders(v map[string]*string) *QueryPoolMonthlyBillInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPoolMonthlyBillInfoResponse) SetStatusCode(v int32) *QueryPoolMonthlyBillInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPoolMonthlyBillInfoResponse) SetBody(v *QueryPoolMonthlyBillInfoResponseBody) *QueryPoolMonthlyBillInfoResponse {
	s.Body = v
	return s
}

type QueryPoolStatisticsInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPoolStatisticsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolStatisticsInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPoolStatisticsInfoRequest) SetOwnerId(v int64) *QueryPoolStatisticsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPoolStatisticsInfoRequest) SetProdCode(v string) *QueryPoolStatisticsInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPoolStatisticsInfoRequest) SetResourceOwnerAccount(v string) *QueryPoolStatisticsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPoolStatisticsInfoRequest) SetResourceOwnerId(v int64) *QueryPoolStatisticsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPoolStatisticsInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPoolStatisticsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolStatisticsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPoolStatisticsInfoResponseBody) SetRequestId(v string) *QueryPoolStatisticsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPoolStatisticsInfoResponseBody) SetCode(v string) *QueryPoolStatisticsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPoolStatisticsInfoResponseBody) SetData(v string) *QueryPoolStatisticsInfoResponseBody {
	s.Data = &v
	return s
}

type QueryPoolStatisticsInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPoolStatisticsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPoolStatisticsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolStatisticsInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPoolStatisticsInfoResponse) SetHeaders(v map[string]*string) *QueryPoolStatisticsInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPoolStatisticsInfoResponse) SetStatusCode(v int32) *QueryPoolStatisticsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPoolStatisticsInfoResponse) SetBody(v *QueryPoolStatisticsInfoResponseBody) *QueryPoolStatisticsInfoResponse {
	s.Body = v
	return s
}

type QueryPoolSummaryInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPoolSummaryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPoolSummaryInfoRequest) SetOwnerId(v int64) *QueryPoolSummaryInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPoolSummaryInfoRequest) SetProdCode(v string) *QueryPoolSummaryInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPoolSummaryInfoRequest) SetResourceOwnerAccount(v string) *QueryPoolSummaryInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPoolSummaryInfoRequest) SetResourceOwnerId(v int64) *QueryPoolSummaryInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPoolSummaryInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPoolSummaryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPoolSummaryInfoResponseBody) SetRequestId(v string) *QueryPoolSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPoolSummaryInfoResponseBody) SetCode(v string) *QueryPoolSummaryInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPoolSummaryInfoResponseBody) SetData(v string) *QueryPoolSummaryInfoResponseBody {
	s.Data = &v
	return s
}

type QueryPoolSummaryInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPoolSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPoolSummaryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPoolSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPoolSummaryInfoResponse) SetHeaders(v map[string]*string) *QueryPoolSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPoolSummaryInfoResponse) SetStatusCode(v int32) *QueryPoolSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPoolSummaryInfoResponse) SetBody(v *QueryPoolSummaryInfoResponseBody) *QueryPoolSummaryInfoResponse {
	s.Body = v
	return s
}

type QueryPurchasedInfoRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPurchasedInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPurchasedInfoRequest) SetBillId(v string) *QueryPurchasedInfoRequest {
	s.BillId = &v
	return s
}

func (s *QueryPurchasedInfoRequest) SetOwnerId(v int64) *QueryPurchasedInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPurchasedInfoRequest) SetProdCode(v string) *QueryPurchasedInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPurchasedInfoRequest) SetResType(v int32) *QueryPurchasedInfoRequest {
	s.ResType = &v
	return s
}

func (s *QueryPurchasedInfoRequest) SetResourceOwnerAccount(v string) *QueryPurchasedInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPurchasedInfoRequest) SetResourceOwnerId(v int64) *QueryPurchasedInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryPurchasedInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPurchasedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPurchasedInfoResponseBody) SetRequestId(v string) *QueryPurchasedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPurchasedInfoResponseBody) SetCode(v string) *QueryPurchasedInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPurchasedInfoResponseBody) SetData(v string) *QueryPurchasedInfoResponseBody {
	s.Data = &v
	return s
}

type QueryPurchasedInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPurchasedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPurchasedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPurchasedInfoResponse) SetHeaders(v map[string]*string) *QueryPurchasedInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPurchasedInfoResponse) SetStatusCode(v int32) *QueryPurchasedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPurchasedInfoResponse) SetBody(v *QueryPurchasedInfoResponseBody) *QueryPurchasedInfoResponse {
	s.Body = v
	return s
}

type QueryPurchasedResListRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	IsDisplayPool        *bool   `json:"IsDisplayPool,omitempty" xml:"IsDisplayPool,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResBindStatus        *int32  `json:"ResBindStatus,omitempty" xml:"ResBindStatus,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s QueryPurchasedResListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedResListRequest) GoString() string {
	return s.String()
}

func (s *QueryPurchasedResListRequest) SetBillId(v string) *QueryPurchasedResListRequest {
	s.BillId = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetCity(v string) *QueryPurchasedResListRequest {
	s.City = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetIsDisplayPool(v bool) *QueryPurchasedResListRequest {
	s.IsDisplayPool = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetOwnerId(v int64) *QueryPurchasedResListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetPageNo(v int32) *QueryPurchasedResListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetPageSize(v int32) *QueryPurchasedResListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetProdCode(v string) *QueryPurchasedResListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetResBindStatus(v int32) *QueryPurchasedResListRequest {
	s.ResBindStatus = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetResType(v int32) *QueryPurchasedResListRequest {
	s.ResType = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetResourceOwnerAccount(v string) *QueryPurchasedResListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetResourceOwnerId(v int64) *QueryPurchasedResListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPurchasedResListRequest) SetSecretNo(v string) *QueryPurchasedResListRequest {
	s.SecretNo = &v
	return s
}

type QueryPurchasedResListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryPurchasedResListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedResListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPurchasedResListResponseBody) SetRequestId(v string) *QueryPurchasedResListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPurchasedResListResponseBody) SetCode(v string) *QueryPurchasedResListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPurchasedResListResponseBody) SetData(v string) *QueryPurchasedResListResponseBody {
	s.Data = &v
	return s
}

type QueryPurchasedResListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPurchasedResListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPurchasedResListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedResListResponse) GoString() string {
	return s.String()
}

func (s *QueryPurchasedResListResponse) SetHeaders(v map[string]*string) *QueryPurchasedResListResponse {
	s.Headers = v
	return s
}

func (s *QueryPurchasedResListResponse) SetStatusCode(v int32) *QueryPurchasedResListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPurchasedResListResponse) SetBody(v *QueryPurchasedResListResponseBody) *QueryPurchasedResListResponse {
	s.Body = v
	return s
}

type QueryQRCodeInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNumber         *string `json:"SecretNumber,omitempty" xml:"SecretNumber,omitempty"`
}

func (s QueryQRCodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryQRCodeInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryQRCodeInfoRequest) SetOwnerId(v int64) *QueryQRCodeInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryQRCodeInfoRequest) SetResourceOwnerAccount(v string) *QueryQRCodeInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryQRCodeInfoRequest) SetResourceOwnerId(v int64) *QueryQRCodeInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryQRCodeInfoRequest) SetSecretNumber(v string) *QueryQRCodeInfoRequest {
	s.SecretNumber = &v
	return s
}

type QueryQRCodeInfoResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s QueryQRCodeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryQRCodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQRCodeInfoResponseBody) SetCode(v string) *QueryQRCodeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryQRCodeInfoResponseBody) SetHttpStatusCode(v int32) *QueryQRCodeInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryQRCodeInfoResponseBody) SetMessage(v string) *QueryQRCodeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryQRCodeInfoResponseBody) SetRequestId(v string) *QueryQRCodeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQRCodeInfoResponseBody) SetSuccess(v bool) *QueryQRCodeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryQRCodeInfoResponseBody) SetToken(v string) *QueryQRCodeInfoResponseBody {
	s.Token = &v
	return s
}

type QueryQRCodeInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryQRCodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryQRCodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryQRCodeInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryQRCodeInfoResponse) SetHeaders(v map[string]*string) *QueryQRCodeInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryQRCodeInfoResponse) SetStatusCode(v int32) *QueryQRCodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQRCodeInfoResponse) SetBody(v *QueryQRCodeInfoResponseBody) *QueryQRCodeInfoResponse {
	s.Body = v
	return s
}

type QueryRecordingUrlRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	CallDate             *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRecordingUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordingUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordingUrlRequest) SetBillId(v string) *QueryRecordingUrlRequest {
	s.BillId = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetCallDate(v string) *QueryRecordingUrlRequest {
	s.CallDate = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetCallId(v string) *QueryRecordingUrlRequest {
	s.CallId = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetOwnerId(v int64) *QueryRecordingUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetProdCode(v string) *QueryRecordingUrlRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetResType(v int32) *QueryRecordingUrlRequest {
	s.ResType = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetResourceOwnerAccount(v string) *QueryRecordingUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRecordingUrlRequest) SetResourceOwnerId(v int64) *QueryRecordingUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryRecordingUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryRecordingUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordingUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordingUrlResponseBody) SetRequestId(v string) *QueryRecordingUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordingUrlResponseBody) SetCode(v string) *QueryRecordingUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordingUrlResponseBody) SetData(v string) *QueryRecordingUrlResponseBody {
	s.Data = &v
	return s
}

type QueryRecordingUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRecordingUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRecordingUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordingUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordingUrlResponse) SetHeaders(v map[string]*string) *QueryRecordingUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRecordingUrlResponse) SetStatusCode(v int32) *QueryRecordingUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecordingUrlResponse) SetBody(v *QueryRecordingUrlResponseBody) *QueryRecordingUrlResponse {
	s.Body = v
	return s
}

type QueryResSummaryInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryResSummaryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryResSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryResSummaryInfoRequest) SetOwnerId(v int64) *QueryResSummaryInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryResSummaryInfoRequest) SetProdCode(v string) *QueryResSummaryInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryResSummaryInfoRequest) SetResourceOwnerAccount(v string) *QueryResSummaryInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryResSummaryInfoRequest) SetResourceOwnerId(v int64) *QueryResSummaryInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryResSummaryInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryResSummaryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResSummaryInfoResponseBody) SetRequestId(v string) *QueryResSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResSummaryInfoResponseBody) SetCode(v string) *QueryResSummaryInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResSummaryInfoResponseBody) SetData(v string) *QueryResSummaryInfoResponseBody {
	s.Data = &v
	return s
}

type QueryResSummaryInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryResSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResSummaryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryResSummaryInfoResponse) SetHeaders(v map[string]*string) *QueryResSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryResSummaryInfoResponse) SetStatusCode(v int32) *QueryResSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResSummaryInfoResponse) SetBody(v *QueryResSummaryInfoResponseBody) *QueryResSummaryInfoResponse {
	s.Body = v
	return s
}

type QueryRingToneUrlRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	FileKey              *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRingToneUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRingToneUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRingToneUrlRequest) SetBillId(v string) *QueryRingToneUrlRequest {
	s.BillId = &v
	return s
}

func (s *QueryRingToneUrlRequest) SetFileKey(v string) *QueryRingToneUrlRequest {
	s.FileKey = &v
	return s
}

func (s *QueryRingToneUrlRequest) SetOwnerId(v int64) *QueryRingToneUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRingToneUrlRequest) SetProdCode(v string) *QueryRingToneUrlRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryRingToneUrlRequest) SetResourceOwnerAccount(v string) *QueryRingToneUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRingToneUrlRequest) SetResourceOwnerId(v int64) *QueryRingToneUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryRingToneUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryRingToneUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRingToneUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRingToneUrlResponseBody) SetRequestId(v string) *QueryRingToneUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRingToneUrlResponseBody) SetCode(v string) *QueryRingToneUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRingToneUrlResponseBody) SetData(v string) *QueryRingToneUrlResponseBody {
	s.Data = &v
	return s
}

type QueryRingToneUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRingToneUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRingToneUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRingToneUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryRingToneUrlResponse) SetHeaders(v map[string]*string) *QueryRingToneUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryRingToneUrlResponse) SetStatusCode(v int32) *QueryRingToneUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRingToneUrlResponse) SetBody(v *QueryRingToneUrlResponseBody) *QueryRingToneUrlResponse {
	s.Body = v
	return s
}

type QueryRingTonesRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlayType             *string `json:"PlayType,omitempty" xml:"PlayType,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRingTonesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRingTonesRequest) GoString() string {
	return s.String()
}

func (s *QueryRingTonesRequest) SetBillId(v string) *QueryRingTonesRequest {
	s.BillId = &v
	return s
}

func (s *QueryRingTonesRequest) SetOwnerId(v int64) *QueryRingTonesRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRingTonesRequest) SetPageNo(v int32) *QueryRingTonesRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRingTonesRequest) SetPageSize(v int32) *QueryRingTonesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRingTonesRequest) SetPlayType(v string) *QueryRingTonesRequest {
	s.PlayType = &v
	return s
}

func (s *QueryRingTonesRequest) SetProdCode(v string) *QueryRingTonesRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryRingTonesRequest) SetResourceOwnerAccount(v string) *QueryRingTonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRingTonesRequest) SetResourceOwnerId(v int64) *QueryRingTonesRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryRingTonesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryRingTonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRingTonesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRingTonesResponseBody) SetRequestId(v string) *QueryRingTonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRingTonesResponseBody) SetCode(v string) *QueryRingTonesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRingTonesResponseBody) SetData(v string) *QueryRingTonesResponseBody {
	s.Data = &v
	return s
}

type QueryRingTonesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRingTonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRingTonesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRingTonesResponse) GoString() string {
	return s.String()
}

func (s *QueryRingTonesResponse) SetHeaders(v map[string]*string) *QueryRingTonesResponse {
	s.Headers = v
	return s
}

func (s *QueryRingTonesResponse) SetStatusCode(v int32) *QueryRingTonesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRingTonesResponse) SetBody(v *QueryRingTonesResponseBody) *QueryRingTonesResponse {
	s.Body = v
	return s
}

type QuerySimplePoolInfoListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySimplePoolInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySimplePoolInfoListRequest) GoString() string {
	return s.String()
}

func (s *QuerySimplePoolInfoListRequest) SetOwnerId(v int64) *QuerySimplePoolInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySimplePoolInfoListRequest) SetPoolName(v string) *QuerySimplePoolInfoListRequest {
	s.PoolName = &v
	return s
}

func (s *QuerySimplePoolInfoListRequest) SetProdCode(v string) *QuerySimplePoolInfoListRequest {
	s.ProdCode = &v
	return s
}

func (s *QuerySimplePoolInfoListRequest) SetResType(v int32) *QuerySimplePoolInfoListRequest {
	s.ResType = &v
	return s
}

func (s *QuerySimplePoolInfoListRequest) SetResourceOwnerAccount(v string) *QuerySimplePoolInfoListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySimplePoolInfoListRequest) SetResourceOwnerId(v int64) *QuerySimplePoolInfoListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QuerySimplePoolInfoListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QuerySimplePoolInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySimplePoolInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySimplePoolInfoListResponseBody) SetRequestId(v string) *QuerySimplePoolInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySimplePoolInfoListResponseBody) SetCode(v string) *QuerySimplePoolInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySimplePoolInfoListResponseBody) SetData(v string) *QuerySimplePoolInfoListResponseBody {
	s.Data = &v
	return s
}

type QuerySimplePoolInfoListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySimplePoolInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySimplePoolInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySimplePoolInfoListResponse) GoString() string {
	return s.String()
}

func (s *QuerySimplePoolInfoListResponse) SetHeaders(v map[string]*string) *QuerySimplePoolInfoListResponse {
	s.Headers = v
	return s
}

func (s *QuerySimplePoolInfoListResponse) SetStatusCode(v int32) *QuerySimplePoolInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySimplePoolInfoListResponse) SetBody(v *QuerySimplePoolInfoListResponseBody) *QuerySimplePoolInfoListResponse {
	s.Body = v
	return s
}

type QueryStatisticsInfoRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryStatisticsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryStatisticsInfoRequest) SetBillId(v string) *QueryStatisticsInfoRequest {
	s.BillId = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetEndDate(v string) *QueryStatisticsInfoRequest {
	s.EndDate = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetOwnerId(v int64) *QueryStatisticsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetPageNo(v int32) *QueryStatisticsInfoRequest {
	s.PageNo = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetPageSize(v int32) *QueryStatisticsInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetProdCode(v string) *QueryStatisticsInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetResType(v int32) *QueryStatisticsInfoRequest {
	s.ResType = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetResourceOwnerAccount(v string) *QueryStatisticsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetResourceOwnerId(v int64) *QueryStatisticsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryStatisticsInfoRequest) SetStartDate(v string) *QueryStatisticsInfoRequest {
	s.StartDate = &v
	return s
}

type QueryStatisticsInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryStatisticsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatisticsInfoResponseBody) SetRequestId(v string) *QueryStatisticsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStatisticsInfoResponseBody) SetCode(v string) *QueryStatisticsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStatisticsInfoResponseBody) SetData(v string) *QueryStatisticsInfoResponseBody {
	s.Data = &v
	return s
}

type QueryStatisticsInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStatisticsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStatisticsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStatisticsInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryStatisticsInfoResponse) SetHeaders(v map[string]*string) *QueryStatisticsInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryStatisticsInfoResponse) SetStatusCode(v int32) *QueryStatisticsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStatisticsInfoResponse) SetBody(v *QueryStatisticsInfoResponseBody) *QueryStatisticsInfoResponse {
	s.Body = v
	return s
}

type QueryTagOpenStatusRequest struct {
	AttributeKey         *string `json:"AttributeKey,omitempty" xml:"AttributeKey,omitempty"`
	BizType              *int32  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SubAttributeKey      *string `json:"SubAttributeKey,omitempty" xml:"SubAttributeKey,omitempty"`
}

func (s QueryTagOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTagOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryTagOpenStatusRequest) SetAttributeKey(v string) *QueryTagOpenStatusRequest {
	s.AttributeKey = &v
	return s
}

func (s *QueryTagOpenStatusRequest) SetBizType(v int32) *QueryTagOpenStatusRequest {
	s.BizType = &v
	return s
}

func (s *QueryTagOpenStatusRequest) SetOwnerId(v int64) *QueryTagOpenStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagOpenStatusRequest) SetResourceOwnerAccount(v string) *QueryTagOpenStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagOpenStatusRequest) SetResourceOwnerId(v int64) *QueryTagOpenStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagOpenStatusRequest) SetSubAttributeKey(v string) *QueryTagOpenStatusRequest {
	s.SubAttributeKey = &v
	return s
}

type QueryTagOpenStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTagOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTagOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagOpenStatusResponseBody) SetCode(v string) *QueryTagOpenStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagOpenStatusResponseBody) SetData(v bool) *QueryTagOpenStatusResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTagOpenStatusResponseBody) SetRequestId(v string) *QueryTagOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryTagOpenStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTagOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTagOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTagOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryTagOpenStatusResponse) SetHeaders(v map[string]*string) *QueryTagOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryTagOpenStatusResponse) SetStatusCode(v int32) *QueryTagOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagOpenStatusResponse) SetBody(v *QueryTagOpenStatusResponseBody) *QueryTagOpenStatusResponse {
	s.Body = v
	return s
}

type QueryTransferDetailsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RecordId             *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTransferDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferDetailsRequest) SetOwnerId(v int64) *QueryTransferDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTransferDetailsRequest) SetPageNo(v int32) *QueryTransferDetailsRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTransferDetailsRequest) SetPageSize(v int32) *QueryTransferDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTransferDetailsRequest) SetProdCode(v string) *QueryTransferDetailsRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryTransferDetailsRequest) SetRecordId(v int64) *QueryTransferDetailsRequest {
	s.RecordId = &v
	return s
}

func (s *QueryTransferDetailsRequest) SetResourceOwnerAccount(v string) *QueryTransferDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTransferDetailsRequest) SetResourceOwnerId(v int64) *QueryTransferDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryTransferDetailsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryTransferDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferDetailsResponseBody) SetRequestId(v string) *QueryTransferDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferDetailsResponseBody) SetCode(v string) *QueryTransferDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTransferDetailsResponseBody) SetData(v string) *QueryTransferDetailsResponseBody {
	s.Data = &v
	return s
}

type QueryTransferDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTransferDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTransferDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferDetailsResponse) SetHeaders(v map[string]*string) *QueryTransferDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferDetailsResponse) SetStatusCode(v int32) *QueryTransferDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferDetailsResponse) SetBody(v *QueryTransferDetailsResponseBody) *QueryTransferDetailsResponse {
	s.Body = v
	return s
}

type QueryTransferRecordRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RecordId             *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTransferRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferRecordRequest) SetOwnerId(v int64) *QueryTransferRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTransferRecordRequest) SetProdCode(v string) *QueryTransferRecordRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryTransferRecordRequest) SetRecordId(v int64) *QueryTransferRecordRequest {
	s.RecordId = &v
	return s
}

func (s *QueryTransferRecordRequest) SetResourceOwnerAccount(v string) *QueryTransferRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTransferRecordRequest) SetResourceOwnerId(v int64) *QueryTransferRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryTransferRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryTransferRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferRecordResponseBody) SetRequestId(v string) *QueryTransferRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferRecordResponseBody) SetCode(v string) *QueryTransferRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTransferRecordResponseBody) SetData(v string) *QueryTransferRecordResponseBody {
	s.Data = &v
	return s
}

type QueryTransferRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTransferRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTransferRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferRecordResponse) SetHeaders(v map[string]*string) *QueryTransferRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferRecordResponse) SetStatusCode(v int32) *QueryTransferRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferRecordResponse) SetBody(v *QueryTransferRecordResponseBody) *QueryTransferRecordResponse {
	s.Body = v
	return s
}

type QueryTransferRecordsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	RecordId             *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTransferRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferRecordsRequest) SetOwnerId(v int64) *QueryTransferRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTransferRecordsRequest) SetPageNo(v int32) *QueryTransferRecordsRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTransferRecordsRequest) SetPageSize(v int32) *QueryTransferRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTransferRecordsRequest) SetProdCode(v string) *QueryTransferRecordsRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryTransferRecordsRequest) SetRecordId(v int64) *QueryTransferRecordsRequest {
	s.RecordId = &v
	return s
}

func (s *QueryTransferRecordsRequest) SetResourceOwnerAccount(v string) *QueryTransferRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTransferRecordsRequest) SetResourceOwnerId(v int64) *QueryTransferRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryTransferRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryTransferRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferRecordsResponseBody) SetRequestId(v string) *QueryTransferRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferRecordsResponseBody) SetCode(v string) *QueryTransferRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTransferRecordsResponseBody) SetData(v string) *QueryTransferRecordsResponseBody {
	s.Data = &v
	return s
}

type QueryTransferRecordsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTransferRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTransferRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTransferRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferRecordsResponse) SetHeaders(v map[string]*string) *QueryTransferRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferRecordsResponse) SetStatusCode(v int32) *QueryTransferRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferRecordsResponse) SetBody(v *QueryTransferRecordsResponseBody) *QueryTransferRecordsResponse {
	s.Body = v
	return s
}

type QueryUserDeleteStatusRequest struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	ProdCode       *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryUserDeleteStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDeleteStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryUserDeleteStatusRequest) SetBid(v string) *QueryUserDeleteStatusRequest {
	s.Bid = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetCountry(v string) *QueryUserDeleteStatusRequest {
	s.Country = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetGmtWakeup(v string) *QueryUserDeleteStatusRequest {
	s.GmtWakeup = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetHid(v int64) *QueryUserDeleteStatusRequest {
	s.Hid = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetInterrupt(v bool) *QueryUserDeleteStatusRequest {
	s.Interrupt = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetInvoker(v string) *QueryUserDeleteStatusRequest {
	s.Invoker = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetLevel(v int64) *QueryUserDeleteStatusRequest {
	s.Level = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetMessage(v string) *QueryUserDeleteStatusRequest {
	s.Message = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetPk(v string) *QueryUserDeleteStatusRequest {
	s.Pk = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetProdCode(v string) *QueryUserDeleteStatusRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetPrompt(v string) *QueryUserDeleteStatusRequest {
	s.Prompt = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetSuccess(v bool) *QueryUserDeleteStatusRequest {
	s.Success = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetTaskExtraData(v string) *QueryUserDeleteStatusRequest {
	s.TaskExtraData = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetTaskIdentifier(v string) *QueryUserDeleteStatusRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *QueryUserDeleteStatusRequest) SetUrl(v string) *QueryUserDeleteStatusRequest {
	s.Url = &v
	return s
}

type QueryUserDeleteStatusResponseBody struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryUserDeleteStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDeleteStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserDeleteStatusResponseBody) SetBid(v string) *QueryUserDeleteStatusResponseBody {
	s.Bid = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetCountry(v string) *QueryUserDeleteStatusResponseBody {
	s.Country = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetGmtWakeup(v string) *QueryUserDeleteStatusResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetHid(v int64) *QueryUserDeleteStatusResponseBody {
	s.Hid = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetInterrupt(v bool) *QueryUserDeleteStatusResponseBody {
	s.Interrupt = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetInvoker(v string) *QueryUserDeleteStatusResponseBody {
	s.Invoker = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetLevel(v int64) *QueryUserDeleteStatusResponseBody {
	s.Level = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetMessage(v string) *QueryUserDeleteStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetPk(v string) *QueryUserDeleteStatusResponseBody {
	s.Pk = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetPrompt(v string) *QueryUserDeleteStatusResponseBody {
	s.Prompt = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetSuccess(v bool) *QueryUserDeleteStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetTaskExtraData(v string) *QueryUserDeleteStatusResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetTaskIdentifier(v string) *QueryUserDeleteStatusResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *QueryUserDeleteStatusResponseBody) SetUrl(v string) *QueryUserDeleteStatusResponseBody {
	s.Url = &v
	return s
}

type QueryUserDeleteStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserDeleteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserDeleteStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDeleteStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryUserDeleteStatusResponse) SetHeaders(v map[string]*string) *QueryUserDeleteStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryUserDeleteStatusResponse) SetStatusCode(v int32) *QueryUserDeleteStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserDeleteStatusResponse) SetBody(v *QueryUserDeleteStatusResponseBody) *QueryUserDeleteStatusResponse {
	s.Body = v
	return s
}

type QueryUserInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoRequest) SetOwnerId(v int64) *QueryUserInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryUserInfoRequest) SetProdCode(v string) *QueryUserInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryUserInfoRequest) SetResourceOwnerAccount(v string) *QueryUserInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryUserInfoRequest) SetResourceOwnerId(v int64) *QueryUserInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryUserInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBody) SetRequestId(v string) *QueryUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserInfoResponseBody) SetCode(v string) *QueryUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBody) SetData(v string) *QueryUserInfoResponseBody {
	s.Data = &v
	return s
}

type QueryUserInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponse) SetHeaders(v map[string]*string) *QueryUserInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoResponse) SetStatusCode(v int32) *QueryUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoResponse) SetBody(v *QueryUserInfoResponseBody) *QueryUserInfoResponse {
	s.Body = v
	return s
}

type QueryUserResPoolInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryUserResPoolInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserResPoolInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUserResPoolInfoRequest) SetOwnerId(v int64) *QueryUserResPoolInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryUserResPoolInfoRequest) SetProdCode(v string) *QueryUserResPoolInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryUserResPoolInfoRequest) SetResourceOwnerAccount(v string) *QueryUserResPoolInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryUserResPoolInfoRequest) SetResourceOwnerId(v int64) *QueryUserResPoolInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryUserResPoolInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryUserResPoolInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserResPoolInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserResPoolInfoResponseBody) SetRequestId(v string) *QueryUserResPoolInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserResPoolInfoResponseBody) SetCode(v string) *QueryUserResPoolInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserResPoolInfoResponseBody) SetData(v string) *QueryUserResPoolInfoResponseBody {
	s.Data = &v
	return s
}

type QueryUserResPoolInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserResPoolInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserResPoolInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserResPoolInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserResPoolInfoResponse) SetHeaders(v map[string]*string) *QueryUserResPoolInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserResPoolInfoResponse) SetStatusCode(v int32) *QueryUserResPoolInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserResPoolInfoResponse) SetBody(v *QueryUserResPoolInfoResponseBody) *QueryUserResPoolInfoResponse {
	s.Body = v
	return s
}

type QueryVirtualOperationShowRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryVirtualOperationShowRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualOperationShowRequest) GoString() string {
	return s.String()
}

func (s *QueryVirtualOperationShowRequest) SetOwnerId(v int64) *QueryVirtualOperationShowRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVirtualOperationShowRequest) SetProdCode(v string) *QueryVirtualOperationShowRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryVirtualOperationShowRequest) SetResourceOwnerAccount(v string) *QueryVirtualOperationShowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVirtualOperationShowRequest) SetResourceOwnerId(v int64) *QueryVirtualOperationShowRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryVirtualOperationShowResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVirtualOperationShowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualOperationShowResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVirtualOperationShowResponseBody) SetCode(v string) *QueryVirtualOperationShowResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVirtualOperationShowResponseBody) SetData(v string) *QueryVirtualOperationShowResponseBody {
	s.Data = &v
	return s
}

func (s *QueryVirtualOperationShowResponseBody) SetRequestId(v string) *QueryVirtualOperationShowResponseBody {
	s.RequestId = &v
	return s
}

type QueryVirtualOperationShowResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryVirtualOperationShowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVirtualOperationShowResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualOperationShowResponse) GoString() string {
	return s.String()
}

func (s *QueryVirtualOperationShowResponse) SetHeaders(v map[string]*string) *QueryVirtualOperationShowResponse {
	s.Headers = v
	return s
}

func (s *QueryVirtualOperationShowResponse) SetStatusCode(v int32) *QueryVirtualOperationShowResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVirtualOperationShowResponse) SetBody(v *QueryVirtualOperationShowResponseBody) *QueryVirtualOperationShowResponse {
	s.Body = v
	return s
}

type QueryWarningListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryWarningListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWarningListRequest) GoString() string {
	return s.String()
}

func (s *QueryWarningListRequest) SetOwnerId(v int64) *QueryWarningListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryWarningListRequest) SetProdCode(v string) *QueryWarningListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryWarningListRequest) SetResourceOwnerAccount(v string) *QueryWarningListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryWarningListRequest) SetResourceOwnerId(v int64) *QueryWarningListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryWarningListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryWarningListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWarningListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWarningListResponseBody) SetRequestId(v string) *QueryWarningListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWarningListResponseBody) SetCode(v string) *QueryWarningListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWarningListResponseBody) SetData(v string) *QueryWarningListResponseBody {
	s.Data = &v
	return s
}

type QueryWarningListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWarningListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWarningListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWarningListResponse) GoString() string {
	return s.String()
}

func (s *QueryWarningListResponse) SetHeaders(v map[string]*string) *QueryWarningListResponse {
	s.Headers = v
	return s
}

func (s *QueryWarningListResponse) SetStatusCode(v int32) *QueryWarningListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWarningListResponse) SetBody(v *QueryWarningListResponseBody) *QueryWarningListResponse {
	s.Body = v
	return s
}

type QueryWaybillOrderInfoRequest struct {
	OuterOrderCode       *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryWaybillOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderInfoRequest) SetOuterOrderCode(v string) *QueryWaybillOrderInfoRequest {
	s.OuterOrderCode = &v
	return s
}

func (s *QueryWaybillOrderInfoRequest) SetOwnerId(v int64) *QueryWaybillOrderInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryWaybillOrderInfoRequest) SetResourceOwnerAccount(v string) *QueryWaybillOrderInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryWaybillOrderInfoRequest) SetResourceOwnerId(v int64) *QueryWaybillOrderInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryWaybillOrderInfoResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryWaybillOrderInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryWaybillOrderInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderInfoResponseBody) SetCode(v string) *QueryWaybillOrderInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBody) SetData(v *QueryWaybillOrderInfoResponseBodyData) *QueryWaybillOrderInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryWaybillOrderInfoResponseBody) SetMessage(v string) *QueryWaybillOrderInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBody) SetRequestId(v string) *QueryWaybillOrderInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryWaybillOrderInfoResponseBodyData struct {
	Aliyunprice          *string `json:"Aliyunprice,omitempty" xml:"Aliyunprice,omitempty"`
	AppointGotEndTime    *string `json:"AppointGotEndTime,omitempty" xml:"AppointGotEndTime,omitempty"`
	AppointGotStartTime  *string `json:"AppointGotStartTime,omitempty" xml:"AppointGotStartTime,omitempty"`
	BizType              *int32  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	CpCode               *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	GotCode              *string `json:"GotCode,omitempty" xml:"GotCode,omitempty"`
	LastLogisticDetail   *string `json:"LastLogisticDetail,omitempty" xml:"LastLogisticDetail,omitempty"`
	LogisticsGmtModified *string `json:"LogisticsGmtModified,omitempty" xml:"LogisticsGmtModified,omitempty"`
	LogisticsStatus      *string `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	LogisticsStatusDesc  *string `json:"LogisticsStatusDesc,omitempty" xml:"LogisticsStatusDesc,omitempty"`
	MailNo               *string `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
	OuterOrderCode       *string `json:"OuterOrderCode,omitempty" xml:"OuterOrderCode,omitempty"`
}

func (s QueryWaybillOrderInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetAliyunprice(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.Aliyunprice = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetAppointGotEndTime(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.AppointGotEndTime = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetAppointGotStartTime(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.AppointGotStartTime = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetBizType(v int32) *QueryWaybillOrderInfoResponseBodyData {
	s.BizType = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetCity(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.City = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetCpCode(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.CpCode = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetGotCode(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.GotCode = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetLastLogisticDetail(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.LastLogisticDetail = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetLogisticsGmtModified(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.LogisticsGmtModified = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetLogisticsStatus(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetLogisticsStatusDesc(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.LogisticsStatusDesc = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetMailNo(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.MailNo = &v
	return s
}

func (s *QueryWaybillOrderInfoResponseBodyData) SetOuterOrderCode(v string) *QueryWaybillOrderInfoResponseBodyData {
	s.OuterOrderCode = &v
	return s
}

type QueryWaybillOrderInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWaybillOrderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWaybillOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderInfoResponse) SetHeaders(v map[string]*string) *QueryWaybillOrderInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryWaybillOrderInfoResponse) SetStatusCode(v int32) *QueryWaybillOrderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWaybillOrderInfoResponse) SetBody(v *QueryWaybillOrderInfoResponseBody) *QueryWaybillOrderInfoResponse {
	s.Body = v
	return s
}

type QueryWaybillOrderStatisticsInfoRequest struct {
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Granularity          *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryWaybillOrderStatisticsInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderStatisticsInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderStatisticsInfoRequest) SetEndTime(v string) *QueryWaybillOrderStatisticsInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoRequest) SetGranularity(v string) *QueryWaybillOrderStatisticsInfoRequest {
	s.Granularity = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoRequest) SetOwnerId(v int64) *QueryWaybillOrderStatisticsInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoRequest) SetResourceOwnerAccount(v string) *QueryWaybillOrderStatisticsInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoRequest) SetResourceOwnerId(v int64) *QueryWaybillOrderStatisticsInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoRequest) SetStartTime(v string) *QueryWaybillOrderStatisticsInfoRequest {
	s.StartTime = &v
	return s
}

type QueryWaybillOrderStatisticsInfoResponseBody struct {
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*QueryWaybillOrderStatisticsInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *string                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWaybillOrderStatisticsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderStatisticsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderStatisticsInfoResponseBody) SetCode(v string) *QueryWaybillOrderStatisticsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBody) SetData(v []*QueryWaybillOrderStatisticsInfoResponseBodyData) *QueryWaybillOrderStatisticsInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBody) SetHttpStatusCode(v string) *QueryWaybillOrderStatisticsInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBody) SetMessage(v string) *QueryWaybillOrderStatisticsInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBody) SetRequestId(v string) *QueryWaybillOrderStatisticsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBody) SetSuccess(v bool) *QueryWaybillOrderStatisticsInfoResponseBody {
	s.Success = &v
	return s
}

type QueryWaybillOrderStatisticsInfoResponseBodyData struct {
	AliyunpriceAmount *float64 `json:"AliyunpriceAmount,omitempty" xml:"AliyunpriceAmount,omitempty"`
	CancelCount       *int32   `json:"CancelCount,omitempty" xml:"CancelCount,omitempty"`
	GmtCreate         *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GotCount          *int32   `json:"GotCount,omitempty" xml:"GotCount,omitempty"`
	OrderTotal        *int32   `json:"OrderTotal,omitempty" xml:"OrderTotal,omitempty"`
}

func (s QueryWaybillOrderStatisticsInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderStatisticsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderStatisticsInfoResponseBodyData) SetAliyunpriceAmount(v float64) *QueryWaybillOrderStatisticsInfoResponseBodyData {
	s.AliyunpriceAmount = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBodyData) SetCancelCount(v int32) *QueryWaybillOrderStatisticsInfoResponseBodyData {
	s.CancelCount = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBodyData) SetGmtCreate(v string) *QueryWaybillOrderStatisticsInfoResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBodyData) SetGotCount(v int32) *QueryWaybillOrderStatisticsInfoResponseBodyData {
	s.GotCount = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponseBodyData) SetOrderTotal(v int32) *QueryWaybillOrderStatisticsInfoResponseBodyData {
	s.OrderTotal = &v
	return s
}

type QueryWaybillOrderStatisticsInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryWaybillOrderStatisticsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWaybillOrderStatisticsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillOrderStatisticsInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryWaybillOrderStatisticsInfoResponse) SetHeaders(v map[string]*string) *QueryWaybillOrderStatisticsInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponse) SetStatusCode(v int32) *QueryWaybillOrderStatisticsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWaybillOrderStatisticsInfoResponse) SetBody(v *QueryWaybillOrderStatisticsInfoResponseBody) *QueryWaybillOrderStatisticsInfoResponse {
	s.Body = v
	return s
}

type ReleaseResourceRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	IsDisplayPool        *bool   `json:"IsDisplayPool,omitempty" xml:"IsDisplayPool,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s ReleaseResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseResourceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseResourceRequest) SetBillId(v string) *ReleaseResourceRequest {
	s.BillId = &v
	return s
}

func (s *ReleaseResourceRequest) SetIsDisplayPool(v bool) *ReleaseResourceRequest {
	s.IsDisplayPool = &v
	return s
}

func (s *ReleaseResourceRequest) SetOwnerId(v int64) *ReleaseResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseResourceRequest) SetProdCode(v string) *ReleaseResourceRequest {
	s.ProdCode = &v
	return s
}

func (s *ReleaseResourceRequest) SetResType(v int32) *ReleaseResourceRequest {
	s.ResType = &v
	return s
}

func (s *ReleaseResourceRequest) SetResourceOwnerAccount(v string) *ReleaseResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseResourceRequest) SetResourceOwnerId(v int64) *ReleaseResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseResourceRequest) SetSecretNo(v string) *ReleaseResourceRequest {
	s.SecretNo = &v
	return s
}

type ReleaseResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ReleaseResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseResourceResponseBody) SetRequestId(v string) *ReleaseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseResourceResponseBody) SetCode(v string) *ReleaseResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseResourceResponseBody) SetData(v string) *ReleaseResourceResponseBody {
	s.Data = &v
	return s
}

type ReleaseResourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseResourceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseResourceResponse) SetHeaders(v map[string]*string) *ReleaseResourceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseResourceResponse) SetStatusCode(v int32) *ReleaseResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseResourceResponse) SetBody(v *ReleaseResourceResponseBody) *ReleaseResourceResponse {
	s.Body = v
	return s
}

type TestTtsRingToneRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tts                  *string `json:"Tts,omitempty" xml:"Tts,omitempty"`
	VoiceSpeed           *string `json:"VoiceSpeed,omitempty" xml:"VoiceSpeed,omitempty"`
	VoiceStyle           *string `json:"VoiceStyle,omitempty" xml:"VoiceStyle,omitempty"`
	VoiceType            *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
	VoiceVolume          *string `json:"VoiceVolume,omitempty" xml:"VoiceVolume,omitempty"`
}

func (s TestTtsRingToneRequest) String() string {
	return tea.Prettify(s)
}

func (s TestTtsRingToneRequest) GoString() string {
	return s.String()
}

func (s *TestTtsRingToneRequest) SetBillId(v string) *TestTtsRingToneRequest {
	s.BillId = &v
	return s
}

func (s *TestTtsRingToneRequest) SetOwnerId(v int64) *TestTtsRingToneRequest {
	s.OwnerId = &v
	return s
}

func (s *TestTtsRingToneRequest) SetProdCode(v string) *TestTtsRingToneRequest {
	s.ProdCode = &v
	return s
}

func (s *TestTtsRingToneRequest) SetResourceOwnerAccount(v string) *TestTtsRingToneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TestTtsRingToneRequest) SetResourceOwnerId(v int64) *TestTtsRingToneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TestTtsRingToneRequest) SetTts(v string) *TestTtsRingToneRequest {
	s.Tts = &v
	return s
}

func (s *TestTtsRingToneRequest) SetVoiceSpeed(v string) *TestTtsRingToneRequest {
	s.VoiceSpeed = &v
	return s
}

func (s *TestTtsRingToneRequest) SetVoiceStyle(v string) *TestTtsRingToneRequest {
	s.VoiceStyle = &v
	return s
}

func (s *TestTtsRingToneRequest) SetVoiceType(v string) *TestTtsRingToneRequest {
	s.VoiceType = &v
	return s
}

func (s *TestTtsRingToneRequest) SetVoiceVolume(v string) *TestTtsRingToneRequest {
	s.VoiceVolume = &v
	return s
}

type TestTtsRingToneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TestTtsRingToneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TestTtsRingToneResponseBody) GoString() string {
	return s.String()
}

func (s *TestTtsRingToneResponseBody) SetRequestId(v string) *TestTtsRingToneResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestTtsRingToneResponseBody) SetCode(v string) *TestTtsRingToneResponseBody {
	s.Code = &v
	return s
}

func (s *TestTtsRingToneResponseBody) SetData(v string) *TestTtsRingToneResponseBody {
	s.Data = &v
	return s
}

type TestTtsRingToneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TestTtsRingToneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TestTtsRingToneResponse) String() string {
	return tea.Prettify(s)
}

func (s TestTtsRingToneResponse) GoString() string {
	return s.String()
}

func (s *TestTtsRingToneResponse) SetHeaders(v map[string]*string) *TestTtsRingToneResponse {
	s.Headers = v
	return s
}

func (s *TestTtsRingToneResponse) SetStatusCode(v int32) *TestTtsRingToneResponse {
	s.StatusCode = &v
	return s
}

func (s *TestTtsRingToneResponse) SetBody(v *TestTtsRingToneResponseBody) *TestTtsRingToneResponse {
	s.Body = v
	return s
}

type UnbindResourceRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	BindIds              *string `json:"BindIds,omitempty" xml:"BindIds,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UnbindResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindResourceRequest) GoString() string {
	return s.String()
}

func (s *UnbindResourceRequest) SetBillId(v string) *UnbindResourceRequest {
	s.BillId = &v
	return s
}

func (s *UnbindResourceRequest) SetBindIds(v string) *UnbindResourceRequest {
	s.BindIds = &v
	return s
}

func (s *UnbindResourceRequest) SetOwnerId(v int64) *UnbindResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindResourceRequest) SetProdCode(v string) *UnbindResourceRequest {
	s.ProdCode = &v
	return s
}

func (s *UnbindResourceRequest) SetResType(v int32) *UnbindResourceRequest {
	s.ResType = &v
	return s
}

func (s *UnbindResourceRequest) SetResourceOwnerAccount(v string) *UnbindResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindResourceRequest) SetResourceOwnerId(v int64) *UnbindResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindResourceRequest) SetSecretNo(v string) *UnbindResourceRequest {
	s.SecretNo = &v
	return s
}

type UnbindResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UnbindResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResourceResponseBody) SetRequestId(v string) *UnbindResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindResourceResponseBody) SetCode(v string) *UnbindResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindResourceResponseBody) SetData(v string) *UnbindResourceResponseBody {
	s.Data = &v
	return s
}

type UnbindResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindResourceResponse) GoString() string {
	return s.String()
}

func (s *UnbindResourceResponse) SetHeaders(v map[string]*string) *UnbindResourceResponse {
	s.Headers = v
	return s
}

func (s *UnbindResourceResponse) SetStatusCode(v int32) *UnbindResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResourceResponse) SetBody(v *UnbindResourceResponseBody) *UnbindResourceResponse {
	s.Body = v
	return s
}

type UnlockResourceRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UnlockResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockResourceRequest) GoString() string {
	return s.String()
}

func (s *UnlockResourceRequest) SetBillId(v string) *UnlockResourceRequest {
	s.BillId = &v
	return s
}

func (s *UnlockResourceRequest) SetOwnerId(v int64) *UnlockResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockResourceRequest) SetProdCode(v string) *UnlockResourceRequest {
	s.ProdCode = &v
	return s
}

func (s *UnlockResourceRequest) SetResourceOwnerAccount(v string) *UnlockResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockResourceRequest) SetResourceOwnerId(v int64) *UnlockResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockResourceRequest) SetSecretNo(v string) *UnlockResourceRequest {
	s.SecretNo = &v
	return s
}

type UnlockResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UnlockResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockResourceResponseBody) SetRequestId(v string) *UnlockResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockResourceResponseBody) SetCode(v string) *UnlockResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UnlockResourceResponseBody) SetData(v string) *UnlockResourceResponseBody {
	s.Data = &v
	return s
}

type UnlockResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnlockResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockResourceResponse) GoString() string {
	return s.String()
}

func (s *UnlockResourceResponse) SetHeaders(v map[string]*string) *UnlockResourceResponse {
	s.Headers = v
	return s
}

func (s *UnlockResourceResponse) SetStatusCode(v int32) *UnlockResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockResourceResponse) SetBody(v *UnlockResourceResponseBody) *UnlockResourceResponse {
	s.Body = v
	return s
}

type UpdateContactsRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactsRequest) GoString() string {
	return s.String()
}

func (s *UpdateContactsRequest) SetBillId(v string) *UpdateContactsRequest {
	s.BillId = &v
	return s
}

func (s *UpdateContactsRequest) SetId(v int64) *UpdateContactsRequest {
	s.Id = &v
	return s
}

func (s *UpdateContactsRequest) SetName(v string) *UpdateContactsRequest {
	s.Name = &v
	return s
}

func (s *UpdateContactsRequest) SetOwnerId(v int64) *UpdateContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContactsRequest) SetPhoneNumber(v string) *UpdateContactsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateContactsRequest) SetProdCode(v string) *UpdateContactsRequest {
	s.ProdCode = &v
	return s
}

func (s *UpdateContactsRequest) SetResourceOwnerAccount(v string) *UpdateContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContactsRequest) SetResourceOwnerId(v int64) *UpdateContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateContactsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactsResponseBody) SetRequestId(v string) *UpdateContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactsResponseBody) SetCode(v string) *UpdateContactsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContactsResponseBody) SetData(v string) *UpdateContactsResponseBody {
	s.Data = &v
	return s
}

type UpdateContactsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContactsResponse) GoString() string {
	return s.String()
}

func (s *UpdateContactsResponse) SetHeaders(v map[string]*string) *UpdateContactsResponse {
	s.Headers = v
	return s
}

func (s *UpdateContactsResponse) SetStatusCode(v int32) *UpdateContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContactsResponse) SetBody(v *UpdateContactsResponseBody) *UpdateContactsResponse {
	s.Body = v
	return s
}

type UpdateGroupDetailRequest struct {
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDetailRequest) SetGroupId(v string) *UpdateGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetId(v string) *UpdateGroupDetailRequest {
	s.Id = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetOwnerId(v int64) *UpdateGroupDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetPoolKey(v string) *UpdateGroupDetailRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetProdCode(v string) *UpdateGroupDetailRequest {
	s.ProdCode = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetRemark(v string) *UpdateGroupDetailRequest {
	s.Remark = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetResourceOwnerAccount(v string) *UpdateGroupDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateGroupDetailRequest) SetResourceOwnerId(v int64) *UpdateGroupDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateGroupDetailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDetailResponseBody) SetRequestId(v string) *UpdateGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupDetailResponseBody) SetCode(v string) *UpdateGroupDetailResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupDetailResponseBody) SetData(v bool) *UpdateGroupDetailResponseBody {
	s.Data = &v
	return s
}

type UpdateGroupDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDetailResponse) SetHeaders(v map[string]*string) *UpdateGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDetailResponse) SetStatusCode(v int32) *UpdateGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupDetailResponse) SetBody(v *UpdateGroupDetailResponseBody) *UpdateGroupDetailResponse {
	s.Body = v
	return s
}

type UpdateGroupInfoRequest struct {
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupInfoRequest) SetId(v string) *UpdateGroupInfoRequest {
	s.Id = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetName(v string) *UpdateGroupInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetOwnerId(v int64) *UpdateGroupInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetPoolKey(v string) *UpdateGroupInfoRequest {
	s.PoolKey = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetProdCode(v string) *UpdateGroupInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetRemark(v string) *UpdateGroupInfoRequest {
	s.Remark = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetResourceOwnerAccount(v string) *UpdateGroupInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateGroupInfoRequest) SetResourceOwnerId(v int64) *UpdateGroupInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateGroupInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupInfoResponseBody) SetRequestId(v string) *UpdateGroupInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupInfoResponseBody) SetCode(v string) *UpdateGroupInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupInfoResponseBody) SetData(v bool) *UpdateGroupInfoResponseBody {
	s.Data = &v
	return s
}

type UpdateGroupInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupInfoResponse) SetHeaders(v map[string]*string) *UpdateGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupInfoResponse) SetStatusCode(v int32) *UpdateGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupInfoResponse) SetBody(v *UpdateGroupInfoResponseBody) *UpdateGroupInfoResponse {
	s.Body = v
	return s
}

type UpdatePoolNameRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdatePoolNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolNameRequest) GoString() string {
	return s.String()
}

func (s *UpdatePoolNameRequest) SetBillId(v string) *UpdatePoolNameRequest {
	s.BillId = &v
	return s
}

func (s *UpdatePoolNameRequest) SetOwnerId(v int64) *UpdatePoolNameRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePoolNameRequest) SetPoolName(v string) *UpdatePoolNameRequest {
	s.PoolName = &v
	return s
}

func (s *UpdatePoolNameRequest) SetProdCode(v string) *UpdatePoolNameRequest {
	s.ProdCode = &v
	return s
}

func (s *UpdatePoolNameRequest) SetResType(v int32) *UpdatePoolNameRequest {
	s.ResType = &v
	return s
}

func (s *UpdatePoolNameRequest) SetResourceOwnerAccount(v string) *UpdatePoolNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePoolNameRequest) SetResourceOwnerId(v int64) *UpdatePoolNameRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdatePoolNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdatePoolNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePoolNameResponseBody) SetRequestId(v string) *UpdatePoolNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePoolNameResponseBody) SetCode(v string) *UpdatePoolNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePoolNameResponseBody) SetData(v string) *UpdatePoolNameResponseBody {
	s.Data = &v
	return s
}

type UpdatePoolNameResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePoolNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePoolNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePoolNameResponse) GoString() string {
	return s.String()
}

func (s *UpdatePoolNameResponse) SetHeaders(v map[string]*string) *UpdatePoolNameResponse {
	s.Headers = v
	return s
}

func (s *UpdatePoolNameResponse) SetStatusCode(v int32) *UpdatePoolNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePoolNameResponse) SetBody(v *UpdatePoolNameResponseBody) *UpdatePoolNameResponse {
	s.Body = v
	return s
}

type UpdateResRemarkRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResType              *int32  `json:"ResType,omitempty" xml:"ResType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecretNo             *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
}

func (s UpdateResRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResRemarkRequest) SetBillId(v string) *UpdateResRemarkRequest {
	s.BillId = &v
	return s
}

func (s *UpdateResRemarkRequest) SetOwnerId(v int64) *UpdateResRemarkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateResRemarkRequest) SetProdCode(v string) *UpdateResRemarkRequest {
	s.ProdCode = &v
	return s
}

func (s *UpdateResRemarkRequest) SetRemark(v string) *UpdateResRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateResRemarkRequest) SetResType(v int32) *UpdateResRemarkRequest {
	s.ResType = &v
	return s
}

func (s *UpdateResRemarkRequest) SetResourceOwnerAccount(v string) *UpdateResRemarkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateResRemarkRequest) SetResourceOwnerId(v int64) *UpdateResRemarkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateResRemarkRequest) SetSecretNo(v string) *UpdateResRemarkRequest {
	s.SecretNo = &v
	return s
}

type UpdateResRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateResRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResRemarkResponseBody) SetRequestId(v string) *UpdateResRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResRemarkResponseBody) SetCode(v string) *UpdateResRemarkResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateResRemarkResponseBody) SetData(v string) *UpdateResRemarkResponseBody {
	s.Data = &v
	return s
}

type UpdateResRemarkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateResRemarkResponse) SetHeaders(v map[string]*string) *UpdateResRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateResRemarkResponse) SetStatusCode(v int32) *UpdateResRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResRemarkResponse) SetBody(v *UpdateResRemarkResponseBody) *UpdateResRemarkResponse {
	s.Body = v
	return s
}

type ValidateOrderRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Data                 *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ValidateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderRequest) GoString() string {
	return s.String()
}

func (s *ValidateOrderRequest) SetOwnerId(v int64) *ValidateOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *ValidateOrderRequest) SetProdCode(v string) *ValidateOrderRequest {
	s.ProdCode = &v
	return s
}

func (s *ValidateOrderRequest) SetResourceOwnerAccount(v string) *ValidateOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ValidateOrderRequest) SetResourceOwnerId(v int64) *ValidateOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ValidateOrderRequest) SetData(v string) *ValidateOrderRequest {
	s.Data = &v
	return s
}

type ValidateOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ValidateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateOrderResponseBody) SetCode(v string) *ValidateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateOrderResponseBody) SetData(v string) *ValidateOrderResponseBody {
	s.Data = &v
	return s
}

func (s *ValidateOrderResponseBody) SetMessage(v string) *ValidateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateOrderResponseBody) SetRequestId(v string) *ValidateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateOrderResponseBody) SetSuccess(v string) *ValidateOrderResponseBody {
	s.Success = &v
	return s
}

type ValidateOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderResponse) GoString() string {
	return s.String()
}

func (s *ValidateOrderResponse) SetHeaders(v map[string]*string) *ValidateOrderResponse {
	s.Headers = v
	return s
}

func (s *ValidateOrderResponse) SetStatusCode(v int32) *ValidateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateOrderResponse) SetBody(v *ValidateOrderResponseBody) *ValidateOrderResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dypls"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyArInvoiceWithSourceWithOptions(request *ApplyArInvoiceWithSourceRequest, runtime *util.RuntimeOptions) (_result *ApplyArInvoiceWithSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		bodyFlat["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		body["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.Applier)) {
		body["Applier"] = request.Applier
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyDate)) {
		body["ApplyDate"] = request.ApplyDate
	}

	if !tea.BoolValue(util.IsUnset(request.CurrencyCode)) {
		body["CurrencyCode"] = request.CurrencyCode
	}

	if !tea.BoolValue(util.IsUnset(request.Customer)) {
		bodyFlat["Customer"] = request.Customer
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptProps)) {
		bodyFlat["EncryptProps"] = request.EncryptProps
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludingTaxAmount)) {
		body["ExcludingTaxAmount"] = request.ExcludingTaxAmount
	}

	if !tea.BoolValue(util.IsUnset(request.InputType)) {
		body["InputType"] = request.InputType
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceType)) {
		body["InvoiceType"] = request.InvoiceType
	}

	if !tea.BoolValue(util.IsUnset(request.IsMerged)) {
		body["IsMerged"] = request.IsMerged
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialType)) {
		body["MaterialType"] = request.MaterialType
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		body["Memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.OuCode)) {
		body["OuCode"] = request.OuCode
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserBankInfo)) {
		body["PurchaserBankInfo"] = request.PurchaserBankInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserContactInfo)) {
		body["PurchaserContactInfo"] = request.PurchaserContactInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserName)) {
		body["PurchaserName"] = request.PurchaserName
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserTaxNo)) {
		body["PurchaserTaxNo"] = request.PurchaserTaxNo
	}

	if !tea.BoolValue(util.IsUnset(request.RequestNo)) {
		body["RequestNo"] = request.RequestNo
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		body["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.SiteId)) {
		body["SiteId"] = request.SiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceList)) {
		body["SourceList"] = request.SourceList
	}

	if !tea.BoolValue(util.IsUnset(request.TaxAmount)) {
		body["TaxAmount"] = request.TaxAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyArInvoiceWithSource"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyArInvoiceWithSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyArInvoiceWithSource(request *ApplyArInvoiceWithSourceRequest) (_result *ApplyArInvoiceWithSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyArInvoiceWithSourceResponse{}
	_body, _err := client.ApplyArInvoiceWithSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyBlackInfoExportWithOptions(request *ApplyBlackInfoExportRequest, runtime *util.RuntimeOptions) (_result *ApplyBlackInfoExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyBlackInfoExport"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyBlackInfoExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyBlackInfoExport(request *ApplyBlackInfoExportRequest) (_result *ApplyBlackInfoExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyBlackInfoExportResponse{}
	_body, _err := client.ApplyBlackInfoExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyCallRecordExportWithOptions(request *ApplyCallRecordExportRequest, runtime *util.RuntimeOptions) (_result *ApplyCallRecordExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.CallDate)) {
		query["CallDate"] = request.CallDate
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyCallRecordExport"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyCallRecordExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyCallRecordExport(request *ApplyCallRecordExportRequest) (_result *ApplyCallRecordExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyCallRecordExportResponse{}
	_body, _err := client.ApplyCallRecordExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyGroupNumberExportWithOptions(request *ApplyGroupNumberExportRequest, runtime *util.RuntimeOptions) (_result *ApplyGroupNumberExportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyGroupNumberExport"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyGroupNumberExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyGroupNumberExport(request *ApplyGroupNumberExportRequest) (_result *ApplyGroupNumberExportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyGroupNumberExportResponse{}
	_body, _err := client.ApplyGroupNumberExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyRingToneWithOptions(request *ApplyRingToneRequest, runtime *util.RuntimeOptions) (_result *ApplyRingToneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayType)) {
		query["PlayType"] = request.PlayType
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyRingTone"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyRingToneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyRingTone(request *ApplyRingToneRequest) (_result *ApplyRingToneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyRingToneResponse{}
	_body, _err := client.ApplyRingToneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchOccupySecretResWithOptions(tmpReq *BatchOccupySecretResRequest, runtime *util.RuntimeOptions) (_result *BatchOccupySecretResResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchOccupySecretResShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BatchOccupyList)) {
		request.BatchOccupyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BatchOccupyList, tea.String("BatchOccupyList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchOccupyListShrink)) {
		query["BatchOccupyList"] = request.BatchOccupyListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchOccupySecretRes"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchOccupySecretResResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchOccupySecretRes(request *BatchOccupySecretResRequest) (_result *BatchOccupySecretResResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchOccupySecretResResponse{}
	_body, _err := client.BatchOccupySecretResWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindResourceWithOptions(request *BindResourceRequest, runtime *util.RuntimeOptions) (_result *BindResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsrModelId)) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !tea.BoolValue(util.IsUnset(request.AsrStatus)) {
		query["AsrStatus"] = request.AsrStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AxnExtensionB)) {
		query["AxnExtensionB"] = request.AxnExtensionB
	}

	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpTime)) {
		query["ExpTime"] = request.ExpTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecord)) {
		query["IsRecord"] = request.IsRecord
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindResource"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindResource(request *BindResourceRequest) (_result *BindResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindResourceResponse{}
	_body, _err := client.BindResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BlackOperateWithOptions(request *BlackOperateRequest, runtime *util.RuntimeOptions) (_result *BlackOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.BlackMap)) {
		query["BlackMap"] = request.BlackMap
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		query["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BlackOperate"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BlackOperateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BlackOperate(request *BlackOperateRequest) (_result *BlackOperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BlackOperateResponse{}
	_body, _err := client.BlackOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCertifyInfoWithOptions(request *CreateCertifyInfoRequest, runtime *util.RuntimeOptions) (_result *CreateCertifyInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNo)) {
		query["PhoneNo"] = request.PhoneNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCertifyInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCertifyInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertifyInfo(request *CreateCertifyInfoRequest) (_result *CreateCertifyInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertifyInfoResponse{}
	_body, _err := client.CreateCertifyInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateContactsWithOptions(request *CreateContactsRequest, runtime *util.RuntimeOptions) (_result *CreateContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateContacts"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateContacts(request *CreateContactsRequest) (_result *CreateContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateContactsResponse{}
	_body, _err := client.CreateContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupDetailWithOptions(request *CreateGroupDetailRequest, runtime *util.RuntimeOptions) (_result *CreateGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NumberList)) {
		query["NumberList"] = request.NumberList
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupDetail"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupDetail(request *CreateGroupDetailRequest) (_result *CreateGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupDetailResponse{}
	_body, _err := client.CreateGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupInfoWithOptions(request *CreateGroupInfoRequest, runtime *util.RuntimeOptions) (_result *CreateGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NumberList)) {
		query["NumberList"] = request.NumberList
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupInfo(request *CreateGroupInfoRequest) (_result *CreateGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupInfoResponse{}
	_body, _err := client.CreateGroupInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogicalDeleteWithOptions(request *CreateLogicalDeleteRequest, runtime *util.RuntimeOptions) (_result *CreateLogicalDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.GmtWakeup)) {
		query["GmtWakeup"] = request.GmtWakeup
	}

	if !tea.BoolValue(util.IsUnset(request.Hid)) {
		query["Hid"] = request.Hid
	}

	if !tea.BoolValue(util.IsUnset(request.Interrupt)) {
		query["Interrupt"] = request.Interrupt
	}

	if !tea.BoolValue(util.IsUnset(request.Invoker)) {
		query["Invoker"] = request.Invoker
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.Success)) {
		query["Success"] = request.Success
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExtraData)) {
		query["TaskExtraData"] = request.TaskExtraData
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdentifier)) {
		query["TaskIdentifier"] = request.TaskIdentifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogicalDelete"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLogicalDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogicalDelete(request *CreateLogicalDeleteRequest) (_result *CreateLogicalDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLogicalDeleteResponse{}
	_body, _err := client.CreateLogicalDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMessageCallbackWithOptions(request *CreateMessageCallbackRequest, runtime *util.RuntimeOptions) (_result *CreateMessageCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMessageCallback"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMessageCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMessageCallback(request *CreateMessageCallbackRequest) (_result *CreateMessageCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMessageCallbackResponse{}
	_body, _err := client.CreateMessageCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMessageQueueWithOptions(request *CreateMessageQueueRequest, runtime *util.RuntimeOptions) (_result *CreateMessageQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillIds)) {
		query["BillIds"] = request.BillIds
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.QueueTitle)) {
		query["QueueTitle"] = request.QueueTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMessageQueue"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMessageQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMessageQueue(request *CreateMessageQueueRequest) (_result *CreateMessageQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMessageQueueResponse{}
	_body, _err := client.CreateMessageQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePhysicalDeleteWithOptions(request *CreatePhysicalDeleteRequest, runtime *util.RuntimeOptions) (_result *CreatePhysicalDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.GmtWakeup)) {
		query["GmtWakeup"] = request.GmtWakeup
	}

	if !tea.BoolValue(util.IsUnset(request.Hid)) {
		query["Hid"] = request.Hid
	}

	if !tea.BoolValue(util.IsUnset(request.Interrupt)) {
		query["Interrupt"] = request.Interrupt
	}

	if !tea.BoolValue(util.IsUnset(request.Invoker)) {
		query["Invoker"] = request.Invoker
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.Success)) {
		query["Success"] = request.Success
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExtraData)) {
		query["TaskExtraData"] = request.TaskExtraData
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdentifier)) {
		query["TaskIdentifier"] = request.TaskIdentifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePhysicalDelete"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePhysicalDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePhysicalDelete(request *CreatePhysicalDeleteRequest) (_result *CreatePhysicalDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePhysicalDeleteResponse{}
	_body, _err := client.CreatePhysicalDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePoolInfoWithOptions(request *CreatePoolInfoRequest, runtime *util.RuntimeOptions) (_result *CreatePoolInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePoolInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePoolInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePoolInfo(request *CreatePoolInfoRequest) (_result *CreatePoolInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePoolInfoResponse{}
	_body, _err := client.CreatePoolInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductWithOptions(request *CreateProductRequest, runtime *util.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProdId)) {
		query["ProdId"] = request.ProdId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProduct"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRingToneWithOptions(request *CreateRingToneRequest, runtime *util.RuntimeOptions) (_result *CreateRingToneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		query["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayType)) {
		query["PlayType"] = request.PlayType
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RingName)) {
		query["RingName"] = request.RingName
	}

	if !tea.BoolValue(util.IsUnset(request.Tts)) {
		query["Tts"] = request.Tts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRingTone"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRingToneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRingTone(request *CreateRingToneRequest) (_result *CreateRingToneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRingToneResponse{}
	_body, _err := client.CreateRingToneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubsTrialWithOptions(request *CreateSubsTrialRequest, runtime *util.RuntimeOptions) (_result *CreateSubsTrialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneA)) {
		query["PhoneA"] = request.PhoneA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneB)) {
		query["PhoneB"] = request.PhoneB
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubsTrial"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubsTrialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubsTrial(request *CreateSubsTrialRequest) (_result *CreateSubsTrialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubsTrialResponse{}
	_body, _err := client.CreateSubsTrialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransferRecordWithOptions(request *CreateTransferRecordRequest, runtime *util.RuntimeOptions) (_result *CreateTransferRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.NumberList)) {
		query["NumberList"] = request.NumberList
	}

	if !tea.BoolValue(util.IsUnset(request.OriginBillId)) {
		query["OriginBillId"] = request.OriginBillId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginName)) {
		query["OriginName"] = request.OriginName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetBillId)) {
		query["TargetBillId"] = request.TargetBillId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetName)) {
		query["TargetName"] = request.TargetName
	}

	if !tea.BoolValue(util.IsUnset(request.Total)) {
		query["Total"] = request.Total
	}

	if !tea.BoolValue(util.IsUnset(request.TransferType)) {
		query["TransferType"] = request.TransferType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransferRecord"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransferRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransferRecord(request *CreateTransferRecordRequest) (_result *CreateTransferRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransferRecordResponse{}
	_body, _err := client.CreateTransferRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCertifyInfoWithOptions(request *DeleteCertifyInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteCertifyInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyId)) {
		query["CertifyId"] = request.CertifyId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCertifyInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCertifyInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCertifyInfo(request *DeleteCertifyInfoRequest) (_result *DeleteCertifyInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCertifyInfoResponse{}
	_body, _err := client.DeleteCertifyInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactsWithOptions(request *DeleteContactsRequest, runtime *util.RuntimeOptions) (_result *DeleteContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContacts"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContacts(request *DeleteContactsRequest) (_result *DeleteContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactsResponse{}
	_body, _err := client.DeleteContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupDetailWithOptions(request *DeleteGroupDetailRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IdList)) {
		query["IdList"] = request.IdList
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupDetail"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupDetail(request *DeleteGroupDetailRequest) (_result *DeleteGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupDetailResponse{}
	_body, _err := client.DeleteGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMessageCallbackWithOptions(request *DeleteMessageCallbackRequest, runtime *util.RuntimeOptions) (_result *DeleteMessageCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMessageCallback"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMessageCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMessageCallback(request *DeleteMessageCallbackRequest) (_result *DeleteMessageCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMessageCallbackResponse{}
	_body, _err := client.DeleteMessageCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRingToneWithOptions(request *DeleteRingToneRequest, runtime *util.RuntimeOptions) (_result *DeleteRingToneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRingTone"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRingToneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRingTone(request *DeleteRingToneRequest) (_result *DeleteRingToneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRingToneResponse{}
	_body, _err := client.DeleteRingToneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadCompleteWithOptions(request *DownloadCompleteRequest, runtime *util.RuntimeOptions) (_result *DownloadCompleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadComplete"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DownloadCompleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadComplete(request *DownloadCompleteRequest) (_result *DownloadCompleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadCompleteResponse{}
	_body, _err := client.DownloadCompleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportResWithOptions(request *ExportResRequest, runtime *util.RuntimeOptions) (_result *ExportResResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResBindStatus)) {
		query["ResBindStatus"] = request.ResBindStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportRes"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportResResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportRes(request *ExportResRequest) (_result *ExportResResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportResResponse{}
	_body, _err := client.ExportResWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEinvoicePdfDataWithOptions(request *GetEinvoicePdfDataRequest, runtime *util.RuntimeOptions) (_result *GetEinvoicePdfDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		body["AppCode"] = request.AppCode
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Customer)) {
		bodyFlat["Customer"] = request.Customer
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptProps)) {
		bodyFlat["EncryptProps"] = request.EncryptProps
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		body["InvoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		body["InvoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		body["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEinvoicePdfData"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEinvoicePdfDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEinvoicePdfData(request *GetEinvoicePdfDataRequest) (_result *GetEinvoicePdfDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEinvoicePdfDataResponse{}
	_body, _err := client.GetEinvoicePdfDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretAsrInfoWithOptions(request *GetSecretAsrInfoRequest, runtime *util.RuntimeOptions) (_result *GetSecretAsrInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CallTime)) {
		query["CallTime"] = request.CallTime
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretAsrInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecretAsrInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecretAsrInfo(request *GetSecretAsrInfoRequest) (_result *GetSecretAsrInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretAsrInfoResponse{}
	_body, _err := client.GetSecretAsrInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserResourceTagStatusWithOptions(request *GetUserResourceTagStatusRequest, runtime *util.RuntimeOptions) (_result *GetUserResourceTagStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserResourceTagStatus"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResourceTagStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserResourceTagStatus(request *GetUserResourceTagStatusRequest) (_result *GetUserResourceTagStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResourceTagStatusResponse{}
	_body, _err := client.GetUserResourceTagStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsrLanguageModelsWithOptions(request *ListAsrLanguageModelsRequest, runtime *util.RuntimeOptions) (_result *ListAsrLanguageModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsrLanguageModels"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsrLanguageModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsrLanguageModels(request *ListAsrLanguageModelsRequest) (_result *ListAsrLanguageModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsrLanguageModelsResponse{}
	_body, _err := client.ListAsrLanguageModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockResourceWithOptions(request *LockResourceRequest, runtime *util.RuntimeOptions) (_result *LockResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockResource"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockResource(request *LockResourceRequest) (_result *LockResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockResourceResponse{}
	_body, _err := client.LockResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OccupySecretResWithOptions(request *OccupySecretResRequest, runtime *util.RuntimeOptions) (_result *OccupySecretResResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.IsDisplayPool)) {
		query["IsDisplayPool"] = request.IsDisplayPool
	}

	if !tea.BoolValue(util.IsUnset(request.NoLike)) {
		query["NoLike"] = request.NoLike
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDetailId)) {
		query["OrderDetailId"] = request.OrderDetailId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerKey)) {
		query["PartnerKey"] = request.PartnerKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNoType)) {
		query["SecretNoType"] = request.SecretNoType
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["secretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OccupySecretRes"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OccupySecretResResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OccupySecretRes(request *OccupySecretResRequest) (_result *OccupySecretResResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OccupySecretResResponse{}
	_body, _err := client.OccupySecretResWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OrderSucceededCallbackWithOptions(request *OrderSucceededCallbackRequest, runtime *util.RuntimeOptions) (_result *OrderSucceededCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OrderSucceededCallback"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OrderSucceededCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OrderSucceededCallback(request *OrderSucceededCallbackRequest) (_result *OrderSucceededCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OrderSucceededCallbackResponse{}
	_body, _err := client.OrderSucceededCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PoolConfigWithOptions(request *PoolConfigRequest, runtime *util.RuntimeOptions) (_result *PoolConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackType)) {
		query["CallbackType"] = request.CallbackType
	}

	if !tea.BoolValue(util.IsUnset(request.FrozenDay)) {
		query["FrozenDay"] = request.FrozenDay
	}

	if !tea.BoolValue(util.IsUnset(request.NeedAllCallRecords)) {
		query["NeedAllCallRecords"] = request.NeedAllCallRecords
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSmsWhite)) {
		query["OpenSmsWhite"] = request.OpenSmsWhite
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolWarningLimit)) {
		query["PoolWarningLimit"] = request.PoolWarningLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SelectXMode)) {
		query["SelectXMode"] = request.SelectXMode
	}

	if !tea.BoolValue(util.IsUnset(request.SmartSmsWhitelist)) {
		query["SmartSmsWhitelist"] = request.SmartSmsWhitelist
	}

	if !tea.BoolValue(util.IsUnset(request.SmsChannel)) {
		query["SmsChannel"] = request.SmsChannel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PoolConfig"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PoolConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PoolConfig(request *PoolConfigRequest) (_result *PoolConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PoolConfigResponse{}
	_body, _err := client.PoolConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PurchaseResourcesWithOptions(request *PurchaseResourcesRequest, runtime *util.RuntimeOptions) (_result *PurchaseResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.BuyNumber)) {
		query["BuyNumber"] = request.BuyNumber
	}

	if !tea.BoolValue(util.IsUnset(request.IsDisplayPool)) {
		query["IsDisplayPool"] = request.IsDisplayPool
	}

	if !tea.BoolValue(util.IsUnset(request.NoLike)) {
		query["NoLike"] = request.NoLike
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionName)) {
		query["RegionName"] = request.RegionName
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	if !tea.BoolValue(util.IsUnset(request.UsageScenarios)) {
		query["UsageScenarios"] = request.UsageScenarios
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PurchaseResources"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PurchaseResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PurchaseResources(request *PurchaseResourcesRequest) (_result *PurchaseResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PurchaseResourcesResponse{}
	_body, _err := client.PurchaseResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBindingDetailsWithOptions(request *QueryBindingDetailsRequest, runtime *util.RuntimeOptions) (_result *QueryBindingDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	if !tea.BoolValue(util.IsUnset(request.SubId)) {
		query["SubId"] = request.SubId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBindingDetails"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBindingDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBindingDetails(request *QueryBindingDetailsRequest) (_result *QueryBindingDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBindingDetailsResponse{}
	_body, _err := client.QueryBindingDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBlackListWithOptions(request *QueryBlackListRequest, runtime *util.RuntimeOptions) (_result *QueryBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.BlackPrefix)) {
		query["BlackPrefix"] = request.BlackPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBlackList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBlackList(request *QueryBlackListRequest) (_result *QueryBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBlackListResponse{}
	_body, _err := client.QueryBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBuyPageInitDataWithOptions(request *QueryBuyPageInitDataRequest, runtime *util.RuntimeOptions) (_result *QueryBuyPageInitDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBuyPageInitData"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBuyPageInitDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBuyPageInitData(request *QueryBuyPageInitDataRequest) (_result *QueryBuyPageInitDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBuyPageInitDataResponse{}
	_body, _err := client.QueryBuyPageInitDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBuyPageResCountWithOptions(request *QueryBuyPageResCountRequest, runtime *util.RuntimeOptions) (_result *QueryBuyPageResCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Like)) {
		query["Like"] = request.Like
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBuyPageResCount"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBuyPageResCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBuyPageResCount(request *QueryBuyPageResCountRequest) (_result *QueryBuyPageResCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBuyPageResCountResponse{}
	_body, _err := client.QueryBuyPageResCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBuyPageResInfoWithOptions(request *QueryBuyPageResInfoRequest, runtime *util.RuntimeOptions) (_result *QueryBuyPageResInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Like)) {
		query["Like"] = request.Like
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBuyPageResInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBuyPageResInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBuyPageResInfo(request *QueryBuyPageResInfoRequest) (_result *QueryBuyPageResInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBuyPageResInfoResponse{}
	_body, _err := client.QueryBuyPageResInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBuyResInfoWithOptions(request *QueryBuyResInfoRequest, runtime *util.RuntimeOptions) (_result *QueryBuyResInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Like)) {
		query["Like"] = request.Like
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBuyResInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBuyResInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBuyResInfo(request *QueryBuyResInfoRequest) (_result *QueryBuyResInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBuyResInfoResponse{}
	_body, _err := client.QueryBuyResInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCallRecordingListWithOptions(request *QueryCallRecordingListRequest, runtime *util.RuntimeOptions) (_result *QueryCallRecordingListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.CallDate)) {
		query["CallDate"] = request.CallDate
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoA)) {
		query["PhoneNoA"] = request.PhoneNoA
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNoB)) {
		query["PhoneNoB"] = request.PhoneNoB
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallRecordingList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallRecordingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCallRecordingList(request *QueryCallRecordingListRequest) (_result *QueryCallRecordingListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallRecordingListResponse{}
	_body, _err := client.QueryCallRecordingListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCertifyInfoListWithOptions(request *QueryCertifyInfoListRequest, runtime *util.RuntimeOptions) (_result *QueryCertifyInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertifyStatus)) {
		query["CertifyStatus"] = request.CertifyStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNo)) {
		query["PhoneNo"] = request.PhoneNo
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCertifyInfoList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCertifyInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCertifyInfoList(request *QueryCertifyInfoListRequest) (_result *QueryCertifyInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCertifyInfoListResponse{}
	_body, _err := client.QueryCertifyInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCertifyOverviewInfoWithOptions(request *QueryCertifyOverviewInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCertifyOverviewInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCertifyOverviewInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCertifyOverviewInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCertifyOverviewInfo(request *QueryCertifyOverviewInfoRequest) (_result *QueryCertifyOverviewInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCertifyOverviewInfoResponse{}
	_body, _err := client.QueryCertifyOverviewInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryContactsListWithOptions(request *QueryContactsListRequest, runtime *util.RuntimeOptions) (_result *QueryContactsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryContactsList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryContactsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryContactsList(request *QueryContactsListRequest) (_result *QueryContactsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryContactsListResponse{}
	_body, _err := client.QueryContactsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCustInfoWithOptions(request *QueryCustInfoRequest, runtime *util.RuntimeOptions) (_result *QueryCustInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCustInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCustInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCustInfo(request *QueryCustInfoRequest) (_result *QueryCustInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCustInfoResponse{}
	_body, _err := client.QueryCustInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDownloadUrlWithOptions(request *QueryDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *QueryDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDownloadUrl"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDownloadUrl(request *QueryDownloadUrlRequest) (_result *QueryDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDownloadUrlResponse{}
	_body, _err := client.QueryDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEffectiveInvoiceListByBillNosWithOptions(request *QueryEffectiveInvoiceListByBillNosRequest, runtime *util.RuntimeOptions) (_result *QueryEffectiveInvoiceListByBillNosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		body["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.BillNo)) {
		body["BillNo"] = request.BillNo
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptProps)) {
		bodyFlat["EncryptProps"] = request.EncryptProps
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.MajorBillNo)) {
		body["MajorBillNo"] = request.MajorBillNo
	}

	if !tea.BoolValue(util.IsUnset(request.OuCode)) {
		body["OuCode"] = request.OuCode
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedSystem)) {
		body["RelatedSystem"] = request.RelatedSystem
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		body["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEffectiveInvoiceListByBillNos"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEffectiveInvoiceListByBillNosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEffectiveInvoiceListByBillNos(request *QueryEffectiveInvoiceListByBillNosRequest) (_result *QueryEffectiveInvoiceListByBillNosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEffectiveInvoiceListByBillNosResponse{}
	_body, _err := client.QueryEffectiveInvoiceListByBillNosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryExportResUrlWithOptions(request *QueryExportResUrlRequest, runtime *util.RuntimeOptions) (_result *QueryExportResUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryExportResUrl"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryExportResUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryExportResUrl(request *QueryExportResUrlRequest) (_result *QueryExportResUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryExportResUrlResponse{}
	_body, _err := client.QueryExportResUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupDetailListWithOptions(request *QueryGroupDetailListRequest, runtime *util.RuntimeOptions) (_result *QueryGroupDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGroupDetailList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupDetailList(request *QueryGroupDetailListRequest) (_result *QueryGroupDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGroupDetailListResponse{}
	_body, _err := client.QueryGroupDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupInfoListWithOptions(request *QueryGroupInfoListRequest, runtime *util.RuntimeOptions) (_result *QueryGroupInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.QueryKey)) {
		query["QueryKey"] = request.QueryKey
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGroupInfoList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupInfoList(request *QueryGroupInfoListRequest) (_result *QueryGroupInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGroupInfoListResponse{}
	_body, _err := client.QueryGroupInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInvoiceInfoByRequestNoWithOptions(request *QueryInvoiceInfoByRequestNoRequest, runtime *util.RuntimeOptions) (_result *QueryInvoiceInfoByRequestNoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		body["AppCode"] = request.AppCode
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncryptProps)) {
		bodyFlat["EncryptProps"] = request.EncryptProps
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedSystem)) {
		body["RelatedSystem"] = request.RelatedSystem
	}

	if !tea.BoolValue(util.IsUnset(request.RequestNo)) {
		body["RequestNo"] = request.RequestNo
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		body["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInvoiceInfoByRequestNo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInvoiceInfoByRequestNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInvoiceInfoByRequestNo(request *QueryInvoiceInfoByRequestNoRequest) (_result *QueryInvoiceInfoByRequestNoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInvoiceInfoByRequestNoResponse{}
	_body, _err := client.QueryInvoiceInfoByRequestNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMessageCallbackInfoWithOptions(request *QueryMessageCallbackInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMessageCallbackInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMessageCallbackInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMessageCallbackInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMessageCallbackInfo(request *QueryMessageCallbackInfoRequest) (_result *QueryMessageCallbackInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMessageCallbackInfoResponse{}
	_body, _err := client.QueryMessageCallbackInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMessageQueueListWithOptions(request *QueryMessageQueueListRequest, runtime *util.RuntimeOptions) (_result *QueryMessageQueueListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMessageQueueList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMessageQueueListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMessageQueueList(request *QueryMessageQueueListRequest) (_result *QueryMessageQueueListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMessageQueueListResponse{}
	_body, _err := client.QueryMessageQueueListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonthlyBillInfoWithOptions(request *QueryMonthlyBillInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMonthlyBillInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillCycle)) {
		query["BillCycle"] = request.BillCycle
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemName)) {
		query["ItemName"] = request.ItemName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubjectItemId)) {
		query["SubjectItemId"] = request.SubjectItemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMonthlyBillInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMonthlyBillInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonthlyBillInfo(request *QueryMonthlyBillInfoRequest) (_result *QueryMonthlyBillInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonthlyBillInfoResponse{}
	_body, _err := client.QueryMonthlyBillInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonthlyStatisticsInfoWithOptions(request *QueryMonthlyStatisticsInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMonthlyStatisticsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMonthlyStatisticsInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMonthlyStatisticsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonthlyStatisticsInfo(request *QueryMonthlyStatisticsInfoRequest) (_result *QueryMonthlyStatisticsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonthlyStatisticsInfoResponse{}
	_body, _err := client.QueryMonthlyStatisticsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryNoBuyTasksWithOptions(request *QueryNoBuyTasksRequest, runtime *util.RuntimeOptions) (_result *QueryNoBuyTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryNoBuyTasks"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryNoBuyTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryNoBuyTasks(request *QueryNoBuyTasksRequest) (_result *QueryNoBuyTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryNoBuyTasksResponse{}
	_body, _err := client.QueryNoBuyTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryNoDistributeWithOptions(request *QueryNoDistributeRequest, runtime *util.RuntimeOptions) (_result *QueryNoDistributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryNoDistribute"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryNoDistributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryNoDistribute(request *QueryNoDistributeRequest) (_result *QueryNoDistributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryNoDistributeResponse{}
	_body, _err := client.QueryNoDistributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOpenStatusWithOptions(request *QueryOpenStatusRequest, runtime *util.RuntimeOptions) (_result *QueryOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusOffer)) {
		query["BusOffer"] = request.BusOffer
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProdId)) {
		query["ProdId"] = request.ProdId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOpenStatus"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOpenStatus(request *QueryOpenStatusRequest) (_result *QueryOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOpenStatusResponse{}
	_body, _err := client.QueryOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPackageDetailWithOptions(request *QueryPackageDetailRequest, runtime *util.RuntimeOptions) (_result *QueryPackageDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPackageDetail"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPackageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPackageDetail(request *QueryPackageDetailRequest) (_result *QueryPackageDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPackageDetailResponse{}
	_body, _err := client.QueryPackageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPackageListWithOptions(request *QueryPackageListRequest, runtime *util.RuntimeOptions) (_result *QueryPackageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillCycle)) {
		query["BillCycle"] = request.BillCycle
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPackageList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPackageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPackageList(request *QueryPackageListRequest) (_result *QueryPackageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPackageListResponse{}
	_body, _err := client.QueryPackageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPackageStatisticsWithOptions(request *QueryPackageStatisticsRequest, runtime *util.RuntimeOptions) (_result *QueryPackageStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPackageStatistics"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPackageStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPackageStatistics(request *QueryPackageStatisticsRequest) (_result *QueryPackageStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPackageStatisticsResponse{}
	_body, _err := client.QueryPackageStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPoolCityListWithOptions(request *QueryPoolCityListRequest, runtime *util.RuntimeOptions) (_result *QueryPoolCityListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPoolCityList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPoolCityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPoolCityList(request *QueryPoolCityListRequest) (_result *QueryPoolCityListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPoolCityListResponse{}
	_body, _err := client.QueryPoolCityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPoolInfoListWithOptions(request *QueryPoolInfoListRequest, runtime *util.RuntimeOptions) (_result *QueryPoolInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsFuzzyQuery)) {
		query["IsFuzzyQuery"] = request.IsFuzzyQuery
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParam)) {
		query["SearchParam"] = request.SearchParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPoolInfoList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPoolInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPoolInfoList(request *QueryPoolInfoListRequest) (_result *QueryPoolInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPoolInfoListResponse{}
	_body, _err := client.QueryPoolInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPoolMonthlyBillInfoWithOptions(request *QueryPoolMonthlyBillInfoRequest, runtime *util.RuntimeOptions) (_result *QueryPoolMonthlyBillInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillCycle)) {
		query["BillCycle"] = request.BillCycle
	}

	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPoolMonthlyBillInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPoolMonthlyBillInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPoolMonthlyBillInfo(request *QueryPoolMonthlyBillInfoRequest) (_result *QueryPoolMonthlyBillInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPoolMonthlyBillInfoResponse{}
	_body, _err := client.QueryPoolMonthlyBillInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPoolStatisticsInfoWithOptions(request *QueryPoolStatisticsInfoRequest, runtime *util.RuntimeOptions) (_result *QueryPoolStatisticsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPoolStatisticsInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPoolStatisticsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPoolStatisticsInfo(request *QueryPoolStatisticsInfoRequest) (_result *QueryPoolStatisticsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPoolStatisticsInfoResponse{}
	_body, _err := client.QueryPoolStatisticsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPoolSummaryInfoWithOptions(request *QueryPoolSummaryInfoRequest, runtime *util.RuntimeOptions) (_result *QueryPoolSummaryInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPoolSummaryInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPoolSummaryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPoolSummaryInfo(request *QueryPoolSummaryInfoRequest) (_result *QueryPoolSummaryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPoolSummaryInfoResponse{}
	_body, _err := client.QueryPoolSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPurchasedInfoWithOptions(request *QueryPurchasedInfoRequest, runtime *util.RuntimeOptions) (_result *QueryPurchasedInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPurchasedInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPurchasedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPurchasedInfo(request *QueryPurchasedInfoRequest) (_result *QueryPurchasedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPurchasedInfoResponse{}
	_body, _err := client.QueryPurchasedInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPurchasedResListWithOptions(request *QueryPurchasedResListRequest, runtime *util.RuntimeOptions) (_result *QueryPurchasedResListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.IsDisplayPool)) {
		query["IsDisplayPool"] = request.IsDisplayPool
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResBindStatus)) {
		query["ResBindStatus"] = request.ResBindStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPurchasedResList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPurchasedResListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPurchasedResList(request *QueryPurchasedResListRequest) (_result *QueryPurchasedResListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPurchasedResListResponse{}
	_body, _err := client.QueryPurchasedResListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryQRCodeInfoWithOptions(request *QueryQRCodeInfoRequest, runtime *util.RuntimeOptions) (_result *QueryQRCodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNumber)) {
		query["SecretNumber"] = request.SecretNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryQRCodeInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryQRCodeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryQRCodeInfo(request *QueryQRCodeInfoRequest) (_result *QueryQRCodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryQRCodeInfoResponse{}
	_body, _err := client.QueryQRCodeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRecordingUrlWithOptions(request *QueryRecordingUrlRequest, runtime *util.RuntimeOptions) (_result *QueryRecordingUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.CallDate)) {
		query["CallDate"] = request.CallDate
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRecordingUrl"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecordingUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRecordingUrl(request *QueryRecordingUrlRequest) (_result *QueryRecordingUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRecordingUrlResponse{}
	_body, _err := client.QueryRecordingUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryResSummaryInfoWithOptions(request *QueryResSummaryInfoRequest, runtime *util.RuntimeOptions) (_result *QueryResSummaryInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryResSummaryInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryResSummaryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResSummaryInfo(request *QueryResSummaryInfoRequest) (_result *QueryResSummaryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryResSummaryInfoResponse{}
	_body, _err := client.QueryResSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRingToneUrlWithOptions(request *QueryRingToneUrlRequest, runtime *util.RuntimeOptions) (_result *QueryRingToneUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		query["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRingToneUrl"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRingToneUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRingToneUrl(request *QueryRingToneUrlRequest) (_result *QueryRingToneUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRingToneUrlResponse{}
	_body, _err := client.QueryRingToneUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRingTonesWithOptions(request *QueryRingTonesRequest, runtime *util.RuntimeOptions) (_result *QueryRingTonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlayType)) {
		query["PlayType"] = request.PlayType
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRingTones"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRingTonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRingTones(request *QueryRingTonesRequest) (_result *QueryRingTonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRingTonesResponse{}
	_body, _err := client.QueryRingTonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySimplePoolInfoListWithOptions(request *QuerySimplePoolInfoListRequest, runtime *util.RuntimeOptions) (_result *QuerySimplePoolInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySimplePoolInfoList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySimplePoolInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySimplePoolInfoList(request *QuerySimplePoolInfoListRequest) (_result *QuerySimplePoolInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySimplePoolInfoListResponse{}
	_body, _err := client.QuerySimplePoolInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStatisticsInfoWithOptions(request *QueryStatisticsInfoRequest, runtime *util.RuntimeOptions) (_result *QueryStatisticsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStatisticsInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStatisticsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStatisticsInfo(request *QueryStatisticsInfoRequest) (_result *QueryStatisticsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStatisticsInfoResponse{}
	_body, _err := client.QueryStatisticsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTagOpenStatusWithOptions(request *QueryTagOpenStatusRequest, runtime *util.RuntimeOptions) (_result *QueryTagOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeKey)) {
		query["AttributeKey"] = request.AttributeKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubAttributeKey)) {
		query["SubAttributeKey"] = request.SubAttributeKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTagOpenStatus"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTagOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTagOpenStatus(request *QueryTagOpenStatusRequest) (_result *QueryTagOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTagOpenStatusResponse{}
	_body, _err := client.QueryTagOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTransferDetailsWithOptions(request *QueryTransferDetailsRequest, runtime *util.RuntimeOptions) (_result *QueryTransferDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTransferDetails"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTransferDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTransferDetails(request *QueryTransferDetailsRequest) (_result *QueryTransferDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTransferDetailsResponse{}
	_body, _err := client.QueryTransferDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTransferRecordWithOptions(request *QueryTransferRecordRequest, runtime *util.RuntimeOptions) (_result *QueryTransferRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTransferRecord"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTransferRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTransferRecord(request *QueryTransferRecordRequest) (_result *QueryTransferRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTransferRecordResponse{}
	_body, _err := client.QueryTransferRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTransferRecordsWithOptions(request *QueryTransferRecordsRequest, runtime *util.RuntimeOptions) (_result *QueryTransferRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTransferRecords"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTransferRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTransferRecords(request *QueryTransferRecordsRequest) (_result *QueryTransferRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTransferRecordsResponse{}
	_body, _err := client.QueryTransferRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserDeleteStatusWithOptions(request *QueryUserDeleteStatusRequest, runtime *util.RuntimeOptions) (_result *QueryUserDeleteStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.GmtWakeup)) {
		query["GmtWakeup"] = request.GmtWakeup
	}

	if !tea.BoolValue(util.IsUnset(request.Hid)) {
		query["Hid"] = request.Hid
	}

	if !tea.BoolValue(util.IsUnset(request.Interrupt)) {
		query["Interrupt"] = request.Interrupt
	}

	if !tea.BoolValue(util.IsUnset(request.Invoker)) {
		query["Invoker"] = request.Invoker
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Success)) {
		query["Success"] = request.Success
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExtraData)) {
		query["TaskExtraData"] = request.TaskExtraData
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdentifier)) {
		query["TaskIdentifier"] = request.TaskIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserDeleteStatus"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserDeleteStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserDeleteStatus(request *QueryUserDeleteStatusRequest) (_result *QueryUserDeleteStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserDeleteStatusResponse{}
	_body, _err := client.QueryUserDeleteStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserInfoWithOptions(request *QueryUserInfoRequest, runtime *util.RuntimeOptions) (_result *QueryUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserInfo(request *QueryUserInfoRequest) (_result *QueryUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserInfoResponse{}
	_body, _err := client.QueryUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserResPoolInfoWithOptions(request *QueryUserResPoolInfoRequest, runtime *util.RuntimeOptions) (_result *QueryUserResPoolInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserResPoolInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserResPoolInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserResPoolInfo(request *QueryUserResPoolInfoRequest) (_result *QueryUserResPoolInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserResPoolInfoResponse{}
	_body, _err := client.QueryUserResPoolInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVirtualOperationShowWithOptions(request *QueryVirtualOperationShowRequest, runtime *util.RuntimeOptions) (_result *QueryVirtualOperationShowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVirtualOperationShow"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVirtualOperationShowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVirtualOperationShow(request *QueryVirtualOperationShowRequest) (_result *QueryVirtualOperationShowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVirtualOperationShowResponse{}
	_body, _err := client.QueryVirtualOperationShowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWarningListWithOptions(request *QueryWarningListRequest, runtime *util.RuntimeOptions) (_result *QueryWarningListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWarningList"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWarningListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWarningList(request *QueryWarningListRequest) (_result *QueryWarningListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWarningListResponse{}
	_body, _err := client.QueryWarningListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWaybillOrderInfoWithOptions(request *QueryWaybillOrderInfoRequest, runtime *util.RuntimeOptions) (_result *QueryWaybillOrderInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OuterOrderCode)) {
		query["OuterOrderCode"] = request.OuterOrderCode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWaybillOrderInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWaybillOrderInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWaybillOrderInfo(request *QueryWaybillOrderInfoRequest) (_result *QueryWaybillOrderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWaybillOrderInfoResponse{}
	_body, _err := client.QueryWaybillOrderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWaybillOrderStatisticsInfoWithOptions(request *QueryWaybillOrderStatisticsInfoRequest, runtime *util.RuntimeOptions) (_result *QueryWaybillOrderStatisticsInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryWaybillOrderStatisticsInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryWaybillOrderStatisticsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWaybillOrderStatisticsInfo(request *QueryWaybillOrderStatisticsInfoRequest) (_result *QueryWaybillOrderStatisticsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWaybillOrderStatisticsInfoResponse{}
	_body, _err := client.QueryWaybillOrderStatisticsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseResourceWithOptions(request *ReleaseResourceRequest, runtime *util.RuntimeOptions) (_result *ReleaseResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.IsDisplayPool)) {
		query["IsDisplayPool"] = request.IsDisplayPool
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseResource"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseResource(request *ReleaseResourceRequest) (_result *ReleaseResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseResourceResponse{}
	_body, _err := client.ReleaseResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TestTtsRingToneWithOptions(request *TestTtsRingToneRequest, runtime *util.RuntimeOptions) (_result *TestTtsRingToneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Tts)) {
		query["Tts"] = request.Tts
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceSpeed)) {
		query["VoiceSpeed"] = request.VoiceSpeed
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceStyle)) {
		query["VoiceStyle"] = request.VoiceStyle
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceType)) {
		query["VoiceType"] = request.VoiceType
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceVolume)) {
		query["VoiceVolume"] = request.VoiceVolume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TestTtsRingTone"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TestTtsRingToneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TestTtsRingTone(request *TestTtsRingToneRequest) (_result *TestTtsRingToneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TestTtsRingToneResponse{}
	_body, _err := client.TestTtsRingToneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindResourceWithOptions(request *UnbindResourceRequest, runtime *util.RuntimeOptions) (_result *UnbindResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.BindIds)) {
		query["BindIds"] = request.BindIds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindResource"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindResource(request *UnbindResourceRequest) (_result *UnbindResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindResourceResponse{}
	_body, _err := client.UnbindResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockResourceWithOptions(request *UnlockResourceRequest, runtime *util.RuntimeOptions) (_result *UnlockResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockResource"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockResource(request *UnlockResourceRequest) (_result *UnlockResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockResourceResponse{}
	_body, _err := client.UnlockResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContactsWithOptions(request *UpdateContactsRequest, runtime *util.RuntimeOptions) (_result *UpdateContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateContacts"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContacts(request *UpdateContactsRequest) (_result *UpdateContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContactsResponse{}
	_body, _err := client.UpdateContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupDetailWithOptions(request *UpdateGroupDetailRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupDetail"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupDetail(request *UpdateGroupDetailRequest) (_result *UpdateGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupDetailResponse{}
	_body, _err := client.UpdateGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupInfoWithOptions(request *UpdateGroupInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolKey)) {
		query["PoolKey"] = request.PoolKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupInfo"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroupInfo(request *UpdateGroupInfoRequest) (_result *UpdateGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupInfoResponse{}
	_body, _err := client.UpdateGroupInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePoolNameWithOptions(request *UpdatePoolNameRequest, runtime *util.RuntimeOptions) (_result *UpdatePoolNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PoolName)) {
		query["PoolName"] = request.PoolName
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePoolName"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePoolNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePoolName(request *UpdatePoolNameRequest) (_result *UpdatePoolNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePoolNameResponse{}
	_body, _err := client.UpdatePoolNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResRemarkWithOptions(request *UpdateResRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateResRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		query["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResType)) {
		query["ResType"] = request.ResType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretNo)) {
		query["SecretNo"] = request.SecretNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResRemark"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResRemark(request *UpdateResRemarkRequest) (_result *UpdateResRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResRemarkResponse{}
	_body, _err := client.UpdateResRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateOrderWithOptions(request *ValidateOrderRequest, runtime *util.RuntimeOptions) (_result *ValidateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateOrder"),
		Version:     tea.String("2017-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateOrder(request *ValidateOrderRequest) (_result *ValidateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateOrderResponse{}
	_body, _err := client.ValidateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
