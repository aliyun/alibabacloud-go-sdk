// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDownloadTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateDownloadTaskResponseBody
	GetStatus() *string
	SetTaskId(v int64) *CreateDownloadTaskResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *CreateDownloadTaskResponseBody
	GetTaskName() *string
}

type CreateDownloadTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E7F333E0-7B70-54DA-A307-4B2B49DEE923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the task. Valid values:
	//
	// finish: The task finished. You can query the task to obtain the download link of the file.
	//
	// start: The task start.
	//
	// error: An error occurred.
	//
	// expire: The task file is invalid and cannot be downloaded.
	//
	// example:
	//
	// start
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique ID of the task.
	//
	// example:
	//
	// 132
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the file download task.
	//
	// example:
	//
	// Internet Boundary Firewall Assets - IPv4
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateDownloadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDownloadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDownloadTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateDownloadTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateDownloadTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateDownloadTaskResponseBody) SetRequestId(v string) *CreateDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) SetStatus(v string) *CreateDownloadTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) SetTaskId(v int64) *CreateDownloadTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) SetTaskName(v string) *CreateDownloadTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
