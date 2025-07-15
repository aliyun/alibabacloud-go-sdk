// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *AttachEndUserRequest
	GetAdDomain() *string
	SetClientType(v int32) *AttachEndUserRequest
	GetClientType() *int32
	SetDeviceId(v string) *AttachEndUserRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *AttachEndUserRequest
	GetDirectoryId() *string
	SetEndUserId(v string) *AttachEndUserRequest
	GetEndUserId() *string
	SetRegionId(v string) *AttachEndUserRequest
	GetRegionId() *string
	SetUserType(v string) *AttachEndUserRequest
	GetUserType() *string
}

type AttachEndUserRequest struct {
	// The address of the Active Directory (AD) office network.
	//
	// example:
	//
	// xn--0zw****
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The type of the client.
	//
	// Valid values:
	//
	// 	- 1: hardware client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The serial number (SN) of the hardware client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111810122200F0C24CF7F1BF-*05AY****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The ID of the convenient office network.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user that you want to bind to the hardware client.
	//
	// This parameter is required.
	//
	// example:
	//
	// moli
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The account type of the user.
	//
	// Valid values:
	//
	// 	- AD: enterprise AD account.
	//
	// 	- SIMPLE: convenience account
	//
	// example:
	//
	// SIMPLE
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AttachEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachEndUserRequest) GoString() string {
	return s.String()
}

func (s *AttachEndUserRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *AttachEndUserRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *AttachEndUserRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AttachEndUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *AttachEndUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *AttachEndUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachEndUserRequest) GetUserType() *string {
	return s.UserType
}

func (s *AttachEndUserRequest) SetAdDomain(v string) *AttachEndUserRequest {
	s.AdDomain = &v
	return s
}

func (s *AttachEndUserRequest) SetClientType(v int32) *AttachEndUserRequest {
	s.ClientType = &v
	return s
}

func (s *AttachEndUserRequest) SetDeviceId(v string) *AttachEndUserRequest {
	s.DeviceId = &v
	return s
}

func (s *AttachEndUserRequest) SetDirectoryId(v string) *AttachEndUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *AttachEndUserRequest) SetEndUserId(v string) *AttachEndUserRequest {
	s.EndUserId = &v
	return s
}

func (s *AttachEndUserRequest) SetRegionId(v string) *AttachEndUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachEndUserRequest) SetUserType(v string) *AttachEndUserRequest {
	s.UserType = &v
	return s
}

func (s *AttachEndUserRequest) Validate() error {
	return dara.Validate(s)
}
