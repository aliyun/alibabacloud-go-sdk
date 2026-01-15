// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetAcpConnectionTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceConnectionModels(v []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) *BatchGetAcpConnectionTicketResponseBody
	GetInstanceConnectionModels() []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels
	SetRequestId(v string) *BatchGetAcpConnectionTicketResponseBody
	GetRequestId() *string
}

type BatchGetAcpConnectionTicketResponseBody struct {
	// The results of the instance connection tasks.
	InstanceConnectionModels []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels `json:"InstanceConnectionModels,omitempty" xml:"InstanceConnectionModels,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetAcpConnectionTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetAcpConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketResponseBody) GetInstanceConnectionModels() []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	return s.InstanceConnectionModels
}

func (s *BatchGetAcpConnectionTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetAcpConnectionTicketResponseBody) SetInstanceConnectionModels(v []*BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) *BatchGetAcpConnectionTicketResponseBody {
	s.InstanceConnectionModels = v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBody) SetRequestId(v string) *BatchGetAcpConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBody) Validate() error {
	if s.InstanceConnectionModels != nil {
		for _, item := range s.InstanceConnectionModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-1uzb6heg797z3****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ErrorCode          *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-ajxvwo1u0hqvd****
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PersistentAppInstanceId *string `json:"PersistentAppInstanceId,omitempty" xml:"PersistentAppInstanceId,omitempty"`
	Port                    *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// cn-hangzhou@c9f5c2e8-f5c4-4b01-8602-000cae94****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The state of the task.
	//
	// example:
	//
	// FINISHED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The ticket used to connect to the cloud phone instance.
	//
	// example:
	//
	// piVE58_AdmVSVW7SEW3*AE5*p8mmO5gvItsNOmv4S_f_cNpoU_BOTwChTBoNM1ZJeedfK9zxYnbN5hossqIZCr6t7SGxRigm2Cb4fGaCdBZWIzmgdHq6sXXZQg4KFWufyvpeV*0*Cm58slMT1tJw3****
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) String() string {
	return dara.Prettify(s)
}

func (s BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetPersistentAppInstanceId() *string {
	return s.PersistentAppInstanceId
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetPort() *int32 {
	return s.Port
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) GetTicket() *string {
	return s.Ticket
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetAppInstanceGroupId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetAppInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.AppInstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetErrorCode(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.ErrorCode = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.InstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetPersistentAppInstanceId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.PersistentAppInstanceId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetPort(v int32) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.Port = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetTaskId(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.TaskId = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetTaskStatus(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.TaskStatus = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) SetTicket(v string) *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels {
	s.Ticket = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponseBodyInstanceConnectionModels) Validate() error {
	return dara.Validate(s)
}
