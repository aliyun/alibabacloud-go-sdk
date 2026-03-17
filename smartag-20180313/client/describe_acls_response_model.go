// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeACLsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeACLsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeACLsResponseBody) *DescribeACLsResponse
	GetBody() *DescribeACLsResponseBody
}

type DescribeACLsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeACLsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeACLsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLsResponse) GoString() string {
	return s.String()
}

func (s *DescribeACLsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeACLsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeACLsResponse) GetBody() *DescribeACLsResponseBody {
	return s.Body
}

func (s *DescribeACLsResponse) SetHeaders(v map[string]*string) *DescribeACLsResponse {
	s.Headers = v
	return s
}

func (s *DescribeACLsResponse) SetStatusCode(v int32) *DescribeACLsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeACLsResponse) SetBody(v *DescribeACLsResponseBody) *DescribeACLsResponse {
	s.Body = v
	return s
}

func (s *DescribeACLsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
