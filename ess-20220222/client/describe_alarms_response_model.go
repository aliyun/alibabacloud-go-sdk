// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlarmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlarmsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlarmsResponseBody) *DescribeAlarmsResponse
	GetBody() *DescribeAlarmsResponseBody
}

type DescribeAlarmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlarmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlarmsResponse) GetBody() *DescribeAlarmsResponseBody {
	return s.Body
}

func (s *DescribeAlarmsResponse) SetHeaders(v map[string]*string) *DescribeAlarmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmsResponse) SetStatusCode(v int32) *DescribeAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmsResponse) SetBody(v *DescribeAlarmsResponseBody) *DescribeAlarmsResponse {
	s.Body = v
	return s
}

func (s *DescribeAlarmsResponse) Validate() error {
	return dara.Validate(s)
}
