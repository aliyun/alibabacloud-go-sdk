// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudClusterServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudClusterServersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudClusterServersResponseBody) *DescribeHybridCloudClusterServersResponse
	GetBody() *DescribeHybridCloudClusterServersResponseBody
}

type DescribeHybridCloudClusterServersResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudClusterServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudClusterServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudClusterServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudClusterServersResponse) GetBody() *DescribeHybridCloudClusterServersResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudClusterServersResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudClusterServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudClusterServersResponse) SetStatusCode(v int32) *DescribeHybridCloudClusterServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudClusterServersResponse) SetBody(v *DescribeHybridCloudClusterServersResponseBody) *DescribeHybridCloudClusterServersResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudClusterServersResponse) Validate() error {
	return dara.Validate(s)
}
