// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberName(v string) *ListDeviceGroupRequest
	GetMemberName() *string
	SetName(v string) *ListDeviceGroupRequest
	GetName() *string
	SetNum(v int32) *ListDeviceGroupRequest
	GetNum() *int32
	SetRegionId(v string) *ListDeviceGroupRequest
	GetRegionId() *string
	SetSize(v int32) *ListDeviceGroupRequest
	GetSize() *int32
	SetStatus(v string) *ListDeviceGroupRequest
	GetStatus() *string
}

type ListDeviceGroupRequest struct {
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size   *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupRequest) GetMemberName() *string {
	return s.MemberName
}

func (s *ListDeviceGroupRequest) GetName() *string {
	return s.Name
}

func (s *ListDeviceGroupRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListDeviceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDeviceGroupRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListDeviceGroupRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDeviceGroupRequest) SetMemberName(v string) *ListDeviceGroupRequest {
	s.MemberName = &v
	return s
}

func (s *ListDeviceGroupRequest) SetName(v string) *ListDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *ListDeviceGroupRequest) SetNum(v int32) *ListDeviceGroupRequest {
	s.Num = &v
	return s
}

func (s *ListDeviceGroupRequest) SetRegionId(v string) *ListDeviceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListDeviceGroupRequest) SetSize(v int32) *ListDeviceGroupRequest {
	s.Size = &v
	return s
}

func (s *ListDeviceGroupRequest) SetStatus(v string) *ListDeviceGroupRequest {
	s.Status = &v
	return s
}

func (s *ListDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
