// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityCheckScheduleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityCheckScheduleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityCheckScheduleConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityCheckScheduleConfigResponseBody) *DescribeSecurityCheckScheduleConfigResponse
	GetBody() *DescribeSecurityCheckScheduleConfigResponseBody
}

type DescribeSecurityCheckScheduleConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityCheckScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityCheckScheduleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityCheckScheduleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityCheckScheduleConfigResponse) GetBody() *DescribeSecurityCheckScheduleConfigResponseBody {
	return s.Body
}

func (s *DescribeSecurityCheckScheduleConfigResponse) SetHeaders(v map[string]*string) *DescribeSecurityCheckScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponse) SetStatusCode(v int32) *DescribeSecurityCheckScheduleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponse) SetBody(v *DescribeSecurityCheckScheduleConfigResponseBody) *DescribeSecurityCheckScheduleConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponse) Validate() error {
	return dara.Validate(s)
}
