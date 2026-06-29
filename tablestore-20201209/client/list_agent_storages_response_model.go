// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStoragesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentStoragesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentStoragesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentStoragesResponseBody) *ListAgentStoragesResponse
	GetBody() *ListAgentStoragesResponseBody
}

type ListAgentStoragesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentStoragesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentStoragesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStoragesResponse) GoString() string {
	return s.String()
}

func (s *ListAgentStoragesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentStoragesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentStoragesResponse) GetBody() *ListAgentStoragesResponseBody {
	return s.Body
}

func (s *ListAgentStoragesResponse) SetHeaders(v map[string]*string) *ListAgentStoragesResponse {
	s.Headers = v
	return s
}

func (s *ListAgentStoragesResponse) SetStatusCode(v int32) *ListAgentStoragesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentStoragesResponse) SetBody(v *ListAgentStoragesResponseBody) *ListAgentStoragesResponse {
	s.Body = v
	return s
}

func (s *ListAgentStoragesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
