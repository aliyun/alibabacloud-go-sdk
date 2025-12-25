// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlussInstance interface {
	dara.Model
	String() string
	GoString() string
	SetClusterState(v *ClusterState) *FlussInstance
	GetClusterState() *ClusterState
	SetClusterStatus(v string) *FlussInstance
	GetClusterStatus() *string
	SetConsoleUrl(v string) *FlussInstance
	GetConsoleUrl() *string
	SetDiskSize(v int64) *FlussInstance
	GetDiskSize() *int64
	SetInstanceId(v string) *FlussInstance
	GetInstanceId() *string
	SetInstanceName(v string) *FlussInstance
	GetInstanceName() *string
	SetOrderState(v string) *FlussInstance
	GetOrderState() *string
	SetRegionId(v string) *FlussInstance
	GetRegionId() *string
	SetResourceCreateTime(v int64) *FlussInstance
	GetResourceCreateTime() *int64
	SetResourceExpiredTime(v int64) *FlussInstance
	GetResourceExpiredTime() *int64
	SetTabletServerModel(v string) *FlussInstance
	GetTabletServerModel() *string
	SetTabletServerNum(v int64) *FlussInstance
	GetTabletServerNum() *int64
	SetTabletServerType(v string) *FlussInstance
	GetTabletServerType() *string
	SetUid(v string) *FlussInstance
	GetUid() *string
	SetVSwitches(v []*FlussVswitch) *FlussInstance
	GetVSwitches() []*FlussVswitch
	SetVpcId(v string) *FlussInstance
	GetVpcId() *string
}

type FlussInstance struct {
	ClusterState        *ClusterState   `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	ClusterStatus       *string         `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	ConsoleUrl          *string         `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	DiskSize            *int64          `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	InstanceId          *string         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OrderState          *string         `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	RegionId            *string         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceCreateTime  *int64          `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceExpiredTime *int64          `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	TabletServerModel   *string         `json:"TabletServerModel,omitempty" xml:"TabletServerModel,omitempty"`
	TabletServerNum     *int64          `json:"TabletServerNum,omitempty" xml:"TabletServerNum,omitempty"`
	TabletServerType    *string         `json:"TabletServerType,omitempty" xml:"TabletServerType,omitempty"`
	Uid                 *string         `json:"Uid,omitempty" xml:"Uid,omitempty"`
	VSwitches           []*FlussVswitch `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId               *string         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s FlussInstance) String() string {
	return dara.Prettify(s)
}

func (s FlussInstance) GoString() string {
	return s.String()
}

func (s *FlussInstance) GetClusterState() *ClusterState {
	return s.ClusterState
}

func (s *FlussInstance) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *FlussInstance) GetConsoleUrl() *string {
	return s.ConsoleUrl
}

func (s *FlussInstance) GetDiskSize() *int64 {
	return s.DiskSize
}

func (s *FlussInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FlussInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *FlussInstance) GetOrderState() *string {
	return s.OrderState
}

func (s *FlussInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *FlussInstance) GetResourceCreateTime() *int64 {
	return s.ResourceCreateTime
}

func (s *FlussInstance) GetResourceExpiredTime() *int64 {
	return s.ResourceExpiredTime
}

func (s *FlussInstance) GetTabletServerModel() *string {
	return s.TabletServerModel
}

func (s *FlussInstance) GetTabletServerNum() *int64 {
	return s.TabletServerNum
}

func (s *FlussInstance) GetTabletServerType() *string {
	return s.TabletServerType
}

func (s *FlussInstance) GetUid() *string {
	return s.Uid
}

func (s *FlussInstance) GetVSwitches() []*FlussVswitch {
	return s.VSwitches
}

func (s *FlussInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *FlussInstance) SetClusterState(v *ClusterState) *FlussInstance {
	s.ClusterState = v
	return s
}

func (s *FlussInstance) SetClusterStatus(v string) *FlussInstance {
	s.ClusterStatus = &v
	return s
}

func (s *FlussInstance) SetConsoleUrl(v string) *FlussInstance {
	s.ConsoleUrl = &v
	return s
}

func (s *FlussInstance) SetDiskSize(v int64) *FlussInstance {
	s.DiskSize = &v
	return s
}

func (s *FlussInstance) SetInstanceId(v string) *FlussInstance {
	s.InstanceId = &v
	return s
}

func (s *FlussInstance) SetInstanceName(v string) *FlussInstance {
	s.InstanceName = &v
	return s
}

func (s *FlussInstance) SetOrderState(v string) *FlussInstance {
	s.OrderState = &v
	return s
}

func (s *FlussInstance) SetRegionId(v string) *FlussInstance {
	s.RegionId = &v
	return s
}

func (s *FlussInstance) SetResourceCreateTime(v int64) *FlussInstance {
	s.ResourceCreateTime = &v
	return s
}

func (s *FlussInstance) SetResourceExpiredTime(v int64) *FlussInstance {
	s.ResourceExpiredTime = &v
	return s
}

func (s *FlussInstance) SetTabletServerModel(v string) *FlussInstance {
	s.TabletServerModel = &v
	return s
}

func (s *FlussInstance) SetTabletServerNum(v int64) *FlussInstance {
	s.TabletServerNum = &v
	return s
}

func (s *FlussInstance) SetTabletServerType(v string) *FlussInstance {
	s.TabletServerType = &v
	return s
}

func (s *FlussInstance) SetUid(v string) *FlussInstance {
	s.Uid = &v
	return s
}

func (s *FlussInstance) SetVSwitches(v []*FlussVswitch) *FlussInstance {
	s.VSwitches = v
	return s
}

func (s *FlussInstance) SetVpcId(v string) *FlussInstance {
	s.VpcId = &v
	return s
}

func (s *FlussInstance) Validate() error {
	if s.ClusterState != nil {
		if err := s.ClusterState.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
