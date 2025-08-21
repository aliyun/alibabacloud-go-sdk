// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConnQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConnQuestionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConnQuestionResponseBody) *DeleteConnQuestionResponse
	GetBody() *DeleteConnQuestionResponseBody
}

type DeleteConnQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConnQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConnQuestionResponse) GetBody() *DeleteConnQuestionResponseBody {
	return s.Body
}

func (s *DeleteConnQuestionResponse) SetHeaders(v map[string]*string) *DeleteConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnQuestionResponse) SetStatusCode(v int32) *DeleteConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnQuestionResponse) SetBody(v *DeleteConnQuestionResponseBody) *DeleteConnQuestionResponse {
	s.Body = v
	return s
}

func (s *DeleteConnQuestionResponse) Validate() error {
	return dara.Validate(s)
}
