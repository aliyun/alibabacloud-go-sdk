// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceStatusResponseBody) *DescribeServiceStatusResponse
	GetBody() *DescribeServiceStatusResponseBody
}

type DescribeServiceStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceStatusResponse) GetBody() *DescribeServiceStatusResponseBody {
	return s.Body
}

func (s *DescribeServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceStatusResponse) SetStatusCode(v int32) *DescribeServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceStatusResponse) SetBody(v *DescribeServiceStatusResponseBody) *DescribeServiceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceStatusResponse) Validate() error {
	return dara.Validate(s)
}
