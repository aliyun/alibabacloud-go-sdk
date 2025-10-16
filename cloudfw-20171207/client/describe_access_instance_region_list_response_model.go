// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceRegionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessInstanceRegionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessInstanceRegionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessInstanceRegionListResponseBody) *DescribeAccessInstanceRegionListResponse
	GetBody() *DescribeAccessInstanceRegionListResponseBody
}

type DescribeAccessInstanceRegionListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessInstanceRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessInstanceRegionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceRegionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceRegionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessInstanceRegionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessInstanceRegionListResponse) GetBody() *DescribeAccessInstanceRegionListResponseBody {
	return s.Body
}

func (s *DescribeAccessInstanceRegionListResponse) SetHeaders(v map[string]*string) *DescribeAccessInstanceRegionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessInstanceRegionListResponse) SetStatusCode(v int32) *DescribeAccessInstanceRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessInstanceRegionListResponse) SetBody(v *DescribeAccessInstanceRegionListResponseBody) *DescribeAccessInstanceRegionListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessInstanceRegionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
