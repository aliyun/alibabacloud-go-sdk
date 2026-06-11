// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePromptVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePromptVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePromptVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreatePromptVersionResponseBody) *CreatePromptVersionResponse
	GetBody() *CreatePromptVersionResponseBody
}

type CreatePromptVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePromptVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePromptVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePromptVersionResponse) GoString() string {
	return s.String()
}

func (s *CreatePromptVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePromptVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePromptVersionResponse) GetBody() *CreatePromptVersionResponseBody {
	return s.Body
}

func (s *CreatePromptVersionResponse) SetHeaders(v map[string]*string) *CreatePromptVersionResponse {
	s.Headers = v
	return s
}

func (s *CreatePromptVersionResponse) SetStatusCode(v int32) *CreatePromptVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePromptVersionResponse) SetBody(v *CreatePromptVersionResponseBody) *CreatePromptVersionResponse {
	s.Body = v
	return s
}

func (s *CreatePromptVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
