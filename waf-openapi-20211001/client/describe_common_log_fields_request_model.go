// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonLogFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCommonLogFieldsRequest
	GetInstanceId() *string
	SetIsDefault(v bool) *DescribeCommonLogFieldsRequest
	GetIsDefault() *bool
	SetIsRequired(v bool) *DescribeCommonLogFieldsRequest
	GetIsRequired() *bool
	SetLogKeyList(v []*string) *DescribeCommonLogFieldsRequest
	GetLogKeyList() []*string
	SetRegionId(v string) *DescribeCommonLogFieldsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCommonLogFieldsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCommonLogFieldsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-l*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// false
	IsRequired *bool     `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	LogKeyList []*string `json:"LogKeyList,omitempty" xml:"LogKeyList,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2sxgs*****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCommonLogFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonLogFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonLogFieldsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCommonLogFieldsRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeCommonLogFieldsRequest) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *DescribeCommonLogFieldsRequest) GetLogKeyList() []*string {
	return s.LogKeyList
}

func (s *DescribeCommonLogFieldsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommonLogFieldsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCommonLogFieldsRequest) SetInstanceId(v string) *DescribeCommonLogFieldsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetIsDefault(v bool) *DescribeCommonLogFieldsRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetIsRequired(v bool) *DescribeCommonLogFieldsRequest {
	s.IsRequired = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetLogKeyList(v []*string) *DescribeCommonLogFieldsRequest {
	s.LogKeyList = v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetRegionId(v string) *DescribeCommonLogFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetResourceManagerResourceGroupId(v string) *DescribeCommonLogFieldsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) Validate() error {
	return dara.Validate(s)
}
