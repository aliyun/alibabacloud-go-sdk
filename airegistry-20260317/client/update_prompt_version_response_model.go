// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePromptVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePromptVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePromptVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePromptVersionResponseBody) *UpdatePromptVersionResponse
	GetBody() *UpdatePromptVersionResponseBody
}

type UpdatePromptVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePromptVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePromptVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePromptVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePromptVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePromptVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePromptVersionResponse) GetBody() *UpdatePromptVersionResponseBody {
	return s.Body
}

func (s *UpdatePromptVersionResponse) SetHeaders(v map[string]*string) *UpdatePromptVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePromptVersionResponse) SetStatusCode(v int32) *UpdatePromptVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePromptVersionResponse) SetBody(v *UpdatePromptVersionResponseBody) *UpdatePromptVersionResponse {
	s.Body = v
	return s
}

func (s *UpdatePromptVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
