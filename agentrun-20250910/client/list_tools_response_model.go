// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListToolsResponse
	GetStatusCode() *int32
	SetBody(v *ListToolsResult) *ListToolsResponse
	GetBody() *ListToolsResult
}

type ListToolsResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListToolsResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListToolsResponse) GoString() string {
	return s.String()
}

func (s *ListToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListToolsResponse) GetBody() *ListToolsResult {
	return s.Body
}

func (s *ListToolsResponse) SetHeaders(v map[string]*string) *ListToolsResponse {
	s.Headers = v
	return s
}

func (s *ListToolsResponse) SetStatusCode(v int32) *ListToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListToolsResponse) SetBody(v *ListToolsResult) *ListToolsResponse {
	s.Body = v
	return s
}

func (s *ListToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
