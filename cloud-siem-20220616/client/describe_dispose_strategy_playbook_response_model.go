// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeStrategyPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDisposeStrategyPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDisposeStrategyPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDisposeStrategyPlaybookResponseBody) *DescribeDisposeStrategyPlaybookResponse
	GetBody() *DescribeDisposeStrategyPlaybookResponseBody
}

type DescribeDisposeStrategyPlaybookResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisposeStrategyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDisposeStrategyPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDisposeStrategyPlaybookResponse) GetBody() *DescribeDisposeStrategyPlaybookResponseBody {
	return s.Body
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetHeaders(v map[string]*string) *DescribeDisposeStrategyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetStatusCode(v int32) *DescribeDisposeStrategyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetBody(v *DescribeDisposeStrategyPlaybookResponseBody) *DescribeDisposeStrategyPlaybookResponse {
	s.Body = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) Validate() error {
	return dara.Validate(s)
}
