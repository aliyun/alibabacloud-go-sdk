// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAiCallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAiCallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAiCallTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartAiCallTaskResponseBody) *StartAiCallTaskResponse
	GetBody() *StartAiCallTaskResponseBody
}

type StartAiCallTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAiCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAiCallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAiCallTaskResponse) GoString() string {
	return s.String()
}

func (s *StartAiCallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAiCallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAiCallTaskResponse) GetBody() *StartAiCallTaskResponseBody {
	return s.Body
}

func (s *StartAiCallTaskResponse) SetHeaders(v map[string]*string) *StartAiCallTaskResponse {
	s.Headers = v
	return s
}

func (s *StartAiCallTaskResponse) SetStatusCode(v int32) *StartAiCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAiCallTaskResponse) SetBody(v *StartAiCallTaskResponseBody) *StartAiCallTaskResponse {
	s.Body = v
	return s
}

func (s *StartAiCallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
