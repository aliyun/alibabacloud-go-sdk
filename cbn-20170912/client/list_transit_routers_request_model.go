// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRoutersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRoutersRequest
	GetCenId() *string
	SetFeatureFilter(v []*ListTransitRoutersRequestFeatureFilter) *ListTransitRoutersRequest
	GetFeatureFilter() []*ListTransitRoutersRequestFeatureFilter
	SetOwnerAccount(v string) *ListTransitRoutersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRoutersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListTransitRoutersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransitRoutersRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTransitRoutersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRoutersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRoutersRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListTransitRoutersRequest
	GetStatus() *string
	SetTag(v []*ListTransitRoutersRequestTag) *ListTransitRoutersRequest
	GetTag() []*ListTransitRoutersRequestTag
	SetTransitRouterId(v string) *ListTransitRoutersRequest
	GetTransitRouterId() *string
	SetTransitRouterName(v string) *ListTransitRoutersRequest
	GetTransitRouterName() *string
	SetType(v string) *ListTransitRoutersRequest
	GetType() *string
}

type ListTransitRoutersRequest struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The field that is used to enable or disable a feature of the transit router.
	FeatureFilter []*ListTransitRoutersRequestFeatureFilter `json:"FeatureFilter,omitempty" xml:"FeatureFilter,omitempty" type:"Repeated"`
	OwnerAccount  *string                                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the transit router is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the transit router. Valid values:
	//
	// 	- **Creating**: The transit router is being created.
	//
	// 	- **Active**: The transit router is available.
	//
	// 	- **Modifying**: The transit router is being modified
	//
	// 	- **Deleting**: The transit router is being deleted.
	//
	// 	- **Upgrading**: The transit router is being upgraded.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListTransitRoutersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The name of the Enterprise Edition transit router.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testname
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
	// The edition of the transit router. Valid values:
	//
	// 	- **Enterprise**: Enhance Edition
	//
	// 	- **Basic**: Basic Edition
	//
	// example:
	//
	// Enterprise
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTransitRoutersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRoutersRequest) GetFeatureFilter() []*ListTransitRoutersRequestFeatureFilter {
	return s.FeatureFilter
}

func (s *ListTransitRoutersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRoutersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRoutersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransitRoutersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransitRoutersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRoutersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRoutersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRoutersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRoutersRequest) GetTag() []*ListTransitRoutersRequestTag {
	return s.Tag
}

func (s *ListTransitRoutersRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRoutersRequest) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *ListTransitRoutersRequest) GetType() *string {
	return s.Type
}

func (s *ListTransitRoutersRequest) SetCenId(v string) *ListTransitRoutersRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetFeatureFilter(v []*ListTransitRoutersRequestFeatureFilter) *ListTransitRoutersRequest {
	s.FeatureFilter = v
	return s
}

func (s *ListTransitRoutersRequest) SetOwnerAccount(v string) *ListTransitRoutersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRoutersRequest) SetOwnerId(v int64) *ListTransitRoutersRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetPageNumber(v int32) *ListTransitRoutersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTransitRoutersRequest) SetPageSize(v int32) *ListTransitRoutersRequest {
	s.PageSize = &v
	return s
}

func (s *ListTransitRoutersRequest) SetRegionId(v string) *ListTransitRoutersRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetResourceOwnerAccount(v string) *ListTransitRoutersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRoutersRequest) SetResourceOwnerId(v int64) *ListTransitRoutersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetStatus(v string) *ListTransitRoutersRequest {
	s.Status = &v
	return s
}

func (s *ListTransitRoutersRequest) SetTag(v []*ListTransitRoutersRequestTag) *ListTransitRoutersRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRoutersRequest) SetTransitRouterId(v string) *ListTransitRoutersRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRoutersRequest) SetTransitRouterName(v string) *ListTransitRoutersRequest {
	s.TransitRouterName = &v
	return s
}

func (s *ListTransitRoutersRequest) SetType(v string) *ListTransitRoutersRequest {
	s.Type = &v
	return s
}

func (s *ListTransitRoutersRequest) Validate() error {
	if s.FeatureFilter != nil {
		for _, item := range s.FeatureFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListTransitRoutersRequestFeatureFilter struct {
	// The value of the field that is used to enable or disable a feature of the transit router. Supported fields:
	//
	// 	- **Multicast**: the multicast feature.
	//
	// example:
	//
	// Multicast
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The fields that are used to enable or disable the features of the transit router. The **Multicast*	- field supports only one value. Valid values:
	//
	// 	- **Enabled**: enables multicast.
	//
	// 	- **Disabled**: disables multicast.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTransitRoutersRequestFeatureFilter) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersRequestFeatureFilter) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersRequestFeatureFilter) GetKey() *string {
	return s.Key
}

func (s *ListTransitRoutersRequestFeatureFilter) GetValue() []*string {
	return s.Value
}

func (s *ListTransitRoutersRequestFeatureFilter) SetKey(v string) *ListTransitRoutersRequestFeatureFilter {
	s.Key = &v
	return s
}

func (s *ListTransitRoutersRequestFeatureFilter) SetValue(v []*string) *ListTransitRoutersRequestFeatureFilter {
	s.Value = v
	return s
}

func (s *ListTransitRoutersRequestFeatureFilter) Validate() error {
	return dara.Validate(s)
}

type ListTransitRoutersRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRoutersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRoutersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRoutersRequestTag) SetKey(v string) *ListTransitRoutersRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRoutersRequestTag) SetValue(v string) *ListTransitRoutersRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRoutersRequestTag) Validate() error {
	return dara.Validate(s)
}
