// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsServiceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeDtsServiceLogRequest
	GetDtsJobId() *string
	SetEndTime(v int64) *DescribeDtsServiceLogRequest
	GetEndTime() *int64
	SetKeyword(v string) *DescribeDtsServiceLogRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeDtsServiceLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDtsServiceLogRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDtsServiceLogRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDtsServiceLogRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDtsServiceLogRequest
	GetStartTime() *int64
	SetStatus(v string) *DescribeDtsServiceLogRequest
	GetStatus() *string
	SetSubJobType(v string) *DescribeDtsServiceLogRequest
	GetSubJobType() *string
	SetZeroEtlJob(v bool) *DescribeDtsServiceLogRequest
	GetZeroEtlJob() *bool
}

type DescribeDtsServiceLogRequest struct {
	// The ID of the data migration or synchronization task.
	//
	// example:
	//
	// c1yr56py103****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The end time of the log information. You can call [DescribePreCheckStatus](https://help.aliyun.com/document_detail/209718.html) to query the end time.
	//
	// > - To query the log information of a DTS subtask within a specific time range, call [DescribePreCheckStatus](https://help.aliyun.com/document_detail/209718.html) to query the execution time of the DTS subtask.
	//
	// - The time is a 13-digit UNIX timestamp in milliseconds. You can use a search engine to find a UNIX timestamp converter.
	//
	// example:
	//
	// 1620897227000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword used to filter query results.
	//
	// > Fuzzy match is used and the keyword is case-sensitive.
	//
	// example:
	//
	// state = IDLE
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries per page. Valid values: **20**, **50**, **100**, **500**, and **1000**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the log information.
	//
	// > - To query the log information of a DTS subtask within a specific time range, call [DescribePreCheckStatus](https://help.aliyun.com/document_detail/209718.html) to query the execution time of the DTS subtask.
	//
	// - The start time is a 13-digit UNIX timestamp in milliseconds. You can use a search engine to find a UNIX timestamp converter.
	//
	// example:
	//
	// 1620896327000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The log level of the log information. Separate multiple values with commas (,). Valid values:
	//
	// - **NORMAL**: Normal.
	//
	// - **WARN**: Warning.
	//
	// - **ERROR**: Error.
	//
	// example:
	//
	// NORMAL,WARN,ERROR
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the DTS task subnode. Valid values:
	//
	//
	// - **DATA_LOAD**: full data migration or initial full data synchronization.
	//
	// - **ONLINE_WRITER**: incremental data migration.
	//
	// - **SYNC_WRITER**: incremental data synchronization.
	//
	// example:
	//
	// SYNC_WRITER
	SubJobType *string `json:"SubJobType,omitempty" xml:"SubJobType,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DescribeDtsServiceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsServiceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsServiceLogRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsServiceLogRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDtsServiceLogRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDtsServiceLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDtsServiceLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDtsServiceLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsServiceLogRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsServiceLogRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDtsServiceLogRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsServiceLogRequest) GetSubJobType() *string {
	return s.SubJobType
}

func (s *DescribeDtsServiceLogRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DescribeDtsServiceLogRequest) SetDtsJobId(v string) *DescribeDtsServiceLogRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetEndTime(v int64) *DescribeDtsServiceLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetKeyword(v string) *DescribeDtsServiceLogRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetPageNumber(v int32) *DescribeDtsServiceLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetPageSize(v int32) *DescribeDtsServiceLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetRegionId(v string) *DescribeDtsServiceLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetResourceGroupId(v string) *DescribeDtsServiceLogRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetStartTime(v int64) *DescribeDtsServiceLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetStatus(v string) *DescribeDtsServiceLogRequest {
	s.Status = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetSubJobType(v string) *DescribeDtsServiceLogRequest {
	s.SubJobType = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) SetZeroEtlJob(v bool) *DescribeDtsServiceLogRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DescribeDtsServiceLogRequest) Validate() error {
	return dara.Validate(s)
}
