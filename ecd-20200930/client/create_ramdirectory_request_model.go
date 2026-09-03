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
	// The method allowed for connecting to cloud computers.
	//
	// example:
	//
	// Internet
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// The directory name. The name must be 2 to 255 characters in length and can contain letters and Chinese characters. The name must start with a letter or a Chinese character and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testDirectoryName
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// Specifies whether to grant local administrator permissions to users who use cloud computers.
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// Specifies whether to enable public network access.
	//
	// example:
	//
	// false
	EnableInternetAccess *bool `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	// The region ID. Call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch ID. Only one vSwitch can be specified.
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
