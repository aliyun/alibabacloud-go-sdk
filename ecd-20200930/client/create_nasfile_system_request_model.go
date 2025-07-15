// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNASFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNASFileSystemRequest
	GetDescription() *string
	SetEncryptType(v string) *CreateNASFileSystemRequest
	GetEncryptType() *string
	SetName(v string) *CreateNASFileSystemRequest
	GetName() *string
	SetOfficeSiteId(v string) *CreateNASFileSystemRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *CreateNASFileSystemRequest
	GetRegionId() *string
	SetStorageType(v string) *CreateNASFileSystemRequest
	GetStorageType() *string
}

type CreateNASFileSystemRequest struct {
	// Description of the NAS file system.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the file system is encrypted. Uses KMS service-managed keys to encrypt the file system\\"s on-disk data. No decryption is required when reading and writing encrypted data. Possible values and their meanings:
	//
	// - 0: Not encrypted.
	//
	// - 1: Encrypted using NAS-managed keys.
	//
	// Default value: 0
	//
	// example:
	//
	// 0
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// Name of the NAS file system.
	//
	// The file name must follow these rules:
	//
	// - Length: 2 to 128 English or Chinese characters.
	//
	// - Must start with an uppercase or lowercase letter or a Chinese character, cannot start with http:// or https://.
	//
	// - Can include numbers, underscores (_), or hyphens (-).
	//
	// example:
	//
	// testNAS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Storage specification type of the NAS file system. Allowed values:
	//
	// - Capacity: Capacity type.
	//
	// - Performance: Performance type.
	//
	// Default value: Capacity
	//
	// example:
	//
	// Capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateNASFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNASFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNASFileSystemRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *CreateNASFileSystemRequest) GetName() *string {
	return s.Name
}

func (s *CreateNASFileSystemRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateNASFileSystemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNASFileSystemRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateNASFileSystemRequest) SetDescription(v string) *CreateNASFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetEncryptType(v string) *CreateNASFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetName(v string) *CreateNASFileSystemRequest {
	s.Name = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetOfficeSiteId(v string) *CreateNASFileSystemRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetRegionId(v string) *CreateNASFileSystemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetStorageType(v string) *CreateNASFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateNASFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
