// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBandwidthPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *ListBandwidthPackagesRequest
	GetBandwidthPackageId() *string
	SetPageNumber(v int32) *ListBandwidthPackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBandwidthPackagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListBandwidthPackagesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListBandwidthPackagesRequest
	GetResourceGroupId() *string
	SetState(v string) *ListBandwidthPackagesRequest
	GetState() *string
	SetTag(v []*ListBandwidthPackagesRequestTag) *ListBandwidthPackagesRequest
	GetTag() []*ListBandwidthPackagesRequestTag
	SetType(v string) *ListBandwidthPackagesRequest
	GetType() *string
}

type ListBandwidthPackagesRequest struct {
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
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
	// rg-aek2ry6mp2c****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the bandwidth plan. Valid values:
	//
	// 	- **init**: The bandwidth plan is being initialized.
	//
	// 	- **active**: The bandwidth plan is available.
	//
	// 	- **binded**: The bandwidth plan is associated.
	//
	// 	- **binding**: The bandwidth plan is being associated.
	//
	// 	- **unbinding**: The bandwidth plan is being disassociated.
	//
	// 	- **updating**: The bandwidth plan is being updated.
	//
	// 	- **finacialLocked**: The bandwidth plan is locked due to overdue payments.
	//
	// 	- **locked**: The bandwidth plan is locked.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag of the bandwidth plan.
	Tag []*ListBandwidthPackagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the bandwidth plan. Valid values:
	//
	// 	- **Basic**: a basic bandwidth plan
	//
	// 	- **CrossDomain**: a cross-border acceleration bandwidth plan
	//
	// If you call this operation on the China site (aliyun.com), you can set Type only to **Basic**.
	//
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBandwidthPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ListBandwidthPackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBandwidthPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBandwidthPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBandwidthPackagesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListBandwidthPackagesRequest) GetState() *string {
	return s.State
}

func (s *ListBandwidthPackagesRequest) GetTag() []*ListBandwidthPackagesRequestTag {
	return s.Tag
}

func (s *ListBandwidthPackagesRequest) GetType() *string {
	return s.Type
}

func (s *ListBandwidthPackagesRequest) SetBandwidthPackageId(v string) *ListBandwidthPackagesRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetPageNumber(v int32) *ListBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetPageSize(v int32) *ListBandwidthPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetRegionId(v string) *ListBandwidthPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetResourceGroupId(v string) *ListBandwidthPackagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetState(v string) *ListBandwidthPackagesRequest {
	s.State = &v
	return s
}

func (s *ListBandwidthPackagesRequest) SetTag(v []*ListBandwidthPackagesRequestTag) *ListBandwidthPackagesRequest {
	s.Tag = v
	return s
}

func (s *ListBandwidthPackagesRequest) SetType(v string) *ListBandwidthPackagesRequest {
	s.Type = &v
	return s
}

func (s *ListBandwidthPackagesRequest) Validate() error {
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

type ListBandwidthPackagesRequestTag struct {
	// The tag key of the bandwidth plan. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the bandwidth plan. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBandwidthPackagesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthPackagesRequestTag) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListBandwidthPackagesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListBandwidthPackagesRequestTag) SetKey(v string) *ListBandwidthPackagesRequestTag {
	s.Key = &v
	return s
}

func (s *ListBandwidthPackagesRequestTag) SetValue(v string) *ListBandwidthPackagesRequestTag {
	s.Value = &v
	return s
}

func (s *ListBandwidthPackagesRequestTag) Validate() error {
	return dara.Validate(s)
}
