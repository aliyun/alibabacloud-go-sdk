// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationPromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationPromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationPromptResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationPromptResponseBody) *DeleteApplicationPromptResponse
	GetBody() *DeleteApplicationPromptResponseBody
}

type DeleteApplicationPromptResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationPromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationPromptResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationPromptResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationPromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationPromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationPromptResponse) GetBody() *DeleteApplicationPromptResponseBody {
	return s.Body
}

func (s *DeleteApplicationPromptResponse) SetHeaders(v map[string]*string) *DeleteApplicationPromptResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationPromptResponse) SetStatusCode(v int32) *DeleteApplicationPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationPromptResponse) SetBody(v *DeleteApplicationPromptResponseBody) *DeleteApplicationPromptResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationPromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
