// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBasicStartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBasicStartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBasicStartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBasicStartResponseBody) *DescribeBasicStartResponse
	GetBody() *DescribeBasicStartResponseBody
}

type DescribeBasicStartResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBasicStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBasicStartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicStartResponse) GoString() string {
	return s.String()
}

func (s *DescribeBasicStartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBasicStartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBasicStartResponse) GetBody() *DescribeBasicStartResponseBody {
	return s.Body
}

func (s *DescribeBasicStartResponse) SetHeaders(v map[string]*string) *DescribeBasicStartResponse {
	s.Headers = v
	return s
}

func (s *DescribeBasicStartResponse) SetStatusCode(v int32) *DescribeBasicStartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBasicStartResponse) SetBody(v *DescribeBasicStartResponseBody) *DescribeBasicStartResponse {
	s.Body = v
	return s
}

func (s *DescribeBasicStartResponse) Validate() error {
	return dara.Validate(s)
}
