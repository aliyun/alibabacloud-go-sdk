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
	// The ID of the shared cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-fh0vdzyh6rdc*****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Specifies whether the file system uses a KMS-managed key to encrypt data at rest. Encrypted data does not need to be decrypted during read and write operations.
	//
	// > NAS-managed keys are supported when `FileSystemType` is set to `standard` or `extreme`. User-managed keys are supported when `FileSystemType` is set to `extreme`.
	//
	// example:
	//
	// 0
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The list of user IDs.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The name of the NAS file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// szy-asp-upm-test
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// The ID of the office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing+dir-15657*****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage specification type of the NAS file system.
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
