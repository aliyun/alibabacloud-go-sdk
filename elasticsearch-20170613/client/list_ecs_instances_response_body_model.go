// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListEcsInstancesResponseBodyHeaders) *ListEcsInstancesResponseBody
	GetHeaders() *ListEcsInstancesResponseBodyHeaders
	SetRequestId(v string) *ListEcsInstancesResponseBody
	GetRequestId() *string
	SetResult(v []*ListEcsInstancesResponseBodyResult) *ListEcsInstancesResponseBody
	GetResult() []*ListEcsInstancesResponseBodyResult
}

type ListEcsInstancesResponseBody struct {
	// The number of returned records.
	Headers *ListEcsInstancesResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The header of the response.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Cloud Assistant the installation status, support:
	//
	// 	- true: The Prometheus agent was installed.
	//
	// 	- false: The Prometheus agent was not installed.
	Result []*ListEcsInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListEcsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBody) GetHeaders() *ListEcsInstancesResponseBodyHeaders {
	return s.Headers
}

func (s *ListEcsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEcsInstancesResponseBody) GetResult() []*ListEcsInstancesResponseBodyResult {
	return s.Result
}

func (s *ListEcsInstancesResponseBody) SetHeaders(v *ListEcsInstancesResponseBodyHeaders) *ListEcsInstancesResponseBody {
	s.Headers = v
	return s
}

func (s *ListEcsInstancesResponseBody) SetRequestId(v string) *ListEcsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsInstancesResponseBody) SetResult(v []*ListEcsInstancesResponseBodyResult) *ListEcsInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListEcsInstancesResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEcsInstancesResponseBodyHeaders struct {
	// The returned data.
	//
	// example:
	//
	// 11
	XTotalCount *int32 `json:"X-Total-Count,omitempty" xml:"X-Total-Count,omitempty"`
}

func (s ListEcsInstancesResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyHeaders) GetXTotalCount() *int32 {
	return s.XTotalCount
}

func (s *ListEcsInstancesResponseBodyHeaders) SetXTotalCount(v int32) *ListEcsInstancesResponseBodyHeaders {
	s.XTotalCount = &v
	return s
}

func (s *ListEcsInstancesResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListEcsInstancesResponseBodyResult struct {
	// The name of the ECS instance.
	//
	// example:
	//
	// true
	CloudAssistantStatus *string `json:"cloudAssistantStatus,omitempty" xml:"cloudAssistantStatus,omitempty"`
	// The ID of the collector instance.
	Collectors []*ListEcsInstancesResponseBodyResultCollectors `json:"collectors,omitempty" xml:"collectors,omitempty" type:"Repeated"`
	// The tags of the ECS instance.
	//
	// example:
	//
	// i-bp14ncqge8wy3l3d****
	EcsInstanceId *string `json:"ecsInstanceId,omitempty" xml:"ecsInstanceId,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// ecsTestName
	EcsInstanceName *string `json:"ecsInstanceName,omitempty" xml:"ecsInstanceName,omitempty"`
	// The type of the IP address that is used by the instance. Valid values:
	//
	// 	- public: public endpoint
	//
	// 	- private: private network address
	IpAddress []*ListEcsInstancesResponseBodyResultIpAddress `json:"ipAddress,omitempty" xml:"ipAddress,omitempty" type:"Repeated"`
	// The status of the ECS instance. Valid values:
	//
	// 	- running: The master instance is running
	//
	// 	- starting
	//
	// 	- stopping: The task is being stopped.
	//
	// 	- stopped: The node is stopped.
	//
	// example:
	//
	// linux
	OsType *string `json:"osType,omitempty" xml:"osType,omitempty"`
	// The IP address of the ECS instance.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The operating system type of the ECS instance. Valid values:
	//
	// 	- windows:Windows operating system
	//
	// 	- linux:Linux operating system
	//
	// example:
	//
	// [ { "tagKey": "a", "tagValue": "b" } ]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListEcsInstancesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResult) GetCloudAssistantStatus() *string {
	return s.CloudAssistantStatus
}

func (s *ListEcsInstancesResponseBodyResult) GetCollectors() []*ListEcsInstancesResponseBodyResultCollectors {
	return s.Collectors
}

func (s *ListEcsInstancesResponseBodyResult) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ListEcsInstancesResponseBodyResult) GetEcsInstanceName() *string {
	return s.EcsInstanceName
}

func (s *ListEcsInstancesResponseBodyResult) GetIpAddress() []*ListEcsInstancesResponseBodyResultIpAddress {
	return s.IpAddress
}

func (s *ListEcsInstancesResponseBodyResult) GetOsType() *string {
	return s.OsType
}

func (s *ListEcsInstancesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListEcsInstancesResponseBodyResult) GetTags() *string {
	return s.Tags
}

func (s *ListEcsInstancesResponseBodyResult) SetCloudAssistantStatus(v string) *ListEcsInstancesResponseBodyResult {
	s.CloudAssistantStatus = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetCollectors(v []*ListEcsInstancesResponseBodyResultCollectors) *ListEcsInstancesResponseBodyResult {
	s.Collectors = v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetEcsInstanceId(v string) *ListEcsInstancesResponseBodyResult {
	s.EcsInstanceId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetEcsInstanceName(v string) *ListEcsInstancesResponseBodyResult {
	s.EcsInstanceName = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetIpAddress(v []*ListEcsInstancesResponseBodyResultIpAddress) *ListEcsInstancesResponseBodyResult {
	s.IpAddress = v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetOsType(v string) *ListEcsInstancesResponseBodyResult {
	s.OsType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetStatus(v string) *ListEcsInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) SetTags(v string) *ListEcsInstancesResponseBodyResult {
	s.Tags = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResult) Validate() error {
	if s.Collectors != nil {
		for _, item := range s.Collectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpAddress != nil {
		for _, item := range s.IpAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEcsInstancesResponseBodyResultCollectors struct {
	CollectorPaths []*string `json:"collectorPaths,omitempty" xml:"collectorPaths,omitempty" type:"Repeated"`
	// The content of the file.
	Configs []*ListEcsInstancesResponseBodyResultCollectorsConfigs `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Whether Monitoring is enabled. This field is displayed when the **configType*	- is **collectorTargetInstance*	- and the **instanceType*	- is **Elasticsearch**. Valid values:
	//
	// 	- true
	//
	// 	- false
	ExtendConfigs []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigs `json:"extendConfigs,omitempty" xml:"extendConfigs,omitempty" type:"Repeated"`
	// The status of the collector. Valid values:
	//
	// 	- activating: The project is taking effect.
	//
	// 	- active: The instance has taken effect.
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtCreatedTime *string `json:"gmtCreatedTime,omitempty" xml:"gmtCreatedTime,omitempty"`
	// Specifies whether to verify and create a crawer. Valid values:
	//
	// 	- true: only verifies and does not create a
	//
	// 	- false: verifies and creates a
	//
	// example:
	//
	// 2020-06-20T07:26:47.000+0000
	GmtUpdateTime *string `json:"gmtUpdateTime,omitempty" xml:"gmtUpdateTime,omitempty"`
	// The configuration file information of the collector.
	//
	// example:
	//
	// ct-testAbc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the Virtual Private Cloud to which the collector belongs.
	//
	// example:
	//
	// 16852***488*****
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// The time when the collector was updated.
	//
	// example:
	//
	// ct-cn-0v3xj86085dvq****
	ResId *string `json:"resId,omitempty" xml:"resId,omitempty"`
	// The version of the collector. If the machine type of the collector is ECS, only **6.8.5_with_community*	- is supported.
	//
	// example:
	//
	// fileBeat
	ResType *string `json:"resType,omitempty" xml:"resType,omitempty"`
	// The time when the crawl collector was created.
	//
	// example:
	//
	// 6.8.5_with_community
	ResVersion *string `json:"resVersion,omitempty" xml:"resVersion,omitempty"`
	// The name of the collector.
	//
	// example:
	//
	// activing
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the collector. FileBeat, metricBeat, heartBeat, and auditBeat are supported.
	//
	// example:
	//
	// vpc-bp16k1dvzxtm******
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectors) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectors) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetCollectorPaths() []*string {
	return s.CollectorPaths
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetConfigs() []*ListEcsInstancesResponseBodyResultCollectorsConfigs {
	return s.Configs
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetExtendConfigs() []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	return s.ExtendConfigs
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetGmtCreatedTime() *string {
	return s.GmtCreatedTime
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetGmtUpdateTime() *string {
	return s.GmtUpdateTime
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetName() *string {
	return s.Name
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetResId() *string {
	return s.ResId
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetResType() *string {
	return s.ResType
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetResVersion() *string {
	return s.ResVersion
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetStatus() *string {
	return s.Status
}

func (s *ListEcsInstancesResponseBodyResultCollectors) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetCollectorPaths(v []*string) *ListEcsInstancesResponseBodyResultCollectors {
	s.CollectorPaths = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetConfigs(v []*ListEcsInstancesResponseBodyResultCollectorsConfigs) *ListEcsInstancesResponseBodyResultCollectors {
	s.Configs = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetDryRun(v bool) *ListEcsInstancesResponseBodyResultCollectors {
	s.DryRun = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetExtendConfigs(v []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) *ListEcsInstancesResponseBodyResultCollectors {
	s.ExtendConfigs = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetGmtCreatedTime(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.GmtCreatedTime = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetGmtUpdateTime(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.GmtUpdateTime = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetName(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.Name = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetOwnerId(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.OwnerId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetResId(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.ResId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetResType(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.ResType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetResVersion(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.ResVersion = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetStatus(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.Status = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) SetVpcId(v string) *ListEcsInstancesResponseBodyResultCollectors {
	s.VpcId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectors) Validate() error {
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

type ListEcsInstancesResponseBodyResultCollectorsConfigs struct {
	// The name of the file.
	//
	// example:
	//
	// - key: log\\n title: Log file content\\n description: >\\n Contains log file lines.\\n ....
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The information about the extended parameter.
	//
	// example:
	//
	// fields.yml
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectorsConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectorsConfigs) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) GetContent() *string {
	return s.Content
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) GetFileName() *string {
	return s.FileName
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) SetContent(v string) *ListEcsInstancesResponseBodyResultCollectorsConfigs {
	s.Content = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) SetFileName(v string) *ListEcsInstancesResponseBodyResultCollectorsConfigs {
	s.FileName = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsConfigs) Validate() error {
	return dara.Validate(s)
}

type ListEcsInstancesResponseBodyResultCollectorsExtendConfigs struct {
	// The instance type specified by Collector Output. Supports Elasticsearch and Logstash. Displayed when the **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// collectorDeployMachine
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The ID of the host group. Displayed when the **configType*	- is **collectorDeployMachine**.
	//
	// example:
	//
	// true
	EnableMonitoring *bool `json:"enableMonitoring,omitempty" xml:"enableMonitoring,omitempty"`
	// The configuration type. Valid values:
	//
	// 	- collectorTargetInstance: Collector Output
	//
	// 	- collectorDeployMachine: Collector Deployment Machine
	//
	// 	- Collector Elasticsearch ForKibana: Elasticsearch instance information that supports the Kibana dashboard
	//
	// example:
	//
	// default_ct-cn-5i2l75bz4776****
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The path in which Filebeat is collected.
	Hosts []*string `json:"hosts,omitempty" xml:"hosts,omitempty" type:"Repeated"`
	// The list of ECS instances on which the collector is deployed. Displayed when the **configType*	- is **collectorDeployMachines*	- and the **type*	- is **ECSInstanceId**.
	//
	// example:
	//
	// es-cn-nif1z89fz003i****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The transmission protocol, which must be the same as the access protocol of the instance specified by Output. HTTP and HTTPS. Displayed when the **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// elasticsearch
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The status of each crawl on the ECS instance. Valid values:
	//
	// 	- heartOk: The heartbeat is normal.
	//
	// 	- heartLost: The heartbeat is abnormal.
	//
	// 	- uninstalled
	//
	// 	- failed: The installation failed.
	Machines []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// The username that is used to access the instance. The default value is elastic. Displayed when the **configType*	- is **collectorTargetInstance*	- or **collectorElasticsearchForKibana**.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The ID of the instance that is associated with the crawker. If the **configType*	- parameter is set to **collectorTargetInstance**, the value of this parameter is the ID of the output collector. If the **configType*	- parameter is set to **collectorDeployMachines*	- and the **type*	- parameter is set to **ACKCluster**, the value of this parameter is the ID of the ACK cluster.
	//
	// example:
	//
	// ECSInstanceId
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The type of the machine on which the Collector is deployed. This parameter is displayed when the **configType*	- is **collectorDeployMachine**. Valid values:
	//
	// 	- ECSInstanceId:ECS
	//
	// 	- ACKCluster: Container Kubernetes
	//
	// example:
	//
	// elastic
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetEnableMonitoring() *bool {
	return s.EnableMonitoring
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetGroupId() *string {
	return s.GroupId
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetHosts() []*string {
	return s.Hosts
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetMachines() []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines {
	return s.Machines
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetType() *string {
	return s.Type
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) GetUserName() *string {
	return s.UserName
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetConfigType(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetEnableMonitoring(v bool) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.EnableMonitoring = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetGroupId(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.GroupId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetHosts(v []*string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Hosts = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetInstanceId(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.InstanceId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetInstanceType(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.InstanceType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetMachines(v []*ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Machines = v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetProtocol(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Protocol = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetType(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.Type = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) SetUserName(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs {
	s.UserName = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigs) Validate() error {
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

type ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines struct {
	// The IDs of ECS instances.
	//
	// example:
	//
	// heartOk
	AgentStatus *string `json:"agentStatus,omitempty" xml:"agentStatus,omitempty"`
	// The list of access addresses of the specified instance for the output of the collector. Displayed when the **configType*	- is **collectorTargetInstance**.
	//
	// example:
	//
	// i-bp13y63575oypr9d****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) SetAgentStatus(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines {
	s.AgentStatus = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) SetInstanceId(v string) *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines {
	s.InstanceId = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultCollectorsExtendConfigsMachines) Validate() error {
	return dara.Validate(s)
}

type ListEcsInstancesResponseBodyResultIpAddress struct {
	// The information about the collectors on the ECS instance.
	//
	// example:
	//
	// 172.16.xx.xx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The IP address of the endpoint.
	//
	// example:
	//
	// private
	IpType *string `json:"ipType,omitempty" xml:"ipType,omitempty"`
}

func (s ListEcsInstancesResponseBodyResultIpAddress) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponseBodyResultIpAddress) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) GetHost() *string {
	return s.Host
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) GetIpType() *string {
	return s.IpType
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) SetHost(v string) *ListEcsInstancesResponseBodyResultIpAddress {
	s.Host = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) SetIpType(v string) *ListEcsInstancesResponseBodyResultIpAddress {
	s.IpType = &v
	return s
}

func (s *ListEcsInstancesResponseBodyResultIpAddress) Validate() error {
	return dara.Validate(s)
}
