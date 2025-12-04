// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubnetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePage(v bool) *ListSubnetsRequest
	GetEnablePage() *bool
	SetPageNumber(v int32) *ListSubnetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSubnetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSubnetsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListSubnetsRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListSubnetsRequest
	GetStatus() *string
	SetSubnetId(v string) *ListSubnetsRequest
	GetSubnetId() *string
	SetSubnetName(v string) *ListSubnetsRequest
	GetSubnetName() *string
	SetTag(v []*ListSubnetsRequestTag) *ListSubnetsRequest
	GetTag() []*ListSubnetsRequestTag
	SetType(v string) *ListSubnetsRequest
	GetType() *string
	SetVpdId(v string) *ListSubnetsRequest
	GetVpdId() *string
	SetZoneId(v string) *ListSubnetsRequest
	GetZoneId() *string
}

type ListSubnetsRequest struct {
	// Specifies whether to query by page. Optional values:
	//
	// 	- **true**: Enable pagination query
	//
	// 	- **false**: Pagination query is disabled
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// The number of the page to return. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
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
	// The status of the CLB instance. Valid values:
	//
	// 	- **Available**: Normal
	//
	// 	- **Not Available**: Unavailable
	//
	// 	- **Executing**: Executing
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Lingjun subnet instance ID
	//
	// example:
	//
	// subnet-anhtskts
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Lingjun subnet instance name
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*ListSubnetsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Lingjun Subnet Usage Type; optional; optional. Valid values:
	//
	// 	- **If you do not set this field for a common type**
	//
	// 	- **OOB*	- :OOB type
	//
	// 	- **LB**: LB type
	//
	// example:
	//
	// Null
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the Lingjun CIDR block.
	//
	// example:
	//
	// vpd-fuliephf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID of the disk.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListSubnetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsRequest) GoString() string {
	return s.String()
}

func (s *ListSubnetsRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListSubnetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSubnetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubnetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSubnetsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSubnetsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSubnetsRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *ListSubnetsRequest) GetSubnetName() *string {
	return s.SubnetName
}

func (s *ListSubnetsRequest) GetTag() []*ListSubnetsRequestTag {
	return s.Tag
}

func (s *ListSubnetsRequest) GetType() *string {
	return s.Type
}

func (s *ListSubnetsRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *ListSubnetsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListSubnetsRequest) SetEnablePage(v bool) *ListSubnetsRequest {
	s.EnablePage = &v
	return s
}

func (s *ListSubnetsRequest) SetPageNumber(v int32) *ListSubnetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubnetsRequest) SetPageSize(v int32) *ListSubnetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubnetsRequest) SetRegionId(v string) *ListSubnetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSubnetsRequest) SetResourceGroupId(v string) *ListSubnetsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSubnetsRequest) SetStatus(v string) *ListSubnetsRequest {
	s.Status = &v
	return s
}

func (s *ListSubnetsRequest) SetSubnetId(v string) *ListSubnetsRequest {
	s.SubnetId = &v
	return s
}

func (s *ListSubnetsRequest) SetSubnetName(v string) *ListSubnetsRequest {
	s.SubnetName = &v
	return s
}

func (s *ListSubnetsRequest) SetTag(v []*ListSubnetsRequestTag) *ListSubnetsRequest {
	s.Tag = v
	return s
}

func (s *ListSubnetsRequest) SetType(v string) *ListSubnetsRequest {
	s.Type = &v
	return s
}

func (s *ListSubnetsRequest) SetVpdId(v string) *ListSubnetsRequest {
	s.VpdId = &v
	return s
}

func (s *ListSubnetsRequest) SetZoneId(v string) *ListSubnetsRequest {
	s.ZoneId = &v
	return s
}

func (s *ListSubnetsRequest) Validate() error {
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

type ListSubnetsRequestTag struct {
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// rg-subnet
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// subnet-group-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSubnetsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsRequestTag) GoString() string {
	return s.String()
}

func (s *ListSubnetsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListSubnetsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListSubnetsRequestTag) SetKey(v string) *ListSubnetsRequestTag {
	s.Key = &v
	return s
}

func (s *ListSubnetsRequestTag) SetValue(v string) *ListSubnetsRequestTag {
	s.Value = &v
	return s
}

func (s *ListSubnetsRequestTag) Validate() error {
	return dara.Validate(s)
}
