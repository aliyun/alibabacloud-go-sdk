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
	SetProtocolType(v string) *CreateNASFileSystemRequest
	GetProtocolType() *string
	SetRegionId(v string) *CreateNASFileSystemRequest
	GetRegionId() *string
	SetStorageType(v string) *CreateNASFileSystemRequest
	GetStorageType() *string
}

type CreateNASFileSystemRequest struct {
	// The description of the NAS file system.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the file system uses a KMS-managed key to encrypt data stored on the file system. Encrypted data does not need to be decrypted during read and write operations.
	//
	// example:
	//
	// 0
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The NAS file system name. The name must meet the following requirements: The name must be 2 to 128 characters in length and can contain letters and Chinese characters. The name must start with a letter or a Chinese character and cannot start with `http://` or `https://`. The name can contain digits, underscores (_), or hyphens (-).
	//
	// example:
	//
	// testNAS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage type of the NAS file system.
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

func (s *CreateNASFileSystemRequest) GetProtocolType() *string {
	return s.ProtocolType
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

func (s *CreateNASFileSystemRequest) SetProtocolType(v string) *CreateNASFileSystemRequest {
	s.ProtocolType = &v
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
