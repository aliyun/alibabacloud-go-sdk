// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetextJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetVideoDetextJobResponseBodyJob) *GetVideoDetextJobResponseBody
	GetJob() *GetVideoDetextJobResponseBodyJob
	SetRequestId(v string) *GetVideoDetextJobResponseBody
	GetRequestId() *string
}

type GetVideoDetextJobResponseBody struct {
	// The video text erasure task.
	Job *GetVideoDetextJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID, which is used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// req-detext-get-20260820-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoDetextJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetextJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoDetextJobResponseBody) GetJob() *GetVideoDetextJobResponseBodyJob {
	return s.Job
}

func (s *GetVideoDetextJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoDetextJobResponseBody) SetJob(v *GetVideoDetextJobResponseBodyJob) *GetVideoDetextJobResponseBody {
	s.Job = v
	return s
}

func (s *GetVideoDetextJobResponseBody) SetRequestId(v string) *GetVideoDetextJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoDetextJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoDetextJobResponseBodyJob struct {
	// The business error code returned when the task fails. This field is typically not returned for non-failure states.
	//
	// example:
	//
	// InvalidInput
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The business error message returned when the task fails. This field is typically not returned for non-failure states.
	//
	// example:
	//
	// Input video is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The normalized input configuration JSON string saved at submission time.
	//
	// example:
	//
	// {"VideoMediaId":"media-video-001"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The video text erasure task ID.
	//
	// example:
	//
	// vdt_0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The normalized text erasure parameter JSON string.
	//
	// example:
	//
	// {"EraseAllText":false,"TextTargets":[{"Box":[0.1,0.8,0.8,0.15],"TimeRanges":[[0,30]]}]}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The task type. The value is fixed to VIDEO_DETEXT.
	//
	// example:
	//
	// VIDEO_DETEXT
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The task output JSON string. When the task succeeds, AiResult.DetextVideoURL contains the URL of the video with text erased.
	//
	// example:
	//
	// {"AiResult":{"DetextVideoURL":"https://example.com/detext/detext.mp4"}}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The task status. Valid values: Created, Queuing, Executing, Finished, and Failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVideoDetextJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetextJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetVideoDetextJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoDetextJobResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoDetextJobResponseBodyJob) GetInput() *string {
	return s.Input
}

func (s *GetVideoDetextJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoDetextJobResponseBodyJob) GetJobParameters() *string {
	return s.JobParameters
}

func (s *GetVideoDetextJobResponseBodyJob) GetJobType() *string {
	return s.JobType
}

func (s *GetVideoDetextJobResponseBodyJob) GetOutput() *string {
	return s.Output
}

func (s *GetVideoDetextJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetVideoDetextJobResponseBodyJob) SetErrorCode(v string) *GetVideoDetextJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetErrorMessage(v string) *GetVideoDetextJobResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetInput(v string) *GetVideoDetextJobResponseBodyJob {
	s.Input = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetJobId(v string) *GetVideoDetextJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetJobParameters(v string) *GetVideoDetextJobResponseBodyJob {
	s.JobParameters = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetJobType(v string) *GetVideoDetextJobResponseBodyJob {
	s.JobType = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetOutput(v string) *GetVideoDetextJobResponseBodyJob {
	s.Output = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) SetStatus(v string) *GetVideoDetextJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetVideoDetextJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
