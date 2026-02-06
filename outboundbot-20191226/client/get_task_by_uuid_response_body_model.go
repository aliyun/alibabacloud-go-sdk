// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskByUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskByUuidResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetTaskByUuidResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTaskByUuidResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskByUuidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskByUuidResponseBody
	GetSuccess() *bool
	SetTask(v *GetTaskByUuidResponseBodyTask) *GetTaskByUuidResponseBody
	GetTask() *GetTaskByUuidResponseBodyTask
}

type GetTaskByUuidResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// {}
	Task *GetTaskByUuidResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetTaskByUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskByUuidResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskByUuidResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTaskByUuidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskByUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskByUuidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskByUuidResponseBody) GetTask() *GetTaskByUuidResponseBodyTask {
	return s.Task
}

func (s *GetTaskByUuidResponseBody) SetCode(v string) *GetTaskByUuidResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskByUuidResponseBody) SetHttpStatusCode(v int32) *GetTaskByUuidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTaskByUuidResponseBody) SetMessage(v string) *GetTaskByUuidResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskByUuidResponseBody) SetRequestId(v string) *GetTaskByUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskByUuidResponseBody) SetSuccess(v bool) *GetTaskByUuidResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskByUuidResponseBody) SetTask(v *GetTaskByUuidResponseBodyTask) *GetTaskByUuidResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskByUuidResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskByUuidResponseBodyTask struct {
	// example:
	//
	// 1640090211434
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 13777777777
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 057190294
	CallingNumber *string                                       `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	Conversations []*GetTaskByUuidResponseBodyTaskConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
	// example:
	//
	// 1640090211434
	EndReason *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// example:
	//
	// 1640090211434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1640090211434
	PlannedTime *int64 `json:"PlannedTime,omitempty" xml:"PlannedTime,omitempty"`
}

func (s GetTaskByUuidResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUuidResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskByUuidResponseBodyTask) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *GetTaskByUuidResponseBodyTask) GetCallId() *string {
	return s.CallId
}

func (s *GetTaskByUuidResponseBodyTask) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *GetTaskByUuidResponseBodyTask) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *GetTaskByUuidResponseBodyTask) GetConversations() []*GetTaskByUuidResponseBodyTaskConversations {
	return s.Conversations
}

func (s *GetTaskByUuidResponseBodyTask) GetEndReason() *string {
	return s.EndReason
}

func (s *GetTaskByUuidResponseBodyTask) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTaskByUuidResponseBodyTask) GetId() *string {
	return s.Id
}

func (s *GetTaskByUuidResponseBodyTask) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTaskByUuidResponseBodyTask) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *GetTaskByUuidResponseBodyTask) GetJobId() *string {
	return s.JobId
}

func (s *GetTaskByUuidResponseBodyTask) GetPlannedTime() *int64 {
	return s.PlannedTime
}

func (s *GetTaskByUuidResponseBodyTask) SetActualTime(v int64) *GetTaskByUuidResponseBodyTask {
	s.ActualTime = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetCallId(v string) *GetTaskByUuidResponseBodyTask {
	s.CallId = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetCalledNumber(v string) *GetTaskByUuidResponseBodyTask {
	s.CalledNumber = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetCallingNumber(v string) *GetTaskByUuidResponseBodyTask {
	s.CallingNumber = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetConversations(v []*GetTaskByUuidResponseBodyTaskConversations) *GetTaskByUuidResponseBodyTask {
	s.Conversations = v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetEndReason(v string) *GetTaskByUuidResponseBodyTask {
	s.EndReason = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetEndTime(v int64) *GetTaskByUuidResponseBodyTask {
	s.EndTime = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetId(v string) *GetTaskByUuidResponseBodyTask {
	s.Id = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetInstanceId(v string) *GetTaskByUuidResponseBodyTask {
	s.InstanceId = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetJobGroupId(v string) *GetTaskByUuidResponseBodyTask {
	s.JobGroupId = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetJobId(v string) *GetTaskByUuidResponseBodyTask {
	s.JobId = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) SetPlannedTime(v int64) *GetTaskByUuidResponseBodyTask {
	s.PlannedTime = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTask) Validate() error {
	if s.Conversations != nil {
		for _, item := range s.Conversations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskByUuidResponseBodyTaskConversations struct {
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Script     *string `json:"Script,omitempty" xml:"Script,omitempty"`
	SequenceId *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	Speaker    *string `json:"Speaker,omitempty" xml:"Speaker,omitempty"`
	Timestamp  *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTaskByUuidResponseBodyTaskConversations) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUuidResponseBodyTaskConversations) GoString() string {
	return s.String()
}

func (s *GetTaskByUuidResponseBodyTaskConversations) GetAction() *string {
	return s.Action
}

func (s *GetTaskByUuidResponseBodyTaskConversations) GetScript() *string {
	return s.Script
}

func (s *GetTaskByUuidResponseBodyTaskConversations) GetSequenceId() *string {
	return s.SequenceId
}

func (s *GetTaskByUuidResponseBodyTaskConversations) GetSpeaker() *string {
	return s.Speaker
}

func (s *GetTaskByUuidResponseBodyTaskConversations) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetTaskByUuidResponseBodyTaskConversations) SetAction(v string) *GetTaskByUuidResponseBodyTaskConversations {
	s.Action = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTaskConversations) SetScript(v string) *GetTaskByUuidResponseBodyTaskConversations {
	s.Script = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTaskConversations) SetSequenceId(v string) *GetTaskByUuidResponseBodyTaskConversations {
	s.SequenceId = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTaskConversations) SetSpeaker(v string) *GetTaskByUuidResponseBodyTaskConversations {
	s.Speaker = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTaskConversations) SetTimestamp(v int64) *GetTaskByUuidResponseBodyTaskConversations {
	s.Timestamp = &v
	return s
}

func (s *GetTaskByUuidResponseBodyTaskConversations) Validate() error {
	return dara.Validate(s)
}
