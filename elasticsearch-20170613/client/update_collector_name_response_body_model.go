// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectorNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCollectorNameResponseBody
	GetRequestId() *string
	SetResult(v *UpdateCollectorNameResponseBodyResult) *UpdateCollectorNameResponseBody
	GetResult() *UpdateCollectorNameResponseBodyResult
}

type UpdateCollectorNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *UpdateCollectorNameResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateCollectorNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCollectorNameResponseBody) GetResult() *UpdateCollectorNameResponseBodyResult {
	return s.Result
}

func (s *UpdateCollectorNameResponseBody) SetRequestId(v string) *UpdateCollectorNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCollectorNameResponseBody) SetResult(v *UpdateCollectorNameResponseBodyResult) *UpdateCollectorNameResponseBody {
	s.Result = v
	return s
}

func (s *UpdateCollectorNameResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCollectorNameResponseBodyResult struct {
	CollectorPaths []*string `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	// The configuration file information of the collector.
	Configs []*UpdateCollectorNameResponseBodyResultConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// Indicates whether the collector is validated only without being created. Valid values:
	//
	// - true: Only validates without updating.
	//
	// - false: Validates and updates.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The extended configurations of the collector.
	ExtendConfigs []*UpdateCollectorNameResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// The time when the collector was created.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	// The time when the collector was last updated.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtUpdateTime *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	// The collector name.
	//
	// example:
	//
	// ct-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 16852099488*****
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// The collector instance ID.
	//
	// example:
	//
	// ct-cn-77uqof2s7rg5c****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
	// The collector type. Valid values: fileBeat, metricBeat, heartBeat, and audiBeat.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// The collector version. The supported versions depend on the type of machine on which the collector is deployed:
	//
	// - ECS: 6.8.5_with_community
	//
	// - ACK: 6.8.13_with_community.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// The collector status. Valid values: activing (taking effect) and active (active).
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the VPC where the collector resides.
	//
	// example:
	//
	// vpc-bp16k1dvzxtma*****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResult) GetCollectorPaths() []*string {
	return s.CollectorPaths
}

func (s *UpdateCollectorNameResponseBodyResult) GetConfigs() []*UpdateCollectorNameResponseBodyResultConfigs {
	return s.Configs
}

func (s *UpdateCollectorNameResponseBodyResult) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCollectorNameResponseBodyResult) GetExtendConfigs() []*UpdateCollectorNameResponseBodyResultExtendConfigs {
	return s.ExtendConfigs
}

func (s *UpdateCollectorNameResponseBodyResult) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *UpdateCollectorNameResponseBodyResult) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *UpdateCollectorNameResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateCollectorNameResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpdateCollectorNameResponseBodyResult) GetResId() *string {
	return s.ResId
}

func (s *UpdateCollectorNameResponseBodyResult) GetResType() *string {
	return s.ResType
}

func (s *UpdateCollectorNameResponseBodyResult) GetResVersion() *string {
	return s.ResVersion
}

func (s *UpdateCollectorNameResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *UpdateCollectorNameResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateCollectorNameResponseBodyResult) SetCollectorPaths(v []*string) *UpdateCollectorNameResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetConfigs(v []*UpdateCollectorNameResponseBodyResultConfigs) *UpdateCollectorNameResponseBodyResult {
	s.Configs = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetDryRun(v bool) *UpdateCollectorNameResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetExtendConfigs(v []*UpdateCollectorNameResponseBodyResultExtendConfigs) *UpdateCollectorNameResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetGmtCreatedTime(v string) *UpdateCollectorNameResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetGmtUpdateTime(v string) *UpdateCollectorNameResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetName(v string) *UpdateCollectorNameResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetOwnerId(v string) *UpdateCollectorNameResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetResId(v string) *UpdateCollectorNameResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetResType(v string) *UpdateCollectorNameResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetResVersion(v string) *UpdateCollectorNameResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetStatus(v string) *UpdateCollectorNameResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) SetVpcId(v string) *UpdateCollectorNameResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResult) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExtendConfigs != nil {
		for _, item := range s.ExtendConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCollectorNameResponseBodyResultConfigs struct {
	// The file content.
	//
	// example:
	//
	// - key: log\\n title: Log file content\\n description: >\\n Contains log file lines.\\n ....
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The file name.
	//
	// example:
	//
	// fields.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResultConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) GetContent() *string {
	return s.Content
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) GetFileName() *string {
	return s.FileName
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) SetContent(v string) *UpdateCollectorNameResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) SetFileName(v string) *UpdateCollectorNameResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultConfigs) Validate() error {
	return dara.Validate(s)
}

type UpdateCollectorNameResponseBodyResultExtendConfigs struct {
	// The configuration type. Valid values:
	//
	// - collectorTargetInstance: the collector Output.
	//
	// - collectorDeployMachine: the machine on which the collector is deployed.
	//
	// - collectorElasticsearchForKibana: the Elasticsearch instance that supports Kibana Dashboard.
	//
	// example:
	//
	// collectorDeployMachine
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// Indicates whether Monitoring is enabled. Displayed when **configType*	- is **collectorTargetInstance*	- and **instanceType*	- is **elasticsearch**. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	EnableMonitoring *bool `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	// The machine group ID. Displayed when **configType*	- is **collectorDeployMachine**.
	//
	// example:
	//
	// default_ct-cn-5i2l75bz4776****
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The internal-facing access address of Kibana on the private network after Kibana Dashboard is enabled. Displayed when **configType*	- is **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// es-cn-4591jumei000u****-kibana.internal.elasticsearch.aliyuncs.com:5601
	Host  *string   `json:"host,omitempty" xml:"host,omitempty"`
	Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	// The ID of the instance associated with the collector. When **configType*	- is **collectorTargetInstance**, this is the instance ID of the collector Output. When **configType*	- is **collectorDeployMachines*	- and **type*	- is **ACKCluster**, this is the ACK cluster ID.
	//
	// example:
	//
	// es-cn-n6w1o1****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The type of instance specified in the collector Output. Valid values: elasticsearch and logstash. Displayed when **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// elasticsearch
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The public network access address of Kibana after Kibana Dashboard is enabled. Displayed when **configType*	- is **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// https://es-cn-4591jumei000u****.kibana.elasticsearch.aliyuncs.com:5601
	KibanaHost *string `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	// The list of ECS machines on which the collector is deployed. Displayed when **configType*	- is **collectorDeployMachines*	- and **type*	- is **ECSInstanceId**.
	Machines []*UpdateCollectorNameResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The transmission protocol. Valid values: **HTTP*	- and **HTTPS**.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The number of pods that are successfully collected in the ACK cluster. Displayed when **configType*	- is **collectorDeployMachines*	- and **type*	- is **ACKCluster**.
	//
	// example:
	//
	// 8
	SuccessPodsCount *string `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	// The total number of pods collected in the ACK cluster. Displayed when **configType*	- is **collectorDeployMachines*	- and **type*	- is **ACKCluster**.
	//
	// example:
	//
	// 10
	TotalPodsCount *string `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	// The type of machine on which the collector is deployed. Displayed when **configType*	- is **collectorDeployMachine**. Valid values:
	//
	// - ECSInstanceId: ECS instance.
	//
	// - ACKCluster: Container Kubernetes cluster.
	//
	// example:
	//
	// ECSInstanceId
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The username used to access the instance specified in the collector Output. Default value: elastic. Displayed when **configType*	- is **collectorTargetInstance*	- or **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetEnableMonitoring() *bool {
	return s.EnableMonitoring
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetHost() *string {
	return s.Host
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetHosts() []*string {
	return s.Hosts
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetKibanaHost() *string {
	return s.KibanaHost
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetMachines() []*UpdateCollectorNameResponseBodyResultExtendConfigsMachines {
	return s.Machines
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetSuccessPodsCount() *string {
	return s.SuccessPodsCount
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetTotalPodsCount() *string {
	return s.TotalPodsCount
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetType() *string {
	return s.Type
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) GetUserName() *string {
	return s.UserName
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetConfigType(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetGroupId(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetHost(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetHosts(v []*string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetInstanceId(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetInstanceType(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetKibanaHost(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetMachines(v []*UpdateCollectorNameResponseBodyResultExtendConfigsMachines) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetProtocol(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetType(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) SetUserName(v string) *UpdateCollectorNameResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigs) Validate() error {
	if s.Machines != nil {
		for _, item := range s.Machines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCollectorNameResponseBodyResultExtendConfigsMachines struct {
	// The status of the collector on the ECS instance. Valid values: **heartOk*	- (normal heartbeat), **heartLost*	- (abnormal heartbeat), **uninstalled*	- (not installed), and **failed*	- (installation failed).
	//
	// example:
	//
	// heartOk
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// The list of ECS machine IDs.
	//
	// example:
	//
	// c1b9fde5172b84f82b9928e825a7b8988
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigsMachines) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorNameResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *UpdateCollectorNameResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *UpdateCollectorNameResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

func (s *UpdateCollectorNameResponseBodyResultExtendConfigsMachines) Validate() error {
	return dara.Validate(s)
}
