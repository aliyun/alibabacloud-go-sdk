// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikePromptExpansionVoiceFixJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobParams(v string) *SubmitYikePromptExpansionVoiceFixJobRequest
	GetJobParams() *string
	SetUserData(v string) *SubmitYikePromptExpansionVoiceFixJobRequest
	GetUserData() *string
}

type SubmitYikePromptExpansionVoiceFixJobRequest struct {
	// This parameter is required.
	JobParams *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitYikePromptExpansionVoiceFixJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikePromptExpansionVoiceFixJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikePromptExpansionVoiceFixJobRequest) GetJobParams() *string {
	return s.JobParams
}

func (s *SubmitYikePromptExpansionVoiceFixJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitYikePromptExpansionVoiceFixJobRequest) SetJobParams(v string) *SubmitYikePromptExpansionVoiceFixJobRequest {
	s.JobParams = &v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobRequest) SetUserData(v string) *SubmitYikePromptExpansionVoiceFixJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobRequest) Validate() error {
	return dara.Validate(s)
}
