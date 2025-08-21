// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSimQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSimQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSimQuestionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSimQuestionResponseBody) *UpdateSimQuestionResponse
	GetBody() *UpdateSimQuestionResponseBody
}

type UpdateSimQuestionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSimQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSimQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSimQuestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSimQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSimQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSimQuestionResponse) GetBody() *UpdateSimQuestionResponseBody {
	return s.Body
}

func (s *UpdateSimQuestionResponse) SetHeaders(v map[string]*string) *UpdateSimQuestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSimQuestionResponse) SetStatusCode(v int32) *UpdateSimQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSimQuestionResponse) SetBody(v *UpdateSimQuestionResponseBody) *UpdateSimQuestionResponse {
	s.Body = v
	return s
}

func (s *UpdateSimQuestionResponse) Validate() error {
	return dara.Validate(s)
}
