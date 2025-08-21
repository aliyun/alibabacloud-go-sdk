// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectableClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConnectableClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConnectableClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConnectableClustersResponseBody) *DescribeConnectableClustersResponse
	GetBody() *DescribeConnectableClustersResponseBody
}

type DescribeConnectableClustersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConnectableClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConnectableClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectableClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectableClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConnectableClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConnectableClustersResponse) GetBody() *DescribeConnectableClustersResponseBody {
	return s.Body
}

func (s *DescribeConnectableClustersResponse) SetHeaders(v map[string]*string) *DescribeConnectableClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectableClustersResponse) SetStatusCode(v int32) *DescribeConnectableClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConnectableClustersResponse) SetBody(v *DescribeConnectableClustersResponseBody) *DescribeConnectableClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeConnectableClustersResponse) Validate() error {
	return dara.Validate(s)
}
