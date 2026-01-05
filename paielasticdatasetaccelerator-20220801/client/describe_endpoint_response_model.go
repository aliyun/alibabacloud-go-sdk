// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEndpointResponseBody) *DescribeEndpointResponse
	GetBody() *DescribeEndpointResponseBody
}

type DescribeEndpointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEndpointResponse) GetBody() *DescribeEndpointResponseBody {
	return s.Body
}

func (s *DescribeEndpointResponse) SetHeaders(v map[string]*string) *DescribeEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointResponse) SetStatusCode(v int32) *DescribeEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointResponse) SetBody(v *DescribeEndpointResponseBody) *DescribeEndpointResponse {
	s.Body = v
	return s
}

func (s *DescribeEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
