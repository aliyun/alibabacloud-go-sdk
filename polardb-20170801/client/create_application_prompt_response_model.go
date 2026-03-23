// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationPromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationPromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationPromptResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationPromptResponseBody) *CreateApplicationPromptResponse
	GetBody() *CreateApplicationPromptResponseBody
}

type CreateApplicationPromptResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationPromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationPromptResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationPromptResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationPromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationPromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationPromptResponse) GetBody() *CreateApplicationPromptResponseBody {
	return s.Body
}

func (s *CreateApplicationPromptResponse) SetHeaders(v map[string]*string) *CreateApplicationPromptResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationPromptResponse) SetStatusCode(v int32) *CreateApplicationPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationPromptResponse) SetBody(v *CreateApplicationPromptResponseBody) *CreateApplicationPromptResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationPromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
