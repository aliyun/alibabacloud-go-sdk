// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmInstanceConfigAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmInstanceConfigAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmInstanceConfigAlertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmInstanceConfigAlertResponseBody) *DescribeCloudGtmInstanceConfigAlertResponse
	GetBody() *DescribeCloudGtmInstanceConfigAlertResponseBody
}

type DescribeCloudGtmInstanceConfigAlertResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmInstanceConfigAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) GetBody() *DescribeCloudGtmInstanceConfigAlertResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmInstanceConfigAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) SetStatusCode(v int32) *DescribeCloudGtmInstanceConfigAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) SetBody(v *DescribeCloudGtmInstanceConfigAlertResponseBody) *DescribeCloudGtmInstanceConfigAlertResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertResponse) Validate() error {
	return dara.Validate(s)
}
