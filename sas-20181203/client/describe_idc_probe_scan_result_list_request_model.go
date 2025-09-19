// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcProbeScanResultListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeIdcProbeScanResultListRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeIdcProbeScanResultListRequest
	GetCurrentPage() *int32
	SetFoundEndTime(v int64) *DescribeIdcProbeScanResultListRequest
	GetFoundEndTime() *int64
	SetFoundStartTime(v int64) *DescribeIdcProbeScanResultListRequest
	GetFoundStartTime() *int64
	SetLogicalExp(v string) *DescribeIdcProbeScanResultListRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *DescribeIdcProbeScanResultListRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeIdcProbeScanResultListRequest
	GetStatus() *string
}

type DescribeIdcProbeScanResultListRequest struct {
	// The search conditions for assets. This parameter is in the JSON format. The value is case-sensitive.
	//
	// >  A search condition can be the instance ID, instance name, VPC ID, region, or public IP address. You can call the [DescribeIdcAssetCriteria](https://help.aliyun.com/document_detail/2842671.html) operation to query supported search conditions.
	//
	// example:
	//
	// [{\\"name\\":\\"scannedIp\\",\\"value\\":\\"192.168.2.11\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the scan.
	//
	// example:
	//
	// 1720006819000
	FoundEndTime *int64 `json:"FoundEndTime,omitempty" xml:"FoundEndTime,omitempty"`
	// The start time of the scan.
	//
	// example:
	//
	// 1720006818000
	FoundStartTime *int64 `json:"FoundStartTime,omitempty" xml:"FoundStartTime,omitempty"`
	// The logical operator that combines multiple search conditions. Valid values:
	//
	// 	- **OR******
	//
	// 	- **AND******
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The statuses of the corresponding probes. Separate multiple values with commas (,). Valid values:
	//
	// 	- **0**: The probe is valid.
	//
	// 	- **1**: The probe is ignored.
	//
	// 	- **2**: The probe is invalid.
	//
	// 	- **3**: The probe expired.
	//
	// 	- **4**: The probe does not exist.
	//
	// example:
	//
	// 0,1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIdcProbeScanResultListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeScanResultListRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeScanResultListRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeIdcProbeScanResultListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeIdcProbeScanResultListRequest) GetFoundEndTime() *int64 {
	return s.FoundEndTime
}

func (s *DescribeIdcProbeScanResultListRequest) GetFoundStartTime() *int64 {
	return s.FoundStartTime
}

func (s *DescribeIdcProbeScanResultListRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeIdcProbeScanResultListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIdcProbeScanResultListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeIdcProbeScanResultListRequest) SetCriteria(v string) *DescribeIdcProbeScanResultListRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) SetCurrentPage(v int32) *DescribeIdcProbeScanResultListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) SetFoundEndTime(v int64) *DescribeIdcProbeScanResultListRequest {
	s.FoundEndTime = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) SetFoundStartTime(v int64) *DescribeIdcProbeScanResultListRequest {
	s.FoundStartTime = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) SetLogicalExp(v string) *DescribeIdcProbeScanResultListRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) SetPageSize(v int32) *DescribeIdcProbeScanResultListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) SetStatus(v string) *DescribeIdcProbeScanResultListRequest {
	s.Status = &v
	return s
}

func (s *DescribeIdcProbeScanResultListRequest) Validate() error {
	return dara.Validate(s)
}
