// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvAccountStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKvAccountStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKvAccountStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKvAccountStatusResponseBody) *DescribeKvAccountStatusResponse
	GetBody() *DescribeKvAccountStatusResponseBody
}

type DescribeKvAccountStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKvAccountStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKvAccountStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvAccountStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeKvAccountStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKvAccountStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKvAccountStatusResponse) GetBody() *DescribeKvAccountStatusResponseBody {
	return s.Body
}

func (s *DescribeKvAccountStatusResponse) SetHeaders(v map[string]*string) *DescribeKvAccountStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeKvAccountStatusResponse) SetStatusCode(v int32) *DescribeKvAccountStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKvAccountStatusResponse) SetBody(v *DescribeKvAccountStatusResponseBody) *DescribeKvAccountStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeKvAccountStatusResponse) Validate() error {
	return dara.Validate(s)
}
