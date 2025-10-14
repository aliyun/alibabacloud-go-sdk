// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSDGsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSDGsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSDGsResponseBody) *DescribeSDGsResponse
	GetBody() *DescribeSDGsResponseBody
}

type DescribeSDGsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSDGsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSDGsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDGsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSDGsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSDGsResponse) GetBody() *DescribeSDGsResponseBody {
	return s.Body
}

func (s *DescribeSDGsResponse) SetHeaders(v map[string]*string) *DescribeSDGsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDGsResponse) SetStatusCode(v int32) *DescribeSDGsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSDGsResponse) SetBody(v *DescribeSDGsResponseBody) *DescribeSDGsResponse {
	s.Body = v
	return s
}

func (s *DescribeSDGsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
