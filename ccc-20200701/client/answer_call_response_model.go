// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnswerCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnswerCallResponse
	GetStatusCode() *int32
	SetBody(v *AnswerCallResponseBody) *AnswerCallResponse
	GetBody() *AnswerCallResponseBody
}

type AnswerCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnswerCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnswerCallResponse) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallResponse) GoString() string {
	return s.String()
}

func (s *AnswerCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnswerCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnswerCallResponse) GetBody() *AnswerCallResponseBody {
	return s.Body
}

func (s *AnswerCallResponse) SetHeaders(v map[string]*string) *AnswerCallResponse {
	s.Headers = v
	return s
}

func (s *AnswerCallResponse) SetStatusCode(v int32) *AnswerCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AnswerCallResponse) SetBody(v *AnswerCallResponseBody) *AnswerCallResponse {
	s.Body = v
	return s
}

func (s *AnswerCallResponse) Validate() error {
	return dara.Validate(s)
}
