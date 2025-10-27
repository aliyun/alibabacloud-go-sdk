// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyClusterListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridProxyClusterListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridProxyClusterListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridProxyClusterListResponseBody) *DescribeHybridProxyClusterListResponse
	GetBody() *DescribeHybridProxyClusterListResponseBody
}

type DescribeHybridProxyClusterListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridProxyClusterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridProxyClusterListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyClusterListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyClusterListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridProxyClusterListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridProxyClusterListResponse) GetBody() *DescribeHybridProxyClusterListResponseBody {
	return s.Body
}

func (s *DescribeHybridProxyClusterListResponse) SetHeaders(v map[string]*string) *DescribeHybridProxyClusterListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridProxyClusterListResponse) SetStatusCode(v int32) *DescribeHybridProxyClusterListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridProxyClusterListResponse) SetBody(v *DescribeHybridProxyClusterListResponseBody) *DescribeHybridProxyClusterListResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridProxyClusterListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
