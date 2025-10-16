// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceZoneListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessInstanceZoneListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessInstanceZoneListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessInstanceZoneListResponseBody) *DescribeAccessInstanceZoneListResponse
	GetBody() *DescribeAccessInstanceZoneListResponseBody
}

type DescribeAccessInstanceZoneListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessInstanceZoneListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessInstanceZoneListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceZoneListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceZoneListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessInstanceZoneListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessInstanceZoneListResponse) GetBody() *DescribeAccessInstanceZoneListResponseBody {
	return s.Body
}

func (s *DescribeAccessInstanceZoneListResponse) SetHeaders(v map[string]*string) *DescribeAccessInstanceZoneListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessInstanceZoneListResponse) SetStatusCode(v int32) *DescribeAccessInstanceZoneListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessInstanceZoneListResponse) SetBody(v *DescribeAccessInstanceZoneListResponseBody) *DescribeAccessInstanceZoneListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessInstanceZoneListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
