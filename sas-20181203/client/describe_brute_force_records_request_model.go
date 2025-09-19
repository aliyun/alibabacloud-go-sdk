// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBruteForceRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *DescribeBruteForceRecordsRequest
	GetBlockIp() *string
	SetCurrentPage(v int32) *DescribeBruteForceRecordsRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *DescribeBruteForceRecordsRequest
	GetInstanceId() *string
	SetPageSize(v int32) *DescribeBruteForceRecordsRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeBruteForceRecordsRequest
	GetRemark() *string
	SetResourceOwnerId(v int64) *DescribeBruteForceRecordsRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *DescribeBruteForceRecordsRequest
	GetStatus() *int32
}

type DescribeBruteForceRecordsRequest struct {
	// The IP address that is blocked.
	//
	// example:
	//
	// 175.106.XX.XX
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page. We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server to query.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the defense rule. Valid values:
	//
	// 	- **0**: invalid
	//
	// 	- **1**: enabled
	//
	// 	- **2**: failed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBruteForceRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceRecordsRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DescribeBruteForceRecordsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBruteForceRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBruteForceRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBruteForceRecordsRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeBruteForceRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBruteForceRecordsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeBruteForceRecordsRequest) SetBlockIp(v string) *DescribeBruteForceRecordsRequest {
	s.BlockIp = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) SetCurrentPage(v int32) *DescribeBruteForceRecordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) SetInstanceId(v string) *DescribeBruteForceRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) SetPageSize(v int32) *DescribeBruteForceRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) SetRemark(v string) *DescribeBruteForceRecordsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) SetResourceOwnerId(v int64) *DescribeBruteForceRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) SetStatus(v int32) *DescribeBruteForceRecordsRequest {
	s.Status = &v
	return s
}

func (s *DescribeBruteForceRecordsRequest) Validate() error {
	return dara.Validate(s)
}
