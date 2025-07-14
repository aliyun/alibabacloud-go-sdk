// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobHistoryResponseBody) *DescribeJobHistoryResponse
	GetBody() *DescribeJobHistoryResponseBody
}

type DescribeJobHistoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobHistoryResponse) GetBody() *DescribeJobHistoryResponseBody {
	return s.Body
}

func (s *DescribeJobHistoryResponse) SetHeaders(v map[string]*string) *DescribeJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobHistoryResponse) SetStatusCode(v int32) *DescribeJobHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobHistoryResponse) SetBody(v *DescribeJobHistoryResponseBody) *DescribeJobHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeJobHistoryResponse) Validate() error {
	return dara.Validate(s)
}
