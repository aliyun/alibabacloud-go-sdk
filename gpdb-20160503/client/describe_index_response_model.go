// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIndexResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIndexResponseBody) *DescribeIndexResponse
	GetBody() *DescribeIndexResponseBody
}

type DescribeIndexResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexResponse) GoString() string {
	return s.String()
}

func (s *DescribeIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIndexResponse) GetBody() *DescribeIndexResponseBody {
	return s.Body
}

func (s *DescribeIndexResponse) SetHeaders(v map[string]*string) *DescribeIndexResponse {
	s.Headers = v
	return s
}

func (s *DescribeIndexResponse) SetStatusCode(v int32) *DescribeIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIndexResponse) SetBody(v *DescribeIndexResponseBody) *DescribeIndexResponse {
	s.Body = v
	return s
}

func (s *DescribeIndexResponse) Validate() error {
	return dara.Validate(s)
}
