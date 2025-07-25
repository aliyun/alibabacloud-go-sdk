// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortRangeListAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortRangeListAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortRangeListAssociationsResponseBody) *DescribePortRangeListAssociationsResponse
	GetBody() *DescribePortRangeListAssociationsResponseBody
}

type DescribePortRangeListAssociationsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortRangeListAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortRangeListAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListAssociationsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortRangeListAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortRangeListAssociationsResponse) GetBody() *DescribePortRangeListAssociationsResponseBody {
	return s.Body
}

func (s *DescribePortRangeListAssociationsResponse) SetHeaders(v map[string]*string) *DescribePortRangeListAssociationsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortRangeListAssociationsResponse) SetStatusCode(v int32) *DescribePortRangeListAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortRangeListAssociationsResponse) SetBody(v *DescribePortRangeListAssociationsResponseBody) *DescribePortRangeListAssociationsResponse {
	s.Body = v
	return s
}

func (s *DescribePortRangeListAssociationsResponse) Validate() error {
	return dara.Validate(s)
}
