// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedHostClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterName(v string) *CreateDedicatedHostClusterRequest
	GetDedicatedHostClusterName() *string
	SetDescription(v string) *CreateDedicatedHostClusterRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateDedicatedHostClusterRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateDedicatedHostClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDedicatedHostClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDedicatedHostClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDedicatedHostClusterRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDedicatedHostClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDedicatedHostClusterRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateDedicatedHostClusterRequestTag) *CreateDedicatedHostClusterRequest
	GetTag() []*CreateDedicatedHostClusterRequestTag
	SetZoneId(v string) *CreateDedicatedHostClusterRequest
	GetZoneId() *string
}

type CreateDedicatedHostClusterRequest struct {
	// The name of the dedicated host cluster. The name must be 2 to 128 characters in length and can contain letters, digits, and Unicode characters that are categorized as letter characters. The name can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// Default value: null.
	//
	// example:
	//
	// myDDHCluster
	DedicatedHostClusterName *string `json:"DedicatedHostClusterName,omitempty" xml:"DedicatedHostClusterName,omitempty"`
	// The description of the dedicated host cluster. The description must be 2 to 256 characters in length. The description cannot start with `http://` or `https://`.
	//
	// Default value: null.
	//
	// example:
	//
	// This-is-my-DDHCluster
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: sends a check request without querying resource status. The system checks the request for potential issues, including invalid AccessKey pairs, the authorization status of the Resource Access Management (RAM) user, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - false: sends a Normal request. If the request passes the check, a 2XX HTTP status code is returned and the resource status is queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the dedicated host cluster. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated host cluster belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreateDedicatedHostClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the dedicated host cluster. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zones in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDedicatedHostClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedHostClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostClusterRequest) GetDedicatedHostClusterName() *string {
	return s.DedicatedHostClusterName
}

func (s *CreateDedicatedHostClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDedicatedHostClusterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDedicatedHostClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDedicatedHostClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDedicatedHostClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDedicatedHostClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDedicatedHostClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDedicatedHostClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDedicatedHostClusterRequest) GetTag() []*CreateDedicatedHostClusterRequestTag {
	return s.Tag
}

func (s *CreateDedicatedHostClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDedicatedHostClusterRequest) SetDedicatedHostClusterName(v string) *CreateDedicatedHostClusterRequest {
	s.DedicatedHostClusterName = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetDescription(v string) *CreateDedicatedHostClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetDryRun(v bool) *CreateDedicatedHostClusterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetOwnerAccount(v string) *CreateDedicatedHostClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetOwnerId(v int64) *CreateDedicatedHostClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetRegionId(v string) *CreateDedicatedHostClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetResourceGroupId(v string) *CreateDedicatedHostClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetResourceOwnerAccount(v string) *CreateDedicatedHostClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetResourceOwnerId(v int64) *CreateDedicatedHostClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetTag(v []*CreateDedicatedHostClusterRequestTag) *CreateDedicatedHostClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDedicatedHostClusterRequest) SetZoneId(v string) *CreateDedicatedHostClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDedicatedHostClusterRequest) Validate() error {
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

type CreateDedicatedHostClusterRequestTag struct {
	// The tag key of the dedicated host cluster. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated host cluster. Valid values of N: 1 to 20. The tag value cannot be an empty string. The tag value can be up to 64 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDedicatedHostClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedHostClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDedicatedHostClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDedicatedHostClusterRequestTag) SetKey(v string) *CreateDedicatedHostClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDedicatedHostClusterRequestTag) SetValue(v string) *CreateDedicatedHostClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDedicatedHostClusterRequestTag) Validate() error {
	return dara.Validate(s)
}
