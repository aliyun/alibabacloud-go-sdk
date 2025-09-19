// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAntiBruteForceRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstanceAntiBruteForceRulesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeInstanceAntiBruteForceRulesRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeInstanceAntiBruteForceRulesRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeInstanceAntiBruteForceRulesRequest
	GetSourceIp() *string
	SetUuidList(v []*string) *DescribeInstanceAntiBruteForceRulesRequest
	GetUuidList() []*string
}

type DescribeInstanceAntiBruteForceRulesRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: **10000**.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 115.238.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUIDs of the servers.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAntiBruteForceRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetCurrentPage(v int32) *DescribeInstanceAntiBruteForceRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetPageSize(v int32) *DescribeInstanceAntiBruteForceRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetResourceOwnerId(v int64) *DescribeInstanceAntiBruteForceRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetSourceIp(v string) *DescribeInstanceAntiBruteForceRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetUuidList(v []*string) *DescribeInstanceAntiBruteForceRulesRequest {
	s.UuidList = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) Validate() error {
	return dara.Validate(s)
}
