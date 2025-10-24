// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetAcpConnectionTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *BatchGetAcpConnectionTicketRequest
	GetEndUserId() *string
	SetInstanceGroupId(v string) *BatchGetAcpConnectionTicketRequest
	GetInstanceGroupId() *string
	SetInstanceIds(v []*string) *BatchGetAcpConnectionTicketRequest
	GetInstanceIds() []*string
	SetInstanceTasks(v []*BatchGetAcpConnectionTicketRequestInstanceTasks) *BatchGetAcpConnectionTicketRequest
	GetInstanceTasks() []*BatchGetAcpConnectionTicketRequestInstanceTasks
}

type BatchGetAcpConnectionTicketRequest struct {
	// The ID of the user to whom the cloud phone instance is assigned.
	//
	// example:
	//
	// user
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-25nt4kk9whjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The IDs of the cloud phone instances. You can specify 1 to 100 IDs of cloud phone instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The instance connection tasks.
	InstanceTasks []*BatchGetAcpConnectionTicketRequestInstanceTasks `json:"InstanceTasks,omitempty" xml:"InstanceTasks,omitempty" type:"Repeated"`
}

func (s BatchGetAcpConnectionTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetAcpConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *BatchGetAcpConnectionTicketRequest) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *BatchGetAcpConnectionTicketRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *BatchGetAcpConnectionTicketRequest) GetInstanceTasks() []*BatchGetAcpConnectionTicketRequestInstanceTasks {
	return s.InstanceTasks
}

func (s *BatchGetAcpConnectionTicketRequest) SetEndUserId(v string) *BatchGetAcpConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) SetInstanceGroupId(v string) *BatchGetAcpConnectionTicketRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) SetInstanceIds(v []*string) *BatchGetAcpConnectionTicketRequest {
	s.InstanceIds = v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) SetInstanceTasks(v []*BatchGetAcpConnectionTicketRequestInstanceTasks) *BatchGetAcpConnectionTicketRequest {
	s.InstanceTasks = v
	return s
}

func (s *BatchGetAcpConnectionTicketRequest) Validate() error {
	if s.InstanceTasks != nil {
		for _, item := range s.InstanceTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetAcpConnectionTicketRequestInstanceTasks struct {
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-fkuit0cmyfvzz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// cn-hangzhou@c9f5c2e8-f5c4-4b01-8602-000cae94****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchGetAcpConnectionTicketRequestInstanceTasks) String() string {
	return dara.Prettify(s)
}

func (s BatchGetAcpConnectionTicketRequestInstanceTasks) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) SetInstanceId(v string) *BatchGetAcpConnectionTicketRequestInstanceTasks {
	s.InstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) SetTaskId(v string) *BatchGetAcpConnectionTicketRequestInstanceTasks {
	s.TaskId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketRequestInstanceTasks) Validate() error {
	return dara.Validate(s)
}
