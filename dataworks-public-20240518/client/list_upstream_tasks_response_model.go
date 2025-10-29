// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUpstreamTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUpstreamTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListUpstreamTasksResponseBody) *ListUpstreamTasksResponse
	GetBody() *ListUpstreamTasksResponseBody
}

type ListUpstreamTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUpstreamTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUpstreamTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksResponse) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUpstreamTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUpstreamTasksResponse) GetBody() *ListUpstreamTasksResponseBody {
	return s.Body
}

func (s *ListUpstreamTasksResponse) SetHeaders(v map[string]*string) *ListUpstreamTasksResponse {
	s.Headers = v
	return s
}

func (s *ListUpstreamTasksResponse) SetStatusCode(v int32) *ListUpstreamTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUpstreamTasksResponse) SetBody(v *ListUpstreamTasksResponseBody) *ListUpstreamTasksResponse {
	s.Body = v
	return s
}

func (s *ListUpstreamTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
