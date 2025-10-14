// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertLogListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertLogListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertLogListResponseBody) *DescribeAlertLogListResponse
	GetBody() *DescribeAlertLogListResponseBody
}

type DescribeAlertLogListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertLogListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertLogListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertLogListResponse) GetBody() *DescribeAlertLogListResponseBody {
	return s.Body
}

func (s *DescribeAlertLogListResponse) SetHeaders(v map[string]*string) *DescribeAlertLogListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertLogListResponse) SetStatusCode(v int32) *DescribeAlertLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertLogListResponse) SetBody(v *DescribeAlertLogListResponseBody) *DescribeAlertLogListResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertLogListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
