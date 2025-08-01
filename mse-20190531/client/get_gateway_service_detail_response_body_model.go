// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayServiceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayServiceDetailResponseBody
	GetCode() *int32
	SetData(v *GetGatewayServiceDetailResponseBodyData) *GetGatewayServiceDetailResponseBody
	GetData() *GetGatewayServiceDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetGatewayServiceDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayServiceDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayServiceDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayServiceDetailResponseBody
	GetSuccess() *bool
}

type GetGatewayServiceDetailResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetGatewayServiceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9802C54E-5CC5-5706-927B-993DBB6DCF2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayServiceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayServiceDetailResponseBody) GetData() *GetGatewayServiceDetailResponseBodyData {
	return s.Data
}

func (s *GetGatewayServiceDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayServiceDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayServiceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayServiceDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayServiceDetailResponseBody) SetCode(v int32) *GetGatewayServiceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBody) SetData(v *GetGatewayServiceDetailResponseBodyData) *GetGatewayServiceDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayServiceDetailResponseBody) SetHttpStatusCode(v int32) *GetGatewayServiceDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBody) SetMessage(v string) *GetGatewayServiceDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBody) SetRequestId(v string) *GetGatewayServiceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBody) SetSuccess(v bool) *GetGatewayServiceDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyData struct {
	DnsServerList []*string `json:"DnsServerList,omitempty" xml:"DnsServerList,omitempty" type:"Repeated"`
	// The ID of the gateway.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The traffic policy of the service.
	//
	// example:
	//
	// {}
	GatewayTrafficPolicy *TrafficPolicy `json:"GatewayTrafficPolicy,omitempty" xml:"GatewayTrafficPolicy,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether the health check is enabled.
	//
	// example:
	//
	// true
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The status of the health check. Valid values:
	//
	// example:
	//
	// true
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address of the service.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The details of the tag.
	LabelDetails []*GetGatewayServiceDetailResponseBodyDataLabelDetails `json:"LabelDetails,omitempty" xml:"LabelDetails,omitempty" type:"Repeated"`
	// The basic information about the service.
	//
	// example:
	//
	// {}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The traffic policy of service ports.
	PortTrafficPolicyList []*GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList `json:"PortTrafficPolicyList,omitempty" xml:"PortTrafficPolicyList,omitempty" type:"Repeated"`
	// The array of service ports.
	Ports       []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	ServiceFQDN *string  `json:"ServiceFQDN,omitempty" xml:"ServiceFQDN,omitempty"`
	// The name of the service registered with the service registry.
	//
	// example:
	//
	// test
	ServiceNameInRegistry *string `json:"ServiceNameInRegistry,omitempty" xml:"ServiceNameInRegistry,omitempty"`
	ServicePort           *int32  `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The protocol of the service.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The ID of the service source.
	//
	// example:
	//
	// 1
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type of the service.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The details of versions.
	VersionDetails []*GetGatewayServiceDetailResponseBodyDataVersionDetails `json:"VersionDetails,omitempty" xml:"VersionDetails,omitempty" type:"Repeated"`
	// The service version. This parameter is deprecated.
	Versions []*GetGatewayServiceDetailResponseBodyDataVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s GetGatewayServiceDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyData) GetDnsServerList() []*string {
	return s.DnsServerList
}

func (s *GetGatewayServiceDetailResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayServiceDetailResponseBodyData) GetGatewayTrafficPolicy() *TrafficPolicy {
	return s.GatewayTrafficPolicy
}

func (s *GetGatewayServiceDetailResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayServiceDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayServiceDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayServiceDetailResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGatewayServiceDetailResponseBodyData) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *GetGatewayServiceDetailResponseBodyData) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *GetGatewayServiceDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayServiceDetailResponseBodyData) GetIps() []*string {
	return s.Ips
}

func (s *GetGatewayServiceDetailResponseBodyData) GetLabelDetails() []*GetGatewayServiceDetailResponseBodyDataLabelDetails {
	return s.LabelDetails
}

func (s *GetGatewayServiceDetailResponseBodyData) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *GetGatewayServiceDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayServiceDetailResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetGatewayServiceDetailResponseBodyData) GetPortTrafficPolicyList() []*GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	return s.PortTrafficPolicyList
}

func (s *GetGatewayServiceDetailResponseBodyData) GetPorts() []*int32 {
	return s.Ports
}

func (s *GetGatewayServiceDetailResponseBodyData) GetServiceFQDN() *string {
	return s.ServiceFQDN
}

func (s *GetGatewayServiceDetailResponseBodyData) GetServiceNameInRegistry() *string {
	return s.ServiceNameInRegistry
}

func (s *GetGatewayServiceDetailResponseBodyData) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *GetGatewayServiceDetailResponseBodyData) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *GetGatewayServiceDetailResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetGatewayServiceDetailResponseBodyData) GetSourceType() *string {
	return s.SourceType
}

func (s *GetGatewayServiceDetailResponseBodyData) GetVersionDetails() []*GetGatewayServiceDetailResponseBodyDataVersionDetails {
	return s.VersionDetails
}

func (s *GetGatewayServiceDetailResponseBodyData) GetVersions() []*GetGatewayServiceDetailResponseBodyDataVersions {
	return s.Versions
}

func (s *GetGatewayServiceDetailResponseBodyData) SetDnsServerList(v []*string) *GetGatewayServiceDetailResponseBodyData {
	s.DnsServerList = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetGatewayId(v int64) *GetGatewayServiceDetailResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetGatewayTrafficPolicy(v *TrafficPolicy) *GetGatewayServiceDetailResponseBodyData {
	s.GatewayTrafficPolicy = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayServiceDetailResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetGmtCreate(v string) *GetGatewayServiceDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetGmtModified(v string) *GetGatewayServiceDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetGroupName(v string) *GetGatewayServiceDetailResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetHealthCheck(v string) *GetGatewayServiceDetailResponseBodyData {
	s.HealthCheck = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetHealthStatus(v string) *GetGatewayServiceDetailResponseBodyData {
	s.HealthStatus = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetId(v int64) *GetGatewayServiceDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetIps(v []*string) *GetGatewayServiceDetailResponseBodyData {
	s.Ips = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetLabelDetails(v []*GetGatewayServiceDetailResponseBodyDataLabelDetails) *GetGatewayServiceDetailResponseBodyData {
	s.LabelDetails = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetMetaInfo(v string) *GetGatewayServiceDetailResponseBodyData {
	s.MetaInfo = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetName(v string) *GetGatewayServiceDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetNamespace(v string) *GetGatewayServiceDetailResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetPortTrafficPolicyList(v []*GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) *GetGatewayServiceDetailResponseBodyData {
	s.PortTrafficPolicyList = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetPorts(v []*int32) *GetGatewayServiceDetailResponseBodyData {
	s.Ports = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetServiceFQDN(v string) *GetGatewayServiceDetailResponseBodyData {
	s.ServiceFQDN = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetServiceNameInRegistry(v string) *GetGatewayServiceDetailResponseBodyData {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetServicePort(v int32) *GetGatewayServiceDetailResponseBodyData {
	s.ServicePort = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetServiceProtocol(v string) *GetGatewayServiceDetailResponseBodyData {
	s.ServiceProtocol = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetSourceId(v int64) *GetGatewayServiceDetailResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetSourceType(v string) *GetGatewayServiceDetailResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetVersionDetails(v []*GetGatewayServiceDetailResponseBodyDataVersionDetails) *GetGatewayServiceDetailResponseBodyData {
	s.VersionDetails = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) SetVersions(v []*GetGatewayServiceDetailResponseBodyDataVersions) *GetGatewayServiceDetailResponseBodyData {
	s.Versions = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyDataLabelDetails struct {
	// The tag.
	//
	// example:
	//
	// label
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetGatewayServiceDetailResponseBodyDataLabelDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyDataLabelDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyDataLabelDetails) GetKey() *string {
	return s.Key
}

func (s *GetGatewayServiceDetailResponseBodyDataLabelDetails) GetValues() []*string {
	return s.Values
}

func (s *GetGatewayServiceDetailResponseBodyDataLabelDetails) SetKey(v string) *GetGatewayServiceDetailResponseBodyDataLabelDetails {
	s.Key = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataLabelDetails) SetValues(v []*string) *GetGatewayServiceDetailResponseBodyDataLabelDetails {
	s.Values = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataLabelDetails) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList struct {
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-2837hfd91h34dbg87364g*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1667460287386
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1667460287386
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The port ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service port number.
	//
	// example:
	//
	// 8080
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The traffic policy.
	TrafficPolicy *TrafficPolicy `json:"TrafficPolicy,omitempty" xml:"TrafficPolicy,omitempty"`
}

func (s GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) GetTrafficPolicy() *TrafficPolicy {
	return s.TrafficPolicy
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetGatewayUniqueId(v string) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetGmtCreate(v string) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetGmtModified(v string) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetId(v int64) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.Id = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetServiceId(v int64) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetServicePort(v int32) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.ServicePort = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) SetTrafficPolicy(v *TrafficPolicy) *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList {
	s.TrafficPolicy = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataPortTrafficPolicyList) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyDataVersionDetails struct {
	// The number of instances.
	//
	// example:
	//
	// 1
	EndpointNum *int32 `json:"EndpointNum,omitempty" xml:"EndpointNum,omitempty"`
	// The percentage of instances.
	//
	// example:
	//
	// 20%
	EndpointNumPercent *string `json:"EndpointNumPercent,omitempty" xml:"EndpointNumPercent,omitempty"`
	// The service version.
	ServiceVersion *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty" type:"Struct"`
}

func (s GetGatewayServiceDetailResponseBodyDataVersionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyDataVersionDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) GetEndpointNum() *int32 {
	return s.EndpointNum
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) GetEndpointNumPercent() *string {
	return s.EndpointNumPercent
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) GetServiceVersion() *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion {
	return s.ServiceVersion
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) SetEndpointNum(v int32) *GetGatewayServiceDetailResponseBodyDataVersionDetails {
	s.EndpointNum = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) SetEndpointNumPercent(v string) *GetGatewayServiceDetailResponseBodyDataVersionDetails {
	s.EndpointNumPercent = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) SetServiceVersion(v *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) *GetGatewayServiceDetailResponseBodyDataVersionDetails {
	s.ServiceVersion = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetails) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion struct {
	// The tag.
	Labels []*GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The version number.
	//
	// example:
	//
	// v2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) GetLabels() []*GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels {
	return s.Labels
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) GetName() *string {
	return s.Name
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) SetLabels(v []*GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion {
	s.Labels = v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) SetName(v string) *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion {
	s.Name = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersion) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels struct {
	// The tag.
	//
	// example:
	//
	// version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) GetKey() *string {
	return s.Key
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) GetValue() *string {
	return s.Value
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) SetKey(v string) *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels {
	s.Key = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) SetValue(v string) *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels {
	s.Value = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersionDetailsServiceVersionLabels) Validate() error {
	return dara.Validate(s)
}

type GetGatewayServiceDetailResponseBodyDataVersions struct {
	// The tag.
	//
	// example:
	//
	// version
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The type.
	//
	// example:
	//
	// test
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetGatewayServiceDetailResponseBodyDataVersions) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayServiceDetailResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) GetLabel() *string {
	return s.Label
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) GetType() *string {
	return s.Type
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) GetValue() *string {
	return s.Value
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) SetLabel(v string) *GetGatewayServiceDetailResponseBodyDataVersions {
	s.Label = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) SetType(v string) *GetGatewayServiceDetailResponseBodyDataVersions {
	s.Type = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) SetValue(v string) *GetGatewayServiceDetailResponseBodyDataVersions {
	s.Value = &v
	return s
}

func (s *GetGatewayServiceDetailResponseBodyDataVersions) Validate() error {
	return dara.Validate(s)
}
