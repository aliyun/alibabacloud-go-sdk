// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeCloudRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeCloudRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeCloudRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeCloudRegionResponseBody) *DescribeEdgeCloudRegionResponse
	GetBody() *DescribeEdgeCloudRegionResponseBody
}

type DescribeEdgeCloudRegionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeCloudRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeCloudRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeCloudRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeCloudRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeCloudRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeCloudRegionResponse) GetBody() *DescribeEdgeCloudRegionResponseBody {
	return s.Body
}

func (s *DescribeEdgeCloudRegionResponse) SetHeaders(v map[string]*string) *DescribeEdgeCloudRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeCloudRegionResponse) SetStatusCode(v int32) *DescribeEdgeCloudRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeCloudRegionResponse) SetBody(v *DescribeEdgeCloudRegionResponseBody) *DescribeEdgeCloudRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeCloudRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
