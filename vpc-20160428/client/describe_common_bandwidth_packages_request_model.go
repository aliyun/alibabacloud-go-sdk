// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonBandwidthPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DescribeCommonBandwidthPackagesRequest
	GetBandwidthPackageId() *string
	SetDryRun(v bool) *DescribeCommonBandwidthPackagesRequest
	GetDryRun() *bool
	SetIncludeReservationData(v bool) *DescribeCommonBandwidthPackagesRequest
	GetIncludeReservationData() *bool
	SetName(v string) *DescribeCommonBandwidthPackagesRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeCommonBandwidthPackagesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCommonBandwidthPackagesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCommonBandwidthPackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCommonBandwidthPackagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCommonBandwidthPackagesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCommonBandwidthPackagesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCommonBandwidthPackagesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCommonBandwidthPackagesRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionEnabled(v bool) *DescribeCommonBandwidthPackagesRequest
	GetSecurityProtectionEnabled() *bool
	SetTag(v []*DescribeCommonBandwidthPackagesRequestTag) *DescribeCommonBandwidthPackagesRequest
	GetTag() []*DescribeCommonBandwidthPackagesRequestTag
}

type DescribeCommonBandwidthPackagesRequest struct {
	// The ID of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqk145****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and instance status. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to return the information about pending orders. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IncludeReservationData *bool `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	// The name of the Internet Shared Bandwidth instance.
	//
	// example:
	//
	// test123
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Internet Shared Bandwidth instance resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
	// rg-acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable Anti-DDoS Pro/Premium. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// The tags to add to the Internet Shared Bandwidth instance.
	Tag []*DescribeCommonBandwidthPackagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCommonBandwidthPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeCommonBandwidthPackagesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeCommonBandwidthPackagesRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeCommonBandwidthPackagesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCommonBandwidthPackagesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCommonBandwidthPackagesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCommonBandwidthPackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCommonBandwidthPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCommonBandwidthPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommonBandwidthPackagesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCommonBandwidthPackagesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCommonBandwidthPackagesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCommonBandwidthPackagesRequest) GetSecurityProtectionEnabled() *bool {
	return s.SecurityProtectionEnabled
}

func (s *DescribeCommonBandwidthPackagesRequest) GetTag() []*DescribeCommonBandwidthPackagesRequestTag {
	return s.Tag
}

func (s *DescribeCommonBandwidthPackagesRequest) SetBandwidthPackageId(v string) *DescribeCommonBandwidthPackagesRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetDryRun(v bool) *DescribeCommonBandwidthPackagesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetIncludeReservationData(v bool) *DescribeCommonBandwidthPackagesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetName(v string) *DescribeCommonBandwidthPackagesRequest {
	s.Name = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetOwnerAccount(v string) *DescribeCommonBandwidthPackagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetOwnerId(v int64) *DescribeCommonBandwidthPackagesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetPageNumber(v int32) *DescribeCommonBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetPageSize(v int32) *DescribeCommonBandwidthPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetRegionId(v string) *DescribeCommonBandwidthPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetResourceGroupId(v string) *DescribeCommonBandwidthPackagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetResourceOwnerAccount(v string) *DescribeCommonBandwidthPackagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetResourceOwnerId(v int64) *DescribeCommonBandwidthPackagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetSecurityProtectionEnabled(v bool) *DescribeCommonBandwidthPackagesRequest {
	s.SecurityProtectionEnabled = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) SetTag(v []*DescribeCommonBandwidthPackagesRequestTag) *DescribeCommonBandwidthPackagesRequest {
	s.Tag = v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequest) Validate() error {
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

type DescribeCommonBandwidthPackagesRequestTag struct {
	// The tag key to add to the Internet Shared Bandwidth instance. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value to add to the Internet Shared Bandwidth instance. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ValueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommonBandwidthPackagesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonBandwidthPackagesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCommonBandwidthPackagesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCommonBandwidthPackagesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCommonBandwidthPackagesRequestTag) SetKey(v string) *DescribeCommonBandwidthPackagesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequestTag) SetValue(v string) *DescribeCommonBandwidthPackagesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCommonBandwidthPackagesRequestTag) Validate() error {
	return dara.Validate(s)
}
