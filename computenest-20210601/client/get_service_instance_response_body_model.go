// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizStatus(v string) *GetServiceInstanceResponseBody
	GetBizStatus() *string
	SetComponents(v string) *GetServiceInstanceResponseBody
	GetComponents() *string
	SetCreateTime(v string) *GetServiceInstanceResponseBody
	GetCreateTime() *string
	SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody
	GetEnableInstanceOps() *bool
	SetEnableUserPrometheus(v bool) *GetServiceInstanceResponseBody
	GetEnableUserPrometheus() *bool
	SetEndTime(v string) *GetServiceInstanceResponseBody
	GetEndTime() *string
	SetGrafanaDashBoardUrl(v string) *GetServiceInstanceResponseBody
	GetGrafanaDashBoardUrl() *string
	SetIsOperated(v bool) *GetServiceInstanceResponseBody
	GetIsOperated() *bool
	SetLicenseEndTime(v string) *GetServiceInstanceResponseBody
	GetLicenseEndTime() *string
	SetMarketInstanceId(v string) *GetServiceInstanceResponseBody
	GetMarketInstanceId() *string
	SetName(v string) *GetServiceInstanceResponseBody
	GetName() *string
	SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody
	GetNetworkConfig() *GetServiceInstanceResponseBodyNetworkConfig
	SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody
	GetOperatedServiceInstanceId() *string
	SetOperationEndTime(v string) *GetServiceInstanceResponseBody
	GetOperationEndTime() *string
	SetOperationStartTime(v string) *GetServiceInstanceResponseBody
	GetOperationStartTime() *string
	SetOutputs(v string) *GetServiceInstanceResponseBody
	GetOutputs() *string
	SetParameters(v string) *GetServiceInstanceResponseBody
	GetParameters() *string
	SetPayType(v string) *GetServiceInstanceResponseBody
	GetPayType() *string
	SetPredefinedParameterName(v string) *GetServiceInstanceResponseBody
	GetPredefinedParameterName() *string
	SetProgress(v int64) *GetServiceInstanceResponseBody
	GetProgress() *int64
	SetRequestId(v string) *GetServiceInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetServiceInstanceResponseBody
	GetResourceGroupId() *string
	SetResources(v string) *GetServiceInstanceResponseBody
	GetResources() *string
	SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody
	GetService() *GetServiceInstanceResponseBodyService
	SetServiceInstanceId(v string) *GetServiceInstanceResponseBody
	GetServiceInstanceId() *string
	SetServiceType(v string) *GetServiceInstanceResponseBody
	GetServiceType() *string
	SetSource(v string) *GetServiceInstanceResponseBody
	GetSource() *string
	SetStatus(v string) *GetServiceInstanceResponseBody
	GetStatus() *string
	SetStatusDetail(v string) *GetServiceInstanceResponseBody
	GetStatusDetail() *string
	SetSupplierUid(v int64) *GetServiceInstanceResponseBody
	GetSupplierUid() *int64
	SetSupportTrialToPrivate(v bool) *GetServiceInstanceResponseBody
	GetSupportTrialToPrivate() *bool
	SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody
	GetTags() []*GetServiceInstanceResponseBodyTags
	SetTemplateName(v string) *GetServiceInstanceResponseBody
	GetTemplateName() *string
	SetUpdateTime(v string) *GetServiceInstanceResponseBody
	GetUpdateTime() *string
	SetUserId(v int64) *GetServiceInstanceResponseBody
	GetUserId() *int64
}

type GetServiceInstanceResponseBody struct {
	// The business state of the service instance. Valid values:
	//
	// 	- Normal
	//
	// 	- Renewing
	//
	// 	- RenewFailed
	//
	// 	- Expired
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// Cloud Marketplace additional billing items.
	//
	// example:
	//
	// {"TiKVServerCount":"3","package_version":"yuncode5398300001","PDServerCount":"3","TiDBServerCount":"2"}
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The time when the serviceInstance was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance supports the operation feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Whether to enable Prometheus monitoring.
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The expiration time of service instance.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// https://g.console.aliyun.com/d/xxxxxxxx-cn-mariadb/mysql-xxxxxx-xxxxxxxx-and-dashboard?orgId=355401&refresh=10s
	GrafanaDashBoardUrl *string `json:"GrafanaDashBoardUrl,omitempty" xml:"GrafanaDashBoardUrl,omitempty"`
	// Indicates whether the hosted O\\&M feature is enabled for the service instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// The expiration time of licence.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	LicenseEndTime *string `json:"LicenseEndTime,omitempty" xml:"LicenseEndTime,omitempty"`
	// The market Instance ID.
	//
	// example:
	//
	// 704***59
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network configurations.
	//
	// >  This parameter is discontinued.
	NetworkConfig *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// The serviceInstance  to be operated.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The operation end time.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The operation start time.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The outputs returned from creating the service instance.
	//
	// 	- If the service is deployed by using a ROS template, all output fields of the template are returned.
	//
	// 	- If the service is deployed by calling an SPI operation, the output fields of the service provider and for the Compute Nest additional features are returned.
	//
	// example:
	//
	// {"InstanceIds":["i-hp38ofxl0dsyfi7z****"]}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method of the instance for market. Valid values:
	//
	// 	- Permanent: Permanent purchase
	//
	// 	- Subscription: Subscription.
	//
	// 	- PayAsYouGo: Pay-as-you-go.
	//
	// 	- CustomFixTime: Merchant custom fixed duration.
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package one
	PredefinedParameterName *string `json:"PredefinedParameterName,omitempty" xml:"PredefinedParameterName,omitempty"`
	// The deployment progress of the service instance. Unit: percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources.
	//
	// example:
	//
	// [{"StackId": "stack-xxx"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The service details.
	Service *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of the service. Valid values:
	//
	// - private: The service is a private service and is deployed within the account of a customer.
	//
	// - managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// - operation: The service is a hosted O&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source of the serviceInstance. Valid values:
	//
	// - User
	//
	// - Market
	//
	// - Supplier
	//
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The deploy status of the serviceInstance. Valid values:
	//
	// - Created
	//
	// - Deploying
	//
	// - DeployedFailed
	//
	// - Deployed
	//
	// - Upgrading
	//
	// - Deleting
	//
	// - Deleted
	//
	// - DeletedFailed
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment state of the service instance.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 158927391332*****
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// Is it supported to convert from trial to private
	SupportTrialToPrivate *bool `json:"SupportTrialToPrivate,omitempty" xml:"SupportTrialToPrivate,omitempty"`
	// The tags.
	Tags []*GetServiceInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// Template one
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time when the serviceInstance  was last updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The AliUid of the user.
	//
	// example:
	//
	// 130920852836****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) GetBizStatus() *string {
	return s.BizStatus
}

func (s *GetServiceInstanceResponseBody) GetComponents() *string {
	return s.Components
}

func (s *GetServiceInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetServiceInstanceResponseBody) GetEnableInstanceOps() *bool {
	return s.EnableInstanceOps
}

func (s *GetServiceInstanceResponseBody) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *GetServiceInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetServiceInstanceResponseBody) GetGrafanaDashBoardUrl() *string {
	return s.GrafanaDashBoardUrl
}

func (s *GetServiceInstanceResponseBody) GetIsOperated() *bool {
	return s.IsOperated
}

func (s *GetServiceInstanceResponseBody) GetLicenseEndTime() *string {
	return s.LicenseEndTime
}

func (s *GetServiceInstanceResponseBody) GetMarketInstanceId() *string {
	return s.MarketInstanceId
}

func (s *GetServiceInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *GetServiceInstanceResponseBody) GetNetworkConfig() *GetServiceInstanceResponseBodyNetworkConfig {
	return s.NetworkConfig
}

func (s *GetServiceInstanceResponseBody) GetOperatedServiceInstanceId() *string {
	return s.OperatedServiceInstanceId
}

func (s *GetServiceInstanceResponseBody) GetOperationEndTime() *string {
	return s.OperationEndTime
}

func (s *GetServiceInstanceResponseBody) GetOperationStartTime() *string {
	return s.OperationStartTime
}

func (s *GetServiceInstanceResponseBody) GetOutputs() *string {
	return s.Outputs
}

func (s *GetServiceInstanceResponseBody) GetParameters() *string {
	return s.Parameters
}

func (s *GetServiceInstanceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetServiceInstanceResponseBody) GetPredefinedParameterName() *string {
	return s.PredefinedParameterName
}

func (s *GetServiceInstanceResponseBody) GetProgress() *int64 {
	return s.Progress
}

func (s *GetServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetServiceInstanceResponseBody) GetResources() *string {
	return s.Resources
}

func (s *GetServiceInstanceResponseBody) GetService() *GetServiceInstanceResponseBodyService {
	return s.Service
}

func (s *GetServiceInstanceResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceInstanceResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceInstanceResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetServiceInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetServiceInstanceResponseBody) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *GetServiceInstanceResponseBody) GetSupplierUid() *int64 {
	return s.SupplierUid
}

func (s *GetServiceInstanceResponseBody) GetSupportTrialToPrivate() *bool {
	return s.SupportTrialToPrivate
}

func (s *GetServiceInstanceResponseBody) GetTags() []*GetServiceInstanceResponseBodyTags {
	return s.Tags
}

func (s *GetServiceInstanceResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceInstanceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetServiceInstanceResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *GetServiceInstanceResponseBody) SetBizStatus(v string) *GetServiceInstanceResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetComponents(v string) *GetServiceInstanceResponseBody {
	s.Components = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableUserPrometheus(v bool) *GetServiceInstanceResponseBody {
	s.EnableUserPrometheus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetGrafanaDashBoardUrl(v string) *GetServiceInstanceResponseBody {
	s.GrafanaDashBoardUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseEndTime(v string) *GetServiceInstanceResponseBody {
	s.LicenseEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetMarketInstanceId(v string) *GetServiceInstanceResponseBody {
	s.MarketInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPredefinedParameterName(v string) *GetServiceInstanceResponseBody {
	s.PredefinedParameterName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResourceGroupId(v string) *GetServiceInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupportTrialToPrivate(v bool) *GetServiceInstanceResponseBody {
	s.SupportTrialToPrivate = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	// The ID of the endpoint for the private connection.
	//
	// >  This parameter is discontinued.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The information about private connections.
	PrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	// The PrivateZone ID.
	//
	// example:
	//
	// cb7f214f80ac348d87daaeac1f35****
	PrivateZoneId *string `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	// The information about the reverse private connection.
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetPrivateVpcConnections() []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	return s.PrivateVpcConnections
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetPrivateZoneId() *string {
	return s.PrivateZoneId
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetReversePrivateVpcConnections() []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	return s.ReversePrivateVpcConnections
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	// The network configurations, which are mainly used for private connections.
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	// The endpoint ID of the private connection.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the private zone of the custom private domain name.
	//
	// example:
	//
	// cb7f214f80ac348d87daaeac1f35****
	PrivateZoneId *string `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	// The custom domain name.
	//
	// example:
	//
	// test.computenest.aliyuncs.com
	PrivateZoneName *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
	// The region ID of the endpoint of the PrivateLink connection.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetConnectionConfigs() []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	return s.ConnectionConfigs
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetPrivateZoneId() *string {
	return s.PrivateZoneId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetPrivateZoneName() *string {
	return s.PrivateZoneName
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetRegionId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs struct {
	// The bandwidth limit for the private connection established based on the private network interconnection mode of Compute Nest.
	//
	// example:
	//
	// 1536Mbps
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The domain name.
	//
	// example:
	//
	// ie-569a9be34f5534f6bc6559b5c1xxxxxx.service-51f80502802e48xxxxxx.cn-hangzhou.computenest.aliyuncs.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP addresses of the endpoints of the private connections.
	EndpointIps []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	// The state of the ingress endpoint. Valid values:
	//
	// 	- Ready: The ingress endpoint is connected.
	//
	// 	- Pending: The ingress endpoint is being connected.
	//
	// 	- Failed: The ingress endpoint fails to be connected.
	//
	// 	- Deleted: The ingress endpoint is deleted.
	//
	// 	- Deleting: The ingress endpoint is being deleted.
	//
	// example:
	//
	// Ready
	IngressEndpointStatus *string `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	// The state of the network service. Valid values:
	//
	// 	- Ready: The network service is connected.
	//
	// 	- Pending: The network service is being connected.
	//
	// 	- Failed: The network service fails to be connected.
	//
	// 	- Deleted: The network service is deleted.
	//
	// 	- Deleting: The network service is being deleted.
	//
	// example:
	//
	// Ready
	NetworkServiceStatus *string `json:"NetworkServiceStatus,omitempty" xml:"NetworkServiceStatus,omitempty"`
	// The region ID of the VPC to which the endpoint of the private connection established based on the private network interconnection mode of Compute Nest belongs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The names of the security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The names of the vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetConnectBandwidth() *int32 {
	return s.ConnectBandwidth
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetDomainName() *string {
	return s.DomainName
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetEndpointIps() []*string {
	return s.EndpointIps
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetIngressEndpointStatus() *string {
	return s.IngressEndpointStatus
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetNetworkServiceStatus() *string {
	return s.NetworkServiceStatus
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetSecurityGroups() []*string {
	return s.SecurityGroups
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetVSwitches() []*string {
	return s.VSwitches
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GetVpcId() *string {
	return s.VpcId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetConnectBandwidth(v int32) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetDomainName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.DomainName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetEndpointIps(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.EndpointIps = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetIngressEndpointStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.IngressEndpointStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetNetworkServiceStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.NetworkServiceStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetRegionId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetSecurityGroups(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.SecurityGroups = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVSwitches(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VSwitches = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVpcId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VpcId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	// The endpoint ID of the reverse private connection.
	//
	// example:
	//
	// ep-m5ei42370541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyService struct {
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// Parameters related to O\\&M operations, including configuration change, prometheus, and log configurations.
	//
	// example:
	//
	// {"SupportBackup":false,"PrometheusConfigMap":{},"ModifyParametersConfig":[{"TemplateName":"Chinese mainland version","Operation":[{"Name":"Change Plan","Description":"Change Plan","Type":"Custom","SupportPredefinedParameters":true,"EnableLogging":false},{"Name":"Parameter configuration change","Description":"Parameter configuration change","Type":"Custom","SupportPredefinedParameters":false,"EnableLogging":false,"Parameters":["DataDiskSize"]}]}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The time when the service version was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The URL of the service documentation.
	//
	// example:
	//
	// http://example.com
	ServiceDocUrl *string `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-9c8a3522528b4fe8****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Draft
	//
	// 	- Submited
	//
	// 	- Approved
	//
	// 	- Online
	//
	// 	- Offline
	//
	// 	- Deleted
	//
	// 	- Launching
	//
	// 	- Beta
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service versions that can be updated.
	UpgradableServiceInfos []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos `json:"UpgradableServiceInfos,omitempty" xml:"UpgradableServiceInfos,omitempty" type:"Repeated"`
	// The service version that can be updated.
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"Allowed Upgrade 11\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\",\\"Resource\\"]}
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// 1.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) GetDeployMetadata() *string {
	return s.DeployMetadata
}

func (s *GetServiceInstanceResponseBodyService) GetDeployType() *string {
	return s.DeployType
}

func (s *GetServiceInstanceResponseBodyService) GetOperationMetadata() *string {
	return s.OperationMetadata
}

func (s *GetServiceInstanceResponseBodyService) GetPublishTime() *string {
	return s.PublishTime
}

func (s *GetServiceInstanceResponseBodyService) GetServiceDocUrl() *string {
	return s.ServiceDocUrl
}

func (s *GetServiceInstanceResponseBodyService) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceInstanceResponseBodyService) GetServiceInfos() []*GetServiceInstanceResponseBodyServiceServiceInfos {
	return s.ServiceInfos
}

func (s *GetServiceInstanceResponseBodyService) GetServiceProductUrl() *string {
	return s.ServiceProductUrl
}

func (s *GetServiceInstanceResponseBodyService) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceInstanceResponseBodyService) GetStatus() *string {
	return s.Status
}

func (s *GetServiceInstanceResponseBodyService) GetSupplierName() *string {
	return s.SupplierName
}

func (s *GetServiceInstanceResponseBodyService) GetSupplierUrl() *string {
	return s.SupplierUrl
}

func (s *GetServiceInstanceResponseBodyService) GetUpgradableServiceInfos() []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	return s.UpgradableServiceInfos
}

func (s *GetServiceInstanceResponseBodyService) GetUpgradableServiceVersions() []*string {
	return s.UpgradableServiceVersions
}

func (s *GetServiceInstanceResponseBodyService) GetUpgradeMetadata() *string {
	return s.UpgradeMetadata
}

func (s *GetServiceInstanceResponseBodyService) GetVersion() *string {
	return s.Version
}

func (s *GetServiceInstanceResponseBodyService) GetVersionName() *string {
	return s.VersionName
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetOperationMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.OperationMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceInfos(v []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceVersions(v []*string) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceVersions = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradeMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.UpgradeMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service instance.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// Docker Community Edition
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// Docker Community Edition (CE) is a free version of the Docker project, aimed at developers, enthusiasts, and individuals and organizations who want to use container technology.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) GetImage() *string {
	return s.Image
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) GetLocale() *string {
	return s.Locale
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) GetName() *string {
	return s.Name
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyServiceUpgradableServiceInfos struct {
	// The service version.
	//
	// example:
	//
	// draft
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version name.
	//
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) GetVersion() *string {
	return s.Version
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) GetVersionName() *string {
	return s.VersionName
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) SetVersion(v string) *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) SetVersionName(v string) *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	s.VersionName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetServiceInstanceResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
