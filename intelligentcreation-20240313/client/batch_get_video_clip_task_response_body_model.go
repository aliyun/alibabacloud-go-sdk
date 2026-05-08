// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoClipTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetVideoClipTaskResponseBody
	GetRequestId() *string
	SetTaskList(v []*BatchGetVideoClipTaskResponseBodyTaskList) *BatchGetVideoClipTaskResponseBody
	GetTaskList() []*BatchGetVideoClipTaskResponseBodyTaskList
}

type BatchGetVideoClipTaskResponseBody struct {
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskList  []*BatchGetVideoClipTaskResponseBodyTaskList `json:"taskList,omitempty" xml:"taskList,omitempty" type:"Repeated"`
}

func (s BatchGetVideoClipTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoClipTaskResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetVideoClipTaskResponseBody) GetTaskList() []*BatchGetVideoClipTaskResponseBodyTaskList {
	return s.TaskList
}

func (s *BatchGetVideoClipTaskResponseBody) SetRequestId(v string) *BatchGetVideoClipTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBody) SetTaskList(v []*BatchGetVideoClipTaskResponseBodyTaskList) *BatchGetVideoClipTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *BatchGetVideoClipTaskResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetVideoClipTaskResponseBodyTaskList struct {
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 864413342857035776
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 43335
	TotalDuration *float64 `json:"totalDuration,omitempty" xml:"totalDuration,omitempty"`
	// example:
	//
	// 11
	TotalToken *int64                                                `json:"totalToken,omitempty" xml:"totalToken,omitempty"`
	VideoList  []*BatchGetVideoClipTaskResponseBodyTaskListVideoList `json:"videoList,omitempty" xml:"videoList,omitempty" type:"Repeated"`
}

func (s BatchGetVideoClipTaskResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoClipTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) GetStatus() *string {
	return s.Status
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) GetTotalDuration() *float64 {
	return s.TotalDuration
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) GetTotalToken() *int64 {
	return s.TotalToken
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) GetVideoList() []*BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	return s.VideoList
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetStatus(v string) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetTaskId(v string) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetTotalDuration(v float64) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.TotalDuration = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetTotalToken(v int64) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.TotalToken = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) SetVideoList(v []*BatchGetVideoClipTaskResponseBodyTaskListVideoList) *BatchGetVideoClipTaskResponseBodyTaskList {
	s.VideoList = v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskList) Validate() error {
	if s.VideoList != nil {
		for _, item := range s.VideoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetVideoClipTaskResponseBodyTaskListVideoList struct {
	// example:
	//
	// 0
	BeginTime   *int32  `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 11110
	EndTime  *int32  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://e-ai.oss-cn-guangzhou.aliyuncs.com/video/jlkasdl.mp4
	VideoDownloadUrl *string `json:"videoDownloadUrl,omitempty" xml:"videoDownloadUrl,omitempty"`
	VideoName        *string `json:"videoName,omitempty" xml:"videoName,omitempty"`
	// example:
	//
	// https://e-ai.oss-cn-guangzhou.aliyuncs.com/video/jlkasdl.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s BatchGetVideoClipTaskResponseBodyTaskListVideoList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoClipTaskResponseBodyTaskListVideoList) GoString() string {
	return s.String()
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetBeginTime() *int32 {
	return s.BeginTime
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetDescription() *string {
	return s.Description
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetEndTime() *int32 {
	return s.EndTime
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetTitle() *string {
	return s.Title
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetVideoDownloadUrl() *string {
	return s.VideoDownloadUrl
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetVideoName() *string {
	return s.VideoName
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetBeginTime(v int32) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.BeginTime = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetDescription(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.Description = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetEndTime(v int32) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.EndTime = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetErrorMsg(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.ErrorMsg = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetTitle(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.Title = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetVideoDownloadUrl(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.VideoDownloadUrl = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetVideoName(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.VideoName = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) SetVideoUrl(v string) *BatchGetVideoClipTaskResponseBodyTaskListVideoList {
	s.VideoUrl = &v
	return s
}

func (s *BatchGetVideoClipTaskResponseBodyTaskListVideoList) Validate() error {
	return dara.Validate(s)
}
