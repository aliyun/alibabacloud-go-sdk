// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationAttributeResponseBody
	GetApplicationId() *string
	SetApplicationType(v string) *DescribeApplicationAttributeResponseBody
	GetApplicationType() *string
	SetArchitecture(v string) *DescribeApplicationAttributeResponseBody
	GetArchitecture() *string
	SetComponents(v []*DescribeApplicationAttributeResponseBodyComponents) *DescribeApplicationAttributeResponseBody
	GetComponents() []*DescribeApplicationAttributeResponseBodyComponents
	SetCreationTime(v string) *DescribeApplicationAttributeResponseBody
	GetCreationTime() *string
	SetDBClusterId(v string) *DescribeApplicationAttributeResponseBody
	GetDBClusterId() *string
	SetDescription(v string) *DescribeApplicationAttributeResponseBody
	GetDescription() *string
	SetEndpoints(v []*DescribeApplicationAttributeResponseBodyEndpoints) *DescribeApplicationAttributeResponseBody
	GetEndpoints() []*DescribeApplicationAttributeResponseBodyEndpoints
	SetExpireTime(v string) *DescribeApplicationAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v bool) *DescribeApplicationAttributeResponseBody
	GetExpired() *bool
	SetLockMode(v string) *DescribeApplicationAttributeResponseBody
	GetLockMode() *string
	SetMaintainEndTime(v string) *DescribeApplicationAttributeResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *DescribeApplicationAttributeResponseBody
	GetMaintainStartTime() *string
	SetPayType(v string) *DescribeApplicationAttributeResponseBody
	GetPayType() *string
	SetPolarFSInstanceId(v string) *DescribeApplicationAttributeResponseBody
	GetPolarFSInstanceId() *string
	SetRegionId(v string) *DescribeApplicationAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeApplicationAttributeResponseBody
	GetRequestId() *string
	SetSecurityGroups(v []*DescribeApplicationAttributeResponseBodySecurityGroups) *DescribeApplicationAttributeResponseBody
	GetSecurityGroups() []*DescribeApplicationAttributeResponseBodySecurityGroups
	SetSecurityIPArrays(v []*DescribeApplicationAttributeResponseBodySecurityIPArrays) *DescribeApplicationAttributeResponseBody
	GetSecurityIPArrays() []*DescribeApplicationAttributeResponseBodySecurityIPArrays
	SetServerlessType(v string) *DescribeApplicationAttributeResponseBody
	GetServerlessType() *string
	SetStatus(v string) *DescribeApplicationAttributeResponseBody
	GetStatus() *string
	SetUpgradeAvailable(v string) *DescribeApplicationAttributeResponseBody
	GetUpgradeAvailable() *string
	SetVPCId(v string) *DescribeApplicationAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribeApplicationAttributeResponseBody
	GetVSwitchId() *string
	SetVersion(v string) *DescribeApplicationAttributeResponseBody
	GetVersion() *string
	SetZoneId(v string) *DescribeApplicationAttributeResponseBody
	GetZoneId() *string
}

type DescribeApplicationAttributeResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// x86
	Architecture *string                                               `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	Components   []*DescribeApplicationAttributeResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-03-25T09:37:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// myapp
	Description *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoints   []*DescribeApplicationAttributeResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-06-25T09:37:10Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 19:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// pfs-**************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId        *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroups   []*DescribeApplicationAttributeResponseBodySecurityGroups   `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	SecurityIPArrays []*DescribeApplicationAttributeResponseBodySecurityIPArrays `json:"SecurityIPArrays,omitempty" xml:"SecurityIPArrays,omitempty" type:"Repeated"`
	ServerlessType   *string                                                     `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// false
	UpgradeAvailable *string `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// VSwitch ID
	//
	// example:
	//
	// vsw-*******************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// cn-beijing-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeApplicationAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationAttributeResponseBody) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeApplicationAttributeResponseBody) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeApplicationAttributeResponseBody) GetComponents() []*DescribeApplicationAttributeResponseBodyComponents {
	return s.Components
}

func (s *DescribeApplicationAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeApplicationAttributeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApplicationAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeApplicationAttributeResponseBody) GetEndpoints() []*DescribeApplicationAttributeResponseBodyEndpoints {
	return s.Endpoints
}

func (s *DescribeApplicationAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeApplicationAttributeResponseBody) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeApplicationAttributeResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeApplicationAttributeResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeApplicationAttributeResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeApplicationAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeApplicationAttributeResponseBody) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *DescribeApplicationAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationAttributeResponseBody) GetSecurityGroups() []*DescribeApplicationAttributeResponseBodySecurityGroups {
	return s.SecurityGroups
}

func (s *DescribeApplicationAttributeResponseBody) GetSecurityIPArrays() []*DescribeApplicationAttributeResponseBodySecurityIPArrays {
	return s.SecurityIPArrays
}

func (s *DescribeApplicationAttributeResponseBody) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeApplicationAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationAttributeResponseBody) GetUpgradeAvailable() *string {
	return s.UpgradeAvailable
}

func (s *DescribeApplicationAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeApplicationAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeApplicationAttributeResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeApplicationAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeApplicationAttributeResponseBody) SetApplicationId(v string) *DescribeApplicationAttributeResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetApplicationType(v string) *DescribeApplicationAttributeResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetArchitecture(v string) *DescribeApplicationAttributeResponseBody {
	s.Architecture = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetComponents(v []*DescribeApplicationAttributeResponseBodyComponents) *DescribeApplicationAttributeResponseBody {
	s.Components = v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetCreationTime(v string) *DescribeApplicationAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetDBClusterId(v string) *DescribeApplicationAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetDescription(v string) *DescribeApplicationAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetEndpoints(v []*DescribeApplicationAttributeResponseBodyEndpoints) *DescribeApplicationAttributeResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetExpireTime(v string) *DescribeApplicationAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetExpired(v bool) *DescribeApplicationAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetLockMode(v string) *DescribeApplicationAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetMaintainEndTime(v string) *DescribeApplicationAttributeResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetMaintainStartTime(v string) *DescribeApplicationAttributeResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetPayType(v string) *DescribeApplicationAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetPolarFSInstanceId(v string) *DescribeApplicationAttributeResponseBody {
	s.PolarFSInstanceId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetRegionId(v string) *DescribeApplicationAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetRequestId(v string) *DescribeApplicationAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetSecurityGroups(v []*DescribeApplicationAttributeResponseBodySecurityGroups) *DescribeApplicationAttributeResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetSecurityIPArrays(v []*DescribeApplicationAttributeResponseBodySecurityIPArrays) *DescribeApplicationAttributeResponseBody {
	s.SecurityIPArrays = v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetServerlessType(v string) *DescribeApplicationAttributeResponseBody {
	s.ServerlessType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetStatus(v string) *DescribeApplicationAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetUpgradeAvailable(v string) *DescribeApplicationAttributeResponseBody {
	s.UpgradeAvailable = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetVPCId(v string) *DescribeApplicationAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetVSwitchId(v string) *DescribeApplicationAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetVersion(v string) *DescribeApplicationAttributeResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetZoneId(v string) *DescribeApplicationAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityGroups != nil {
		for _, item := range s.SecurityGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityIPArrays != nil {
		for _, item := range s.SecurityIPArrays {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationAttributeResponseBodyComponents struct {
	// example:
	//
	// polar.app.g2.medium
	ComponentClass *string `json:"ComponentClass,omitempty" xml:"ComponentClass,omitempty"`
	// example:
	//
	// 2C4GB
	ComponentClassDescription *string `json:"ComponentClassDescription,omitempty" xml:"ComponentClassDescription,omitempty"`
	// example:
	//
	// pac-*******************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// example:
	//
	// 1
	ComponentMaxReplica *int64 `json:"ComponentMaxReplica,omitempty" xml:"ComponentMaxReplica,omitempty"`
	// example:
	//
	// 1
	ComponentReplica *int64 `json:"ComponentReplica,omitempty" xml:"ComponentReplica,omitempty"`
	// example:
	//
	// default
	ComponentReplicaGroupName *string `json:"ComponentReplicaGroupName,omitempty" xml:"ComponentReplicaGroupName,omitempty"`
	// example:
	//
	// gateway
	ComponentType    *string                                                               `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	SecurityGroups   []*DescribeApplicationAttributeResponseBodyComponentsSecurityGroups   `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	SecurityIPArrays []*DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays `json:"SecurityIPArrays,omitempty" xml:"SecurityIPArrays,omitempty" type:"Repeated"`
	// example:
	//
	// Activated
	Status   *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Topology *DescribeApplicationAttributeResponseBodyComponentsTopology `json:"Topology,omitempty" xml:"Topology,omitempty" type:"Struct"`
}

func (s DescribeApplicationAttributeResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentClass() *string {
	return s.ComponentClass
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentClassDescription() *string {
	return s.ComponentClassDescription
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentMaxReplica() *int64 {
	return s.ComponentMaxReplica
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentReplica() *int64 {
	return s.ComponentReplica
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentReplicaGroupName() *string {
	return s.ComponentReplicaGroupName
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetSecurityGroups() []*DescribeApplicationAttributeResponseBodyComponentsSecurityGroups {
	return s.SecurityGroups
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetSecurityIPArrays() []*DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays {
	return s.SecurityIPArrays
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationAttributeResponseBodyComponents) GetTopology() *DescribeApplicationAttributeResponseBodyComponentsTopology {
	return s.Topology
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentClass(v string) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentClass = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentClassDescription(v string) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentClassDescription = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentId(v string) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentMaxReplica(v int64) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentMaxReplica = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentReplica(v int64) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentReplica = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentReplicaGroupName(v string) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentReplicaGroupName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetComponentType(v string) *DescribeApplicationAttributeResponseBodyComponents {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetSecurityGroups(v []*DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) *DescribeApplicationAttributeResponseBodyComponents {
	s.SecurityGroups = v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetSecurityIPArrays(v []*DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) *DescribeApplicationAttributeResponseBodyComponents {
	s.SecurityIPArrays = v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetStatus(v string) *DescribeApplicationAttributeResponseBodyComponents {
	s.Status = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) SetTopology(v *DescribeApplicationAttributeResponseBodyComponentsTopology) *DescribeApplicationAttributeResponseBodyComponents {
	s.Topology = v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponents) Validate() error {
	if s.SecurityGroups != nil {
		for _, item := range s.SecurityGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityIPArrays != nil {
		for _, item := range s.SecurityIPArrays {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Topology != nil {
		if err := s.Topology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationAttributeResponseBodyComponentsSecurityGroups struct {
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-*******************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// MyGroupName
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) GetNetType() *string {
	return s.NetType
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) SetNetType(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups {
	s.NetType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) SetRegionId(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) SetSecurityGroupId(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) SetSecurityGroupName(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays struct {
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// example:
	//
	// mytag
	SecurityIPArrayTag *string `json:"SecurityIPArrayTag,omitempty" xml:"SecurityIPArrayTag,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// mix
	SecurityIPNetType *string `json:"SecurityIPNetType,omitempty" xml:"SecurityIPNetType,omitempty"`
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) GetSecurityIPArrayTag() *string {
	return s.SecurityIPArrayTag
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) GetSecurityIPNetType() *string {
	return s.SecurityIPNetType
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) SetSecurityIPArrayName(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays {
	s.SecurityIPArrayName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) SetSecurityIPArrayTag(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays {
	s.SecurityIPArrayTag = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) SetSecurityIPList(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) SetSecurityIPNetType(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays {
	s.SecurityIPNetType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) SetSecurityIPType(v string) *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays {
	s.SecurityIPType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodyComponentsTopology struct {
	Children []*string `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Layer   *string   `json:"Layer,omitempty" xml:"Layer,omitempty"`
	Parents []*string `json:"Parents,omitempty" xml:"Parents,omitempty" type:"Repeated"`
}

func (s DescribeApplicationAttributeResponseBodyComponentsTopology) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyComponentsTopology) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) GetChildren() []*string {
	return s.Children
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) GetLayer() *string {
	return s.Layer
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) GetParents() []*string {
	return s.Parents
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) SetChildren(v []*string) *DescribeApplicationAttributeResponseBodyComponentsTopology {
	s.Children = v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) SetLayer(v string) *DescribeApplicationAttributeResponseBodyComponentsTopology {
	s.Layer = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) SetParents(v []*string) *DescribeApplicationAttributeResponseBodyComponentsTopology {
	s.Parents = v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyComponentsTopology) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodyEndpoints struct {
	// example:
	//
	// myendpoint
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// pa-**************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// 172.31.95.252
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) GetDescription() *string {
	return s.Description
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) GetIP() *string {
	return s.IP
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) GetNetType() *string {
	return s.NetType
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) GetPort() *string {
	return s.Port
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) SetDescription(v string) *DescribeApplicationAttributeResponseBodyEndpoints {
	s.Description = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) SetEndpointId(v string) *DescribeApplicationAttributeResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) SetIP(v string) *DescribeApplicationAttributeResponseBodyEndpoints {
	s.IP = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) SetNetType(v string) *DescribeApplicationAttributeResponseBodyEndpoints {
	s.NetType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) SetPort(v string) *DescribeApplicationAttributeResponseBodyEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodySecurityGroups struct {
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// MyGroupName
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodySecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) GetNetType() *string {
	return s.NetType
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) SetNetType(v string) *DescribeApplicationAttributeResponseBodySecurityGroups {
	s.NetType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) SetRegionId(v string) *DescribeApplicationAttributeResponseBodySecurityGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) SetSecurityGroupId(v string) *DescribeApplicationAttributeResponseBodySecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) SetSecurityGroupName(v string) *DescribeApplicationAttributeResponseBodySecurityGroups {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodySecurityIPArrays struct {
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// example:
	//
	// mytag
	SecurityIPArrayTag *string `json:"SecurityIPArrayTag,omitempty" xml:"SecurityIPArrayTag,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// mix
	SecurityIPNetType *string `json:"SecurityIPNetType,omitempty" xml:"SecurityIPNetType,omitempty"`
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodySecurityIPArrays) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodySecurityIPArrays) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) GetSecurityIPArrayTag() *string {
	return s.SecurityIPArrayTag
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) GetSecurityIPNetType() *string {
	return s.SecurityIPNetType
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) SetSecurityIPArrayName(v string) *DescribeApplicationAttributeResponseBodySecurityIPArrays {
	s.SecurityIPArrayName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) SetSecurityIPArrayTag(v string) *DescribeApplicationAttributeResponseBodySecurityIPArrays {
	s.SecurityIPArrayTag = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) SetSecurityIPList(v string) *DescribeApplicationAttributeResponseBodySecurityIPArrays {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) SetSecurityIPNetType(v string) *DescribeApplicationAttributeResponseBodySecurityIPArrays {
	s.SecurityIPNetType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) SetSecurityIPType(v string) *DescribeApplicationAttributeResponseBodySecurityIPArrays {
	s.SecurityIPType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodySecurityIPArrays) Validate() error {
	return dara.Validate(s)
}
