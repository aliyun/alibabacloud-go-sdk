// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ListVccsRequest
	GetBandwidth() *int32
	SetCenId(v string) *ListVccsRequest
	GetCenId() *string
	SetEnablePage(v bool) *ListVccsRequest
	GetEnablePage() *bool
	SetExStatus(v string) *ListVccsRequest
	GetExStatus() *string
	SetFilterErId(v string) *ListVccsRequest
	GetFilterErId() *string
	SetPageNumber(v int32) *ListVccsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVccsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVccsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVccsRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListVccsRequest
	GetStatus() *string
	SetTag(v []*ListVccsRequestTag) *ListVccsRequest
	GetTag() []*ListVccsRequestTag
	SetVccId(v string) *ListVccsRequest
	GetVccId() *string
	SetVpcId(v string) *ListVccsRequest
	GetVpcId() *string
	SetVpdId(v string) *ListVccsRequest
	GetVpdId() *string
}

type ListVccsRequest struct {
	// The peak bandwidth of the Lingjun connection instance. Unit: Mbit/s. Valid values: 1000 to 400000
	//
	// example:
	//
	// 5000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the CEN instance; [What is the CEN?](https://help.aliyun.com/document_detail/181681.html)
	//
	// You can call the [DescribeCens](https://help.aliyun.com/document_detail/468215.htm) to query the information of CEN instances under the current Alibaba Cloud account.
	//
	// example:
	//
	// cen-95iwtpyvj3kk1v0ao0
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Specifies whether to enable paged query. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// Excludes all data in the specified status. If the status parameter exists, ExStatus does not take effect.
	//
	// example:
	//
	// Prepaid
	ExStatus *string `json:"ExStatus,omitempty" xml:"ExStatus,omitempty"`
	// Filter queries by Lingjun HUB instance ID
	//
	// example:
	//
	// er-a7rqv1rq
	FilterErId *string `json:"FilterErId,omitempty" xml:"FilterErId,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*ListVccsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-bp1nrtkmamy329u6a1z0i
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-omqutbff
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s ListVccsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVccsRequest) GoString() string {
	return s.String()
}

func (s *ListVccsRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListVccsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListVccsRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListVccsRequest) GetExStatus() *string {
	return s.ExStatus
}

func (s *ListVccsRequest) GetFilterErId() *string {
	return s.FilterErId
}

func (s *ListVccsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVccsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVccsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVccsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVccsRequest) GetTag() []*ListVccsRequestTag {
	return s.Tag
}

func (s *ListVccsRequest) GetVccId() *string {
	return s.VccId
}

func (s *ListVccsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVccsRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *ListVccsRequest) SetBandwidth(v int32) *ListVccsRequest {
	s.Bandwidth = &v
	return s
}

func (s *ListVccsRequest) SetCenId(v string) *ListVccsRequest {
	s.CenId = &v
	return s
}

func (s *ListVccsRequest) SetEnablePage(v bool) *ListVccsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListVccsRequest) SetExStatus(v string) *ListVccsRequest {
	s.ExStatus = &v
	return s
}

func (s *ListVccsRequest) SetFilterErId(v string) *ListVccsRequest {
	s.FilterErId = &v
	return s
}

func (s *ListVccsRequest) SetPageNumber(v int32) *ListVccsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVccsRequest) SetPageSize(v int32) *ListVccsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVccsRequest) SetRegionId(v string) *ListVccsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVccsRequest) SetResourceGroupId(v string) *ListVccsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVccsRequest) SetStatus(v string) *ListVccsRequest {
	s.Status = &v
	return s
}

func (s *ListVccsRequest) SetTag(v []*ListVccsRequestTag) *ListVccsRequest {
	s.Tag = v
	return s
}

func (s *ListVccsRequest) SetVccId(v string) *ListVccsRequest {
	s.VccId = &v
	return s
}

func (s *ListVccsRequest) SetVpcId(v string) *ListVccsRequest {
	s.VpcId = &v
	return s
}

func (s *ListVccsRequest) SetVpdId(v string) *ListVccsRequest {
	s.VpdId = &v
	return s
}

func (s *ListVccsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVccsRequestTag struct {
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVccsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVccsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVccsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVccsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVccsRequestTag) SetKey(v string) *ListVccsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVccsRequestTag) SetValue(v string) *ListVccsRequestTag {
	s.Value = &v
	return s
}

func (s *ListVccsRequestTag) Validate() error {
	return dara.Validate(s)
}
