// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConnQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConnQuestionResponse
	GetStatusCode() *int32
	SetBody(v *CreateConnQuestionResponseBody) *CreateConnQuestionResponse
	GetBody() *CreateConnQuestionResponseBody
}

type CreateConnQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConnQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConnQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConnQuestionResponse) GoString() string {
	return s.String()
}

func (s *CreateConnQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConnQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConnQuestionResponse) GetBody() *CreateConnQuestionResponseBody {
	return s.Body
}

func (s *CreateConnQuestionResponse) SetHeaders(v map[string]*string) *CreateConnQuestionResponse {
	s.Headers = v
	return s
}

func (s *CreateConnQuestionResponse) SetStatusCode(v int32) *CreateConnQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConnQuestionResponse) SetBody(v *CreateConnQuestionResponseBody) *CreateConnQuestionResponse {
	s.Body = v
	return s
}

func (s *CreateConnQuestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
