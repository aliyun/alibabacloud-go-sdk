// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingInstancesResponseBody) *ListRenderingInstancesResponse
	GetBody() *ListRenderingInstancesResponseBody
}

type ListRenderingInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingInstancesResponse) GetBody() *ListRenderingInstancesResponseBody {
	return s.Body
}

func (s *ListRenderingInstancesResponse) SetHeaders(v map[string]*string) *ListRenderingInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingInstancesResponse) SetStatusCode(v int32) *ListRenderingInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingInstancesResponse) SetBody(v *ListRenderingInstancesResponseBody) *ListRenderingInstancesResponse {
	s.Body = v
	return s
}

func (s *ListRenderingInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
