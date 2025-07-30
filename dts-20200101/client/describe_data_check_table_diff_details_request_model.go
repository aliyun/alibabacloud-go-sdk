// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckTableDiffDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckType(v int32) *DescribeDataCheckTableDiffDetailsRequest
	GetCheckType() *int32
	SetDbName(v string) *DescribeDataCheckTableDiffDetailsRequest
	GetDbName() *string
	SetDtsJobId(v string) *DescribeDataCheckTableDiffDetailsRequest
	GetDtsJobId() *string
	SetPageNumber(v int64) *DescribeDataCheckTableDiffDetailsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDataCheckTableDiffDetailsRequest
	GetPageSize() *int64
	SetResourceGroupId(v string) *DescribeDataCheckTableDiffDetailsRequest
	GetResourceGroupId() *string
	SetTbName(v string) *DescribeDataCheckTableDiffDetailsRequest
	GetTbName() *string
}

type DescribeDataCheckTableDiffDetailsRequest struct {
	// The data verification method. Valid values:
	//
	// 	- **1**: full data verification.
	//
	// 	- **2**: incremental data verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The name of the database to which the table that contains inconsistent data belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// db_dtstest
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the data migration or data synchronization task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// xd4e4xb419q****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The page number of the page to return. The value must be an integer greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the table that contains inconsistent data exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_person
	TbName *string `json:"TbName,omitempty" xml:"TbName,omitempty"`
}

func (s DescribeDataCheckTableDiffDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckTableDiffDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDataCheckTableDiffDetailsRequest) GetTbName() *string {
	return s.TbName
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetCheckType(v int32) *DescribeDataCheckTableDiffDetailsRequest {
	s.CheckType = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetDbName(v string) *DescribeDataCheckTableDiffDetailsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetDtsJobId(v string) *DescribeDataCheckTableDiffDetailsRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetPageNumber(v int64) *DescribeDataCheckTableDiffDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetPageSize(v int64) *DescribeDataCheckTableDiffDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetResourceGroupId(v string) *DescribeDataCheckTableDiffDetailsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) SetTbName(v string) *DescribeDataCheckTableDiffDetailsRequest {
	s.TbName = &v
	return s
}

func (s *DescribeDataCheckTableDiffDetailsRequest) Validate() error {
	return dara.Validate(s)
}
