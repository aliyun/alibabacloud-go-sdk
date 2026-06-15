// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifySnapshotCategoryRequest
	GetCategory() *string
	SetOwnerAccount(v string) *ModifySnapshotCategoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySnapshotCategoryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySnapshotCategoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySnapshotCategoryRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *ModifySnapshotCategoryRequest
	GetRetentionDays() *int32
	SetSnapshotId(v string) *ModifySnapshotCategoryRequest
	GetSnapshotId() *string
}

type ModifySnapshotCategoryRequest struct {
	// The snapshot type.
	//
	// - Archive: archive snapshot.
	//
	// example:
	//
	// Archive
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of days for which the snapshot is retained. The retention period starts from the snapshot creation time (CreationTime). A standard snapshot must have been retained for at least 14 days after creation before it can be archived.
	//
	// Archive snapshots must be retained for at least 60 days. When the retention period of an archive snapshot is calculated, the retention period of the standard snapshot is deducted. If an archive snapshot is deleted before 60 days, you are charged for 60 days of archive storage. For more information, see [Snapshot billing](https://help.aliyun.com/document_detail/56159.html).
	//
	// Valid values: [74, 65536].
	//
	// > If you do not specify this parameter, the snapshot is permanently retained.
	//
	// example:
	//
	// 100
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-123**sd
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ModifySnapshotCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotCategoryRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotCategoryRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifySnapshotCategoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySnapshotCategoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySnapshotCategoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySnapshotCategoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySnapshotCategoryRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifySnapshotCategoryRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifySnapshotCategoryRequest) SetCategory(v string) *ModifySnapshotCategoryRequest {
	s.Category = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetOwnerAccount(v string) *ModifySnapshotCategoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetOwnerId(v int64) *ModifySnapshotCategoryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetResourceOwnerAccount(v string) *ModifySnapshotCategoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetResourceOwnerId(v int64) *ModifySnapshotCategoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetRetentionDays(v int32) *ModifySnapshotCategoryRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetSnapshotId(v string) *ModifySnapshotCategoryRequest {
	s.SnapshotId = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) Validate() error {
	return dara.Validate(s)
}
