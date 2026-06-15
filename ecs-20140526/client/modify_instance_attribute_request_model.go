// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuOptions(v *ModifyInstanceAttributeRequestCpuOptions) *ModifyInstanceAttributeRequest
	GetCpuOptions() *ModifyInstanceAttributeRequestCpuOptions
	SetCreditSpecification(v string) *ModifyInstanceAttributeRequest
	GetCreditSpecification() *string
	SetDeletionProtection(v bool) *ModifyInstanceAttributeRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *ModifyInstanceAttributeRequest
	GetDescription() *string
	SetEnableJumboFrame(v bool) *ModifyInstanceAttributeRequest
	GetEnableJumboFrame() *bool
	SetEnableNetworkEncryption(v bool) *ModifyInstanceAttributeRequest
	GetEnableNetworkEncryption() *bool
	SetHostName(v string) *ModifyInstanceAttributeRequest
	GetHostName() *string
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceAttributeRequest
	GetInstanceName() *string
	SetNetworkInterfaceQueueNumber(v int32) *ModifyInstanceAttributeRequest
	GetNetworkInterfaceQueueNumber() *int32
	SetOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifyInstanceAttributeRequest
	GetPassword() *string
	SetPrivateDnsNameOptions(v *ModifyInstanceAttributeRequestPrivateDnsNameOptions) *ModifyInstanceAttributeRequest
	GetPrivateDnsNameOptions() *ModifyInstanceAttributeRequestPrivateDnsNameOptions
	SetRecyclable(v bool) *ModifyInstanceAttributeRequest
	GetRecyclable() *bool
	SetRemoteConnectionOptions(v *ModifyInstanceAttributeRequestRemoteConnectionOptions) *ModifyInstanceAttributeRequest
	GetRemoteConnectionOptions() *ModifyInstanceAttributeRequestRemoteConnectionOptions
	SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupIds(v []*string) *ModifyInstanceAttributeRequest
	GetSecurityGroupIds() []*string
	SetUserData(v string) *ModifyInstanceAttributeRequest
	GetUserData() *string
}

type ModifyInstanceAttributeRequest struct {
	CpuOptions *ModifyInstanceAttributeRequestCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// The performance mode of the burstable performance instance. Valid values:
	//
	// - Standard: standard mode.
	//
	// - Unlimited: unlimited mode.
	//
	// For more information about the performance modes of burstable performance instances, see [Overview of burstable performance instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// Specifies whether to enable deletion protection for the instance. This setting prevents the instance from being released from the console or by calling the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation.
	//
	// > This feature applies only to pay-as-you-go instances and protects instances only from manual release operations. It does not affect system-initiated release operations.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The instance description. The description must be 2 to 256 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testInstanceDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable jumbo frames. When jumbo frames are enabled, the MTU of the instance is 8500. When jumbo frames are disabled, the MTU of the instance is 1500. Valid values:
	//
	// - true
	//
	// - false
	//
	// Note the following when you use this parameter:
	//
	// - The instance must be in the `Running` or `Stopped` state.
	//
	// - The instance must be in a VPC.
	//
	// - Only some instance types support jumbo frames. For more information, see [ECS instance MTU](https://help.aliyun.com/document_detail/200512.html).
	//
	// example:
	//
	// false
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// Specifies whether to enable VPC network traffic encryption. Valid values:
	//
	// - true
	//
	// - false
	//
	// > This parameter is in invitation-only preview and is not publicly available.
	//
	// example:
	//
	// true
	EnableNetworkEncryption *bool `json:"EnableNetworkEncryption,omitempty" xml:"EnableNetworkEncryption,omitempty"`
	// The hostname of the operating system. Note the following when you use this parameter:
	//
	// - The instance cannot be in the `Pending` or `Starting` state. Otherwise, the specified hostname and the `/etc/hosts` configuration may not take effect. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/25506.html) operation to query the current state of the instance.
	//
	// - The new hostname takes effect after you restart the instance from the ECS console (see [Restart an instance](https://help.aliyun.com/document_detail/25440.html)) or by calling the [RebootInstance](https://help.aliyun.com/document_detail/25502.html) operation. Restarting the instance from within its operating system does not apply the change.
	//
	// The hostname must meet the following requirements for different operating systems:
	//
	// - For Windows Server instances: The hostname must be 2 to 15 characters long and contain letters, digits, and hyphens (-). It cannot start or end with a hyphen, contain consecutive hyphens, or consist of only digits.
	//
	// - For other types of instances (such as Linux): The hostname must be 2 to 64 characters long. You can use periods (.) to separate the hostname into segments. Each segment can contain letters, digits, and hyphens (-). The hostname cannot start or end with a period or hyphen, and cannot contain consecutive periods or hyphens.
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name. The name must be 2 to 128 characters long. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of queues for the primary network interface. Note the following when you use this parameter:
	//
	// - The instance must be in the `Stopped` state.
	//
	// - The value cannot exceed the maximum number of queues that the instance type supports for a single network interface. The total number of queues across all network interfaces of the instance cannot exceed the total queue quota that the instance type supports. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the maximum number of queues per network interface and the total queue quota for an instance type.
	//
	// - If you set this parameter to -1, the number of queues for the primary network interface is reset to the default value for the instance type. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the default number of queues for an elastic network interface of an instance type.
	//
	// example:
	//
	// 8
	NetworkInterfaceQueueNumber *int32  `json:"NetworkInterfaceQueueNumber,omitempty" xml:"NetworkInterfaceQueueNumber,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters long and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The supported special characters are **()\\~!@#$%^&\\*-_+=|{}[]:;\\"<>,.?/**. For a Windows instance, the password cannot start with a forward slash (/). Note the following when you use this parameter:
	//
	// - The instance cannot be in the `Starting` state.
	//
	// - The new password takes effect after you restart the instance from the ECS console (see [Restart an instance](https://help.aliyun.com/document_detail/25440.html)) or by calling the [RebootInstance](https://help.aliyun.com/document_detail/25502.html) operation. Restarting the instance from within its operating system does not apply the change.
	//
	// > To prevent password exposure, send requests that include the `Password` parameter over HTTPS.
	//
	// example:
	//
	// Test123456&$
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The private DNS name settings for the instance.
	//
	// For more information about private DNS name resolution, see [ECS private DNS resolution
	//
	// ](https://help.aliyun.com/document_detail/2844797.html).
	PrivateDnsNameOptions *ModifyInstanceAttributeRequestPrivateDnsNameOptions `json:"PrivateDnsNameOptions,omitempty" xml:"PrivateDnsNameOptions,omitempty" type:"Struct"`
	// > This parameter is in invitation-only preview and is not publicly available.
	//
	// example:
	//
	// true
	Recyclable *bool `json:"Recyclable,omitempty" xml:"Recyclable,omitempty"`
	// > This parameter is in invitation-only preview and is not publicly available.
	RemoteConnectionOptions *ModifyInstanceAttributeRequestRemoteConnectionOptions `json:"RemoteConnectionOptions,omitempty" xml:"RemoteConnectionOptions,omitempty" type:"Struct"`
	ResourceOwnerAccount    *string                                                `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                                 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of security groups to assign to the instance. Note the following when you use this parameter:
	//
	// - The security group IDs in the array must be unique. The number of security groups that an instance can join is limited. For more information, see [Limits](~~25412#SecurityGroupQuota1~~).
	//
	// - Specifying this parameter removes the instance from its current security groups. To retain existing security group associations, you must include their IDs in this array.
	//
	// - You can switch the security group type. However, the specified security groups cannot include both basic and enterprise security groups.
	//
	// - The security groups must belong to the same VPC as the instance.
	//
	// - This parameter is not supported for instances in the classic network.
	//
	// > The change takes effect on the instance after a short delay.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7o****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The user data of the instance. User data should be Base64-encoded before it is passed. Note the following when you use this parameter:
	//
	// - The user data must comply with the limits described in [Generate user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// - After you restart the instance, the new user data is available on the instance but is not executed.
	//
	// > The raw data cannot exceed 32 KB before being encoded. Do not pass confidential information, such as passwords and private keys, in plaintext. If you must pass confidential information, encrypt and then Base64-encode it. On the instance, decrypt the information by using the corresponding decryption method.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetCpuOptions() *ModifyInstanceAttributeRequestCpuOptions {
	return s.CpuOptions
}

func (s *ModifyInstanceAttributeRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *ModifyInstanceAttributeRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyInstanceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyInstanceAttributeRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *ModifyInstanceAttributeRequest) GetEnableNetworkEncryption() *bool {
	return s.EnableNetworkEncryption
}

func (s *ModifyInstanceAttributeRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAttributeRequest) GetNetworkInterfaceQueueNumber() *int32 {
	return s.NetworkInterfaceQueueNumber
}

func (s *ModifyInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceAttributeRequest) GetPrivateDnsNameOptions() *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	return s.PrivateDnsNameOptions
}

func (s *ModifyInstanceAttributeRequest) GetRecyclable() *bool {
	return s.Recyclable
}

func (s *ModifyInstanceAttributeRequest) GetRemoteConnectionOptions() *ModifyInstanceAttributeRequestRemoteConnectionOptions {
	return s.RemoteConnectionOptions
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAttributeRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyInstanceAttributeRequest) GetUserData() *string {
	return s.UserData
}

func (s *ModifyInstanceAttributeRequest) SetCpuOptions(v *ModifyInstanceAttributeRequestCpuOptions) *ModifyInstanceAttributeRequest {
	s.CpuOptions = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetCreditSpecification(v string) *ModifyInstanceAttributeRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDeletionProtection(v bool) *ModifyInstanceAttributeRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetEnableJumboFrame(v bool) *ModifyInstanceAttributeRequest {
	s.EnableJumboFrame = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetEnableNetworkEncryption(v bool) *ModifyInstanceAttributeRequest {
	s.EnableNetworkEncryption = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetHostName(v string) *ModifyInstanceAttributeRequest {
	s.HostName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetNetworkInterfaceQueueNumber(v int32) *ModifyInstanceAttributeRequest {
	s.NetworkInterfaceQueueNumber = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPassword(v string) *ModifyInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPrivateDnsNameOptions(v *ModifyInstanceAttributeRequestPrivateDnsNameOptions) *ModifyInstanceAttributeRequest {
	s.PrivateDnsNameOptions = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRecyclable(v bool) *ModifyInstanceAttributeRequest {
	s.Recyclable = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRemoteConnectionOptions(v *ModifyInstanceAttributeRequestRemoteConnectionOptions) *ModifyInstanceAttributeRequest {
	s.RemoteConnectionOptions = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetSecurityGroupIds(v []*string) *ModifyInstanceAttributeRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetUserData(v string) *ModifyInstanceAttributeRequest {
	s.UserData = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateDnsNameOptions != nil {
		if err := s.PrivateDnsNameOptions.Validate(); err != nil {
			return err
		}
	}
	if s.RemoteConnectionOptions != nil {
		if err := s.RemoteConnectionOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceAttributeRequestCpuOptions struct {
	// The number of CPU cores. This parameter is not customizable and uses a default value.
	//
	// <props="china">
	//
	// For information about the default value, see [Custom CPU options](https://help.aliyun.com/document_detail/145895.html).
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The number of threads per core. The total number of vCPUs for an ECS instance is the value of `CpuOptions.Core` multiplied by the value of `CpuOptions.ThreadsPerCore`.
	//
	// - Setting `CpuOptions.ThreadsPerCore` to 1 disables hyper-threading.
	//
	// - Only some instance types support specifying the number of threads per core.
	//
	// <props="china">
	//
	// For information about the valid values and default value, see [Custom CPU options](https://help.aliyun.com/document_detail/145895.html).
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	// The CPU topology type of the instance. Valid values:
	//
	// - ContinuousCoreToHTMapping: The hyper-threads of the same core are contiguous.
	//
	// - DiscreteCoreToHTMapping: The hyper-threads of the same core are discrete.
	//
	// Default value: None.
	//
	// Note the following when you use this parameter:
	//
	// - The instance must be in the `Stopped` state.
	//
	// > This parameter is supported by only some instance families. See [View and modify the CPU topology](https://help.aliyun.com/document_detail/2636059.html) for a list of supported families.
	//
	// example:
	//
	// DiscreteCoreToHTMapping
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// > This parameter is in invitation-only preview and is not publicly available.
	//
	// example:
	//
	// enabled
	NestedVirtualization *string `json:"NestedVirtualization,omitempty" xml:"NestedVirtualization,omitempty"`
}

func (s ModifyInstanceAttributeRequestCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestCpuOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetTopologyType() *string {
	return s.TopologyType
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetNestedVirtualization() *string {
	return s.NestedVirtualization
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetCore(v int32) *ModifyInstanceAttributeRequestCpuOptions {
	s.Core = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetThreadsPerCore(v int32) *ModifyInstanceAttributeRequestCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetTopologyType(v string) *ModifyInstanceAttributeRequestCpuOptions {
	s.TopologyType = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetNestedVirtualization(v string) *ModifyInstanceAttributeRequestCpuOptions {
	s.NestedVirtualization = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceAttributeRequestPrivateDnsNameOptions struct {
	// Specifies whether to enable resolution of the instance ID-based domain name to an IPv6 address. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsAAAARecord *bool `json:"EnableInstanceIdDnsAAAARecord,omitempty" xml:"EnableInstanceIdDnsAAAARecord,omitempty"`
	// Specifies whether to enable resolution of the instance ID-based domain name to an IPv4 address. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsARecord *bool `json:"EnableInstanceIdDnsARecord,omitempty" xml:"EnableInstanceIdDnsARecord,omitempty"`
	// Specifies whether to enable resolution of the IP address-based domain name to an IPv4 address. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableIpDnsARecord *bool `json:"EnableIpDnsARecord,omitempty" xml:"EnableIpDnsARecord,omitempty"`
	// Specifies whether to enable reverse DNS resolution of an IPv4 address to an IP address-based domain name. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableIpDnsPtrRecord *bool `json:"EnableIpDnsPtrRecord,omitempty" xml:"EnableIpDnsPtrRecord,omitempty"`
	// The hostname type. Valid values:
	//
	// - Custom: a custom hostname.
	//
	// - IpBased: an IP address-based hostname.
	//
	// - InstanceIdBased: an instance ID-based hostname.
	//
	// Default value: Custom.
	//
	// example:
	//
	// Custom
	HostnameType *string `json:"HostnameType,omitempty" xml:"HostnameType,omitempty"`
}

func (s ModifyInstanceAttributeRequestPrivateDnsNameOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestPrivateDnsNameOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsAAAARecord() *bool {
	return s.EnableInstanceIdDnsAAAARecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsARecord() *bool {
	return s.EnableInstanceIdDnsARecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableIpDnsARecord() *bool {
	return s.EnableIpDnsARecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableIpDnsPtrRecord() *bool {
	return s.EnableIpDnsPtrRecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetHostnameType() *string {
	return s.HostnameType
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsAAAARecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsAAAARecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsARecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsARecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableIpDnsARecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableIpDnsARecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableIpDnsPtrRecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableIpDnsPtrRecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetHostnameType(v string) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.HostnameType = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceAttributeRequestRemoteConnectionOptions struct {
	// > This parameter is in invitation-only preview and is not publicly available.
	//
	// example:
	//
	// hide
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// > This parameter is in invitation-only preview and is not publicly available.
	//
	// example:
	//
	// hide
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyInstanceAttributeRequestRemoteConnectionOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestRemoteConnectionOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) GetType() *string {
	return s.Type
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) SetPassword(v string) *ModifyInstanceAttributeRequestRemoteConnectionOptions {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) SetType(v string) *ModifyInstanceAttributeRequestRemoteConnectionOptions {
	s.Type = &v
	return s
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) Validate() error {
	return dara.Validate(s)
}
