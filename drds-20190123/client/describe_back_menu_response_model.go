// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackMenuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackMenuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackMenuResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackMenuResponseBody) *DescribeBackMenuResponse
	GetBody() *DescribeBackMenuResponseBody
}

type DescribeBackMenuResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackMenuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackMenuResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackMenuResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackMenuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackMenuResponse) GetBody() *DescribeBackMenuResponseBody {
	return s.Body
}

func (s *DescribeBackMenuResponse) SetHeaders(v map[string]*string) *DescribeBackMenuResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackMenuResponse) SetStatusCode(v int32) *DescribeBackMenuResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackMenuResponse) SetBody(v *DescribeBackMenuResponseBody) *DescribeBackMenuResponse {
	s.Body = v
	return s
}

func (s *DescribeBackMenuResponse) Validate() error {
	return dara.Validate(s)
}
