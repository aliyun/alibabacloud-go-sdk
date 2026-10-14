// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *CreateAsyncTaskRequest
	GetResourceGroupId() *string
	SetTaskParams(v string) *CreateAsyncTaskRequest
	GetTaskParams() *string
	SetTaskType(v int32) *CreateAsyncTaskRequest
	GetTaskType() *int32
}

type CreateAsyncTaskRequest struct {
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. If this parameter is empty, the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The task parameter information. The value is a JSON-formatted string. The required task parameters vary based on the value of **TaskType**.
	//
	// If **TaskType*	- is set to **1**, **3**, **4**, **5**, or **6**, the task parameters use the following structure:
	//
	// - **instanceId**: String type. Required. The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// If **TaskType*	- is set to **2**, the task parameters use the following structure:
	//
	// - **domain**: String type. Optional. The domain name of the website. If this parameter is not specified, the forwarding rules of all websites are exported.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"instanceId": "ddoscoo-cn-mp91j1ao****"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// The type of the export task to create. Valid values:
	//
	// - **1**: Layer 4 export task. Exports port forwarding rules of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **2**: Layer 7 export task. Exports website forwarding rules.
	//
	// - **3**: Session persistence and health check export task. Exports session persistence and health check configurations of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **4**: DDoS mitigation policy export task. Exports DDoS mitigation policy configurations of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **5**: Blacklist (destination IP address-based) download task. Exports blacklisted IP addresses for the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **6**: Whitelist (destination IP address-based) download task. Exports whitelisted IP addresses for the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAsyncTaskRequest) GetTaskParams() *string {
	return s.TaskParams
}

func (s *CreateAsyncTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateAsyncTaskRequest) SetResourceGroupId(v string) *CreateAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskParams(v string) *CreateAsyncTaskRequest {
	s.TaskParams = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskType(v int32) *CreateAsyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
