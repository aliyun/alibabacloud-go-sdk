// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnswerLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnswerLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnswerLibResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnswerLibResponseBody) *CreateAnswerLibResponse
	GetBody() *CreateAnswerLibResponseBody
}

type CreateAnswerLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnswerLibResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnswerLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnswerLibResponse) GetBody() *CreateAnswerLibResponseBody {
	return s.Body
}

func (s *CreateAnswerLibResponse) SetHeaders(v map[string]*string) *CreateAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *CreateAnswerLibResponse) SetStatusCode(v int32) *CreateAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnswerLibResponse) SetBody(v *CreateAnswerLibResponseBody) *CreateAnswerLibResponse {
	s.Body = v
	return s
}

func (s *CreateAnswerLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
