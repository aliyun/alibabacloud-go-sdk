// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConnQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConnQuestionResponse
	GetStatusCode() *int32
	SetBody(v *ListConnQuestionResponseBody) *ListConnQuestionResponse
	GetBody() *ListConnQuestionResponseBody
}

type ListConnQuestionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *ListConnQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConnQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConnQuestionResponse) GetBody() *ListConnQuestionResponseBody {
	return s.Body
}

func (s *ListConnQuestionResponse) SetHeaders(v map[string]*string) *ListConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *ListConnQuestionResponse) SetStatusCode(v int32) *ListConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnQuestionResponse) SetBody(v *ListConnQuestionResponseBody) *ListConnQuestionResponse {
	s.Body = v
	return s
}

func (s *ListConnQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
