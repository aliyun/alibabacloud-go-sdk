// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScheduleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScheduleResponseBody) *DescribeScheduleResponse
	GetBody() *DescribeScheduleResponseBody
}

type DescribeScheduleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScheduleResponse) GetBody() *DescribeScheduleResponseBody {
	return s.Body
}

func (s *DescribeScheduleResponse) SetHeaders(v map[string]*string) *DescribeScheduleResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleResponse) SetStatusCode(v int32) *DescribeScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduleResponse) SetBody(v *DescribeScheduleResponseBody) *DescribeScheduleResponse {
	s.Body = v
	return s
}

func (s *DescribeScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
