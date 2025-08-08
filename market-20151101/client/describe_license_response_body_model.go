// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicense(v *DescribeLicenseResponseBodyLicense) *DescribeLicenseResponseBody
	GetLicense() *DescribeLicenseResponseBodyLicense
	SetRequestId(v string) *DescribeLicenseResponseBody
	GetRequestId() *string
}

type DescribeLicenseResponseBody struct {
	License *DescribeLicenseResponseBodyLicense `json:"License,omitempty" xml:"License,omitempty" type:"Struct"`
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBody) GetLicense() *DescribeLicenseResponseBodyLicense {
	return s.License
}

func (s *DescribeLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLicenseResponseBody) SetLicense(v *DescribeLicenseResponseBodyLicense) *DescribeLicenseResponseBody {
	s.License = v
	return s
}

func (s *DescribeLicenseResponseBody) SetRequestId(v string) *DescribeLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLicenseResponseBodyLicense struct {
	// example:
	//
	// 2019-05-25 09:00:00
	ActivateTime *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	// example:
	//
	// 2019-05-25 09:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2019-06-25 09:00:00
	ExpiredTime *string                                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ExtendArray *DescribeLicenseResponseBodyLicenseExtendArray `json:"ExtendArray,omitempty" xml:"ExtendArray,omitempty" type:"Struct"`
	ExtendInfo  *DescribeLicenseResponseBodyLicenseExtendInfo  `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1551111111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// -
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// example:
	//
	// ACTIVATED
	LicenseStatus *string `json:"LicenseStatus,omitempty" xml:"LicenseStatus,omitempty"`
	// example:
	//
	// cmgj02****
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// cmgj02****-prepay
	ProductSkuId *string `json:"ProductSkuId,omitempty" xml:"ProductSkuId,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeLicenseResponseBodyLicense) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicense) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicense) GetActivateTime() *string {
	return s.ActivateTime
}

func (s *DescribeLicenseResponseBodyLicense) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLicenseResponseBodyLicense) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeLicenseResponseBodyLicense) GetExtendArray() *DescribeLicenseResponseBodyLicenseExtendArray {
	return s.ExtendArray
}

func (s *DescribeLicenseResponseBodyLicense) GetExtendInfo() *DescribeLicenseResponseBodyLicenseExtendInfo {
	return s.ExtendInfo
}

func (s *DescribeLicenseResponseBodyLicense) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLicenseResponseBodyLicense) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeLicenseResponseBodyLicense) GetLicenseStatus() *string {
	return s.LicenseStatus
}

func (s *DescribeLicenseResponseBodyLicense) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeLicenseResponseBodyLicense) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeLicenseResponseBodyLicense) GetProductSkuId() *string {
	return s.ProductSkuId
}

func (s *DescribeLicenseResponseBodyLicense) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeLicenseResponseBodyLicense) SetActivateTime(v string) *DescribeLicenseResponseBodyLicense {
	s.ActivateTime = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetCreateTime(v string) *DescribeLicenseResponseBodyLicense {
	s.CreateTime = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetExpiredTime(v string) *DescribeLicenseResponseBodyLicense {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetExtendArray(v *DescribeLicenseResponseBodyLicenseExtendArray) *DescribeLicenseResponseBodyLicense {
	s.ExtendArray = v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetExtendInfo(v *DescribeLicenseResponseBodyLicenseExtendInfo) *DescribeLicenseResponseBodyLicense {
	s.ExtendInfo = v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetInstanceId(v string) *DescribeLicenseResponseBodyLicense {
	s.InstanceId = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetLicenseCode(v string) *DescribeLicenseResponseBodyLicense {
	s.LicenseCode = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetLicenseStatus(v string) *DescribeLicenseResponseBodyLicense {
	s.LicenseStatus = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetProductCode(v string) *DescribeLicenseResponseBodyLicense {
	s.ProductCode = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetProductName(v string) *DescribeLicenseResponseBodyLicense {
	s.ProductName = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetProductSkuId(v string) *DescribeLicenseResponseBodyLicense {
	s.ProductSkuId = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) SetSupplierName(v string) *DescribeLicenseResponseBodyLicense {
	s.SupplierName = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicense) Validate() error {
	return dara.Validate(s)
}

type DescribeLicenseResponseBodyLicenseExtendArray struct {
	LicenseAttribute []*DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute `json:"LicenseAttribute,omitempty" xml:"LicenseAttribute,omitempty" type:"Repeated"`
}

func (s DescribeLicenseResponseBodyLicenseExtendArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicenseExtendArray) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicenseExtendArray) GetLicenseAttribute() []*DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute {
	return s.LicenseAttribute
}

func (s *DescribeLicenseResponseBodyLicenseExtendArray) SetLicenseAttribute(v []*DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) *DescribeLicenseResponseBodyLicenseExtendArray {
	s.LicenseAttribute = v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendArray) Validate() error {
	return dara.Validate(s)
}

type DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute struct {
	// example:
	//
	// -
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) GetCode() *string {
	return s.Code
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) GetValue() *string {
	return s.Value
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) SetCode(v string) *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute {
	s.Code = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) SetValue(v string) *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute {
	s.Value = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendArrayLicenseAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeLicenseResponseBodyLicenseExtendInfo struct {
	// example:
	//
	// -
	AccountQuantity *int64 `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	// example:
	//
	// 190311111111****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// id***@**.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 129****1111
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s DescribeLicenseResponseBodyLicenseExtendInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseResponseBodyLicenseExtendInfo) GoString() string {
	return s.String()
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) GetAccountQuantity() *int64 {
	return s.AccountQuantity
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) GetEmail() *string {
	return s.Email
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) GetMobile() *string {
	return s.Mobile
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetAccountQuantity(v int64) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.AccountQuantity = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetAliUid(v int64) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.AliUid = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetEmail(v string) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.Email = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) SetMobile(v string) *DescribeLicenseResponseBodyLicenseExtendInfo {
	s.Mobile = &v
	return s
}

func (s *DescribeLicenseResponseBodyLicenseExtendInfo) Validate() error {
	return dara.Validate(s)
}
