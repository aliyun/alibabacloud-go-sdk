// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePackageProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResourcePackageProductResponseBody
	GetCode() *string
	SetData(v *DescribeResourcePackageProductResponseBodyData) *DescribeResourcePackageProductResponseBody
	GetData() *DescribeResourcePackageProductResponseBodyData
	SetMessage(v string) *DescribeResourcePackageProductResponseBody
	GetMessage() *string
	SetOrderId(v int64) *DescribeResourcePackageProductResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *DescribeResourcePackageProductResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResourcePackageProductResponseBody
	GetSuccess() *bool
}

type DescribeResourcePackageProductResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeResourcePackageProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 72353765387
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BBEF51A3-E933-4F40-A534-C673CBDB9C80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeResourcePackageProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResourcePackageProductResponseBody) GetData() *DescribeResourcePackageProductResponseBodyData {
	return s.Data
}

func (s *DescribeResourcePackageProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourcePackageProductResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeResourcePackageProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourcePackageProductResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResourcePackageProductResponseBody) SetCode(v string) *DescribeResourcePackageProductResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetData(v *DescribeResourcePackageProductResponseBodyData) *DescribeResourcePackageProductResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetMessage(v string) *DescribeResourcePackageProductResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetOrderId(v int64) *DescribeResourcePackageProductResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetRequestId(v string) *DescribeResourcePackageProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetSuccess(v bool) *DescribeResourcePackageProductResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyData struct {
	// The details about the resource plans.
	ResourcePackages *DescribeResourcePackageProductResponseBodyDataResourcePackages `json:"ResourcePackages,omitempty" xml:"ResourcePackages,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyData) GetResourcePackages() *DescribeResourcePackageProductResponseBodyDataResourcePackages {
	return s.ResourcePackages
}

func (s *DescribeResourcePackageProductResponseBodyData) SetResourcePackages(v *DescribeResourcePackageProductResponseBodyDataResourcePackages) *DescribeResourcePackageProductResponseBodyData {
	s.ResourcePackages = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyData) Validate() error {
	if s.ResourcePackages != nil {
		if err := s.ResourcePackages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackages struct {
	ResourcePackage []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage `json:"ResourcePackage,omitempty" xml:"ResourcePackage,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackages) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackages) GetResourcePackage() []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	return s.ResourcePackage
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackages) SetResourcePackage(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) *DescribeResourcePackageProductResponseBodyDataResourcePackages {
	s.ResourcePackage = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackages) Validate() error {
	if s.ResourcePackage != nil {
		for _, item := range s.ResourcePackage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage struct {
	// The name of the resource plan.
	//
	// example:
	//
	// Object Storage Service (OSS) resource plan (monthly)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The types of the resource plans.
	PackageTypes *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes `json:"PackageTypes,omitempty" xml:"PackageTypes,omitempty" type:"Struct"`
	// The code of the service.
	//
	// example:
	//
	// ossbag
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// ossbag
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) GetName() *string {
	return s.Name
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) GetPackageTypes() *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes {
	return s.PackageTypes
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetPackageTypes(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.PackageTypes = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetProductCode(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.ProductCode = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetProductType(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.ProductType = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) Validate() error {
	if s.PackageTypes != nil {
		if err := s.PackageTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes struct {
	PackageType []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType `json:"PackageType,omitempty" xml:"PackageType,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) GetPackageType() []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	return s.PackageType
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) SetPackageType(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes {
	s.PackageType = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) Validate() error {
	if s.PackageType != nil {
		for _, item := range s.PackageType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType struct {
	// The code of the resource plan.
	//
	// example:
	//
	// FPT_ossbag_deadlineAcc_CdnOut_common_sz
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the resource plan type.
	//
	// example:
	//
	// Back-to-origin traffic plan - China (Shenzhen)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The properties of the resource plan.
	Properties *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// The specifications of the resource plan.
	Specifications *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) GetCode() *string {
	return s.Code
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) GetName() *string {
	return s.Name
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) GetProperties() *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties {
	return s.Properties
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) GetSpecifications() *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications {
	return s.Specifications
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetCode(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Code = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetProperties(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Properties = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetSpecifications(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Specifications = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) Validate() error {
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.Specifications != nil {
		if err := s.Specifications.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties struct {
	Property []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) GetProperty() []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
	return s.Property
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) SetProperty(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties {
	s.Property = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) Validate() error {
	if s.Property != nil {
		for _, item := range s.Property {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty struct {
	// The name of the property.
	//
	// example:
	//
	// region
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// cn-shenzhen
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) GetName() *string {
	return s.Name
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) GetValue() *string {
	return s.Value
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) SetValue(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
	s.Value = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) Validate() error {
	return dara.Validate(s)
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications struct {
	Specification []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification `json:"Specification,omitempty" xml:"Specification,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) GetSpecification() []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	return s.Specification
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) SetSpecification(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications {
	s.Specification = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) Validate() error {
	if s.Specification != nil {
		for _, item := range s.Specification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification struct {
	// The validity periods available for the resource plan.
	AvailableDurations *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations `json:"AvailableDurations,omitempty" xml:"AvailableDurations,omitempty" type:"Struct"`
	// The name of the specification.
	//
	// example:
	//
	// 1TB
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the specification.
	//
	// example:
	//
	// 1024
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) GetAvailableDurations() *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations {
	return s.AvailableDurations
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) GetName() *string {
	return s.Name
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) GetValue() *string {
	return s.Value
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetAvailableDurations(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	s.AvailableDurations = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetValue(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	s.Value = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) Validate() error {
	if s.AvailableDurations != nil {
		if err := s.AvailableDurations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations struct {
	AvailableDuration []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration `json:"AvailableDuration,omitempty" xml:"AvailableDuration,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) GetAvailableDuration() []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	return s.AvailableDuration
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) SetAvailableDuration(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations {
	s.AvailableDuration = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) Validate() error {
	if s.AvailableDuration != nil {
		for _, item := range s.AvailableDuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration struct {
	// The name of the validity period.
	//
	// example:
	//
	// 6 Month
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the validity period for the resource plan. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the validity period.
	//
	// example:
	//
	// 6
	Value *int32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) GetName() *string {
	return s.Name
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) GetUnit() *string {
	return s.Unit
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) GetValue() *int32 {
	return s.Value
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetUnit(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	s.Unit = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetValue(v int32) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	s.Value = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) Validate() error {
	return dara.Validate(s)
}
