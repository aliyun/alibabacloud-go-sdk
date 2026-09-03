// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIDBClusterId(v string) *CreateApplicationShrinkRequest
	GetAIDBClusterId() *string
	SetApplicationType(v string) *CreateApplicationShrinkRequest
	GetApplicationType() *string
	SetArchitecture(v string) *CreateApplicationShrinkRequest
	GetArchitecture() *string
	SetAuthProvider(v string) *CreateApplicationShrinkRequest
	GetAuthProvider() *string
	SetAuthProviderConfig(v string) *CreateApplicationShrinkRequest
	GetAuthProviderConfig() *string
	SetAutoAllocatePublicEip(v bool) *CreateApplicationShrinkRequest
	GetAutoAllocatePublicEip() *bool
	SetAutoCreatePolarFs(v bool) *CreateApplicationShrinkRequest
	GetAutoCreatePolarFs() *bool
	SetAutoRenew(v bool) *CreateApplicationShrinkRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *CreateApplicationShrinkRequest
	GetAutoUseCoupon() *bool
	SetComponentsShrink(v string) *CreateApplicationShrinkRequest
	GetComponentsShrink() *string
	SetDBClusterId(v string) *CreateApplicationShrinkRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateApplicationShrinkRequest
	GetDescription() *string
	SetDnatEntriesShrink(v string) *CreateApplicationShrinkRequest
	GetDnatEntriesShrink() *string
	SetDnatIpAddress(v string) *CreateApplicationShrinkRequest
	GetDnatIpAddress() *string
	SetDryRun(v bool) *CreateApplicationShrinkRequest
	GetDryRun() *bool
	SetEndpointsShrink(v string) *CreateApplicationShrinkRequest
	GetEndpointsShrink() *string
	SetKnowledgeApplicationSpecShrink(v string) *CreateApplicationShrinkRequest
	GetKnowledgeApplicationSpecShrink() *string
	SetMemApplicationSpecShrink(v string) *CreateApplicationShrinkRequest
	GetMemApplicationSpecShrink() *string
	SetModelApi(v string) *CreateApplicationShrinkRequest
	GetModelApi() *string
	SetModelApiKey(v string) *CreateApplicationShrinkRequest
	GetModelApiKey() *string
	SetModelBaseUrl(v string) *CreateApplicationShrinkRequest
	GetModelBaseUrl() *string
	SetModelFrom(v string) *CreateApplicationShrinkRequest
	GetModelFrom() *string
	SetModelName(v string) *CreateApplicationShrinkRequest
	GetModelName() *string
	SetParametersShrink(v string) *CreateApplicationShrinkRequest
	GetParametersShrink() *string
	SetPayType(v string) *CreateApplicationShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CreateApplicationShrinkRequest
	GetPeriod() *string
	SetPolarFSInstanceId(v string) *CreateApplicationShrinkRequest
	GetPolarFSInstanceId() *string
	SetPromotionCode(v string) *CreateApplicationShrinkRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateApplicationShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApplicationShrinkRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateApplicationShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityIPArrayName(v string) *CreateApplicationShrinkRequest
	GetSecurityIPArrayName() *string
	SetSecurityIPList(v string) *CreateApplicationShrinkRequest
	GetSecurityIPList() *string
	SetSecurityIPType(v string) *CreateApplicationShrinkRequest
	GetSecurityIPType() *string
	SetSkillTemplateId(v string) *CreateApplicationShrinkRequest
	GetSkillTemplateId() *string
	SetTag(v []*CreateApplicationShrinkRequestTag) *CreateApplicationShrinkRequest
	GetTag() []*CreateApplicationShrinkRequestTag
	SetTargetVersion(v string) *CreateApplicationShrinkRequest
	GetTargetVersion() *string
	SetUsedTime(v string) *CreateApplicationShrinkRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateApplicationShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateApplicationShrinkRequest
	GetVpcId() *string
	SetVpcNatGatewayId(v string) *CreateApplicationShrinkRequest
	GetVpcNatGatewayId() *string
	SetZoneId(v string) *CreateApplicationShrinkRequest
	GetZoneId() *string
}

type CreateApplicationShrinkRequest struct {
	// The ID of an existing template operator instance to associate. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// pm-xxxxxx
	AIDBClusterId *string `json:"AIDBClusterId,omitempty" xml:"AIDBClusterId,omitempty"`
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
	// The configuration of the authentication provider.
	//
	// example:
	//
	// xxx
	AuthProviderConfig *string `json:"AuthProviderConfig,omitempty" xml:"AuthProviderConfig,omitempty"`
	// Specifies whether to automatically create and associate with an elastic IP address (EIP).
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
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
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
	DnatEntriesShrink *string `json:"DnatEntries,omitempty" xml:"DnatEntries,omitempty"`
	// The dedicated DNAT NAT IP address that is allocated by the customer (separate from the SNAT IP address) for NAT mapping. The IP address must belong to the specified gateway and be in the available state. The vSwitch of the gateway must belong to the primary CIDR block that is reachable from the office network. Specify this parameter together with VpcNatGatewayId. Prerequisite: An SNAT entry is bound to the vSwitch where the application resides.
	//
	// example:
	//
	// 10.64.0.10
	DnatIpAddress *string `json:"DnatIpAddress,omitempty" xml:"DnatIpAddress,omitempty"`
	// Default value: `false`. If you set this parameter to `true`, only parameter and resource validation is performed without actually creating resources.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of user-defined service endpoints. By default, a VPC endpoint is created.
	EndpointsShrink *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// Required for knowledge applications.
	KnowledgeApplicationSpecShrink *string `json:"KnowledgeApplicationSpec,omitempty" xml:"KnowledgeApplicationSpec,omitempty"`
	// Required for mem0 applications.
	MemApplicationSpecShrink *string `json:"MemApplicationSpec,omitempty" xml:"MemApplicationSpec,omitempty"`
	// The API of the model. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// openai-completions
	ModelApi *string `json:"ModelApi,omitempty" xml:"ModelApi,omitempty"`
	// The API key of the model. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// sk-xxxxxx
	ModelApiKey *string `json:"ModelApiKey,omitempty" xml:"ModelApiKey,omitempty"`
	// The URL of the model. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	ModelBaseUrl *string `json:"ModelBaseUrl,omitempty" xml:"ModelBaseUrl,omitempty"`
	// The model source. Valid values:
	//
	// 	- bailian: Bailian model.
	//
	// 	- custom: Custom model.
	//
	// 	- maas: PolarDB model operator.
	//
	// example:
	//
	// bailian
	ModelFrom *string `json:"ModelFrom,omitempty" xml:"ModelFrom,omitempty"`
	// The name of the model. This parameter takes effect only when ApplicationType is set to polarclaw.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The list of parameters.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing type.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type, such as yearly or monthly.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The instance ID of the Polarlakebase cold storage or high-performance instance. Default value: empty. If specified, the corresponding storage is mounted to the application.
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
	// The IP whitelist. If you do not specify this parameter, the default value `127.0.0.1` is used.
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
	// The tags.
	Tag []*CreateApplicationShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The vSwitch. Default value: the current vSwitch in the primary zone of the instance.
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
	// The VPC NAT gateway ID for NAT mapping. If specified, NAT mapping is enabled when the instance is created. The NAT gateway must be in the same VPC as the application, use the private network type (intranet), and be in the active state.
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

func (s CreateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) GetAIDBClusterId() *string {
	return s.AIDBClusterId
}

func (s *CreateApplicationShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateApplicationShrinkRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateApplicationShrinkRequest) GetAuthProvider() *string {
	return s.AuthProvider
}

func (s *CreateApplicationShrinkRequest) GetAuthProviderConfig() *string {
	return s.AuthProviderConfig
}

func (s *CreateApplicationShrinkRequest) GetAutoAllocatePublicEip() *bool {
	return s.AutoAllocatePublicEip
}

func (s *CreateApplicationShrinkRequest) GetAutoCreatePolarFs() *bool {
	return s.AutoCreatePolarFs
}

func (s *CreateApplicationShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateApplicationShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreateApplicationShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *CreateApplicationShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApplicationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationShrinkRequest) GetDnatEntriesShrink() *string {
	return s.DnatEntriesShrink
}

func (s *CreateApplicationShrinkRequest) GetDnatIpAddress() *string {
	return s.DnatIpAddress
}

func (s *CreateApplicationShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateApplicationShrinkRequest) GetEndpointsShrink() *string {
	return s.EndpointsShrink
}

func (s *CreateApplicationShrinkRequest) GetKnowledgeApplicationSpecShrink() *string {
	return s.KnowledgeApplicationSpecShrink
}

func (s *CreateApplicationShrinkRequest) GetMemApplicationSpecShrink() *string {
	return s.MemApplicationSpecShrink
}

func (s *CreateApplicationShrinkRequest) GetModelApi() *string {
	return s.ModelApi
}

func (s *CreateApplicationShrinkRequest) GetModelApiKey() *string {
	return s.ModelApiKey
}

func (s *CreateApplicationShrinkRequest) GetModelBaseUrl() *string {
	return s.ModelBaseUrl
}

func (s *CreateApplicationShrinkRequest) GetModelFrom() *string {
	return s.ModelFrom
}

func (s *CreateApplicationShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateApplicationShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateApplicationShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateApplicationShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateApplicationShrinkRequest) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *CreateApplicationShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateApplicationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateApplicationShrinkRequest) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *CreateApplicationShrinkRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateApplicationShrinkRequest) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *CreateApplicationShrinkRequest) GetSkillTemplateId() *string {
	return s.SkillTemplateId
}

func (s *CreateApplicationShrinkRequest) GetTag() []*CreateApplicationShrinkRequestTag {
	return s.Tag
}

func (s *CreateApplicationShrinkRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateApplicationShrinkRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateApplicationShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateApplicationShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateApplicationShrinkRequest) GetVpcNatGatewayId() *string {
	return s.VpcNatGatewayId
}

func (s *CreateApplicationShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateApplicationShrinkRequest) SetAIDBClusterId(v string) *CreateApplicationShrinkRequest {
	s.AIDBClusterId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetApplicationType(v string) *CreateApplicationShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetArchitecture(v string) *CreateApplicationShrinkRequest {
	s.Architecture = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAuthProvider(v string) *CreateApplicationShrinkRequest {
	s.AuthProvider = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAuthProviderConfig(v string) *CreateApplicationShrinkRequest {
	s.AuthProviderConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAutoAllocatePublicEip(v bool) *CreateApplicationShrinkRequest {
	s.AutoAllocatePublicEip = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAutoCreatePolarFs(v bool) *CreateApplicationShrinkRequest {
	s.AutoCreatePolarFs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAutoRenew(v bool) *CreateApplicationShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAutoUseCoupon(v bool) *CreateApplicationShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetComponentsShrink(v string) *CreateApplicationShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDBClusterId(v string) *CreateApplicationShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDescription(v string) *CreateApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDnatEntriesShrink(v string) *CreateApplicationShrinkRequest {
	s.DnatEntriesShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDnatIpAddress(v string) *CreateApplicationShrinkRequest {
	s.DnatIpAddress = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDryRun(v bool) *CreateApplicationShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEndpointsShrink(v string) *CreateApplicationShrinkRequest {
	s.EndpointsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetKnowledgeApplicationSpecShrink(v string) *CreateApplicationShrinkRequest {
	s.KnowledgeApplicationSpecShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMemApplicationSpecShrink(v string) *CreateApplicationShrinkRequest {
	s.MemApplicationSpecShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetModelApi(v string) *CreateApplicationShrinkRequest {
	s.ModelApi = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetModelApiKey(v string) *CreateApplicationShrinkRequest {
	s.ModelApiKey = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetModelBaseUrl(v string) *CreateApplicationShrinkRequest {
	s.ModelBaseUrl = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetModelFrom(v string) *CreateApplicationShrinkRequest {
	s.ModelFrom = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetModelName(v string) *CreateApplicationShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetParametersShrink(v string) *CreateApplicationShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPayType(v string) *CreateApplicationShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPeriod(v string) *CreateApplicationShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPolarFSInstanceId(v string) *CreateApplicationShrinkRequest {
	s.PolarFSInstanceId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPromotionCode(v string) *CreateApplicationShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetRegionId(v string) *CreateApplicationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceGroupId(v string) *CreateApplicationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSecurityGroupId(v string) *CreateApplicationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSecurityIPArrayName(v string) *CreateApplicationShrinkRequest {
	s.SecurityIPArrayName = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSecurityIPList(v string) *CreateApplicationShrinkRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSecurityIPType(v string) *CreateApplicationShrinkRequest {
	s.SecurityIPType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSkillTemplateId(v string) *CreateApplicationShrinkRequest {
	s.SkillTemplateId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTag(v []*CreateApplicationShrinkRequestTag) *CreateApplicationShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTargetVersion(v string) *CreateApplicationShrinkRequest {
	s.TargetVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetUsedTime(v string) *CreateApplicationShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVSwitchId(v string) *CreateApplicationShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVpcId(v string) *CreateApplicationShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVpcNatGatewayId(v string) *CreateApplicationShrinkRequest {
	s.VpcNatGatewayId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetZoneId(v string) *CreateApplicationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) Validate() error {
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

type CreateApplicationShrinkRequestTag struct {
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

func (s CreateApplicationShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateApplicationShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateApplicationShrinkRequestTag) SetKey(v string) *CreateApplicationShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateApplicationShrinkRequestTag) SetValue(v string) *CreateApplicationShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateApplicationShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
