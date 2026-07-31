// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySnapshotAttributeRequest
	GetDescription() *string
	SetDisableInstantAccess(v bool) *ModifySnapshotAttributeRequest
	GetDisableInstantAccess() *bool
	SetOwnerAccount(v string) *ModifySnapshotAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySnapshotAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySnapshotAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySnapshotAttributeRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *ModifySnapshotAttributeRequest
	GetRetentionDays() *int32
	SetSnapshotId(v string) *ModifySnapshotAttributeRequest
	GetSnapshotId() *string
	SetSnapshotName(v string) *ModifySnapshotAttributeRequest
	GetSnapshotName() *string
}

type ModifySnapshotAttributeRequest struct {
	// The description of the snapshot. The description can be empty and can be up to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable the snapshot instant access feature. Valid values:
	//
	// - true: Disables the snapshot instant access feature.
	//
	// - false: Does not disable the snapshot instant access feature.
	//
	// Default value: false.
	//
	// >This parameter is deprecated. Standard snapshots of enterprise SSDs have been upgraded to [instant access by default](https://help.aliyun.com/document_detail/193667.html). No additional configuration or fees are required.
	//
	// example:
	//
	// false
	DisableInstantAccess *bool   `json:"DisableInstantAccess,omitempty" xml:"DisableInstantAccess,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of days for which the snapshot is retained. The retention period is calculated from the snapshot **creation time*	- (represented in the ISO 8601 standard and in UTC+0 time in the yyyy-MM-ddTHH:mm:ssZ format). Valid values: 1 to 65536.
	//
	// >The snapshot retention period can only be extended. Shortening the existing retention period of a snapshot is not supported.
	//
	// example:
	//
	// 10
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp199lyny9bb47pa****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The display name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. The name can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// The name cannot start with auto to avoid conflicts with automatic snapshot names.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s ModifySnapshotAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySnapshotAttributeRequest) GetDisableInstantAccess() *bool {
	return s.DisableInstantAccess
}

func (s *ModifySnapshotAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySnapshotAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySnapshotAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySnapshotAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySnapshotAttributeRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifySnapshotAttributeRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifySnapshotAttributeRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *ModifySnapshotAttributeRequest) SetDescription(v string) *ModifySnapshotAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetDisableInstantAccess(v bool) *ModifySnapshotAttributeRequest {
	s.DisableInstantAccess = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetOwnerAccount(v string) *ModifySnapshotAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetOwnerId(v int64) *ModifySnapshotAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetResourceOwnerAccount(v string) *ModifySnapshotAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetResourceOwnerId(v int64) *ModifySnapshotAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetRetentionDays(v int32) *ModifySnapshotAttributeRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotId(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotName(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotName = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) Validate() error {
	return dara.Validate(s)
}
