// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatMemberDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStatMemberDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStatMemberDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStatMemberDeviceInfoResponseBody) *DescribeStatMemberDeviceInfoResponse
	GetBody() *DescribeStatMemberDeviceInfoResponseBody
}

type DescribeStatMemberDeviceInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStatMemberDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStatMemberDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatMemberDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatMemberDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStatMemberDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStatMemberDeviceInfoResponse) GetBody() *DescribeStatMemberDeviceInfoResponseBody {
	return s.Body
}

func (s *DescribeStatMemberDeviceInfoResponse) SetHeaders(v map[string]*string) *DescribeStatMemberDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponse) SetStatusCode(v int32) *DescribeStatMemberDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponse) SetBody(v *DescribeStatMemberDeviceInfoResponseBody) *DescribeStatMemberDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeStatMemberDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
