// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSiemOrderStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserSiemOrderStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserSiemOrderStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserSiemOrderStatusResponseBody) *DescribeUserSiemOrderStatusResponse
	GetBody() *DescribeUserSiemOrderStatusResponseBody
}

type DescribeUserSiemOrderStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserSiemOrderStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserSiemOrderStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSiemOrderStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserSiemOrderStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserSiemOrderStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserSiemOrderStatusResponse) GetBody() *DescribeUserSiemOrderStatusResponseBody {
	return s.Body
}

func (s *DescribeUserSiemOrderStatusResponse) SetHeaders(v map[string]*string) *DescribeUserSiemOrderStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserSiemOrderStatusResponse) SetStatusCode(v int32) *DescribeUserSiemOrderStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponse) SetBody(v *DescribeUserSiemOrderStatusResponseBody) *DescribeUserSiemOrderStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserSiemOrderStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
