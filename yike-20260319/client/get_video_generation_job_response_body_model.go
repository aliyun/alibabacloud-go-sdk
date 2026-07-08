// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVideoGenerationJobResponseBody
	GetRequestId() *string
	SetVideoGenerationJob(v *GetVideoGenerationJobResponseBodyVideoGenerationJob) *GetVideoGenerationJobResponseBody
	GetVideoGenerationJob() *GetVideoGenerationJobResponseBodyVideoGenerationJob
}

type GetVideoGenerationJobResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The video generation task.
	VideoGenerationJob *GetVideoGenerationJobResponseBodyVideoGenerationJob `json:"VideoGenerationJob,omitempty" xml:"VideoGenerationJob,omitempty" type:"Struct"`
}

func (s GetVideoGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoGenerationJobResponseBody) GetVideoGenerationJob() *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	return s.VideoGenerationJob
}

func (s *GetVideoGenerationJobResponseBody) SetRequestId(v string) *GetVideoGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoGenerationJobResponseBody) SetVideoGenerationJob(v *GetVideoGenerationJobResponseBodyVideoGenerationJob) *GetVideoGenerationJobResponseBody {
	s.VideoGenerationJob = v
	return s
}

func (s *GetVideoGenerationJobResponseBody) Validate() error {
	if s.VideoGenerationJob != nil {
		if err := s.VideoGenerationJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoGenerationJobResponseBodyVideoGenerationJob struct {
	// The aspect ratio.
	//
	// example:
	//
	// 16:9
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The video duration. Unit: seconds.
	//
	// example:
	//
	// 5
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The error message. This parameter is returned when the task is in the Failed state.
	//
	// example:
	//
	// Input file not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task input.
	//
	// example:
	//
	// {"Prompt":"图1在篮球场上，用图2来了个灌篮","Medias":[{"Type":"image","Url":"https://xxx/xxx.jpg"},{"Type":"image","Url":"https://xxx/xxx.jpg"}]}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fdc7f121056249c2b64e04bba27bcc8c
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task feature configuration. This parameter does not need to be set.
	//
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The task type.
	//
	// example:
	//
	// first_last_frame
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The model name.
	//
	// example:
	//
	// happyhorse-1.1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of generated videos.
	//
	// example:
	//
	// 1
	N *int32 `json:"N,omitempty" xml:"N,omitempty"`
	// The output result in JsonString format. The following fields are included:
	//
	// Medias: a list of media information (Media objects). The Media object contains the following fields:
	//
	// MediaId: String. The media asset ID.
	//
	// OutputUrl: String. The media URL (with the authentication string).
	//
	// example:
	//
	// {\\"Medias\\":[{\\"MediaId\\":\\"*****470732171f1bfcaf7f6d44*****\\",\\"OutputUrl\\":\\"https://xxxxxxx/.../xxxxx.mp4?Expires=xxxx&OSSAccessKeyId=xxx&Signature=xxxx\\"}]}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The resolution.
	//
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The scene type. Currently, only general is supported.
	//
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The task status. Valid values:
	//
	// - Created: The task is created.
	//
	// - Queuing: The task is queuing.
	//
	// - Executing: The task is being executed.
	//
	// - Finished: The task is completed.
	//
	// - Failed: The task failed.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user business information.
	//
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetVideoGenerationJobResponseBodyVideoGenerationJob) String() string {
	return dara.Prettify(s)
}

func (s GetVideoGenerationJobResponseBodyVideoGenerationJob) GoString() string {
	return s.String()
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetDuration() *string {
	return s.Duration
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetInput() *string {
	return s.Input
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetJobParameters() *string {
	return s.JobParameters
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetJobType() *string {
	return s.JobType
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetModel() *string {
	return s.Model
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetN() *int32 {
	return s.N
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetOutput() *string {
	return s.Output
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetResolution() *string {
	return s.Resolution
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetScene() *string {
	return s.Scene
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetStatus() *string {
	return s.Status
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetAspectRatio(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.AspectRatio = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetDuration(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Duration = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetErrorMessage(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetInput(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Input = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetJobId(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.JobId = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetJobParameters(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.JobParameters = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetJobType(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.JobType = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetModel(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Model = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetN(v int32) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.N = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetOutput(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Output = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetResolution(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Resolution = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetScene(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Scene = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetStatus(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.Status = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) SetUserData(v string) *GetVideoGenerationJobResponseBodyVideoGenerationJob {
	s.UserData = &v
	return s
}

func (s *GetVideoGenerationJobResponseBodyVideoGenerationJob) Validate() error {
	return dara.Validate(s)
}
