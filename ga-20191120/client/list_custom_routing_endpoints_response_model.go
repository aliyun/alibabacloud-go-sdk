// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomRoutingEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomRoutingEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomRoutingEndpointsResponseBody) *ListCustomRoutingEndpointsResponse
	GetBody() *ListCustomRoutingEndpointsResponseBody
}

type ListCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomRoutingEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomRoutingEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomRoutingEndpointsResponse) GetBody() *ListCustomRoutingEndpointsResponseBody {
	return s.Body
}

func (s *ListCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointsResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointsResponse) SetBody(v *ListCustomRoutingEndpointsResponseBody) *ListCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListCustomRoutingEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
