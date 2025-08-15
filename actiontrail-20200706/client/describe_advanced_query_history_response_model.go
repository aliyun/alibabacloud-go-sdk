// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedQueryHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvancedQueryHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvancedQueryHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvancedQueryHistoryResponseBody) *DescribeAdvancedQueryHistoryResponse
	GetBody() *DescribeAdvancedQueryHistoryResponseBody
}

type DescribeAdvancedQueryHistoryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvancedQueryHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvancedQueryHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvancedQueryHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvancedQueryHistoryResponse) GetBody() *DescribeAdvancedQueryHistoryResponseBody {
	return s.Body
}

func (s *DescribeAdvancedQueryHistoryResponse) SetHeaders(v map[string]*string) *DescribeAdvancedQueryHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponse) SetStatusCode(v int32) *DescribeAdvancedQueryHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponse) SetBody(v *DescribeAdvancedQueryHistoryResponseBody) *DescribeAdvancedQueryHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponse) Validate() error {
	return dara.Validate(s)
}
