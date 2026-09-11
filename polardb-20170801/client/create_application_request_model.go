// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIDBClusterId(v string) *CreateApplicationRequest
	GetAIDBClusterId() *string
	SetAgenticDBBranchSpec(v *CreateApplicationRequestAgenticDBBranchSpec) *CreateApplicationRequest
	GetAgenticDBBranchSpec() *CreateApplicationRequestAgenticDBBranchSpec
	SetApplicationType(v string) *CreateApplicationRequest
	GetApplicationType() *string
	SetArchitecture(v string) *CreateApplicationRequest
	GetArchitecture() *string
	SetAuthProvider(v string) *CreateApplicationRequest
	GetAuthProvider() *string
	SetAuthProviderConfig(v string) *CreateApplicationRequest
	GetAuthProviderConfig() *string
	SetAutoAllocatePublicEip(v bool) *CreateApplicationRequest
	GetAutoAllocatePublicEip() *bool
	SetAutoCreatePolarFs(v bool) *CreateApplicationRequest
	GetAutoCreatePolarFs() *bool
	SetAutoRenew(v bool) *CreateApplicationRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *CreateApplicationRequest
	GetAutoUseCoupon() *bool
	SetComponents(v []*CreateApplicationRequestComponents) *CreateApplicationRequest
	GetComponents() []*CreateApplicationRequestComponents
	SetDBClusterId(v string) *CreateApplicationRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateApplicationRequest
	GetDescription() *string
	SetDnatEntries(v []*CreateApplicationRequestDnatEntries) *CreateApplicationRequest
	GetDnatEntries() []*CreateApplicationRequestDnatEntries
	SetDnatIpAddress(v string) *CreateApplicationRequest
	GetDnatIpAddress() *string
	SetDryRun(v bool) *CreateApplicationRequest
	GetDryRun() *bool
	SetEndpoints(v []*CreateApplicationRequestEndpoints) *CreateApplicationRequest
	GetEndpoints() []*CreateApplicationRequestEndpoints
	SetKnowledgeApplicationSpec(v *CreateApplicationRequestKnowledgeApplicationSpec) *CreateApplicationRequest
	GetKnowledgeApplicationSpec() *CreateApplicationRequestKnowledgeApplicationSpec
	SetMemApplicationSpec(v *CreateApplicationRequestMemApplicationSpec) *CreateApplicationRequest
	GetMemApplicationSpec() *CreateApplicationRequestMemApplicationSpec
	SetModelApi(v string) *CreateApplicationRequest
	GetModelApi() *string
	SetModelApiKey(v string) *CreateApplicationRequest
	GetModelApiKey() *string
	SetModelBaseUrl(v string) *CreateApplicationRequest
	GetModelBaseUrl() *string
	SetModelFrom(v string) *CreateApplicationRequest
	GetModelFrom() *string
	SetModelName(v string) *CreateApplicationRequest
	GetModelName() *string
	SetParameters(v []*CreateApplicationRequestParameters) *CreateApplicationRequest
	GetParameters() []*CreateApplicationRequestParameters
	SetPayType(v string) *CreateApplicationRequest
	GetPayType() *string
	SetPeriod(v string) *CreateApplicationRequest
	GetPeriod() *string
	SetPolarFSInstanceId(v string) *CreateApplicationRequest
	GetPolarFSInstanceId() *string
	SetPromotionCode(v string) *CreateApplicationRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateApplicationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApplicationRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateApplicationRequest
	GetSecurityGroupId() *string
	SetSecurityIPArrayName(v string) *CreateApplicationRequest
	GetSecurityIPArrayName() *string
	SetSecurityIPList(v string) *CreateApplicationRequest
	GetSecurityIPList() *string
	SetSecurityIPType(v string) *CreateApplicationRequest
	GetSecurityIPType() *string
	SetSkillTemplateId(v string) *CreateApplicationRequest
	GetSkillTemplateId() *string
	SetStorages(v []*CreateApplicationRequestStorages) *CreateApplicationRequest
	GetStorages() []*CreateApplicationRequestStorages
	SetTag(v []*CreateApplicationRequestTag) *CreateApplicationRequest
	GetTag() []*CreateApplicationRequestTag
	SetTargetVersion(v string) *CreateApplicationRequest
	GetTargetVersion() *string
	SetUsedTime(v string) *CreateApplicationRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateApplicationRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateApplicationRequest
	GetVpcId() *string
	SetVpcNatGatewayId(v string) *CreateApplicationRequest
	GetVpcNatGatewayId() *string
	SetZoneId(v string) *CreateApplicationRequest
	GetZoneId() *string
}

type CreateApplicationRequest struct {
	// The ID of an existing model operator instance to associate. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// pm-xxxxxx
	AIDBClusterId *string `json:"AIDBClusterId,omitempty" xml:"AIDBClusterId,omitempty"`
	// The AgenticDB branch specification.
	//
	// example:
	//
	// {"DBClusterId":"pagc-2zea920mcvd5o87","TenantId":"t-cfc2d7df0e59439681f0087f51","ProjectId":"proj-d7849d0050664c758af795d468","BranchId":"br-9054b3b7649e4c0d977bd0df37","ForkFromBranch":true,"ForkFromApplicationId":"pa-source"}
	AgenticDBBranchSpec *CreateApplicationRequestAgenticDBBranchSpec `json:"AgenticDBBranchSpec,omitempty" xml:"AgenticDBBranchSpec,omitempty" type:"Struct"`
	// The application type. Valid values:
	//
	// - supabase: Set this value to create a managed Supabase application.
	//
	// - raycluster: Set this value to create a managed Ray Cluster application.
	//
	// - polarclaw: Set this value to create a managed PolarClaw application.
	//
	// This parameter is required.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The CPU architecture. Valid values:
	//
	// - x86
	//
	// This parameter is required.
	//
	// example:
	//
	// x86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The authentication service provider.
	//
	// example:
	//
	// feishu
	AuthProvider *string `json:"AuthProvider,omitempty" xml:"AuthProvider,omitempty"`
	// The authentication provider configuration.
	//
	// example:
	//
	// xxx
	AuthProviderConfig *string `json:"AuthProviderConfig,omitempty" xml:"AuthProviderConfig,omitempty"`
	// Specifies whether to enable automatic creation of an elastic IP address (EIP) and attach it to the instance. This is equivalent to associate with an EIP.
	//
	// example:
	//
	// qwen3-max
	AutoAllocatePublicEip *bool `json:"AutoAllocatePublicEip,omitempty" xml:"AutoAllocatePublicEip,omitempty"`
	// Specifies whether to enable automatic creation of a cold storage Polarlakebase instance. Valid values:
	//
	// 	- false (default): Automatic creation is disabled.
	//
	// 	- true: Automatic creation is enabled.
	//
	// example:
	//
	// false
	AutoCreatePolarFs *bool `json:"AutoCreatePolarFs,omitempty" xml:"AutoCreatePolarFs,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- true (default): Use coupons.
	//
	// 	- false: Do not use coupons.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The list of user-defined application subcomponents.
	Components []*CreateApplicationRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The instance ID of the PolarDB instance on which the application depends.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of expected DNAT entries for NAT mapping. Specify this parameter together with VpcNatGatewayId. This parameter can be left empty, which indicates that no DNAT entries are created.
	DnatEntries []*CreateApplicationRequestDnatEntries `json:"DnatEntries,omitempty" xml:"DnatEntries,omitempty" type:"Repeated"`
	// The DNAT-dedicated NAT IP address that has been allocated (separate from the SNAT IP address) for NAT mapping. The IP address must belong to the specified gateway and be in an available state. The vSwitch of the gateway must belong to a primary CIDR block that is reachable from the office network. Specify this parameter together with VpcNatGatewayId. Prerequisite: An SNAT entry has been bound to the vSwitch where the application resides.
	//
	// example:
	//
	// 10.64.0.10
	DnatIpAddress *string `json:"DnatIpAddress,omitempty" xml:"DnatIpAddress,omitempty"`
	// Default value: `false`. If you set this parameter to `true`, only parameter and resource validation is performed without actually creating the resource.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of user-defined service endpoints. By default, a VPC endpoint is created.
	Endpoints []*CreateApplicationRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// Required for knowledge applications.
	KnowledgeApplicationSpec *CreateApplicationRequestKnowledgeApplicationSpec `json:"KnowledgeApplicationSpec,omitempty" xml:"KnowledgeApplicationSpec,omitempty" type:"Struct"`
	// Required for mem0 applications.
	MemApplicationSpec *CreateApplicationRequestMemApplicationSpec `json:"MemApplicationSpec,omitempty" xml:"MemApplicationSpec,omitempty" type:"Struct"`
	// The model API. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// openai-completions
	ModelApi *string `json:"ModelApi,omitempty" xml:"ModelApi,omitempty"`
	// The model API key. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// sk-xxxxxx
	ModelApiKey *string `json:"ModelApiKey,omitempty" xml:"ModelApiKey,omitempty"`
	// The model base URL. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	ModelBaseUrl *string `json:"ModelBaseUrl,omitempty" xml:"ModelBaseUrl,omitempty"`
	// The model source. Valid values:
	//
	// 	- bailian: Alibaba Cloud Model Studio model.
	//
	// 	- custom: Custom model.
	//
	// 	- maas: PolarDB model operator.
	//
	// example:
	//
	// bailian
	ModelFrom *string `json:"ModelFrom,omitempty" xml:"ModelFrom,omitempty"`
	// The model name. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The list of parameters.
	Parameters []*CreateApplicationRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The billing type.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type (yearly or monthly).
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The instance ID of the Polarlakebase cold storage or high-performance edition. Default value: empty. If specified, the corresponding storage is mounted to the application.
	//
	// Currently, only the following applications support this parameter:
	//
	// - supabase
	//
	// - raycluster
	//
	// example:
	//
	// pcs-********************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// The coupon code. If you do not specify this parameter, the default coupon is used.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region. Default value: the region of the instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-********************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the IP whitelist group. Default value: `default`.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The IP whitelist. If you do not specify this parameter, the default value is `127.0.0.1`.
	//
	// example:
	//
	// 127.0.0.1,172.17.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The type of the IP address.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The skill template ID.
	//
	// example:
	//
	// xxx
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
	// The list of application storages.
	//
	// example:
	//
	// [{"StorageType":"oss","StorageInstanceId":"pfs-xxxx","EndpointId":"pe-xxxx"}]
	Storages []*CreateApplicationRequestStorages `json:"Storages,omitempty" xml:"Storages,omitempty" type:"Repeated"`
	// The tags.
	Tag []*CreateApplicationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The target version.
	//
	// example:
	//
	// latest
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch. Default value: the vSwitch in the primary zone of the instance.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-********************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC NAT gateway ID for NAT mapping. If specified, NAT mapping is enabled when the instance is created. The NAT gateway must be in the same VPC as the application, use the private network type (intranet), and be in an active state.
	//
	// example:
	//
	// ngw-xxx
	VpcNatGatewayId *string `json:"VpcNatGatewayId,omitempty" xml:"VpcNatGatewayId,omitempty"`
	// The zone. Default value: the primary zone of the instance.
	//
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetAIDBClusterId() *string {
	return s.AIDBClusterId
}

func (s *CreateApplicationRequest) GetAgenticDBBranchSpec() *CreateApplicationRequestAgenticDBBranchSpec {
	return s.AgenticDBBranchSpec
}

func (s *CreateApplicationRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateApplicationRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateApplicationRequest) GetAuthProvider() *string {
	return s.AuthProvider
}

func (s *CreateApplicationRequest) GetAuthProviderConfig() *string {
	return s.AuthProviderConfig
}

func (s *CreateApplicationRequest) GetAutoAllocatePublicEip() *bool {
	return s.AutoAllocatePublicEip
}

func (s *CreateApplicationRequest) GetAutoCreatePolarFs() *bool {
	return s.AutoCreatePolarFs
}

func (s *CreateApplicationRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateApplicationRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateApplicationRequest) GetComponents() []*CreateApplicationRequestComponents {
	return s.Components
}

func (s *CreateApplicationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequest) GetDnatEntries() []*CreateApplicationRequestDnatEntries {
	return s.DnatEntries
}

func (s *CreateApplicationRequest) GetDnatIpAddress() *string {
	return s.DnatIpAddress
}

func (s *CreateApplicationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateApplicationRequest) GetEndpoints() []*CreateApplicationRequestEndpoints {
	return s.Endpoints
}

func (s *CreateApplicationRequest) GetKnowledgeApplicationSpec() *CreateApplicationRequestKnowledgeApplicationSpec {
	return s.KnowledgeApplicationSpec
}

func (s *CreateApplicationRequest) GetMemApplicationSpec() *CreateApplicationRequestMemApplicationSpec {
	return s.MemApplicationSpec
}

func (s *CreateApplicationRequest) GetModelApi() *string {
	return s.ModelApi
}

func (s *CreateApplicationRequest) GetModelApiKey() *string {
	return s.ModelApiKey
}

func (s *CreateApplicationRequest) GetModelBaseUrl() *string {
	return s.ModelBaseUrl
}

func (s *CreateApplicationRequest) GetModelFrom() *string {
	return s.ModelFrom
}

func (s *CreateApplicationRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateApplicationRequest) GetParameters() []*CreateApplicationRequestParameters {
	return s.Parameters
}

func (s *CreateApplicationRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateApplicationRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateApplicationRequest) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *CreateApplicationRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateApplicationRequest) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *CreateApplicationRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateApplicationRequest) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *CreateApplicationRequest) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *CreateApplicationRequest) GetStorages() []*CreateApplicationRequestStorages {
	return s.Storages
}

func (s *CreateApplicationRequest) GetTag() []*CreateApplicationRequestTag {
	return s.Tag
}

func (s *CreateApplicationRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateApplicationRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateApplicationRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateApplicationRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateApplicationRequest) GetVpcNatGatewayId() *string {
	return s.VpcNatGatewayId
}

func (s *CreateApplicationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateApplicationRequest) SetAIDBClusterId(v string) *CreateApplicationRequest {
	s.AIDBClusterId = &v
	return s
}

func (s *CreateApplicationRequest) SetAgenticDBBranchSpec(v *CreateApplicationRequestAgenticDBBranchSpec) *CreateApplicationRequest {
	s.AgenticDBBranchSpec = v
	return s
}

func (s *CreateApplicationRequest) SetApplicationType(v string) *CreateApplicationRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateApplicationRequest) SetArchitecture(v string) *CreateApplicationRequest {
	s.Architecture = &v
	return s
}

func (s *CreateApplicationRequest) SetAuthProvider(v string) *CreateApplicationRequest {
	s.AuthProvider = &v
	return s
}

func (s *CreateApplicationRequest) SetAuthProviderConfig(v string) *CreateApplicationRequest {
	s.AuthProviderConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoAllocatePublicEip(v bool) *CreateApplicationRequest {
	s.AutoAllocatePublicEip = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoCreatePolarFs(v bool) *CreateApplicationRequest {
	s.AutoCreatePolarFs = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoRenew(v bool) *CreateApplicationRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoUseCoupon(v bool) *CreateApplicationRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateApplicationRequest) SetComponents(v []*CreateApplicationRequestComponents) *CreateApplicationRequest {
	s.Components = v
	return s
}

func (s *CreateApplicationRequest) SetDBClusterId(v string) *CreateApplicationRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetDnatEntries(v []*CreateApplicationRequestDnatEntries) *CreateApplicationRequest {
	s.DnatEntries = v
	return s
}

func (s *CreateApplicationRequest) SetDnatIpAddress(v string) *CreateApplicationRequest {
	s.DnatIpAddress = &v
	return s
}

func (s *CreateApplicationRequest) SetDryRun(v bool) *CreateApplicationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateApplicationRequest) SetEndpoints(v []*CreateApplicationRequestEndpoints) *CreateApplicationRequest {
	s.Endpoints = v
	return s
}

func (s *CreateApplicationRequest) SetKnowledgeApplicationSpec(v *CreateApplicationRequestKnowledgeApplicationSpec) *CreateApplicationRequest {
	s.KnowledgeApplicationSpec = v
	return s
}

func (s *CreateApplicationRequest) SetMemApplicationSpec(v *CreateApplicationRequestMemApplicationSpec) *CreateApplicationRequest {
	s.MemApplicationSpec = v
	return s
}

func (s *CreateApplicationRequest) SetModelApi(v string) *CreateApplicationRequest {
	s.ModelApi = &v
	return s
}

func (s *CreateApplicationRequest) SetModelApiKey(v string) *CreateApplicationRequest {
	s.ModelApiKey = &v
	return s
}

func (s *CreateApplicationRequest) SetModelBaseUrl(v string) *CreateApplicationRequest {
	s.ModelBaseUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetModelFrom(v string) *CreateApplicationRequest {
	s.ModelFrom = &v
	return s
}

func (s *CreateApplicationRequest) SetModelName(v string) *CreateApplicationRequest {
	s.ModelName = &v
	return s
}

func (s *CreateApplicationRequest) SetParameters(v []*CreateApplicationRequestParameters) *CreateApplicationRequest {
	s.Parameters = v
	return s
}

func (s *CreateApplicationRequest) SetPayType(v string) *CreateApplicationRequest {
	s.PayType = &v
	return s
}

func (s *CreateApplicationRequest) SetPeriod(v string) *CreateApplicationRequest {
	s.Period = &v
	return s
}

func (s *CreateApplicationRequest) SetPolarFSInstanceId(v string) *CreateApplicationRequest {
	s.PolarFSInstanceId = &v
	return s
}

func (s *CreateApplicationRequest) SetPromotionCode(v string) *CreateApplicationRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateApplicationRequest) SetRegionId(v string) *CreateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetSecurityGroupId(v string) *CreateApplicationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetSecurityIPArrayName(v string) *CreateApplicationRequest {
	s.SecurityIPArrayName = &v
	return s
}

func (s *CreateApplicationRequest) SetSecurityIPList(v string) *CreateApplicationRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateApplicationRequest) SetSecurityIPType(v string) *CreateApplicationRequest {
	s.SecurityIPType = &v
	return s
}

func (s *CreateApplicationRequest) SetSkillTemplateId(v string) *CreateApplicationRequest {
	s.SkillTemplateId = &v
	return s
}

func (s *CreateApplicationRequest) SetStorages(v []*CreateApplicationRequestStorages) *CreateApplicationRequest {
	s.Storages = v
	return s
}

func (s *CreateApplicationRequest) SetTag(v []*CreateApplicationRequestTag) *CreateApplicationRequest {
	s.Tag = v
	return s
}

func (s *CreateApplicationRequest) SetTargetVersion(v string) *CreateApplicationRequest {
	s.TargetVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetUsedTime(v string) *CreateApplicationRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateApplicationRequest) SetVSwitchId(v string) *CreateApplicationRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationRequest) SetVpcId(v string) *CreateApplicationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateApplicationRequest) SetVpcNatGatewayId(v string) *CreateApplicationRequest {
	s.VpcNatGatewayId = &v
	return s
}

func (s *CreateApplicationRequest) SetZoneId(v string) *CreateApplicationRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	if s.AgenticDBBranchSpec != nil {
		if err := s.AgenticDBBranchSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DnatEntries != nil {
		for _, item := range s.DnatEntries {
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
	if s.KnowledgeApplicationSpec != nil {
		if err := s.KnowledgeApplicationSpec.Validate(); err != nil {
			return err
		}
	}
	if s.MemApplicationSpec != nil {
		if err := s.MemApplicationSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApplicationRequestAgenticDBBranchSpec struct {
	// The AgenticDB branch ID.
	//
	// example:
	//
	// br-9054b3b7649e4c0d977bd0df37
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The AgenticDB cluster ID.
	//
	// example:
	//
	// pagc-2zea920mcvd5o87
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the source application.
	//
	// example:
	//
	// pa-source
	ForkFromApplicationId *string `json:"ForkFromApplicationId,omitempty" xml:"ForkFromApplicationId,omitempty"`
	// Specifies whether to create the application based on a specified AgenticDB branch.
	//
	// example:
	//
	// true
	ForkFromBranch *bool `json:"ForkFromBranch,omitempty" xml:"ForkFromBranch,omitempty"`
	// The AgenticDB project ID.
	//
	// example:
	//
	// proj-d7849d0050664c758af795d468
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The AgenticDB tenant ID.
	//
	// example:
	//
	// t-cfc2d7df0e59439681f0087f51
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateApplicationRequestAgenticDBBranchSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestAgenticDBBranchSpec) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) GetBranchId() *string {
	return s.BranchId
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) GetForkFromApplicationId() *string {
	return s.ForkFromApplicationId
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) GetForkFromBranch() *bool {
	return s.ForkFromBranch
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) SetBranchId(v string) *CreateApplicationRequestAgenticDBBranchSpec {
	s.BranchId = &v
	return s
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) SetDBClusterId(v string) *CreateApplicationRequestAgenticDBBranchSpec {
	s.DBClusterId = &v
	return s
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) SetForkFromApplicationId(v string) *CreateApplicationRequestAgenticDBBranchSpec {
	s.ForkFromApplicationId = &v
	return s
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) SetForkFromBranch(v bool) *CreateApplicationRequestAgenticDBBranchSpec {
	s.ForkFromBranch = &v
	return s
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) SetProjectId(v string) *CreateApplicationRequestAgenticDBBranchSpec {
	s.ProjectId = &v
	return s
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) SetTenantId(v string) *CreateApplicationRequestAgenticDBBranchSpec {
	s.TenantId = &v
	return s
}

func (s *CreateApplicationRequestAgenticDBBranchSpec) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestComponents struct {
	// The specification of the application subcomponent.
	//
	// example:
	//
	// polar.app.g2.medium
	ComponentClass *string `json:"ComponentClass,omitempty" xml:"ComponentClass,omitempty"`
	// The maximum number of replicas for the application subcomponent with the same specification. Default value: the value of ComponentReplica.
	//
	// - Only raycluster supports this parameter.
	//
	// example:
	//
	// 1
	ComponentMaxReplica *int64 `json:"ComponentMaxReplica,omitempty" xml:"ComponentMaxReplica,omitempty"`
	// The number of replicas for the application subcomponent. Default value: 1.
	//
	// example:
	//
	// 1
	ComponentReplica *int64 `json:"ComponentReplica,omitempty" xml:"ComponentReplica,omitempty"`
	// The type of the application subcomponent.
	//
	// For supabase, valid values:
	//
	// - gateway
	//
	// - backend
	//
	// For raycluster, valid values:
	//
	// - head
	//
	// - worker
	//
	// - gpuworker
	//
	// example:
	//
	// gateway
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The maximum number of replicas for component scaling.
	//
	// example:
	//
	// 16
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum number of replicas for component scaling.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The list of security groups for the application subcomponent, separated by commas (,).
	//
	// example:
	//
	// sg-********************
	SecurityGroups *string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty"`
	// The name of the whitelist IP address group for the application subcomponent. Default value: default.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The whitelist IP addresses of the application subcomponent, separated by commas (,).
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The type of the whitelist IP addresses for the application subcomponent. Default value: ipv4.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
}

func (s CreateApplicationRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestComponents) GetComponentClass() *string {
	return s.ComponentClass
}

func (s *CreateApplicationRequestComponents) GetComponentMaxReplica() *int64 {
	return s.ComponentMaxReplica
}

func (s *CreateApplicationRequestComponents) GetComponentReplica() *int64 {
	return s.ComponentReplica
}

func (s *CreateApplicationRequestComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *CreateApplicationRequestComponents) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *CreateApplicationRequestComponents) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *CreateApplicationRequestComponents) GetSecurityGroups() *string {
	return s.SecurityGroups
}

func (s *CreateApplicationRequestComponents) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *CreateApplicationRequestComponents) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateApplicationRequestComponents) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *CreateApplicationRequestComponents) SetComponentClass(v string) *CreateApplicationRequestComponents {
	s.ComponentClass = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetComponentMaxReplica(v int64) *CreateApplicationRequestComponents {
	s.ComponentMaxReplica = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetComponentReplica(v int64) *CreateApplicationRequestComponents {
	s.ComponentReplica = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetComponentType(v string) *CreateApplicationRequestComponents {
	s.ComponentType = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetScaleMax(v string) *CreateApplicationRequestComponents {
	s.ScaleMax = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetScaleMin(v string) *CreateApplicationRequestComponents {
	s.ScaleMin = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityGroups(v string) *CreateApplicationRequestComponents {
	s.SecurityGroups = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityIPArrayName(v string) *CreateApplicationRequestComponents {
	s.SecurityIPArrayName = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityIPList(v string) *CreateApplicationRequestComponents {
	s.SecurityIPList = &v
	return s
}

func (s *CreateApplicationRequestComponents) SetSecurityIPType(v string) *CreateApplicationRequestComponents {
	s.SecurityIPType = &v
	return s
}

func (s *CreateApplicationRequestComponents) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestDnatEntries struct {
	// The frontend port. This parameter is optional. If not specified, the system automatically assigns a port that does not conflict with ports already in use on the gateway. You can query the assignment result by calling the DescribeApplicationAttribute operation.
	//
	// example:
	//
	// 10001
	FrontPort *int32 `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	// The port name. Valid values: webui, hermesagent, dashboard, and ssh.
	//
	// example:
	//
	// webui
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
}

func (s CreateApplicationRequestDnatEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestDnatEntries) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestDnatEntries) GetFrontPort() *int32 {
	return s.FrontPort
}

func (s *CreateApplicationRequestDnatEntries) GetPortName() *string {
	return s.PortName
}

func (s *CreateApplicationRequestDnatEntries) SetFrontPort(v int32) *CreateApplicationRequestDnatEntries {
	s.FrontPort = &v
	return s
}

func (s *CreateApplicationRequestDnatEntries) SetPortName(v string) *CreateApplicationRequestDnatEntries {
	s.PortName = &v
	return s
}

func (s *CreateApplicationRequestDnatEntries) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestEndpoints struct {
	// The description of the service endpoint.
	//
	// example:
	//
	// my_endpoint
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the service endpoint. The value is fixed as Primary.
	//
	// example:
	//
	// Primary
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s CreateApplicationRequestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestEndpoints) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequestEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateApplicationRequestEndpoints) SetDescription(v string) *CreateApplicationRequestEndpoints {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequestEndpoints) SetEndpointType(v string) *CreateApplicationRequestEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateApplicationRequestEndpoints) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestKnowledgeApplicationSpec struct {
	// The dashboard password.
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The password.
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// Required for knowledge applications. The LLM model name, such as qwen3-max.
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
}

func (s CreateApplicationRequestKnowledgeApplicationSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestKnowledgeApplicationSpec) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) GetDbPassword() *string {
	return s.DbPassword
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) GetLlmModel() *string {
	return s.LlmModel
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) SetDashboardPassword(v string) *CreateApplicationRequestKnowledgeApplicationSpec {
	s.DashboardPassword = &v
	return s
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) SetDbPassword(v string) *CreateApplicationRequestKnowledgeApplicationSpec {
	s.DbPassword = &v
	return s
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) SetLlmModel(v string) *CreateApplicationRequestKnowledgeApplicationSpec {
	s.LlmModel = &v
	return s
}

func (s *CreateApplicationRequestKnowledgeApplicationSpec) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestMemApplicationSpec struct {
	// The database name.
	//
	// example:
	//
	// test-database-name
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The password.
	//
	// example:
	//
	// test-user-password
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// The username.
	//
	// example:
	//
	// test-user
	DbUser *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	// Required for mem0 applications. The embedder model name, such as text-embedding-v4.
	//
	// example:
	//
	// text-embedding-v4
	EmbedderModel *string `json:"EmbedderModel,omitempty" xml:"EmbedderModel,omitempty"`
	// The vector dimensions.
	//
	// example:
	//
	// 1024
	EmbedderModelDimension *int32 `json:"EmbedderModelDimension,omitempty" xml:"EmbedderModelDimension,omitempty"`
	// The graph LLM model.
	//
	// example:
	//
	// qwen-plus
	GraphLlmModel *string `json:"GraphLlmModel,omitempty" xml:"GraphLlmModel,omitempty"`
	// Required for mem0 applications. The LLM model name, such as qwen3-max.
	//
	// example:
	//
	// qwen3-max
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The project name, which corresponds to the database schema that stores project data.
	//
	// example:
	//
	// test-project-name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Required for mem0 applications. The reranker model name, such as qwen3-rerank.
	//
	// example:
	//
	// qwen3-rerank
	RerankerModel *string `json:"RerankerModel,omitempty" xml:"RerankerModel,omitempty"`
	// The number of table shards.
	//
	// example:
	//
	// 1
	Shard *int32 `json:"Shard,omitempty" xml:"Shard,omitempty"`
}

func (s CreateApplicationRequestMemApplicationSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestMemApplicationSpec) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestMemApplicationSpec) GetDbName() *string {
	return s.DbName
}

func (s *CreateApplicationRequestMemApplicationSpec) GetDbPassword() *string {
	return s.DbPassword
}

func (s *CreateApplicationRequestMemApplicationSpec) GetDbUser() *string {
	return s.DbUser
}

func (s *CreateApplicationRequestMemApplicationSpec) GetEmbedderModel() *string {
	return s.EmbedderModel
}

func (s *CreateApplicationRequestMemApplicationSpec) GetEmbedderModelDimension() *int32 {
	return s.EmbedderModelDimension
}

func (s *CreateApplicationRequestMemApplicationSpec) GetGraphLlmModel() *string {
	return s.GraphLlmModel
}

func (s *CreateApplicationRequestMemApplicationSpec) GetLlmModel() *string {
	return s.LlmModel
}

func (s *CreateApplicationRequestMemApplicationSpec) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateApplicationRequestMemApplicationSpec) GetRerankerModel() *string {
	return s.RerankerModel
}

func (s *CreateApplicationRequestMemApplicationSpec) GetShard() *int32 {
	return s.Shard
}

func (s *CreateApplicationRequestMemApplicationSpec) SetDbName(v string) *CreateApplicationRequestMemApplicationSpec {
	s.DbName = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetDbPassword(v string) *CreateApplicationRequestMemApplicationSpec {
	s.DbPassword = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetDbUser(v string) *CreateApplicationRequestMemApplicationSpec {
	s.DbUser = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetEmbedderModel(v string) *CreateApplicationRequestMemApplicationSpec {
	s.EmbedderModel = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetEmbedderModelDimension(v int32) *CreateApplicationRequestMemApplicationSpec {
	s.EmbedderModelDimension = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetGraphLlmModel(v string) *CreateApplicationRequestMemApplicationSpec {
	s.GraphLlmModel = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetLlmModel(v string) *CreateApplicationRequestMemApplicationSpec {
	s.LlmModel = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetProjectName(v string) *CreateApplicationRequestMemApplicationSpec {
	s.ProjectName = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetRerankerModel(v string) *CreateApplicationRequestMemApplicationSpec {
	s.RerankerModel = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) SetShard(v int32) *CreateApplicationRequestMemApplicationSpec {
	s.Shard = &v
	return s
}

func (s *CreateApplicationRequestMemApplicationSpec) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestParameters struct {
	// The parameter name.
	//
	// example:
	//
	// secret.gateway.auth.token
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// TK***
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateApplicationRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *CreateApplicationRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateApplicationRequestParameters) SetParameterName(v string) *CreateApplicationRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateApplicationRequestParameters) SetParameterValue(v string) *CreateApplicationRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateApplicationRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestStorages struct {
	// The mount path inside the container.
	//
	// example:
	//
	// /data/container
	ContainerMountPath *string `json:"ContainerMountPath,omitempty" xml:"ContainerMountPath,omitempty"`
	// The storage endpoint ID.
	//
	// example:
	//
	// pe-xxxx
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The storage mount path.
	//
	// example:
	//
	// /data/source
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 100
	StorageCapacity *string `json:"StorageCapacity,omitempty" xml:"StorageCapacity,omitempty"`
	// The storage access endpoint.
	//
	// example:
	//
	// polarfs.example.com
	StorageEndpoint *string `json:"StorageEndpoint,omitempty" xml:"StorageEndpoint,omitempty"`
	// The storage instance ID.
	//
	// example:
	//
	// pfs-xxxx
	StorageInstanceId *string `json:"StorageInstanceId,omitempty" xml:"StorageInstanceId,omitempty"`
	// The storage performance level.
	//
	// example:
	//
	// PL1
	StoragePerformanceLevel *string `json:"StoragePerformanceLevel,omitempty" xml:"StoragePerformanceLevel,omitempty"`
	// The storage type.
	//
	// example:
	//
	// oss
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateApplicationRequestStorages) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestStorages) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestStorages) GetContainerMountPath() *string {
	return s.ContainerMountPath
}

func (s *CreateApplicationRequestStorages) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateApplicationRequestStorages) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateApplicationRequestStorages) GetStorageCapacity() *string {
	return s.StorageCapacity
}

func (s *CreateApplicationRequestStorages) GetStorageEndpoint() *string {
	return s.StorageEndpoint
}

func (s *CreateApplicationRequestStorages) GetStorageInstanceId() *string {
	return s.StorageInstanceId
}

func (s *CreateApplicationRequestStorages) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *CreateApplicationRequestStorages) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateApplicationRequestStorages) SetContainerMountPath(v string) *CreateApplicationRequestStorages {
	s.ContainerMountPath = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetEndpointId(v string) *CreateApplicationRequestStorages {
	s.EndpointId = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetMountPath(v string) *CreateApplicationRequestStorages {
	s.MountPath = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetStorageCapacity(v string) *CreateApplicationRequestStorages {
	s.StorageCapacity = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetStorageEndpoint(v string) *CreateApplicationRequestStorages {
	s.StorageEndpoint = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetStorageInstanceId(v string) *CreateApplicationRequestStorages {
	s.StorageInstanceId = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetStoragePerformanceLevel(v string) *CreateApplicationRequestStorages {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *CreateApplicationRequestStorages) SetStorageType(v string) *CreateApplicationRequestStorages {
	s.StorageType = &v
	return s
}

func (s *CreateApplicationRequestStorages) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApplicationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateApplicationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateApplicationRequestTag) SetKey(v string) *CreateApplicationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateApplicationRequestTag) SetValue(v string) *CreateApplicationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateApplicationRequestTag) Validate() error {
	return dara.Validate(s)
}
