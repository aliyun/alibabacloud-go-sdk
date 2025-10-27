// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScheduleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyScheduleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyScheduleConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyScheduleConfigResponseBody) *DescribePropertyScheduleConfigResponse
	GetBody() *DescribePropertyScheduleConfigResponseBody
}

type DescribePropertyScheduleConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyScheduleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyScheduleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyScheduleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyScheduleConfigResponse) GetBody() *DescribePropertyScheduleConfigResponseBody {
	return s.Body
}

func (s *DescribePropertyScheduleConfigResponse) SetHeaders(v map[string]*string) *DescribePropertyScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyScheduleConfigResponse) SetStatusCode(v int32) *DescribePropertyScheduleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyScheduleConfigResponse) SetBody(v *DescribePropertyScheduleConfigResponseBody) *DescribePropertyScheduleConfigResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyScheduleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
