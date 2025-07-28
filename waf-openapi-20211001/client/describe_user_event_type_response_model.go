// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEventTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserEventTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserEventTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserEventTypeResponseBody) *DescribeUserEventTypeResponse
	GetBody() *DescribeUserEventTypeResponseBody
}

type DescribeUserEventTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserEventTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserEventTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserEventTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserEventTypeResponse) GetBody() *DescribeUserEventTypeResponseBody {
	return s.Body
}

func (s *DescribeUserEventTypeResponse) SetHeaders(v map[string]*string) *DescribeUserEventTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserEventTypeResponse) SetStatusCode(v int32) *DescribeUserEventTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserEventTypeResponse) SetBody(v *DescribeUserEventTypeResponseBody) *DescribeUserEventTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeUserEventTypeResponse) Validate() error {
	return dara.Validate(s)
}
