// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoTranslationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitVideoTranslationJobRequest
	GetClientToken() *string
	SetDescription(v string) *SubmitVideoTranslationJobRequest
	GetDescription() *string
	SetInput(v string) *SubmitVideoTranslationJobRequest
	GetInput() *string
	SetJobParameters(v string) *SubmitVideoTranslationJobRequest
	GetJobParameters() *string
	SetJobType(v string) *SubmitVideoTranslationJobRequest
	GetJobType() *string
	SetOutput(v string) *SubmitVideoTranslationJobRequest
	GetOutput() *string
	SetTitle(v string) *SubmitVideoTranslationJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitVideoTranslationJobRequest
	GetUserData() *string
}

type SubmitVideoTranslationJobRequest struct {
	// The user-level idempotency token, which can be up to 40 characters in length. If the same user submits the request repeatedly with the same token, the original task is returned.
	//
	// example:
	//
	// vt-client-20260820-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The job description, which is used to record the business purpose or processing requirements.
	//
	// example:
	//
	// Translate the Chinese product introduction video into English
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input configuration JSON string. You must specify either Video or VideoMediaId, but not both. You can specify at most one of Audio and AudioMediaId. Subtitle is optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"VideoMediaId":"media-video-001"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The job parameters JSON string. It must include at least SourceLanguage and TargetLanguage. You can also configure primary subtitle erasure, voice translation, on-screen text translation, and final editing.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"SourceLanguage":"zh","TargetLanguage":"en","NeedDetext":false,"NeedVisualTranslate":false}
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
	// The job type. SubtitleTranslate indicates subtitle translation, and VoiceTranslate indicates voice translation.
	//
	// This parameter is required.
	//
	// example:
	//
	// VoiceTranslate
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The output configuration JSON string. OssUri is the optional customer OSS output directory. If OssUri is not specified, a signed URL for the service-owned output is returned.
	//
	// example:
	//
	// {"OssUri":"oss://example-bucket/video-translation/output/"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The job title. If not specified, the service generates a default title.
	//
	// example:
	//
	// English translation of product introduction video
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data JSON string. When specifying an MNS queue to receive asynchronous notifications, set `NotifyAddress` to the pure queue name starting with `yike-callback-` instead of the full queue URL. You can optionally set `NotifyMnsEndpoint` to the MNS endpoint of the task account.
	//
	// example:
	//
	// {"NotifyAddress":"yike-callback-test","NotifyMnsEndpoint":"https://1234567890123456.mns.cn-shanghai.aliyuncs.com"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoTranslationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoTranslationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoTranslationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitVideoTranslationJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitVideoTranslationJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitVideoTranslationJobRequest) GetJobParameters() *string {
	return s.JobParameters
}

func (s *SubmitVideoTranslationJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitVideoTranslationJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitVideoTranslationJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitVideoTranslationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoTranslationJobRequest) SetClientToken(v string) *SubmitVideoTranslationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetDescription(v string) *SubmitVideoTranslationJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetInput(v string) *SubmitVideoTranslationJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetJobParameters(v string) *SubmitVideoTranslationJobRequest {
	s.JobParameters = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetJobType(v string) *SubmitVideoTranslationJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetOutput(v string) *SubmitVideoTranslationJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetTitle(v string) *SubmitVideoTranslationJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) SetUserData(v string) *SubmitVideoTranslationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoTranslationJobRequest) Validate() error {
	return dara.Validate(s)
}
