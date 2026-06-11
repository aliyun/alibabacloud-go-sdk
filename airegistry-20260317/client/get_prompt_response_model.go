// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPromptResponse
	GetStatusCode() *int32
	SetBody(v *GetPromptResponseBody) *GetPromptResponse
	GetBody() *GetPromptResponseBody
}

type GetPromptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPromptResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPromptResponse) GoString() string {
	return s.String()
}

func (s *GetPromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPromptResponse) GetBody() *GetPromptResponseBody {
	return s.Body
}

func (s *GetPromptResponse) SetHeaders(v map[string]*string) *GetPromptResponse {
	s.Headers = v
	return s
}

func (s *GetPromptResponse) SetStatusCode(v int32) *GetPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPromptResponse) SetBody(v *GetPromptResponseBody) *GetPromptResponse {
	s.Body = v
	return s
}

func (s *GetPromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
