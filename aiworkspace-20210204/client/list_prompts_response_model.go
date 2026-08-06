// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPromptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPromptsResponse
	GetStatusCode() *int32
	SetBody(v *ListPromptsResponseBody) *ListPromptsResponse
	GetBody() *ListPromptsResponseBody
}

type ListPromptsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPromptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPromptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsResponse) GoString() string {
	return s.String()
}

func (s *ListPromptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPromptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPromptsResponse) GetBody() *ListPromptsResponseBody {
	return s.Body
}

func (s *ListPromptsResponse) SetHeaders(v map[string]*string) *ListPromptsResponse {
	s.Headers = v
	return s
}

func (s *ListPromptsResponse) SetStatusCode(v int32) *ListPromptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromptsResponse) SetBody(v *ListPromptsResponseBody) *ListPromptsResponse {
	s.Body = v
	return s
}

func (s *ListPromptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
