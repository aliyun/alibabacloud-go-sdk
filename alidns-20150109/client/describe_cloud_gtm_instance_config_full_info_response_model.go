// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmInstanceConfigFullInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmInstanceConfigFullInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmInstanceConfigFullInfoResponseBody) *DescribeCloudGtmInstanceConfigFullInfoResponse
	GetBody() *DescribeCloudGtmInstanceConfigFullInfoResponseBody
}

type DescribeCloudGtmInstanceConfigFullInfoResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmInstanceConfigFullInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) GetBody() *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmInstanceConfigFullInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) SetStatusCode(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) SetBody(v *DescribeCloudGtmInstanceConfigFullInfoResponseBody) *DescribeCloudGtmInstanceConfigFullInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponse) Validate() error {
	return dara.Validate(s)
}
