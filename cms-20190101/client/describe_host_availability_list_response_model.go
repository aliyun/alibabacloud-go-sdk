// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostAvailabilityListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHostAvailabilityListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHostAvailabilityListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHostAvailabilityListResponseBody) *DescribeHostAvailabilityListResponse
	GetBody() *DescribeHostAvailabilityListResponseBody
}

type DescribeHostAvailabilityListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHostAvailabilityListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHostAvailabilityListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostAvailabilityListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHostAvailabilityListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHostAvailabilityListResponse) GetBody() *DescribeHostAvailabilityListResponseBody {
	return s.Body
}

func (s *DescribeHostAvailabilityListResponse) SetHeaders(v map[string]*string) *DescribeHostAvailabilityListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostAvailabilityListResponse) SetStatusCode(v int32) *DescribeHostAvailabilityListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHostAvailabilityListResponse) SetBody(v *DescribeHostAvailabilityListResponseBody) *DescribeHostAvailabilityListResponse {
	s.Body = v
	return s
}

func (s *DescribeHostAvailabilityListResponse) Validate() error {
	return dara.Validate(s)
}
