// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *DetachEndUserRequest
	GetAdDomain() *string
	SetClientType(v string) *DetachEndUserRequest
	GetClientType() *string
	SetDeviceId(v string) *DetachEndUserRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *DetachEndUserRequest
	GetDirectoryId() *string
	SetEndUserId(v string) *DetachEndUserRequest
	GetEndUserId() *string
	SetRegion(v string) *DetachEndUserRequest
	GetRegion() *string
}

type DetachEndUserRequest struct {
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
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The serial number (SN) of the hardware client.
	//
	// This parameter is required.
	//
	// example:
	//
	// F9E52EDCCB2B****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The ID of the convenient office network.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user that you want to unbind from the hardware client.
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
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DetachEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachEndUserRequest) GoString() string {
	return s.String()
}

func (s *DetachEndUserRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *DetachEndUserRequest) GetClientType() *string {
	return s.ClientType
}

func (s *DetachEndUserRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DetachEndUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DetachEndUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DetachEndUserRequest) GetRegion() *string {
	return s.Region
}

func (s *DetachEndUserRequest) SetAdDomain(v string) *DetachEndUserRequest {
	s.AdDomain = &v
	return s
}

func (s *DetachEndUserRequest) SetClientType(v string) *DetachEndUserRequest {
	s.ClientType = &v
	return s
}

func (s *DetachEndUserRequest) SetDeviceId(v string) *DetachEndUserRequest {
	s.DeviceId = &v
	return s
}

func (s *DetachEndUserRequest) SetDirectoryId(v string) *DetachEndUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *DetachEndUserRequest) SetEndUserId(v string) *DetachEndUserRequest {
	s.EndUserId = &v
	return s
}

func (s *DetachEndUserRequest) SetRegion(v string) *DetachEndUserRequest {
	s.Region = &v
	return s
}

func (s *DetachEndUserRequest) Validate() error {
	return dara.Validate(s)
}
