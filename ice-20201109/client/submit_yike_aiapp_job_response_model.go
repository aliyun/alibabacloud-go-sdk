// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeAIAppJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikeAIAppJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikeAIAppJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikeAIAppJobResponseBody) *SubmitYikeAIAppJobResponse
	GetBody() *SubmitYikeAIAppJobResponseBody
}

type SubmitYikeAIAppJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikeAIAppJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikeAIAppJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeAIAppJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikeAIAppJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikeAIAppJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikeAIAppJobResponse) GetBody() *SubmitYikeAIAppJobResponseBody {
	return s.Body
}

func (s *SubmitYikeAIAppJobResponse) SetHeaders(v map[string]*string) *SubmitYikeAIAppJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikeAIAppJobResponse) SetStatusCode(v int32) *SubmitYikeAIAppJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikeAIAppJobResponse) SetBody(v *SubmitYikeAIAppJobResponseBody) *SubmitYikeAIAppJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikeAIAppJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
