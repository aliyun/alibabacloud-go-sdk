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
	// 1. Query cluster machine list scenario: Set the retrieval type to comm_instance_id_list or instance_id_list. comm_instance_id_list: Retrieves only the list of machines in the specified cluster that participate in collective communication. instance_id_list: Retrieves the list of all machines in the specified cluster.
	//
	// 2. Query cluster AI task list scenario: Set the retrieval type to comm_task_id_list or task_id_list. comm_task_id_list: Retrieves only the list of tasks in the specified cluster that involve collective communication. task_id_list: Retrieves the list of all tasks in the specified cluster.
	//
	// 3. Query cluster AI job pod list scenario: Set the retrieval type to comm_pod_list or pod_list. comm_pod_list: Retrieves only the list of pods in the specified cluster that participate in collective communication. pod_list: Retrieves the list of all pods in the specified cluster.
	//
	// 4. Query collective communication topology scenario: communication: Retrieves all communication group IDs. channel: Retrieves the channel list under the specified communication group ID. topoType: Retrieves the topology type under the specified channel of the specified communication group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance_id_list
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// The filter conditions, which is a JSON object. The cluster_id field is required. Other filter conditions support the following values:
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
