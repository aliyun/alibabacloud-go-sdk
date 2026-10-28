// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceInfos interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ResourceInfos
	GetClusterId() *string
	SetEcsSpec(v string) *ResourceInfos
	GetEcsSpec() *string
	SetGpuCardType(v string) *ResourceInfos
	GetGpuCardType() *string
	SetMachineModel(v string) *ResourceInfos
	GetMachineModel() *string
	SetMaxQuota(v int64) *ResourceInfos
	GetMaxQuota() *int64
	SetNetworkPodId(v string) *ResourceInfos
	GetNetworkPodId() *string
	SetRegionId(v string) *ResourceInfos
	GetRegionId() *string
	SetUsedQuota(v int64) *ResourceInfos
	GetUsedQuota() *int64
	SetUserId(v string) *ResourceInfos
	GetUserId() *string
	SetUserName(v string) *ResourceInfos
	GetUserName() *string
}

type ResourceInfos struct {
	// The cluster ID.
	//
	// example:
	//
	// c1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ECS instance type information.
	//
	// example:
	//
	// ecs.gn5-c8g1.2xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The GPU card type.
	GpuCardType *string `json:"GpuCardType,omitempty" xml:"GpuCardType,omitempty"`
	// The machine model.
	MachineModel *string `json:"MachineModel,omitempty" xml:"MachineModel,omitempty"`
	// The maximum quota.
	//
	// example:
	//
	// 100
	MaxQuota *int64 `json:"MaxQuota,omitempty" xml:"MaxQuota,omitempty"`
	// The network pod ID.
	NetworkPodId *string `json:"NetworkPodId,omitempty" xml:"NetworkPodId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 10
	UsedQuota *int64 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ResourceInfos) String() string {
	return dara.Prettify(s)
}

func (s ResourceInfos) GoString() string {
	return s.String()
}

func (s *ResourceInfos) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResourceInfos) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *ResourceInfos) GetGpuCardType() *string {
	return s.GpuCardType
}

func (s *ResourceInfos) GetMachineModel() *string {
	return s.MachineModel
}

func (s *ResourceInfos) GetMaxQuota() *int64 {
	return s.MaxQuota
}

func (s *ResourceInfos) GetNetworkPodId() *string {
	return s.NetworkPodId
}

func (s *ResourceInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ResourceInfos) GetUsedQuota() *int64 {
	return s.UsedQuota
}

func (s *ResourceInfos) GetUserId() *string {
	return s.UserId
}

func (s *ResourceInfos) GetUserName() *string {
	return s.UserName
}

func (s *ResourceInfos) SetClusterId(v string) *ResourceInfos {
	s.ClusterId = &v
	return s
}

func (s *ResourceInfos) SetEcsSpec(v string) *ResourceInfos {
	s.EcsSpec = &v
	return s
}

func (s *ResourceInfos) SetGpuCardType(v string) *ResourceInfos {
	s.GpuCardType = &v
	return s
}

func (s *ResourceInfos) SetMachineModel(v string) *ResourceInfos {
	s.MachineModel = &v
	return s
}

func (s *ResourceInfos) SetMaxQuota(v int64) *ResourceInfos {
	s.MaxQuota = &v
	return s
}

func (s *ResourceInfos) SetNetworkPodId(v string) *ResourceInfos {
	s.NetworkPodId = &v
	return s
}

func (s *ResourceInfos) SetRegionId(v string) *ResourceInfos {
	s.RegionId = &v
	return s
}

func (s *ResourceInfos) SetUsedQuota(v int64) *ResourceInfos {
	s.UsedQuota = &v
	return s
}

func (s *ResourceInfos) SetUserId(v string) *ResourceInfos {
	s.UserId = &v
	return s
}

func (s *ResourceInfos) SetUserName(v string) *ResourceInfos {
	s.UserName = &v
	return s
}

func (s *ResourceInfos) Validate() error {
	return dara.Validate(s)
}
