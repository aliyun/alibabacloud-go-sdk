// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeContainerAppStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeContainerAppStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeContainerAppStatsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeContainerAppStatsResponseBody) *DescribeEdgeContainerAppStatsResponse
	GetBody() *DescribeEdgeContainerAppStatsResponseBody
}

type DescribeEdgeContainerAppStatsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeContainerAppStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeContainerAppStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeContainerAppStatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeContainerAppStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeContainerAppStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeContainerAppStatsResponse) GetBody() *DescribeEdgeContainerAppStatsResponseBody {
	return s.Body
}

func (s *DescribeEdgeContainerAppStatsResponse) SetHeaders(v map[string]*string) *DescribeEdgeContainerAppStatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponse) SetStatusCode(v int32) *DescribeEdgeContainerAppStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponse) SetBody(v *DescribeEdgeContainerAppStatsResponseBody) *DescribeEdgeContainerAppStatsResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
