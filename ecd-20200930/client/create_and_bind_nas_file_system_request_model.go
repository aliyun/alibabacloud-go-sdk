// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndBindNasFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAndBindNasFileSystemRequest
	GetDescription() *string
	SetDesktopGroupId(v string) *CreateAndBindNasFileSystemRequest
	GetDesktopGroupId() *string
	SetEncryptType(v int32) *CreateAndBindNasFileSystemRequest
	GetEncryptType() *int32
	SetEndUserIds(v []*string) *CreateAndBindNasFileSystemRequest
	GetEndUserIds() []*string
	SetFileSystemName(v string) *CreateAndBindNasFileSystemRequest
	GetFileSystemName() *string
	SetOfficeSiteId(v string) *CreateAndBindNasFileSystemRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *CreateAndBindNasFileSystemRequest
	GetRegionId() *string
	SetStorageType(v string) *CreateAndBindNasFileSystemRequest
	GetStorageType() *string
}

type CreateAndBindNasFileSystemRequest struct {
	// The description of the NAS file system.
	//
	// example:
	//
	// newDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the desktop group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-fh0vdzyh6rdc*****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Specifies whether to encrypt data in the NAS file system. You can use keys that are hosted by Key Management Service (KMS) to encrypt data in a file system. When you read and write the encrypted data, the data is automatically decrypted. Valid values:
	//
	// 	- 0: does not encrypt data in the NAS file system.
	//
	// 	- 1: encrypts data in the NAS file system by using a NAS-managed key. ` If you set  `FileSystemType`  to  `standard`  or  `extreme`, you can use a NAS-managed key to encrypt data in a NAS file system.`
	//
	// 	- 2: encrypts data in the NAS file system by using a KMS-managed key. `If` you set FileSystemType`  to  `extreme`, you can use a KMS-managed key to encrypt data in a NAS file system.`
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The list of users.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The name of the NAS file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// szy-asp-upm-test
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing+dir-15657*****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage type of the NAS file system. Valid values:
	//
	// 	- Capacity
	//
	// 	- Performance
	//
	// Default value: Capacity.
	//
	// This parameter is required.
	//
	// example:
	//
	// Capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateAndBindNasFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindNasFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateAndBindNasFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAndBindNasFileSystemRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *CreateAndBindNasFileSystemRequest) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *CreateAndBindNasFileSystemRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *CreateAndBindNasFileSystemRequest) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *CreateAndBindNasFileSystemRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateAndBindNasFileSystemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAndBindNasFileSystemRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateAndBindNasFileSystemRequest) SetDescription(v string) *CreateAndBindNasFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetDesktopGroupId(v string) *CreateAndBindNasFileSystemRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetEncryptType(v int32) *CreateAndBindNasFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetEndUserIds(v []*string) *CreateAndBindNasFileSystemRequest {
	s.EndUserIds = v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetFileSystemName(v string) *CreateAndBindNasFileSystemRequest {
	s.FileSystemName = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetOfficeSiteId(v string) *CreateAndBindNasFileSystemRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetRegionId(v string) *CreateAndBindNasFileSystemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) SetStorageType(v string) *CreateAndBindNasFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateAndBindNasFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
