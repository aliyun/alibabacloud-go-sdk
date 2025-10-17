// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAITaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenAITaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenAITaskResponse
	GetStatusCode() *int32
	SetBody(v *OpenAITaskResponseBody) *OpenAITaskResponse
	GetBody() *OpenAITaskResponseBody
}

type OpenAITaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenAITaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenAITaskResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenAITaskResponse) GoString() string {
	return s.String()
}

func (s *OpenAITaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenAITaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenAITaskResponse) GetBody() *OpenAITaskResponseBody {
	return s.Body
}

func (s *OpenAITaskResponse) SetHeaders(v map[string]*string) *OpenAITaskResponse {
	s.Headers = v
	return s
}

func (s *OpenAITaskResponse) SetStatusCode(v int32) *OpenAITaskResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAITaskResponse) SetBody(v *OpenAITaskResponseBody) *OpenAITaskResponse {
	s.Body = v
	return s
}

func (s *OpenAITaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
