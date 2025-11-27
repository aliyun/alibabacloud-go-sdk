// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceRCInstanceSystemDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ReplaceRCInstanceSystemDiskRequest
	GetImageId() *string
	SetInstanceId(v string) *ReplaceRCInstanceSystemDiskRequest
	GetInstanceId() *string
	SetIsLocalDisk(v bool) *ReplaceRCInstanceSystemDiskRequest
	GetIsLocalDisk() *bool
	SetKeyPairName(v string) *ReplaceRCInstanceSystemDiskRequest
	GetKeyPairName() *string
	SetPassword(v string) *ReplaceRCInstanceSystemDiskRequest
	GetPassword() *string
	SetRegionId(v string) *ReplaceRCInstanceSystemDiskRequest
	GetRegionId() *string
}

type ReplaceRCInstanceSystemDiskRequest struct {
	// The image ID that is used when you reinstall the OS.
	//
	// example:
	//
	// m-2zec4lvlhcdkyd13****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-m5ei7b1w38w2l91x****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// None
	IsLocalDisk *bool `json:"IsLocalDisk,omitempty" xml:"IsLocalDisk,omitempty"`
	// The name of the new key pair. If you do not specify this parameter, you must reset the key pair after the OS is reinstalled.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The new logon password of the RDS Custom instance. If you do not specify this parameter, you must reset the logon password after the OS is reinstalled.
	//
	// 	- The value must be 8 to 30 characters in length.
	//
	// 	- The value must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Supported special characters include: ( ) \\` ~ ! @ # $ % ^ & \\	- - _ + =
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReplaceRCInstanceSystemDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceRCInstanceSystemDiskRequest) GoString() string {
	return s.String()
}

func (s *ReplaceRCInstanceSystemDiskRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ReplaceRCInstanceSystemDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReplaceRCInstanceSystemDiskRequest) GetIsLocalDisk() *bool {
	return s.IsLocalDisk
}

func (s *ReplaceRCInstanceSystemDiskRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ReplaceRCInstanceSystemDiskRequest) GetPassword() *string {
	return s.Password
}

func (s *ReplaceRCInstanceSystemDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReplaceRCInstanceSystemDiskRequest) SetImageId(v string) *ReplaceRCInstanceSystemDiskRequest {
	s.ImageId = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskRequest) SetInstanceId(v string) *ReplaceRCInstanceSystemDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskRequest) SetIsLocalDisk(v bool) *ReplaceRCInstanceSystemDiskRequest {
	s.IsLocalDisk = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskRequest) SetKeyPairName(v string) *ReplaceRCInstanceSystemDiskRequest {
	s.KeyPairName = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskRequest) SetPassword(v string) *ReplaceRCInstanceSystemDiskRequest {
	s.Password = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskRequest) SetRegionId(v string) *ReplaceRCInstanceSystemDiskRequest {
	s.RegionId = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskRequest) Validate() error {
	return dara.Validate(s)
}
