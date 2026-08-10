// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoTranslationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetVideoTranslationJobResponseBodyJob) *GetVideoTranslationJobResponseBody
	GetJob() *GetVideoTranslationJobResponseBodyJob
	SetRequestId(v string) *GetVideoTranslationJobResponseBody
	GetRequestId() *string
}

type GetVideoTranslationJobResponseBody struct {
	Job *GetVideoTranslationJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoTranslationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTranslationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoTranslationJobResponseBody) GetJob() *GetVideoTranslationJobResponseBodyJob {
	return s.Job
}

func (s *GetVideoTranslationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoTranslationJobResponseBody) SetJob(v *GetVideoTranslationJobResponseBodyJob) *GetVideoTranslationJobResponseBody {
	s.Job = v
	return s
}

func (s *GetVideoTranslationJobResponseBody) SetRequestId(v string) *GetVideoTranslationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoTranslationJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoTranslationJobResponseBodyJob struct {
	// example:
	//
	// 10.0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// ba50304145fd411c827239c398820267
	EditingProjectId *string `json:"EditingProjectId,omitempty" xml:"EditingProjectId,omitempty"`
	// example:
	//
	// InvalidInput
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Input is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// {"Video":"https://example.com/input.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// vtj_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// {"NeedDetext":true,"SubtitleFrom":"default","SourceLanguage":"zh","TargetLanguage":"en","NeedVisualTranslate":true}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// example:
	//
	// VoiceTranslate
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {"AiResult":{"ResultMap":{"ja":{"EditingProjectId":"editing-project-xxx","MediaURL":"https://example.com/bucket/prefix/ja/result.mp4"}}}}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVideoTranslationJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTranslationJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetVideoTranslationJobResponseBodyJob) GetDuration() *float64 {
	return s.Duration
}

func (s *GetVideoTranslationJobResponseBodyJob) GetEditingProjectId() *string {
	return s.EditingProjectId
}

func (s *GetVideoTranslationJobResponseBodyJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoTranslationJobResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoTranslationJobResponseBodyJob) GetInput() *string {
	return s.Input
}

func (s *GetVideoTranslationJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoTranslationJobResponseBodyJob) GetJobParameters() *string {
	return s.JobParameters
}

func (s *GetVideoTranslationJobResponseBodyJob) GetJobType() *string {
	return s.JobType
}

func (s *GetVideoTranslationJobResponseBodyJob) GetOutput() *string {
	return s.Output
}

func (s *GetVideoTranslationJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetVideoTranslationJobResponseBodyJob) SetDuration(v float64) *GetVideoTranslationJobResponseBodyJob {
	s.Duration = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetEditingProjectId(v string) *GetVideoTranslationJobResponseBodyJob {
	s.EditingProjectId = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetErrorCode(v string) *GetVideoTranslationJobResponseBodyJob {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetErrorMessage(v string) *GetVideoTranslationJobResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetInput(v string) *GetVideoTranslationJobResponseBodyJob {
	s.Input = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetJobId(v string) *GetVideoTranslationJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetJobParameters(v string) *GetVideoTranslationJobResponseBodyJob {
	s.JobParameters = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetJobType(v string) *GetVideoTranslationJobResponseBodyJob {
	s.JobType = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetOutput(v string) *GetVideoTranslationJobResponseBodyJob {
	s.Output = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) SetStatus(v string) *GetVideoTranslationJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetVideoTranslationJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
