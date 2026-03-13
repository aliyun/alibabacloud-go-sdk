// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseRuleChangeLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeBaseRuleChangeLogRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeBaseRuleChangeLogRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeBaseRuleChangeLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBaseRuleChangeLogRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBaseRuleChangeLogRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeBaseRuleChangeLogRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeBaseRuleChangeLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeBaseRuleChangeLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseRuleChangeLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeBaseRuleChangeLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBaseRuleChangeLogRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBaseRuleChangeLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBaseRuleChangeLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBaseRuleChangeLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBaseRuleChangeLogRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeBaseRuleChangeLogRequest) SetInstanceId(v string) *DescribeBaseRuleChangeLogRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBaseRuleChangeLogRequest) SetLang(v string) *DescribeBaseRuleChangeLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBaseRuleChangeLogRequest) SetPageNumber(v int32) *DescribeBaseRuleChangeLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBaseRuleChangeLogRequest) SetPageSize(v int32) *DescribeBaseRuleChangeLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBaseRuleChangeLogRequest) SetRegionId(v string) *DescribeBaseRuleChangeLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBaseRuleChangeLogRequest) SetResourceManagerResourceGroupId(v string) *DescribeBaseRuleChangeLogRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeBaseRuleChangeLogRequest) Validate() error {
	return dara.Validate(s)
}
