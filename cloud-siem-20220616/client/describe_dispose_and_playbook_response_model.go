// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeAndPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDisposeAndPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDisposeAndPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDisposeAndPlaybookResponseBody) *DescribeDisposeAndPlaybookResponse
	GetBody() *DescribeDisposeAndPlaybookResponseBody
}

type DescribeDisposeAndPlaybookResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisposeAndPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDisposeAndPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDisposeAndPlaybookResponse) GetBody() *DescribeDisposeAndPlaybookResponseBody {
	return s.Body
}

func (s *DescribeDisposeAndPlaybookResponse) SetHeaders(v map[string]*string) *DescribeDisposeAndPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) SetStatusCode(v int32) *DescribeDisposeAndPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) SetBody(v *DescribeDisposeAndPlaybookResponseBody) *DescribeDisposeAndPlaybookResponse {
	s.Body = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) Validate() error {
	return dara.Validate(s)
}
