// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateSnapshotRequest
	GetCategory() *string
	SetClientToken(v string) *CreateSnapshotRequest
	GetClientToken() *string
	SetDescription(v string) *CreateSnapshotRequest
	GetDescription() *string
	SetDiskId(v string) *CreateSnapshotRequest
	GetDiskId() *string
	SetInstantAccess(v bool) *CreateSnapshotRequest
	GetInstantAccess() *bool
	SetInstantAccessRetentionDays(v int32) *CreateSnapshotRequest
	GetInstantAccessRetentionDays() *int32
	SetOwnerAccount(v string) *CreateSnapshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSnapshotRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateSnapshotRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateSnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSnapshotRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *CreateSnapshotRequest
	GetRetentionDays() *int32
	SetSnapshotName(v string) *CreateSnapshotRequest
	GetSnapshotName() *string
	SetStorageLocationArn(v string) *CreateSnapshotRequest
	GetStorageLocationArn() *string
	SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest
	GetTag() []*CreateSnapshotRequestTag
}

type CreateSnapshotRequest struct {
	// The snapshot type. Valid values:
	//
	// - Standard: standard snapshot.
	//
	// - Flash: local snapshot.
	//
	// > This parameter will be deprecated. Standard snapshots of enterprise SSDs have been upgraded to [instant access by default](https://help.aliyun.com/document_detail/193667.html). No additional configuration or fees are required. This applies to enterprise SSDs, ESSD AutoPL disks, ESSD Entry disks, and regional enterprise SSDs. Standard snapshots of standard SSDs are also active by default.
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the snapshot. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: null.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1s5fnvk4gn2tws0****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to enable the snapshot instant access feature. Valid values:
	//
	// - true: enables the feature. Only enterprise SSDs support this feature.
	//
	// - false: shutdown. A standard snapshot is created.
	//
	// Default value: false.
	//
	// > This parameter is deprecated. Standard snapshots of enterprise SSDs have been upgraded to [instant access by default](https://help.aliyun.com/document_detail/193667.html). No additional configuration or fees are required. This applies to enterprise SSDs, ESSD AutoPL disks, ESSD Entry disks, and regional enterprise SSDs. Standard snapshots of standard SSDs are also active by default.
	//
	// example:
	//
	// false
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// Settings for the retention period of the snapshot instant access feature. After the retention period expires, the snapshot is subject to automatic release. This parameter takes effect only when `InstantAccess=true`. Unit: days. Valid values: 1 to 65535.
	//
	// Default value: the same as the value of the `RetentionDays` parameter.
	//
	// > This parameter is deprecated. Standard snapshots of enterprise SSDs have been upgraded to [instant access by default](https://help.aliyun.com/document_detail/193667.html). No additional configuration or fees are required. This applies to enterprise SSDs, ESSD AutoPL disks, ESSD Entry disks, and regional enterprise SSDs. Standard snapshots of standard SSDs are also active by default.
	//
	// example:
	//
	// 1
	InstantAccessRetentionDays *int32  `json:"InstantAccessRetentionDays,omitempty" xml:"InstantAccessRetentionDays,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the snapshot belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Settings for the retention period of the snapshot. Unit: days. Valid values: 1 to 65536. The snapshot is subject to automatic release when the retention period expires.
	//
	// Default value: null, which indicates that the snapshot is not subject to automatic release.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The name can contain Unicode characters under the letter category (including letters in English and Chinese), ASCII digits (0-9), colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// > The name cannot start with `auto` to avoid conflicts with the names of automatic snapshots.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The tags.
	Tag []*CreateSnapshotRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSnapshotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateSnapshotRequest) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *CreateSnapshotRequest) GetInstantAccessRetentionDays() *int32 {
	return s.InstantAccessRetentionDays
}

func (s *CreateSnapshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSnapshotRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSnapshotRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateSnapshotRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *CreateSnapshotRequest) GetTag() []*CreateSnapshotRequestTag {
	return s.Tag
}

func (s *CreateSnapshotRequest) SetCategory(v string) *CreateSnapshotRequest {
	s.Category = &v
	return s
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetDiskId(v string) *CreateSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateSnapshotRequest) SetInstantAccess(v bool) *CreateSnapshotRequest {
	s.InstantAccess = &v
	return s
}

func (s *CreateSnapshotRequest) SetInstantAccessRetentionDays(v int32) *CreateSnapshotRequest {
	s.InstantAccessRetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetOwnerAccount(v string) *CreateSnapshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSnapshotRequest) SetOwnerId(v int64) *CreateSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceGroupId(v string) *CreateSnapshotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceOwnerAccount(v string) *CreateSnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceOwnerId(v int64) *CreateSnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetStorageLocationArn(v string) *CreateSnapshotRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *CreateSnapshotRequest) SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest {
	s.Tag = v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
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

type CreateSnapshotRequestTag struct {
	// The tag key of the snapshot. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSnapshotRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSnapshotRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSnapshotRequestTag) SetKey(v string) *CreateSnapshotRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSnapshotRequestTag) SetValue(v string) *CreateSnapshotRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSnapshotRequestTag) Validate() error {
	return dara.Validate(s)
}
