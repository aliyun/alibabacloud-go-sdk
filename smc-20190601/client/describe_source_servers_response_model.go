// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSourceServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSourceServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSourceServersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSourceServersResponseBody) *DescribeSourceServersResponse
	GetBody() *DescribeSourceServersResponseBody
}

type DescribeSourceServersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSourceServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSourceServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSourceServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSourceServersResponse) GetBody() *DescribeSourceServersResponseBody {
	return s.Body
}

func (s *DescribeSourceServersResponse) SetHeaders(v map[string]*string) *DescribeSourceServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSourceServersResponse) SetStatusCode(v int32) *DescribeSourceServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSourceServersResponse) SetBody(v *DescribeSourceServersResponseBody) *DescribeSourceServersResponse {
	s.Body = v
	return s
}

func (s *DescribeSourceServersResponse) Validate() error {
	return dara.Validate(s)
}
