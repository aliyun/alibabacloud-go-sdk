// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConnQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConnQuestionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConnQuestionResponseBody) *UpdateConnQuestionResponse
	GetBody() *UpdateConnQuestionResponseBody
}

type UpdateConnQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConnQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConnQuestionResponse) GetBody() *UpdateConnQuestionResponseBody {
	return s.Body
}

func (s *UpdateConnQuestionResponse) SetHeaders(v map[string]*string) *UpdateConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnQuestionResponse) SetStatusCode(v int32) *UpdateConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnQuestionResponse) SetBody(v *UpdateConnQuestionResponseBody) *UpdateConnQuestionResponse {
	s.Body = v
	return s
}

func (s *UpdateConnQuestionResponse) Validate() error {
	return dara.Validate(s)
}
