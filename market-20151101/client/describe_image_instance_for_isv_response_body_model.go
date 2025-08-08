// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInstanceForIsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddress(v string) *DescribeImageInstanceForIsvResponseBody
	GetActiveAddress() *string
	SetAppJson(v string) *DescribeImageInstanceForIsvResponseBody
	GetAppJson() *string
	SetAutoRenewal(v string) *DescribeImageInstanceForIsvResponseBody
	GetAutoRenewal() *string
	SetBeganOn(v int64) *DescribeImageInstanceForIsvResponseBody
	GetBeganOn() *int64
	SetComponentJson(v string) *DescribeImageInstanceForIsvResponseBody
	GetComponentJson() *string
	SetConstraints(v string) *DescribeImageInstanceForIsvResponseBody
	GetConstraints() *string
	SetCreatedOn(v int64) *DescribeImageInstanceForIsvResponseBody
	GetCreatedOn() *int64
	SetEndOn(v int64) *DescribeImageInstanceForIsvResponseBody
	GetEndOn() *int64
	SetExtendJson(v string) *DescribeImageInstanceForIsvResponseBody
	GetExtendJson() *string
	SetHostJson(v string) *DescribeImageInstanceForIsvResponseBody
	GetHostJson() *string
	SetInstanceId(v int64) *DescribeImageInstanceForIsvResponseBody
	GetInstanceId() *int64
	SetIsTrial(v bool) *DescribeImageInstanceForIsvResponseBody
	GetIsTrial() *bool
	SetLicenseCode(v string) *DescribeImageInstanceForIsvResponseBody
	GetLicenseCode() *string
	SetModules(v []*DescribeImageInstanceForIsvResponseBodyModules) *DescribeImageInstanceForIsvResponseBody
	GetModules() []*DescribeImageInstanceForIsvResponseBodyModules
	SetOrderId(v int64) *DescribeImageInstanceForIsvResponseBody
	GetOrderId() *int64
	SetProductCode(v string) *DescribeImageInstanceForIsvResponseBody
	GetProductCode() *string
	SetProductName(v string) *DescribeImageInstanceForIsvResponseBody
	GetProductName() *string
	SetProductSkuCode(v string) *DescribeImageInstanceForIsvResponseBody
	GetProductSkuCode() *string
	SetProductType(v string) *DescribeImageInstanceForIsvResponseBody
	GetProductType() *string
	SetRelationalData(v *DescribeImageInstanceForIsvResponseBodyRelationalData) *DescribeImageInstanceForIsvResponseBody
	GetRelationalData() *DescribeImageInstanceForIsvResponseBodyRelationalData
	SetRequestId(v string) *DescribeImageInstanceForIsvResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeImageInstanceForIsvResponseBody
	GetStatus() *string
	SetSupplierName(v string) *DescribeImageInstanceForIsvResponseBody
	GetSupplierName() *string
}

type DescribeImageInstanceForIsvResponseBody struct {
	ActiveAddress  *string                                                `json:"ActiveAddress,omitempty" xml:"ActiveAddress,omitempty"`
	AppJson        *string                                                `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	AutoRenewal    *string                                                `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	BeganOn        *int64                                                 `json:"BeganOn,omitempty" xml:"BeganOn,omitempty"`
	ComponentJson  *string                                                `json:"ComponentJson,omitempty" xml:"ComponentJson,omitempty"`
	Constraints    *string                                                `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	CreatedOn      *int64                                                 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	EndOn          *int64                                                 `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	ExtendJson     *string                                                `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	HostJson       *string                                                `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	InstanceId     *int64                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsTrial        *bool                                                  `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	LicenseCode    *string                                                `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	Modules        []*DescribeImageInstanceForIsvResponseBodyModules      `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Repeated"`
	OrderId        *int64                                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductCode    *string                                                `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName    *string                                                `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductSkuCode *string                                                `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ProductType    *string                                                `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RelationalData *DescribeImageInstanceForIsvResponseBodyRelationalData `json:"RelationalData,omitempty" xml:"RelationalData,omitempty" type:"Struct"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName   *string                                                `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeImageInstanceForIsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvResponseBody) GetActiveAddress() *string {
	return s.ActiveAddress
}

func (s *DescribeImageInstanceForIsvResponseBody) GetAppJson() *string {
	return s.AppJson
}

func (s *DescribeImageInstanceForIsvResponseBody) GetAutoRenewal() *string {
	return s.AutoRenewal
}

func (s *DescribeImageInstanceForIsvResponseBody) GetBeganOn() *int64 {
	return s.BeganOn
}

func (s *DescribeImageInstanceForIsvResponseBody) GetComponentJson() *string {
	return s.ComponentJson
}

func (s *DescribeImageInstanceForIsvResponseBody) GetConstraints() *string {
	return s.Constraints
}

func (s *DescribeImageInstanceForIsvResponseBody) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeImageInstanceForIsvResponseBody) GetEndOn() *int64 {
	return s.EndOn
}

func (s *DescribeImageInstanceForIsvResponseBody) GetExtendJson() *string {
	return s.ExtendJson
}

func (s *DescribeImageInstanceForIsvResponseBody) GetHostJson() *string {
	return s.HostJson
}

func (s *DescribeImageInstanceForIsvResponseBody) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeImageInstanceForIsvResponseBody) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *DescribeImageInstanceForIsvResponseBody) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeImageInstanceForIsvResponseBody) GetModules() []*DescribeImageInstanceForIsvResponseBodyModules {
	return s.Modules
}

func (s *DescribeImageInstanceForIsvResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeImageInstanceForIsvResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeImageInstanceForIsvResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeImageInstanceForIsvResponseBody) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeImageInstanceForIsvResponseBody) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeImageInstanceForIsvResponseBody) GetRelationalData() *DescribeImageInstanceForIsvResponseBodyRelationalData {
	return s.RelationalData
}

func (s *DescribeImageInstanceForIsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageInstanceForIsvResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageInstanceForIsvResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeImageInstanceForIsvResponseBody) SetActiveAddress(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ActiveAddress = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetAppJson(v string) *DescribeImageInstanceForIsvResponseBody {
	s.AppJson = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetAutoRenewal(v string) *DescribeImageInstanceForIsvResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetBeganOn(v int64) *DescribeImageInstanceForIsvResponseBody {
	s.BeganOn = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetComponentJson(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ComponentJson = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetConstraints(v string) *DescribeImageInstanceForIsvResponseBody {
	s.Constraints = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetCreatedOn(v int64) *DescribeImageInstanceForIsvResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetEndOn(v int64) *DescribeImageInstanceForIsvResponseBody {
	s.EndOn = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetExtendJson(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ExtendJson = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetHostJson(v string) *DescribeImageInstanceForIsvResponseBody {
	s.HostJson = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetInstanceId(v int64) *DescribeImageInstanceForIsvResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetIsTrial(v bool) *DescribeImageInstanceForIsvResponseBody {
	s.IsTrial = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetLicenseCode(v string) *DescribeImageInstanceForIsvResponseBody {
	s.LicenseCode = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetModules(v []*DescribeImageInstanceForIsvResponseBodyModules) *DescribeImageInstanceForIsvResponseBody {
	s.Modules = v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetOrderId(v int64) *DescribeImageInstanceForIsvResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetProductCode(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetProductName(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetProductSkuCode(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetProductType(v string) *DescribeImageInstanceForIsvResponseBody {
	s.ProductType = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetRelationalData(v *DescribeImageInstanceForIsvResponseBodyRelationalData) *DescribeImageInstanceForIsvResponseBody {
	s.RelationalData = v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetRequestId(v string) *DescribeImageInstanceForIsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetStatus(v string) *DescribeImageInstanceForIsvResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) SetSupplierName(v string) *DescribeImageInstanceForIsvResponseBody {
	s.SupplierName = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageInstanceForIsvResponseBodyModules struct {
	Code       *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Id         *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties []*DescribeImageInstanceForIsvResponseBodyModulesProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s DescribeImageInstanceForIsvResponseBodyModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvResponseBodyModules) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) GetCode() *string {
	return s.Code
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) GetId() *string {
	return s.Id
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) GetName() *string {
	return s.Name
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) GetProperties() []*DescribeImageInstanceForIsvResponseBodyModulesProperties {
	return s.Properties
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) SetCode(v string) *DescribeImageInstanceForIsvResponseBodyModules {
	s.Code = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) SetId(v string) *DescribeImageInstanceForIsvResponseBodyModules {
	s.Id = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) SetName(v string) *DescribeImageInstanceForIsvResponseBodyModules {
	s.Name = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) SetProperties(v []*DescribeImageInstanceForIsvResponseBodyModulesProperties) *DescribeImageInstanceForIsvResponseBodyModules {
	s.Properties = v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModules) Validate() error {
	return dara.Validate(s)
}

type DescribeImageInstanceForIsvResponseBodyModulesProperties struct {
	DisplayUnit    *string                                                                   `json:"DisplayUnit,omitempty" xml:"DisplayUnit,omitempty"`
	Key            *string                                                                   `json:"Key,omitempty" xml:"Key,omitempty"`
	Name           *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyValues []*DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
	ShowType       *string                                                                   `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
}

func (s DescribeImageInstanceForIsvResponseBodyModulesProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvResponseBodyModulesProperties) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) GetDisplayUnit() *string {
	return s.DisplayUnit
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) GetKey() *string {
	return s.Key
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) GetName() *string {
	return s.Name
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) GetPropertyValues() []*DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	return s.PropertyValues
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) GetShowType() *string {
	return s.ShowType
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) SetDisplayUnit(v string) *DescribeImageInstanceForIsvResponseBodyModulesProperties {
	s.DisplayUnit = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) SetKey(v string) *DescribeImageInstanceForIsvResponseBodyModulesProperties {
	s.Key = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) SetName(v string) *DescribeImageInstanceForIsvResponseBodyModulesProperties {
	s.Name = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) SetPropertyValues(v []*DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) *DescribeImageInstanceForIsvResponseBodyModulesProperties {
	s.PropertyValues = v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) SetShowType(v string) *DescribeImageInstanceForIsvResponseBodyModulesProperties {
	s.ShowType = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesProperties) Validate() error {
	return dara.Validate(s)
}

type DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Max         *string `json:"Max,omitempty" xml:"Max,omitempty"`
	Min         *string `json:"Min,omitempty" xml:"Min,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Step        *string `json:"Step,omitempty" xml:"Step,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetMax() *string {
	return s.Max
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetMin() *string {
	return s.Min
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetRemark() *string {
	return s.Remark
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetStep() *string {
	return s.Step
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetType() *string {
	return s.Type
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) GetValue() *string {
	return s.Value
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetDisplayName(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.DisplayName = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetMax(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.Max = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetMin(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.Min = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetRemark(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.Remark = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetStep(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.Step = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetType(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.Type = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) SetValue(v string) *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues {
	s.Value = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyModulesPropertiesPropertyValues) Validate() error {
	return dara.Validate(s)
}

type DescribeImageInstanceForIsvResponseBodyRelationalData struct {
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s DescribeImageInstanceForIsvResponseBodyRelationalData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvResponseBodyRelationalData) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvResponseBodyRelationalData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeImageInstanceForIsvResponseBodyRelationalData) SetServiceStatus(v string) *DescribeImageInstanceForIsvResponseBodyRelationalData {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponseBodyRelationalData) Validate() error {
	return dara.Validate(s)
}
