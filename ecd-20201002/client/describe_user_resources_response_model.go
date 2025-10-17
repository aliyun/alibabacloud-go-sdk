// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserResourcesResponseBody) *DescribeUserResourcesResponse
	GetBody() *DescribeUserResourcesResponseBody
}

type DescribeUserResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserResourcesResponse) GetBody() *DescribeUserResourcesResponseBody {
	return s.Body
}

func (s *DescribeUserResourcesResponse) SetHeaders(v map[string]*string) *DescribeUserResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserResourcesResponse) SetStatusCode(v int32) *DescribeUserResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserResourcesResponse) SetBody(v *DescribeUserResourcesResponseBody) *DescribeUserResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeUserResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
