// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIJobList(v *ListAIJobResponseBodyAIJobList) *ListAIJobResponseBody
	GetAIJobList() *ListAIJobResponseBodyAIJobList
	SetNonExistAIJobIds(v *ListAIJobResponseBodyNonExistAIJobIds) *ListAIJobResponseBody
	GetNonExistAIJobIds() *ListAIJobResponseBodyNonExistAIJobIds
	SetRequestId(v string) *ListAIJobResponseBody
	GetRequestId() *string
}

type ListAIJobResponseBody struct {
	// The list of jobs.
	AIJobList *ListAIJobResponseBodyAIJobList `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	// The IDs of the jobs that do not exist.
	NonExistAIJobIds *ListAIJobResponseBodyNonExistAIJobIds `json:"NonExistAIJobIds,omitempty" xml:"NonExistAIJobIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8233A0E4-E112-44*****58-2BCED1B88173
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBody) GetAIJobList() *ListAIJobResponseBodyAIJobList {
	return s.AIJobList
}

func (s *ListAIJobResponseBody) GetNonExistAIJobIds() *ListAIJobResponseBodyNonExistAIJobIds {
	return s.NonExistAIJobIds
}

func (s *ListAIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIJobResponseBody) SetAIJobList(v *ListAIJobResponseBodyAIJobList) *ListAIJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *ListAIJobResponseBody) SetNonExistAIJobIds(v *ListAIJobResponseBodyNonExistAIJobIds) *ListAIJobResponseBody {
	s.NonExistAIJobIds = v
	return s
}

func (s *ListAIJobResponseBody) SetRequestId(v string) *ListAIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIJobResponseBodyAIJobList struct {
	AIJob []*ListAIJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s ListAIJobResponseBodyAIJobList) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyAIJobList) GetAIJob() []*ListAIJobResponseBodyAIJobListAIJob {
	return s.AIJob
}

func (s *ListAIJobResponseBodyAIJobList) SetAIJob(v []*ListAIJobResponseBodyAIJobListAIJob) *ListAIJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

func (s *ListAIJobResponseBodyAIJobList) Validate() error {
	return dara.Validate(s)
}

type ListAIJobResponseBodyAIJobListAIJob struct {
	// The error code. This parameter is returned if the value of Status is fail.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-28T02:04:47Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the job was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-28T02:04:32Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The returned data. The value is a JSON string. For more information, see [AITemplateConfig](https://help.aliyun.com/document_detail/89863.html).
	//
	// example:
	//
	// {"OrigASRData":{"AsrTextList":[{"EndTime":700,"StartTime":0,"Text":"Yes.","ChannelId":0,"SpeechRate":85},{"EndTime":3750,"StartTime":1630,"Text":"No.","ChannelId":0,"SpeechRate":28},{"EndTime":5910,"StartTime":4020,"Text":"Of course.","ChannelId":0,"SpeechRate":95},{"EndTime":12750,"StartTime":10090,"Text":"Message.","ChannelId":0,"SpeechRate":45},{"EndTime":25230,"StartTime":13590,"Text":"Hello, good afternoon.","ChannelId":0,"SpeechRate":20},{"EndTime":30000,"StartTime":28220,"Text":"Yes.","ChannelId":0,"SpeechRate":33}],"Duration":"30016"}}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The job ID.
	//
	// example:
	//
	// a718a3a1e8bb42ee3bc88921e94****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the video file.
	//
	// example:
	//
	// 3D3D12340d9401fab46a0b847****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The error message. This parameter is returned if the value of Status is fail.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **success**: The job is successful.
	//
	// 	- **fail**: The job failed.
	//
	// 	- **init**: The job is being initialized.
	//
	// 	- **Processing**: The job is in progress.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the job. Valid values:
	//
	// 	- **AIMediaDNA**: video fingerprinting
	//
	// 	- **AIVideoTag**: smart tagging
	//
	// example:
	//
	// AIVideoTag
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAIJobResponseBodyAIJobListAIJob) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetCode() *string {
	return s.Code
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetData() *string {
	return s.Data
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetJobId() *string {
	return s.JobId
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetMediaId() *string {
	return s.MediaId
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetMessage() *string {
	return s.Message
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetStatus() *string {
	return s.Status
}

func (s *ListAIJobResponseBodyAIJobListAIJob) GetType() *string {
	return s.Type
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCode(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Code = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCompleteTime(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.CompleteTime = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCreationTime(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetData(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Data = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetJobId(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetMediaId(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetMessage(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Message = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetStatus(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Status = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetType(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Type = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) Validate() error {
	return dara.Validate(s)
}

type ListAIJobResponseBodyNonExistAIJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIJobResponseBodyNonExistAIJobIds) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponseBodyNonExistAIJobIds) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) GetString_() []*string {
	return s.String_
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) SetString_(v []*string) *ListAIJobResponseBodyNonExistAIJobIds {
	s.String_ = v
	return s
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) Validate() error {
	return dara.Validate(s)
}
