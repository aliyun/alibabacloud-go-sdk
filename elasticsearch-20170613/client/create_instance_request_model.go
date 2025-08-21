// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientNodeConfiguration(v *ClientNodeConfiguration) *CreateInstanceRequest
	GetClientNodeConfiguration() *ClientNodeConfiguration
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetElasticDataNodeConfiguration(v *ElasticDataNodeConfiguration) *CreateInstanceRequest
	GetElasticDataNodeConfiguration() *ElasticDataNodeConfiguration
	SetEsAdminPassword(v string) *CreateInstanceRequest
	GetEsAdminPassword() *string
	SetEsVersion(v string) *CreateInstanceRequest
	GetEsVersion() *string
	SetInstanceCategory(v string) *CreateInstanceRequest
	GetInstanceCategory() *string
	SetKibanaConfiguration(v *KibanaNodeConfiguration) *CreateInstanceRequest
	GetKibanaConfiguration() *KibanaNodeConfiguration
	SetMasterConfiguration(v *MasterNodeConfiguration) *CreateInstanceRequest
	GetMasterConfiguration() *MasterNodeConfiguration
	SetNetworkConfig(v *NetworkConfig) *CreateInstanceRequest
	GetNetworkConfig() *NetworkConfig
	SetNodeAmount(v int32) *CreateInstanceRequest
	GetNodeAmount() *int32
	SetNodeSpec(v *NodeSpec) *CreateInstanceRequest
	GetNodeSpec() *NodeSpec
	SetPaymentInfo(v *PaymentInfo) *CreateInstanceRequest
	GetPaymentInfo() *PaymentInfo
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
	SetWarmNodeConfiguration(v *WarmNodeConfiguration) *CreateInstanceRequest
	GetWarmNodeConfiguration() *WarmNodeConfiguration
	SetZoneCount(v int32) *CreateInstanceRequest
	GetZoneCount() *int32
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
}

type CreateInstanceRequest struct {
	ClientNodeConfiguration *ClientNodeConfiguration `json:"clientNodeConfiguration,omitempty" xml:"clientNodeConfiguration,omitempty"`
	// example:
	//
	// es
	Description                  *string                       `json:"description,omitempty" xml:"description,omitempty"`
	ElasticDataNodeConfiguration *ElasticDataNodeConfiguration `json:"elasticDataNodeConfiguration,omitempty" xml:"elasticDataNodeConfiguration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// es_password
	EsAdminPassword *string `json:"esAdminPassword,omitempty" xml:"esAdminPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5.5.3_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// example:
	//
	// advanced
	InstanceCategory    *string                  `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	KibanaConfiguration *KibanaNodeConfiguration `json:"kibanaConfiguration,omitempty" xml:"kibanaConfiguration,omitempty"`
	MasterConfiguration *MasterNodeConfiguration `json:"masterConfiguration,omitempty" xml:"masterConfiguration,omitempty"`
	// This parameter is required.
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	NodeAmount  *int32       `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec    *NodeSpec    `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty"`
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty" xml:"paymentInfo,omitempty"`
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// rg-aekzu7tsu4nem4q
	ResourceGroupId       *string                      `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags                  []*CreateInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	WarmNodeConfiguration *WarmNodeConfiguration       `json:"warmNodeConfiguration,omitempty" xml:"warmNodeConfiguration,omitempty"`
	// example:
	//
	// 2
	ZoneCount *int32 `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetClientNodeConfiguration() *ClientNodeConfiguration {
	return s.ClientNodeConfiguration
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetElasticDataNodeConfiguration() *ElasticDataNodeConfiguration {
	return s.ElasticDataNodeConfiguration
}

func (s *CreateInstanceRequest) GetEsAdminPassword() *string {
	return s.EsAdminPassword
}

func (s *CreateInstanceRequest) GetEsVersion() *string {
	return s.EsVersion
}

func (s *CreateInstanceRequest) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *CreateInstanceRequest) GetKibanaConfiguration() *KibanaNodeConfiguration {
	return s.KibanaConfiguration
}

func (s *CreateInstanceRequest) GetMasterConfiguration() *MasterNodeConfiguration {
	return s.MasterConfiguration
}

func (s *CreateInstanceRequest) GetNetworkConfig() *NetworkConfig {
	return s.NetworkConfig
}

func (s *CreateInstanceRequest) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *CreateInstanceRequest) GetNodeSpec() *NodeSpec {
	return s.NodeSpec
}

func (s *CreateInstanceRequest) GetPaymentInfo() *PaymentInfo {
	return s.PaymentInfo
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) GetWarmNodeConfiguration() *WarmNodeConfiguration {
	return s.WarmNodeConfiguration
}

func (s *CreateInstanceRequest) GetZoneCount() *int32 {
	return s.ZoneCount
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) SetClientNodeConfiguration(v *ClientNodeConfiguration) *CreateInstanceRequest {
	s.ClientNodeConfiguration = v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetElasticDataNodeConfiguration(v *ElasticDataNodeConfiguration) *CreateInstanceRequest {
	s.ElasticDataNodeConfiguration = v
	return s
}

func (s *CreateInstanceRequest) SetEsAdminPassword(v string) *CreateInstanceRequest {
	s.EsAdminPassword = &v
	return s
}

func (s *CreateInstanceRequest) SetEsVersion(v string) *CreateInstanceRequest {
	s.EsVersion = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceCategory(v string) *CreateInstanceRequest {
	s.InstanceCategory = &v
	return s
}

func (s *CreateInstanceRequest) SetKibanaConfiguration(v *KibanaNodeConfiguration) *CreateInstanceRequest {
	s.KibanaConfiguration = v
	return s
}

func (s *CreateInstanceRequest) SetMasterConfiguration(v *MasterNodeConfiguration) *CreateInstanceRequest {
	s.MasterConfiguration = v
	return s
}

func (s *CreateInstanceRequest) SetNetworkConfig(v *NetworkConfig) *CreateInstanceRequest {
	s.NetworkConfig = v
	return s
}

func (s *CreateInstanceRequest) SetNodeAmount(v int32) *CreateInstanceRequest {
	s.NodeAmount = &v
	return s
}

func (s *CreateInstanceRequest) SetNodeSpec(v *NodeSpec) *CreateInstanceRequest {
	s.NodeSpec = v
	return s
}

func (s *CreateInstanceRequest) SetPaymentInfo(v *PaymentInfo) *CreateInstanceRequest {
	s.PaymentInfo = v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetWarmNodeConfiguration(v *WarmNodeConfiguration) *CreateInstanceRequest {
	s.WarmNodeConfiguration = v
	return s
}

func (s *CreateInstanceRequest) SetZoneCount(v int32) *CreateInstanceRequest {
	s.ZoneCount = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateInstanceRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateInstanceRequestTags) SetTagKey(v string) *CreateInstanceRequestTags {
	s.TagKey = &v
	return s
}

func (s *CreateInstanceRequestTags) SetTagValue(v string) *CreateInstanceRequestTags {
	s.TagValue = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
