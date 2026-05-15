// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikePromptExpansionVoiceFixJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikePromptExpansionVoiceFixJobRequest
	GetJobId() *string
}

type GetYikePromptExpansionVoiceFixJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikePromptExpansionVoiceFixJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikePromptExpansionVoiceFixJobRequest) GoString() string {
	return s.String()
}

func (s *GetYikePromptExpansionVoiceFixJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikePromptExpansionVoiceFixJobRequest) SetJobId(v string) *GetYikePromptExpansionVoiceFixJobRequest {
	s.JobId = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobRequest) Validate() error {
	return dara.Validate(s)
}
