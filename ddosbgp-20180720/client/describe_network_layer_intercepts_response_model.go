// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkLayerInterceptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkLayerInterceptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkLayerInterceptsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkLayerInterceptsResponseBody) *DescribeNetworkLayerInterceptsResponse
	GetBody() *DescribeNetworkLayerInterceptsResponseBody
}

type DescribeNetworkLayerInterceptsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkLayerInterceptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkLayerInterceptsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkLayerInterceptsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkLayerInterceptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkLayerInterceptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkLayerInterceptsResponse) GetBody() *DescribeNetworkLayerInterceptsResponseBody {
	return s.Body
}

func (s *DescribeNetworkLayerInterceptsResponse) SetHeaders(v map[string]*string) *DescribeNetworkLayerInterceptsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponse) SetStatusCode(v int32) *DescribeNetworkLayerInterceptsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponse) SetBody(v *DescribeNetworkLayerInterceptsResponseBody) *DescribeNetworkLayerInterceptsResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
