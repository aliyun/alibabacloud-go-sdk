// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikePromptExpansionVoiceFixJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitYikePromptExpansionVoiceFixJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitYikePromptExpansionVoiceFixJobResponseBody
	GetRequestId() *string
}

type SubmitYikePromptExpansionVoiceFixJobResponseBody struct {
	// example:
	//
	// 68ca759e798b40b4903b255********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitYikePromptExpansionVoiceFixJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikePromptExpansionVoiceFixJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponseBody) SetJobId(v string) *SubmitYikePromptExpansionVoiceFixJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponseBody) SetRequestId(v string) *SubmitYikePromptExpansionVoiceFixJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponseBody) Validate() error {
	return dara.Validate(s)
}
