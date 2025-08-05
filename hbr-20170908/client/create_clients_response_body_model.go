// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateClientsResponseBody
	GetCode() *string
	SetInstanceStatuses(v *CreateClientsResponseBodyInstanceStatuses) *CreateClientsResponseBody
	GetInstanceStatuses() *CreateClientsResponseBodyInstanceStatuses
	SetMessage(v string) *CreateClientsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateClientsResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateClientsResponseBody
	GetTaskId() *string
}

type CreateClientsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the ECS instance. If you specify more than one instance IDs in the request and the status of an ECS instance does not meet the requirements to install an HBR client, an error message is returned based on the value of this parameter.
	InstanceStatuses *CreateClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Struct"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4A8A9AE4-F798-5E6D-853E-10F9F5A1BD4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of an asynchronous job.
	//
	// example:
	//
	// t-000h9x5t02vhyksf1x7k
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateClientsResponseBody) GetInstanceStatuses() *CreateClientsResponseBodyInstanceStatuses {
	return s.InstanceStatuses
}

func (s *CreateClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateClientsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateClientsResponseBody) SetCode(v string) *CreateClientsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateClientsResponseBody) SetInstanceStatuses(v *CreateClientsResponseBodyInstanceStatuses) *CreateClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *CreateClientsResponseBody) SetMessage(v string) *CreateClientsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClientsResponseBody) SetRequestId(v string) *CreateClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientsResponseBody) SetSuccess(v bool) *CreateClientsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateClientsResponseBody) SetTaskId(v string) *CreateClientsResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateClientsResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateClientsResponseBodyInstanceStatuses struct {
	InstanceStatus []*CreateClientsResponseBodyInstanceStatusesInstanceStatus `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty" type:"Repeated"`
}

func (s CreateClientsResponseBodyInstanceStatuses) String() string {
	return dara.Prettify(s)
}

func (s CreateClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *CreateClientsResponseBodyInstanceStatuses) GetInstanceStatus() []*CreateClientsResponseBodyInstanceStatusesInstanceStatus {
	return s.InstanceStatus
}

func (s *CreateClientsResponseBodyInstanceStatuses) SetInstanceStatus(v []*CreateClientsResponseBodyInstanceStatusesInstanceStatus) *CreateClientsResponseBodyInstanceStatuses {
	s.InstanceStatus = v
	return s
}

func (s *CreateClientsResponseBodyInstanceStatuses) Validate() error {
	return dara.Validate(s)
}

type CreateClientsResponseBodyInstanceStatusesInstanceStatus struct {
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-2zegp3cdu******uj9i
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether an HBR client can be installed on the ECS instance. Valid values:
	//
	// 	- true: An HBR client can be installed on the ECS instance.
	//
	// 	- false: An HBR client cannot be installed on the ECS instance.
	//
	// example:
	//
	// true
	ValidInstance *bool `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s CreateClientsResponseBodyInstanceStatusesInstanceStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateClientsResponseBodyInstanceStatusesInstanceStatus) GoString() string {
	return s.String()
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) GetValidInstance() *bool {
	return s.ValidInstance
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) SetInstanceId(v string) *CreateClientsResponseBodyInstanceStatusesInstanceStatus {
	s.InstanceId = &v
	return s
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) SetValidInstance(v bool) *CreateClientsResponseBodyInstanceStatusesInstanceStatus {
	s.ValidInstance = &v
	return s
}

func (s *CreateClientsResponseBodyInstanceStatusesInstanceStatus) Validate() error {
	return dara.Validate(s)
}
