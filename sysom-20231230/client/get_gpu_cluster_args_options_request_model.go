// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterArgsOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterType(v string) *GetGpuClusterArgsOptionsRequest
	GetFilterType() *string
	SetFilterValues(v interface{}) *GetGpuClusterArgsOptionsRequest
	GetFilterValues() interface{}
}

type GetGpuClusterArgsOptionsRequest struct {
	// 1. Query the machine list of a cluster
	//
	// Set the retrieval type to comm_instance_id_list or instance_id_list.
	//
	// comm_instance_id_list: retrieves only the machines that participate in collective communication in the specified cluster.
	//
	// instance_id_list: retrieves all machines in the specified cluster.
	//
	// 2. Query the AI task list of a cluster
	//
	// Set the retrieval type to comm_task_id_list or task_id_list.
	//
	// comm_task_id_list: retrieves only the tasks that involve collective communication in the specified cluster.
	//
	// task_id_list: retrieves all tasks in the specified cluster.
	//
	// 3. Query the AI job pod list of a cluster
	//
	// Set the retrieval type to comm_pod_list or pod_list.
	//
	// comm_pod_list: retrieves only the pods that participate in collective communication in the specified cluster.
	//
	// pod_list: retrieves all pods in the specified cluster.
	//
	// 4. Query the collective communication topology
	//
	// communication: retrieves all communication group IDs.
	//
	// channel: retrieves the channel list under the specified communication group ID.
	//
	// topoType: retrieves the topology type under the specified channel of the specified communication group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance_id_list
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// The filter conditions. The value is a JSON object. The cluster_id field is required. The following optional filter conditions are supported:
	//
	// - instance_id_list
	//
	// - comm_instance_id_list
	//
	// - pod_list
	//
	// - comm_pod_list
	//
	// - communication
	//
	// - task_id_list
	//
	// - comm_task_id_list
	//
	// - channel
	//
	// - topoType
	//
	// example:
	//
	// {
	//
	// "cluster_id": "cluster1"
	//
	// }
	FilterValues interface{} `json:"filterValues,omitempty" xml:"filterValues,omitempty"`
}

func (s GetGpuClusterArgsOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterArgsOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetGpuClusterArgsOptionsRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *GetGpuClusterArgsOptionsRequest) GetFilterValues() interface{} {
	return s.FilterValues
}

func (s *GetGpuClusterArgsOptionsRequest) SetFilterType(v string) *GetGpuClusterArgsOptionsRequest {
	s.FilterType = &v
	return s
}

func (s *GetGpuClusterArgsOptionsRequest) SetFilterValues(v interface{}) *GetGpuClusterArgsOptionsRequest {
	s.FilterValues = v
	return s
}

func (s *GetGpuClusterArgsOptionsRequest) Validate() error {
	return dara.Validate(s)
}
