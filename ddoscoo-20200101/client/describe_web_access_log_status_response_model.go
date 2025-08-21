// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebAccessLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebAccessLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebAccessLogStatusResponseBody) *DescribeWebAccessLogStatusResponse
	GetBody() *DescribeWebAccessLogStatusResponseBody
}

type DescribeWebAccessLogStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebAccessLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebAccessLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebAccessLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebAccessLogStatusResponse) GetBody() *DescribeWebAccessLogStatusResponseBody {
	return s.Body
}

func (s *DescribeWebAccessLogStatusResponse) SetHeaders(v map[string]*string) *DescribeWebAccessLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAccessLogStatusResponse) SetStatusCode(v int32) *DescribeWebAccessLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponse) SetBody(v *DescribeWebAccessLogStatusResponseBody) *DescribeWebAccessLogStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeWebAccessLogStatusResponse) Validate() error {
	return dara.Validate(s)
}
