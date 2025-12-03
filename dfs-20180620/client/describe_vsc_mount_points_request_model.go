// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscMountPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeVscMountPointsRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *DescribeVscMountPointsRequest
	GetInputRegionId() *string
	SetInstanceId(v string) *DescribeVscMountPointsRequest
	GetInstanceId() *string
	SetMountPointId(v string) *DescribeVscMountPointsRequest
	GetMountPointId() *string
	SetStatus(v string) *DescribeVscMountPointsRequest
	GetStatus() *string
	SetVscId(v string) *DescribeVscMountPointsRequest
	GetVscId() *string
}

type DescribeVscMountPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 037****e1d
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// ["ecs-instance1", "ecs-instance2"]
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 037cb49e1d-c***5
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VscId  *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeVscMountPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscMountPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeVscMountPointsRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DescribeVscMountPointsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVscMountPointsRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DescribeVscMountPointsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscMountPointsRequest) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscMountPointsRequest) SetFileSystemId(v string) *DescribeVscMountPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetInputRegionId(v string) *DescribeVscMountPointsRequest {
	s.InputRegionId = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetInstanceId(v string) *DescribeVscMountPointsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetMountPointId(v string) *DescribeVscMountPointsRequest {
	s.MountPointId = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetStatus(v string) *DescribeVscMountPointsRequest {
	s.Status = &v
	return s
}

func (s *DescribeVscMountPointsRequest) SetVscId(v string) *DescribeVscMountPointsRequest {
	s.VscId = &v
	return s
}

func (s *DescribeVscMountPointsRequest) Validate() error {
	return dara.Validate(s)
}
