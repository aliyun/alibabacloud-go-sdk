// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRemakeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetVideoRemakeJobResponseBodyJob) *GetVideoRemakeJobResponseBody
	GetJob() *GetVideoRemakeJobResponseBodyJob
	SetRequestId(v string) *GetVideoRemakeJobResponseBody
	GetRequestId() *string
}

type GetVideoRemakeJobResponseBody struct {
	// The unified remake task details.
	Job *GetVideoRemakeJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request tracking identifier.
	//
	// example:
	//
	// 12345678-1234-4321-ABCD-123456789012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoRemakeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRemakeJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoRemakeJobResponseBody) GetJob() *GetVideoRemakeJobResponseBodyJob {
	return s.Job
}

func (s *GetVideoRemakeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoRemakeJobResponseBody) SetJob(v *GetVideoRemakeJobResponseBodyJob) *GetVideoRemakeJobResponseBody {
	s.Job = v
	return s
}

func (s *GetVideoRemakeJobResponseBody) SetRequestId(v string) *GetVideoRemakeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoRemakeJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoRemakeJobResponseBodyJob struct {
	// The error code for the failed asynchronous task. This parameter has a value only when the task status is Failed.
	//
	// example:
	//
	// NOT_ENOUGH_ALLOCATE_CREDIT_QUOTA
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message for the failed asynchronous task. This parameter has a value only when the task status is Failed.
	//
	// example:
	//
	// lockCredit business failed: NOT_ENOUGH_ALLOCATE_CREDIT_QUOTA
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unified remake task ID.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The content understanding and script outputs returned when the status is Finished. This is not the rendered video.
	JobResult *GetVideoRemakeJobResponseBodyJobJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Struct"`
	// The list of input media asset IDs.
	//
	// example:
	//
	// ["1234567890abcdef1234567890abcdef"]
	MediaIds []*string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty" type:"Repeated"`
	// The task status. Valid values: Created, Executing, Finished, and Failed. Finished indicates that content understanding and script generation are completed. If the status is Failed, check ErrorCode and ErrorMessage. If the task has been deleted, the query returns RecordNotFound.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UserData JSON string submitted with the task, returned as is.
	//
	// example:
	//
	// {"NotifyAddress":"https://example.com/callback"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetVideoRemakeJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRemakeJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetVideoRemakeJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoRemakeJobResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoRemakeJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoRemakeJobResponseBodyJob) GetJobResult() *GetVideoRemakeJobResponseBodyJobJobResult {
	return s.JobResult
}

func (s *GetVideoRemakeJobResponseBodyJob) GetMediaIds() []*string {
	return s.MediaIds
}

func (s *GetVideoRemakeJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetVideoRemakeJobResponseBodyJob) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoRemakeJobResponseBodyJob) SetErrorCode(v string) *GetVideoRemakeJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) SetErrorMessage(v string) *GetVideoRemakeJobResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) SetJobId(v string) *GetVideoRemakeJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) SetJobResult(v *GetVideoRemakeJobResponseBodyJobJobResult) *GetVideoRemakeJobResponseBodyJob {
	s.JobResult = v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) SetMediaIds(v []*string) *GetVideoRemakeJobResponseBodyJob {
	s.MediaIds = v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) SetStatus(v string) *GetVideoRemakeJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) SetUserData(v string) *GetVideoRemakeJobResponseBodyJob {
	s.UserData = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJob) Validate() error {
	if s.JobResult != nil {
		if err := s.JobResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoRemakeJobResponseBodyJobJobResult struct {
	// The download URL of the JSON file for the content understanding result.
	//
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/comprehension-result.json
	AiResult *string `json:"AiResult,omitempty" xml:"AiResult,omitempty"`
	// The download URL of the JSON file for the creative script. Use the file content as the Script parameter for SubmitVideoRenderJob.
	//
	// example:
	//
	// https://example.com/creative-script.json
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetVideoRemakeJobResponseBodyJobJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRemakeJobResponseBodyJobJobResult) GoString() string {
	return s.String()
}

func (s *GetVideoRemakeJobResponseBodyJobJobResult) GetAiResult() *string {
	return s.AiResult
}

func (s *GetVideoRemakeJobResponseBodyJobJobResult) GetScript() *string {
	return s.Script
}

func (s *GetVideoRemakeJobResponseBodyJobJobResult) SetAiResult(v string) *GetVideoRemakeJobResponseBodyJobJobResult {
	s.AiResult = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJobJobResult) SetScript(v string) *GetVideoRemakeJobResponseBodyJobJobResult {
	s.Script = &v
	return s
}

func (s *GetVideoRemakeJobResponseBodyJobJobResult) Validate() error {
	return dara.Validate(s)
}
