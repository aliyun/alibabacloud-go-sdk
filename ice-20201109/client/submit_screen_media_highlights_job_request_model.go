// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitScreenMediaHighlightsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditingConfig(v string) *SubmitScreenMediaHighlightsJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitScreenMediaHighlightsJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitScreenMediaHighlightsJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitScreenMediaHighlightsJobRequest
	GetUserData() *string
}

type SubmitScreenMediaHighlightsJobRequest struct {
	// The editing configuration. For detailed parameters, see [EditingConfig](~~2863940#9b05519d46e0x~~).
	//
	// example:
	//
	// {
	//
	// 	"MediaConfig": {
	//
	// 		"Volume": 1
	//
	// 	},
	//
	// 	"ProcessConfig": {
	//
	// 		"AllowTransition": true,
	//
	// 		"TransitionList": ["fadecolor"]
	//
	// 	}
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// The input configuration. For detailed parameters, see [InputConfig](~~2863940#dda38bf6ec2pk~~).
	//
	// example:
	//
	// {
	//
	// 	"MediaArray": [
	//
	// 		"****9d46c886b45481030f6e****",
	//
	// 		"****6c886b4549d481030f6e****"
	//
	// 	]
	//
	// }
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output configuration. For detailed parameters, see [OutputConfig](~~2863940#4111a373d0xbz~~).
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://xxx.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Count": 1,
	//
	//   "Width": 1080,
	//
	//   "Height": 1920
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The user-defined data, including the business and callback configurations. For more information, see [UserData](https://help.aliyun.com/document_detail/357745.html).
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitScreenMediaHighlightsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitScreenMediaHighlightsJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitScreenMediaHighlightsJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitScreenMediaHighlightsJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitScreenMediaHighlightsJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitScreenMediaHighlightsJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitScreenMediaHighlightsJobRequest) SetEditingConfig(v string) *SubmitScreenMediaHighlightsJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobRequest) SetInputConfig(v string) *SubmitScreenMediaHighlightsJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobRequest) SetOutputConfig(v string) *SubmitScreenMediaHighlightsJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobRequest) SetUserData(v string) *SubmitScreenMediaHighlightsJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitScreenMediaHighlightsJobRequest) Validate() error {
	return dara.Validate(s)
}
