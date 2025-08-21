// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlackholeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBlackholeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBlackholeStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBlackholeStatusResponseBody) *DescribeBlackholeStatusResponse
	GetBody() *DescribeBlackholeStatusResponseBody
}

type DescribeBlackholeStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBlackholeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBlackholeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackholeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlackholeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBlackholeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBlackholeStatusResponse) GetBody() *DescribeBlackholeStatusResponseBody {
	return s.Body
}

func (s *DescribeBlackholeStatusResponse) SetHeaders(v map[string]*string) *DescribeBlackholeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlackholeStatusResponse) SetStatusCode(v int32) *DescribeBlackholeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlackholeStatusResponse) SetBody(v *DescribeBlackholeStatusResponseBody) *DescribeBlackholeStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeBlackholeStatusResponse) Validate() error {
	return dara.Validate(s)
}
