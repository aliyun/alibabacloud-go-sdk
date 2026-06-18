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
	SetDescription(v string) *DescribeDBInstancesRequest
	GetDescription() *string
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
	// The description or remarks of the database.
	//
	// example:
	//
	// app-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dinga93c84f4d***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the instance must have a log engine.
	//
	// example:
	//
	// true
	MustHasCdc *bool `json:"MustHasCdc,omitempty" xml:"MustHasCdc,omitempty"`
	// The page number. Starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - 30
	//
	// - 50
	//
	// - 100.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyst47******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance edition. Valid values:
	//
	// - **enterprise**: Enterprise Edition.
	//
	// - **standard**: Standard Edition.
	//
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The list of tags.
	//
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

func (s *DescribeDBInstancesRequest) GetDescription() *string {
	return s.Description
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

func (s *DescribeDBInstancesRequest) SetDescription(v string) *DescribeDBInstancesRequest {
	s.Description = &v
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
