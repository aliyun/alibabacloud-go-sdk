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
	// The ID of the channel. Used to view the collective communication topology and the topology relationships of a specified channel ID under a specified communication group, such as the communication, tree, or ring topology among ranks within a communication group. This parameter must be used together with communication_id and topo_type. You can call the GetGpuClusterArgsOptions operation with the channel filter parameter to obtain the list of channel_id values under a specified communication_id.
	//
	// example:
	//
	// communicationId =0x412523235123
	//
	// channelId=1
	//
	// topoType=ring
	ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty"`
	// The name of the cluster. You can select a cluster from the cluster list obtained by calling the GetListGpuCluster operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// "cluster1"
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The ID of the communication group. Used to view the collective communication topology and the topology relationships of a specified communication group ID, such as the communication, tree, or ring topology among ranks within a communication group. This parameter must be used together with channel_id and topo_type. You can call the GetGpuClusterArgsOptions operation with the communication filter type to obtain the list of communication_id values.
	//
	// example:
	//
	// communicationId =0x412523235123
	//
	// channelId=1
	//
	// topoType=ring
	CommunicationId *string `json:"communicationId,omitempty" xml:"communicationId,omitempty"`
	// The list of instances for which to query GPU topology information.
	//
	// example:
	//
	// ["instance1", "instance2","instance3"]
	InstanceIdList interface{} `json:"instanceIdList,omitempty" xml:"instanceIdList,omitempty"`
	// The list of pods for which to query GPU topology information.
	//
	// example:
	//
	// ["pod1", "pod2","pod3"]
	PodList interface{} `json:"podList,omitempty" xml:"podList,omitempty"`
	// The query timestamp. Default value: the current time. The timestamp must be accurate to seconds.
	//
	// example:
	//
	// 1709740800
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The list of tasks for which to query GPU topology information.
	//
	// example:
	//
	// ["task1", "task2"]
	TaskIdList interface{} `json:"taskIdList,omitempty" xml:"taskIdList,omitempty"`
	// The topology type. Used to view the collective communication topology and the topology relationships of a specified topology type under a specified communication group and channel ID, such as the communication, tree, or ring topology among ranks within a communication group. This parameter must be used together with communication_id and channel_id. You can call the GetGpuClusterArgsOptions operation with the topoType filter parameter to obtain the list of topo_type values under a specified communication_id and channel_id.
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
