// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPromptVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPromptVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListPromptVersionsResponseBody) *ListPromptVersionsResponse
	GetBody() *ListPromptVersionsResponseBody
}

type ListPromptVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPromptVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPromptVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPromptVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListPromptVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPromptVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPromptVersionsResponse) GetBody() *ListPromptVersionsResponseBody {
	return s.Body
}

func (s *ListPromptVersionsResponse) SetHeaders(v map[string]*string) *ListPromptVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListPromptVersionsResponse) SetStatusCode(v int32) *ListPromptVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromptVersionsResponse) SetBody(v *ListPromptVersionsResponseBody) *ListPromptVersionsResponse {
	s.Body = v
	return s
}

func (s *ListPromptVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
