// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScheduledTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteScheduledTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*DeleteScheduledTaskResponseBodyTasks) *DeleteScheduledTaskResponseBody
	GetTasks() []*DeleteScheduledTaskResponseBodyTasks
	SetTotalCount(v int32) *DeleteScheduledTaskResponseBody
	GetTotalCount() *int32
}

type DeleteScheduledTaskResponseBody struct {
	// The status code. A value of 200 indicates success.
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
	// The list of deletion results.
	Tasks []*DeleteScheduledTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduledTaskResponseBody) GetTasks() []*DeleteScheduledTaskResponseBodyTasks {
	return s.Tasks
}

func (s *DeleteScheduledTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DeleteScheduledTaskResponseBody) SetCode(v string) *DeleteScheduledTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) SetMessage(v string) *DeleteScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) SetRequestId(v string) *DeleteScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) SetTasks(v []*DeleteScheduledTaskResponseBodyTasks) *DeleteScheduledTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *DeleteScheduledTaskResponseBody) SetTotalCount(v int32) *DeleteScheduledTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DeleteScheduledTaskResponseBody) Validate() error {
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

type DeleteScheduledTaskResponseBodyTasks struct {
	// The ID of the scheduled task.
	//
	// example:
	//
	// sch-260705-agbx1eev
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// Indicates whether the scheduled task is deleted successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScheduledTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponseBodyTasks) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *DeleteScheduledTaskResponseBodyTasks) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScheduledTaskResponseBodyTasks) SetScheduledId(v string) *DeleteScheduledTaskResponseBodyTasks {
	s.ScheduledId = &v
	return s
}

func (s *DeleteScheduledTaskResponseBodyTasks) SetSuccess(v bool) *DeleteScheduledTaskResponseBodyTasks {
	s.Success = &v
	return s
}

func (s *DeleteScheduledTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
