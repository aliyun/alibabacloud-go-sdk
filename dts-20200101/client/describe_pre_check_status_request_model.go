// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribePreCheckStatusRequest
	GetDtsJobId() *string
	SetJobCode(v string) *DescribePreCheckStatusRequest
	GetJobCode() *string
	SetName(v string) *DescribePreCheckStatusRequest
	GetName() *string
	SetPageNo(v string) *DescribePreCheckStatusRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribePreCheckStatusRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribePreCheckStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePreCheckStatusRequest
	GetResourceGroupId() *string
	SetStructPhase(v string) *DescribePreCheckStatusRequest
	GetStructPhase() *string
	SetStructType(v string) *DescribePreCheckStatusRequest
	GetStructType() *string
	SetZeroEtlJob(v bool) *DescribePreCheckStatusRequest
	GetZeroEtlJob() *bool
}

type DescribePreCheckStatusRequest struct {
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The task code. Set this parameter to specify the type of the subtask for the DTS task to query. Valid values:
	//
	// - **01**: Precheck.
	//
	// - **02**: Database and table structure migration or initialization.
	//
	// - **03**: Full data migration or initialization.
	//
	// - **04**: Incremental data migration or synchronization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	JobCode *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	// The filter condition used to perform a fuzzy match on table names.
	//
	// example:
	//
	// dewuprop
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Valid values: an integer greater than **0*	- and less than or equal to the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the DTS instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz4us4iruleja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The filter condition for schema migration. Valid values: tables, views, and functions.
	//
	// example:
	//
	// View
	StructPhase *string `json:"StructPhase,omitempty" xml:"StructPhase,omitempty"`
	// The type of the schema definition. Valid values:
	//
	// - **before**: Schema migration or initial schema synchronization.
	//
	// - **after**: DDL operations in incremental data migration or synchronization.
	//
	// example:
	//
	// before
	StructType *string `json:"StructType,omitempty" xml:"StructType,omitempty"`
	// Specifies whether the task is a seamless integration (Zero-ETL) task. Valid values:
	//
	// - **false**: No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// false
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DescribePreCheckStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribePreCheckStatusRequest) GetJobCode() *string {
	return s.JobCode
}

func (s *DescribePreCheckStatusRequest) GetName() *string {
	return s.Name
}

func (s *DescribePreCheckStatusRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribePreCheckStatusRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePreCheckStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePreCheckStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePreCheckStatusRequest) GetStructPhase() *string {
	return s.StructPhase
}

func (s *DescribePreCheckStatusRequest) GetStructType() *string {
	return s.StructType
}

func (s *DescribePreCheckStatusRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DescribePreCheckStatusRequest) SetDtsJobId(v string) *DescribePreCheckStatusRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetJobCode(v string) *DescribePreCheckStatusRequest {
	s.JobCode = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetName(v string) *DescribePreCheckStatusRequest {
	s.Name = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetPageNo(v string) *DescribePreCheckStatusRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetPageSize(v string) *DescribePreCheckStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetRegionId(v string) *DescribePreCheckStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetResourceGroupId(v string) *DescribePreCheckStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetStructPhase(v string) *DescribePreCheckStatusRequest {
	s.StructPhase = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetStructType(v string) *DescribePreCheckStatusRequest {
	s.StructType = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetZeroEtlJob(v bool) *DescribePreCheckStatusRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DescribePreCheckStatusRequest) Validate() error {
	return dara.Validate(s)
}
