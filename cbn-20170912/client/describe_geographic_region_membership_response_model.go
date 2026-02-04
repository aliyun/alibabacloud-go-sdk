// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGeographicRegionMembershipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGeographicRegionMembershipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGeographicRegionMembershipResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGeographicRegionMembershipResponseBody) *DescribeGeographicRegionMembershipResponse
	GetBody() *DescribeGeographicRegionMembershipResponseBody
}

type DescribeGeographicRegionMembershipResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGeographicRegionMembershipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponse) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGeographicRegionMembershipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGeographicRegionMembershipResponse) GetBody() *DescribeGeographicRegionMembershipResponseBody {
	return s.Body
}

func (s *DescribeGeographicRegionMembershipResponse) SetHeaders(v map[string]*string) *DescribeGeographicRegionMembershipResponse {
	s.Headers = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponse) SetStatusCode(v int32) *DescribeGeographicRegionMembershipResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponse) SetBody(v *DescribeGeographicRegionMembershipResponseBody) *DescribeGeographicRegionMembershipResponse {
	s.Body = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
