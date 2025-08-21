// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSimQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSimQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSimQuestionResponse
	GetStatusCode() *int32
	SetBody(v *ListSimQuestionResponseBody) *ListSimQuestionResponse
	GetBody() *ListSimQuestionResponseBody
}

type ListSimQuestionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSimQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *ListSimQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSimQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSimQuestionResponse) GetBody() *ListSimQuestionResponseBody {
	return s.Body
}

func (s *ListSimQuestionResponse) SetHeaders(v map[string]*string) *ListSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *ListSimQuestionResponse) SetStatusCode(v int32) *ListSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSimQuestionResponse) SetBody(v *ListSimQuestionResponseBody) *ListSimQuestionResponse {
	s.Body = v
	return s
}

func (s *ListSimQuestionResponse) Validate() error {
	return dara.Validate(s)
}
