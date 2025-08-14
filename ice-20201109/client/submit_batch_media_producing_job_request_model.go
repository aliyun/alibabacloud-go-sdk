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
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The editing configurations. For more information, see [EditingConfig](~~2692547#1be9bba03b7qu~~).
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
	// The input configurations. For more information, see [InputConfig](~~2692547#2faed1559549n~~).
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output configurations. For more information, see [OutputConfig](~~2692547#447b928fcbuoa~~).
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
	OutputConfig   *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The user-defined data, including the business and callback configurations. For more information, see [UserData](https://help.aliyun.com/document_detail/357745.html).
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
