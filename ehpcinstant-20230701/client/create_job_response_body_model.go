// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateJobResponseBody
	GetRequestId() *string
	SetTasks(v []*CreateJobResponseBodyTasks) *CreateJobResponseBody
	GetTasks() []*CreateJobResponseBodyTasks
}

type CreateJobResponseBody struct {
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*CreateJobResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CreateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobResponseBody) GetTasks() []*CreateJobResponseBodyTasks {
	return s.Tasks
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetTasks(v []*CreateJobResponseBodyTasks) *CreateJobResponseBody {
	s.Tasks = v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateJobResponseBodyTasks struct {
	ExecutorIds []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	TaskName    *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateJobResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBodyTasks) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *CreateJobResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateJobResponseBodyTasks) SetExecutorIds(v []*string) *CreateJobResponseBodyTasks {
	s.ExecutorIds = v
	return s
}

func (s *CreateJobResponseBodyTasks) SetTaskName(v string) *CreateJobResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *CreateJobResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
