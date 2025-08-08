// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddress(v string) *DescribeInstanceResponseBody
	GetActiveAddress() *string
	SetAppJson(v string) *DescribeInstanceResponseBody
	GetAppJson() *string
	SetAutoRenewal(v string) *DescribeInstanceResponseBody
	GetAutoRenewal() *string
	SetBeganOn(v int64) *DescribeInstanceResponseBody
	GetBeganOn() *int64
	SetComponentJson(v string) *DescribeInstanceResponseBody
	GetComponentJson() *string
	SetConstraints(v string) *DescribeInstanceResponseBody
	GetConstraints() *string
	SetCreatedOn(v int64) *DescribeInstanceResponseBody
	GetCreatedOn() *int64
	SetEndOn(v int64) *DescribeInstanceResponseBody
	GetEndOn() *int64
	SetExtendJson(v string) *DescribeInstanceResponseBody
	GetExtendJson() *string
	SetHostJson(v string) *DescribeInstanceResponseBody
	GetHostJson() *string
	SetInstanceId(v int64) *DescribeInstanceResponseBody
	GetInstanceId() *int64
	SetIsTrial(v bool) *DescribeInstanceResponseBody
	GetIsTrial() *bool
	SetLicenseCode(v string) *DescribeInstanceResponseBody
	GetLicenseCode() *string
	SetModules(v *DescribeInstanceResponseBodyModules) *DescribeInstanceResponseBody
	GetModules() *DescribeInstanceResponseBodyModules
	SetOrderId(v int64) *DescribeInstanceResponseBody
	GetOrderId() *int64
	SetProductCode(v string) *DescribeInstanceResponseBody
	GetProductCode() *string
	SetProductName(v string) *DescribeInstanceResponseBody
	GetProductName() *string
	SetProductSkuCode(v string) *DescribeInstanceResponseBody
	GetProductSkuCode() *string
	SetProductType(v string) *DescribeInstanceResponseBody
	GetProductType() *string
	SetRelationalData(v *DescribeInstanceResponseBodyRelationalData) *DescribeInstanceResponseBody
	GetRelationalData() *DescribeInstanceResponseBodyRelationalData
	SetStatus(v string) *DescribeInstanceResponseBody
	GetStatus() *string
	SetSupplierName(v string) *DescribeInstanceResponseBody
	GetSupplierName() *string
}

type DescribeInstanceResponseBody struct {
	ActiveAddress *string `json:"ActiveAddress,omitempty" xml:"ActiveAddress,omitempty"`
	// example:
	//
	// {"frontEndUrl":"https://****.aliyundoc.com","password":"Sjtv***","adminUrl":"https://****.aliyundoc.com","username":"aliyun***"}
	AppJson     *string `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	AutoRenewal *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
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
	// {}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// example:
	//
	// 1570634018000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// example:
	//
	// 1602259200000
	EndOn      *int64  `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	ExtendJson *string `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	// example:
	//
	// {"password":"***","ip":"118.31.***.41","innerIp":"118.31.***.41","region":"","username":"***","beianInfo":""}
	HostJson *string `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	// example:
	//
	// 1551111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsTrial     *bool                                `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	LicenseCode *string                              `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	Modules     *DescribeInstanceResponseBodyModules `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Struct"`
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
	ProductType    *string                                     `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RelationalData *DescribeInstanceResponseBodyRelationalData `json:"RelationalData,omitempty" xml:"RelationalData,omitempty" type:"Struct"`
	// example:
	//
	// OPENED
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetActiveAddress() *string {
	return s.ActiveAddress
}

func (s *DescribeInstanceResponseBody) GetAppJson() *string {
	return s.AppJson
}

func (s *DescribeInstanceResponseBody) GetAutoRenewal() *string {
	return s.AutoRenewal
}

func (s *DescribeInstanceResponseBody) GetBeganOn() *int64 {
	return s.BeganOn
}

func (s *DescribeInstanceResponseBody) GetComponentJson() *string {
	return s.ComponentJson
}

func (s *DescribeInstanceResponseBody) GetConstraints() *string {
	return s.Constraints
}

func (s *DescribeInstanceResponseBody) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeInstanceResponseBody) GetEndOn() *int64 {
	return s.EndOn
}

func (s *DescribeInstanceResponseBody) GetExtendJson() *string {
	return s.ExtendJson
}

func (s *DescribeInstanceResponseBody) GetHostJson() *string {
	return s.HostJson
}

func (s *DescribeInstanceResponseBody) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBody) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *DescribeInstanceResponseBody) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstanceResponseBody) GetModules() *DescribeInstanceResponseBodyModules {
	return s.Modules
}

func (s *DescribeInstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeInstanceResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceResponseBody) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeInstanceResponseBody) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstanceResponseBody) GetRelationalData() *DescribeInstanceResponseBodyRelationalData {
	return s.RelationalData
}

func (s *DescribeInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeInstanceResponseBody) SetActiveAddress(v string) *DescribeInstanceResponseBody {
	s.ActiveAddress = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetAppJson(v string) *DescribeInstanceResponseBody {
	s.AppJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetAutoRenewal(v string) *DescribeInstanceResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetBeganOn(v int64) *DescribeInstanceResponseBody {
	s.BeganOn = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetComponentJson(v string) *DescribeInstanceResponseBody {
	s.ComponentJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetConstraints(v string) *DescribeInstanceResponseBody {
	s.Constraints = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedOn(v int64) *DescribeInstanceResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEndOn(v int64) *DescribeInstanceResponseBody {
	s.EndOn = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExtendJson(v string) *DescribeInstanceResponseBody {
	s.ExtendJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetHostJson(v string) *DescribeInstanceResponseBody {
	s.HostJson = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v int64) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsTrial(v bool) *DescribeInstanceResponseBody {
	s.IsTrial = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetLicenseCode(v string) *DescribeInstanceResponseBody {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModules(v *DescribeInstanceResponseBodyModules) *DescribeInstanceResponseBody {
	s.Modules = v
	return s
}

func (s *DescribeInstanceResponseBody) SetOrderId(v int64) *DescribeInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductCode(v string) *DescribeInstanceResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductName(v string) *DescribeInstanceResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductSkuCode(v string) *DescribeInstanceResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProductType(v string) *DescribeInstanceResponseBody {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRelationalData(v *DescribeInstanceResponseBodyRelationalData) *DescribeInstanceResponseBody {
	s.RelationalData = v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSupplierName(v string) *DescribeInstanceResponseBody {
	s.SupplierName = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyModules struct {
	Module []*DescribeInstanceResponseBodyModulesModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyModules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModules) GetModule() []*DescribeInstanceResponseBodyModulesModule {
	return s.Module
}

func (s *DescribeInstanceResponseBodyModules) SetModule(v []*DescribeInstanceResponseBodyModulesModule) *DescribeInstanceResponseBodyModules {
	s.Module = v
	return s
}

func (s *DescribeInstanceResponseBodyModules) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyModulesModule struct {
	// example:
	//
	// package_config
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 101*********026
	Id         *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties *DescribeInstanceResponseBodyModulesModuleProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
}

func (s DescribeInstanceResponseBodyModulesModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModule) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModule) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceResponseBodyModulesModule) GetId() *string {
	return s.Id
}

func (s *DescribeInstanceResponseBodyModulesModule) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyModulesModule) GetProperties() *DescribeInstanceResponseBodyModulesModuleProperties {
	return s.Properties
}

func (s *DescribeInstanceResponseBodyModulesModule) SetCode(v string) *DescribeInstanceResponseBodyModulesModule {
	s.Code = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) SetId(v string) *DescribeInstanceResponseBodyModulesModule {
	s.Id = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) SetName(v string) *DescribeInstanceResponseBodyModulesModule {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) SetProperties(v *DescribeInstanceResponseBodyModulesModuleProperties) *DescribeInstanceResponseBodyModulesModule {
	s.Properties = v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModule) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyModulesModuleProperties struct {
	Property []*DescribeInstanceResponseBodyModulesModulePropertiesProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyModulesModuleProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModuleProperties) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModuleProperties) GetProperty() []*DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	return s.Property
}

func (s *DescribeInstanceResponseBodyModulesModuleProperties) SetProperty(v []*DescribeInstanceResponseBodyModulesModulePropertiesProperty) *DescribeInstanceResponseBodyModulesModuleProperties {
	s.Property = v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModuleProperties) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyModulesModulePropertiesProperty struct {
	// example:
	//
	// 12
	DisplayUnit *string `json:"DisplayUnit,omitempty" xml:"DisplayUnit,omitempty"`
	// example:
	//
	// 12
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 12
	Name           *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyValues *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Struct"`
	// example:
	//
	// 12
	ShowType *string `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesProperty) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesProperty) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) GetDisplayUnit() *string {
	return s.DisplayUnit
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) GetPropertyValues() *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	return s.PropertyValues
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) GetShowType() *string {
	return s.ShowType
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetDisplayUnit(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.DisplayUnit = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetKey(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetName(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetPropertyValues(v *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.PropertyValues = v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) SetShowType(v string) *DescribeInstanceResponseBodyModulesModulePropertiesProperty {
	s.ShowType = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesProperty) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues struct {
	PropertyValue []*DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) GetPropertyValue() []*DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	return s.PropertyValue
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) SetPropertyValue(v []*DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	s.PropertyValue = v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue struct {
	// example:
	//
	// 12
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 12
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 12
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// 12
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 12
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// example:
	//
	// 12
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 12
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetMax() *string {
	return s.Max
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetMin() *string {
	return s.Min
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetStep() *string {
	return s.Step
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetDisplayName(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.DisplayName = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMax(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Max = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMin(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Min = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetRemark(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetStep(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Step = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetType(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetValue(v string) *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Value = &v
	return s
}

func (s *DescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyRelationalData struct {
	// example:
	//
	// STARTED
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s DescribeInstanceResponseBodyRelationalData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyRelationalData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyRelationalData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeInstanceResponseBodyRelationalData) SetServiceStatus(v string) *DescribeInstanceResponseBodyRelationalData {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyRelationalData) Validate() error {
	return dara.Validate(s)
}
