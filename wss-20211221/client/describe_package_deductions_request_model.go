// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageDeductionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePackageDeductionsRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribePackageDeductionsRequest
	GetInstanceIds() []*string
	SetPackageIds(v []*string) *DescribePackageDeductionsRequest
	GetPackageIds() []*string
	SetPageNum(v int32) *DescribePackageDeductionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribePackageDeductionsRequest
	GetPageSize() *int32
	SetResourceType(v string) *DescribePackageDeductionsRequest
	GetResourceType() *string
	SetResourceTypes(v []*string) *DescribePackageDeductionsRequest
	GetResourceTypes() []*string
	SetStartTime(v int64) *DescribePackageDeductionsRequest
	GetStartTime() *int64
}

type DescribePackageDeductionsRequest struct {
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PackageIds  []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// CorePackage
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePackageDeductionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageDeductionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageDeductionsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePackageDeductionsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribePackageDeductionsRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DescribePackageDeductionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribePackageDeductionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePackageDeductionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePackageDeductionsRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribePackageDeductionsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePackageDeductionsRequest) SetEndTime(v int64) *DescribePackageDeductionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetInstanceIds(v []*string) *DescribePackageDeductionsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPackageIds(v []*string) *DescribePackageDeductionsRequest {
	s.PackageIds = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageNum(v int32) *DescribePackageDeductionsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetPageSize(v int32) *DescribePackageDeductionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceType(v string) *DescribePackageDeductionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePackageDeductionsRequest) SetResourceTypes(v []*string) *DescribePackageDeductionsRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribePackageDeductionsRequest) SetStartTime(v int64) *DescribePackageDeductionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePackageDeductionsRequest) Validate() error {
	return dara.Validate(s)
}
