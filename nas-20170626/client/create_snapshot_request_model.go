// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSnapshotRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateSnapshotRequest
	GetFileSystemId() *string
	SetRetentionDays(v int32) *CreateSnapshotRequest
	GetRetentionDays() *int32
	SetSnapshotName(v string) *CreateSnapshotRequest
	GetSnapshotName() *string
}

type CreateSnapshotRequest struct {
	// The snapshot description.
	//
	// Limits:
	//
	// - The description must be 2 to 256 characters in length.
	//
	// - The description cannot start with `http://` or `https://`.
	//
	// - Default value: empty.
	//
	// example:
	//
	// FinanceDepet
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Advanced Extreme NAS file system. The value must start with `extreme-`, such as `extreme-01dd****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme-01dd****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The retention period of the snapshot.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// - -1 (default): The snapshot is permanently retained and is automatically deleted when the snapshot quota is reached.
	//
	// - 1 to 65536: The snapshot is retained for the specified number of days. The snapshot is subject to automatic release after the retention period expires.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot name.
	//
	// Limits:
	//
	// - The name must be 2 to 128 characters in length and can contain letters and digits. The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// - The name can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// - The name cannot start with auto to avoid conflicts with automatic snapshot names.
	//
	// example:
	//
	// FinanceJoshua
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateSnapshotRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetFileSystemId(v string) *CreateSnapshotRequest {
	s.FileSystemId = &v
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

func (s *CreateSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
