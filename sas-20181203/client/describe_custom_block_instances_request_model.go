// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBlockInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockIp(v string) *DescribeCustomBlockInstancesRequest
	GetBlockIp() *string
	SetBound(v string) *DescribeCustomBlockInstancesRequest
	GetBound() *string
	SetCurrentPage(v int32) *DescribeCustomBlockInstancesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeCustomBlockInstancesRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeCustomBlockInstancesRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *DescribeCustomBlockInstancesRequest
	GetStatus() *int32
}

type DescribeCustomBlockInstancesRequest struct {
	// The IP address that you want to specify in the rule.
	//
	// >  You can call the [DescribeCustomBlockRecords](~~DescribeCustomBlockRecords~~) operation to obtain the IP address.
	//
	// example:
	//
	// 47.92.33.1xx
	BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
	// The traffic direction that you want to specify in the custom rule. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// example:
	//
	// in
	Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 8
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether the rule is enabled for the server.
	//
	// 	- **2**: enabling failed
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCustomBlockInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockInstancesRequest) GetBlockIp() *string {
	return s.BlockIp
}

func (s *DescribeCustomBlockInstancesRequest) GetBound() *string {
	return s.Bound
}

func (s *DescribeCustomBlockInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCustomBlockInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomBlockInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCustomBlockInstancesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCustomBlockInstancesRequest) SetBlockIp(v string) *DescribeCustomBlockInstancesRequest {
	s.BlockIp = &v
	return s
}

func (s *DescribeCustomBlockInstancesRequest) SetBound(v string) *DescribeCustomBlockInstancesRequest {
	s.Bound = &v
	return s
}

func (s *DescribeCustomBlockInstancesRequest) SetCurrentPage(v int32) *DescribeCustomBlockInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCustomBlockInstancesRequest) SetPageSize(v int32) *DescribeCustomBlockInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomBlockInstancesRequest) SetResourceOwnerId(v int64) *DescribeCustomBlockInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCustomBlockInstancesRequest) SetStatus(v int32) *DescribeCustomBlockInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCustomBlockInstancesRequest) Validate() error {
	return dara.Validate(s)
}
