// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteConflictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouteConflictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouteConflictResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouteConflictResponseBody) *DescribeRouteConflictResponse
	GetBody() *DescribeRouteConflictResponseBody
}

type DescribeRouteConflictResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouteConflictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouteConflictResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteConflictResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouteConflictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouteConflictResponse) GetBody() *DescribeRouteConflictResponseBody {
	return s.Body
}

func (s *DescribeRouteConflictResponse) SetHeaders(v map[string]*string) *DescribeRouteConflictResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteConflictResponse) SetStatusCode(v int32) *DescribeRouteConflictResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouteConflictResponse) SetBody(v *DescribeRouteConflictResponseBody) *DescribeRouteConflictResponse {
	s.Body = v
	return s
}

func (s *DescribeRouteConflictResponse) Validate() error {
	return dara.Validate(s)
}
