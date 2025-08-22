// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogserviceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserLogserviceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserLogserviceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserLogserviceStatusResponseBody) *DescribeUserLogserviceStatusResponse
	GetBody() *DescribeUserLogserviceStatusResponseBody
}

type DescribeUserLogserviceStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserLogserviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserLogserviceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogserviceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLogserviceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserLogserviceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserLogserviceStatusResponse) GetBody() *DescribeUserLogserviceStatusResponseBody {
	return s.Body
}

func (s *DescribeUserLogserviceStatusResponse) SetHeaders(v map[string]*string) *DescribeUserLogserviceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLogserviceStatusResponse) SetStatusCode(v int32) *DescribeUserLogserviceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserLogserviceStatusResponse) SetBody(v *DescribeUserLogserviceStatusResponseBody) *DescribeUserLogserviceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserLogserviceStatusResponse) Validate() error {
	return dara.Validate(s)
}
