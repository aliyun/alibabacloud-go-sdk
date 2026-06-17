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
	SetZoneId(v string) *CreateApplicationRequest
	GetZoneId() *string
}

type CreateApplicationRequest struct {
	// The ID of an existing model operator instance to associate. This parameter is effective only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// pm-xxxxxx
	AIDBClusterId *string `json:"AIDBClusterId,omitempty" xml:"AIDBClusterId,omitempty"`
	// The type of the application. Valid values:
	//
	// - supabase: Creates a managed Supabase application.
	//
	// - raycluster: Creates a managed Ray Cluster application.
	//
	// - polarclaw: Creates a managed PolarClaw application.
	//
	// This parameter is required.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The CPU architecture. Valid value:
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
	// The configuration of the authentication provider.
	//
	// example:
	//
	// xxx
	AuthProviderConfig *string `json:"AuthProviderConfig,omitempty" xml:"AuthProviderConfig,omitempty"`
	// Specifies whether to automatically create and bind an Elastic IP Address (EIP).
	//
	// example:
	//
	// qwen3-max
	AutoAllocatePublicEip *bool `json:"AutoAllocatePublicEip,omitempty" xml:"AutoAllocatePublicEip,omitempty"`
	// Specifies whether to automatically create a PolarFS cold storage instance. Valid values:
	//
	// - false (default): Does not automatically create the instance.
	//
	// - true: Automatically creates the instance.
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
	// Specifies whether to automatically use a coupon. Valid values:
	//
	// - true (default): Uses a coupon.
	//
	// - false: Does not use a coupon.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// A list of custom child components for the application.
	Components []*CreateApplicationRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The ID of the PolarDB instance that the application depends on.
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
	// The default value is `false`. If you set this parameter to `true`, the system only checks the parameters and resources without creating the actual resources.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// A list of custom server-side endpoints. By default, a VPC Endpoint is created.
	Endpoints []*CreateApplicationRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// This parameter is required for knowledge applications.
	KnowledgeApplicationSpec *CreateApplicationRequestKnowledgeApplicationSpec `json:"KnowledgeApplicationSpec,omitempty" xml:"KnowledgeApplicationSpec,omitempty" type:"Struct"`
	// This parameter is required for mem0 applications.
	MemApplicationSpec *CreateApplicationRequestMemApplicationSpec `json:"MemApplicationSpec,omitempty" xml:"MemApplicationSpec,omitempty" type:"Struct"`
	// The model API. This parameter is effective only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// openai-completions
	ModelApi *string `json:"ModelApi,omitempty" xml:"ModelApi,omitempty"`
	// The API key for the model. This parameter is effective only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// sk-xxxxxx
	ModelApiKey *string `json:"ModelApiKey,omitempty" xml:"ModelApiKey,omitempty"`
	// The URL of the model. This parameter is effective only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	ModelBaseUrl *string `json:"ModelBaseUrl,omitempty" xml:"ModelBaseUrl,omitempty"`
	// The source of the model. Valid values:
	//
	// - bailian: Alibaba Cloud Model Studio model.
	//
	// - custom: A custom model.
	//
	// - maas: PolarDB model operator.
	//
	// example:
	//
	// bailian
	ModelFrom *string `json:"ModelFrom,omitempty" xml:"ModelFrom,omitempty"`
	// The name of the model. This parameter is effective only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// A list of parameters.
	Parameters []*CreateApplicationRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The billing method.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period type.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the PolarFileSystem (PolarFS) cold storage or high-performance instance. This parameter is empty by default. If you specify this parameter, the corresponding storage is mounted to the application.
	//
	// This feature is currently supported only by the following applications:
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
	// The region. The default value is the region of the instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-********************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the IP address whitelist group. The default value is `default`.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The IP address whitelist. If you do not specify this parameter, the default value `127.0.0.1` is used.
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
	// The ID of the skill template.
	//
	// example:
	//
	// xxx
	SkillTemplateId *string `json:"SkillTemplateId,omitempty" xml:"SkillTemplateId,omitempty"`
	// The tag.
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
	// The vSwitch. The default value is the current vSwitch in the primary zone of the instance.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the Virtual Private Cloud (VPC).
	//
	// example:
	//
	// vpc-********************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone. The default value is the primary zone of the instance.
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

func (s *CreateApplicationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateApplicationRequest) SetAIDBClusterId(v string) *CreateApplicationRequest {
	s.AIDBClusterId = &v
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

func (s *CreateApplicationRequest) SetZoneId(v string) *CreateApplicationRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
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

type CreateApplicationRequestComponents struct {
	// The specifications of the child component.
	//
	// example:
	//
	// polar.app.g2.medium
	ComponentClass *string `json:"ComponentClass,omitempty" xml:"ComponentClass,omitempty"`
	// The maximum number of child components with the same specifications. The default value is the value of ComponentReplica.
	//
	// - This parameter is supported only for raycluster.
	//
	// example:
	//
	// 1
	ComponentMaxReplica *int64 `json:"ComponentMaxReplica,omitempty" xml:"ComponentMaxReplica,omitempty"`
	// The number of replicas for the child component. The default value is 1.
	//
	// example:
	//
	// 1
	ComponentReplica *int64 `json:"ComponentReplica,omitempty" xml:"ComponentReplica,omitempty"`
	// The type of the child component.
	//
	// For supabase, valid values are:
	//
	// - gateway
	//
	// - backend
	//
	// For raycluster, valid values are:
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
	// The maximum number of component replicas for scaling.
	//
	// example:
	//
	// 16
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum number of component replicas for scaling.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The security groups for the child component. Separate multiple security group IDs with commas (,).
	//
	// example:
	//
	// sg-********************
	SecurityGroups *string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty"`
	// The name of the IP address whitelist group for the child component. The default value is default.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The IP address whitelist for the child component. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The type of the IP address in the whitelist for the child component. The default value is ipv4.
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

type CreateApplicationRequestEndpoints struct {
	// The description of the server-side endpoint.
	//
	// example:
	//
	// my_endpoint
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the server-side endpoint. This value is fixed to Primary.
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
	// The password for the dashboard.
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The password.
	DbPassword *string `json:"DbPassword,omitempty" xml:"DbPassword,omitempty"`
	// This parameter is required for knowledge applications. It specifies the name of the LLM, such as qwen3-max.
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
	// The name of the database.
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
	// This parameter is required for mem0 applications. It specifies the name of the embedder model, such as text-embedding-v4.
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
	// The graph LLM.
	//
	// example:
	//
	// qwen-plus
	GraphLlmModel *string `json:"GraphLlmModel,omitempty" xml:"GraphLlmModel,omitempty"`
	// This parameter is required for mem0 applications. It specifies the name of the large language model (LLM), such as qwen3-max.
	//
	// example:
	//
	// qwen3-max
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The project name. This corresponds to the schema in the database where project data is stored.
	//
	// example:
	//
	// test-project-name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required for mem0 applications. It specifies the name of the reranker model, such as qwen3-rerank.
	//
	// example:
	//
	// qwen3-rerank
	RerankerModel *string `json:"RerankerModel,omitempty" xml:"RerankerModel,omitempty"`
	// The number of sharded tables.
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
	// The name of the parameter.
	//
	// example:
	//
	// secret.gateway.auth.token
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter.
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

type CreateApplicationRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
