// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotificationConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNotificationConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNotificationConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNotificationConfigurationsResponseBody) *DescribeNotificationConfigurationsResponse
	GetBody() *DescribeNotificationConfigurationsResponseBody
}

type DescribeNotificationConfigurationsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNotificationConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNotificationConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNotificationConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNotificationConfigurationsResponse) GetBody() *DescribeNotificationConfigurationsResponseBody {
	return s.Body
}

func (s *DescribeNotificationConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeNotificationConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotificationConfigurationsResponse) SetStatusCode(v int32) *DescribeNotificationConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponse) SetBody(v *DescribeNotificationConfigurationsResponseBody) *DescribeNotificationConfigurationsResponse {
	s.Body = v
	return s
}

func (s *DescribeNotificationConfigurationsResponse) Validate() error {
	return dara.Validate(s)
}
