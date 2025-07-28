// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResponseCodeTrendGraphResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResponseCodeTrendGraphResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResponseCodeTrendGraphResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResponseCodeTrendGraphResponseBody) *DescribeResponseCodeTrendGraphResponse
	GetBody() *DescribeResponseCodeTrendGraphResponseBody
}

type DescribeResponseCodeTrendGraphResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResponseCodeTrendGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResponseCodeTrendGraphResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResponseCodeTrendGraphResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResponseCodeTrendGraphResponse) GetBody() *DescribeResponseCodeTrendGraphResponseBody {
	return s.Body
}

func (s *DescribeResponseCodeTrendGraphResponse) SetHeaders(v map[string]*string) *DescribeResponseCodeTrendGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) SetStatusCode(v int32) *DescribeResponseCodeTrendGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) SetBody(v *DescribeResponseCodeTrendGraphResponseBody) *DescribeResponseCodeTrendGraphResponse {
	s.Body = v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) Validate() error {
	return dara.Validate(s)
}
