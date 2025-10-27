// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterShardNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterShardNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterShardNumberResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterShardNumberResponseBody) *DescribeDBClusterShardNumberResponse
	GetBody() *DescribeDBClusterShardNumberResponseBody
}

type DescribeDBClusterShardNumberResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterShardNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterShardNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterShardNumberResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterShardNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterShardNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterShardNumberResponse) GetBody() *DescribeDBClusterShardNumberResponseBody {
	return s.Body
}

func (s *DescribeDBClusterShardNumberResponse) SetHeaders(v map[string]*string) *DescribeDBClusterShardNumberResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterShardNumberResponse) SetStatusCode(v int32) *DescribeDBClusterShardNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterShardNumberResponse) SetBody(v *DescribeDBClusterShardNumberResponseBody) *DescribeDBClusterShardNumberResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterShardNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
