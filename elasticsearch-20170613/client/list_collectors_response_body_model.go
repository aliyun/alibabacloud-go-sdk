// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListCollectorsResponseBodyHeaders) *ListCollectorsResponseBody
	GetHeaders() *ListCollectorsResponseBodyHeaders
	SetRequestId(v string) *ListCollectorsResponseBody
	GetRequestId() *string
	SetResult(v []*ListCollectorsResponseBodyResult) *ListCollectorsResponseBody
	GetResult() []*ListCollectorsResponseBodyResult
}

type ListCollectorsResponseBody struct {
	// The header of the response.
	Headers *ListCollectorsResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListCollectorsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCollectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBody) GetHeaders() *ListCollectorsResponseBodyHeaders {
	return s.Headers
}

func (s *ListCollectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCollectorsResponseBody) GetResult() []*ListCollectorsResponseBodyResult {
	return s.Result
}

func (s *ListCollectorsResponseBody) SetHeaders(v *ListCollectorsResponseBodyHeaders) *ListCollectorsResponseBody {
	s.Headers = v
	return s
}

func (s *ListCollectorsResponseBody) SetRequestId(v string) *ListCollectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCollectorsResponseBody) SetResult(v []*ListCollectorsResponseBodyResult) *ListCollectorsResponseBody {
	s.Result = v
	return s
}

func (s *ListCollectorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCollectorsResponseBodyHeaders struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListCollectorsResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListCollectorsResponseBodyHeaders) SetXTotalCount(v int32) *ListCollectorsResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListCollectorsResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListCollectorsResponseBodyResult struct {
	CollectorPaths []*string `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	// The information about the configuration file of the shipper.
	Configs []*ListCollectorsResponseBodyResultConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// Indicates whether a dry run is performed. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The extended configurations of the shipper.
	ExtendConfigs []*ListCollectorsResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// The time when the shipper was created.
	//
	// example:
	//
	// 2020-08-18T02:06:12.000+0000
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	// The time when the shipper was updated.
	//
	// example:
	//
	// 2020-08-18T09:40:43.000+0000
	GmtUpdateTime *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	// The name of the shipper.
	//
	// example:
	//
	// FileBeat001
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 168520994880****
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// The ID of the shipper.
	//
	// example:
	//
	// ct-cn-0v3xj86085dvq****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
	// The type of the shipper. Valid values: fileBeat, metricBeat, heartBeat, and auditBeat.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// The version of the shipper.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// The status of the shipper. Valid values:
	//
	// 	- activating
	//
	// 	- active
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the virtual private cloud (VPC) where the shipper resides.
	//
	// example:
	//
	// vpc-bp16k1dvzxtma*****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListCollectorsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResult) GetCollectorPaths() []*string {
	return s.CollectorPaths
}

func (s *ListCollectorsResponseBodyResult) GetConfigs() []*ListCollectorsResponseBodyResultConfigs {
	return s.Configs
}

func (s *ListCollectorsResponseBodyResult) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListCollectorsResponseBodyResult) GetExtendConfigs() []*ListCollectorsResponseBodyResultExtendConfigs {
	return s.ExtendConfigs
}

func (s *ListCollectorsResponseBodyResult) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *ListCollectorsResponseBodyResult) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *ListCollectorsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListCollectorsResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListCollectorsResponseBodyResult) GetResId() *string {
	return s.ResId
}

func (s *ListCollectorsResponseBodyResult) GetResType() *string {
	return s.ResType
}

func (s *ListCollectorsResponseBodyResult) GetResVersion() *string {
	return s.ResVersion
}

func (s *ListCollectorsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListCollectorsResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *ListCollectorsResponseBodyResult) SetCollectorPaths(v []*string) *ListCollectorsResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetConfigs(v []*ListCollectorsResponseBodyResultConfigs) *ListCollectorsResponseBodyResult {
	s.Configs = v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetDryRun(v bool) *ListCollectorsResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetExtendConfigs(v []*ListCollectorsResponseBodyResultExtendConfigs) *ListCollectorsResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetGmtCreatedTime(v string) *ListCollectorsResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetGmtUpdateTime(v string) *ListCollectorsResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetName(v string) *ListCollectorsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetOwnerId(v string) *ListCollectorsResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetResId(v string) *ListCollectorsResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetResType(v string) *ListCollectorsResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetResVersion(v string) *ListCollectorsResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetStatus(v string) *ListCollectorsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) SetVpcId(v string) *ListCollectorsResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *ListCollectorsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListCollectorsResponseBodyResultConfigs struct {
	// The content of the file.
	//
	// example:
	//
	// - key: log\\n title: Log file content\\n description: >\\n Contains log file lines.\\n ....
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// fields.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ListCollectorsResponseBodyResultConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResultConfigs) GetContent() *string {
	return s.Content
}

func (s *ListCollectorsResponseBodyResultConfigs) GetFileName() *string {
	return s.FileName
}

func (s *ListCollectorsResponseBodyResultConfigs) SetContent(v string) *ListCollectorsResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *ListCollectorsResponseBodyResultConfigs) SetFileName(v string) *ListCollectorsResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

func (s *ListCollectorsResponseBodyResultConfigs) Validate() error {
	return dara.Validate(s)
}

type ListCollectorsResponseBodyResultExtendConfigs struct {
	// The configuration type. Valid values:
	//
	// 	- collectorTargetInstance
	//
	// 	- collectorDeployMachine
	//
	// 	- collectorElasticsearchForKibana
	//
	// example:
	//
	// collectorDeployMachine
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// Indicates whether monitoring is enabled. This parameter is returned if the value of **configType*	- is **collectorTargetInstance*	- and the value of **instanceType*	- is **elasticsearch**. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableMonitoring *bool `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	// The ID of the machine group. This parameter is returned if the value of **configType*	- is **collectorDeployMachine**.
	//
	// example:
	//
	// default_ct-cn-5i2l75bz4776****
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The internal endpoint of Kibana after you enable the Kibana dashboard. This parameter is returned if the value of **configType*	- is **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****-kibana.internal.elasticsearch.aliyuncs.com:5601
	Host  *string   `json:"host,omitempty" xml:"host,omitempty"`
	Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	// The ID of the resource with which the shipper is associated. If the value of **configType*	- is **collectorTargetInstance**, the value of this parameter is the ID of the resource specified in the output configuration part of the shipper. If the value of **configType*	- is **collectorDeployMachine*	- and the value of **type*	- is **ACKCluster**, the value of this parameter is the ID of the ACK cluster.
	//
	// example:
	//
	// es-cn-nif1z89fz003i****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The type of the cluster specified in the output configuration part of the shipper. Valid values: elasticsearch and logstash. This parameter is returned if the value of **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// elasticsearch
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The public endpoint of Kibana after you enable the Kibana dashboard. This parameter is returned if the value of **configType*	- is **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// https://es-cn-nif1z89fz003i****.kibana.elasticsearch.aliyuncs.com:5601
	KibanaHost *string `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	// The information about the ECS instances on which the shipper is deployed. This parameter is returned if the value of **configType*	- is **collectorDeployMachine*	- and the value of **type*	- is **ECSInstanceId**.
	Machines []*ListCollectorsResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The transmission protocol, which must be the same as the access protocol of the resource specified in the output configuration part of the shipper. Valid values: HTTP and HTTPS. This parameter is returned if the value of **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The number of pods from which data is successfully collected in the ACK cluster. This parameter is returned if the value of **configType*	- is **collectorDeployMachine*	- and the value of **type*	- is **ACKCluster**.
	//
	// example:
	//
	// 8
	SuccessPodsCount *string `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	// The total number of pods from which data is collected in the ACK cluster. This parameter is returned if the value of **configType*	- is **collectorDeployMachine*	- and the value of **type*	- is **ACKCluster**.
	//
	// example:
	//
	// 10
	TotalPodsCount *string `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	// The type of the machine on which the shipper is deployed. This parameter is returned if the value of **configType*	- is **collectorDeployMachine**. Valid values:
	//
	// 	- ECSInstanceId
	//
	// 	- ACKCluster
	//
	// example:
	//
	// ECSInstanceId
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The username that is used to access the resource specified in the output configuration part of the shipper. The default value is elastic. This parameter is returned if the value of **configType*	- is **collectorTargetInstance*	- or **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListCollectorsResponseBodyResultExtendConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetEnableMonitoring() *bool {
	return s.EnableMonitoring
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetGroupId() *string {
	return s.GroupId
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetHost() *string {
	return s.Host
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetHosts() []*string {
	return s.Hosts
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetKibanaHost() *string {
	return s.KibanaHost
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetMachines() []*ListCollectorsResponseBodyResultExtendConfigsMachines {
	return s.Machines
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetSuccessPodsCount() *string {
	return s.SuccessPodsCount
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetTotalPodsCount() *string {
	return s.TotalPodsCount
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetType() *string {
	return s.Type
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) GetUserName() *string {
	return s.UserName
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetConfigType(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *ListCollectorsResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetGroupId(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetHost(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetHosts(v []*string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetInstanceId(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetInstanceType(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetKibanaHost(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetMachines(v []*ListCollectorsResponseBodyResultExtendConfigsMachines) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetProtocol(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetType(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) SetUserName(v string) *ListCollectorsResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigs) Validate() error {
	return dara.Validate(s)
}

type ListCollectorsResponseBodyResultExtendConfigsMachines struct {
	// The status of the shipper on the ECS instance. Valid values:
	//
	// 	- heartOk
	//
	// 	- heartLost
	//
	// 	- uninstalled
	//
	// 	- failed
	//
	// example:
	//
	// heartOk
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// The IDs of the ECS instances.
	//
	// example:
	//
	// i-bp13y63575oypr9d****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListCollectorsResponseBodyResultExtendConfigsMachines) String() string {
	return dara.Prettify(s)
}

func (s ListCollectorsResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *ListCollectorsResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *ListCollectorsResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

func (s *ListCollectorsResponseBodyResultExtendConfigsMachines) Validate() error {
	return dara.Validate(s)
}
