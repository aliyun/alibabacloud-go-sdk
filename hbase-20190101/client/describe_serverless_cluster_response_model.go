// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerlessClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServerlessClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServerlessClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServerlessClusterResponseBody) *DescribeServerlessClusterResponse
	GetBody() *DescribeServerlessClusterResponseBody
}

type DescribeServerlessClusterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerlessClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerlessClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerlessClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerlessClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServerlessClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServerlessClusterResponse) GetBody() *DescribeServerlessClusterResponseBody {
	return s.Body
}

func (s *DescribeServerlessClusterResponse) SetHeaders(v map[string]*string) *DescribeServerlessClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerlessClusterResponse) SetStatusCode(v int32) *DescribeServerlessClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerlessClusterResponse) SetBody(v *DescribeServerlessClusterResponseBody) *DescribeServerlessClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeServerlessClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
