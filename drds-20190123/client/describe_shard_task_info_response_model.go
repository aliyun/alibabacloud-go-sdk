// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShardTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeShardTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeShardTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeShardTaskInfoResponseBody) *DescribeShardTaskInfoResponse
	GetBody() *DescribeShardTaskInfoResponseBody
}

type DescribeShardTaskInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeShardTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeShardTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeShardTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeShardTaskInfoResponse) GetBody() *DescribeShardTaskInfoResponseBody {
	return s.Body
}

func (s *DescribeShardTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeShardTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardTaskInfoResponse) SetStatusCode(v int32) *DescribeShardTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShardTaskInfoResponse) SetBody(v *DescribeShardTaskInfoResponseBody) *DescribeShardTaskInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeShardTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
