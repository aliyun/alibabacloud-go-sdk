// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourcesResponseBody) *DescribeResourcesResponse
	GetBody() *DescribeResourcesResponseBody
}

type DescribeResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourcesResponse) GetBody() *DescribeResourcesResponseBody {
	return s.Body
}

func (s *DescribeResourcesResponse) SetHeaders(v map[string]*string) *DescribeResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcesResponse) SetStatusCode(v int32) *DescribeResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcesResponse) SetBody(v *DescribeResourcesResponseBody) *DescribeResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
