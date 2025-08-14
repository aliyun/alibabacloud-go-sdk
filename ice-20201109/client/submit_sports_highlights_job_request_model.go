// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSportsHighlightsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SubmitSportsHighlightsJobRequest
	GetClientToken() *string
	SetInputConfig(v string) *SubmitSportsHighlightsJobRequest
	GetInputConfig() *string
	SetOutputConfig(v string) *SubmitSportsHighlightsJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitSportsHighlightsJobRequest
	GetUserData() *string
}

type SubmitSportsHighlightsJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The input configurations.
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// The output configurations.
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
	// The user-defined data.
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSportsHighlightsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSportsHighlightsJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSportsHighlightsJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SubmitSportsHighlightsJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitSportsHighlightsJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitSportsHighlightsJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSportsHighlightsJobRequest) SetClientToken(v string) *SubmitSportsHighlightsJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitSportsHighlightsJobRequest) SetInputConfig(v string) *SubmitSportsHighlightsJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSportsHighlightsJobRequest) SetOutputConfig(v string) *SubmitSportsHighlightsJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSportsHighlightsJobRequest) SetUserData(v string) *SubmitSportsHighlightsJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSportsHighlightsJobRequest) Validate() error {
	return dara.Validate(s)
}
