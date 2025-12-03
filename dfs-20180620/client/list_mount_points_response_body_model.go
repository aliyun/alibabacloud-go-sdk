// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMountPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoints(v []*ListMountPointsResponseBodyMountPoints) *ListMountPointsResponseBody
	GetMountPoints() []*ListMountPointsResponseBodyMountPoints
	SetNextToken(v string) *ListMountPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMountPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMountPointsResponseBody
	GetTotalCount() *int32
}

type ListMountPointsResponseBody struct {
	MountPoints []*ListMountPointsResponseBodyMountPoints `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMountPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMountPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMountPointsResponseBody) GetMountPoints() []*ListMountPointsResponseBodyMountPoints {
	return s.MountPoints
}

func (s *ListMountPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMountPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMountPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMountPointsResponseBody) SetMountPoints(v []*ListMountPointsResponseBodyMountPoints) *ListMountPointsResponseBody {
	s.MountPoints = v
	return s
}

func (s *ListMountPointsResponseBody) SetNextToken(v string) *ListMountPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMountPointsResponseBody) SetRequestId(v string) *ListMountPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMountPointsResponseBody) SetTotalCount(v int32) *ListMountPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMountPointsResponseBody) Validate() error {
	if s.MountPoints != nil {
		for _, item := range s.MountPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMountPointsResponseBodyMountPoints struct {
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

func (s ListMountPointsResponseBodyMountPoints) String() string {
	return dara.Prettify(s)
}

func (s ListMountPointsResponseBodyMountPoints) GoString() string {
	return s.String()
}

func (s *ListMountPointsResponseBodyMountPoints) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ListMountPointsResponseBodyMountPoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMountPointsResponseBodyMountPoints) GetDescription() *string {
	return s.Description
}

func (s *ListMountPointsResponseBodyMountPoints) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListMountPointsResponseBodyMountPoints) GetMountPointAlias() *string {
	return s.MountPointAlias
}

func (s *ListMountPointsResponseBodyMountPoints) GetMountPointDomain() *string {
	return s.MountPointDomain
}

func (s *ListMountPointsResponseBodyMountPoints) GetMountPointId() *string {
	return s.MountPointId
}

func (s *ListMountPointsResponseBodyMountPoints) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListMountPointsResponseBodyMountPoints) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMountPointsResponseBodyMountPoints) GetStatus() *string {
	return s.Status
}

func (s *ListMountPointsResponseBodyMountPoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListMountPointsResponseBodyMountPoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListMountPointsResponseBodyMountPoints) SetAccessGroupId(v string) *ListMountPointsResponseBodyMountPoints {
	s.AccessGroupId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetCreateTime(v string) *ListMountPointsResponseBodyMountPoints {
	s.CreateTime = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetDescription(v string) *ListMountPointsResponseBodyMountPoints {
	s.Description = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetFileSystemId(v string) *ListMountPointsResponseBodyMountPoints {
	s.FileSystemId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetMountPointAlias(v string) *ListMountPointsResponseBodyMountPoints {
	s.MountPointAlias = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetMountPointDomain(v string) *ListMountPointsResponseBodyMountPoints {
	s.MountPointDomain = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetMountPointId(v string) *ListMountPointsResponseBodyMountPoints {
	s.MountPointId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetNetworkType(v string) *ListMountPointsResponseBodyMountPoints {
	s.NetworkType = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetRegionId(v string) *ListMountPointsResponseBodyMountPoints {
	s.RegionId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetStatus(v string) *ListMountPointsResponseBodyMountPoints {
	s.Status = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetVSwitchId(v string) *ListMountPointsResponseBodyMountPoints {
	s.VSwitchId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) SetVpcId(v string) *ListMountPointsResponseBodyMountPoints {
	s.VpcId = &v
	return s
}

func (s *ListMountPointsResponseBodyMountPoints) Validate() error {
	return dara.Validate(s)
}
