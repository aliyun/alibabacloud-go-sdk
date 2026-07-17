// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageGenerationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageGenerationJob(v *GetImageGenerationJobResponseBodyImageGenerationJob) *GetImageGenerationJobResponseBody
	GetImageGenerationJob() *GetImageGenerationJobResponseBodyImageGenerationJob
	SetRequestId(v string) *GetImageGenerationJobResponseBody
	GetRequestId() *string
}

type GetImageGenerationJobResponseBody struct {
	// The image generation task.
	ImageGenerationJob *GetImageGenerationJobResponseBodyImageGenerationJob `json:"ImageGenerationJob,omitempty" xml:"ImageGenerationJob,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageGenerationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageGenerationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageGenerationJobResponseBody) GetImageGenerationJob() *GetImageGenerationJobResponseBodyImageGenerationJob {
	return s.ImageGenerationJob
}

func (s *GetImageGenerationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageGenerationJobResponseBody) SetImageGenerationJob(v *GetImageGenerationJobResponseBodyImageGenerationJob) *GetImageGenerationJobResponseBody {
	s.ImageGenerationJob = v
	return s
}

func (s *GetImageGenerationJobResponseBody) SetRequestId(v string) *GetImageGenerationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageGenerationJobResponseBody) Validate() error {
	if s.ImageGenerationJob != nil {
		if err := s.ImageGenerationJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageGenerationJobResponseBodyImageGenerationJob struct {
	// The video aspect ratio.
	//
	// example:
	//
	// 16:9
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The error message. This parameter is returned only when the task is in the Failed state.
	//
	// example:
	//
	// Input file not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The task input.
	//
	// example:
	//
	// {"Prompt":"图1在篮球场上，用图2来了个灌篮"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fdc7f121056249c2b64e04bba27bcc8c
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task feature configuration. No configuration is required at this time.
	//
	// example:
	//
	// {}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The task type.
	//
	// example:
	//
	// text_to_image
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The model name.
	//
	// example:
	//
	// wan2.7-image
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of generated images.
	//
	// example:
	//
	// 1
	N *string `json:"N,omitempty" xml:"N,omitempty"`
	// The generation result in JSON string format. Fields:
	//
	// - Medias: a list of media information (Media objects). Fields of a Media object:
	//
	//   - MediaId: String. The media asset ID.
	//
	//   - OutputUrl: String. The media URL (with authentication string).
	//
	// example:
	//
	// {"Medias":[{"MediaId":"***e3700761971f19c32e7e7d5496***","OutputUrl":"https://**bucket**.oss-ap-southeast-1.aliyuncs.com/xxx.prd"}]}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The resolution of the generated video.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The scenario type. Currently, only `general` is supported.
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
	// The custom business information.
	//
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetImageGenerationJobResponseBodyImageGenerationJob) String() string {
	return dara.Prettify(s)
}

func (s GetImageGenerationJobResponseBodyImageGenerationJob) GoString() string {
	return s.String()
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetInput() *string {
	return s.Input
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetJobId() *string {
	return s.JobId
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetJobParameters() *string {
	return s.JobParameters
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetJobType() *string {
	return s.JobType
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetModel() *string {
	return s.Model
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetN() *string {
	return s.N
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetOutput() *string {
	return s.Output
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetResolution() *string {
	return s.Resolution
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetScene() *string {
	return s.Scene
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetStatus() *string {
	return s.Status
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) GetUserData() *string {
	return s.UserData
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetAspectRatio(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.AspectRatio = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetErrorMessage(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetInput(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.Input = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetJobId(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.JobId = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetJobParameters(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.JobParameters = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetJobType(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.JobType = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetModel(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.Model = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetN(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.N = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetOutput(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.Output = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetResolution(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.Resolution = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetScene(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.Scene = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetStatus(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.Status = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) SetUserData(v string) *GetImageGenerationJobResponseBodyImageGenerationJob {
	s.UserData = &v
	return s
}

func (s *GetImageGenerationJobResponseBodyImageGenerationJob) Validate() error {
	return dara.Validate(s)
}
