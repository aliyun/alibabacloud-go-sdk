// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorTrialConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudVendorTrialConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudVendorTrialConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudVendorTrialConfigResponseBody) *DescribeCloudVendorTrialConfigResponse
	GetBody() *DescribeCloudVendorTrialConfigResponseBody
}

type DescribeCloudVendorTrialConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudVendorTrialConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudVendorTrialConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorTrialConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorTrialConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudVendorTrialConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudVendorTrialConfigResponse) GetBody() *DescribeCloudVendorTrialConfigResponseBody {
	return s.Body
}

func (s *DescribeCloudVendorTrialConfigResponse) SetHeaders(v map[string]*string) *DescribeCloudVendorTrialConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudVendorTrialConfigResponse) SetStatusCode(v int32) *DescribeCloudVendorTrialConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudVendorTrialConfigResponse) SetBody(v *DescribeCloudVendorTrialConfigResponseBody) *DescribeCloudVendorTrialConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudVendorTrialConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
