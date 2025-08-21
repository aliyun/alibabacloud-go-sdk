// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnBlackholeCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUnBlackholeCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUnBlackholeCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUnBlackholeCountResponseBody) *DescribeUnBlackholeCountResponse
	GetBody() *DescribeUnBlackholeCountResponseBody
}

type DescribeUnBlackholeCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUnBlackholeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUnBlackholeCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnBlackholeCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnBlackholeCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUnBlackholeCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUnBlackholeCountResponse) GetBody() *DescribeUnBlackholeCountResponseBody {
	return s.Body
}

func (s *DescribeUnBlackholeCountResponse) SetHeaders(v map[string]*string) *DescribeUnBlackholeCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnBlackholeCountResponse) SetStatusCode(v int32) *DescribeUnBlackholeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUnBlackholeCountResponse) SetBody(v *DescribeUnBlackholeCountResponseBody) *DescribeUnBlackholeCountResponse {
	s.Body = v
	return s
}

func (s *DescribeUnBlackholeCountResponse) Validate() error {
	return dara.Validate(s)
}
