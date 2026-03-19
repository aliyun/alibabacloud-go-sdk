// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHealthCheckListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHealthCheckListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHealthCheckListResponseBody) *DescribeHealthCheckListResponse
	GetBody() *DescribeHealthCheckListResponseBody
}

type DescribeHealthCheckListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthCheckListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHealthCheckListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHealthCheckListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHealthCheckListResponse) GetBody() *DescribeHealthCheckListResponseBody {
	return s.Body
}

func (s *DescribeHealthCheckListResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckListResponse) SetStatusCode(v int32) *DescribeHealthCheckListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthCheckListResponse) SetBody(v *DescribeHealthCheckListResponseBody) *DescribeHealthCheckListResponse {
	s.Body = v
	return s
}

func (s *DescribeHealthCheckListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
