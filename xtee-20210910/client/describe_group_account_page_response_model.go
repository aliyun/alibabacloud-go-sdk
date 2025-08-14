// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupAccountPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupAccountPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupAccountPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupAccountPageResponseBody) *DescribeGroupAccountPageResponse
	GetBody() *DescribeGroupAccountPageResponseBody
}

type DescribeGroupAccountPageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupAccountPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupAccountPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupAccountPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupAccountPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupAccountPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupAccountPageResponse) GetBody() *DescribeGroupAccountPageResponseBody {
	return s.Body
}

func (s *DescribeGroupAccountPageResponse) SetHeaders(v map[string]*string) *DescribeGroupAccountPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupAccountPageResponse) SetStatusCode(v int32) *DescribeGroupAccountPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupAccountPageResponse) SetBody(v *DescribeGroupAccountPageResponseBody) *DescribeGroupAccountPageResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupAccountPageResponse) Validate() error {
	return dara.Validate(s)
}
