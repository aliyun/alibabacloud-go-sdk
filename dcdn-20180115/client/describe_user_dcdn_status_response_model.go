// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDcdnStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserDcdnStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserDcdnStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserDcdnStatusResponseBody) *DescribeUserDcdnStatusResponse
	GetBody() *DescribeUserDcdnStatusResponseBody
}

type DescribeUserDcdnStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserDcdnStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserDcdnStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDcdnStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDcdnStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserDcdnStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserDcdnStatusResponse) GetBody() *DescribeUserDcdnStatusResponseBody {
	return s.Body
}

func (s *DescribeUserDcdnStatusResponse) SetHeaders(v map[string]*string) *DescribeUserDcdnStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDcdnStatusResponse) SetStatusCode(v int32) *DescribeUserDcdnStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserDcdnStatusResponse) SetBody(v *DescribeUserDcdnStatusResponseBody) *DescribeUserDcdnStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserDcdnStatusResponse) Validate() error {
	return dara.Validate(s)
}
