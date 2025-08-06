// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTranscodeTaskResponseBody
	GetRequestId() *string
	SetTranscodeTaskList(v []*ListTranscodeTaskResponseBodyTranscodeTaskList) *ListTranscodeTaskResponseBody
	GetTranscodeTaskList() []*ListTranscodeTaskResponseBodyTranscodeTaskList
}

type ListTranscodeTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about transcoding tasks.
	TranscodeTaskList []*ListTranscodeTaskResponseBodyTranscodeTaskList `json:"TranscodeTaskList,omitempty" xml:"TranscodeTaskList,omitempty" type:"Repeated"`
}

func (s ListTranscodeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTranscodeTaskResponseBody) GetTranscodeTaskList() []*ListTranscodeTaskResponseBodyTranscodeTaskList {
	return s.TranscodeTaskList
}

func (s *ListTranscodeTaskResponseBody) SetRequestId(v string) *ListTranscodeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranscodeTaskResponseBody) SetTranscodeTaskList(v []*ListTranscodeTaskResponseBodyTranscodeTaskList) *ListTranscodeTaskResponseBody {
	s.TranscodeTaskList = v
	return s
}

func (s *ListTranscodeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeTaskResponseBodyTranscodeTaskList struct {
	// The time when the transcoding task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-23T12:40:12Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the transcoding task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-23T12:35:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The status of the transcoding task. Valid values:
	//
	// 	- **Processing**: In progress.
	//
	// 	- **Partial**: Some transcoding jobs were complete.
	//
	// 	- **CompleteAllSucc**: All transcoding jobs were successful.
	//
	// 	- **CompleteAllFail**: All transcoding jobs failed. If an exception occurs in the source file, no transcoding job is initiated and the transcoding task fails.
	//
	// 	- **CompletePartialSucc**: All transcoding jobs were complete but only some were successful.
	//
	// example:
	//
	// Processing
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The ID of the transcoding task.
	//
	// example:
	//
	// b1b65ab107*****ba3dbb900f6c1fe0
	TranscodeTaskId *string `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// b500c7094bd24*****f3e9900752d7c3
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The mode in which the transcoding task is triggered. Valid values:
	//
	// 	- **Auto**: The transcoding task is automatically triggered when the video is uploaded.
	//
	// 	- **Manual**: The transcoding task is triggered by calling the SubmitTranscodeJobs operation.
	//
	// example:
	//
	// Auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The ID of the audio or video file.
	//
	// example:
	//
	// d4860fcc6a5*****bce9fed52e893824
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListTranscodeTaskResponseBodyTranscodeTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTaskResponseBodyTranscodeTaskList) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetTranscodeTaskId() *string {
	return s.TranscodeTaskId
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetTrigger() *string {
	return s.Trigger
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) GetVideoId() *string {
	return s.VideoId
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetCompleteTime(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.CompleteTime = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetCreationTime(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.CreationTime = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTaskStatus(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.TaskStatus = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTranscodeTaskId(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.TranscodeTaskId = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTranscodeTemplateGroupId(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTrigger(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.Trigger = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetVideoId(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.VideoId = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) Validate() error {
	return dara.Validate(s)
}
