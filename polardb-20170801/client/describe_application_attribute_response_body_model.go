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
	SetCanDisableSnat(v bool) *DescribeApplicationAttributeResponseBody
	GetCanDisableSnat() *bool
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
	SetIsLatestVersion(v bool) *DescribeApplicationAttributeResponseBody
	GetIsLatestVersion() *bool
	SetLatestVersion(v string) *DescribeApplicationAttributeResponseBody
	GetLatestVersion() *string
	SetLockMode(v string) *DescribeApplicationAttributeResponseBody
	GetLockMode() *string
	SetMaintainEndTime(v string) *DescribeApplicationAttributeResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *DescribeApplicationAttributeResponseBody
	GetMaintainStartTime() *string
	SetMemApplicationAttribute(v *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) *DescribeApplicationAttributeResponseBody
	GetMemApplicationAttribute() *DescribeApplicationAttributeResponseBodyMemApplicationAttribute
	SetMinorVersion(v string) *DescribeApplicationAttributeResponseBody
	GetMinorVersion() *string
	SetNatGatewayId(v string) *DescribeApplicationAttributeResponseBody
	GetNatGatewayId() *string
	SetPayType(v string) *DescribeApplicationAttributeResponseBody
	GetPayType() *string
	SetPolarClawSaaSApplicationAttribute(v *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) *DescribeApplicationAttributeResponseBody
	GetPolarClawSaaSApplicationAttribute() *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute
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
	SetSnatStatus(v string) *DescribeApplicationAttributeResponseBody
	GetSnatStatus() *string
	SetStatus(v string) *DescribeApplicationAttributeResponseBody
	GetStatus() *string
	SetStorages(v []*DescribeApplicationAttributeResponseBodyStorages) *DescribeApplicationAttributeResponseBody
	GetStorages() []*DescribeApplicationAttributeResponseBodyStorages
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
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The CPU architecture. The value is:
	//
	// - `x86`
	//
	// example:
	//
	// x86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Indicates whether SNAT can be disabled.
	CanDisableSnat *bool `json:"CanDisableSnat,omitempty" xml:"CanDisableSnat,omitempty"`
	// The child components.
	Components []*DescribeApplicationAttributeResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2025-03-25T09:37:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the PolarDB instance that the application depends on.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The application description.
	//
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application endpoints.
	Endpoints []*DescribeApplicationAttributeResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The expiration time.
	//
	// This parameter is not returned for Postpaid instances.
	//
	// example:
	//
	// 2025-06-25T09:37:10Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the application has expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Indicates whether the application is the latest version.
	//
	// example:
	//
	// true
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// v2026.3.13-1#20260320
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The lock mode. Valid values:
	//
	// - Unlock: The application is not locked.
	//
	// - Lock: The application is locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The maintenance end time.
	//
	// example:
	//
	// 19:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The maintenance start time.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The attributes of the Mem0 application.
	MemApplicationAttribute *DescribeApplicationAttributeResponseBodyMemApplicationAttribute `json:"MemApplicationAttribute,omitempty" xml:"MemApplicationAttribute,omitempty" type:"Struct"`
	// The minor version number.
	//
	// example:
	//
	// v2026.3.13-1#20260320
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The ID of the NAT Gateway.
	//
	// example:
	//
	// pc-xxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The billing method.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The attributes of the PolarClaw SaaS application.
	PolarClawSaaSApplicationAttribute *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute `json:"PolarClawSaaSApplicationAttribute,omitempty" xml:"PolarClawSaaSApplicationAttribute,omitempty" type:"Struct"`
	// The ID of the PolarFS Cold Storage or PolarFS High-performance instance.
	//
	// example:
	//
	// pfs-**************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The application-level security groups.
	SecurityGroups []*DescribeApplicationAttributeResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The application-level whitelists.
	SecurityIPArrays []*DescribeApplicationAttributeResponseBodySecurityIPArrays `json:"SecurityIPArrays,omitempty" xml:"SecurityIPArrays,omitempty" type:"Repeated"`
	// The Serverless type. Valid values:
	//
	// - 2: Agile
	//
	// - 3: Stable
	//
	// example:
	//
	// 2
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The SNAT status.
	//
	// example:
	//
	// off
	SnatStatus *string `json:"SnatStatus,omitempty" xml:"SnatStatus,omitempty"`
	// The application status. Valid values:
	//
	// - Creating: The application is being created.
	//
	// - Activated: The application is running.
	//
	// - Maintaining: The application is being maintained.
	//
	// - ClassChanging: The application configuration is being changed.
	//
	// - Transing: The application is being migrated.
	//
	// - MinorVersionUpgrading: The minor version of the application is being upgraded.
	//
	// - NetCreating: The endpoint is being created.
	//
	// - NetDeleting: The endpoint is being deleted.
	//
	// - NetModifying: The endpoint is being modified.
	//
	// - Restarting: The application is restarting.
	//
	// - Locking: The application is being locked.
	//
	// - Locked: The application is locked.
	//
	// - Unlocking: The application is being unlocked.
	//
	// - Deleting: The application is being deleted.
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the storage resources.
	Storages []*DescribeApplicationAttributeResponseBodyStorages `json:"Storages,omitempty" xml:"Storages,omitempty" type:"Repeated"`
	// Indicates whether an upgrade is available.
	//
	// example:
	//
	// false
	UpgradeAvailable *string `json:"UpgradeAvailable,omitempty" xml:"UpgradeAvailable,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The VSwitch ID.
	//
	// example:
	//
	// vsw-*******************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The application version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The zone ID.
	//
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

func (s *DescribeApplicationAttributeResponseBody) GetCanDisableSnat() *bool {
	return s.CanDisableSnat
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

func (s *DescribeApplicationAttributeResponseBody) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeApplicationAttributeResponseBody) GetLatestVersion() *string {
	return s.LatestVersion
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

func (s *DescribeApplicationAttributeResponseBody) GetMemApplicationAttribute() *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	return s.MemApplicationAttribute
}

func (s *DescribeApplicationAttributeResponseBody) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeApplicationAttributeResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeApplicationAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeApplicationAttributeResponseBody) GetPolarClawSaaSApplicationAttribute() *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute {
	return s.PolarClawSaaSApplicationAttribute
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

func (s *DescribeApplicationAttributeResponseBody) GetSnatStatus() *string {
	return s.SnatStatus
}

func (s *DescribeApplicationAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationAttributeResponseBody) GetStorages() []*DescribeApplicationAttributeResponseBodyStorages {
	return s.Storages
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

func (s *DescribeApplicationAttributeResponseBody) SetCanDisableSnat(v bool) *DescribeApplicationAttributeResponseBody {
	s.CanDisableSnat = &v
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

func (s *DescribeApplicationAttributeResponseBody) SetIsLatestVersion(v bool) *DescribeApplicationAttributeResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetLatestVersion(v string) *DescribeApplicationAttributeResponseBody {
	s.LatestVersion = &v
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

func (s *DescribeApplicationAttributeResponseBody) SetMemApplicationAttribute(v *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) *DescribeApplicationAttributeResponseBody {
	s.MemApplicationAttribute = v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetMinorVersion(v string) *DescribeApplicationAttributeResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetNatGatewayId(v string) *DescribeApplicationAttributeResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetPayType(v string) *DescribeApplicationAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetPolarClawSaaSApplicationAttribute(v *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) *DescribeApplicationAttributeResponseBody {
	s.PolarClawSaaSApplicationAttribute = v
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

func (s *DescribeApplicationAttributeResponseBody) SetSnatStatus(v string) *DescribeApplicationAttributeResponseBody {
	s.SnatStatus = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetStatus(v string) *DescribeApplicationAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBody) SetStorages(v []*DescribeApplicationAttributeResponseBodyStorages) *DescribeApplicationAttributeResponseBody {
	s.Storages = v
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
	if s.MemApplicationAttribute != nil {
		if err := s.MemApplicationAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.PolarClawSaaSApplicationAttribute != nil {
		if err := s.PolarClawSaaSApplicationAttribute.Validate(); err != nil {
			return err
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
	if s.Storages != nil {
		for _, item := range s.Storages {
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
	// The class of the child component.
	//
	// example:
	//
	// polar.app.g2.medium
	ComponentClass *string `json:"ComponentClass,omitempty" xml:"ComponentClass,omitempty"`
	// The description of the child component\\"s class.
	//
	// example:
	//
	// 2C4GB
	ComponentClassDescription *string `json:"ComponentClassDescription,omitempty" xml:"ComponentClassDescription,omitempty"`
	// The child component ID.
	//
	// example:
	//
	// pac-*******************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The maximum number of replicas for the child component.
	//
	// example:
	//
	// 1
	ComponentMaxReplica *int64 `json:"ComponentMaxReplica,omitempty" xml:"ComponentMaxReplica,omitempty"`
	// The number of replicas of the child component.
	//
	// example:
	//
	// 1
	ComponentReplica *int64 `json:"ComponentReplica,omitempty" xml:"ComponentReplica,omitempty"`
	// The group name of the child component replicas.
	//
	// example:
	//
	// default
	ComponentReplicaGroupName *string `json:"ComponentReplicaGroupName,omitempty" xml:"ComponentReplicaGroupName,omitempty"`
	// The type of the child component.
	//
	// example:
	//
	// gateway
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The component-level security groups.
	//
	// This parameter is not returned if the component-level security groups are the same as the application-level security groups.
	SecurityGroups []*DescribeApplicationAttributeResponseBodyComponentsSecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The component-level whitelists.
	//
	// This parameter is not returned if the component-level whitelists are the same as the application-level whitelists.
	SecurityIPArrays []*DescribeApplicationAttributeResponseBodyComponentsSecurityIPArrays `json:"SecurityIPArrays,omitempty" xml:"SecurityIPArrays,omitempty" type:"Repeated"`
	// The component status. The valid values are the same as those for the application status.
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topology of the child component.
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
	// The network type. The value is:
	//
	// - vpc
	//
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-*******************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group name.
	//
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
	// The name of the IP address group. The default value is `default`.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The tag of the IP address group.
	//
	// example:
	//
	// mytag
	SecurityIPArrayTag *string `json:"SecurityIPArrayTag,omitempty" xml:"SecurityIPArrayTag,omitempty"`
	// The whitelisted IP addresses, separated by commas.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The network type of the whitelisted IP addresses. The default value is `mix`.
	//
	// example:
	//
	// mix
	SecurityIPNetType *string `json:"SecurityIPNetType,omitempty" xml:"SecurityIPNetType,omitempty"`
	// The IP address type.
	//
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
	// The IDs or component types of the child nodes for this child component.
	Children []*string `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The topology layer of the child component.
	//
	// example:
	//
	// 0
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The IDs or component types of the parent nodes for this child component.
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
	// The endpoint description.
	//
	// example:
	//
	// myendpoint
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// pa-**************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.31.95.252
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The endpoint type. Valid values:
	//
	// - Private: a VPC endpoint
	//
	// - Public: a public endpoint
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The port description.
	//
	// example:
	//
	// kong_http
	PortDescription *string `json:"PortDescription,omitempty" xml:"PortDescription,omitempty"`
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

func (s *DescribeApplicationAttributeResponseBodyEndpoints) GetPortDescription() *string {
	return s.PortDescription
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

func (s *DescribeApplicationAttributeResponseBodyEndpoints) SetPortDescription(v string) *DescribeApplicationAttributeResponseBodyEndpoints {
	s.PortDescription = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodyMemApplicationAttribute struct {
	// The database name.
	//
	// example:
	//
	// test-database-name
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the embedder model.
	//
	// example:
	//
	// text-embedding-v4
	EmbedderModelName *string `json:"EmbedderModelName,omitempty" xml:"EmbedderModelName,omitempty"`
	// The name of the graph LLM model.
	//
	// example:
	//
	// qwen3-max
	GraphLlmModelName *string `json:"GraphLlmModelName,omitempty" xml:"GraphLlmModelName,omitempty"`
	// The name of the LLM model.
	//
	// example:
	//
	// qwen3-max
	LlmModelName *string `json:"LlmModelName,omitempty" xml:"LlmModelName,omitempty"`
	// The project name. It corresponds to the database schema where project data is stored.
	//
	// example:
	//
	// test-project-name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the reranker model.
	//
	// example:
	//
	// qwen3-rerank
	RerankerModelName *string `json:"RerankerModelName,omitempty" xml:"RerankerModelName,omitempty"`
	// The username.
	//
	// example:
	//
	// test-user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodyMemApplicationAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetDbName() *string {
	return s.DbName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetEmbedderModelName() *string {
	return s.EmbedderModelName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetGraphLlmModelName() *string {
	return s.GraphLlmModelName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetLlmModelName() *string {
	return s.LlmModelName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetRerankerModelName() *string {
	return s.RerankerModelName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) GetUserName() *string {
	return s.UserName
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetDbName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.DbName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetEmbedderModelName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.EmbedderModelName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetGraphLlmModelName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.GraphLlmModelName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetLlmModelName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.LlmModelName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetProjectName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.ProjectName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetRerankerModelName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.RerankerModelName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) SetUserName(v string) *DescribeApplicationAttributeResponseBodyMemApplicationAttribute {
	s.UserName = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyMemApplicationAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute struct {
	// The authentication callback URL.
	//
	// example:
	//
	// http://8.xxx.xxx.xxx.xxx/xxx
	AuthCallbackURL *string `json:"AuthCallbackURL,omitempty" xml:"AuthCallbackURL,omitempty"`
	// The enabled authentication providers.
	AuthProviders []*string `json:"AuthProviders,omitempty" xml:"AuthProviders,omitempty" type:"Repeated"`
	// The Supabase cluster ID.
	//
	// example:
	//
	// pa-**************
	SupabaseClusterId *string `json:"SupabaseClusterId,omitempty" xml:"SupabaseClusterId,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) GetAuthCallbackURL() *string {
	return s.AuthCallbackURL
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) GetAuthProviders() []*string {
	return s.AuthProviders
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) GetSupabaseClusterId() *string {
	return s.SupabaseClusterId
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) SetAuthCallbackURL(v string) *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute {
	s.AuthCallbackURL = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) SetAuthProviders(v []*string) *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute {
	s.AuthProviders = v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) SetSupabaseClusterId(v string) *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute {
	s.SupabaseClusterId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyPolarClawSaaSApplicationAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAttributeResponseBodySecurityGroups struct {
	// The network type. The value is:
	//
	// - vpc
	//
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group name.
	//
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
	// The name of the IP address group. The default value is `default`.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The tag of the IP address group.
	//
	// example:
	//
	// mytag
	SecurityIPArrayTag *string `json:"SecurityIPArrayTag,omitempty" xml:"SecurityIPArrayTag,omitempty"`
	// The whitelisted IP addresses, separated by commas.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The network type of the whitelisted IP addresses. The default value is `mix`.
	//
	// example:
	//
	// mix
	SecurityIPNetType *string `json:"SecurityIPNetType,omitempty" xml:"SecurityIPNetType,omitempty"`
	// The IP address type.
	//
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

type DescribeApplicationAttributeResponseBodyStorages struct {
	// The storage capacity.
	//
	// example:
	//
	// 20Gi
	StorageCapacity *string `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	// The storage instance ID.
	//
	// example:
	//
	// pa-**************
	StorageInstanceId *string `json:"StorageInstanceId,omitempty" xml:"StorageInstanceId,omitempty"`
	// The storage performance level.
	//
	// example:
	//
	// PL0
	StoragePerformanceLevel *string `json:"StoragePerformanceLevel,omitempty" xml:"StoragePerformanceLevel,omitempty"`
	// The storage type.
	//
	// example:
	//
	// essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeApplicationAttributeResponseBodyStorages) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponseBodyStorages) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponseBodyStorages) GetStorageCapacity() *string {
	return s.StorageCapacity
}

func (s *DescribeApplicationAttributeResponseBodyStorages) GetStorageInstanceId() *string {
	return s.StorageInstanceId
}

func (s *DescribeApplicationAttributeResponseBodyStorages) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *DescribeApplicationAttributeResponseBodyStorages) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeApplicationAttributeResponseBodyStorages) SetStorageCapacity(v string) *DescribeApplicationAttributeResponseBodyStorages {
	s.StorageCapacity = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyStorages) SetStorageInstanceId(v string) *DescribeApplicationAttributeResponseBodyStorages {
	s.StorageInstanceId = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyStorages) SetStoragePerformanceLevel(v string) *DescribeApplicationAttributeResponseBodyStorages {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyStorages) SetStorageType(v string) *DescribeApplicationAttributeResponseBodyStorages {
	s.StorageType = &v
	return s
}

func (s *DescribeApplicationAttributeResponseBodyStorages) Validate() error {
	return dara.Validate(s)
}
