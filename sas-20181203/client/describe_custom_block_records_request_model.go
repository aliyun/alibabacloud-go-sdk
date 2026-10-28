// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBlockRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *DescribeCustomBlockRecordsRequest
	GetBlockIp() *string
	SetCurrentPage(v int32) *DescribeCustomBlockRecordsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeCustomBlockRecordsRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeCustomBlockRecordsRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *DescribeCustomBlockRecordsRequest
	GetStatus() *int32
}

type DescribeCustomBlockRecordsRequest struct {
	// The IP address to be blocked to prevent brute-force attacks.
	//
	// example:
	//
	// 117.66.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// Specifies the page number of the returned results. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Settings for paging. Specifies the number of records to return on each page during a paged query. Default value: **20**, which indicates that 20 records are displayed on each page.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the defense rule against brute-force attacks. Valid values:
	//
	// - **0**: invalid
	//
	// - **1**: enabled
	//
	// - **2**: failed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCustomBlockRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockRecordsRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DescribeCustomBlockRecordsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCustomBlockRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomBlockRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCustomBlockRecordsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCustomBlockRecordsRequest) SetBlockIp(v string) *DescribeCustomBlockRecordsRequest {
	s.BlockIp = &v
	return s
}

func (s *DescribeCustomBlockRecordsRequest) SetCurrentPage(v int32) *DescribeCustomBlockRecordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCustomBlockRecordsRequest) SetPageSize(v int32) *DescribeCustomBlockRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomBlockRecordsRequest) SetResourceOwnerId(v int64) *DescribeCustomBlockRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCustomBlockRecordsRequest) SetStatus(v int32) *DescribeCustomBlockRecordsRequest {
	s.Status = &v
	return s
}

func (s *DescribeCustomBlockRecordsRequest) Validate() error {
	return dara.Validate(s)
}
