// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudSdkServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridCloudSdkServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridCloudSdkServersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridCloudSdkServersResponseBody) *DescribeHybridCloudSdkServersResponse
	GetBody() *DescribeHybridCloudSdkServersResponseBody
}

type DescribeHybridCloudSdkServersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridCloudSdkServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridCloudSdkServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSdkServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSdkServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridCloudSdkServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridCloudSdkServersResponse) GetBody() *DescribeHybridCloudSdkServersResponseBody {
	return s.Body
}

func (s *DescribeHybridCloudSdkServersResponse) SetHeaders(v map[string]*string) *DescribeHybridCloudSdkServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridCloudSdkServersResponse) SetStatusCode(v int32) *DescribeHybridCloudSdkServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponse) SetBody(v *DescribeHybridCloudSdkServersResponseBody) *DescribeHybridCloudSdkServersResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridCloudSdkServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
