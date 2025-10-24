// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudServerRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudServerRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudServerRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudServerRegionsResponseBody) *DescribeHybridCloudServerRegionsResponse
	GetBody() *DescribeHybridCloudServerRegionsResponseBody
}

type DescribeHybridCloudServerRegionsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudServerRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudServerRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudServerRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudServerRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudServerRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudServerRegionsResponse) GetBody() *DescribeHybridCloudServerRegionsResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudServerRegionsResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudServerRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponse) SetStatusCode(v int32) *DescribeHybridCloudServerRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponse) SetBody(v *DescribeHybridCloudServerRegionsResponseBody) *DescribeHybridCloudServerRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudServerRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
