// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRAMDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopAccessType(v string) *CreateRAMDirectoryRequest
	GetDesktopAccessType() *string
	SetDirectoryName(v string) *CreateRAMDirectoryRequest
	GetDirectoryName() *string
	SetEnableAdminAccess(v bool) *CreateRAMDirectoryRequest
	GetEnableAdminAccess() *bool
	SetEnableInternetAccess(v bool) *CreateRAMDirectoryRequest
	GetEnableInternetAccess() *bool
	SetRegionId(v string) *CreateRAMDirectoryRequest
	GetRegionId() *string
	SetVSwitchId(v []*string) *CreateRAMDirectoryRequest
	GetVSwitchId() []*string
}

type CreateRAMDirectoryRequest struct {
	// The method in which the cloud computer is connected.
	//
	// Valid values:
	//
	// 	- VPC
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Internet (default)
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Any
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The directory name. The name must be 2 to 255 characters in length. It must start with a letter but cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testDirectoryName
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// Specifies whether to grant the local administrator permissions to users that are authorized to use cloud computers in the office network.
	//
	// Valid values:
	//
	// 	- <!-- -->
	//
	//     true
	//
	//     <!-- -->
	//
	//     (default)
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable Internet access.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	EnableInternetAccess *bool `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch IDs. You can configure only one vSwitch.
	//
	// This parameter is required.
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s CreateRAMDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRAMDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryRequest) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *CreateRAMDirectoryRequest) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *CreateRAMDirectoryRequest) GetEnableAdminAccess() *bool {
	return s.EnableAdminAccess
}

func (s *CreateRAMDirectoryRequest) GetEnableInternetAccess() *bool {
	return s.EnableInternetAccess
}

func (s *CreateRAMDirectoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRAMDirectoryRequest) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *CreateRAMDirectoryRequest) SetDesktopAccessType(v string) *CreateRAMDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDirectoryName(v string) *CreateRAMDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableAdminAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableInternetAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetRegionId(v string) *CreateRAMDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetVSwitchId(v []*string) *CreateRAMDirectoryRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateRAMDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
