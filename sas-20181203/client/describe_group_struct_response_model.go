// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupStructResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupStructResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupStructResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupStructResponseBody) *DescribeGroupStructResponse
	GetBody() *DescribeGroupStructResponseBody
}

type DescribeGroupStructResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupStructResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupStructResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupStructResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupStructResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupStructResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupStructResponse) GetBody() *DescribeGroupStructResponseBody {
	return s.Body
}

func (s *DescribeGroupStructResponse) SetHeaders(v map[string]*string) *DescribeGroupStructResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupStructResponse) SetStatusCode(v int32) *DescribeGroupStructResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupStructResponse) SetBody(v *DescribeGroupStructResponseBody) *DescribeGroupStructResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupStructResponse) Validate() error {
	return dara.Validate(s)
}
