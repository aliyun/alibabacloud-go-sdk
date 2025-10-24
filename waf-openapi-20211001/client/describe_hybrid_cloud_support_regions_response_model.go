// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudSupportRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudSupportRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudSupportRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudSupportRegionsResponseBody) *DescribeHybridCloudSupportRegionsResponse
	GetBody() *DescribeHybridCloudSupportRegionsResponseBody
}

type DescribeHybridCloudSupportRegionsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudSupportRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudSupportRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSupportRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSupportRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudSupportRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudSupportRegionsResponse) GetBody() *DescribeHybridCloudSupportRegionsResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudSupportRegionsResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudSupportRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudSupportRegionsResponse) SetStatusCode(v int32) *DescribeHybridCloudSupportRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudSupportRegionsResponse) SetBody(v *DescribeHybridCloudSupportRegionsResponseBody) *DescribeHybridCloudSupportRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudSupportRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
