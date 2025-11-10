// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiMcpServerSystemToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiMcpServerSystemToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiMcpServerSystemToolsResponse
	GetStatusCode() *int32
	SetBody(v *ListApiMcpServerSystemToolsResponseBody) *ListApiMcpServerSystemToolsResponse
	GetBody() *ListApiMcpServerSystemToolsResponseBody
}

type ListApiMcpServerSystemToolsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiMcpServerSystemToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiMcpServerSystemToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServerSystemToolsResponse) GoString() string {
	return s.String()
}

func (s *ListApiMcpServerSystemToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiMcpServerSystemToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiMcpServerSystemToolsResponse) GetBody() *ListApiMcpServerSystemToolsResponseBody {
	return s.Body
}

func (s *ListApiMcpServerSystemToolsResponse) SetHeaders(v map[string]*string) *ListApiMcpServerSystemToolsResponse {
	s.Headers = v
	return s
}

func (s *ListApiMcpServerSystemToolsResponse) SetStatusCode(v int32) *ListApiMcpServerSystemToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiMcpServerSystemToolsResponse) SetBody(v *ListApiMcpServerSystemToolsResponseBody) *ListApiMcpServerSystemToolsResponse {
	s.Body = v
	return s
}

func (s *ListApiMcpServerSystemToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
