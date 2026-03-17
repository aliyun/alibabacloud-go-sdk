// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudConnectNetworksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudConnectNetworksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudConnectNetworksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudConnectNetworksResponseBody) *DescribeCloudConnectNetworksResponse
	GetBody() *DescribeCloudConnectNetworksResponseBody
}

type DescribeCloudConnectNetworksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudConnectNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudConnectNetworksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudConnectNetworksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudConnectNetworksResponse) GetBody() *DescribeCloudConnectNetworksResponseBody {
	return s.Body
}

func (s *DescribeCloudConnectNetworksResponse) SetHeaders(v map[string]*string) *DescribeCloudConnectNetworksResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudConnectNetworksResponse) SetStatusCode(v int32) *DescribeCloudConnectNetworksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponse) SetBody(v *DescribeCloudConnectNetworksResponseBody) *DescribeCloudConnectNetworksResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudConnectNetworksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
