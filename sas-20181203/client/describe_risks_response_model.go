// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRisksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRisksResponseBody) *DescribeRisksResponse
	GetBody() *DescribeRisksResponseBody
}

type DescribeRisksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRisksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeRisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRisksResponse) GetBody() *DescribeRisksResponseBody {
	return s.Body
}

func (s *DescribeRisksResponse) SetHeaders(v map[string]*string) *DescribeRisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeRisksResponse) SetStatusCode(v int32) *DescribeRisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRisksResponse) SetBody(v *DescribeRisksResponseBody) *DescribeRisksResponse {
	s.Body = v
	return s
}

func (s *DescribeRisksResponse) Validate() error {
	return dara.Validate(s)
}
