// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterV2UserKubeconfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterV2UserKubeconfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterV2UserKubeconfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterV2UserKubeconfigResponseBody) *DescribeClusterV2UserKubeconfigResponse
	GetBody() *DescribeClusterV2UserKubeconfigResponseBody
}

type DescribeClusterV2UserKubeconfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterV2UserKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterV2UserKubeconfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterV2UserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2UserKubeconfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterV2UserKubeconfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterV2UserKubeconfigResponse) GetBody() *DescribeClusterV2UserKubeconfigResponseBody {
	return s.Body
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeClusterV2UserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetStatusCode(v int32) *DescribeClusterV2UserKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponse) SetBody(v *DescribeClusterV2UserKubeconfigResponseBody) *DescribeClusterV2UserKubeconfigResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterV2UserKubeconfigResponse) Validate() error {
	return dara.Validate(s)
}
