// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcludeSystemPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExcludeSystemPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExcludeSystemPathResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExcludeSystemPathResponseBody) *DescribeExcludeSystemPathResponse
	GetBody() *DescribeExcludeSystemPathResponseBody
}

type DescribeExcludeSystemPathResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExcludeSystemPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExcludeSystemPathResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcludeSystemPathResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcludeSystemPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExcludeSystemPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExcludeSystemPathResponse) GetBody() *DescribeExcludeSystemPathResponseBody {
	return s.Body
}

func (s *DescribeExcludeSystemPathResponse) SetHeaders(v map[string]*string) *DescribeExcludeSystemPathResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcludeSystemPathResponse) SetStatusCode(v int32) *DescribeExcludeSystemPathResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExcludeSystemPathResponse) SetBody(v *DescribeExcludeSystemPathResponseBody) *DescribeExcludeSystemPathResponse {
	s.Body = v
	return s
}

func (s *DescribeExcludeSystemPathResponse) Validate() error {
	return dara.Validate(s)
}
