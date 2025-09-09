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
	SetLicenseMetadata(v string) *GetServiceInstanceResponseBody
	GetLicenseMetadata() *string
	SetName(v string) *GetServiceInstanceResponseBody
	GetName() *string
	SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody
	GetNetworkConfig() *GetServiceInstanceResponseBodyNetworkConfig
	SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody
	GetOperatedServiceInstanceId() *string
	SetOperationEndTime(v string) *GetServiceInstanceResponseBody
	GetOperationEndTime() *string
	SetOperationExtraInfo(v string) *GetServiceInstanceResponseBody
	GetOperationExtraInfo() *string
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
	SetRdAccountLoginUrl(v string) *GetServiceInstanceResponseBody
	GetRdAccountLoginUrl() *string
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
	// The business status of the service instance. Valid values:
	//
	// 	- Normal: The service instance is normal.
	//
	// 	- Renewing: The service instance is being renewed.
	//
	// 	- RenewFailed: The service instance failed to be renewed.
	//
	// 	- Expired: The service instance expired.
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The time when the service instance was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Indicates whether the Prometheus monitoring feature is enabled on the user side.
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The time when the service instance expires.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// https://g.xxxxxx.aliyun.com/d/xxxxxxxx-cn-mariadb/mysql-xxxxxxxx-quickstart-and-dashboard?orgId=355401&refresh=10s
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
	// The license metadata.
	//
	// example:
	//
	// {"renewType":"MONTHLY"}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
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
	// The ID of the service instance that is used to implement hosted O\\&M.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The end of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// Operate extra info
	//
	// example:
	//
	// 6cc5ce3590ed7f5529e85fec
	OperationExtraInfo *string `json:"OperationExtraInfo,omitempty" xml:"OperationExtraInfo,omitempty"`
	// The beginning of the time range during which hosted O\\&M is implemented.
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
	// {
	//
	//       "InstanceIds": [
	//
	//             "i-hp38ofxl0dsyfi7z****"
	//
	//       ]
	//
	// }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {
	//
	//       "param": "value"
	//
	// }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method of the service. Valid values:
	//
	// 	- Permanent: Once you purchase the service, you can use it permanently.
	//
	// 	- Subscription: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a subscription basis.
	//
	// 	- PayAsYouGo: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a pay-as-you-go basis.
	//
	// 	- CustomFixTime: You are charged for the service based on a custom duration fixed by the service provider.
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package 1
	PredefinedParameterName *string `json:"PredefinedParameterName,omitempty" xml:"PredefinedParameterName,omitempty"`
	// The deployment progress of the service instance. Unit: percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The logon URL for the accounts in the resource directory corresponding to the service instance.
	//
	// example:
	//
	// https://signin.aliyun.com/switchRole.htm?parentAlias=1307782318217247.onaliyun.com&roleName=ResourceDirectoryAccountAccessRole
	RdAccountLoginUrl *string `json:"RdAccountLoginUrl,omitempty" xml:"RdAccountLoginUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resources.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "StackId": "stack-xxx"
	//
	//       }
	//
	// ]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The information about the service to which the service instance belongs.
	Service *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source of the service instance. Valid values:
	//
	// 	- User: Compute Nest customer
	//
	// 	- Market: Alibaba Cloud Marketplace
	//
	// 	- Supplier: Compute Nest service provider
	//
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The deployment state of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
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
	// 1964460391538545
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The custom tags.
	Tags []*GetServiceInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time when the service instance was updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// example:
	//
	// 1234567
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

func (s *GetServiceInstanceResponseBody) GetLicenseMetadata() *string {
	return s.LicenseMetadata
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

func (s *GetServiceInstanceResponseBody) GetOperationExtraInfo() *string {
	return s.OperationExtraInfo
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

func (s *GetServiceInstanceResponseBody) GetRdAccountLoginUrl() *string {
	return s.RdAccountLoginUrl
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

func (s *GetServiceInstanceResponseBody) SetLicenseMetadata(v string) *GetServiceInstanceResponseBody {
	s.LicenseMetadata = &v
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

func (s *GetServiceInstanceResponseBody) SetOperationExtraInfo(v string) *GetServiceInstanceResponseBody {
	s.OperationExtraInfo = &v
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

func (s *GetServiceInstanceResponseBody) SetRdAccountLoginUrl(v string) *GetServiceInstanceResponseBody {
	s.RdAccountLoginUrl = &v
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
	// The ID of the endpoint service for the private connection.
	//
	// >  This parameter is discontinued.
	//
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	// The information about private connections.
	PrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
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

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetEndpointServiceId() *string {
	return s.EndpointServiceId
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetPrivateVpcConnections() []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	return s.PrivateVpcConnections
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) GetReversePrivateVpcConnections() []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	return s.ReversePrivateVpcConnections
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
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
	// The network configurations, which are mainly used for the private connection.
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	// The ID of the endpoint for the private connection.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the endpoint service for the private connection.
	//
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	// The custom domain name.
	//
	// example:
	//
	// test.computenest.aliyuncs.com
	PrivateZoneName *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
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

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetEndpointServiceId() *string {
	return s.EndpointServiceId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GetPrivateZoneName() *string {
	return s.PrivateZoneName
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
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
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP addresses of the endpoints for private connections.
	EndpointIps []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	// The status of the Ingress endpoint. Valid values:
	//
	// 	- Ready: The Ingress endpoint is connected.
	//
	// 	- Pending: The Ingress endpoint is being connected.
	//
	// 	- Failed: The Ingress endpoint fails to be connected.
	//
	// 	- Deleted: The Ingress endpoint is deleted.
	//
	// 	- Deleting: The Ingress endpoint is being deleted.
	//
	// example:
	//
	// Ready
	IngressEndpointStatus *string `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	// The status of the network service. Valid values:
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
	// The names of the security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The names of the vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
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
	// The ID of the endpoint for the reverse private connection.
	//
	// example:
	//
	// ep-m5ei42370541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the endpoint service for the reverse private connection.
	//
	// example:
	//
	// epsrv-5ei07324541816bxxxx
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
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

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GetEndpointServiceId() *string {
	return s.EndpointServiceId
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointServiceId = &v
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
	// The time when the service was published.
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
	// Draft: The registration request of the service is pending to be submitted.
	//
	// Submitted: The registration request of the service is submitted.
	//
	// Approved: The registration request of the service is approved.
	//
	// Online: The service is published.
	//
	// Offline: The service is unpublished.
	//
	// Deleted: The service is deleted.
	//
	// Launching: The service is being published.
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
	// The upgradable service Info.
	UpgradableServiceInfos []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos `json:"UpgradableServiceInfos,omitempty" xml:"UpgradableServiceInfos,omitempty" type:"Repeated"`
	// The service versions that can be updated.
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
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
	// 20231127020225
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
	// The service name.
	//
	// example:
	//
	// WordPress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
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
	// The upgradable service version.
	//
	// example:
	//
	// 4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version name of an upgradable service version.
	//
	// example:
	//
	// Init version
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
