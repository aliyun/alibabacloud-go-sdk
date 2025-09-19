// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridProxyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridProxyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridProxyListResponseBody) *DescribeHybridProxyListResponse
	GetBody() *DescribeHybridProxyListResponseBody
}

type DescribeHybridProxyListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridProxyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridProxyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridProxyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridProxyListResponse) GetBody() *DescribeHybridProxyListResponseBody {
	return s.Body
}

func (s *DescribeHybridProxyListResponse) SetHeaders(v map[string]*string) *DescribeHybridProxyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridProxyListResponse) SetStatusCode(v int32) *DescribeHybridProxyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridProxyListResponse) SetBody(v *DescribeHybridProxyListResponseBody) *DescribeHybridProxyListResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridProxyListResponse) Validate() error {
	return dara.Validate(s)
}
