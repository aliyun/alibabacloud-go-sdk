// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientsResponseBody) *DescribeClientsResponse
	GetBody() *DescribeClientsResponseBody
}

type DescribeClientsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientsResponse) GetBody() *DescribeClientsResponseBody {
	return s.Body
}

func (s *DescribeClientsResponse) SetHeaders(v map[string]*string) *DescribeClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientsResponse) SetStatusCode(v int32) *DescribeClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientsResponse) SetBody(v *DescribeClientsResponseBody) *DescribeClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeClientsResponse) Validate() error {
	return dara.Validate(s)
}
