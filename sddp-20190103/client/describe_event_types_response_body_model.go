// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventTypeList(v []*DescribeEventTypesResponseBodyEventTypeList) *DescribeEventTypesResponseBody
	GetEventTypeList() []*DescribeEventTypesResponseBodyEventTypeList
	SetRequestId(v string) *DescribeEventTypesResponseBody
	GetRequestId() *string
}

type DescribeEventTypesResponseBody struct {
	// The list of anomalous activity types.
	//
	// > If ParentTypeId is empty, the parent anomalous activity types are returned. If ParentTypeId is not empty, the child anomalous activity types are returned.
	EventTypeList []*DescribeEventTypesResponseBodyEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEventTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBody) GetEventTypeList() []*DescribeEventTypesResponseBodyEventTypeList {
	return s.EventTypeList
}

func (s *DescribeEventTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventTypesResponseBody) SetEventTypeList(v []*DescribeEventTypesResponseBodyEventTypeList) *DescribeEventTypesResponseBody {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventTypesResponseBody) SetRequestId(v string) *DescribeEventTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventTypesResponseBody) Validate() error {
	if s.EventTypeList != nil {
		for _, item := range s.EventTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventTypesResponseBodyEventTypeList struct {
	// The code of the parent anomalous activity type.
	//
	// example:
	//
	// 01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the parent anomalous activity type.
	//
	// example:
	//
	// Permission usage anomaly, ****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the parent anomalous activity type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the parent anomalous activity type.
	//
	// example:
	//
	// Permission usage anomaly
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of child anomalous activity types.
	SubTypeList []*DescribeEventTypesResponseBodyEventTypeListSubTypeList `json:"SubTypeList,omitempty" xml:"SubTypeList,omitempty" type:"Repeated"`
}

func (s DescribeEventTypesResponseBodyEventTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTypesResponseBodyEventTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBodyEventTypeList) GetCode() *string {
	return s.Code
}

func (s *DescribeEventTypesResponseBodyEventTypeList) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventTypesResponseBodyEventTypeList) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventTypesResponseBodyEventTypeList) GetName() *string {
	return s.Name
}

func (s *DescribeEventTypesResponseBodyEventTypeList) GetSubTypeList() []*DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	return s.SubTypeList
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetCode(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Code = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetDescription(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Description = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetId(v int64) *DescribeEventTypesResponseBodyEventTypeList {
	s.Id = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetName(v string) *DescribeEventTypesResponseBodyEventTypeList {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) SetSubTypeList(v []*DescribeEventTypesResponseBodyEventTypeListSubTypeList) *DescribeEventTypesResponseBodyEventTypeList {
	s.SubTypeList = v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeList) Validate() error {
	if s.SubTypeList != nil {
		for _, item := range s.SubTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventTypesResponseBodyEventTypeListSubTypeList struct {
	// The products to which the rule applies, including MaxCompute, OSS, AnalyticDB for MySQL, Tablestore, and ApsaraDB RDS.
	//
	// example:
	//
	// RDS
	AdaptedProduct *string `json:"AdaptedProduct,omitempty" xml:"AdaptedProduct,omitempty"`
	// The code of the child anomalous activity type.
	//
	// example:
	//
	// 020008
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration code.
	//
	// example:
	//
	// 0100**
	ConfigCode *string `json:"ConfigCode,omitempty" xml:"ConfigCode,omitempty"`
	// The format of the rule item. Valid values:
	//
	// - **0**: numeric (such as a threshold).
	//
	// - **1**: text (such as an IP address).
	//
	// example:
	//
	// 1
	ConfigContentType *int32 `json:"ConfigContentType,omitempty" xml:"ConfigContentType,omitempty"`
	// The configuration description.
	//
	// example:
	//
	// Permission idle period exceeds threshold: current threshold is defined as 7 natural days
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	// The configuration value.
	//
	// example:
	//
	// 90
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The description of the child anomalous activity type.
	//
	// example:
	//
	// Configuration error - MaxCompute sensitive project not protected，****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of times the rule is hit.
	//
	// example:
	//
	// 2
	EventHitCount *int32 `json:"EventHitCount,omitempty" xml:"EventHitCount,omitempty"`
	// The unique ID of the child anomalous activity type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the child anomalous activity type.
	//
	// example:
	//
	// Configuration error - MaxCompute sensitive project not protected
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detection feature of Data Security Center (DSC) for the child anomalous activity type. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventTypesResponseBodyEventTypeListSubTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTypesResponseBodyEventTypeListSubTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetAdaptedProduct() *string {
	return s.AdaptedProduct
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetCode() *string {
	return s.Code
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetConfigCode() *string {
	return s.ConfigCode
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetConfigContentType() *int32 {
	return s.ConfigContentType
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetConfigDescription() *string {
	return s.ConfigDescription
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetEventHitCount() *int32 {
	return s.EventHitCount
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetName() *string {
	return s.Name
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetAdaptedProduct(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.AdaptedProduct = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetCode(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Code = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigCode(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigCode = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigContentType(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigContentType = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigDescription(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigDescription = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetConfigValue(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.ConfigValue = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetDescription(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Description = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetEventHitCount(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.EventHitCount = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetId(v int64) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Id = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetName(v string) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Name = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) SetStatus(v int32) *DescribeEventTypesResponseBodyEventTypeListSubTypeList {
	s.Status = &v
	return s
}

func (s *DescribeEventTypesResponseBodyEventTypeListSubTypeList) Validate() error {
	return dara.Validate(s)
}
