// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeVoiceNarratorJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikeVoiceNarratorJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikeVoiceNarratorJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikeVoiceNarratorJobResponseBody) *SubmitYikeVoiceNarratorJobResponse
	GetBody() *SubmitYikeVoiceNarratorJobResponseBody
}

type SubmitYikeVoiceNarratorJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikeVoiceNarratorJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikeVoiceNarratorJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeVoiceNarratorJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikeVoiceNarratorJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikeVoiceNarratorJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikeVoiceNarratorJobResponse) GetBody() *SubmitYikeVoiceNarratorJobResponseBody {
	return s.Body
}

func (s *SubmitYikeVoiceNarratorJobResponse) SetHeaders(v map[string]*string) *SubmitYikeVoiceNarratorJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikeVoiceNarratorJobResponse) SetStatusCode(v int32) *SubmitYikeVoiceNarratorJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikeVoiceNarratorJobResponse) SetBody(v *SubmitYikeVoiceNarratorJobResponseBody) *SubmitYikeVoiceNarratorJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikeVoiceNarratorJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
