// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAICoachDebugResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAICoachDebugResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAICoachDebugResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAICoachDebugResponseBody) *SubmitAICoachDebugResponse
	GetBody() *SubmitAICoachDebugResponseBody
}

type SubmitAICoachDebugResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAICoachDebugResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAICoachDebugResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICoachDebugResponse) GoString() string {
	return s.String()
}

func (s *SubmitAICoachDebugResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAICoachDebugResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAICoachDebugResponse) GetBody() *SubmitAICoachDebugResponseBody {
	return s.Body
}

func (s *SubmitAICoachDebugResponse) SetHeaders(v map[string]*string) *SubmitAICoachDebugResponse {
	s.Headers = v
	return s
}

func (s *SubmitAICoachDebugResponse) SetStatusCode(v int32) *SubmitAICoachDebugResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAICoachDebugResponse) SetBody(v *SubmitAICoachDebugResponseBody) *SubmitAICoachDebugResponse {
	s.Body = v
	return s
}

func (s *SubmitAICoachDebugResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
