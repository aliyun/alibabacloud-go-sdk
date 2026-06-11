// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPromptVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPromptVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetPromptVersionResponseBody) *GetPromptVersionResponse
	GetBody() *GetPromptVersionResponseBody
}

type GetPromptVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPromptVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPromptVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPromptVersionResponse) GoString() string {
	return s.String()
}

func (s *GetPromptVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPromptVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPromptVersionResponse) GetBody() *GetPromptVersionResponseBody {
	return s.Body
}

func (s *GetPromptVersionResponse) SetHeaders(v map[string]*string) *GetPromptVersionResponse {
	s.Headers = v
	return s
}

func (s *GetPromptVersionResponse) SetStatusCode(v int32) *GetPromptVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPromptVersionResponse) SetBody(v *GetPromptVersionResponseBody) *GetPromptVersionResponse {
	s.Body = v
	return s
}

func (s *GetPromptVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
