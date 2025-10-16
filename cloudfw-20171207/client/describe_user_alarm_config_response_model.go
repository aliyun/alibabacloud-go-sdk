// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAlarmConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAlarmConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAlarmConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAlarmConfigResponseBody) *DescribeUserAlarmConfigResponse
	GetBody() *DescribeUserAlarmConfigResponseBody
}

type DescribeUserAlarmConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAlarmConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAlarmConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlarmConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAlarmConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAlarmConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAlarmConfigResponse) GetBody() *DescribeUserAlarmConfigResponseBody {
	return s.Body
}

func (s *DescribeUserAlarmConfigResponse) SetHeaders(v map[string]*string) *DescribeUserAlarmConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAlarmConfigResponse) SetStatusCode(v int32) *DescribeUserAlarmConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAlarmConfigResponse) SetBody(v *DescribeUserAlarmConfigResponseBody) *DescribeUserAlarmConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAlarmConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
