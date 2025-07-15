// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGuestApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGuestApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGuestApplicationsResponseBody) *DescribeGuestApplicationsResponse
	GetBody() *DescribeGuestApplicationsResponseBody
}

type DescribeGuestApplicationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGuestApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGuestApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestApplicationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGuestApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGuestApplicationsResponse) GetBody() *DescribeGuestApplicationsResponseBody {
	return s.Body
}

func (s *DescribeGuestApplicationsResponse) SetHeaders(v map[string]*string) *DescribeGuestApplicationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestApplicationsResponse) SetStatusCode(v int32) *DescribeGuestApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestApplicationsResponse) SetBody(v *DescribeGuestApplicationsResponseBody) *DescribeGuestApplicationsResponse {
	s.Body = v
	return s
}

func (s *DescribeGuestApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
