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
	// The end of the time range to query. You can call the [DescribePreCheckStatus](https://help.aliyun.com/document_detail/209718.html) operation to query the execution time of the subtasks.
	//
	// > 	- To obtain the logs that are generated for DTS subtasks within a specific period of time, you can call the [DescribePreCheckStatus](https://help.aliyun.com/document_detail/209718.html) operation to query the execution time of the subtasks.
	//
	// >	- Specify the time in the 13-digit UNIX timestamp format. Unit: milliseconds. You can use a search engine to obtain a UNIX timestamp converter.
	//
	// example:
	//
	// 1620897227000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword that is passed to specify the query content.
	//
	// >  Fuzzy match is used and the keyword is case-sensitive.
	//
	// example:
	//
	// state = IDLE
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0 and less than or equal to the maximum value supported by the integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries to return on each page. Valid values: **20**, **50**, **100**, **500**, and **1000**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query.
	//
	// > 	- To obtain the logs that are generated for Data Transmission Service (DTS) subtasks within a specific period of time, you can call the [DescribePreCheckStatus](https://help.aliyun.com/document_detail/209718.html) operation to query the execution time of the subtasks.
	//
	// >	- Specify the time in the 13-digit UNIX timestamp format. Unit: milliseconds. You can use a search engine to obtain a UNIX timestamp converter.
	//
	// example:
	//
	// 1620896327000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The log level. Separate multiple log levels with commas (,). Valid values:
	//
	// 	- **NORMAL**: displays the logs that are generated when the DTS task runs as expected.
	//
	// 	- **WARN**: displays the logs about severe issues that stop the DTS task from running.
	//
	// 	- **ERROR**: displays the logs about unexpected issues that stop specific processes form running.
	//
	// example:
	//
	// NORMAL,WARN,ERROR
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of a DTS subtask. Valid values:
	//
	// 	- **DATA_LOAD**: full migration or full synchronization
	//
	// 	- **ONLINE_WRITER**: incremental migration
	//
	// 	- **SYNC_WRITER**: incremental synchronization
	//
	// example:
	//
	// SYNC_WRITER
	SubJobType *string `json:"SubJobType,omitempty" xml:"SubJobType,omitempty"`
	// Whether it is a seamless integration (Zero-ETL) task, the value can be: - **true**: Yes. - **false**: No.
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
