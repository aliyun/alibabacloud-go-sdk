// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCollectorResponseBody
	GetRequestId() *string
	SetResult(v *UpdateCollectorResponseBodyResult) *UpdateCollectorResponseBody
	GetResult() *UpdateCollectorResponseBodyResult
}

type UpdateCollectorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *UpdateCollectorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCollectorResponseBody) GetResult() *UpdateCollectorResponseBodyResult {
	return s.Result
}

func (s *UpdateCollectorResponseBody) SetRequestId(v string) *UpdateCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCollectorResponseBody) SetResult(v *UpdateCollectorResponseBodyResult) *UpdateCollectorResponseBody {
	s.Result = v
	return s
}

func (s *UpdateCollectorResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCollectorResponseBodyResult struct {
	CollectorPaths []*string `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	// The configuration file information of the collector.
	Configs []*UpdateCollectorResponseBodyResultConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// Indicates whether the collector is validated and created. Valid values:
	//
	// - true: Only validated, not created.
	//
	// - false: Validated and created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The extended parameter information.
	ExtendConfigs []*UpdateCollectorResponseBodyResultExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
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
	// ct-cn-0v3xj86085dvq****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
	// The collector type. Valid values: fileBeat, metricBeat, heartBeat, and auditBeat.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// The collector version.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// The collector status. Valid values:
	//
	// - activing: Taking effect.
	//
	// - active: Active.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the virtual private cloud (VPC) where the collector resides.
	//
	// example:
	//
	// vpc-bp16k1dvzxtma*****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateCollectorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResult) GetCollectorPaths() []*string {
	return s.CollectorPaths
}

func (s *UpdateCollectorResponseBodyResult) GetConfigs() []*UpdateCollectorResponseBodyResultConfigs {
	return s.Configs
}

func (s *UpdateCollectorResponseBodyResult) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCollectorResponseBodyResult) GetExtendConfigs() []*UpdateCollectorResponseBodyResultExtendConfigs {
	return s.ExtendConfigs
}

func (s *UpdateCollectorResponseBodyResult) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *UpdateCollectorResponseBodyResult) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *UpdateCollectorResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateCollectorResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpdateCollectorResponseBodyResult) GetResId() *string {
	return s.ResId
}

func (s *UpdateCollectorResponseBodyResult) GetResType() *string {
	return s.ResType
}

func (s *UpdateCollectorResponseBodyResult) GetResVersion() *string {
	return s.ResVersion
}

func (s *UpdateCollectorResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *UpdateCollectorResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateCollectorResponseBodyResult) SetCollectorPaths(v []*string) *UpdateCollectorResponseBodyResult {
	s.CollectorPaths = v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetConfigs(v []*UpdateCollectorResponseBodyResultConfigs) *UpdateCollectorResponseBodyResult {
	s.Configs = v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetDryRun(v bool) *UpdateCollectorResponseBodyResult {
	s.DryRun = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetExtendConfigs(v []*UpdateCollectorResponseBodyResultExtendConfigs) *UpdateCollectorResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetGmtCreatedTime(v string) *UpdateCollectorResponseBodyResult {
	s.GmtCreatedTime = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetGmtUpdateTime(v string) *UpdateCollectorResponseBodyResult {
	s.GmtUpdateTime = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetName(v string) *UpdateCollectorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetOwnerId(v string) *UpdateCollectorResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetResId(v string) *UpdateCollectorResponseBodyResult {
	s.ResId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetResType(v string) *UpdateCollectorResponseBodyResult {
	s.ResType = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetResVersion(v string) *UpdateCollectorResponseBodyResult {
	s.ResVersion = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetStatus(v string) *UpdateCollectorResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) SetVpcId(v string) *UpdateCollectorResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResult) Validate() error {
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

type UpdateCollectorResponseBodyResultConfigs struct {
	// The file content.
	//
	// example:
	//
	// filebeat.inputs:xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The file name.
	//
	// example:
	//
	// filebeat.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s UpdateCollectorResponseBodyResultConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorResponseBodyResultConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResultConfigs) GetContent() *string {
	return s.Content
}

func (s *UpdateCollectorResponseBodyResultConfigs) GetFileName() *string {
	return s.FileName
}

func (s *UpdateCollectorResponseBodyResultConfigs) SetContent(v string) *UpdateCollectorResponseBodyResultConfigs {
	s.Content = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultConfigs) SetFileName(v string) *UpdateCollectorResponseBodyResultConfigs {
	s.FileName = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultConfigs) Validate() error {
	return dara.Validate(s)
}

type UpdateCollectorResponseBodyResultExtendConfigs struct {
	// The configuration type. Valid values:
	//
	// - collectorTargetInstance: the collector Output.
	//
	// - collectorDeployMachine: the machine on which the collector is deployed.
	//
	// - collectorElasticsearchForKibana: the Elasticsearch instance information that supports Kibana Dashboard.
	//
	// example:
	//
	// collectorDeployMachine
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// Indicates whether Monitoring is enabled. This parameter is displayed when **configType*	- is set to **collectorTargetInstance*	- and **instanceType*	- is set to **elasticsearch**. Valid values: true (enabled) and false (disabled).
	//
	// example:
	//
	// true
	EnableMonitoring *bool `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	// The machine group ID. This parameter is displayed when **configType*	- is set to **collectorDeployMachine**.
	//
	// example:
	//
	// default_ct-cn-5i2l75bz4776****
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The internal-facing access address of Kibana on the private network after Kibana Dashboard is enabled. This parameter is displayed when **configType*	- is set to **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// es-cn-n6w1o1x0w001c****-kibana.internal.elasticsearch.aliyuncs.com:5601
	Host  *string   `json:"host,omitempty" xml:"host,omitempty"`
	Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	// The ID of the instance associated with the collector. When **configType*	- is set to **collectorTargetInstance**, this parameter indicates the instance ID of the collector Output. When **configType*	- is set to **collectorDeployMachines*	- and **type*	- is set to **ACKCluster**, this parameter indicates the ACK (Container Kubernetes) cluster ID.
	//
	// example:
	//
	// es-cn-nif1z89fz003i****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The type of instance specified by the collector Output. Valid values: elasticsearch and logstash. This parameter is displayed when **configType*	- is set to **collectorTargetInstance**.
	//
	// example:
	//
	// elasticsearch
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The public network access address of Kibana after Kibana Dashboard is enabled. This parameter is displayed when **configType*	- is set to **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// https://es-cn-nif1z89fz003i****.kibana.elasticsearch.aliyuncs.com:5601
	KibanaHost *string `json:"kibanaHost,omitempty" xml:"kibanaHost,omitempty"`
	// Specific to the collectorDeployMachine type:
	//
	// The information about the ECS instances or ACK clusters on which the collector is deployed.
	Machines []*UpdateCollectorResponseBodyResultExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The transport protocol, which must be consistent with the access protocol of the instance specified by the collector Output. Valid values: HTTP and HTTPS. This parameter is displayed when **configType*	- is set to **collectorTargetInstance**.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The number of pods that are successfully collected in the ACK cluster. This parameter is displayed when **configType*	- is set to **collectorDeployMachines*	- and **type*	- is set to **ACKCluster**.
	//
	// example:
	//
	// 8
	SuccessPodsCount *string `json:"successPodsCount,omitempty" xml:"successPodsCount,omitempty"`
	// The total number of pods collected in the ACK cluster. This parameter is displayed when **configType*	- is set to **collectorDeployMachines*	- and **type*	- is set to **ACKCluster**.
	//
	// example:
	//
	// 10
	TotalPodsCount *string `json:"totalPodsCount,omitempty" xml:"totalPodsCount,omitempty"`
	// The type of machine on which the collector is deployed. This parameter is displayed when **configType*	- is set to **collectorDeployMachine**. Valid values:
	//
	// - ECSInstanceId: ECS.
	//
	// - ACKCluster: Container Kubernetes.
	//
	// example:
	//
	// ECSInstanceId
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The username used to access the instance specified by the collector Output. Default value: elastic. This parameter is displayed when **configType*	- is set to **collectorTargetInstance*	- or **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s UpdateCollectorResponseBodyResultExtendConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorResponseBodyResultExtendConfigs) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetEnableMonitoring() *bool {
	return s.EnableMonitoring
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetHost() *string {
	return s.Host
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetHosts() []*string {
	return s.Hosts
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetKibanaHost() *string {
	return s.KibanaHost
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetMachines() []*UpdateCollectorResponseBodyResultExtendConfigsMachines {
	return s.Machines
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetSuccessPodsCount() *string {
	return s.SuccessPodsCount
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetTotalPodsCount() *string {
	return s.TotalPodsCount
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetType() *string {
	return s.Type
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) GetUserName() *string {
	return s.UserName
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetConfigType(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetEnableMonitoring(v bool) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetGroupId(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetHost(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Host = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetHosts(v []*string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Hosts = v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetInstanceId(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetInstanceType(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetKibanaHost(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.KibanaHost = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetMachines(v []*UpdateCollectorResponseBodyResultExtendConfigsMachines) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Machines = v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetProtocol(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetSuccessPodsCount(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.SuccessPodsCount = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetTotalPodsCount(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.TotalPodsCount = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetType(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.Type = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) SetUserName(v string) *UpdateCollectorResponseBodyResultExtendConfigs {
	s.UserName = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigs) Validate() error {
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

type UpdateCollectorResponseBodyResultExtendConfigsMachines struct {
	// The status of each collector on the ECS instance. Valid values:
	//
	// - heartOk: The heartbeat is normal.
	//
	// - heartLost: The heartbeat is abnormal.
	//
	// - uninstalled: Not installed.
	//
	// - failed: Installation failed.
	//
	// example:
	//
	// heartOk
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// The list of ECS instance IDs.
	//
	// example:
	//
	// i-bp13y63575oypr9d****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateCollectorResponseBodyResultExtendConfigsMachines) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorResponseBodyResultExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) SetAgentStatus(v string) *UpdateCollectorResponseBodyResultExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) SetInstanceId(v string) *UpdateCollectorResponseBodyResultExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

func (s *UpdateCollectorResponseBodyResultExtendConfigsMachines) Validate() error {
	return dara.Validate(s)
}
