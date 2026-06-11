// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePromptResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePromptResponseBody) *UpdatePromptResponse
	GetBody() *UpdatePromptResponseBody
}

type UpdatePromptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePromptResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptResponse) GoString() string {
	return s.String()
}

func (s *UpdatePromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePromptResponse) GetBody() *UpdatePromptResponseBody {
	return s.Body
}

func (s *UpdatePromptResponse) SetHeaders(v map[string]*string) *UpdatePromptResponse {
	s.Headers = v
	return s
}

func (s *UpdatePromptResponse) SetStatusCode(v int32) *UpdatePromptResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePromptResponse) SetBody(v *UpdatePromptResponseBody) *UpdatePromptResponse {
	s.Body = v
	return s
}

func (s *UpdatePromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
