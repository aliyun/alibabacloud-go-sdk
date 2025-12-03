// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoint(v *GetMountPointResponseBodyMountPoint) *GetMountPointResponseBody
	GetMountPoint() *GetMountPointResponseBodyMountPoint
	SetRequestId(v string) *GetMountPointResponseBody
	GetRequestId() *string
}

type GetMountPointResponseBody struct {
	MountPoint *GetMountPointResponseBodyMountPoint `json:"MountPoint,omitempty" xml:"MountPoint,omitempty" type:"Struct"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetMountPointResponseBody) GetMountPoint() *GetMountPointResponseBodyMountPoint {
	return s.MountPoint
}

func (s *GetMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMountPointResponseBody) SetMountPoint(v *GetMountPointResponseBodyMountPoint) *GetMountPointResponseBody {
	s.MountPoint = v
	return s
}

func (s *GetMountPointResponseBody) SetRequestId(v string) *GetMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMountPointResponseBody) Validate() error {
	if s.MountPoint != nil {
		if err := s.MountPoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMountPointResponseBodyMountPoint struct {
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId    *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId     *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountPointAlias  *string `json:"MountPointAlias,omitempty" xml:"MountPointAlias,omitempty"`
	MountPointDomain *string `json:"MountPointDomain,omitempty" xml:"MountPointDomain,omitempty"`
	MountPointId     *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-iq8fymi327krd14mt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-iq8hhsk3ymzv9m4wn****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetMountPointResponseBodyMountPoint) String() string {
	return dara.Prettify(s)
}

func (s GetMountPointResponseBodyMountPoint) GoString() string {
	return s.String()
}

func (s *GetMountPointResponseBodyMountPoint) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *GetMountPointResponseBodyMountPoint) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMountPointResponseBodyMountPoint) GetDescription() *string {
	return s.Description
}

func (s *GetMountPointResponseBodyMountPoint) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetMountPointResponseBodyMountPoint) GetMountPointAlias() *string {
	return s.MountPointAlias
}

func (s *GetMountPointResponseBodyMountPoint) GetMountPointDomain() *string {
	return s.MountPointDomain
}

func (s *GetMountPointResponseBodyMountPoint) GetMountPointId() *string {
	return s.MountPointId
}

func (s *GetMountPointResponseBodyMountPoint) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetMountPointResponseBodyMountPoint) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMountPointResponseBodyMountPoint) GetStatus() *string {
	return s.Status
}

func (s *GetMountPointResponseBodyMountPoint) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetMountPointResponseBodyMountPoint) GetVpcId() *string {
	return s.VpcId
}

func (s *GetMountPointResponseBodyMountPoint) SetAccessGroupId(v string) *GetMountPointResponseBodyMountPoint {
	s.AccessGroupId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetCreateTime(v string) *GetMountPointResponseBodyMountPoint {
	s.CreateTime = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetDescription(v string) *GetMountPointResponseBodyMountPoint {
	s.Description = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetFileSystemId(v string) *GetMountPointResponseBodyMountPoint {
	s.FileSystemId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetMountPointAlias(v string) *GetMountPointResponseBodyMountPoint {
	s.MountPointAlias = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetMountPointDomain(v string) *GetMountPointResponseBodyMountPoint {
	s.MountPointDomain = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetMountPointId(v string) *GetMountPointResponseBodyMountPoint {
	s.MountPointId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetNetworkType(v string) *GetMountPointResponseBodyMountPoint {
	s.NetworkType = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetRegionId(v string) *GetMountPointResponseBodyMountPoint {
	s.RegionId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetStatus(v string) *GetMountPointResponseBodyMountPoint {
	s.Status = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetVSwitchId(v string) *GetMountPointResponseBodyMountPoint {
	s.VSwitchId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) SetVpcId(v string) *GetMountPointResponseBodyMountPoint {
	s.VpcId = &v
	return s
}

func (s *GetMountPointResponseBodyMountPoint) Validate() error {
	return dara.Validate(s)
}
