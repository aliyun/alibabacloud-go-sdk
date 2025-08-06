// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataBusinessesValue interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *DataBusinessesValue
	GetId() *int32
	SetName(v string) *DataBusinessesValue
	GetName() *string
	SetCname(v string) *DataBusinessesValue
	GetCname() *string
	SetSname(v string) *DataBusinessesValue
	GetSname() *string
	SetSdkCodes(v []*DataBusinessesValueSdkCodes) *DataBusinessesValue
	GetSdkCodes() []*DataBusinessesValueSdkCodes
	SetSdkFeatures(v []*DataBusinessesValueSdkFeatures) *DataBusinessesValue
	GetSdkFeatures() []*DataBusinessesValueSdkFeatures
	SetVersion(v string) *DataBusinessesValue
	GetVersion() *string
	SetIceOpen(v int32) *DataBusinessesValue
	GetIceOpen() *int32
	SetSupportPlatform(v int32) *DataBusinessesValue
	GetSupportPlatform() *int32
	SetFunctionDesc(v string) *DataBusinessesValue
	GetFunctionDesc() *string
	SetProductDesc(v string) *DataBusinessesValue
	GetProductDesc() *string
	SetSize(v string) *DataBusinessesValue
	GetSize() *string
	SetAuthorized(v int32) *DataBusinessesValue
	GetAuthorized() *int32
	SetAuthorizedApp(v []*string) *DataBusinessesValue
	GetAuthorizedApp() []*string
}

type DataBusinessesValue struct {
	Id              *int32                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Cname           *string                           `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Sname           *string                           `json:"Sname,omitempty" xml:"Sname,omitempty"`
	SdkCodes        []*DataBusinessesValueSdkCodes    `json:"SdkCodes,omitempty" xml:"SdkCodes,omitempty" type:"Repeated"`
	SdkFeatures     []*DataBusinessesValueSdkFeatures `json:"SdkFeatures,omitempty" xml:"SdkFeatures,omitempty" type:"Repeated"`
	Version         *string                           `json:"Version,omitempty" xml:"Version,omitempty"`
	IceOpen         *int32                            `json:"IceOpen,omitempty" xml:"IceOpen,omitempty"`
	SupportPlatform *int32                            `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
	FunctionDesc    *string                           `json:"FunctionDesc,omitempty" xml:"FunctionDesc,omitempty"`
	ProductDesc     *string                           `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	Size            *string                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Authorized      *int32                            `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	AuthorizedApp   []*string                         `json:"AuthorizedApp,omitempty" xml:"AuthorizedApp,omitempty" type:"Repeated"`
}

func (s DataBusinessesValue) String() string {
	return dara.Prettify(s)
}

func (s DataBusinessesValue) GoString() string {
	return s.String()
}

func (s *DataBusinessesValue) GetId() *int32 {
	return s.Id
}

func (s *DataBusinessesValue) GetName() *string {
	return s.Name
}

func (s *DataBusinessesValue) GetCname() *string {
	return s.Cname
}

func (s *DataBusinessesValue) GetSname() *string {
	return s.Sname
}

func (s *DataBusinessesValue) GetSdkCodes() []*DataBusinessesValueSdkCodes {
	return s.SdkCodes
}

func (s *DataBusinessesValue) GetSdkFeatures() []*DataBusinessesValueSdkFeatures {
	return s.SdkFeatures
}

func (s *DataBusinessesValue) GetVersion() *string {
	return s.Version
}

func (s *DataBusinessesValue) GetIceOpen() *int32 {
	return s.IceOpen
}

func (s *DataBusinessesValue) GetSupportPlatform() *int32 {
	return s.SupportPlatform
}

func (s *DataBusinessesValue) GetFunctionDesc() *string {
	return s.FunctionDesc
}

func (s *DataBusinessesValue) GetProductDesc() *string {
	return s.ProductDesc
}

func (s *DataBusinessesValue) GetSize() *string {
	return s.Size
}

func (s *DataBusinessesValue) GetAuthorized() *int32 {
	return s.Authorized
}

func (s *DataBusinessesValue) GetAuthorizedApp() []*string {
	return s.AuthorizedApp
}

func (s *DataBusinessesValue) SetId(v int32) *DataBusinessesValue {
	s.Id = &v
	return s
}

func (s *DataBusinessesValue) SetName(v string) *DataBusinessesValue {
	s.Name = &v
	return s
}

func (s *DataBusinessesValue) SetCname(v string) *DataBusinessesValue {
	s.Cname = &v
	return s
}

func (s *DataBusinessesValue) SetSname(v string) *DataBusinessesValue {
	s.Sname = &v
	return s
}

func (s *DataBusinessesValue) SetSdkCodes(v []*DataBusinessesValueSdkCodes) *DataBusinessesValue {
	s.SdkCodes = v
	return s
}

func (s *DataBusinessesValue) SetSdkFeatures(v []*DataBusinessesValueSdkFeatures) *DataBusinessesValue {
	s.SdkFeatures = v
	return s
}

func (s *DataBusinessesValue) SetVersion(v string) *DataBusinessesValue {
	s.Version = &v
	return s
}

func (s *DataBusinessesValue) SetIceOpen(v int32) *DataBusinessesValue {
	s.IceOpen = &v
	return s
}

func (s *DataBusinessesValue) SetSupportPlatform(v int32) *DataBusinessesValue {
	s.SupportPlatform = &v
	return s
}

func (s *DataBusinessesValue) SetFunctionDesc(v string) *DataBusinessesValue {
	s.FunctionDesc = &v
	return s
}

func (s *DataBusinessesValue) SetProductDesc(v string) *DataBusinessesValue {
	s.ProductDesc = &v
	return s
}

func (s *DataBusinessesValue) SetSize(v string) *DataBusinessesValue {
	s.Size = &v
	return s
}

func (s *DataBusinessesValue) SetAuthorized(v int32) *DataBusinessesValue {
	s.Authorized = &v
	return s
}

func (s *DataBusinessesValue) SetAuthorizedApp(v []*string) *DataBusinessesValue {
	s.AuthorizedApp = v
	return s
}

func (s *DataBusinessesValue) Validate() error {
	return dara.Validate(s)
}

type DataBusinessesValueSdkCodes struct {
	Id              *int32                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	SellPrice       *string                                 `json:"SellPrice,omitempty" xml:"SellPrice,omitempty"`
	IceOpen         *int32                                  `json:"IceOpen,omitempty" xml:"IceOpen,omitempty"`
	SupportPlatform *int32                                  `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
	Size            *string                                 `json:"Size,omitempty" xml:"Size,omitempty"`
	Authorized      *int32                                  `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	AuthorizedApp   []*string                               `json:"AuthorizedApp,omitempty" xml:"AuthorizedApp,omitempty" type:"Repeated"`
	DefaultFeature  *string                                 `json:"DefaultFeature,omitempty" xml:"DefaultFeature,omitempty"`
	Integrate       []*DataBusinessesValueSdkCodesIntegrate `json:"Integrate,omitempty" xml:"Integrate,omitempty" type:"Repeated"`
	Subscription    *string                                 `json:"Subscription,omitempty" xml:"Subscription,omitempty"`
	ProductDesc     *string                                 `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	SubscriptionPkg *string                                 `json:"SubscriptionPkg,omitempty" xml:"SubscriptionPkg,omitempty"`
	SubscriptionImp *string                                 `json:"SubscriptionImp,omitempty" xml:"SubscriptionImp,omitempty"`
}

func (s DataBusinessesValueSdkCodes) String() string {
	return dara.Prettify(s)
}

func (s DataBusinessesValueSdkCodes) GoString() string {
	return s.String()
}

func (s *DataBusinessesValueSdkCodes) GetId() *int32 {
	return s.Id
}

func (s *DataBusinessesValueSdkCodes) GetName() *string {
	return s.Name
}

func (s *DataBusinessesValueSdkCodes) GetSellPrice() *string {
	return s.SellPrice
}

func (s *DataBusinessesValueSdkCodes) GetIceOpen() *int32 {
	return s.IceOpen
}

func (s *DataBusinessesValueSdkCodes) GetSupportPlatform() *int32 {
	return s.SupportPlatform
}

func (s *DataBusinessesValueSdkCodes) GetSize() *string {
	return s.Size
}

func (s *DataBusinessesValueSdkCodes) GetAuthorized() *int32 {
	return s.Authorized
}

func (s *DataBusinessesValueSdkCodes) GetAuthorizedApp() []*string {
	return s.AuthorizedApp
}

func (s *DataBusinessesValueSdkCodes) GetDefaultFeature() *string {
	return s.DefaultFeature
}

func (s *DataBusinessesValueSdkCodes) GetIntegrate() []*DataBusinessesValueSdkCodesIntegrate {
	return s.Integrate
}

func (s *DataBusinessesValueSdkCodes) GetSubscription() *string {
	return s.Subscription
}

func (s *DataBusinessesValueSdkCodes) GetProductDesc() *string {
	return s.ProductDesc
}

func (s *DataBusinessesValueSdkCodes) GetSubscriptionPkg() *string {
	return s.SubscriptionPkg
}

func (s *DataBusinessesValueSdkCodes) GetSubscriptionImp() *string {
	return s.SubscriptionImp
}

func (s *DataBusinessesValueSdkCodes) SetId(v int32) *DataBusinessesValueSdkCodes {
	s.Id = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetName(v string) *DataBusinessesValueSdkCodes {
	s.Name = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetSellPrice(v string) *DataBusinessesValueSdkCodes {
	s.SellPrice = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetIceOpen(v int32) *DataBusinessesValueSdkCodes {
	s.IceOpen = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetSupportPlatform(v int32) *DataBusinessesValueSdkCodes {
	s.SupportPlatform = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetSize(v string) *DataBusinessesValueSdkCodes {
	s.Size = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetAuthorized(v int32) *DataBusinessesValueSdkCodes {
	s.Authorized = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetAuthorizedApp(v []*string) *DataBusinessesValueSdkCodes {
	s.AuthorizedApp = v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetDefaultFeature(v string) *DataBusinessesValueSdkCodes {
	s.DefaultFeature = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetIntegrate(v []*DataBusinessesValueSdkCodesIntegrate) *DataBusinessesValueSdkCodes {
	s.Integrate = v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetSubscription(v string) *DataBusinessesValueSdkCodes {
	s.Subscription = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetProductDesc(v string) *DataBusinessesValueSdkCodes {
	s.ProductDesc = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetSubscriptionPkg(v string) *DataBusinessesValueSdkCodes {
	s.SubscriptionPkg = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) SetSubscriptionImp(v string) *DataBusinessesValueSdkCodes {
	s.SubscriptionImp = &v
	return s
}

func (s *DataBusinessesValueSdkCodes) Validate() error {
	return dara.Validate(s)
}

type DataBusinessesValueSdkCodesIntegrate struct {
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Pkg      *string `json:"Pkg,omitempty" xml:"Pkg,omitempty"`
	Size     *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DataBusinessesValueSdkCodesIntegrate) String() string {
	return dara.Prettify(s)
}

func (s DataBusinessesValueSdkCodesIntegrate) GoString() string {
	return s.String()
}

func (s *DataBusinessesValueSdkCodesIntegrate) GetPlatform() *string {
	return s.Platform
}

func (s *DataBusinessesValueSdkCodesIntegrate) GetCode() *string {
	return s.Code
}

func (s *DataBusinessesValueSdkCodesIntegrate) GetPkg() *string {
	return s.Pkg
}

func (s *DataBusinessesValueSdkCodesIntegrate) GetSize() *string {
	return s.Size
}

func (s *DataBusinessesValueSdkCodesIntegrate) SetPlatform(v string) *DataBusinessesValueSdkCodesIntegrate {
	s.Platform = &v
	return s
}

func (s *DataBusinessesValueSdkCodesIntegrate) SetCode(v string) *DataBusinessesValueSdkCodesIntegrate {
	s.Code = &v
	return s
}

func (s *DataBusinessesValueSdkCodesIntegrate) SetPkg(v string) *DataBusinessesValueSdkCodesIntegrate {
	s.Pkg = &v
	return s
}

func (s *DataBusinessesValueSdkCodesIntegrate) SetSize(v string) *DataBusinessesValueSdkCodesIntegrate {
	s.Size = &v
	return s
}

func (s *DataBusinessesValueSdkCodesIntegrate) Validate() error {
	return dara.Validate(s)
}

type DataBusinessesValueSdkFeatures struct {
	Id     *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Scode  *string `json:"Scode,omitempty" xml:"Scode,omitempty"`
	Svalue *string `json:"Svalue,omitempty" xml:"Svalue,omitempty"`
}

func (s DataBusinessesValueSdkFeatures) String() string {
	return dara.Prettify(s)
}

func (s DataBusinessesValueSdkFeatures) GoString() string {
	return s.String()
}

func (s *DataBusinessesValueSdkFeatures) GetId() *int32 {
	return s.Id
}

func (s *DataBusinessesValueSdkFeatures) GetName() *string {
	return s.Name
}

func (s *DataBusinessesValueSdkFeatures) GetScode() *string {
	return s.Scode
}

func (s *DataBusinessesValueSdkFeatures) GetSvalue() *string {
	return s.Svalue
}

func (s *DataBusinessesValueSdkFeatures) SetId(v int32) *DataBusinessesValueSdkFeatures {
	s.Id = &v
	return s
}

func (s *DataBusinessesValueSdkFeatures) SetName(v string) *DataBusinessesValueSdkFeatures {
	s.Name = &v
	return s
}

func (s *DataBusinessesValueSdkFeatures) SetScode(v string) *DataBusinessesValueSdkFeatures {
	s.Scode = &v
	return s
}

func (s *DataBusinessesValueSdkFeatures) SetSvalue(v string) *DataBusinessesValueSdkFeatures {
	s.Svalue = &v
	return s
}

func (s *DataBusinessesValueSdkFeatures) Validate() error {
	return dara.Validate(s)
}
