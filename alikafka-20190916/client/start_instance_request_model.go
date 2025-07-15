// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *StartInstanceRequest
	GetConfig() *string
	SetCrossZone(v bool) *StartInstanceRequest
	GetCrossZone() *bool
	SetDeployModule(v string) *StartInstanceRequest
	GetDeployModule() *string
	SetInstanceId(v string) *StartInstanceRequest
	GetInstanceId() *string
	SetIsEipInner(v bool) *StartInstanceRequest
	GetIsEipInner() *bool
	SetIsForceSelectedZones(v bool) *StartInstanceRequest
	GetIsForceSelectedZones() *bool
	SetIsSetUserAndPassword(v bool) *StartInstanceRequest
	GetIsSetUserAndPassword() *bool
	SetKMSKeyId(v string) *StartInstanceRequest
	GetKMSKeyId() *string
	SetName(v string) *StartInstanceRequest
	GetName() *string
	SetNotifier(v string) *StartInstanceRequest
	GetNotifier() *string
	SetPassword(v string) *StartInstanceRequest
	GetPassword() *string
	SetRegionId(v string) *StartInstanceRequest
	GetRegionId() *string
	SetSecurityGroup(v string) *StartInstanceRequest
	GetSecurityGroup() *string
	SetSelectedZones(v string) *StartInstanceRequest
	GetSelectedZones() *string
	SetServiceVersion(v string) *StartInstanceRequest
	GetServiceVersion() *string
	SetUserPhoneNum(v string) *StartInstanceRequest
	GetUserPhoneNum() *string
	SetUsername(v string) *StartInstanceRequest
	GetUsername() *string
	SetVSwitchId(v string) *StartInstanceRequest
	GetVSwitchId() *string
	SetVSwitchIds(v []*string) *StartInstanceRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *StartInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *StartInstanceRequest
	GetZoneId() *string
}

type StartInstanceRequest struct {
	// The initial configurations of the ApsaraMQ for Kafka instance. The values must be valid JSON strings. If you do not specify this parameter, it is left empty.
	//
	// > - You cannot configure this parameter when you deploy an ApsaraMQ for Confluent instance.
	//
	// > - You cannot configure enable.acl for instances whose versions are earlier than 2.2.0.
	//
	// The **Config*	- parameter supports the following parameters:
	//
	// 	- **enable.vpc_sasl_ssl**: specifies whether to enable VPC transmission encryption. Valid values:
	//
	//     	- **true**: enables VPC transmission encryption. If you enable VPC transmission encryption, you must also enable access control list (ACL).
	//
	//     	- **false**: disables VPC transmission encryption. This is the default value.
	//
	// 	- **enable.acl**: specifies whether to enable ACL. Valid values:
	//
	//     	- **true**: enables ACL.
	//
	//     	- **false**: disables the ACL feature. This is the default value.
	//
	// 	- **kafka.log.retention.hours**: the maximum message retention period when the disk capacity is sufficient. Unit: hours. Valid values: 24 to 480. Default value: **72**. When the disk usage reaches 85%, the disk capacity is insufficient. In this case, the system deletes the earliest stored messages to ensure service availability.
	//
	// 	- **kafka.message.max.bytes**: the maximum size of a message that can be sent and received by ApsaraMQ for Kafka. Unit: bytes. Valid values: 1048576 to 10485760. Default value: **1048576**. Before you change the maximum message size to a new value, make sure that the new value matches the configurations of the producers and consumers.
	//
	// example:
	//
	// {"kafka.log.retention.hours":"33"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether cross-zone deployment is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// false
	CrossZone *bool `json:"CrossZone,omitempty" xml:"CrossZone,omitempty"`
	// The deployment mode. If the instance is an ApsaraMQ for Kafka V2 instance, this parameter is required. If the instance is an ApsaraMQ for Kafka V3 instance or an ApsaraMQ for Confluent instance, this parameter is optional. Valid values:
	//
	// 	- **vpc**: deploys the instance in a virtual private cloud (VPC).
	//
	// 	- **eip**: deploys the instance over the Internet and in the VPC.
	//
	// The deployment mode of the ApsaraMQ for Kafka instance must be consistent with the instance type. If the instance is a VPC-connected instance, set this parameter to **vpc**. If the instance is an Internet- and VPC-connected instance, set this parameter to **eip**.
	//
	// example:
	//
	// vpc
	DeployModule *string `json:"DeployModule,omitempty" xml:"DeployModule,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the instance supports elastic IP addresses (EIPs). Valid values:
	//
	// 	- **true**: supports EIPs and allows access from the Internet and a VPC.
	//
	// 	- **false**: does not support EIPs and allows access only from a VPC.
	//
	// The value of this parameter must match the type of the instance. For example, if the instance allows access only from a VPC, set this parameter to **false**.
	//
	// example:
	//
	// false
	IsEipInner *bool `json:"IsEipInner,omitempty" xml:"IsEipInner,omitempty"`
	// Specifies whether to forcibly deploy the instance in the selected zones.
	//
	// example:
	//
	// false
	IsForceSelectedZones *bool `json:"IsForceSelectedZones,omitempty" xml:"IsForceSelectedZones,omitempty"`
	// Specifies whether to set a new username and password. Valid values:
	//
	// 	- **true**: sets a new username and password.
	//
	// 	- **false**: does not set a new username or password.
	//
	// This parameter is available only if you deploy an instance that allows access from the Internet and a VPC.
	//
	// example:
	//
	// false
	IsSetUserAndPassword *bool `json:"IsSetUserAndPassword,omitempty" xml:"IsSetUserAndPassword,omitempty"`
	// The ID of the key that is used for disk encryption in the region where the instance is deployed. You can obtain the ID of the key in the [Key Management Service (KMS) console](https://kms.console.aliyun.com/?spm=a2c4g.11186623.2.5.336745b8hfiU21) or create a key. For more information, see [Manage CMKs](https://help.aliyun.com/document_detail/181610.html).
	//
	// If this parameter is configured, disk encryption is enabled for the instance. You cannot disable disk encryption after disk encryption is enabled. When you call this operation, the system checks whether the AliyunServiceRoleForAlikafkaInstanceEncryption service-linked role is created. If the role is not created, the system automatically creates the role. For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/190460.html).
	//
	// > When you deploy a serverless ApsaraMQ for Kafka V3 instance, you cannot configure this parameter.
	//
	// example:
	//
	// 0d24xxxx-da7b-4786-b981-9a164dxxxxxx
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The name of the instance.
	//
	// >  If you specify a value for this parameter, make sure that the specified value is unique in the region of the instance.
	//
	// example:
	//
	// newInstanceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert contact.
	//
	// example:
	//
	// Mr. Wang
	Notifier *string `json:"Notifier,omitempty" xml:"Notifier,omitempty"`
	// The instance password.
	//
	// 	- This parameter is available only for Internet- and VPC- connected ApsaraMQ for Kafka V2 and V3 instances.
	//
	// 	- If the instance is an ApsaraMQ for Confluent instance, this parameter is required. The value of this parameter must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported: ! @ # $ % ^ & \\	- () _ + - =
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group of the instance.
	//
	// If you do not specify this parameter, ApsaraMQ for Kafka automatically configures a security group for your instance. If you specify this parameter, you must create a security group in advance. For more information, see [Create a security group](https://help.aliyun.com/document_detail/25468.html).
	//
	// example:
	//
	// sg-bp13wfx7kz9pko****
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The two-dimensional arrays that consist of the candidate set for primary zones and the candidate set for secondary zones. Custom code in the `zone {zone}` format and standard code in the `cn-RegionID-{zone}` format are supported.
	//
	// 	- If you set CrossZone to true and specify Zone H and Zone F as the candidate set for primary zones and Zone K as the candidate set for secondary zones, set this parameter to `[["zoneh","zonef"],["zonek"]]`.
	//
	// > If you specify multiple zones as the primary or secondary zones, the system deploys the instance in one of the zones without prioritizing them. For example, if you set this parameter to `[["zoneh","zonef"],["zonek"]]`, the primary zone in which the instance is deployed can be Zone H or Zone F, and the secondary zone is Zone K.
	//
	// 	- If you set CrossZone to false and want to deploy the instance in Zone K, set this parameter to `[["zonek"],[]]`. In this case, the value of this parameter must still be two-dimensional arrays, but the array that specifies the candidate for secondary zones is left empty.
	//
	// example:
	//
	// [[\\"zonel\\"],[\\"zonek\\"]]
	SelectedZones *string `json:"SelectedZones,omitempty" xml:"SelectedZones,omitempty"`
	// The version of the ApsaraMQ for Kafka instance. Valid values:
	//
	// 	- ApsaraMQ for Kafka V2 instances: 2.2.0 and 2.6.2.
	//
	// 	- ApsaraMQ for Kafka V3 instances: 3.3.1.
	//
	// 	- ApsaraMQ for Confluent instances: 7.4.0.
	//
	// Default value:
	//
	// 	- ApsaraMQ for Kafka V2 instances: 2.2.0.
	//
	// 	- ApsaraMQ for Kafka V3 instances: 3.3.1.
	//
	// 	- ApsaraMQ for Confluent instances: 7.4.0.
	//
	// example:
	//
	// ApsaraMQ for Kafka V2 instances: 2.2.0
	//
	// ApsaraMQ for Kafka V3 instances: 3.3.1
	//
	// ApsaraMQ for Confluent instances: 7.4.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The mobile phone number of the alert contact.
	//
	// example:
	//
	// 1581234****
	UserPhoneNum *string `json:"UserPhoneNum,omitempty" xml:"UserPhoneNum,omitempty"`
	// The instance username.
	//
	// 	- This parameter is available only for Internet- and VPC- connected ApsaraMQ for Kafka V2 and V3 instances.
	//
	// 	- If the instance is an ApsaraMQ for Confluent instance, set this parameter to root or leave this parameter empty.
	//
	// Default value for ApsaraMQ for Kafka V2 and V3 instances: username. Default value for ApsaraMQ for Confluent instances: root.
	//
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The ID of the vSwitch to which you want to connect the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1j3sg5979fstnpl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of the vSwitches with which the instance is associated. If the instance is an ApsaraMQ for Kafka V2 or V3 instance, this parameter is required. If the instance is an ApsaraMQ for Confluent instance, you must configure one of VSwitchIds and VSwitchId. If you configure both of the parameters, the value of VSwitchIds takes effect.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which you want to deploy the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1r4eg3yrxmygv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone where you want to deploy the ApsaraMQ for Kafka instance.
	//
	// 	- The zone ID of the ApsaraMQ for Kafka instance must be the same as that of the vSwitch.
	//
	// 	- The value must be in the zoneX or Region ID-X format. Examples: zonea and cn-hangzhou-k.
	//
	// >  If resources in the specified zone is insufficient, the instance may be deployed in another zone.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetConfig() *string {
	return s.Config
}

func (s *StartInstanceRequest) GetCrossZone() *bool {
	return s.CrossZone
}

func (s *StartInstanceRequest) GetDeployModule() *string {
	return s.DeployModule
}

func (s *StartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceRequest) GetIsEipInner() *bool {
	return s.IsEipInner
}

func (s *StartInstanceRequest) GetIsForceSelectedZones() *bool {
	return s.IsForceSelectedZones
}

func (s *StartInstanceRequest) GetIsSetUserAndPassword() *bool {
	return s.IsSetUserAndPassword
}

func (s *StartInstanceRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *StartInstanceRequest) GetName() *string {
	return s.Name
}

func (s *StartInstanceRequest) GetNotifier() *string {
	return s.Notifier
}

func (s *StartInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *StartInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstanceRequest) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *StartInstanceRequest) GetSelectedZones() *string {
	return s.SelectedZones
}

func (s *StartInstanceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *StartInstanceRequest) GetUserPhoneNum() *string {
	return s.UserPhoneNum
}

func (s *StartInstanceRequest) GetUsername() *string {
	return s.Username
}

func (s *StartInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *StartInstanceRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *StartInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *StartInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *StartInstanceRequest) SetConfig(v string) *StartInstanceRequest {
	s.Config = &v
	return s
}

func (s *StartInstanceRequest) SetCrossZone(v bool) *StartInstanceRequest {
	s.CrossZone = &v
	return s
}

func (s *StartInstanceRequest) SetDeployModule(v string) *StartInstanceRequest {
	s.DeployModule = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetIsEipInner(v bool) *StartInstanceRequest {
	s.IsEipInner = &v
	return s
}

func (s *StartInstanceRequest) SetIsForceSelectedZones(v bool) *StartInstanceRequest {
	s.IsForceSelectedZones = &v
	return s
}

func (s *StartInstanceRequest) SetIsSetUserAndPassword(v bool) *StartInstanceRequest {
	s.IsSetUserAndPassword = &v
	return s
}

func (s *StartInstanceRequest) SetKMSKeyId(v string) *StartInstanceRequest {
	s.KMSKeyId = &v
	return s
}

func (s *StartInstanceRequest) SetName(v string) *StartInstanceRequest {
	s.Name = &v
	return s
}

func (s *StartInstanceRequest) SetNotifier(v string) *StartInstanceRequest {
	s.Notifier = &v
	return s
}

func (s *StartInstanceRequest) SetPassword(v string) *StartInstanceRequest {
	s.Password = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) SetSecurityGroup(v string) *StartInstanceRequest {
	s.SecurityGroup = &v
	return s
}

func (s *StartInstanceRequest) SetSelectedZones(v string) *StartInstanceRequest {
	s.SelectedZones = &v
	return s
}

func (s *StartInstanceRequest) SetServiceVersion(v string) *StartInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *StartInstanceRequest) SetUserPhoneNum(v string) *StartInstanceRequest {
	s.UserPhoneNum = &v
	return s
}

func (s *StartInstanceRequest) SetUsername(v string) *StartInstanceRequest {
	s.Username = &v
	return s
}

func (s *StartInstanceRequest) SetVSwitchId(v string) *StartInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *StartInstanceRequest) SetVSwitchIds(v []*string) *StartInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *StartInstanceRequest) SetVpcId(v string) *StartInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *StartInstanceRequest) SetZoneId(v string) *StartInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
