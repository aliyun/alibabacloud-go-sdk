// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserCertificateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateOrderList(v []*ListUserCertificateOrderResponseBodyCertificateOrderList) *ListUserCertificateOrderResponseBody
	GetCertificateOrderList() []*ListUserCertificateOrderResponseBodyCertificateOrderList
	SetCurrentPage(v int64) *ListUserCertificateOrderResponseBody
	GetCurrentPage() *int64
	SetRequestId(v string) *ListUserCertificateOrderResponseBody
	GetRequestId() *string
	SetShowSize(v int64) *ListUserCertificateOrderResponseBody
	GetShowSize() *int64
	SetTotalCount(v int64) *ListUserCertificateOrderResponseBody
	GetTotalCount() *int64
}

type ListUserCertificateOrderResponseBody struct {
	// The certificates and orders.
	CertificateOrderList []*ListUserCertificateOrderResponseBodyCertificateOrderList `json:"CertificateOrderList,omitempty" xml:"CertificateOrderList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	ShowSize *int64 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserCertificateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserCertificateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderResponseBody) GetCertificateOrderList() []*ListUserCertificateOrderResponseBodyCertificateOrderList {
	return s.CertificateOrderList
}

func (s *ListUserCertificateOrderResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListUserCertificateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserCertificateOrderResponseBody) GetShowSize() *int64 {
	return s.ShowSize
}

func (s *ListUserCertificateOrderResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserCertificateOrderResponseBody) SetCertificateOrderList(v []*ListUserCertificateOrderResponseBodyCertificateOrderList) *ListUserCertificateOrderResponseBody {
	s.CertificateOrderList = v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetCurrentPage(v int64) *ListUserCertificateOrderResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetRequestId(v string) *ListUserCertificateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetShowSize(v int64) *ListUserCertificateOrderResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) SetTotalCount(v int64) *ListUserCertificateOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserCertificateOrderResponseBody) Validate() error {
	if s.CertificateOrderList != nil {
		for _, item := range s.CertificateOrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserCertificateOrderResponseBodyCertificateOrderList struct {
	// The algorithm. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the Alibaba Cloud order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 234567
	AliyunOrderId *int64 `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// The time at which the order was placed. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1634283958000
	BuyDate *int64 `json:"BuyDate,omitempty" xml:"BuyDate,omitempty"`
	// The time at which the certificate expires. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1665819958000
	CertEndTime *int64 `json:"CertEndTime,omitempty" xml:"CertEndTime,omitempty"`
	// The time at which the certificate starts to take effect. Unit: milliseconds. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1665819958000
	CertStartTime *int64 `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The type of the certificate. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **DV**: domain validated (DV) certificate
	//
	// 	- **EV**: extended validation (EV) certificate
	//
	// 	- **OV**: organization validated (OV) certificate **FREE**: free certificate, available only on the China site (aliyun.com)
	//
	// example:
	//
	// FREE
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The ID of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 896521
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The city in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The parent domain name of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The domain name. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The total number of domain names that can be bound to the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 1
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The type of the domain name. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **ONE**: single domain name
	//
	// 	- **MULTIPLE**: multiple domain names
	//
	// 	- **WILDCARD**: single wildcard domain name
	//
	// 	- **M_WILDCARD**: multiple wildcard domain names
	//
	// 	- **MIX**: hybrid domain name
	//
	// example:
	//
	// ONE
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The time at which the certificate expires. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 2022-11-17
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether the certificate expires. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// true
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The fingerprint of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// CC6B3696E7C7CA715BD26E28E45FF3E3DF435C03
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cas-instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issuer of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// MyIssuer
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The name of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// cert-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 2345687
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The name of the organization that is associated with the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// Alibaba Cloud
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The ID of the third-party certificate authority (CA) order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// ca-123456
	PartnerOrderId *string `json:"PartnerOrderId,omitempty" xml:"PartnerOrderId,omitempty"`
	// The specification ID of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// bykj123456
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specification name of the order. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// CFCA
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The province or autonomous region in which the organization is located. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// Zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The ID of the resource group. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The brand of the certificate. Valid values: WoSign, CFCA, DigiCert, and vTrus. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// CFCA
	RootBrand *string `json:"RootBrand,omitempty" xml:"RootBrand,omitempty"`
	// All domain names that are bound to the certificate. Multiple domain names are separated by commas (,). This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// aliyun.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 040a6e493cffdda6d744acf99b6576cf
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// The SHA-2 value of the certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 56B4DED2243A81DD909D7C39824FFE4DDBD87F91BFA46CD333FF212FE0E7CB11
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The type of the order. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **cpack**: virtual resource order
	//
	// 	- **buy**: purchase order
	//
	// example:
	//
	// buy
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time at which the certificate starts to take effect. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// 2021-11-16
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The certificate status of the order. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **PAYED**: pending application
	//
	// 	- **CHECKING**: reviewing
	//
	// 	- **CHECKED_FAIL**: review failed
	//
	// 	- **ISSUED**: issued
	//
	// 	- **WILLEXPIRED**: about to expire
	//
	// 	- **EXPIRED**: expired
	//
	// 	- **NOTACTIVATED**: not activated
	//
	// 	- **REVOKED**: revoked
	//
	// example:
	//
	// PAYED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The hosting status of the certificate. This parameter is returned only if OrderType is set to CPACK or BUY. Valid values:
	//
	// 	- **unTrustee**: not hosted
	//
	// 	- **trustee**: hosted
	//
	// example:
	//
	// unTrustee
	TrusteeStatus *string `json:"TrusteeStatus,omitempty" xml:"TrusteeStatus,omitempty"`
	// Indicates whether the certificate is an uploaded certificate. This parameter is returned only if OrderType is set to CERT or UPLOAD.
	//
	// example:
	//
	// false
	Upload *bool `json:"Upload,omitempty" xml:"Upload,omitempty"`
	// The number of wildcard domain names that can be bound to the certificate. This parameter is returned only if OrderType is set to CPACK or BUY.
	//
	// example:
	//
	// 0
	WildDomainCount *int64 `json:"WildDomainCount,omitempty" xml:"WildDomainCount,omitempty"`
}

func (s ListUserCertificateOrderResponseBodyCertificateOrderList) String() string {
	return dara.Prettify(s)
}

func (s ListUserCertificateOrderResponseBodyCertificateOrderList) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetAliyunOrderId() *int64 {
	return s.AliyunOrderId
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetBuyDate() *int64 {
	return s.BuyDate
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCertEndTime() *int64 {
	return s.CertEndTime
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCertStartTime() *int64 {
	return s.CertStartTime
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCertType() *string {
	return s.CertType
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCertificateId() *int64 {
	return s.CertificateId
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCity() *string {
	return s.City
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetCountry() *string {
	return s.Country
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetDomain() *string {
	return s.Domain
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetDomainType() *string {
	return s.DomainType
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetEndDate() *string {
	return s.EndDate
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetExpired() *bool {
	return s.Expired
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetIssuer() *string {
	return s.Issuer
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetName() *string {
	return s.Name
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetOrgName() *string {
	return s.OrgName
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetPartnerOrderId() *string {
	return s.PartnerOrderId
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetProductName() *string {
	return s.ProductName
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetProvince() *string {
	return s.Province
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetRootBrand() *string {
	return s.RootBrand
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetSans() *string {
	return s.Sans
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetSha2() *string {
	return s.Sha2
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetStartDate() *string {
	return s.StartDate
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetStatus() *string {
	return s.Status
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetTrusteeStatus() *string {
	return s.TrusteeStatus
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetUpload() *bool {
	return s.Upload
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) GetWildDomainCount() *int64 {
	return s.WildDomainCount
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetAlgorithm(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Algorithm = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetAliyunOrderId(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.AliyunOrderId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetBuyDate(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.BuyDate = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertEndTime(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertEndTime = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertStartTime(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertStartTime = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertType(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertType = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCertificateId(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CertificateId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCity(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.City = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCommonName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.CommonName = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetCountry(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Country = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetDomain(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Domain = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetDomainCount(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.DomainCount = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetDomainType(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.DomainType = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetEndDate(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.EndDate = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetExpired(v bool) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Expired = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetFingerprint(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Fingerprint = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetInstanceId(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.InstanceId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetIssuer(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Issuer = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Name = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetOrderId(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.OrderId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetOrgName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.OrgName = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetPartnerOrderId(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.PartnerOrderId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetProductCode(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.ProductCode = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetProductName(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.ProductName = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetProvince(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Province = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetResourceGroupId(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetRootBrand(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.RootBrand = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSans(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Sans = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSerialNo(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.SerialNo = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSha2(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Sha2 = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetSourceType(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.SourceType = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetStartDate(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.StartDate = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetStatus(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Status = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetTrusteeStatus(v string) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.TrusteeStatus = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetUpload(v bool) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.Upload = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) SetWildDomainCount(v int64) *ListUserCertificateOrderResponseBodyCertificateOrderList {
	s.WildDomainCount = &v
	return s
}

func (s *ListUserCertificateOrderResponseBodyCertificateOrderList) Validate() error {
	return dara.Validate(s)
}
