// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchMediaProducingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitBatchMediaProducingJobRequest
	GetClientToken() *string
	SetEditingConfig(v string) *SubmitBatchMediaProducingJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitBatchMediaProducingJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitBatchMediaProducingJobRequest
	GetOutputConfig() *string
	SetTemplateConfig(v string) *SubmitBatchMediaProducingJobRequest
	GetTemplateConfig() *string
	SetUserData(v string) *SubmitBatchMediaProducingJobRequest
	GetUserData() *string
}

type SubmitBatchMediaProducingJobRequest struct {
	// The client token that is used to ensure the idempotency of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The editing configuration. For the specific structure definition, see [EditingConfig](~~2692547#1be9bba03b7qu~~).
	//
	// example:
	//
	// {
	//
	//   "MediaConfig": {
	//
	//       "Volume": 0
	//
	//   },
	//
	//   "SpeechConfig": {
	//
	//       "Volume": 1
	//
	//   },
	//
	//  "BackgroundMusicConfig": {
	//
	//       "Volume": 0.3
	//
	//   }
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configuration. For the specific structure definition, see [InputConfig](~~2692547#2faed1559549n~~).
	//
	// example:
	//
	// {
	//
	//   "MediaGroupArray": [{
	//
	//       "GroupName": "MediaGroup1",
	//
	//       "MediaArray": [
	//
	//         "****9d46c886b45481030f6e****",
	//
	//         "****6c886b4549d481030f6e****" ]
	//
	//     }, {
	//
	//       "GroupName": "MediaGroup2",
	//
	//       "MediaArray": [
	//
	//         "****d46c886810b454930f6e****",
	//
	//         "****4549d886810b46c30f6e****" ]
	//
	//   }],
	//
	//   "TitleArray": [
	//
	//       "Hema Fresh grand opening in Huilongguan",
	//
	//       "Hema Fresh grand opening" ],
	//
	//   "SpeechTextArray": [
	//
	//       "A new Hema Fresh store just opened in the nearby mall, today is the first day of business"
	//
	//       "There are quite a few people in the mall, snacks and beverages are relatively cheap, come check it out" ]
	//
	// }
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output configuration. For the specific structure definition, see [OutputConfig](~~2692547#447b928fcbuoa~~).
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://xxx.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Count": 20,
	//
	//   "MaxDuration": 15,
	//
	//   "Width": 1080,
	//
	//   "Height": 1920,
	//
	//   "Video": {"Crf": 27}
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The template parameters. You can configure multiple one-click video production templates, and one is randomly selected and applied. For details, see [TemplateConfig metric description](https://www.alibabacloud.com/help/en/ims/use-cases/batch-video-production-public-parameters#32c3bea6182sy).
	//
	// example:
	//
	// ["****b4549d46c88681030f6e****","****549d46c88b4681030f6e****"]
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The user business configuration and callback configuration. For the specific structure definition, see [UserData configuration](~~357745#section-urj-v3f-0s1~~).
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"} or {"NotifyAddress":"https://xx.xx.xxx"} or {"NotifyAddress":"ice-callback-demo"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitBatchMediaProducingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchMediaProducingJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitBatchMediaProducingJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitBatchMediaProducingJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitBatchMediaProducingJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitBatchMediaProducingJobRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *SubmitBatchMediaProducingJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitBatchMediaProducingJobRequest) SetClientToken(v string) *SubmitBatchMediaProducingJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitBatchMediaProducingJobRequest) SetEditingConfig(v string) *SubmitBatchMediaProducingJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitBatchMediaProducingJobRequest) SetInputConfig(v string) *SubmitBatchMediaProducingJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitBatchMediaProducingJobRequest) SetOutputConfig(v string) *SubmitBatchMediaProducingJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitBatchMediaProducingJobRequest) SetTemplateConfig(v string) *SubmitBatchMediaProducingJobRequest {
	s.TemplateConfig = &v
	return s
}

func (s *SubmitBatchMediaProducingJobRequest) SetUserData(v string) *SubmitBatchMediaProducingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitBatchMediaProducingJobRequest) Validate() error {
	return dara.Validate(s)
}
