// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimQuestionResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimQuestionResponseBody) *CreateSimQuestionResponse
	GetBody() *CreateSimQuestionResponseBody
}

type CreateSimQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *CreateSimQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimQuestionResponse) GetBody() *CreateSimQuestionResponseBody {
	return s.Body
}

func (s *CreateSimQuestionResponse) SetHeaders(v map[string]*string) *CreateSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *CreateSimQuestionResponse) SetStatusCode(v int32) *CreateSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimQuestionResponse) SetBody(v *CreateSimQuestionResponseBody) *CreateSimQuestionResponse {
	s.Body = v
	return s
}

func (s *CreateSimQuestionResponse) Validate() error {
	return dara.Validate(s)
}
