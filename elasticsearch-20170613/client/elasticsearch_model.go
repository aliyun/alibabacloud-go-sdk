// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElasticsearch interface {
  dara.Model
  String() string
  GoString() string
  SetAdvancedDedicateMaster(v bool) *Elasticsearch
  GetAdvancedDedicateMaster() *bool 
  SetAdvancedSetting(v *ElasticsearchAdvancedSetting) *Elasticsearch
  GetAdvancedSetting() *ElasticsearchAdvancedSetting 
  SetAliwsDicts(v []*DictInfo) *Elasticsearch
  GetAliwsDicts() []*DictInfo 
  SetClientNodeConfiguration(v *ClientNodeConfiguration) *Elasticsearch
  GetClientNodeConfiguration() *ClientNodeConfiguration 
  SetCreatedAt(v string) *Elasticsearch
  GetCreatedAt() *string 
  SetDataNode(v bool) *Elasticsearch
  GetDataNode() *bool 
  SetDedicateMaster(v bool) *Elasticsearch
  GetDedicateMaster() *bool 
  SetDescription(v string) *Elasticsearch
  GetDescription() *string 
  SetDictList(v []*DictInfo) *Elasticsearch
  GetDictList() []*DictInfo 
  SetDomain(v string) *Elasticsearch
  GetDomain() *string 
  SetElasticDataNodeConfiguration(v *ElasticDataNodeConfiguration) *Elasticsearch
  GetElasticDataNodeConfiguration() *ElasticDataNodeConfiguration 
  SetEnableKibanaPrivateNetwork(v bool) *Elasticsearch
  GetEnableKibanaPrivateNetwork() *bool 
  SetEnableKibanaPublicNetwork(v bool) *Elasticsearch
  GetEnableKibanaPublicNetwork() *bool 
  SetEnablePublic(v bool) *Elasticsearch
  GetEnablePublic() *bool 
  SetEndTime(v int64) *Elasticsearch
  GetEndTime() *int64 
  SetEsConfig(v map[string]*string) *Elasticsearch
  GetEsConfig() map[string]*string 
  SetEsIPWhitelist(v []*string) *Elasticsearch
  GetEsIPWhitelist() []*string 
  SetEsVersion(v string) *Elasticsearch
  GetEsVersion() *string 
  SetExtendConfigs(v []map[string]interface{}) *Elasticsearch
  GetExtendConfigs() []map[string]interface{} 
  SetHaveClientNode(v bool) *Elasticsearch
  GetHaveClientNode() *bool 
  SetHaveElasticDataNode(v bool) *Elasticsearch
  GetHaveElasticDataNode() *bool 
  SetHaveKibana(v bool) *Elasticsearch
  GetHaveKibana() *bool 
  SetIkHotDicts(v []*DictInfo) *Elasticsearch
  GetIkHotDicts() []*DictInfo 
  SetInstanceId(v string) *Elasticsearch
  GetInstanceId() *string 
  SetKibanaConfiguration(v *KibanaNodeConfiguration) *Elasticsearch
  GetKibanaConfiguration() *KibanaNodeConfiguration 
  SetKibanaDomain(v string) *Elasticsearch
  GetKibanaDomain() *string 
  SetKibanaIPWhitelist(v []*string) *Elasticsearch
  GetKibanaIPWhitelist() []*string 
  SetKibanaPort(v int64) *Elasticsearch
  GetKibanaPort() *int64 
  SetKibanaPrivateDomain(v string) *Elasticsearch
  GetKibanaPrivateDomain() *string 
  SetKibanaPrivateIPWhitelist(v []*string) *Elasticsearch
  GetKibanaPrivateIPWhitelist() []*string 
  SetKibanaPrivatePort(v int64) *Elasticsearch
  GetKibanaPrivatePort() *int64 
  SetKibanaProtocol(v string) *Elasticsearch
  GetKibanaProtocol() *string 
  SetMasterConfiguration(v *MasterNodeConfiguration) *Elasticsearch
  GetMasterConfiguration() *MasterNodeConfiguration 
  SetNetworkConfig(v *NetworkConfig) *Elasticsearch
  GetNetworkConfig() *NetworkConfig 
  SetNodeAmount(v int64) *Elasticsearch
  GetNodeAmount() *int64 
  SetNodeSpec(v *NodeSpec) *Elasticsearch
  GetNodeSpec() *NodeSpec 
  SetPaymentType(v string) *Elasticsearch
  GetPaymentType() *string 
  SetPort(v int64) *Elasticsearch
  GetPort() *int64 
  SetPrivateNetworkIpWhiteList(v []*string) *Elasticsearch
  GetPrivateNetworkIpWhiteList() []*string 
  SetProductType(v string) *Elasticsearch
  GetProductType() *string 
  SetProtocol(v string) *Elasticsearch
  GetProtocol() *string 
  SetPublicDomain(v string) *Elasticsearch
  GetPublicDomain() *string 
  SetPublicIpWhitelist(v []*string) *Elasticsearch
  GetPublicIpWhitelist() []*string 
  SetPublicPort(v int64) *Elasticsearch
  GetPublicPort() *int64 
  SetReadWritePolicy(v *ReadWritePolicy) *Elasticsearch
  GetReadWritePolicy() *ReadWritePolicy 
  SetResourceGroupId(v string) *Elasticsearch
  GetResourceGroupId() *string 
  SetServiceVpc(v bool) *Elasticsearch
  GetServiceVpc() *bool 
  SetStatus(v string) *Elasticsearch
  GetStatus() *string 
  SetSynonymsDicts(v []*DictInfo) *Elasticsearch
  GetSynonymsDicts() []*DictInfo 
  SetTags(v []*Tag) *Elasticsearch
  GetTags() []*Tag 
  SetUpdatedAt(v string) *Elasticsearch
  GetUpdatedAt() *string 
  SetWarmNode(v bool) *Elasticsearch
  GetWarmNode() *bool 
  SetWarmNodeConfiguration(v *WarmNodeConfiguration) *Elasticsearch
  GetWarmNodeConfiguration() *WarmNodeConfiguration 
  SetZoneCount(v int64) *Elasticsearch
  GetZoneCount() *int64 
  SetZoneInfos(v []*ZoneInfo) *Elasticsearch
  GetZoneInfos() []*ZoneInfo 
}

type Elasticsearch struct {
  AdvancedDedicateMaster *bool `json:"advancedDedicateMaster,omitempty" xml:"advancedDedicateMaster,omitempty"`
  AdvancedSetting *ElasticsearchAdvancedSetting `json:"advancedSetting,omitempty" xml:"advancedSetting,omitempty" type:"Struct"`
  AliwsDicts []*DictInfo `json:"aliwsDicts,omitempty" xml:"aliwsDicts,omitempty" type:"Repeated"`
  ClientNodeConfiguration *ClientNodeConfiguration `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty"`
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  DataNode *bool `json:"dataNode,omitempty" xml:"dataNode,omitempty"`
  DedicateMaster *bool `json:"dedicateMaster,omitempty" xml:"dedicateMaster,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  DictList []*DictInfo `json:"dictList,omitempty" xml:"dictList,omitempty" type:"Repeated"`
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
  ElasticDataNodeConfiguration *ElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty"`
  EnableKibanaPrivateNetwork *bool `json:"enableKibanaPrivateNetwork,omitempty" xml:"enableKibanaPrivateNetwork,omitempty"`
  EnableKibanaPublicNetwork *bool `json:"enableKibanaPublicNetwork,omitempty" xml:"enableKibanaPublicNetwork,omitempty"`
  EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
  EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
  EsConfig map[string]*string `json:"esConfig,omitempty" xml:"esConfig,omitempty"`
  EsIPWhitelist []*string `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
  EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
  ExtendConfigs []map[string]interface{} `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
  HaveClientNode *bool `json:"haveClientNode,omitempty" xml:"haveClientNode,omitempty"`
  HaveElasticDataNode *bool `json:"haveElasticDataNode,omitempty" xml:"haveElasticDataNode,omitempty"`
  HaveKibana *bool `json:"haveKibana,omitempty" xml:"haveKibana,omitempty"`
  IkHotDicts []*DictInfo `json:"ikHotDicts,omitempty" xml:"ikHotDicts,omitempty" type:"Repeated"`
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
  KibanaConfiguration *KibanaNodeConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty"`
  KibanaDomain *string `json:"kibanaDomain,omitempty" xml:"kibanaDomain,omitempty"`
  KibanaIPWhitelist []*string `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
  KibanaPort *int64 `json:"kibanaPort,omitempty" xml:"kibanaPort,omitempty"`
  KibanaPrivateDomain *string `json:"kibanaPrivateDomain,omitempty" xml:"kibanaPrivateDomain,omitempty"`
  KibanaPrivateIPWhitelist []*string `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
  KibanaPrivatePort *int64 `json:"kibanaPrivatePort,omitempty" xml:"kibanaPrivatePort,omitempty"`
  KibanaProtocol *string `json:"kibanaProtocol,omitempty" xml:"kibanaProtocol,omitempty"`
  MasterConfiguration *MasterNodeConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty"`
  NetworkConfig *NetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty"`
  NodeAmount *int64 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
  NodeSpec *NodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty"`
  PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
  Port *int64 `json:"port,omitempty" xml:"port,omitempty"`
  PrivateNetworkIpWhiteList []*string `json:"privateNetworkIpWhiteList,omitempty" xml:"privateNetworkIpWhiteList,omitempty" type:"Repeated"`
  ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  PublicDomain *string `json:"publicDomain,omitempty" xml:"publicDomain,omitempty"`
  PublicIpWhitelist []*string `json:"publicIpWhitelist,omitempty" xml:"publicIpWhitelist,omitempty" type:"Repeated"`
  PublicPort *int64 `json:"publicPort,omitempty" xml:"publicPort,omitempty"`
  ReadWritePolicy *ReadWritePolicy `json:"readWritePolicy,omitempty" xml:"readWritePolicy,omitempty"`
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
  ServiceVpc *bool `json:"serviceVpc,omitempty" xml:"serviceVpc,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  SynonymsDicts []*DictInfo `json:"synonymsDicts,omitempty" xml:"synonymsDicts,omitempty" type:"Repeated"`
  Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
  UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
  WarmNode *bool `json:"warmNode,omitempty" xml:"warmNode,omitempty"`
  WarmNodeConfiguration *WarmNodeConfiguration `json:"warmNodeConfiguration,omitempty" xml:"warmNodeConfiguration,omitempty"`
  ZoneCount *int64 `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
  ZoneInfos []*ZoneInfo `json:"zoneInfos,omitempty" xml:"zoneInfos,omitempty" type:"Repeated"`
}

func (s Elasticsearch) String() string {
  return dara.Prettify(s)
}

func (s Elasticsearch) GoString() string {
  return s.String()
}

func (s *Elasticsearch) GetAdvancedDedicateMaster() *bool  {
  return s.AdvancedDedicateMaster
}

func (s *Elasticsearch) GetAdvancedSetting() *ElasticsearchAdvancedSetting  {
  return s.AdvancedSetting
}

func (s *Elasticsearch) GetAliwsDicts() []*DictInfo  {
  return s.AliwsDicts
}

func (s *Elasticsearch) GetClientNodeConfiguration() *ClientNodeConfiguration  {
  return s.ClientNodeConfiguration
}

func (s *Elasticsearch) GetCreatedAt() *string  {
  return s.CreatedAt
}

func (s *Elasticsearch) GetDataNode() *bool  {
  return s.DataNode
}

func (s *Elasticsearch) GetDedicateMaster() *bool  {
  return s.DedicateMaster
}

func (s *Elasticsearch) GetDescription() *string  {
  return s.Description
}

func (s *Elasticsearch) GetDictList() []*DictInfo  {
  return s.DictList
}

func (s *Elasticsearch) GetDomain() *string  {
  return s.Domain
}

func (s *Elasticsearch) GetElasticDataNodeConfiguration() *ElasticDataNodeConfiguration  {
  return s.ElasticDataNodeConfiguration
}

func (s *Elasticsearch) GetEnableKibanaPrivateNetwork() *bool  {
  return s.EnableKibanaPrivateNetwork
}

func (s *Elasticsearch) GetEnableKibanaPublicNetwork() *bool  {
  return s.EnableKibanaPublicNetwork
}

func (s *Elasticsearch) GetEnablePublic() *bool  {
  return s.EnablePublic
}

func (s *Elasticsearch) GetEndTime() *int64  {
  return s.EndTime
}

func (s *Elasticsearch) GetEsConfig() map[string]*string  {
  return s.EsConfig
}

func (s *Elasticsearch) GetEsIPWhitelist() []*string  {
  return s.EsIPWhitelist
}

func (s *Elasticsearch) GetEsVersion() *string  {
  return s.EsVersion
}

func (s *Elasticsearch) GetExtendConfigs() []map[string]interface{}  {
  return s.ExtendConfigs
}

func (s *Elasticsearch) GetHaveClientNode() *bool  {
  return s.HaveClientNode
}

func (s *Elasticsearch) GetHaveElasticDataNode() *bool  {
  return s.HaveElasticDataNode
}

func (s *Elasticsearch) GetHaveKibana() *bool  {
  return s.HaveKibana
}

func (s *Elasticsearch) GetIkHotDicts() []*DictInfo  {
  return s.IkHotDicts
}

func (s *Elasticsearch) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *Elasticsearch) GetKibanaConfiguration() *KibanaNodeConfiguration  {
  return s.KibanaConfiguration
}

func (s *Elasticsearch) GetKibanaDomain() *string  {
  return s.KibanaDomain
}

func (s *Elasticsearch) GetKibanaIPWhitelist() []*string  {
  return s.KibanaIPWhitelist
}

func (s *Elasticsearch) GetKibanaPort() *int64  {
  return s.KibanaPort
}

func (s *Elasticsearch) GetKibanaPrivateDomain() *string  {
  return s.KibanaPrivateDomain
}

func (s *Elasticsearch) GetKibanaPrivateIPWhitelist() []*string  {
  return s.KibanaPrivateIPWhitelist
}

func (s *Elasticsearch) GetKibanaPrivatePort() *int64  {
  return s.KibanaPrivatePort
}

func (s *Elasticsearch) GetKibanaProtocol() *string  {
  return s.KibanaProtocol
}

func (s *Elasticsearch) GetMasterConfiguration() *MasterNodeConfiguration  {
  return s.MasterConfiguration
}

func (s *Elasticsearch) GetNetworkConfig() *NetworkConfig  {
  return s.NetworkConfig
}

func (s *Elasticsearch) GetNodeAmount() *int64  {
  return s.NodeAmount
}

func (s *Elasticsearch) GetNodeSpec() *NodeSpec  {
  return s.NodeSpec
}

func (s *Elasticsearch) GetPaymentType() *string  {
  return s.PaymentType
}

func (s *Elasticsearch) GetPort() *int64  {
  return s.Port
}

func (s *Elasticsearch) GetPrivateNetworkIpWhiteList() []*string  {
  return s.PrivateNetworkIpWhiteList
}

func (s *Elasticsearch) GetProductType() *string  {
  return s.ProductType
}

func (s *Elasticsearch) GetProtocol() *string  {
  return s.Protocol
}

func (s *Elasticsearch) GetPublicDomain() *string  {
  return s.PublicDomain
}

func (s *Elasticsearch) GetPublicIpWhitelist() []*string  {
  return s.PublicIpWhitelist
}

func (s *Elasticsearch) GetPublicPort() *int64  {
  return s.PublicPort
}

func (s *Elasticsearch) GetReadWritePolicy() *ReadWritePolicy  {
  return s.ReadWritePolicy
}

func (s *Elasticsearch) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *Elasticsearch) GetServiceVpc() *bool  {
  return s.ServiceVpc
}

func (s *Elasticsearch) GetStatus() *string  {
  return s.Status
}

func (s *Elasticsearch) GetSynonymsDicts() []*DictInfo  {
  return s.SynonymsDicts
}

func (s *Elasticsearch) GetTags() []*Tag  {
  return s.Tags
}

func (s *Elasticsearch) GetUpdatedAt() *string  {
  return s.UpdatedAt
}

func (s *Elasticsearch) GetWarmNode() *bool  {
  return s.WarmNode
}

func (s *Elasticsearch) GetWarmNodeConfiguration() *WarmNodeConfiguration  {
  return s.WarmNodeConfiguration
}

func (s *Elasticsearch) GetZoneCount() *int64  {
  return s.ZoneCount
}

func (s *Elasticsearch) GetZoneInfos() []*ZoneInfo  {
  return s.ZoneInfos
}

func (s *Elasticsearch) SetAdvancedDedicateMaster(v bool) *Elasticsearch {
  s.AdvancedDedicateMaster = &v
  return s
}

func (s *Elasticsearch) SetAdvancedSetting(v *ElasticsearchAdvancedSetting) *Elasticsearch {
  s.AdvancedSetting = v
  return s
}

func (s *Elasticsearch) SetAliwsDicts(v []*DictInfo) *Elasticsearch {
  s.AliwsDicts = v
  return s
}

func (s *Elasticsearch) SetClientNodeConfiguration(v *ClientNodeConfiguration) *Elasticsearch {
  s.ClientNodeConfiguration = v
  return s
}

func (s *Elasticsearch) SetCreatedAt(v string) *Elasticsearch {
  s.CreatedAt = &v
  return s
}

func (s *Elasticsearch) SetDataNode(v bool) *Elasticsearch {
  s.DataNode = &v
  return s
}

func (s *Elasticsearch) SetDedicateMaster(v bool) *Elasticsearch {
  s.DedicateMaster = &v
  return s
}

func (s *Elasticsearch) SetDescription(v string) *Elasticsearch {
  s.Description = &v
  return s
}

func (s *Elasticsearch) SetDictList(v []*DictInfo) *Elasticsearch {
  s.DictList = v
  return s
}

func (s *Elasticsearch) SetDomain(v string) *Elasticsearch {
  s.Domain = &v
  return s
}

func (s *Elasticsearch) SetElasticDataNodeConfiguration(v *ElasticDataNodeConfiguration) *Elasticsearch {
  s.ElasticDataNodeConfiguration = v
  return s
}

func (s *Elasticsearch) SetEnableKibanaPrivateNetwork(v bool) *Elasticsearch {
  s.EnableKibanaPrivateNetwork = &v
  return s
}

func (s *Elasticsearch) SetEnableKibanaPublicNetwork(v bool) *Elasticsearch {
  s.EnableKibanaPublicNetwork = &v
  return s
}

func (s *Elasticsearch) SetEnablePublic(v bool) *Elasticsearch {
  s.EnablePublic = &v
  return s
}

func (s *Elasticsearch) SetEndTime(v int64) *Elasticsearch {
  s.EndTime = &v
  return s
}

func (s *Elasticsearch) SetEsConfig(v map[string]*string) *Elasticsearch {
  s.EsConfig = v
  return s
}

func (s *Elasticsearch) SetEsIPWhitelist(v []*string) *Elasticsearch {
  s.EsIPWhitelist = v
  return s
}

func (s *Elasticsearch) SetEsVersion(v string) *Elasticsearch {
  s.EsVersion = &v
  return s
}

func (s *Elasticsearch) SetExtendConfigs(v []map[string]interface{}) *Elasticsearch {
  s.ExtendConfigs = v
  return s
}

func (s *Elasticsearch) SetHaveClientNode(v bool) *Elasticsearch {
  s.HaveClientNode = &v
  return s
}

func (s *Elasticsearch) SetHaveElasticDataNode(v bool) *Elasticsearch {
  s.HaveElasticDataNode = &v
  return s
}

func (s *Elasticsearch) SetHaveKibana(v bool) *Elasticsearch {
  s.HaveKibana = &v
  return s
}

func (s *Elasticsearch) SetIkHotDicts(v []*DictInfo) *Elasticsearch {
  s.IkHotDicts = v
  return s
}

func (s *Elasticsearch) SetInstanceId(v string) *Elasticsearch {
  s.InstanceId = &v
  return s
}

func (s *Elasticsearch) SetKibanaConfiguration(v *KibanaNodeConfiguration) *Elasticsearch {
  s.KibanaConfiguration = v
  return s
}

func (s *Elasticsearch) SetKibanaDomain(v string) *Elasticsearch {
  s.KibanaDomain = &v
  return s
}

func (s *Elasticsearch) SetKibanaIPWhitelist(v []*string) *Elasticsearch {
  s.KibanaIPWhitelist = v
  return s
}

func (s *Elasticsearch) SetKibanaPort(v int64) *Elasticsearch {
  s.KibanaPort = &v
  return s
}

func (s *Elasticsearch) SetKibanaPrivateDomain(v string) *Elasticsearch {
  s.KibanaPrivateDomain = &v
  return s
}

func (s *Elasticsearch) SetKibanaPrivateIPWhitelist(v []*string) *Elasticsearch {
  s.KibanaPrivateIPWhitelist = v
  return s
}

func (s *Elasticsearch) SetKibanaPrivatePort(v int64) *Elasticsearch {
  s.KibanaPrivatePort = &v
  return s
}

func (s *Elasticsearch) SetKibanaProtocol(v string) *Elasticsearch {
  s.KibanaProtocol = &v
  return s
}

func (s *Elasticsearch) SetMasterConfiguration(v *MasterNodeConfiguration) *Elasticsearch {
  s.MasterConfiguration = v
  return s
}

func (s *Elasticsearch) SetNetworkConfig(v *NetworkConfig) *Elasticsearch {
  s.NetworkConfig = v
  return s
}

func (s *Elasticsearch) SetNodeAmount(v int64) *Elasticsearch {
  s.NodeAmount = &v
  return s
}

func (s *Elasticsearch) SetNodeSpec(v *NodeSpec) *Elasticsearch {
  s.NodeSpec = v
  return s
}

func (s *Elasticsearch) SetPaymentType(v string) *Elasticsearch {
  s.PaymentType = &v
  return s
}

func (s *Elasticsearch) SetPort(v int64) *Elasticsearch {
  s.Port = &v
  return s
}

func (s *Elasticsearch) SetPrivateNetworkIpWhiteList(v []*string) *Elasticsearch {
  s.PrivateNetworkIpWhiteList = v
  return s
}

func (s *Elasticsearch) SetProductType(v string) *Elasticsearch {
  s.ProductType = &v
  return s
}

func (s *Elasticsearch) SetProtocol(v string) *Elasticsearch {
  s.Protocol = &v
  return s
}

func (s *Elasticsearch) SetPublicDomain(v string) *Elasticsearch {
  s.PublicDomain = &v
  return s
}

func (s *Elasticsearch) SetPublicIpWhitelist(v []*string) *Elasticsearch {
  s.PublicIpWhitelist = v
  return s
}

func (s *Elasticsearch) SetPublicPort(v int64) *Elasticsearch {
  s.PublicPort = &v
  return s
}

func (s *Elasticsearch) SetReadWritePolicy(v *ReadWritePolicy) *Elasticsearch {
  s.ReadWritePolicy = v
  return s
}

func (s *Elasticsearch) SetResourceGroupId(v string) *Elasticsearch {
  s.ResourceGroupId = &v
  return s
}

func (s *Elasticsearch) SetServiceVpc(v bool) *Elasticsearch {
  s.ServiceVpc = &v
  return s
}

func (s *Elasticsearch) SetStatus(v string) *Elasticsearch {
  s.Status = &v
  return s
}

func (s *Elasticsearch) SetSynonymsDicts(v []*DictInfo) *Elasticsearch {
  s.SynonymsDicts = v
  return s
}

func (s *Elasticsearch) SetTags(v []*Tag) *Elasticsearch {
  s.Tags = v
  return s
}

func (s *Elasticsearch) SetUpdatedAt(v string) *Elasticsearch {
  s.UpdatedAt = &v
  return s
}

func (s *Elasticsearch) SetWarmNode(v bool) *Elasticsearch {
  s.WarmNode = &v
  return s
}

func (s *Elasticsearch) SetWarmNodeConfiguration(v *WarmNodeConfiguration) *Elasticsearch {
  s.WarmNodeConfiguration = v
  return s
}

func (s *Elasticsearch) SetZoneCount(v int64) *Elasticsearch {
  s.ZoneCount = &v
  return s
}

func (s *Elasticsearch) SetZoneInfos(v []*ZoneInfo) *Elasticsearch {
  s.ZoneInfos = v
  return s
}

func (s *Elasticsearch) Validate() error {
  return dara.Validate(s)
}

type ElasticsearchAdvancedSetting struct {
  GcName *string `json:"gcName,omitempty" xml:"gcName,omitempty"`
}

func (s ElasticsearchAdvancedSetting) String() string {
  return dara.Prettify(s)
}

func (s ElasticsearchAdvancedSetting) GoString() string {
  return s.String()
}

func (s *ElasticsearchAdvancedSetting) GetGcName() *string  {
  return s.GcName
}

func (s *ElasticsearchAdvancedSetting) SetGcName(v string) *ElasticsearchAdvancedSetting {
  s.GcName = &v
  return s
}

func (s *ElasticsearchAdvancedSetting) Validate() error {
  return dara.Validate(s)
}

