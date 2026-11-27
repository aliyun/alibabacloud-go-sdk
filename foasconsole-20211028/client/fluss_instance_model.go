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
	SetHa(v bool) *FlussInstance
	GetHa() *bool
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
	SetTieringPostCu(v int64) *FlussInstance
	GetTieringPostCu() *int64
	SetTieringPreCu(v int64) *FlussInstance
	GetTieringPreCu() *int64
	SetUid(v string) *FlussInstance
	GetUid() *string
	SetVSwitches(v []*FlussVswitch) *FlussInstance
	GetVSwitches() []*FlussVswitch
	SetVpcId(v string) *FlussInstance
	GetVpcId() *string
}

type FlussInstance struct {
	// The running state information of the cluster.
	ClusterState *ClusterState `json:"ClusterState,omitempty" xml:"ClusterState,omitempty"`
	// The status of the Fluss cluster.
	//
	// example:
	//
	// RUNNING
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The console access URL of the Fluss instance.
	ConsoleUrl *string `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	// The disk size of a single TabletServer node, in GiB.
	//
	// example:
	//
	// 900
	DiskSize *int64 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// Specifies whether multi-zone high availability is enabled.
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// The ID of the Fluss instance.
	//
	// example:
	//
	// f-cn-zvp2q0z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Fluss instance.
	//
	// example:
	//
	// fluss-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The order status of the instance.
	//
	// example:
	//
	// NORMAL
	OrderState *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The creation time of the instance, represented as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1629879567394
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The time when the instance expires, represented as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1637337600000
	ResourceExpiredTime *int64 `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	// The TabletServer node specifications.
	//
	// example:
	//
	// 4 RCU
	TabletServerModel *string `json:"TabletServerModel,omitempty" xml:"TabletServerModel,omitempty"`
	// The number of TabletServer nodes.
	//
	// example:
	//
	// 3
	TabletServerNum *int64 `json:"TabletServerNum,omitempty" xml:"TabletServerNum,omitempty"`
	// The type of TabletServer node specifications.
	//
	// example:
	//
	// 4 RCU
	TabletServerType *string `json:"TabletServerType,omitempty" xml:"TabletServerType,omitempty"`
	// The upper limit of elastic compute resources for the lake-stream synchronization service, in CUs.
	//
	// example:
	//
	// 100
	TieringPostCu *int64 `json:"TieringPostCu,omitempty" xml:"TieringPostCu,omitempty"`
	// The base compute resource count for the lake-stream synchronization service, in CUs.
	TieringPreCu *int64 `json:"TieringPreCu,omitempty" xml:"TieringPreCu,omitempty"`
	// The UID of the Alibaba Cloud account to which the instance belongs.
	//
	// example:
	//
	// 183899668736****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The list of zone and vSwitch configurations used by the instance.
	VSwitches []*FlussVswitch `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *FlussInstance) GetHa() *bool {
	return s.Ha
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

func (s *FlussInstance) GetTieringPostCu() *int64 {
	return s.TieringPostCu
}

func (s *FlussInstance) GetTieringPreCu() *int64 {
	return s.TieringPreCu
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

func (s *FlussInstance) SetHa(v bool) *FlussInstance {
	s.Ha = &v
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

func (s *FlussInstance) SetTieringPostCu(v int64) *FlussInstance {
	s.TieringPostCu = &v
	return s
}

func (s *FlussInstance) SetTieringPreCu(v int64) *FlussInstance {
	s.TieringPreCu = &v
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
