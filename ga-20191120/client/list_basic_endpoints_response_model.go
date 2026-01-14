// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBasicEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBasicEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListBasicEndpointsResponseBody) *ListBasicEndpointsResponse
	GetBody() *ListBasicEndpointsResponseBody
}

type ListBasicEndpointsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBasicEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBasicEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBasicEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBasicEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBasicEndpointsResponse) GetBody() *ListBasicEndpointsResponseBody {
	return s.Body
}

func (s *ListBasicEndpointsResponse) SetHeaders(v map[string]*string) *ListBasicEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicEndpointsResponse) SetStatusCode(v int32) *ListBasicEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicEndpointsResponse) SetBody(v *ListBasicEndpointsResponseBody) *ListBasicEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListBasicEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
