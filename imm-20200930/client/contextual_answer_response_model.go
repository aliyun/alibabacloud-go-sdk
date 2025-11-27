// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualAnswerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContextualAnswerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContextualAnswerResponse
	GetStatusCode() *int32
	SetBody(v *ContextualAnswerResponseBody) *ContextualAnswerResponse
	GetBody() *ContextualAnswerResponseBody
}

type ContextualAnswerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContextualAnswerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContextualAnswerResponse) String() string {
	return dara.Prettify(s)
}

func (s ContextualAnswerResponse) GoString() string {
	return s.String()
}

func (s *ContextualAnswerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContextualAnswerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContextualAnswerResponse) GetBody() *ContextualAnswerResponseBody {
	return s.Body
}

func (s *ContextualAnswerResponse) SetHeaders(v map[string]*string) *ContextualAnswerResponse {
	s.Headers = v
	return s
}

func (s *ContextualAnswerResponse) SetStatusCode(v int32) *ContextualAnswerResponse {
	s.StatusCode = &v
	return s
}

func (s *ContextualAnswerResponse) SetBody(v *ContextualAnswerResponseBody) *ContextualAnswerResponse {
	s.Body = v
	return s
}

func (s *ContextualAnswerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
