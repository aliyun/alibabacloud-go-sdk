// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorNamespaceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeHybridMonitorNamespaceListRequest
	GetKeyword() *string
	SetNamespace(v string) *DescribeHybridMonitorNamespaceListRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeHybridMonitorNamespaceListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridMonitorNamespaceListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHybridMonitorNamespaceListRequest
	GetRegionId() *string
	SetShowTaskStatistic(v bool) *DescribeHybridMonitorNamespaceListRequest
	GetShowTaskStatistic() *bool
}

type DescribeHybridMonitorNamespaceListRequest struct {
	// The search keyword.
	//
	// example:
	//
	// aliyun
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The name of the namespace.
	//
	// The name can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// aliyun-test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// Page numbers start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Page numbers start from 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return the configuration details of metric import tasks for Alibaba Cloud services and the number of metric import tasks for third-party services. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	ShowTaskStatistic *bool `json:"ShowTaskStatistic,omitempty" xml:"ShowTaskStatistic,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeHybridMonitorNamespaceListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeHybridMonitorNamespaceListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridMonitorNamespaceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridMonitorNamespaceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridMonitorNamespaceListRequest) GetShowTaskStatistic() *bool {
	return s.ShowTaskStatistic
}

func (s *DescribeHybridMonitorNamespaceListRequest) SetKeyword(v string) *DescribeHybridMonitorNamespaceListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListRequest) SetNamespace(v string) *DescribeHybridMonitorNamespaceListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListRequest) SetPageNumber(v int32) *DescribeHybridMonitorNamespaceListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListRequest) SetPageSize(v int32) *DescribeHybridMonitorNamespaceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListRequest) SetRegionId(v string) *DescribeHybridMonitorNamespaceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListRequest) SetShowTaskStatistic(v bool) *DescribeHybridMonitorNamespaceListRequest {
	s.ShowTaskStatistic = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListRequest) Validate() error {
	return dara.Validate(s)
}
