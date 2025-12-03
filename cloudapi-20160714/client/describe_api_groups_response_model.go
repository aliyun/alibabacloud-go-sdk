// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiGroupsResponseBody) *DescribeApiGroupsResponse
	GetBody() *DescribeApiGroupsResponseBody
}

type DescribeApiGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiGroupsResponse) GetBody() *DescribeApiGroupsResponseBody {
	return s.Body
}

func (s *DescribeApiGroupsResponse) SetHeaders(v map[string]*string) *DescribeApiGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupsResponse) SetStatusCode(v int32) *DescribeApiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGroupsResponse) SetBody(v *DescribeApiGroupsResponseBody) *DescribeApiGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeApiGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
