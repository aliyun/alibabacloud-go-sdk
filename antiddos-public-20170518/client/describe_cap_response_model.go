// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCapResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCapResponseBody) *DescribeCapResponse
	GetBody() *DescribeCapResponseBody
}

type DescribeCapResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCapResponse) GetBody() *DescribeCapResponseBody {
	return s.Body
}

func (s *DescribeCapResponse) SetHeaders(v map[string]*string) *DescribeCapResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapResponse) SetStatusCode(v int32) *DescribeCapResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapResponse) SetBody(v *DescribeCapResponseBody) *DescribeCapResponse {
	s.Body = v
	return s
}

func (s *DescribeCapResponse) Validate() error {
	return dara.Validate(s)
}
