// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscMountPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountPoints(v []*DescribeVscMountPointsResponseBodyMountPoints) *DescribeVscMountPointsResponseBody
	GetMountPoints() []*DescribeVscMountPointsResponseBodyMountPoints
	SetRequestId(v string) *DescribeVscMountPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVscMountPointsResponseBody
	GetTotalCount() *int32
}

type DescribeVscMountPointsResponseBody struct {
	MountPoints []*DescribeVscMountPointsResponseBodyMountPoints `json:"MountPoints,omitempty" xml:"MountPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVscMountPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscMountPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBody) GetMountPoints() []*DescribeVscMountPointsResponseBodyMountPoints {
	return s.MountPoints
}

func (s *DescribeVscMountPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVscMountPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVscMountPointsResponseBody) SetMountPoints(v []*DescribeVscMountPointsResponseBodyMountPoints) *DescribeVscMountPointsResponseBody {
	s.MountPoints = v
	return s
}

func (s *DescribeVscMountPointsResponseBody) SetRequestId(v string) *DescribeVscMountPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscMountPointsResponseBody) SetTotalCount(v int32) *DescribeVscMountPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVscMountPointsResponseBody) Validate() error {
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

type DescribeVscMountPointsResponseBodyMountPoints struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	InstanceTotalCount *int32                                                    `json:"InstanceTotalCount,omitempty" xml:"InstanceTotalCount,omitempty"`
	Instances          []*DescribeVscMountPointsResponseBodyMountPointsInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// hdfs1
	MountPointAlias *string `json:"MountPointAlias,omitempty" xml:"MountPointAlias,omitempty"`
	// example:
	//
	// 037cb49e1d-c***5
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
}

func (s DescribeVscMountPointsResponseBodyMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscMountPointsResponseBodyMountPoints) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) GetDescription() *string {
	return s.Description
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) GetInstanceTotalCount() *int32 {
	return s.InstanceTotalCount
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) GetInstances() []*DescribeVscMountPointsResponseBodyMountPointsInstances {
	return s.Instances
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) GetMountPointAlias() *string {
	return s.MountPointAlias
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) SetDescription(v string) *DescribeVscMountPointsResponseBodyMountPoints {
	s.Description = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) SetInstanceTotalCount(v int32) *DescribeVscMountPointsResponseBodyMountPoints {
	s.InstanceTotalCount = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) SetInstances(v []*DescribeVscMountPointsResponseBodyMountPointsInstances) *DescribeVscMountPointsResponseBodyMountPoints {
	s.Instances = v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) SetMountPointAlias(v string) *DescribeVscMountPointsResponseBodyMountPoints {
	s.MountPointAlias = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) SetMountPointId(v string) *DescribeVscMountPointsResponseBodyMountPoints {
	s.MountPointId = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPoints) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVscMountPointsResponseBodyMountPointsInstances struct {
	// example:
	//
	// ["ecs-instance1", "ecs-instance2"]
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Vscs   []*DescribeVscMountPointsResponseBodyMountPointsInstancesVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s DescribeVscMountPointsResponseBodyMountPointsInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscMountPointsResponseBodyMountPointsInstances) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) GetVscs() []*DescribeVscMountPointsResponseBodyMountPointsInstancesVscs {
	return s.Vscs
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) SetInstanceId(v string) *DescribeVscMountPointsResponseBodyMountPointsInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) SetStatus(v string) *DescribeVscMountPointsResponseBodyMountPointsInstances {
	s.Status = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) SetVscs(v []*DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) *DescribeVscMountPointsResponseBodyMountPointsInstances {
	s.Vscs = v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstances) Validate() error {
	if s.Vscs != nil {
		for _, item := range s.Vscs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVscMountPointsResponseBodyMountPointsInstancesVscs struct {
	VscId     *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscName   *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscStatus *string `json:"VscStatus,omitempty" xml:"VscStatus,omitempty"`
	VscType   *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) GoString() string {
	return s.String()
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) GetVscName() *string {
	return s.VscName
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) GetVscStatus() *string {
	return s.VscStatus
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) GetVscType() *string {
	return s.VscType
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) SetVscId(v string) *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs {
	s.VscId = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) SetVscName(v string) *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs {
	s.VscName = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) SetVscStatus(v string) *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs {
	s.VscStatus = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) SetVscType(v string) *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs {
	s.VscType = &v
	return s
}

func (s *DescribeVscMountPointsResponseBodyMountPointsInstancesVscs) Validate() error {
	return dara.Validate(s)
}
