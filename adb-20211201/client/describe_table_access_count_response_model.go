// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableAccessCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTableAccessCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTableAccessCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTableAccessCountResponseBody) *DescribeTableAccessCountResponse
	GetBody() *DescribeTableAccessCountResponseBody
}

type DescribeTableAccessCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableAccessCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableAccessCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableAccessCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableAccessCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTableAccessCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTableAccessCountResponse) GetBody() *DescribeTableAccessCountResponseBody {
	return s.Body
}

func (s *DescribeTableAccessCountResponse) SetHeaders(v map[string]*string) *DescribeTableAccessCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableAccessCountResponse) SetStatusCode(v int32) *DescribeTableAccessCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableAccessCountResponse) SetBody(v *DescribeTableAccessCountResponseBody) *DescribeTableAccessCountResponse {
	s.Body = v
	return s
}

func (s *DescribeTableAccessCountResponse) Validate() error {
	return dara.Validate(s)
}
