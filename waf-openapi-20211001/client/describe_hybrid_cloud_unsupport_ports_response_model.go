// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUnsupportPortsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudUnsupportPortsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudUnsupportPortsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudUnsupportPortsResponseBody) *DescribeHybridCloudUnsupportPortsResponse
	GetBody() *DescribeHybridCloudUnsupportPortsResponseBody
}

type DescribeHybridCloudUnsupportPortsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudUnsupportPortsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudUnsupportPortsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnsupportPortsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnsupportPortsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudUnsupportPortsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudUnsupportPortsResponse) GetBody() *DescribeHybridCloudUnsupportPortsResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudUnsupportPortsResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudUnsupportPortsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsResponse) SetStatusCode(v int32) *DescribeHybridCloudUnsupportPortsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsResponse) SetBody(v *DescribeHybridCloudUnsupportPortsResponseBody) *DescribeHybridCloudUnsupportPortsResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
