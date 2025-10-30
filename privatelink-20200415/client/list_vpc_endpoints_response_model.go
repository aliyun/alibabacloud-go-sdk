// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointsResponseBody) *ListVpcEndpointsResponse
	GetBody() *ListVpcEndpointsResponseBody
}

type ListVpcEndpointsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointsResponse) GetBody() *ListVpcEndpointsResponseBody {
	return s.Body
}

func (s *ListVpcEndpointsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointsResponse) SetStatusCode(v int32) *ListVpcEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointsResponse) SetBody(v *ListVpcEndpointsResponseBody) *ListVpcEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
