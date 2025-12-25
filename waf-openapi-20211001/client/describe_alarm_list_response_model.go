// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlarmListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlarmListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlarmListResponseBody) *DescribeAlarmListResponse
	GetBody() *DescribeAlarmListResponseBody
}

type DescribeAlarmListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlarmListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlarmListResponse) GetBody() *DescribeAlarmListResponseBody {
	return s.Body
}

func (s *DescribeAlarmListResponse) SetHeaders(v map[string]*string) *DescribeAlarmListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmListResponse) SetStatusCode(v int32) *DescribeAlarmListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmListResponse) SetBody(v *DescribeAlarmListResponseBody) *DescribeAlarmListResponse {
	s.Body = v
	return s
}

func (s *DescribeAlarmListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
