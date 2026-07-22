// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyScheduledTaskResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyScheduledTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*ModifyScheduledTaskResponseBodyTasks) *ModifyScheduledTaskResponseBody
	GetTasks() []*ModifyScheduledTaskResponseBodyTasks
	SetTotalCount(v int32) *ModifyScheduledTaskResponseBody
	GetTotalCount() *int32
}

type ModifyScheduledTaskResponseBody struct {
	// The API status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A51B1DF-96FF-3BCC-B08C-783161D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of modification results.
	Tasks []*ModifyScheduledTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScheduledTaskResponseBody) GetTasks() []*ModifyScheduledTaskResponseBodyTasks {
	return s.Tasks
}

func (s *ModifyScheduledTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModifyScheduledTaskResponseBody) SetCode(v string) *ModifyScheduledTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetMessage(v string) *ModifyScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetTasks(v []*ModifyScheduledTaskResponseBodyTasks) *ModifyScheduledTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetTotalCount(v int32) *ModifyScheduledTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyScheduledTaskResponseBodyTasks struct {
	// The list of batch delivery results for instance scheduled tasks.
	InstanceResults []*ModifyScheduledTaskResponseBodyTasksInstanceResults `json:"InstanceResults,omitempty" xml:"InstanceResults,omitempty" type:"Repeated"`
	// The updated CAS version number.
	//
	// example:
	//
	// 2
	NewVersion *int32 `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	// The scheduled task ID.
	//
	// example:
	//
	// sch-260705-agbx*****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
}

func (s ModifyScheduledTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBodyTasks) GetInstanceResults() []*ModifyScheduledTaskResponseBodyTasksInstanceResults {
	return s.InstanceResults
}

func (s *ModifyScheduledTaskResponseBodyTasks) GetNewVersion() *int32 {
	return s.NewVersion
}

func (s *ModifyScheduledTaskResponseBodyTasks) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ModifyScheduledTaskResponseBodyTasks) SetInstanceResults(v []*ModifyScheduledTaskResponseBodyTasksInstanceResults) *ModifyScheduledTaskResponseBodyTasks {
	s.InstanceResults = v
	return s
}

func (s *ModifyScheduledTaskResponseBodyTasks) SetNewVersion(v int32) *ModifyScheduledTaskResponseBodyTasks {
	s.NewVersion = &v
	return s
}

func (s *ModifyScheduledTaskResponseBodyTasks) SetScheduledId(v string) *ModifyScheduledTaskResponseBodyTasks {
	s.ScheduledId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBodyTasks) Validate() error {
	if s.InstanceResults != nil {
		for _, item := range s.InstanceResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyScheduledTaskResponseBodyTasksInstanceResults struct {
	// The error message.
	//
	// example:
	//
	// Instance not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// acp-4dkmkip0l0uw*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScheduledTaskResponseBodyTasksInstanceResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskResponseBodyTasksInstanceResults) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) SetErrorMessage(v string) *ModifyScheduledTaskResponseBodyTasksInstanceResults {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) SetInstanceId(v string) *ModifyScheduledTaskResponseBodyTasksInstanceResults {
	s.InstanceId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) SetSuccess(v bool) *ModifyScheduledTaskResponseBodyTasksInstanceResults {
	s.Success = &v
	return s
}

func (s *ModifyScheduledTaskResponseBodyTasksInstanceResults) Validate() error {
	return dara.Validate(s)
}
