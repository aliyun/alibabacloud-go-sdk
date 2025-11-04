// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertConfigurationResponseBody) *DescribeAlertConfigurationResponse
	GetBody() *DescribeAlertConfigurationResponseBody
}

type DescribeAlertConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertConfigurationResponse) GetBody() *DescribeAlertConfigurationResponseBody {
	return s.Body
}

func (s *DescribeAlertConfigurationResponse) SetHeaders(v map[string]*string) *DescribeAlertConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertConfigurationResponse) SetStatusCode(v int32) *DescribeAlertConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertConfigurationResponse) SetBody(v *DescribeAlertConfigurationResponseBody) *DescribeAlertConfigurationResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
