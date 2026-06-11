// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePromptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePromptResponse
	GetStatusCode() *int32
	SetBody(v *DeletePromptResponseBody) *DeletePromptResponse
	GetBody() *DeletePromptResponseBody
}

type DeletePromptResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePromptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePromptResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptResponse) GoString() string {
	return s.String()
}

func (s *DeletePromptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePromptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePromptResponse) GetBody() *DeletePromptResponseBody {
	return s.Body
}

func (s *DeletePromptResponse) SetHeaders(v map[string]*string) *DeletePromptResponse {
	s.Headers = v
	return s
}

func (s *DeletePromptResponse) SetStatusCode(v int32) *DeletePromptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePromptResponse) SetBody(v *DeletePromptResponseBody) *DeletePromptResponse {
	s.Body = v
	return s
}

func (s *DeletePromptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
