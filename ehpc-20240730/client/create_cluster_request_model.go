// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPackages(v []*CreateClusterRequestAdditionalPackages) *CreateClusterRequest
	GetAdditionalPackages() []*CreateClusterRequestAdditionalPackages
	SetAddons(v []*CreateClusterRequestAddons) *CreateClusterRequest
	GetAddons() []*CreateClusterRequestAddons
	SetClientVersion(v string) *CreateClusterRequest
	GetClientVersion() *string
	SetClusterCategory(v string) *CreateClusterRequest
	GetClusterCategory() *string
	SetClusterCredentials(v *CreateClusterRequestClusterCredentials) *CreateClusterRequest
	GetClusterCredentials() *CreateClusterRequestClusterCredentials
	SetClusterCustomConfiguration(v *CreateClusterRequestClusterCustomConfiguration) *CreateClusterRequest
	GetClusterCustomConfiguration() *CreateClusterRequestClusterCustomConfiguration
	SetClusterDescription(v string) *CreateClusterRequest
	GetClusterDescription() *string
	SetClusterMode(v string) *CreateClusterRequest
	GetClusterMode() *string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetClusterVSwitchId(v string) *CreateClusterRequest
	GetClusterVSwitchId() *string
	SetClusterVpcId(v string) *CreateClusterRequest
	GetClusterVpcId() *string
	SetDeletionProtection(v bool) *CreateClusterRequest
	GetDeletionProtection() *bool
	SetIsEnterpriseSecurityGroup(v bool) *CreateClusterRequest
	GetIsEnterpriseSecurityGroup() *bool
	SetManager(v *CreateClusterRequestManager) *CreateClusterRequest
	GetManager() *CreateClusterRequestManager
	SetMaxCoreCount(v int32) *CreateClusterRequest
	GetMaxCoreCount() *int32
	SetMaxCount(v int32) *CreateClusterRequest
	GetMaxCount() *int32
	SetQueues(v []*QueueTemplate) *CreateClusterRequest
	GetQueues() []*QueueTemplate
	SetResourceGroupId(v string) *CreateClusterRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateClusterRequest
	GetSecurityGroupId() *string
	SetSharedStorages(v []*SharedStorageTemplate) *CreateClusterRequest
	GetSharedStorages() []*SharedStorageTemplate
	SetTags(v []*CreateClusterRequestTags) *CreateClusterRequest
	GetTags() []*CreateClusterRequestTags
}

type CreateClusterRequest struct {
	// The list of software that you want to install in the cluster. Valid values of N: 0 to 10.
	AdditionalPackages []*CreateClusterRequestAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The configurations of the custom addons in the cluster. Only one addon is supported.
	Addons []*CreateClusterRequestAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The client version. By default, the latest version is used.
	//
	// example:
	//
	// 2.1.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The access credentials of the cluster.
	ClusterCredentials *CreateClusterRequestClusterCredentials `json:"ClusterCredentials,omitempty" xml:"ClusterCredentials,omitempty" type:"Struct"`
	// The post-processing script of the cluster.
	ClusterCustomConfiguration *CreateClusterRequestClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The cluster description. The description must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The deployment mode of the cluster. Valid values:
	//
	// 	- Integrated
	//
	// 	- Hybrid
	//
	// 	- Custom
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The cluster name. The name must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// slurm22.05.8-cluster-20240718
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the vSwitch that you want the cluster to use. The vSwitch must reside in the VPC that is specified by the `ClusterVpcId` parameter.
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html) operation to query information about the created VPCs and vSwitches.
	//
	// example:
	//
	// vsw-f8za5p0mwzgdu3wgx****
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the cluster resides.
	//
	// example:
	//
	// vpc-m5efjevmclc0xdmys****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Specifies whether to enable deletion protection for the cluster. Deletion protection decides whether the cluster can be deleted in the console or by calling the [DeleteCluster](https://help.aliyun.com/document_detail/424406.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to use an advanced security group. Valid values:
	//
	// 	- true: automatically creates and uses an advanced security group.
	//
	// 	- false: automatically creates and uses a basic security group.
	//
	// For more information, see [Basic security groups and advanced security groups](https://help.aliyun.com/document_detail/605897.html).
	//
	// example:
	//
	// false
	IsEnterpriseSecurityGroup *bool `json:"IsEnterpriseSecurityGroup,omitempty" xml:"IsEnterpriseSecurityGroup,omitempty"`
	// The configurations of the cluster management node.
	Manager *CreateClusterRequestManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The maximum number of vCPUs that can be used by compute nodes in the cluster. Valid values: 0 to 100,000.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that the cluster can manage. Valid values: 0 to 5,000.
	//
	// example:
	//
	// 500
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The queues in the cluster. The number of queues can be 0 to 8.
	Queues []*QueueTemplate `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The ID of the resource group to which the cluster belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to obtain the IDs of the resource groups.
	//
	// example:
	//
	// rg-acfmxazb4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group to which the cluster belongs.
	//
	// You can call the [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to query available security groups in the current region.
	//
	// example:
	//
	// sg-bp13n61xsydodfyg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The shared storage resources of the cluster.
	SharedStorages []*SharedStorageTemplate `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
	// The tags of the cluster.
	Tags []*CreateClusterRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetAdditionalPackages() []*CreateClusterRequestAdditionalPackages {
	return s.AdditionalPackages
}

func (s *CreateClusterRequest) GetAddons() []*CreateClusterRequestAddons {
	return s.Addons
}

func (s *CreateClusterRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *CreateClusterRequest) GetClusterCategory() *string {
	return s.ClusterCategory
}

func (s *CreateClusterRequest) GetClusterCredentials() *CreateClusterRequestClusterCredentials {
	return s.ClusterCredentials
}

func (s *CreateClusterRequest) GetClusterCustomConfiguration() *CreateClusterRequestClusterCustomConfiguration {
	return s.ClusterCustomConfiguration
}

func (s *CreateClusterRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *CreateClusterRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetClusterVSwitchId() *string {
	return s.ClusterVSwitchId
}

func (s *CreateClusterRequest) GetClusterVpcId() *string {
	return s.ClusterVpcId
}

func (s *CreateClusterRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateClusterRequest) GetIsEnterpriseSecurityGroup() *bool {
	return s.IsEnterpriseSecurityGroup
}

func (s *CreateClusterRequest) GetManager() *CreateClusterRequestManager {
	return s.Manager
}

func (s *CreateClusterRequest) GetMaxCoreCount() *int32 {
	return s.MaxCoreCount
}

func (s *CreateClusterRequest) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *CreateClusterRequest) GetQueues() []*QueueTemplate {
	return s.Queues
}

func (s *CreateClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateClusterRequest) GetSharedStorages() []*SharedStorageTemplate {
	return s.SharedStorages
}

func (s *CreateClusterRequest) GetTags() []*CreateClusterRequestTags {
	return s.Tags
}

func (s *CreateClusterRequest) SetAdditionalPackages(v []*CreateClusterRequestAdditionalPackages) *CreateClusterRequest {
	s.AdditionalPackages = v
	return s
}

func (s *CreateClusterRequest) SetAddons(v []*CreateClusterRequestAddons) *CreateClusterRequest {
	s.Addons = v
	return s
}

func (s *CreateClusterRequest) SetClientVersion(v string) *CreateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterRequest) SetClusterCategory(v string) *CreateClusterRequest {
	s.ClusterCategory = &v
	return s
}

func (s *CreateClusterRequest) SetClusterCredentials(v *CreateClusterRequestClusterCredentials) *CreateClusterRequest {
	s.ClusterCredentials = v
	return s
}

func (s *CreateClusterRequest) SetClusterCustomConfiguration(v *CreateClusterRequestClusterCustomConfiguration) *CreateClusterRequest {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *CreateClusterRequest) SetClusterDescription(v string) *CreateClusterRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterRequest) SetClusterMode(v string) *CreateClusterRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVSwitchId(v string) *CreateClusterRequest {
	s.ClusterVSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVpcId(v string) *CreateClusterRequest {
	s.ClusterVpcId = &v
	return s
}

func (s *CreateClusterRequest) SetDeletionProtection(v bool) *CreateClusterRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateClusterRequest) SetIsEnterpriseSecurityGroup(v bool) *CreateClusterRequest {
	s.IsEnterpriseSecurityGroup = &v
	return s
}

func (s *CreateClusterRequest) SetManager(v *CreateClusterRequestManager) *CreateClusterRequest {
	s.Manager = v
	return s
}

func (s *CreateClusterRequest) SetMaxCoreCount(v int32) *CreateClusterRequest {
	s.MaxCoreCount = &v
	return s
}

func (s *CreateClusterRequest) SetMaxCount(v int32) *CreateClusterRequest {
	s.MaxCount = &v
	return s
}

func (s *CreateClusterRequest) SetQueues(v []*QueueTemplate) *CreateClusterRequest {
	s.Queues = v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSharedStorages(v []*SharedStorageTemplate) *CreateClusterRequest {
	s.SharedStorages = v
	return s
}

func (s *CreateClusterRequest) SetTags(v []*CreateClusterRequestTags) *CreateClusterRequest {
	s.Tags = v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	if s.AdditionalPackages != nil {
		for _, item := range s.AdditionalPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClusterCredentials != nil {
		if err := s.ClusterCredentials.Validate(); err != nil {
			return err
		}
	}
	if s.ClusterCustomConfiguration != nil {
		if err := s.ClusterCustomConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Manager != nil {
		if err := s.Manager.Validate(); err != nil {
			return err
		}
	}
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SharedStorages != nil {
		for _, item := range s.SharedStorages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterRequestAdditionalPackages struct {
	// The name of the software that you want to install in the cluster.
	//
	// example:
	//
	// mpich
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software that you want to install in the cluster.
	//
	// example:
	//
	// 4.0.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestAdditionalPackages) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestAdditionalPackages) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalPackages) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequestAdditionalPackages) GetVersion() *string {
	return s.Version
}

func (s *CreateClusterRequestAdditionalPackages) SetName(v string) *CreateClusterRequestAdditionalPackages {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestAdditionalPackages) SetVersion(v string) *CreateClusterRequestAdditionalPackages {
	s.Version = &v
	return s
}

func (s *CreateClusterRequestAdditionalPackages) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestAddons struct {
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	//
	// example:
	//
	// "{\\\\"EipResource\\\\": {\\\\"AutoCreate\\\\": true}, \\\\"EcsResources\\\\": [{\\\\"InstanceType\\\\": \\\\"ecs.c7.xlarge\\\\", \\\\"ImageId\\\\": \\\\"centos_7_6_x64_20G_alibase_20211130.vhd\\\\", \\\\"SystemDisk\\\\": {\\\\"Category\\\\": \\\\"cloud_essd\\\\", \\\\"Size\\\\": 40, \\\\"Level\\\\": \\\\"PL0\\\\"}, \\\\"EnableHT\\\\": true, \\\\"InstanceChargeType\\\\": \\\\"PostPaid\\\\", \\\\"SpotStrategy\\\\": \\\\"NoSpot\\\\"}]}"
	ResourcesSpec *string `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty"`
	// The service configurations of the addon.
	//
	// example:
	//
	// "[{\\\\"ServiceName\\\\": \\\\"SSH\\\\", \\\\"ServiceAccessType\\\\": null, \\\\"ServiceAccessUrl\\\\": null, \\\\"NetworkACL\\\\": [{\\\\"IpProtocol\\\\": \\\\"TCP\\\\", \\\\"Port\\\\": 22, \\\\"SourceCidrIp\\\\": \\\\"0.0.0.0/0\\\\"}]}, {\\\\"ServiceName\\\\": \\\\"VNC\\\\", \\\\"ServiceAccessType\\\\": null, \\\\"ServiceAccessUrl\\\\": null, \\\\"NetworkACL\\\\": [{\\\\"IpProtocol\\\\": \\\\"TCP\\\\", \\\\"Port\\\\": 12016, \\\\"SourceCidrIp\\\\": \\\\"0.0.0.0/0\\\\"}]}, {\\\\"ServiceName\\\\": \\\\"CLIENT\\\\", \\\\"ServiceAccessType\\\\": \\\\"URL\\\\", \\\\"ServiceAccessUrl\\\\": \\\\"\\\\", \\\\"NetworkACL\\\\": [{\\\\"IpProtocol\\\\": \\\\"TCP\\\\", \\\\"Port\\\\": 12011, \\\\"SourceCidrIp\\\\": \\\\"0.0.0.0/0\\\\"}]}]"
	ServicesSpec *string `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestAddons) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestAddons) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAddons) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequestAddons) GetResourcesSpec() *string {
	return s.ResourcesSpec
}

func (s *CreateClusterRequestAddons) GetServicesSpec() *string {
	return s.ServicesSpec
}

func (s *CreateClusterRequestAddons) GetVersion() *string {
	return s.Version
}

func (s *CreateClusterRequestAddons) SetName(v string) *CreateClusterRequestAddons {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestAddons) SetResourcesSpec(v string) *CreateClusterRequestAddons {
	s.ResourcesSpec = &v
	return s
}

func (s *CreateClusterRequestAddons) SetServicesSpec(v string) *CreateClusterRequestAddons {
	s.ServicesSpec = &v
	return s
}

func (s *CreateClusterRequestAddons) SetVersion(v string) *CreateClusterRequestAddons {
	s.Version = &v
	return s
}

func (s *CreateClusterRequestAddons) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestClusterCredentials struct {
	// The name of the key pair. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain digits, letters, colons (:), underscores (_), and hyphens (-).
	//
	// >  For more information, see [Create a key pair](https://help.aliyun.com/document_detail/51793.html).
	//
	// example:
	//
	// ali0824
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password for the root user to log on to the node. The password must be 8 to 20 characters in length, and must contain at least 3 of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported: `() ~ ! @ # $ % ^ & 	- - = + { } [ ] : ; \\" < > , . ? /`
	//
	// >  We recommend that you use HTTPS to call the API operation to prevent password leakage.
	//
	// example:
	//
	// **********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateClusterRequestClusterCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestClusterCredentials) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestClusterCredentials) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateClusterRequestClusterCredentials) GetPassword() *string {
	return s.Password
}

func (s *CreateClusterRequestClusterCredentials) SetKeyPairName(v string) *CreateClusterRequestClusterCredentials {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequestClusterCredentials) SetPassword(v string) *CreateClusterRequestClusterCredentials {
	s.Password = &v
	return s
}

func (s *CreateClusterRequestClusterCredentials) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestClusterCustomConfiguration struct {
	// The runtime parameters of the script after the cluster is created.
	//
	// example:
	//
	// E-HPC cn-hangzhou
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// http://*****
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateClusterRequestClusterCustomConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestClusterCustomConfiguration) GetArgs() *string {
	return s.Args
}

func (s *CreateClusterRequestClusterCustomConfiguration) GetScript() *string {
	return s.Script
}

func (s *CreateClusterRequestClusterCustomConfiguration) SetArgs(v string) *CreateClusterRequestClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *CreateClusterRequestClusterCustomConfiguration) SetScript(v string) *CreateClusterRequestClusterCustomConfiguration {
	s.Script = &v
	return s
}

func (s *CreateClusterRequestClusterCustomConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestManager struct {
	// The configurations of the domain name resolution service.
	DNS *CreateClusterRequestManagerDNS `json:"DNS,omitempty" xml:"DNS,omitempty" type:"Struct"`
	// The configurations of the domain account service.
	DirectoryService *CreateClusterRequestManagerDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// The hardware configurations of the management node.
	ManagerNode *NodeTemplate `json:"ManagerNode,omitempty" xml:"ManagerNode,omitempty"`
	// The configurations of the scheduler service.
	Scheduler *CreateClusterRequestManagerScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s CreateClusterRequestManager) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestManager) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManager) GetDNS() *CreateClusterRequestManagerDNS {
	return s.DNS
}

func (s *CreateClusterRequestManager) GetDirectoryService() *CreateClusterRequestManagerDirectoryService {
	return s.DirectoryService
}

func (s *CreateClusterRequestManager) GetManagerNode() *NodeTemplate {
	return s.ManagerNode
}

func (s *CreateClusterRequestManager) GetScheduler() *CreateClusterRequestManagerScheduler {
	return s.Scheduler
}

func (s *CreateClusterRequestManager) SetDNS(v *CreateClusterRequestManagerDNS) *CreateClusterRequestManager {
	s.DNS = v
	return s
}

func (s *CreateClusterRequestManager) SetDirectoryService(v *CreateClusterRequestManagerDirectoryService) *CreateClusterRequestManager {
	s.DirectoryService = v
	return s
}

func (s *CreateClusterRequestManager) SetManagerNode(v *NodeTemplate) *CreateClusterRequestManager {
	s.ManagerNode = v
	return s
}

func (s *CreateClusterRequestManager) SetScheduler(v *CreateClusterRequestManagerScheduler) *CreateClusterRequestManager {
	s.Scheduler = v
	return s
}

func (s *CreateClusterRequestManager) Validate() error {
	if s.DNS != nil {
		if err := s.DNS.Validate(); err != nil {
			return err
		}
	}
	if s.DirectoryService != nil {
		if err := s.DirectoryService.Validate(); err != nil {
			return err
		}
	}
	if s.ManagerNode != nil {
		if err := s.ManagerNode.Validate(); err != nil {
			return err
		}
	}
	if s.Scheduler != nil {
		if err := s.Scheduler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestManagerDNS struct {
	// The domain name resolution type.
	//
	// Valid values:
	//
	// 	- NIS
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain name resolution service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestManagerDNS) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestManagerDNS) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManagerDNS) GetType() *string {
	return s.Type
}

func (s *CreateClusterRequestManagerDNS) GetVersion() *string {
	return s.Version
}

func (s *CreateClusterRequestManagerDNS) SetType(v string) *CreateClusterRequestManagerDNS {
	s.Type = &v
	return s
}

func (s *CreateClusterRequestManagerDNS) SetVersion(v string) *CreateClusterRequestManagerDNS {
	s.Version = &v
	return s
}

func (s *CreateClusterRequestManagerDNS) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestManagerDirectoryService struct {
	// The type of the domain account.
	//
	// Valid values:
	//
	// 	- NIS
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestManagerDirectoryService) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestManagerDirectoryService) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManagerDirectoryService) GetType() *string {
	return s.Type
}

func (s *CreateClusterRequestManagerDirectoryService) GetVersion() *string {
	return s.Version
}

func (s *CreateClusterRequestManagerDirectoryService) SetType(v string) *CreateClusterRequestManagerDirectoryService {
	s.Type = &v
	return s
}

func (s *CreateClusterRequestManagerDirectoryService) SetVersion(v string) *CreateClusterRequestManagerDirectoryService {
	s.Version = &v
	return s
}

func (s *CreateClusterRequestManagerDirectoryService) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestManagerScheduler struct {
	// The scheduler type. Valid values:
	//
	// 	- SLURM
	//
	// 	- PBS
	//
	// 	- OPENGRIDSCHEDULER
	//
	// 	- LSF_PLUGIN
	//
	// 	- PBS_PLUGIN
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateClusterRequestManagerScheduler) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestManagerScheduler) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestManagerScheduler) GetType() *string {
	return s.Type
}

func (s *CreateClusterRequestManagerScheduler) GetVersion() *string {
	return s.Version
}

func (s *CreateClusterRequestManagerScheduler) SetType(v string) *CreateClusterRequestManagerScheduler {
	s.Type = &v
	return s
}

func (s *CreateClusterRequestManagerScheduler) SetVersion(v string) *CreateClusterRequestManagerScheduler {
	s.Version = &v
	return s
}

func (s *CreateClusterRequestManagerScheduler) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestTags struct {
	// The tag key. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ClusterId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ehpc-hz-******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestTags) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateClusterRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateClusterRequestTags) SetKey(v string) *CreateClusterRequestTags {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTags) SetValue(v string) *CreateClusterRequestTags {
	s.Value = &v
	return s
}

func (s *CreateClusterRequestTags) Validate() error {
	return dara.Validate(s)
}
