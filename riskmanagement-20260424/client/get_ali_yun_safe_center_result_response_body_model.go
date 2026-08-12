// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliYunSafeCenterResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAliYunSafeCenterResultResponseBody
	GetCode() *string
	SetData(v *GetAliYunSafeCenterResultResponseBodyData) *GetAliYunSafeCenterResultResponseBody
	GetData() *GetAliYunSafeCenterResultResponseBodyData
	SetMessage(v string) *GetAliYunSafeCenterResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAliYunSafeCenterResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAliYunSafeCenterResultResponseBody
	GetSuccess() *bool
}

type GetAliYunSafeCenterResultResponseBody struct {
	// The status code.
	//
	// > 200: Success. Other values (500, 400, etc.): Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAliYunSafeCenterResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3C107939-59BD-5EB9-B250-39559C830A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current API call is successful. Valid values: True/False. This does not indicate the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAliYunSafeCenterResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAliYunSafeCenterResultResponseBody) GetData() *GetAliYunSafeCenterResultResponseBodyData {
	return s.Data
}

func (s *GetAliYunSafeCenterResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAliYunSafeCenterResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAliYunSafeCenterResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAliYunSafeCenterResultResponseBody) SetCode(v string) *GetAliYunSafeCenterResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBody) SetData(v *GetAliYunSafeCenterResultResponseBodyData) *GetAliYunSafeCenterResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBody) SetMessage(v string) *GetAliYunSafeCenterResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBody) SetRequestId(v string) *GetAliYunSafeCenterResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBody) SetSuccess(v bool) *GetAliYunSafeCenterResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAliYunSafeCenterResultResponseBodyData struct {
	// The status of the Security Center Agent.
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	AgentStatus *bool `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The running status of the ECS instance.
	//
	// - **true**: Running.
	//
	// - **false**: Not running.
	//
	// example:
	//
	// true
	EcsInstanceStatus *bool `json:"EcsInstanceStatus,omitempty" xml:"EcsInstanceStatus,omitempty"`
	// The list of instance IDs of identical security alert events in Security Center.
	InstanceIds []*int64 `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8169D779-3391-541F-936B-11F4EC09AD0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The running status of the simple application server.
	//
	// - **true**: Running.
	//
	// - **false**: Not running.
	//
	// example:
	//
	// true
	SwasInstanceStatus *bool `json:"SwasInstanceStatus,omitempty" xml:"SwasInstanceStatus,omitempty"`
	// The task ID for querying security alerting events triggered by the same rule or alerting type.
	//
	// example:
	//
	// 12313123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Indicates whether the list of security alerting events that match the same IP rule or same alerting type as the alerting event to be handled is empty.
	//
	// - **true**: Empty.
	//
	// - **false**: Not empty.
	//
	// example:
	//
	// true
	TaskStatus *bool `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetAliYunSafeCenterResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAliYunSafeCenterResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetAgentStatus() *bool {
	return s.AgentStatus
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetEcsInstanceStatus() *bool {
	return s.EcsInstanceStatus
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetInstanceIds() []*int64 {
	return s.InstanceIds
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetSwasInstanceStatus() *bool {
	return s.SwasInstanceStatus
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAliYunSafeCenterResultResponseBodyData) GetTaskStatus() *bool {
	return s.TaskStatus
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetAgentStatus(v bool) *GetAliYunSafeCenterResultResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetEcsInstanceStatus(v bool) *GetAliYunSafeCenterResultResponseBodyData {
	s.EcsInstanceStatus = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetInstanceIds(v []*int64) *GetAliYunSafeCenterResultResponseBodyData {
	s.InstanceIds = v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetRequestId(v string) *GetAliYunSafeCenterResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetSwasInstanceStatus(v bool) *GetAliYunSafeCenterResultResponseBodyData {
	s.SwasInstanceStatus = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetTaskId(v int64) *GetAliYunSafeCenterResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) SetTaskStatus(v bool) *GetAliYunSafeCenterResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetAliYunSafeCenterResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
