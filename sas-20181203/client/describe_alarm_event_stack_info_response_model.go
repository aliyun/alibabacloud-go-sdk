// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmEventStackInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlarmEventStackInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlarmEventStackInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlarmEventStackInfoResponseBody) *DescribeAlarmEventStackInfoResponse
	GetBody() *DescribeAlarmEventStackInfoResponseBody
}

type DescribeAlarmEventStackInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmEventStackInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmEventStackInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlarmEventStackInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlarmEventStackInfoResponse) GetBody() *DescribeAlarmEventStackInfoResponseBody {
	return s.Body
}

func (s *DescribeAlarmEventStackInfoResponse) SetHeaders(v map[string]*string) *DescribeAlarmEventStackInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmEventStackInfoResponse) SetStatusCode(v int32) *DescribeAlarmEventStackInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmEventStackInfoResponse) SetBody(v *DescribeAlarmEventStackInfoResponseBody) *DescribeAlarmEventStackInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeAlarmEventStackInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
