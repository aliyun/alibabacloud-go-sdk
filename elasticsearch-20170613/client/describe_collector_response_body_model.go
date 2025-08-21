// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCollectorResponseBody
	GetRequestId() *string
	SetResult(v *DescribeCollectorResponseBodyResult) *DescribeCollectorResponseBody
	GetResult() *DescribeCollectorResponseBodyResult
}

type DescribeCollectorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DescribeCollectorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCollectorResponseBody) GetResult() *DescribeCollectorResponseBodyResult {
	return s.Result
}

func (s *DescribeCollectorResponseBody) SetRequestId(v string) *DescribeCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCollectorResponseBody) SetResult(v *DescribeCollectorResponseBodyResult) *DescribeCollectorResponseBody {
	s.Result = v
	return s
}

func (s *DescribeCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCollectorResponseBodyResult struct {
	CollectorPaths []*string `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	// The information about the configuration file of the shipper.
	Configs []*DescribeCollectorResponseBodyResultConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
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
	ExtendConfigs []*DescribeCollectorResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// The time when the shipper was created.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	// The time when the shipper was updated.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtUpdateTime *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	// The name of the shipper.
	//
	// example:
	//
	// ct-cn-4135is2tj194p****
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 16852099488*****
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// The ID of the shipper.
	//
	// example:
	//
	// ct-cn-rg31ahn82m0qd****
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

func (s DescribeCollectorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResult) GetCollectorPaths() []*string {
	return s.CollectorPaths
}

func (s *DescribeCollectorResponseBodyResult) GetConfigs() []*DescribeCollectorResponseBodyResultConfigs {
	return s.Configs
}

func (s *DescribeCollectorResponseBodyResult) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeCollectorResponseBodyResult) GetExtendConfigs() []*DescribeCollectorResponseBodyResultExtendConfigs {
	return s.ExtendConfigs
}

func (s *DescribeCollectorResponseBodyResult) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *DescribeCollectorResponseBodyResult) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *DescribeCollectorResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeCollectorResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeCollectorResponseBodyResult) GetResId() *string {
	return s.ResId
}

func (s *DescribeCollectorResponseBodyResult) GetResType() *string {
	return s.ResType
}

func (s *DescribeCollectorResponseBodyResult) GetResVersion() *string {
	return s.ResVersion
}

func (s *DescribeCollectorResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeCollectorResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeCollectorResponseBodyResult) SetCollectorPaths(v []*string) *DescribeCollectorResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetConfigs(v []*DescribeCollectorResponseBodyResultConfigs) *DescribeCollectorResponseBodyResult {
	s.Configs = v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetDryRun(v bool) *DescribeCollectorResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetExtendConfigs(v []*DescribeCollectorResponseBodyResultExtendConfigs) *DescribeCollectorResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetGmtCreatedTime(v string) *DescribeCollectorResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetGmtUpdateTime(v string) *DescribeCollectorResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetName(v string) *DescribeCollectorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetOwnerId(v string) *DescribeCollectorResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetResId(v string) *DescribeCollectorResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetResType(v string) *DescribeCollectorResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetResVersion(v string) *DescribeCollectorResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetStatus(v string) *DescribeCollectorResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) SetVpcId(v string) *DescribeCollectorResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type DescribeCollectorResponseBodyResultConfigs struct {
	// The content of the file.
	//
	// example:
	//
	// fileBeat.inputs:xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// filebeat.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s DescribeCollectorResponseBodyResultConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectorResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResultConfigs) GetContent() *string {
	return s.Content
}

func (s *DescribeCollectorResponseBodyResultConfigs) GetFileName() *string {
	return s.FileName
}

func (s *DescribeCollectorResponseBodyResultConfigs) SetContent(v string) *DescribeCollectorResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultConfigs) SetFileName(v string) *DescribeCollectorResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeCollectorResponseBodyResultExtendConfigs struct {
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
	// Indicates whether monitoring is enabled. This parameter is returned if the value of **configType*	- is **collectorTargetInstance**. Valid values:
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
	// The private endpoint of Kibana after you enable the Kibana dashboard. This parameter is returned if the value of **configType*	- is **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****-kibana.internal.elasticsearch.aliyuncs.com:5601
	Host  *string   `json:"host,omitempty" xml:"host,omitempty"`
	Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	// The ID of the resource that is associated with the shipper. If the value of **configType*	- is **collectorTargetInstance**, the value of this parameter is the ID of the resource specified in the output configuration part of the shipper. If the value of **configType*	- is **collectorDeployMachines*	- and the value of **type*	- is **ACKCluster**, the value of this parameter is the ID of the ACK cluster.
	//
	// example:
	//
	// es-cn-n6w1o1****
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
	// The information about the Elastic Compute Service (ECS) instances on which the shipper is deployed. This parameter is returned if the value of **configType*	- is **collectorDeployMachines*	- and the value of **type*	- is **ECSInstanceId**.
	Machines []*DescribeCollectorResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The transmission protocol, which must be the same as the access protocol of the resource specified in the output configuration part of the shipper. Valid values: HTTP and HTTPS. This parameter is returned if the value of **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The number of pods from which data is succcessfully collected in the Container Service for Kubernetes (ACK) cluster.
	//
	// example:
	//
	// 8
	SuccessPodsCount *string `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	// The total number of pods from which data is collected in the ACK cluster.
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

func (s DescribeCollectorResponseBodyResultExtendConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectorResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetEnableMonitoring() *bool {
	return s.EnableMonitoring
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetHost() *string {
	return s.Host
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetHosts() []*string {
	return s.Hosts
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetKibanaHost() *string {
	return s.KibanaHost
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetMachines() []*DescribeCollectorResponseBodyResultExtendConfigsMachines {
	return s.Machines
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetSuccessPodsCount() *string {
	return s.SuccessPodsCount
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetTotalPodsCount() *string {
	return s.TotalPodsCount
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetType() *string {
	return s.Type
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) GetUserName() *string {
	return s.UserName
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetConfigType(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetGroupId(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetHost(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetHosts(v []*string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetInstanceId(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetInstanceType(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetKibanaHost(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetMachines(v []*DescribeCollectorResponseBodyResultExtendConfigsMachines) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetProtocol(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetType(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) SetUserName(v string) *DescribeCollectorResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeCollectorResponseBodyResultExtendConfigsMachines struct {
	// The status of the shipper on the ECS instance. Valid values:
	//
	// 	- heartOk: The heartbeat is normal.
	//
	// 	- heartLost: The heartbeat is abnormal.
	//
	// 	- uninstalled: The shipper is not installed.
	//
	// 	- failed: The shipper fails to be installed.
	//
	// example:
	//
	// heartOk
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// The IDs of the ECS instances.
	//
	// example:
	//
	// i-bp1gyhphjaj73jsr****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DescribeCollectorResponseBodyResultExtendConfigsMachines) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectorResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *DescribeCollectorResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *DescribeCollectorResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

func (s *DescribeCollectorResponseBodyResultExtendConfigsMachines) Validate() error {
	return dara.Validate(s)
}
