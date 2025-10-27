// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudProductFieldStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupedFields(v *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) *DescribeCloudProductFieldStatisticsResponseBody
	GetGroupedFields() *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields
	SetRequestId(v string) *DescribeCloudProductFieldStatisticsResponseBody
	GetRequestId() *string
}

type DescribeCloudProductFieldStatisticsResponseBody struct {
	// The statistics of cloud services that are protected by Security Center.
	GroupedFields *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudProductFieldStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) GetGroupedFields() *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	return s.GroupedFields
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) SetGroupedFields(v *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) *DescribeCloudProductFieldStatisticsResponseBody {
	s.GroupedFields = v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) SetRequestId(v string) *DescribeCloudProductFieldStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) Validate() error {
	if s.GroupedFields != nil {
		if err := s.GroupedFields.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudProductFieldStatisticsResponseBodyGroupedFields struct {
	// The statistics of different types of assets. **MachineType*	- indicates the type of the asset. **Count*	- indicates the number of assets of a specific type.
	//
	// Valid values of **MachineType**:
	//
	// 	- **1**: Server Load Balancer (SLB) instance
	//
	// 	- **2**: NAT gateway instance
	//
	// 	- **3**: ApsaraDB RDS instance
	//
	// 	- **4**: ApsaraDB for MongoDB instance
	//
	// example:
	//
	// [{"MachineType":1,"Count":11}]
	CategoryCount *string `json:"CategoryCount,omitempty" xml:"CategoryCount,omitempty"`
	// The total number of cloud services that are protected by Security Center.
	//
	// example:
	//
	// 100
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The number of cloud services that are at risk.
	//
	// example:
	//
	// 90
	RiskInstanceCount *int32 `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
}

func (s DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) GetCategoryCount() *string {
	return s.CategoryCount
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) GetRiskInstanceCount() *int32 {
	return s.RiskInstanceCount
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) SetCategoryCount(v string) *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	s.CategoryCount = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) SetInstanceCount(v int32) *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	s.InstanceCount = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) SetRiskInstanceCount(v int32) *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) Validate() error {
	return dara.Validate(s)
}
