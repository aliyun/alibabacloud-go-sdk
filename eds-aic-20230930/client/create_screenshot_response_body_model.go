// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScreenshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateScreenshotResponseBody
	GetRequestId() *string
	SetTasks(v []*CreateScreenshotResponseBodyTasks) *CreateScreenshotResponseBody
	GetTasks() []*CreateScreenshotResponseBodyTasks
}

type CreateScreenshotResponseBody struct {
	// The ID of the request. If the request fails, share this ID with technical support to help diagnose the issue.
	//
	// example:
	//
	// 3AF82CE1-2801-52CE-BF64-B491DD7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*CreateScreenshotResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CreateScreenshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScreenshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScreenshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScreenshotResponseBody) GetTasks() []*CreateScreenshotResponseBodyTasks {
	return s.Tasks
}

func (s *CreateScreenshotResponseBody) SetRequestId(v string) *CreateScreenshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScreenshotResponseBody) SetTasks(v []*CreateScreenshotResponseBodyTasks) *CreateScreenshotResponseBody {
	s.Tasks = v
	return s
}

func (s *CreateScreenshotResponseBody) Validate() error {
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

type CreateScreenshotResponseBodyTasks struct {
	// The ID of the cloud phone instance.
	//
	// example:
	//
	// acp-bwhtebzah2fse****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The ID of the task. You can use the task ID with the DescribeTasks operation to get the download link for the screenshot.
	//
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateScreenshotResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s CreateScreenshotResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *CreateScreenshotResponseBodyTasks) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *CreateScreenshotResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateScreenshotResponseBodyTasks) SetAndroidInstanceId(v string) *CreateScreenshotResponseBodyTasks {
	s.AndroidInstanceId = &v
	return s
}

func (s *CreateScreenshotResponseBodyTasks) SetTaskId(v string) *CreateScreenshotResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *CreateScreenshotResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
