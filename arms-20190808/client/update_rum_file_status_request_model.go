// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRumFileStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *UpdateRumFileStatusRequest
	GetFileName() *string
	SetPid(v string) *UpdateRumFileStatusRequest
	GetPid() *string
	SetRegionId(v string) *UpdateRumFileStatusRequest
	GetRegionId() *string
	SetSize(v string) *UpdateRumFileStatusRequest
	GetSize() *string
	SetStatus(v string) *UpdateRumFileStatusRequest
	GetStatus() *string
	SetUuid(v string) *UpdateRumFileStatusRequest
	GetUuid() *string
	SetVersionId(v string) *UpdateRumFileStatusRequest
	GetVersionId() *string
}

type UpdateRumFileStatusRequest struct {
	// The file name.
	//
	// example:
	//
	// test.js.map
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The application ID.
	//
	// example:
	//
	// atc8xxxx
	//
	// cf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 20
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the file. Valid values: SUCCESS and INIT.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique ID of the file. If you do not set this parameter, the system automatically sets a UUID for you.
	//
	// example:
	//
	// MS4wLjAtbWFpbi4wZjM0NzRlOSxxxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The version number of the file.
	//
	// example:
	//
	// 1.0.0
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateRumFileStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRumFileStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRumFileStatusRequest) GetFileName() *string {
	return s.FileName
}

func (s *UpdateRumFileStatusRequest) GetPid() *string {
	return s.Pid
}

func (s *UpdateRumFileStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateRumFileStatusRequest) GetSize() *string {
	return s.Size
}

func (s *UpdateRumFileStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateRumFileStatusRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateRumFileStatusRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateRumFileStatusRequest) SetFileName(v string) *UpdateRumFileStatusRequest {
	s.FileName = &v
	return s
}

func (s *UpdateRumFileStatusRequest) SetPid(v string) *UpdateRumFileStatusRequest {
	s.Pid = &v
	return s
}

func (s *UpdateRumFileStatusRequest) SetRegionId(v string) *UpdateRumFileStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateRumFileStatusRequest) SetSize(v string) *UpdateRumFileStatusRequest {
	s.Size = &v
	return s
}

func (s *UpdateRumFileStatusRequest) SetStatus(v string) *UpdateRumFileStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateRumFileStatusRequest) SetUuid(v string) *UpdateRumFileStatusRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateRumFileStatusRequest) SetVersionId(v string) *UpdateRumFileStatusRequest {
	s.VersionId = &v
	return s
}

func (s *UpdateRumFileStatusRequest) Validate() error {
	return dara.Validate(s)
}
