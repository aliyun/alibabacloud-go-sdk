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
	// An array that consists of the types of anomalous events.
	//
	// > If you leave the ParentTypeId parameter empty, anomalous event types are returned. If you set the ParentTypeId parameter, anomalous event subtypes under the specified anomalous event type are returned.
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
	// The code of the anomalous event type.
	//
	// example:
	//
	// 01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the anomalous event type.
	//
	// example:
	//
	// Anomalous permission usage,\\*\\*\\*\\*
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the anomalous event type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anomalous event type.
	//
	// example:
	//
	// Anomalous permission usage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array that consists of anomalous event subtypes.
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
	// The service to which the anomalous event detection rule applies. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	AdaptedProduct *string `json:"AdaptedProduct,omitempty" xml:"AdaptedProduct,omitempty"`
	// The code of the anomalous event subtype.
	//
	// example:
	//
	// 020008
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The code of the configuration.
	//
	// example:
	//
	// 0100**
	ConfigCode *string `json:"ConfigCode,omitempty" xml:"ConfigCode,omitempty"`
	// The content format of anomalous event detection rule. Valid values:
	//
	// 	- **0**: numeric values such as thresholds
	//
	// 	- **1**: text such as IP addresses
	//
	// example:
	//
	// 1
	ConfigContentType *int32 `json:"ConfigContentType,omitempty" xml:"ConfigContentType,omitempty"`
	// The description of the configuration.
	//
	// example:
	//
	// The period of time for which the permission is not used exceeds the threshold. The specified threshold is ${value} calendar days.
	ConfigDescription *string `json:"ConfigDescription,omitempty" xml:"ConfigDescription,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// 90
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The description of the anomalous event subtype.
	//
	// example:
	//
	// Inappropriate configuration-No protection for the MaxCompute sensitive project, \\*\\*\\*\\*
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of times that the anomalous event hits the anomalous event detection rule.
	//
	// example:
	//
	// 2
	EventHitCount *int32 `json:"EventHitCount,omitempty" xml:"EventHitCount,omitempty"`
	// The ID of the anomalous event subtype.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anomalous event subtype.
	//
	// example:
	//
	// Inappropriate configuration-No protection for the MaxCompute sensitive project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether detection is enabled for the anomalous event subtype. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
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
