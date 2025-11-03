// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToEcrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGrantRulesToEcrRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeGrantRulesToEcrRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeGrantRulesToEcrRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeGrantRulesToEcrRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGrantRulesToEcrRequest
	GetResourceGroupId() *string
	SetTags(v []*DescribeGrantRulesToEcrRequestTags) *DescribeGrantRulesToEcrRequest
	GetTags() []*DescribeGrantRulesToEcrRequestTags
}

type DescribeGrantRulesToEcrRequest struct {
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	Tags []*DescribeGrantRulesToEcrRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeGrantRulesToEcrRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToEcrRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToEcrRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGrantRulesToEcrRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeGrantRulesToEcrRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeGrantRulesToEcrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantRulesToEcrRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGrantRulesToEcrRequest) GetTags() []*DescribeGrantRulesToEcrRequestTags {
	return s.Tags
}

func (s *DescribeGrantRulesToEcrRequest) SetInstanceId(v string) *DescribeGrantRulesToEcrRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequest) SetPageNumber(v int64) *DescribeGrantRulesToEcrRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequest) SetPageSize(v int64) *DescribeGrantRulesToEcrRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequest) SetRegionId(v string) *DescribeGrantRulesToEcrRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequest) SetResourceGroupId(v string) *DescribeGrantRulesToEcrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequest) SetTags(v []*DescribeGrantRulesToEcrRequestTags) *DescribeGrantRulesToEcrRequest {
	s.Tags = v
	return s
}

func (s *DescribeGrantRulesToEcrRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGrantRulesToEcrRequestTags struct {
	// The tag keys. You must specify at least one tag key and at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeGrantRulesToEcrRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToEcrRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToEcrRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeGrantRulesToEcrRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeGrantRulesToEcrRequestTags) SetKey(v string) *DescribeGrantRulesToEcrRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequestTags) SetValue(v string) *DescribeGrantRulesToEcrRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeGrantRulesToEcrRequestTags) Validate() error {
	return dara.Validate(s)
}
