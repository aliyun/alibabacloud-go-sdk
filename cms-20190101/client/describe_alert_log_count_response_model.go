// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertLogCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertLogCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertLogCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertLogCountResponseBody) *DescribeAlertLogCountResponse
	GetBody() *DescribeAlertLogCountResponseBody
}

type DescribeAlertLogCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertLogCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertLogCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertLogCountResponse) GetBody() *DescribeAlertLogCountResponseBody {
	return s.Body
}

func (s *DescribeAlertLogCountResponse) SetHeaders(v map[string]*string) *DescribeAlertLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertLogCountResponse) SetStatusCode(v int32) *DescribeAlertLogCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertLogCountResponse) SetBody(v *DescribeAlertLogCountResponseBody) *DescribeAlertLogCountResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertLogCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
