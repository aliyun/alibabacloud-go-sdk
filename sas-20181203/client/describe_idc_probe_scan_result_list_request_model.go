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
	// The search conditions for assets. This parameter is in JSON format. The parameter names are case-sensitive.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, or public IP address.
	//
	// example:
	//
	// [{\\"name\\":\\"scannedIp\\",\\"value\\":\\"192.168.2.11\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the scan discovery. Specify a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1720006819000
	FoundEndTime *int64 `json:"FoundEndTime,omitempty" xml:"FoundEndTime,omitempty"`
	// The start time of the scan discovery. Specify a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1720006818000
	FoundStartTime *int64 `json:"FoundStartTime,omitempty" xml:"FoundStartTime,omitempty"`
	// The logical relationship among multiple search conditions. Valid values:
	//
	// - **OR**: The search conditions are in a logical **OR*	- relationship.
	//
	// - **AND**: The search conditions are in a logical **AND*	- relationship.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The maximum number of entries per page in a paging query. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// > Specify a value for PageSize.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status list of the corresponding probes. Separate multiple values with commas. Valid values:
	//
	// - **0**: active
	//
	// - **1**: ignored
	//
	// - **2**: invalid
	//
	// - **3**: expired
	//
	// - **4**: probe does not exist
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
