// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikePromptExpansionVoiceFixJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikePromptExpansionVoiceFixJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikePromptExpansionVoiceFixJobResponse
	GetStatusCode() *int32
	SetBody(v *GetYikePromptExpansionVoiceFixJobResponseBody) *GetYikePromptExpansionVoiceFixJobResponse
	GetBody() *GetYikePromptExpansionVoiceFixJobResponseBody
}

type GetYikePromptExpansionVoiceFixJobResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikePromptExpansionVoiceFixJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikePromptExpansionVoiceFixJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikePromptExpansionVoiceFixJobResponse) GoString() string {
	return s.String()
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) GetBody() *GetYikePromptExpansionVoiceFixJobResponseBody {
	return s.Body
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) SetHeaders(v map[string]*string) *GetYikePromptExpansionVoiceFixJobResponse {
	s.Headers = v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) SetStatusCode(v int32) *GetYikePromptExpansionVoiceFixJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) SetBody(v *GetYikePromptExpansionVoiceFixJobResponseBody) *GetYikePromptExpansionVoiceFixJobResponse {
	s.Body = v
	return s
}

func (s *GetYikePromptExpansionVoiceFixJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
