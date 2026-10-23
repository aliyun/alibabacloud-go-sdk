// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetAcpConnectionTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionMode(v string) *BatchGetAcpConnectionTicketRequest
	GetConnectionMode() *string
	SetEndUserId(v string) *BatchGetAcpConnectionTicketRequest
	GetEndUserId() *string
	SetInstanceGroupId(v string) *BatchGetAcpConnectionTicketRequest
	GetInstanceGroupId() *string
	SetInstanceIds(v []*string) *BatchGetAcpConnectionTicketRequest
	GetInstanceIds() []*string
	SetInstanceTasks(v []*BatchGetAcpConnectionTicketRequestInstanceTasks) *BatchGetAcpConnectionTicketRequest
	GetInstanceTasks() []*BatchGetAcpConnectionTicketRequestInstanceTasks
	SetPorts(v []*string) *BatchGetAcpConnectionTicketRequest
	GetPorts() []*string
}

type BatchGetAcpConnectionTicketRequest struct {
	// The connection mode.
	//
	// example:
	//
	// custom
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The ID of the assigned user. If you do not specify this parameter, the system user assigned in the background is used to obtain the ticket.
	//
	// <props="china">
	//
	// The usage of this parameter varies based on the Cloud Phone edition and the instance streaming mode:
	//
	// - Cloud Phone Instance Edition or Cloud Phone Matrix Edition (preemptive mode): You can obtain a `Ticket` only by using the same `EnduserId`.
	//
	// - Cloud Phone Matrix Edition (collaborative mode): You can pass different `EnduserId` values to obtain `Ticket` values and stream for different users (up to 5) at the same time. You can pass only one `EnduserId` at a time.
	//
	// > You can define the instance streaming mode by using the `StreamMode` parameter of the `ModifyCloudPhoneNode` operation.
	//
	// example:
	//
	// user
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The instance group ID.
	//
	// example:
	//
	// ag-25nt4kk9whjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The list of Cloud Phone instance IDs. You can specify 1 to 100 IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The instance connection tasks. This parameter is used to query whether the ticket retrieval task is completed.
	InstanceTasks []*BatchGetAcpConnectionTicketRequestInstanceTasks `json:"InstanceTasks,omitempty" xml:"InstanceTasks,omitempty" type:"Repeated"`
	// The list of custom connection ports.
	Ports []*string `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
}

func (s BatchGetAcpConnectionTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetAcpConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketRequest) GetConnectionMode() *string {
	return s.ConnectionMode
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

func (s *BatchGetAcpConnectionTicketRequest) GetPorts() []*string {
	return s.Ports
}

func (s *BatchGetAcpConnectionTicketRequest) SetConnectionMode(v string) *BatchGetAcpConnectionTicketRequest {
	s.ConnectionMode = &v
	return s
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

func (s *BatchGetAcpConnectionTicketRequest) SetPorts(v []*string) *BatchGetAcpConnectionTicketRequest {
	s.Ports = v
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
	// The instance ID.
	//
	// example:
	//
	// acp-fkuit0cmyfvzz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task ID.
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
