// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperatorListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperatorListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperatorListResponseBody) *DescribeOperatorListResponse
	GetBody() *DescribeOperatorListResponseBody
}

type DescribeOperatorListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperatorListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperatorListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperatorListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperatorListResponse) GetBody() *DescribeOperatorListResponseBody {
	return s.Body
}

func (s *DescribeOperatorListResponse) SetHeaders(v map[string]*string) *DescribeOperatorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorListResponse) SetStatusCode(v int32) *DescribeOperatorListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorListResponse) SetBody(v *DescribeOperatorListResponseBody) *DescribeOperatorListResponse {
	s.Body = v
	return s
}

func (s *DescribeOperatorListResponse) Validate() error {
	return dara.Validate(s)
}
