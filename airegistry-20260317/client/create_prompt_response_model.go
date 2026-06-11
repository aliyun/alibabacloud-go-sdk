// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePromptResponse
	GetStatusCode() *int32
	SetBody(v *CreatePromptResponseBody) *CreatePromptResponse
	GetBody() *CreatePromptResponseBody
}

type CreatePromptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePromptResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptResponse) GoString() string {
	return s.String()
}

func (s *CreatePromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePromptResponse) GetBody() *CreatePromptResponseBody {
	return s.Body
}

func (s *CreatePromptResponse) SetHeaders(v map[string]*string) *CreatePromptResponse {
	s.Headers = v
	return s
}

func (s *CreatePromptResponse) SetStatusCode(v int32) *CreatePromptResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePromptResponse) SetBody(v *CreatePromptResponseBody) *CreatePromptResponse {
	s.Body = v
	return s
}

func (s *CreatePromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
