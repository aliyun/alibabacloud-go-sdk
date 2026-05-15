// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikePromptExpansionVoiceFixJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikePromptExpansionVoiceFixJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikePromptExpansionVoiceFixJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikePromptExpansionVoiceFixJobResponseBody) *SubmitYikePromptExpansionVoiceFixJobResponse
	GetBody() *SubmitYikePromptExpansionVoiceFixJobResponseBody
}

type SubmitYikePromptExpansionVoiceFixJobResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikePromptExpansionVoiceFixJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikePromptExpansionVoiceFixJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikePromptExpansionVoiceFixJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) GetBody() *SubmitYikePromptExpansionVoiceFixJobResponseBody {
	return s.Body
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) SetHeaders(v map[string]*string) *SubmitYikePromptExpansionVoiceFixJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) SetStatusCode(v int32) *SubmitYikePromptExpansionVoiceFixJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) SetBody(v *SubmitYikePromptExpansionVoiceFixJobResponseBody) *SubmitYikePromptExpansionVoiceFixJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikePromptExpansionVoiceFixJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
