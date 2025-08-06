// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaRefreshJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaRefreshJobs(v []*GetMediaRefreshJobsResponseBodyMediaRefreshJobs) *GetMediaRefreshJobsResponseBody
	GetMediaRefreshJobs() []*GetMediaRefreshJobsResponseBodyMediaRefreshJobs
	SetRequestId(v string) *GetMediaRefreshJobsResponseBody
	GetRequestId() *string
}

type GetMediaRefreshJobsResponseBody struct {
	// The media refresh or prefetch jobs.
	MediaRefreshJobs []*GetMediaRefreshJobsResponseBodyMediaRefreshJobs `json:"MediaRefreshJobs,omitempty" xml:"MediaRefreshJobs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaRefreshJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaRefreshJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaRefreshJobsResponseBody) GetMediaRefreshJobs() []*GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	return s.MediaRefreshJobs
}

func (s *GetMediaRefreshJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaRefreshJobsResponseBody) SetMediaRefreshJobs(v []*GetMediaRefreshJobsResponseBodyMediaRefreshJobs) *GetMediaRefreshJobsResponseBody {
	s.MediaRefreshJobs = v
	return s
}

func (s *GetMediaRefreshJobsResponseBody) SetRequestId(v string) *GetMediaRefreshJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaRefreshJobsResponseBodyMediaRefreshJobs struct {
	// The error code. This parameter is returned if the refresh or prefetch task fails.
	//
	// example:
	//
	// PreloadQueueFull
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. This parameter is returned if the refresh or prefetch task fails.
	//
	// example:
	//
	// Preload queue is full, please try again later!
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The filtering conditions for stream playback. The value is a JSON string. This parameter is used as a request parameter of the [RefreshMediaPlayUrls](~~RefreshMediaPlayUrls~~) operation.
	//
	// example:
	//
	// {"Formats":"mp4,m3u8", "Definitions":"HD,SD",  " StreamType":"video",  "ResultType":"Single",  " SliceFlag":false, "SliceCount": 3}
	FilterPolicy *string `json:"FilterPolicy,omitempty" xml:"FilterPolicy,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-05-20T08:23:22Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2022-05-21T08:23:22Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the media file.
	//
	// example:
	//
	// ca3a8f6e4957b658067095869****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 41d465e31957****
	MediaRefreshJobId *string `json:"MediaRefreshJobId,omitempty" xml:"MediaRefreshJobId,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The playback URLs that were refreshed or prefetched.
	//
	// example:
	//
	// https://shenzhen.****.aliyuncdn.com/74401a4f546007bf845cd8840****.m3u8,https://shenzhen.****.aliyuncdn.com/24041e7d13582d86604d8****.m3u8
	SuccessPlayUrls *string `json:"SuccessPlayUrls,omitempty" xml:"SuccessPlayUrls,omitempty"`
	// The IDs of the refresh or prefetch tasks for the playback URLs of media files. Only one URL can be refreshed or prefetched in a task. This value is used in the [DescribeVodRefreshTasks](~~DescribeVodRefreshTasks~~) operation, which queries the status of refresh or prefetch tasks for playback URLs of media files.
	//
	// example:
	//
	// 70422****,9524****
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	// The type of the job. Valid values:
	//
	// 	- **Refresh**
	//
	// 	- **Preload**
	//
	// example:
	//
	// Preload
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The user data that you passed when you submit a refresh or prefetch task.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"}, "Extend":{"localId":"xxx","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetMediaRefreshJobsResponseBodyMediaRefreshJobs) String() string {
	return dara.Prettify(s)
}

func (s GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GoString() string {
	return s.String()
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetFilterPolicy() *string {
	return s.FilterPolicy
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetMediaRefreshJobId() *string {
	return s.MediaRefreshJobId
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetStatus() *string {
	return s.Status
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetSuccessPlayUrls() *string {
	return s.SuccessPlayUrls
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetTaskIds() *string {
	return s.TaskIds
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetTaskType() *string {
	return s.TaskType
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) GetUserData() *string {
	return s.UserData
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetErrorCode(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.ErrorCode = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetErrorMessage(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.ErrorMessage = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetFilterPolicy(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.FilterPolicy = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetGmtCreate(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.GmtCreate = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetGmtModified(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.GmtModified = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetMediaId(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.MediaId = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetMediaRefreshJobId(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.MediaRefreshJobId = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetStatus(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.Status = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetSuccessPlayUrls(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.SuccessPlayUrls = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetTaskIds(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.TaskIds = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetTaskType(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.TaskType = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) SetUserData(v string) *GetMediaRefreshJobsResponseBodyMediaRefreshJobs {
	s.UserData = &v
	return s
}

func (s *GetMediaRefreshJobsResponseBodyMediaRefreshJobs) Validate() error {
	return dara.Validate(s)
}
