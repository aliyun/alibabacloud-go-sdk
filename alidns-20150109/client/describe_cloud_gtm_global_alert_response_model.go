// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmGlobalAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmGlobalAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmGlobalAlertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmGlobalAlertResponseBody) *DescribeCloudGtmGlobalAlertResponse
	GetBody() *DescribeCloudGtmGlobalAlertResponseBody
}

type DescribeCloudGtmGlobalAlertResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmGlobalAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmGlobalAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmGlobalAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmGlobalAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmGlobalAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmGlobalAlertResponse) GetBody() *DescribeCloudGtmGlobalAlertResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmGlobalAlertResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmGlobalAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponse) SetStatusCode(v int32) *DescribeCloudGtmGlobalAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponse) SetBody(v *DescribeCloudGtmGlobalAlertResponseBody) *DescribeCloudGtmGlobalAlertResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmGlobalAlertResponse) Validate() error {
	return dara.Validate(s)
}
