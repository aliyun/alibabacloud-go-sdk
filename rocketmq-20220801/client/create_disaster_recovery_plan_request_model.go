// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDisasterRecoveryPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSyncCheckpoint(v bool) *CreateDisasterRecoveryPlanRequest
	GetAutoSyncCheckpoint() *bool
	SetInstances(v []*CreateDisasterRecoveryPlanRequestInstances) *CreateDisasterRecoveryPlanRequest
	GetInstances() []*CreateDisasterRecoveryPlanRequestInstances
	SetPlanDesc(v string) *CreateDisasterRecoveryPlanRequest
	GetPlanDesc() *string
	SetPlanName(v string) *CreateDisasterRecoveryPlanRequest
	GetPlanName() *string
	SetPlanType(v string) *CreateDisasterRecoveryPlanRequest
	GetPlanType() *string
	SetSyncCheckpointEnabled(v bool) *CreateDisasterRecoveryPlanRequest
	GetSyncCheckpointEnabled() *bool
}

type CreateDisasterRecoveryPlanRequest struct {
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
	// The instances involved in the Global Replicator task. You must specify this parameter.
	Instances []*CreateDisasterRecoveryPlanRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The description of the Global Replicator task.
	//
	// example:
	//
	// xxx
	PlanDesc *string `json:"planDesc,omitempty" xml:"planDesc,omitempty"`
	// The name of the Global Replicator task. You must specify this parameter.
	//
	// example:
	//
	// xxx
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// The type of the Global Replicator task. You must specify this parameter. For more information, see [Global Replicator](https://help.aliyun.com/document_detail/2843187.html). Valid values:
	//
	// 	- ACTIVE_PASSIVE: one-way backup
	//
	// 	- ACTIVE_ACTIVE: two-way backup
	//
	// example:
	//
	// ACTIVE_ACTIVE
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

func (s CreateDisasterRecoveryPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDisasterRecoveryPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDisasterRecoveryPlanRequest) GetAutoSyncCheckpoint() *bool {
	return s.AutoSyncCheckpoint
}

func (s *CreateDisasterRecoveryPlanRequest) GetInstances() []*CreateDisasterRecoveryPlanRequestInstances {
	return s.Instances
}

func (s *CreateDisasterRecoveryPlanRequest) GetPlanDesc() *string {
	return s.PlanDesc
}

func (s *CreateDisasterRecoveryPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateDisasterRecoveryPlanRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *CreateDisasterRecoveryPlanRequest) GetSyncCheckpointEnabled() *bool {
	return s.SyncCheckpointEnabled
}

func (s *CreateDisasterRecoveryPlanRequest) SetAutoSyncCheckpoint(v bool) *CreateDisasterRecoveryPlanRequest {
	s.AutoSyncCheckpoint = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequest) SetInstances(v []*CreateDisasterRecoveryPlanRequestInstances) *CreateDisasterRecoveryPlanRequest {
	s.Instances = v
	return s
}

func (s *CreateDisasterRecoveryPlanRequest) SetPlanDesc(v string) *CreateDisasterRecoveryPlanRequest {
	s.PlanDesc = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequest) SetPlanName(v string) *CreateDisasterRecoveryPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequest) SetPlanType(v string) *CreateDisasterRecoveryPlanRequest {
	s.PlanType = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequest) SetSyncCheckpointEnabled(v bool) *CreateDisasterRecoveryPlanRequest {
	s.SyncCheckpointEnabled = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequest) Validate() error {
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

type CreateDisasterRecoveryPlanRequestInstances struct {
	// The authentication method. If you set instanceType to ALIYUN_ROCKETMQ and the instance is an ApsaraMQ for RocketMQ 4.0 instance, you do not need to specify this parameter.
	//
	// 	- NO_AUTH: no authentication
	//
	// 	- ACL_AUTH: access control list (ACL)-based authentication
	//
	// Valid values:
	//
	// 	- NO_AUTH: no authentication
	//
	// 	- ACL_AUTH: access control list (ACL)-based authentication
	//
	// example:
	//
	// ACL_AUTH
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The ID of the consumer group.
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
	// The instance ID. This parameter is required only if you set instanceType to ALIYUN_ROCKETMQ.
	//
	// example:
	//
	// rmq-cn-******
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance role. Valid values:
	//
	// 	- ACTIVE: primary instance
	//
	// 	- Passive: secondary instance
	//
	// example:
	//
	// PASSIVE
	InstanceRole *string `json:"instanceRole,omitempty" xml:"instanceRole,omitempty"`
	// The instance type. Valid values:
	//
	// 	- ALIYUN_ROCKETMQ: ApsaraMQ for RocketMQ instance
	//
	// 	- EXTERNAL_ROCKETMQ: external RocketMQ instance
	//
	// Valid values:
	//
	// 	- ALIYUN_ROCKETMQ: ApsaraMQ for RocketMQ instance
	//
	// 	- EXTERNAL_ROCKETMQ: external RocketMQ instance
	//
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The message attribute. When you synchronize a message to the destination cluster, the system automatically adds the attribute to the message for SQL-based filtering.
	MessageProperty *CreateDisasterRecoveryPlanRequestInstancesMessageProperty `json:"messageProperty,omitempty" xml:"messageProperty,omitempty" type:"Struct"`
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
	// The password used for authentication. This parameter is required only if you set authType to ACL_AUTH. If you set instanceType to ALIYUN_ROCKETMQ, you do not need to specify this parameter.
	//
	// example:
	//
	// xxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The region where the instance resides.
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

func (s CreateDisasterRecoveryPlanRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s CreateDisasterRecoveryPlanRequestInstances) GoString() string {
	return s.String()
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetInstanceRole() *string {
	return s.InstanceRole
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetMessageProperty() *CreateDisasterRecoveryPlanRequestInstancesMessageProperty {
	return s.MessageProperty
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetPassword() *string {
	return s.Password
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetUsername() *string {
	return s.Username
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDisasterRecoveryPlanRequestInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetAuthType(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.AuthType = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetConsumerGroupId(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.ConsumerGroupId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetEndpointUrl(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.EndpointUrl = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetInstanceId(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetInstanceRole(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.InstanceRole = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetInstanceType(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.InstanceType = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetMessageProperty(v *CreateDisasterRecoveryPlanRequestInstancesMessageProperty) *CreateDisasterRecoveryPlanRequestInstances {
	s.MessageProperty = v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetNetworkType(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.NetworkType = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetPassword(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.Password = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetRegionId(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.RegionId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetSecurityGroupId(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetUsername(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.Username = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetVSwitchId(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.VSwitchId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) SetVpcId(v string) *CreateDisasterRecoveryPlanRequestInstances {
	s.VpcId = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstances) Validate() error {
	if s.MessageProperty != nil {
		if err := s.MessageProperty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDisasterRecoveryPlanRequestInstancesMessageProperty struct {
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

func (s CreateDisasterRecoveryPlanRequestInstancesMessageProperty) String() string {
	return dara.Prettify(s)
}

func (s CreateDisasterRecoveryPlanRequestInstancesMessageProperty) GoString() string {
	return s.String()
}

func (s *CreateDisasterRecoveryPlanRequestInstancesMessageProperty) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *CreateDisasterRecoveryPlanRequestInstancesMessageProperty) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *CreateDisasterRecoveryPlanRequestInstancesMessageProperty) SetPropertyKey(v string) *CreateDisasterRecoveryPlanRequestInstancesMessageProperty {
	s.PropertyKey = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstancesMessageProperty) SetPropertyValue(v string) *CreateDisasterRecoveryPlanRequestInstancesMessageProperty {
	s.PropertyValue = &v
	return s
}

func (s *CreateDisasterRecoveryPlanRequestInstancesMessageProperty) Validate() error {
	return dara.Validate(s)
}
