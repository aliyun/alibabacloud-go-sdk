// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPromptVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPromptVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPromptVersionResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPromptVersionResponseBody) *SubmitPromptVersionResponse
	GetBody() *SubmitPromptVersionResponseBody
}

type SubmitPromptVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPromptVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPromptVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPromptVersionResponse) GoString() string {
	return s.String()
}

func (s *SubmitPromptVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPromptVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPromptVersionResponse) GetBody() *SubmitPromptVersionResponseBody {
	return s.Body
}

func (s *SubmitPromptVersionResponse) SetHeaders(v map[string]*string) *SubmitPromptVersionResponse {
	s.Headers = v
	return s
}

func (s *SubmitPromptVersionResponse) SetStatusCode(v int32) *SubmitPromptVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPromptVersionResponse) SetBody(v *SubmitPromptVersionResponseBody) *SubmitPromptVersionResponse {
	s.Body = v
	return s
}

func (s *SubmitPromptVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
