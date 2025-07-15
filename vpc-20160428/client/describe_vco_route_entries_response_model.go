// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVcoRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVcoRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVcoRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVcoRouteEntriesResponseBody) *DescribeVcoRouteEntriesResponse
	GetBody() *DescribeVcoRouteEntriesResponseBody
}

type DescribeVcoRouteEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVcoRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVcoRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVcoRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVcoRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVcoRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVcoRouteEntriesResponse) GetBody() *DescribeVcoRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeVcoRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeVcoRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVcoRouteEntriesResponse) SetStatusCode(v int32) *DescribeVcoRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponse) SetBody(v *DescribeVcoRouteEntriesResponseBody) *DescribeVcoRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeVcoRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
