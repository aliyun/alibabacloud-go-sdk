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
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The details of the asynchronous export task. The value is a JSON string. The field in the value varies with **TaskType**.
	//
	// If **TaskType*	- is set to **1**, **3**, **4**, **5**, or **6**, the following filed is returned:
	//
	// 	- **instanceId**: the ID of the instance. This field is required and must be of the STRING type.
	//
	// If **TaskType*	- is set to **2**, the following field is returned:
	//
	// 	- **domain**: the domain name of the website, which must be of the STRING type. If you do not specify this field, the forwarding rules of all websites are exported.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"instanceId": "ddoscoo-cn-mp91j1ao****"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// The type of the asynchronous export task that you want to create. Valid values:
	//
	// 	- **1**: the task to export the port forwarding rules of an instance
	//
	// 	- **2**: the task to export the forwarding rules of a website protected by an instance
	//
	// 	- **3**: the task to export the session persistence and health check settings of an instance
	//
	// 	- **4**: the task to export the anti-DDoS mitigation policies of an instance
	//
	// 	- **5**: the task to download the blacklist for destination IP addresses of an instance
	//
	// 	- **6**: the task to download the whitelist for destination IP addresses of an instance
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
