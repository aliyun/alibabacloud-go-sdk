// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *GetGpuClusterTopologyRequest
	GetChannelId() *string
	SetClusterId(v string) *GetGpuClusterTopologyRequest
	GetClusterId() *string
	SetCommunicationId(v string) *GetGpuClusterTopologyRequest
	GetCommunicationId() *string
	SetInstanceIdList(v interface{}) *GetGpuClusterTopologyRequest
	GetInstanceIdList() interface{}
	SetPodList(v interface{}) *GetGpuClusterTopologyRequest
	GetPodList() interface{}
	SetStartTime(v string) *GetGpuClusterTopologyRequest
	GetStartTime() *string
	SetTaskIdList(v interface{}) *GetGpuClusterTopologyRequest
	GetTaskIdList() interface{}
	SetTopoType(v string) *GetGpuClusterTopologyRequest
	GetTopoType() *string
}

type GetGpuClusterTopologyRequest struct {
	// Used when viewing collective communication topology. Specifies the topology relationship for a given channel ID within a communication group, such as the inter-rank communication topology (tree or ring topology). Must be used together with communication_id and topo_type.
	//
	// ⚠️ Use the channel filter parameter of the GetGpuClusterArgsOptions operation to retrieve the list of channel_id values for a specified communication_id.
	//
	// example:
	//
	// communicationId =0x412523235123
	//
	// channelId=1
	//
	// topoType=ring
	ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty"`
	// The cluster name. Select a cluster from the cluster list returned by the GetListGpuCluster operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// "cluster1"
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// Used when viewing collective communication topology. Specifies the topology relationship for a given communication group ID, such as the inter-rank communication topology (tree or ring topology) within a communication group. Must be used together with channel_id and topo_type.
	//
	// ⚠️ Use the communication filter type of the GetGpuClusterArgsOptions operation to retrieve the list of communication_id values.
	//
	// example:
	//
	// communicationId =0x412523235123
	//
	// channelId=1
	//
	// topoType=ring
	CommunicationId *string `json:"communicationId,omitempty" xml:"communicationId,omitempty"`
	// The GPU topology information for the specified list of instances.
	//
	// example:
	//
	// ["instance1", "instance2","instance3"]
	InstanceIdList interface{} `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty"`
	// The GPU topology information for the specified list of pods.
	//
	// example:
	//
	// ["pod1", "pod2","pod3"]
	PodList interface{} `json:"podList,omitempty" xml:"podList,omitempty"`
	// The timestamp for the query. Defaults to the current time. Must be a second-level timestamp.
	//
	// example:
	//
	// 1709740800
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The GPU topology information for the specified list of nodes.
	//
	// example:
	//
	// ["task1", "task2"]
	TaskIdList interface{} `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
	// Used when viewing collective communication topology. Specifies the topology relationship for a given topology type within a communication group and channel, such as the inter-rank communication topology (tree or ring topology). Must be used together with communication_id and channel_id.
	//
	// ⚠️ Use the topoType filter parameter of the GetGpuClusterArgsOptions operation to retrieve the list of topo_type values for a specified communication_id and channel_id.
	//
	// example:
	//
	// communicationId =0x412523235123
	//
	// channelId=1
	//
	// topoType=ring
	TopoType *string `json:"topoType,omitempty" xml:"topoType,omitempty"`
}

func (s GetGpuClusterTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterTopologyRequest) GoString() string {
	return s.String()
}

func (s *GetGpuClusterTopologyRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetGpuClusterTopologyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetGpuClusterTopologyRequest) GetCommunicationId() *string {
	return s.CommunicationId
}

func (s *GetGpuClusterTopologyRequest) GetInstanceIdList() interface{} {
	return s.InstanceIdList
}

func (s *GetGpuClusterTopologyRequest) GetPodList() interface{} {
	return s.PodList
}

func (s *GetGpuClusterTopologyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGpuClusterTopologyRequest) GetTaskIdList() interface{} {
	return s.TaskIdList
}

func (s *GetGpuClusterTopologyRequest) GetTopoType() *string {
	return s.TopoType
}

func (s *GetGpuClusterTopologyRequest) SetChannelId(v string) *GetGpuClusterTopologyRequest {
	s.ChannelId = &v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetClusterId(v string) *GetGpuClusterTopologyRequest {
	s.ClusterId = &v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetCommunicationId(v string) *GetGpuClusterTopologyRequest {
	s.CommunicationId = &v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetInstanceIdList(v interface{}) *GetGpuClusterTopologyRequest {
	s.InstanceIdList = v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetPodList(v interface{}) *GetGpuClusterTopologyRequest {
	s.PodList = v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetStartTime(v string) *GetGpuClusterTopologyRequest {
	s.StartTime = &v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetTaskIdList(v interface{}) *GetGpuClusterTopologyRequest {
	s.TaskIdList = v
	return s
}

func (s *GetGpuClusterTopologyRequest) SetTopoType(v string) *GetGpuClusterTopologyRequest {
	s.TopoType = &v
	return s
}

func (s *GetGpuClusterTopologyRequest) Validate() error {
	return dara.Validate(s)
}
