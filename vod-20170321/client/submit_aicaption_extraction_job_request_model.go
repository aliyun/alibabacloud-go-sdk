// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAICaptionExtractionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIPipelineId(v string) *SubmitAICaptionExtractionJobRequest
	GetAIPipelineId() *string
	SetJobConfig(v string) *SubmitAICaptionExtractionJobRequest
	GetJobConfig() *string
	SetUserData(v string) *SubmitAICaptionExtractionJobRequest
	GetUserData() *string
	SetVideoId(v string) *SubmitAICaptionExtractionJobRequest
	GetVideoId() *string
}

type SubmitAICaptionExtractionJobRequest struct {
	AIPipelineId *string `json:"AIPipelineId,omitempty" xml:"AIPipelineId,omitempty"`
	// This parameter is required.
	JobConfig *string `json:"JobConfig,omitempty" xml:"JobConfig,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// This parameter is required.
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitAICaptionExtractionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICaptionExtractionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAICaptionExtractionJobRequest) GetAIPipelineId() *string {
	return s.AIPipelineId
}

func (s *SubmitAICaptionExtractionJobRequest) GetJobConfig() *string {
	return s.JobConfig
}

func (s *SubmitAICaptionExtractionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAICaptionExtractionJobRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitAICaptionExtractionJobRequest) SetAIPipelineId(v string) *SubmitAICaptionExtractionJobRequest {
	s.AIPipelineId = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) SetJobConfig(v string) *SubmitAICaptionExtractionJobRequest {
	s.JobConfig = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) SetUserData(v string) *SubmitAICaptionExtractionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) SetVideoId(v string) *SubmitAICaptionExtractionJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) Validate() error {
	return dara.Validate(s)
}
