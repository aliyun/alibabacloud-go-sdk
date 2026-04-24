// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeVoiceNarratorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobParams(v string) *SubmitYikeVoiceNarratorJobRequest
	GetJobParams() *string
	SetUserData(v string) *SubmitYikeVoiceNarratorJobRequest
	GetUserData() *string
}

type SubmitYikeVoiceNarratorJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {\\"TextType\\":2,\\"TextContent\\":\\"Today, Beijing held a press conference to announce plans to further optimize the city\\"s transportation network, including adding three new subway lines within the next three years....\\",\\"AspectRatio\\":\\"16:9\\", \\"Resolution\\":\\"720P\\", \\"OutputLanguages\\":[\\"CN\\",\\"YUE\\"]"}
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	// example:
	//
	// {\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://cms.example.com/callback/video-task\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitYikeVoiceNarratorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeVoiceNarratorJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikeVoiceNarratorJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitYikeVoiceNarratorJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitYikeVoiceNarratorJobRequest) SetJobParams(v string) *SubmitYikeVoiceNarratorJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitYikeVoiceNarratorJobRequest) SetUserData(v string) *SubmitYikeVoiceNarratorJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitYikeVoiceNarratorJobRequest) Validate() error {
	return dara.Validate(s)
}
