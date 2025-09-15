// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseGroupValidResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseGroupValidResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseGroupValidResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseGroupValidResourcesResponseBody) *DescribeDefenseGroupValidResourcesResponse
	GetBody() *DescribeDefenseGroupValidResourcesResponseBody
}

type DescribeDefenseGroupValidResourcesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseGroupValidResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseGroupValidResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseGroupValidResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseGroupValidResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseGroupValidResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseGroupValidResourcesResponse) GetBody() *DescribeDefenseGroupValidResourcesResponseBody {
	return s.Body
}

func (s *DescribeDefenseGroupValidResourcesResponse) SetHeaders(v map[string]*string) *DescribeDefenseGroupValidResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseGroupValidResourcesResponse) SetStatusCode(v int32) *DescribeDefenseGroupValidResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesResponse) SetBody(v *DescribeDefenseGroupValidResourcesResponseBody) *DescribeDefenseGroupValidResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseGroupValidResourcesResponse) Validate() error {
	return dara.Validate(s)
}
