// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingSessionsResponseBody) *ListRenderingSessionsResponse
	GetBody() *ListRenderingSessionsResponseBody
}

type ListRenderingSessionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingSessionsResponse) GetBody() *ListRenderingSessionsResponseBody {
	return s.Body
}

func (s *ListRenderingSessionsResponse) SetHeaders(v map[string]*string) *ListRenderingSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingSessionsResponse) SetStatusCode(v int32) *ListRenderingSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingSessionsResponse) SetBody(v *ListRenderingSessionsResponseBody) *ListRenderingSessionsResponse {
	s.Body = v
	return s
}

func (s *ListRenderingSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
