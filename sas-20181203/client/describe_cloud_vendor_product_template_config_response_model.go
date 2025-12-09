// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorProductTemplateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudVendorProductTemplateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudVendorProductTemplateConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudVendorProductTemplateConfigResponseBody) *DescribeCloudVendorProductTemplateConfigResponse
	GetBody() *DescribeCloudVendorProductTemplateConfigResponseBody
}

type DescribeCloudVendorProductTemplateConfigResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudVendorProductTemplateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudVendorProductTemplateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorProductTemplateConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) GetBody() *DescribeCloudVendorProductTemplateConfigResponseBody {
	return s.Body
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) SetHeaders(v map[string]*string) *DescribeCloudVendorProductTemplateConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) SetStatusCode(v int32) *DescribeCloudVendorProductTemplateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) SetBody(v *DescribeCloudVendorProductTemplateConfigResponseBody) *DescribeCloudVendorProductTemplateConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
