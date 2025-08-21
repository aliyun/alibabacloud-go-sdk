// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetResult(v *DescribeInstanceResponseBodyResult) *DescribeInstanceResponseBody
	GetResult() *DescribeInstanceResponseBodyResult
}

type DescribeInstanceResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetResult() *DescribeInstanceResponseBodyResult {
	return s.Result
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetResult(v *DescribeInstanceResponseBodyResult) *DescribeInstanceResponseBody {
	s.Result = v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResult struct {
	// example:
	//
	// true
	AdvancedDedicateMaster  *bool                                                      `json:"advancedDedicateMaster,omitempty" xml:"advancedDedicateMaster,omitempty"`
	AdvancedSetting         *DescribeInstanceResponseBodyResultAdvancedSetting         `json:"advancedSetting,omitempty" xml:"advancedSetting,omitempty" type:"Struct"`
	AliwsDicts              []*DescribeInstanceResponseBodyResultAliwsDicts            `json:"aliwsDicts,omitempty" xml:"aliwsDicts,omitempty" type:"Repeated"`
	ArchType                *string                                                    `json:"archType,omitempty" xml:"archType,omitempty"`
	ClientNodeConfiguration *DescribeInstanceResponseBodyResultClientNodeConfiguration `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// false
	DedicateMaster *bool `json:"dedicateMaster,omitempty" xml:"dedicateMaster,omitempty"`
	// example:
	//
	// es-cn-abc
	Description *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	DictList    []*DescribeInstanceResponseBodyResultDictList `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
	// example:
	//
	// es-cn-3h4k3axh33th9****.elasticsearch.aliyuncs.com
	Domain                       *string                                                         `json:"domain,omitempty" xml:"domain,omitempty"`
	ElasticDataNodeConfiguration *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// false
	EnableKibanaPrivateNetwork *bool `json:"enableKibanaPrivateNetwork,omitempty" xml:"enableKibanaPrivateNetwork,omitempty"`
	// example:
	//
	// true
	EnableKibanaPublicNetwork *bool `json:"enableKibanaPublicNetwork,omitempty" xml:"enableKibanaPublicNetwork,omitempty"`
	// example:
	//
	// true
	EnablePublic *bool  `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	EndTime      *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// {"http.cors.allow-credentials":"false"}
	EsConfig      map[string]interface{} `json:"esConfig,omitempty" xml:"esConfig,omitempty"`
	EsIPBlacklist []*string              `json:"esIPBlacklist,omitempty" xml:"esIPBlacklist,omitempty" type:"Repeated"`
	EsIPWhitelist []*string              `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	// example:
	//
	// 6.3.2_with_X-Pack
	EsVersion     *string                  `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	ExtendConfigs []map[string]interface{} `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HaveClientNode *bool `json:"haveClientNode,omitempty" xml:"haveClientNode,omitempty"`
	// example:
	//
	// true
	HaveKibana *bool                                           `json:"haveKibana,omitempty" xml:"haveKibana,omitempty"`
	IkHotDicts []*DescribeInstanceResponseBodyResultIkHotDicts `json:"ikHotDicts,omitempty" xml:"ikHotDicts,omitempty" type:"Repeated"`
	// example:
	//
	// advanced
	InstanceCategory *string `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	// example:
	//
	// es-cn-3h4k3axh33th9****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	IsNewDeployment     *bool                                                  `json:"isNewDeployment,omitempty" xml:"isNewDeployment,omitempty"`
	KibanaConfiguration *DescribeInstanceResponseBodyResultKibanaConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// es-cn-3h4k3axh33th9****.kibana.elasticsearch.aliyuncs.com
	KibanaDomain      *string   `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
	KibanaIPWhitelist []*string `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	// example:
	//
	// 5601
	KibanaPort               *int32                                                 `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
	KibanaPrivateDomain      *string                                                `json:"kibanaPrivateDomain,omitempty" xml:"kibanaPrivateDomain,omitempty"`
	KibanaPrivateIPWhitelist []*string                                              `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
	KibanaPrivatePort        *string                                                `json:"kibanaPrivatePort,omitempty" xml:"kibanaPrivatePort,omitempty"`
	MasterConfiguration      *DescribeInstanceResponseBodyResultMasterConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty" type:"Struct"`
	NetworkConfig            *DescribeInstanceResponseBodyResultNetworkConfig       `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2
	NodeAmount *int32                                      `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec   *DescribeInstanceResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// active
	PostpaidServiceStatus     *string   `json:"postpaidServiceStatus,omitempty" xml:"postpaidServiceStatus,omitempty"`
	PrivateNetworkIpWhiteList []*string `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// es-cn-3h4k3axh33th9****.elasticsearch.aliyuncs.com
	PublicDomain      *string   `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
	PublicIpWhitelist []*string `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
	// example:
	//
	// 9200
	PublicPort *int32 `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
	// example:
	//
	// rg-aekzvowej3i****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// true
	ServiceVpc *bool `json:"serviceVpc,omitempty" xml:"serviceVpc,omitempty"`
	// example:
	//
	// active
	Status        *string                                            `json:"status,omitempty" xml:"status,omitempty"`
	SynonymsDicts []*DescribeInstanceResponseBodyResultSynonymsDicts `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
	Tags          []*DescribeInstanceResponseBodyResultTags          `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// vpc-bp1uag5jj38c****
	VpcInstanceId *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
	// example:
	//
	// true
	WarmNode              *bool                                                    `json:"warmNode,omitempty" xml:"warmNode,omitempty"`
	WarmNodeConfiguration *DescribeInstanceResponseBodyResultWarmNodeConfiguration `json:"warmNodeConfiguration,omitempty" xml:"warmNodeConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 2
	ZoneCount *int32                                         `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
	ZoneInfos []*DescribeInstanceResponseBodyResultZoneInfos `json:"zoneInfos,omitempty" xml:"zoneInfos,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResult) GetAdvancedDedicateMaster() *bool {
	return s.AdvancedDedicateMaster
}

func (s *DescribeInstanceResponseBodyResult) GetAdvancedSetting() *DescribeInstanceResponseBodyResultAdvancedSetting {
	return s.AdvancedSetting
}

func (s *DescribeInstanceResponseBodyResult) GetAliwsDicts() []*DescribeInstanceResponseBodyResultAliwsDicts {
	return s.AliwsDicts
}

func (s *DescribeInstanceResponseBodyResult) GetArchType() *string {
	return s.ArchType
}

func (s *DescribeInstanceResponseBodyResult) GetClientNodeConfiguration() *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	return s.ClientNodeConfiguration
}

func (s *DescribeInstanceResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeInstanceResponseBodyResult) GetDedicateMaster() *bool {
	return s.DedicateMaster
}

func (s *DescribeInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceResponseBodyResult) GetDictList() []*DescribeInstanceResponseBodyResultDictList {
	return s.DictList
}

func (s *DescribeInstanceResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *DescribeInstanceResponseBodyResult) GetElasticDataNodeConfiguration() *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	return s.ElasticDataNodeConfiguration
}

func (s *DescribeInstanceResponseBodyResult) GetEnableKibanaPrivateNetwork() *bool {
	return s.EnableKibanaPrivateNetwork
}

func (s *DescribeInstanceResponseBodyResult) GetEnableKibanaPublicNetwork() *bool {
	return s.EnableKibanaPublicNetwork
}

func (s *DescribeInstanceResponseBodyResult) GetEnablePublic() *bool {
	return s.EnablePublic
}

func (s *DescribeInstanceResponseBodyResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeInstanceResponseBodyResult) GetEsConfig() map[string]interface{} {
	return s.EsConfig
}

func (s *DescribeInstanceResponseBodyResult) GetEsIPBlacklist() []*string {
	return s.EsIPBlacklist
}

func (s *DescribeInstanceResponseBodyResult) GetEsIPWhitelist() []*string {
	return s.EsIPWhitelist
}

func (s *DescribeInstanceResponseBodyResult) GetEsVersion() *string {
	return s.EsVersion
}

func (s *DescribeInstanceResponseBodyResult) GetExtendConfigs() []map[string]interface{} {
	return s.ExtendConfigs
}

func (s *DescribeInstanceResponseBodyResult) GetHaveClientNode() *bool {
	return s.HaveClientNode
}

func (s *DescribeInstanceResponseBodyResult) GetHaveKibana() *bool {
	return s.HaveKibana
}

func (s *DescribeInstanceResponseBodyResult) GetIkHotDicts() []*DescribeInstanceResponseBodyResultIkHotDicts {
	return s.IkHotDicts
}

func (s *DescribeInstanceResponseBodyResult) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *DescribeInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBodyResult) GetIsNewDeployment() *bool {
	return s.IsNewDeployment
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaConfiguration() *DescribeInstanceResponseBodyResultKibanaConfiguration {
	return s.KibanaConfiguration
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaDomain() *string {
	return s.KibanaDomain
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaIPWhitelist() []*string {
	return s.KibanaIPWhitelist
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaPort() *int32 {
	return s.KibanaPort
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaPrivateDomain() *string {
	return s.KibanaPrivateDomain
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaPrivateIPWhitelist() []*string {
	return s.KibanaPrivateIPWhitelist
}

func (s *DescribeInstanceResponseBodyResult) GetKibanaPrivatePort() *string {
	return s.KibanaPrivatePort
}

func (s *DescribeInstanceResponseBodyResult) GetMasterConfiguration() *DescribeInstanceResponseBodyResultMasterConfiguration {
	return s.MasterConfiguration
}

func (s *DescribeInstanceResponseBodyResult) GetNetworkConfig() *DescribeInstanceResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *DescribeInstanceResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *DescribeInstanceResponseBodyResult) GetNodeSpec() *DescribeInstanceResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *DescribeInstanceResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeInstanceResponseBodyResult) GetPort() *int32 {
	return s.Port
}

func (s *DescribeInstanceResponseBodyResult) GetPostpaidServiceStatus() *string {
	return s.PostpaidServiceStatus
}

func (s *DescribeInstanceResponseBodyResult) GetPrivateNetworkIpWhiteList() []*string {
	return s.PrivateNetworkIpWhiteList
}

func (s *DescribeInstanceResponseBodyResult) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeInstanceResponseBodyResult) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *DescribeInstanceResponseBodyResult) GetPublicIpWhitelist() []*string {
	return s.PublicIpWhitelist
}

func (s *DescribeInstanceResponseBodyResult) GetPublicPort() *int32 {
	return s.PublicPort
}

func (s *DescribeInstanceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceResponseBodyResult) GetServiceVpc() *bool {
	return s.ServiceVpc
}

func (s *DescribeInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceResponseBodyResult) GetSynonymsDicts() []*DescribeInstanceResponseBodyResultSynonymsDicts {
	return s.SynonymsDicts
}

func (s *DescribeInstanceResponseBodyResult) GetTags() []*DescribeInstanceResponseBodyResultTags {
	return s.Tags
}

func (s *DescribeInstanceResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeInstanceResponseBodyResult) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeInstanceResponseBodyResult) GetWarmNode() *bool {
	return s.WarmNode
}

func (s *DescribeInstanceResponseBodyResult) GetWarmNodeConfiguration() *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	return s.WarmNodeConfiguration
}

func (s *DescribeInstanceResponseBodyResult) GetZoneCount() *int32 {
	return s.ZoneCount
}

func (s *DescribeInstanceResponseBodyResult) GetZoneInfos() []*DescribeInstanceResponseBodyResultZoneInfos {
	return s.ZoneInfos
}

func (s *DescribeInstanceResponseBodyResult) SetAdvancedDedicateMaster(v bool) *DescribeInstanceResponseBodyResult {
	s.AdvancedDedicateMaster = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetAdvancedSetting(v *DescribeInstanceResponseBodyResultAdvancedSetting) *DescribeInstanceResponseBodyResult {
	s.AdvancedSetting = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetAliwsDicts(v []*DescribeInstanceResponseBodyResultAliwsDicts) *DescribeInstanceResponseBodyResult {
	s.AliwsDicts = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetArchType(v string) *DescribeInstanceResponseBodyResult {
	s.ArchType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetClientNodeConfiguration(v *DescribeInstanceResponseBodyResultClientNodeConfiguration) *DescribeInstanceResponseBodyResult {
	s.ClientNodeConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetCreatedAt(v string) *DescribeInstanceResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDedicateMaster(v bool) *DescribeInstanceResponseBodyResult {
	s.DedicateMaster = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDescription(v string) *DescribeInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDictList(v []*DescribeInstanceResponseBodyResultDictList) *DescribeInstanceResponseBodyResult {
	s.DictList = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetDomain(v string) *DescribeInstanceResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetElasticDataNodeConfiguration(v *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) *DescribeInstanceResponseBodyResult {
	s.ElasticDataNodeConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEnableKibanaPrivateNetwork(v bool) *DescribeInstanceResponseBodyResult {
	s.EnableKibanaPrivateNetwork = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEnableKibanaPublicNetwork(v bool) *DescribeInstanceResponseBodyResult {
	s.EnableKibanaPublicNetwork = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEnablePublic(v bool) *DescribeInstanceResponseBodyResult {
	s.EnablePublic = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEndTime(v int64) *DescribeInstanceResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsConfig(v map[string]interface{}) *DescribeInstanceResponseBodyResult {
	s.EsConfig = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsIPBlacklist(v []*string) *DescribeInstanceResponseBodyResult {
	s.EsIPBlacklist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsIPWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.EsIPWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetEsVersion(v string) *DescribeInstanceResponseBodyResult {
	s.EsVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetExtendConfigs(v []map[string]interface{}) *DescribeInstanceResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetHaveClientNode(v bool) *DescribeInstanceResponseBodyResult {
	s.HaveClientNode = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetHaveKibana(v bool) *DescribeInstanceResponseBodyResult {
	s.HaveKibana = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetIkHotDicts(v []*DescribeInstanceResponseBodyResultIkHotDicts) *DescribeInstanceResponseBodyResult {
	s.IkHotDicts = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetInstanceCategory(v string) *DescribeInstanceResponseBodyResult {
	s.InstanceCategory = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetInstanceId(v string) *DescribeInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetIsNewDeployment(v bool) *DescribeInstanceResponseBodyResult {
	s.IsNewDeployment = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaConfiguration(v *DescribeInstanceResponseBodyResultKibanaConfiguration) *DescribeInstanceResponseBodyResult {
	s.KibanaConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaDomain(v string) *DescribeInstanceResponseBodyResult {
	s.KibanaDomain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaIPWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.KibanaIPWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaPort(v int32) *DescribeInstanceResponseBodyResult {
	s.KibanaPort = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaPrivateDomain(v string) *DescribeInstanceResponseBodyResult {
	s.KibanaPrivateDomain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaPrivateIPWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.KibanaPrivateIPWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetKibanaPrivatePort(v string) *DescribeInstanceResponseBodyResult {
	s.KibanaPrivatePort = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetMasterConfiguration(v *DescribeInstanceResponseBodyResultMasterConfiguration) *DescribeInstanceResponseBodyResult {
	s.MasterConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetNetworkConfig(v *DescribeInstanceResponseBodyResultNetworkConfig) *DescribeInstanceResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetNodeAmount(v int32) *DescribeInstanceResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetNodeSpec(v *DescribeInstanceResponseBodyResultNodeSpec) *DescribeInstanceResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPaymentType(v string) *DescribeInstanceResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPort(v int32) *DescribeInstanceResponseBodyResult {
	s.Port = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPostpaidServiceStatus(v string) *DescribeInstanceResponseBodyResult {
	s.PostpaidServiceStatus = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPrivateNetworkIpWhiteList(v []*string) *DescribeInstanceResponseBodyResult {
	s.PrivateNetworkIpWhiteList = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetProtocol(v string) *DescribeInstanceResponseBodyResult {
	s.Protocol = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPublicDomain(v string) *DescribeInstanceResponseBodyResult {
	s.PublicDomain = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPublicIpWhitelist(v []*string) *DescribeInstanceResponseBodyResult {
	s.PublicIpWhitelist = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetPublicPort(v int32) *DescribeInstanceResponseBodyResult {
	s.PublicPort = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetResourceGroupId(v string) *DescribeInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetServiceVpc(v bool) *DescribeInstanceResponseBodyResult {
	s.ServiceVpc = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetStatus(v string) *DescribeInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetSynonymsDicts(v []*DescribeInstanceResponseBodyResultSynonymsDicts) *DescribeInstanceResponseBodyResult {
	s.SynonymsDicts = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetTags(v []*DescribeInstanceResponseBodyResultTags) *DescribeInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetUpdatedAt(v string) *DescribeInstanceResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetVpcInstanceId(v string) *DescribeInstanceResponseBodyResult {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetWarmNode(v bool) *DescribeInstanceResponseBodyResult {
	s.WarmNode = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetWarmNodeConfiguration(v *DescribeInstanceResponseBodyResultWarmNodeConfiguration) *DescribeInstanceResponseBodyResult {
	s.WarmNodeConfiguration = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetZoneCount(v int32) *DescribeInstanceResponseBodyResult {
	s.ZoneCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResult) SetZoneInfos(v []*DescribeInstanceResponseBodyResultZoneInfos) *DescribeInstanceResponseBodyResult {
	s.ZoneInfos = v
	return s
}

func (s *DescribeInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultAdvancedSetting struct {
	// example:
	//
	// CMS
	GcName *string `json:"gcName,omitempty" xml:"gcName,omitempty"`
}

func (s DescribeInstanceResponseBodyResultAdvancedSetting) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultAdvancedSetting) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultAdvancedSetting) GetGcName() *string {
	return s.GcName
}

func (s *DescribeInstanceResponseBodyResultAdvancedSetting) SetGcName(v string) *DescribeInstanceResponseBodyResultAdvancedSetting {
	s.GcName = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAdvancedSetting) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultAliwsDicts struct {
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// aliws_ext_dict.txt
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// ALI_WS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultAliwsDicts) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultAliwsDicts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetFileSize(v int64) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetName(v string) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetSourceType(v string) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) SetType(v string) *DescribeInstanceResponseBodyResultAliwsDicts {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultAliwsDicts) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultClientNodeConfiguration struct {
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 40
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s DescribeInstanceResponseBodyResultClientNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultClientNodeConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) SetSpecInfo(v string) *DescribeInstanceResponseBodyResultClientNodeConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultClientNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultDictList struct {
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultDictList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultDictList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultDictList) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeInstanceResponseBodyResultDictList) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyResultDictList) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeInstanceResponseBodyResultDictList) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBodyResultDictList) SetFileSize(v int64) *DescribeInstanceResponseBodyResultDictList {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) SetName(v string) *DescribeInstanceResponseBodyResultDictList {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) SetSourceType(v string) *DescribeInstanceResponseBodyResultDictList {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) SetType(v string) *DescribeInstanceResponseBodyResultDictList {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultDictList) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultElasticDataNodeConfiguration struct {
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskEncryption(v bool) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) SetSpecInfo(v string) *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultElasticDataNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultIkHotDicts struct {
	// example:
	//
	// 6
	FileSize *int32 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// deploy_0.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultIkHotDicts) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultIkHotDicts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) GetFileSize() *int32 {
	return s.FileSize
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) SetFileSize(v int32) *DescribeInstanceResponseBodyResultIkHotDicts {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) SetName(v string) *DescribeInstanceResponseBodyResultIkHotDicts {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) SetSourceType(v string) *DescribeInstanceResponseBodyResultIkHotDicts {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) SetType(v string) *DescribeInstanceResponseBodyResultIkHotDicts {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultIkHotDicts) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultKibanaConfiguration struct {
	// example:
	//
	// 1
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s DescribeInstanceResponseBodyResultKibanaConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultKibanaConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultKibanaConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultKibanaConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultKibanaConfiguration {
	s.Spec = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) SetSpecInfo(v string) *DescribeInstanceResponseBodyResultKibanaConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultKibanaConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultMasterConfiguration struct {
	// example:
	//
	// 3
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 40
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s DescribeInstanceResponseBodyResultMasterConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultMasterConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.Spec = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) SetSpecInfo(v string) *DescribeInstanceResponseBodyResultMasterConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultMasterConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultNetworkConfig struct {
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-abc
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// example:
	//
	// vsw-abc
	VswitchId        *string                                                            `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	WhiteIpGroupList []*DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) GetWhiteIpGroupList() []*DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	return s.WhiteIpGroupList
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetType(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetVpcId(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetVsArea(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetVswitchId(v string) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) SetWhiteIpGroupList(v []*DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) *DescribeInstanceResponseBodyResultNetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	// example:
	//
	// default
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// example:
	//
	// PRIVATE_ES
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GetIps() []*string {
	return s.Ips
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetGroupName(v string) *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetIps(v []*string) *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.Ips = v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) SetWhiteIpType(v string) *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.WhiteIpType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNetworkConfigWhiteIpGroupList) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultNodeSpec struct {
	// example:
	//
	// 0
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s DescribeInstanceResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetDisk(v int32) *DescribeInstanceResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *DescribeInstanceResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetDiskType(v string) *DescribeInstanceResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetPerformanceLevel(v string) *DescribeInstanceResponseBodyResultNodeSpec {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetSpec(v string) *DescribeInstanceResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) SetSpecInfo(v string) *DescribeInstanceResponseBodyResultNodeSpec {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultSynonymsDicts struct {
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// SYSTEM_MAIN.dic
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ORIGIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// STOP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeInstanceResponseBodyResultSynonymsDicts) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultSynonymsDicts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetFileSize(v int64) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.FileSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetName(v string) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetSourceType(v string) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.SourceType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) SetType(v string) *DescribeInstanceResponseBodyResultSynonymsDicts {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultSynonymsDicts) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultTags struct {
	// example:
	//
	// env
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// example:
	//
	// dev
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s DescribeInstanceResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeInstanceResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeInstanceResponseBodyResultTags) SetTagKey(v string) *DescribeInstanceResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultTags) SetTagValue(v string) *DescribeInstanceResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultWarmNodeConfiguration struct {
	// example:
	//
	// 6
	Amount *int32 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 500
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// elasticsearch.n4.small
	Spec     *string `json:"spec,omitempty" xml:"spec,omitempty"`
	SpecInfo *string `json:"specInfo,omitempty" xml:"specInfo,omitempty"`
}

func (s DescribeInstanceResponseBodyResultWarmNodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultWarmNodeConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) GetSpec() *string {
	return s.Spec
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) GetSpecInfo() *string {
	return s.SpecInfo
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetAmount(v int32) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetDisk(v int32) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.Disk = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetDiskEncryption(v bool) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetDiskType(v string) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetSpec(v string) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.Spec = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) SetSpecInfo(v string) *DescribeInstanceResponseBodyResultWarmNodeConfiguration {
	s.SpecInfo = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultWarmNodeConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyResultZoneInfos struct {
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeInstanceResponseBodyResultZoneInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyResultZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) SetStatus(v string) *DescribeInstanceResponseBodyResultZoneInfos {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) SetZoneId(v string) *DescribeInstanceResponseBodyResultZoneInfos {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceResponseBodyResultZoneInfos) Validate() error {
	return dara.Validate(s)
}
