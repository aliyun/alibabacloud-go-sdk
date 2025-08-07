// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDataNetworkListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalDataNetworkListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalDataNetworkListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalDataNetworkListResponseBody) *DescribeGlobalDataNetworkListResponse
	GetBody() *DescribeGlobalDataNetworkListResponseBody
}

type DescribeGlobalDataNetworkListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDataNetworkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDataNetworkListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalDataNetworkListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalDataNetworkListResponse) GetBody() *DescribeGlobalDataNetworkListResponseBody {
	return s.Body
}

func (s *DescribeGlobalDataNetworkListResponse) SetHeaders(v map[string]*string) *DescribeGlobalDataNetworkListResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponse) SetStatusCode(v int32) *DescribeGlobalDataNetworkListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponse) SetBody(v *DescribeGlobalDataNetworkListResponseBody) *DescribeGlobalDataNetworkListResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponse) Validate() error {
	return dara.Validate(s)
}
