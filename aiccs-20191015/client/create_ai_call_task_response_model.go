// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiCallTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiCallTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiCallTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiCallTaskResponseBody) *CreateAiCallTaskResponse
	GetBody() *CreateAiCallTaskResponseBody
}

type CreateAiCallTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiCallTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiCallTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAiCallTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiCallTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiCallTaskResponse) GetBody() *CreateAiCallTaskResponseBody {
	return s.Body
}

func (s *CreateAiCallTaskResponse) SetHeaders(v map[string]*string) *CreateAiCallTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAiCallTaskResponse) SetStatusCode(v int32) *CreateAiCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiCallTaskResponse) SetBody(v *CreateAiCallTaskResponseBody) *CreateAiCallTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAiCallTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
