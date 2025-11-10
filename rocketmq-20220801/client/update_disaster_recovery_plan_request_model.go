// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDisasterRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSyncCheckpoint(v bool) *UpdateDisasterRecoveryPlanRequest
	GetAutoSyncCheckpoint() *bool
	SetInstances(v []*UpdateDisasterRecoveryPlanRequestInstances) *UpdateDisasterRecoveryPlanRequest
	GetInstances() []*UpdateDisasterRecoveryPlanRequestInstances
	SetPlanDesc(v string) *UpdateDisasterRecoveryPlanRequest
	GetPlanDesc() *string
	SetPlanName(v string) *UpdateDisasterRecoveryPlanRequest
	GetPlanName() *string
	SetPlanType(v string) *UpdateDisasterRecoveryPlanRequest
	GetPlanType() *string
	SetSyncCheckpointEnabled(v bool) *UpdateDisasterRecoveryPlanRequest
	GetSyncCheckpointEnabled() *bool
}

type UpdateDisasterRecoveryPlanRequest struct {
	// Specifies whether to enable automatic consumer progress synchronization.
	//
	// >  This parameter takes effect only when you set `syncCheckpointEnabled` to true.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoSyncCheckpoint *bool `json:"autoSyncCheckpoint,omitempty" xml:"autoSyncCheckpoint,omitempty"`
	// The instances involved in the Global Replicator task. After you create a Global Replicator task, you cannot change the instances involved in the task. You can change only the message attribute and authentication type of the task.
	Instances []*UpdateDisasterRecoveryPlanRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The description of the Global Replicator task.
	//
	// example:
	//
	// xxx
	PlanDesc *string `json:"planDesc,omitempty" xml:"planDesc,omitempty"`
	// The name of the Global Replicator task.
	//
	// example:
	//
	// xxx
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// The type of the Global Replicator task. After you create a Global Replicator task, you cannot change the type of the task. Valid values:
	//
	// 	- ACTIVE_PASSIVE: one-way backup
	//
	// 	- ACTIVE_ACTIVE: two-way backup
	//
	// example:
	//
	// ACTIVE_PASSIVE
	PlanType *string `json:"planType,omitempty" xml:"planType,omitempty"`
	// Specifies whether to enable consumer progress synchronization.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SyncCheckpointEnabled *bool `json:"syncCheckpointEnabled,omitempty" xml:"syncCheckpointEnabled,omitempty"`
}

func (s UpdateDisasterRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryPlanRequest) GetAutoSyncCheckpoint() *bool {
	return s.AutoSyncCheckpoint
}

func (s *UpdateDisasterRecoveryPlanRequest) GetInstances() []*UpdateDisasterRecoveryPlanRequestInstances {
	return s.Instances
}

func (s *UpdateDisasterRecoveryPlanRequest) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *UpdateDisasterRecoveryPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *UpdateDisasterRecoveryPlanRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *UpdateDisasterRecoveryPlanRequest) GetSyncCheckpointEnabled() *bool {
	return s.SyncCheckpointEnabled
}

func (s *UpdateDisasterRecoveryPlanRequest) SetAutoSyncCheckpoint(v bool) *UpdateDisasterRecoveryPlanRequest {
	s.AutoSyncCheckpoint = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequest) SetInstances(v []*UpdateDisasterRecoveryPlanRequestInstances) *UpdateDisasterRecoveryPlanRequest {
	s.Instances = v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequest) SetPlanDesc(v string) *UpdateDisasterRecoveryPlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequest) SetPlanName(v string) *UpdateDisasterRecoveryPlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequest) SetPlanType(v string) *UpdateDisasterRecoveryPlanRequest {
	s.PlanType = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequest) SetSyncCheckpointEnabled(v bool) *UpdateDisasterRecoveryPlanRequest {
	s.SyncCheckpointEnabled = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequest) Validate() error {
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

type UpdateDisasterRecoveryPlanRequestInstances struct {
	// The authentication type. Valid values:
	//
	// 	- NO_AUTH: no authentication
	//
	// 	- ACL_AUTH: access control list (ACL)-based authentication
	//
	// <!---->
	//
	// *
	//
	// *
	//
	// example:
	//
	// NO_AUTH
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// GID_DS_XXX_YYY
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The instance endpoint. This parameter is required only if you set instanceType to EXTERNAL_ROCKETMQ.
	//
	// example:
	//
	// xxx
	EndpointUrl *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-83l3r0xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance role. Valid values:
	//
	// 	- ACTIVE: primary instance
	//
	// 	- Passive: secondary instance
	//
	// example:
	//
	// ACTIVE
	InstanceRole *string `json:"instanceRole,omitempty" xml:"instanceRole,omitempty"`
	// The instance type. Valid values:
	//
	// 	- ALIYUN_ROCKETMQ: ApsaraMQ for RocketMQ instance
	//
	// 	- EXTERNAL_ROCKETMQ: open source RocketMQ cluster
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The message attribute. When you synchronize a message to the destination cluster, the system automatically adds the attribute to the message for SQL-based filtering.
	MessageProperty *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty `json:"messageProperty,omitempty" xml:"messageProperty,omitempty" type:"Struct"`
	// The network type. This parameter is required only if you set instanceType to EXTERNAL_ROCKETMQ. Valid values:
	//
	// 	- TCP_INTERNET: Internet over TCP
	//
	// 	- TCP_VPC: virtual private cloud (VPC) over TCP.
	//
	// example:
	//
	// TCP_INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The password used for authentication. This parameter is required only if you set authType to ACL_AUTH.
	//
	// example:
	//
	// xxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the security group to which the instance belongs. This parameter is required only if you set instanceType to EXTERNAL_ROCKETMQ and networkType to TCP_VPC.
	//
	// example:
	//
	// sg-bp17hpmgz9******
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The username used for authentication. This parameter is required only if you set authType to ACL_AUTH.
	//
	// example:
	//
	// xxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
	// The ID of the vSwitch with which the instance is associated. This parameter is required only if you set instanceType to EXTERNAL_ROCKETMQ and networkType to TCP_VPC.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC with which the instance is associated. This parameter is required only if you set instanceType to EXTERNAL_ROCKETMQ and networkType to TCP_VPC.
	//
	// example:
	//
	// vpc-wz9qt50xhtj9krb******
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateDisasterRecoveryPlanRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryPlanRequestInstances) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetInstanceRole() *string {
	return s.InstanceRole
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetMessageProperty() *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty {
	return s.MessageProperty
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetPassword() *string {
	return s.Password
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetUsername() *string {
	return s.Username
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetAuthType(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.AuthType = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetConsumerGroupId(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.ConsumerGroupId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetEndpointUrl(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.EndpointUrl = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetInstanceId(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetInstanceRole(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.InstanceRole = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetInstanceType(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.InstanceType = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetMessageProperty(v *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) *UpdateDisasterRecoveryPlanRequestInstances {
	s.MessageProperty = v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetNetworkType(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.NetworkType = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetPassword(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.Password = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetRegionId(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.RegionId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetSecurityGroupId(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetUsername(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.Username = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetVSwitchId(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.VSwitchId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) SetVpcId(v string) *UpdateDisasterRecoveryPlanRequestInstances {
	s.VpcId = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstances) Validate() error {
	if s.MessageProperty != nil {
		if err := s.MessageProperty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDisasterRecoveryPlanRequestInstancesMessageProperty struct {
	// The attribute key.
	//
	// example:
	//
	// aaa
	PropertyKey *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// bbb
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) GoString() string {
	return s.String()
}

func (s *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) SetPropertyKey(v string) *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty {
	s.PropertyKey = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) SetPropertyValue(v string) *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty {
	s.PropertyValue = &v
	return s
}

func (s *UpdateDisasterRecoveryPlanRequestInstancesMessageProperty) Validate() error {
	return dara.Validate(s)
}
