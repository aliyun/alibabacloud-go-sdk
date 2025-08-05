// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNasFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DetachNasFileSystemRequest
	GetCreateTime() *string
	SetCrossAccountRoleName(v string) *DetachNasFileSystemRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *DetachNasFileSystemRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *DetachNasFileSystemRequest
	GetCrossAccountUserId() *int64
	SetFileSystemId(v string) *DetachNasFileSystemRequest
	GetFileSystemId() *string
}

type DetachNasFileSystemRequest struct {
	// The time when the file system was created. The value must be a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1607436917
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up and restore data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up and restored within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up and restored within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up and restored across Alibaba Cloud accounts.
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up and restore data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 005494
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DetachNasFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachNasFileSystemRequest) GoString() string {
	return s.String()
}

func (s *DetachNasFileSystemRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DetachNasFileSystemRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DetachNasFileSystemRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DetachNasFileSystemRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DetachNasFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DetachNasFileSystemRequest) SetCreateTime(v string) *DetachNasFileSystemRequest {
	s.CreateTime = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetCrossAccountRoleName(v string) *DetachNasFileSystemRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetCrossAccountType(v string) *DetachNasFileSystemRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetCrossAccountUserId(v int64) *DetachNasFileSystemRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DetachNasFileSystemRequest) SetFileSystemId(v string) *DetachNasFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *DetachNasFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
