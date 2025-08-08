// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceForIsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddress(v string) *DescribeInstanceForIsvResponseBody
	GetActiveAddress() *string
	SetAppJson(v string) *DescribeInstanceForIsvResponseBody
	GetAppJson() *string
	SetAutoRenewal(v string) *DescribeInstanceForIsvResponseBody
	GetAutoRenewal() *string
	SetBeganOn(v int64) *DescribeInstanceForIsvResponseBody
	GetBeganOn() *int64
	SetComponentJson(v string) *DescribeInstanceForIsvResponseBody
	GetComponentJson() *string
	SetCreatedOn(v int64) *DescribeInstanceForIsvResponseBody
	GetCreatedOn() *int64
	SetEndOn(v int64) *DescribeInstanceForIsvResponseBody
	GetEndOn() *int64
	SetExtendJson(v string) *DescribeInstanceForIsvResponseBody
	GetExtendJson() *string
	SetHostJson(v string) *DescribeInstanceForIsvResponseBody
	GetHostJson() *string
	SetImageJson(v string) *DescribeInstanceForIsvResponseBody
	GetImageJson() *string
	SetInstanceId(v int64) *DescribeInstanceForIsvResponseBody
	GetInstanceId() *int64
	SetIsTrial(v bool) *DescribeInstanceForIsvResponseBody
	GetIsTrial() *bool
	SetLicenseCode(v string) *DescribeInstanceForIsvResponseBody
	GetLicenseCode() *string
	SetOrderId(v int64) *DescribeInstanceForIsvResponseBody
	GetOrderId() *int64
	SetProductCode(v string) *DescribeInstanceForIsvResponseBody
	GetProductCode() *string
	SetProductName(v string) *DescribeInstanceForIsvResponseBody
	GetProductName() *string
	SetProductSkuCode(v string) *DescribeInstanceForIsvResponseBody
	GetProductSkuCode() *string
	SetProductType(v string) *DescribeInstanceForIsvResponseBody
	GetProductType() *string
	SetRelationalData(v *DescribeInstanceForIsvResponseBodyRelationalData) *DescribeInstanceForIsvResponseBody
	GetRelationalData() *DescribeInstanceForIsvResponseBodyRelationalData
	SetRequestId(v string) *DescribeInstanceForIsvResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeInstanceForIsvResponseBody
	GetStatus() *string
	SetSupplierName(v string) *DescribeInstanceForIsvResponseBody
	GetSupplierName() *string
}

type DescribeInstanceForIsvResponseBody struct {
	ActiveAddress *string `json:"ActiveAddress,omitempty" xml:"ActiveAddress,omitempty"`
	AppJson       *string `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	AutoRenewal   *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// 1570634021000
	BeganOn *int64 `json:"BeganOn,omitempty" xml:"BeganOn,omitempty"`
	// example:
	//
	// {"package_version":"yuncode000111"}
	ComponentJson *string `json:"ComponentJson,omitempty" xml:"ComponentJson,omitempty"`
	// example:
	//
	// 1570634018000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// example:
	//
	// 1602259200000
	EndOn      *int64  `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	ExtendJson *string `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	HostJson   *string `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	ImageJson  *string `json:"ImageJson,omitempty" xml:"ImageJson,omitempty"`
	// example:
	//
	// 1551111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsTrial     *bool   `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// example:
	//
	// 204211111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cmgj00**11
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// cmgj00**11-prepay
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	// example:
	//
	// APP
	ProductType    *string                                           `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RelationalData *DescribeInstanceForIsvResponseBodyRelationalData `json:"RelationalData,omitempty" xml:"RelationalData,omitempty" type:"Struct"`
	// example:
	//
	// 6EF60BEC-****-****-****-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OPENED
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeInstanceForIsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceForIsvResponseBody) GetActiveAddress() *string {
	return s.ActiveAddress
}

func (s *DescribeInstanceForIsvResponseBody) GetAppJson() *string {
	return s.AppJson
}

func (s *DescribeInstanceForIsvResponseBody) GetAutoRenewal() *string {
	return s.AutoRenewal
}

func (s *DescribeInstanceForIsvResponseBody) GetBeganOn() *int64 {
	return s.BeganOn
}

func (s *DescribeInstanceForIsvResponseBody) GetComponentJson() *string {
	return s.ComponentJson
}

func (s *DescribeInstanceForIsvResponseBody) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeInstanceForIsvResponseBody) GetEndOn() *int64 {
	return s.EndOn
}

func (s *DescribeInstanceForIsvResponseBody) GetExtendJson() *string {
	return s.ExtendJson
}

func (s *DescribeInstanceForIsvResponseBody) GetHostJson() *string {
	return s.HostJson
}

func (s *DescribeInstanceForIsvResponseBody) GetImageJson() *string {
	return s.ImageJson
}

func (s *DescribeInstanceForIsvResponseBody) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeInstanceForIsvResponseBody) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *DescribeInstanceForIsvResponseBody) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstanceForIsvResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeInstanceForIsvResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceForIsvResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceForIsvResponseBody) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeInstanceForIsvResponseBody) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstanceForIsvResponseBody) GetRelationalData() *DescribeInstanceForIsvResponseBodyRelationalData {
	return s.RelationalData
}

func (s *DescribeInstanceForIsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceForIsvResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceForIsvResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeInstanceForIsvResponseBody) SetActiveAddress(v string) *DescribeInstanceForIsvResponseBody {
	s.ActiveAddress = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetAppJson(v string) *DescribeInstanceForIsvResponseBody {
	s.AppJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetAutoRenewal(v string) *DescribeInstanceForIsvResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetBeganOn(v int64) *DescribeInstanceForIsvResponseBody {
	s.BeganOn = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetComponentJson(v string) *DescribeInstanceForIsvResponseBody {
	s.ComponentJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetCreatedOn(v int64) *DescribeInstanceForIsvResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetEndOn(v int64) *DescribeInstanceForIsvResponseBody {
	s.EndOn = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetExtendJson(v string) *DescribeInstanceForIsvResponseBody {
	s.ExtendJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetHostJson(v string) *DescribeInstanceForIsvResponseBody {
	s.HostJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetImageJson(v string) *DescribeInstanceForIsvResponseBody {
	s.ImageJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetInstanceId(v int64) *DescribeInstanceForIsvResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetIsTrial(v bool) *DescribeInstanceForIsvResponseBody {
	s.IsTrial = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetLicenseCode(v string) *DescribeInstanceForIsvResponseBody {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetOrderId(v int64) *DescribeInstanceForIsvResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductCode(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductName(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductSkuCode(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductType(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetRelationalData(v *DescribeInstanceForIsvResponseBodyRelationalData) *DescribeInstanceForIsvResponseBody {
	s.RelationalData = v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetRequestId(v string) *DescribeInstanceForIsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetStatus(v string) *DescribeInstanceForIsvResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetSupplierName(v string) *DescribeInstanceForIsvResponseBody {
	s.SupplierName = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceForIsvResponseBodyRelationalData struct {
	// example:
	//
	// STARTED
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s DescribeInstanceForIsvResponseBodyRelationalData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceForIsvResponseBodyRelationalData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceForIsvResponseBodyRelationalData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeInstanceForIsvResponseBodyRelationalData) SetServiceStatus(v string) *DescribeInstanceForIsvResponseBodyRelationalData {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBodyRelationalData) Validate() error {
	return dara.Validate(s)
}
