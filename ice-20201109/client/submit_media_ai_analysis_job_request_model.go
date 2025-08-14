// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaAiAnalysisJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisParams(v string) *SubmitMediaAiAnalysisJobRequest
	GetAnalysisParams() *string
	SetInput(v string) *SubmitMediaAiAnalysisJobRequest
	GetInput() *string
	SetUserData(v string) *SubmitMediaAiAnalysisJobRequest
	GetUserData() *string
}

type SubmitMediaAiAnalysisJobRequest struct {
	// The analysis parameters.
	//
	// example:
	//
	// {"nlpParams":{"sourceLanguage":"cn","diarizationEnabled":true,"speakerCount":0,"summarizationEnabled":false,"translationEnabled":false}}
	AnalysisParams *string `json:"AnalysisParams,omitempty" xml:"AnalysisParams,omitempty"`
	// The media asset that you want to analyze. You can specify an Object Storage Service (OSS) URL, a media asset ID, or an external URL.
	//
	// example:
	//
	// {"MediaType":"video","Media":"https://xxx.com/your_movie.mp4"}
	Input    *string `json:"Input,omitempty" xml:"Input,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaAiAnalysisJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaAiAnalysisJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaAiAnalysisJobRequest) GetAnalysisParams() *string {
	return s.AnalysisParams
}

func (s *SubmitMediaAiAnalysisJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitMediaAiAnalysisJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaAiAnalysisJobRequest) SetAnalysisParams(v string) *SubmitMediaAiAnalysisJobRequest {
	s.AnalysisParams = &v
	return s
}

func (s *SubmitMediaAiAnalysisJobRequest) SetInput(v string) *SubmitMediaAiAnalysisJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitMediaAiAnalysisJobRequest) SetUserData(v string) *SubmitMediaAiAnalysisJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaAiAnalysisJobRequest) Validate() error {
	return dara.Validate(s)
}
