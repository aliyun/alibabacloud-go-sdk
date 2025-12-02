// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnswerLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnswerLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnswerLibResponse
	GetStatusCode() *int32
	SetBody(v *ListAnswerLibResponseBody) *ListAnswerLibResponse
	GetBody() *ListAnswerLibResponseBody
}

type ListAnswerLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnswerLibResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *ListAnswerLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnswerLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnswerLibResponse) GetBody() *ListAnswerLibResponseBody {
	return s.Body
}

func (s *ListAnswerLibResponse) SetHeaders(v map[string]*string) *ListAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *ListAnswerLibResponse) SetStatusCode(v int32) *ListAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnswerLibResponse) SetBody(v *ListAnswerLibResponseBody) *ListAnswerLibResponse {
	s.Body = v
	return s
}

func (s *ListAnswerLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
