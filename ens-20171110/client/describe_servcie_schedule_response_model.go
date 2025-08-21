// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServcieScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServcieScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServcieScheduleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServcieScheduleResponseBody) *DescribeServcieScheduleResponse
	GetBody() *DescribeServcieScheduleResponseBody
}

type DescribeServcieScheduleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServcieScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServcieScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServcieScheduleResponse) GoString() string {
	return s.String()
}

func (s *DescribeServcieScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServcieScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServcieScheduleResponse) GetBody() *DescribeServcieScheduleResponseBody {
	return s.Body
}

func (s *DescribeServcieScheduleResponse) SetHeaders(v map[string]*string) *DescribeServcieScheduleResponse {
	s.Headers = v
	return s
}

func (s *DescribeServcieScheduleResponse) SetStatusCode(v int32) *DescribeServcieScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServcieScheduleResponse) SetBody(v *DescribeServcieScheduleResponseBody) *DescribeServcieScheduleResponse {
	s.Body = v
	return s
}

func (s *DescribeServcieScheduleResponse) Validate() error {
	return dara.Validate(s)
}
