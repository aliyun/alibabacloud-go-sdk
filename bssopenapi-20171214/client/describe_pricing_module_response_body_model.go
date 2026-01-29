// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePricingModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePricingModuleResponseBody
	GetCode() *string
	SetData(v *DescribePricingModuleResponseBodyData) *DescribePricingModuleResponseBody
	GetData() *DescribePricingModuleResponseBodyData
	SetMessage(v string) *DescribePricingModuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePricingModuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePricingModuleResponseBody
	GetSuccess() *bool
}

type DescribePricingModuleResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribePricingModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// This API is not applicable for caller.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C40A8EE0-8084-49FE-B66E-5E1C3B6AE025
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePricingModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePricingModuleResponseBody) GetData() *DescribePricingModuleResponseBodyData {
	return s.Data
}

func (s *DescribePricingModuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePricingModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePricingModuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePricingModuleResponseBody) SetCode(v string) *DescribePricingModuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetData(v *DescribePricingModuleResponseBodyData) *DescribePricingModuleResponseBody {
	s.Data = v
	return s
}

func (s *DescribePricingModuleResponseBody) SetMessage(v string) *DescribePricingModuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetRequestId(v string) *DescribePricingModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetSuccess(v bool) *DescribePricingModuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePricingModuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyData struct {
	// The module attributes.
	AttributeList *DescribePricingModuleResponseBodyDataAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Struct"`
	// The pricing information of modules.
	ModuleList *DescribePricingModuleResponseBodyDataModuleList `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Struct"`
}

func (s DescribePricingModuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyData) GetAttributeList() *DescribePricingModuleResponseBodyDataAttributeList {
	return s.AttributeList
}

func (s *DescribePricingModuleResponseBodyData) GetModuleList() *DescribePricingModuleResponseBodyDataModuleList {
	return s.ModuleList
}

func (s *DescribePricingModuleResponseBodyData) SetAttributeList(v *DescribePricingModuleResponseBodyDataAttributeList) *DescribePricingModuleResponseBodyData {
	s.AttributeList = v
	return s
}

func (s *DescribePricingModuleResponseBodyData) SetModuleList(v *DescribePricingModuleResponseBodyDataModuleList) *DescribePricingModuleResponseBodyData {
	s.ModuleList = v
	return s
}

func (s *DescribePricingModuleResponseBodyData) Validate() error {
	if s.AttributeList != nil {
		if err := s.AttributeList.Validate(); err != nil {
			return err
		}
	}
	if s.ModuleList != nil {
		if err := s.ModuleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyDataAttributeList struct {
	Attribute []*DescribePricingModuleResponseBodyDataAttributeListAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataAttributeList) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeList) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeList) GetAttribute() []*DescribePricingModuleResponseBodyDataAttributeListAttribute {
	return s.Attribute
}

func (s *DescribePricingModuleResponseBodyDataAttributeList) SetAttribute(v []*DescribePricingModuleResponseBodyDataAttributeListAttribute) *DescribePricingModuleResponseBodyDataAttributeList {
	s.Attribute = v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeList) Validate() error {
	if s.Attribute != nil {
		for _, item := range s.Attribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyDataAttributeListAttribute struct {
	// The code of the attribute.
	//
	// example:
	//
	// DBInstanceStorage
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the attribute.
	//
	// example:
	//
	// Capacity
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the attribute.
	//
	// example:
	//
	// GB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The attribute values.
	Values *DescribePricingModuleResponseBodyDataAttributeListAttributeValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttribute) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) GetCode() *string {
	return s.Code
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) GetName() *string {
	return s.Name
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) GetUnit() *string {
	return s.Unit
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) GetValues() *DescribePricingModuleResponseBodyDataAttributeListAttributeValues {
	return s.Values
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetCode(v string) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Code = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetName(v string) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Name = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetUnit(v string) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Unit = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetValues(v *DescribePricingModuleResponseBodyDataAttributeListAttributeValues) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Values = v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) Validate() error {
	if s.Values != nil {
		if err := s.Values.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyDataAttributeListAttributeValues struct {
	AttributeValue []*DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValues) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValues) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValues) GetAttributeValue() []*DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	return s.AttributeValue
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValues) SetAttributeValue(v []*DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) *DescribePricingModuleResponseBodyDataAttributeListAttributeValues {
	s.AttributeValue = v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValues) Validate() error {
	if s.AttributeValue != nil {
		for _, item := range s.AttributeValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue struct {
	// The attribute value that corresponds to the module code.
	//
	// example:
	//
	// 2 Cores and 4 GB Memory (Basic Edition)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the module values.
	//
	// example:
	//
	// Connections: 4,000 IOPS is related to storage space
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the attribute value that corresponds to the module code. Valid values:
	//
	// 	- single_float: single value
	//
	// 	- range_float: range value
	//
	// example:
	//
	// single_string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute value that corresponds to the module code.
	//
	// >  If the Type parameter is set to range_float, the valid values of this parameter range from 1024 to 1024000. A value of 1024 indicates that the step size is 1024.
	//
	// example:
	//
	// mysql.n2.medium.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) GetName() *string {
	return s.Name
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) GetRemark() *string {
	return s.Remark
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) GetType() *string {
	return s.Type
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) GetValue() *string {
	return s.Value
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetName(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Name = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetRemark(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Remark = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetType(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Type = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetValue(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Value = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) Validate() error {
	return dara.Validate(s)
}

type DescribePricingModuleResponseBodyDataModuleList struct {
	Module []*DescribePricingModuleResponseBodyDataModuleListModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataModuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataModuleList) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataModuleList) GetModule() []*DescribePricingModuleResponseBodyDataModuleListModule {
	return s.Module
}

func (s *DescribePricingModuleResponseBodyDataModuleList) SetModule(v []*DescribePricingModuleResponseBodyDataModuleListModule) *DescribePricingModuleResponseBodyDataModuleList {
	s.Module = v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleList) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyDataModuleListModule struct {
	ConfigList *DescribePricingModuleResponseBodyDataModuleListModuleConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Struct"`
	// The currency. Default value: CNY.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The code of the pricing module.
	//
	// example:
	//
	// InstanceType
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The name of the pricing module.
	//
	// example:
	//
	// Instance
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The price type. Valid values:
	//
	// 	- Usage: usage price
	//
	// 	- Hour: hourly price
	//
	// 	- Day: daily price
	//
	// 	- Week: weekly price
	//
	// 	- Month: monthly price
	//
	// 	- Year: annual price
	//
	// example:
	//
	// Month
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
}

func (s DescribePricingModuleResponseBodyDataModuleListModule) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataModuleListModule) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) GetConfigList() *DescribePricingModuleResponseBodyDataModuleListModuleConfigList {
	return s.ConfigList
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) GetPriceType() *string {
	return s.PriceType
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetConfigList(v *DescribePricingModuleResponseBodyDataModuleListModuleConfigList) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.ConfigList = v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetCurrency(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.Currency = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetModuleCode(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.ModuleCode = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetModuleName(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.ModuleName = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetPriceType(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.PriceType = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) Validate() error {
	if s.ConfigList != nil {
		if err := s.ConfigList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePricingModuleResponseBodyDataModuleListModuleConfigList struct {
	ConfigList []*string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataModuleListModuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataModuleListModuleConfigList) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataModuleListModuleConfigList) GetConfigList() []*string {
	return s.ConfigList
}

func (s *DescribePricingModuleResponseBodyDataModuleListModuleConfigList) SetConfigList(v []*string) *DescribePricingModuleResponseBodyDataModuleListModuleConfigList {
	s.ConfigList = v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModuleConfigList) Validate() error {
	return dara.Validate(s)
}
