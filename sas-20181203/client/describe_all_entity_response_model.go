// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllEntityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllEntityResponseBody) *DescribeAllEntityResponse
	GetBody() *DescribeAllEntityResponseBody
}

type DescribeAllEntityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEntityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllEntityResponse) GetBody() *DescribeAllEntityResponseBody {
	return s.Body
}

func (s *DescribeAllEntityResponse) SetHeaders(v map[string]*string) *DescribeAllEntityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllEntityResponse) SetStatusCode(v int32) *DescribeAllEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllEntityResponse) SetBody(v *DescribeAllEntityResponseBody) *DescribeAllEntityResponse {
	s.Body = v
	return s
}

func (s *DescribeAllEntityResponse) Validate() error {
	return dara.Validate(s)
}
