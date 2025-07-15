// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppResponseBody) *DescribeAppResponse
	GetBody() *DescribeAppResponseBody
}

type DescribeAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppResponse) GetBody() *DescribeAppResponseBody {
	return s.Body
}

func (s *DescribeAppResponse) SetHeaders(v map[string]*string) *DescribeAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppResponse) SetStatusCode(v int32) *DescribeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppResponse) SetBody(v *DescribeAppResponseBody) *DescribeAppResponse {
	s.Body = v
	return s
}

func (s *DescribeAppResponse) Validate() error {
	return dara.Validate(s)
}
