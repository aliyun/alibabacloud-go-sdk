// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceIdsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceIdsResponseBody) *DescribeInstanceIdsResponse
	GetBody() *DescribeInstanceIdsResponseBody
}

type DescribeInstanceIdsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceIdsResponse) GetBody() *DescribeInstanceIdsResponseBody {
	return s.Body
}

func (s *DescribeInstanceIdsResponse) SetHeaders(v map[string]*string) *DescribeInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceIdsResponse) SetStatusCode(v int32) *DescribeInstanceIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceIdsResponse) SetBody(v *DescribeInstanceIdsResponseBody) *DescribeInstanceIdsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
