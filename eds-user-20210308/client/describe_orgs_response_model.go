// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOrgsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOrgsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOrgsResponseBody) *DescribeOrgsResponse
	GetBody() *DescribeOrgsResponseBody
}

type DescribeOrgsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOrgsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOrgsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrgsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOrgsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOrgsResponse) GetBody() *DescribeOrgsResponseBody {
	return s.Body
}

func (s *DescribeOrgsResponse) SetHeaders(v map[string]*string) *DescribeOrgsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrgsResponse) SetStatusCode(v int32) *DescribeOrgsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrgsResponse) SetBody(v *DescribeOrgsResponseBody) *DescribeOrgsResponse {
	s.Body = v
	return s
}

func (s *DescribeOrgsResponse) Validate() error {
	return dara.Validate(s)
}
