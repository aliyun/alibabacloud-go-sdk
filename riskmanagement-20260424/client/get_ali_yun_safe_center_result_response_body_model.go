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
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAliYunSafeCenterResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3C107939-59BD-5EB9-B250-39559C830A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// true
	AgentStatus *bool `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// example:
	//
	// true
	EcsInstanceStatus *bool    `json:"EcsInstanceStatus,omitempty" xml:"EcsInstanceStatus,omitempty"`
	InstanceIds       []*int64 `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 8169D779-3391-541F-936B-11F4EC09AD0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SwasInstanceStatus *bool `json:"SwasInstanceStatus,omitempty" xml:"SwasInstanceStatus,omitempty"`
	// example:
	//
	// 12313123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
