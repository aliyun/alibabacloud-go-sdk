// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackend(v string) *DescribeDomainsRequest
	GetBackend() *string
	SetDomain(v string) *DescribeDomainsRequest
	GetDomain() *string
	SetInstanceId(v string) *DescribeDomainsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeDomainsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDomainsRequest
	GetResourceManagerResourceGroupId() *string
	SetTag(v []*DescribeDomainsRequestTag) *DescribeDomainsRequest
	GetTag() []*DescribeDomainsRequestTag
}

type DescribeDomainsRequest struct {
	// An array of HTTPS listener ports.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tag of the resource. You can specify up to 20 tags.
	Tag []*DescribeDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) GetBackend() *string {
	return s.Backend
}

func (s *DescribeDomainsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDomainsRequest) GetTag() []*DescribeDomainsRequestTag {
	return s.Tag
}

func (s *DescribeDomainsRequest) SetBackend(v string) *DescribeDomainsRequest {
	s.Backend = &v
	return s
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceId(v string) *DescribeDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetRegionId(v string) *DescribeDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetTag(v []*DescribeDomainsRequestTag) *DescribeDomainsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDomainsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// TagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDomainsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainsRequestTag) SetKey(v string) *DescribeDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDomainsRequestTag) SetValue(v string) *DescribeDomainsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDomainsRequestTag) Validate() error {
	return dara.Validate(s)
}
