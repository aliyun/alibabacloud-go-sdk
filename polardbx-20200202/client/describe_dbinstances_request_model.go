// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbVersion(v string) *DescribeDBInstancesRequest
	GetDbVersion() *string
	SetInstanceId(v string) *DescribeDBInstancesRequest
	GetInstanceId() *string
	SetMustHasCdc(v bool) *DescribeDBInstancesRequest
	GetMustHasCdc() *bool
	SetPageNumber(v int32) *DescribeDBInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesRequest
	GetResourceGroupId() *string
	SetSeries(v string) *DescribeDBInstancesRequest
	GetSeries() *string
	SetTags(v string) *DescribeDBInstancesRequest
	GetTags() *string
}

type DescribeDBInstancesRequest struct {
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// dinga93c84f4d***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MustHasCdc *bool   `json:"MustHasCdc,omitempty" xml:"MustHasCdc,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyst47hjw***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// [{\\"TagKey\\":\\"test\\",\\"TagValue\\":\\"test-value\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeDBInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDBInstancesRequest) GetMustHasCdc() *bool {
	return s.MustHasCdc
}

func (s *DescribeDBInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesRequest) GetSeries() *string {
	return s.Series
}

func (s *DescribeDBInstancesRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeDBInstancesRequest) SetDbVersion(v string) *DescribeDBInstancesRequest {
	s.DbVersion = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceId(v string) *DescribeDBInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetMustHasCdc(v bool) *DescribeDBInstancesRequest {
	s.MustHasCdc = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetSeries(v string) *DescribeDBInstancesRequest {
	s.Series = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTags(v string) *DescribeDBInstancesRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDBInstancesRequest) Validate() error {
	return dara.Validate(s)
}
