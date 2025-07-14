// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSlbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationSlbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationSlbsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationSlbsResponseBody) *DescribeApplicationSlbsResponse
	GetBody() *DescribeApplicationSlbsResponseBody
}

type DescribeApplicationSlbsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationSlbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationSlbsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSlbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationSlbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationSlbsResponse) GetBody() *DescribeApplicationSlbsResponseBody {
	return s.Body
}

func (s *DescribeApplicationSlbsResponse) SetHeaders(v map[string]*string) *DescribeApplicationSlbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationSlbsResponse) SetStatusCode(v int32) *DescribeApplicationSlbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationSlbsResponse) SetBody(v *DescribeApplicationSlbsResponseBody) *DescribeApplicationSlbsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationSlbsResponse) Validate() error {
	return dara.Validate(s)
}
