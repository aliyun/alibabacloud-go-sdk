// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomerSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindCustomerSnapshotResponseBody
	GetCode() *string
	SetCustomerSnapshot(v *FindCustomerSnapshotResponseBodyCustomerSnapshot) *FindCustomerSnapshotResponseBody
	GetCustomerSnapshot() *FindCustomerSnapshotResponseBodyCustomerSnapshot
	SetMessage(v string) *FindCustomerSnapshotResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindCustomerSnapshotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindCustomerSnapshotResponseBody
	GetSuccess() *bool
}

type FindCustomerSnapshotResponseBody struct {
	Code             *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	CustomerSnapshot *FindCustomerSnapshotResponseBodyCustomerSnapshot `json:"CustomerSnapshot,omitempty" xml:"CustomerSnapshot,omitempty" type:"Struct"`
	Message          *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindCustomerSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *FindCustomerSnapshotResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindCustomerSnapshotResponseBody) GetCustomerSnapshot() *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	return s.CustomerSnapshot
}

func (s *FindCustomerSnapshotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindCustomerSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindCustomerSnapshotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindCustomerSnapshotResponseBody) SetCode(v string) *FindCustomerSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *FindCustomerSnapshotResponseBody) SetCustomerSnapshot(v *FindCustomerSnapshotResponseBodyCustomerSnapshot) *FindCustomerSnapshotResponseBody {
	s.CustomerSnapshot = v
	return s
}

func (s *FindCustomerSnapshotResponseBody) SetMessage(v string) *FindCustomerSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *FindCustomerSnapshotResponseBody) SetRequestId(v string) *FindCustomerSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindCustomerSnapshotResponseBody) SetSuccess(v bool) *FindCustomerSnapshotResponseBody {
	s.Success = &v
	return s
}

func (s *FindCustomerSnapshotResponseBody) Validate() error {
	if s.CustomerSnapshot != nil {
		if err := s.CustomerSnapshot.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindCustomerSnapshotResponseBodyCustomerSnapshot struct {
	AccountInfoSnapshotModel *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel `json:"AccountInfoSnapshotModel,omitempty" xml:"AccountInfoSnapshotModel,omitempty" type:"Struct"`
	AccountTaxSnapshotModel  *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel  `json:"AccountTaxSnapshotModel,omitempty" xml:"AccountTaxSnapshotModel,omitempty" type:"Struct"`
	GmtCreate                *string                                                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                       *int64                                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	InfoType                 *string                                                                   `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	KpId                     *int64                                                                    `json:"KpId,omitempty" xml:"KpId,omitempty"`
}

func (s FindCustomerSnapshotResponseBodyCustomerSnapshot) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerSnapshotResponseBodyCustomerSnapshot) GoString() string {
	return s.String()
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) GetAccountInfoSnapshotModel() *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	return s.AccountInfoSnapshotModel
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) GetAccountTaxSnapshotModel() *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	return s.AccountTaxSnapshotModel
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) GetId() *int64 {
	return s.Id
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) GetInfoType() *string {
	return s.InfoType
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) GetKpId() *int64 {
	return s.KpId
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) SetAccountInfoSnapshotModel(v *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	s.AccountInfoSnapshotModel = v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) SetAccountTaxSnapshotModel(v *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	s.AccountTaxSnapshotModel = v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) SetGmtCreate(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	s.GmtCreate = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) SetId(v int64) *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	s.Id = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) SetInfoType(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	s.InfoType = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) SetKpId(v int64) *FindCustomerSnapshotResponseBodyCustomerSnapshot {
	s.KpId = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshot) Validate() error {
	if s.AccountInfoSnapshotModel != nil {
		if err := s.AccountInfoSnapshotModel.Validate(); err != nil {
			return err
		}
	}
	if s.AccountTaxSnapshotModel != nil {
		if err := s.AccountTaxSnapshotModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel struct {
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2     *string `json:"Address2,omitempty" xml:"Address2,omitempty"`
	Address3     *string `json:"Address3,omitempty" xml:"Address3,omitempty"`
	Address4     *string `json:"Address4,omitempty" xml:"Address4,omitempty"`
	Address5     *string `json:"Address5,omitempty" xml:"Address5,omitempty"`
	Address6     *string `json:"Address6,omitempty" xml:"Address6,omitempty"`
	CityId       *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	CityName     *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	PostCode     *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	ProvinceId   *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
	ProvinceName *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	TrueName     *string `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
}

func (s FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GoString() string {
	return s.String()
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetAddress() *string {
	return s.Address
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetAddress2() *string {
	return s.Address2
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetAddress3() *string {
	return s.Address3
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetAddress4() *string {
	return s.Address4
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetAddress5() *string {
	return s.Address5
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetAddress6() *string {
	return s.Address6
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetCityId() *string {
	return s.CityId
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetCityName() *string {
	return s.CityName
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetPostCode() *string {
	return s.PostCode
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetProvinceId() *string {
	return s.ProvinceId
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) GetTrueName() *string {
	return s.TrueName
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetAddress(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.Address = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetAddress2(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.Address2 = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetAddress3(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.Address3 = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetAddress4(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.Address4 = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetAddress5(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.Address5 = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetAddress6(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.Address6 = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetCityId(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.CityId = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetCityName(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.CityName = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetPostCode(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.PostCode = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetProvinceId(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.ProvinceId = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetProvinceName(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.ProvinceName = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) SetTrueName(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel {
	s.TrueName = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountInfoSnapshotModel) Validate() error {
	return dara.Validate(s)
}

type FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel struct {
	FinanceTaxCertificateImgName       *string `json:"FinanceTaxCertificateImgName,omitempty" xml:"FinanceTaxCertificateImgName,omitempty"`
	FinanceTaxCertificateImgUrl        *string `json:"FinanceTaxCertificateImgUrl,omitempty" xml:"FinanceTaxCertificateImgUrl,omitempty"`
	SecondFinanceTax                   *string `json:"SecondFinanceTax,omitempty" xml:"SecondFinanceTax,omitempty"`
	SecondFinanceTaxCertificateImgName *string `json:"SecondFinanceTaxCertificateImgName,omitempty" xml:"SecondFinanceTaxCertificateImgName,omitempty"`
	SecondFinanceTaxCertificateImgUrl  *string `json:"SecondFinanceTaxCertificateImgUrl,omitempty" xml:"SecondFinanceTaxCertificateImgUrl,omitempty"`
	Tax                                *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
}

func (s FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) String() string {
	return dara.Prettify(s)
}

func (s FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GoString() string {
	return s.String()
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GetFinanceTaxCertificateImgName() *string {
	return s.FinanceTaxCertificateImgName
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GetFinanceTaxCertificateImgUrl() *string {
	return s.FinanceTaxCertificateImgUrl
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GetSecondFinanceTax() *string {
	return s.SecondFinanceTax
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GetSecondFinanceTaxCertificateImgName() *string {
	return s.SecondFinanceTaxCertificateImgName
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GetSecondFinanceTaxCertificateImgUrl() *string {
	return s.SecondFinanceTaxCertificateImgUrl
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) GetTax() *string {
	return s.Tax
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) SetFinanceTaxCertificateImgName(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	s.FinanceTaxCertificateImgName = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) SetFinanceTaxCertificateImgUrl(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	s.FinanceTaxCertificateImgUrl = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) SetSecondFinanceTax(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	s.SecondFinanceTax = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) SetSecondFinanceTaxCertificateImgName(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	s.SecondFinanceTaxCertificateImgName = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) SetSecondFinanceTaxCertificateImgUrl(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	s.SecondFinanceTaxCertificateImgUrl = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) SetTax(v string) *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel {
	s.Tax = &v
	return s
}

func (s *FindCustomerSnapshotResponseBodyCustomerSnapshotAccountTaxSnapshotModel) Validate() error {
	return dara.Validate(s)
}
