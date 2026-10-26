// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMailTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMailTaskStatusResponseBody
	GetRequestId() *string
	SetResult(v []*GetMailTaskStatusResponseBodyResult) *GetMailTaskStatusResponseBody
	GetResult() []*GetMailTaskStatusResponseBodyResult
	SetSuccess(v bool) *GetMailTaskStatusResponseBody
	GetSuccess() *bool
}

type GetMailTaskStatusResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 38C0FEC8-****-415C-A9F1-****422BDB65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	Result []*GetMailTaskStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMailTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMailTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMailTaskStatusResponseBody) GetResult() []*GetMailTaskStatusResponseBodyResult {
	return s.Result
}

func (s *GetMailTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMailTaskStatusResponseBody) SetRequestId(v string) *GetMailTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMailTaskStatusResponseBody) SetResult(v []*GetMailTaskStatusResponseBodyResult) *GetMailTaskStatusResponseBody {
	s.Result = v
	return s
}

func (s *GetMailTaskStatusResponseBody) SetSuccess(v bool) *GetMailTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetMailTaskStatusResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMailTaskStatusResponseBodyResult struct {
	// Execution time, in the format yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2024-10-09 17:34:11
	ExecTime *string `json:"execTime,omitempty" xml:"execTime,omitempty"`
	// Mail ID
	//
	// example:
	//
	// c38f73f4c5*****c808c41b3f4d23b7852
	MailId *string `json:"mailId,omitempty" xml:"mailId,omitempty"`
	// Mail status. Possible values:
	//
	// - Success: SENT
	//
	// - Failure: FAILED
	//
	// - In Progress: PROCESSING
	//
	// example:
	//
	// SENT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Task ID
	//
	// example:
	//
	// 1282xxx610816
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetMailTaskStatusResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMailTaskStatusResponseBodyResult) GetExecTime() *string {
	return s.ExecTime
}

func (s *GetMailTaskStatusResponseBodyResult) GetMailId() *string {
	return s.MailId
}

func (s *GetMailTaskStatusResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetMailTaskStatusResponseBodyResult) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetMailTaskStatusResponseBodyResult) SetExecTime(v string) *GetMailTaskStatusResponseBodyResult {
	s.ExecTime = &v
	return s
}

func (s *GetMailTaskStatusResponseBodyResult) SetMailId(v string) *GetMailTaskStatusResponseBodyResult {
	s.MailId = &v
	return s
}

func (s *GetMailTaskStatusResponseBodyResult) SetStatus(v string) *GetMailTaskStatusResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetMailTaskStatusResponseBodyResult) SetTaskId(v int64) *GetMailTaskStatusResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetMailTaskStatusResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
