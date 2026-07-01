// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyExRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyExRequest
	GetAutoSnapshotPolicyId() *string
	SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPolicyExRequest
	GetAutoSnapshotPolicyName() *string
	SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoSnapshotPolicyExRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotPolicyExRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoSnapshotPolicyExRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAutoSnapshotPolicyExRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeAutoSnapshotPolicyExRequestTag) *DescribeAutoSnapshotPolicyExRequest
	GetTag() []*DescribeAutoSnapshotPolicyExRequestTag
}

type DescribeAutoSnapshotPolicyExRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-bp67acfmxazb4ph****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). It must support characters in the Unicode letter category, which includes characters from various languages such as English and Chinese.
	//
	// example:
	//
	// TestName
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the automatic snapshot policy list.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page when automatic snapshot policies are displayed in paging mode.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the automatic snapshot policies that you want to query. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. When you use this parameter to filter resources, the resource count cannot exceed 1000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-aek2kkmhmhs****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeAutoSnapshotPolicyExRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPolicyExRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetTag() []*DescribeAutoSnapshotPolicyExRequestTag {
	return s.Tag
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetPageNumber(v int32) *DescribeAutoSnapshotPolicyExRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetPageSize(v int32) *DescribeAutoSnapshotPolicyExRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetRegionId(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetResourceGroupId(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetTag(v []*DescribeAutoSnapshotPolicyExRequestTag) *DescribeAutoSnapshotPolicyExRequest {
	s.Tag = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) Validate() error {
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

type DescribeAutoSnapshotPolicyExRequestTag struct {
	// The tag key of the automatic snapshot policy. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the automatic snapshot policy. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) SetKey(v string) *DescribeAutoSnapshotPolicyExRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) SetValue(v string) *DescribeAutoSnapshotPolicyExRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) Validate() error {
	return dara.Validate(s)
}
