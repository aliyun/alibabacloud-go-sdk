// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestGraphResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRequestGraphResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRequestGraphResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRequestGraphResponseBody) *DescribeRequestGraphResponse
	GetBody() *DescribeRequestGraphResponseBody
}

type DescribeRequestGraphResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRequestGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRequestGraphResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRequestGraphResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRequestGraphResponse) GetBody() *DescribeRequestGraphResponseBody {
	return s.Body
}

func (s *DescribeRequestGraphResponse) SetHeaders(v map[string]*string) *DescribeRequestGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestGraphResponse) SetStatusCode(v int32) *DescribeRequestGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRequestGraphResponse) SetBody(v *DescribeRequestGraphResponseBody) *DescribeRequestGraphResponse {
	s.Body = v
	return s
}

func (s *DescribeRequestGraphResponse) Validate() error {
	return dara.Validate(s)
}
