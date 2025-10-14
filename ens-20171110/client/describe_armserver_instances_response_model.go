// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeARMServerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeARMServerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeARMServerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeARMServerInstancesResponseBody) *DescribeARMServerInstancesResponse
	GetBody() *DescribeARMServerInstancesResponseBody
}

type DescribeARMServerInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeARMServerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeARMServerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeARMServerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeARMServerInstancesResponse) GetBody() *DescribeARMServerInstancesResponseBody {
	return s.Body
}

func (s *DescribeARMServerInstancesResponse) SetHeaders(v map[string]*string) *DescribeARMServerInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeARMServerInstancesResponse) SetStatusCode(v int32) *DescribeARMServerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeARMServerInstancesResponse) SetBody(v *DescribeARMServerInstancesResponseBody) *DescribeARMServerInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeARMServerInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
