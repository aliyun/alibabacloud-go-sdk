// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterUserKubeconfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterUserKubeconfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterUserKubeconfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterUserKubeconfigResponseBody) *DescribeClusterUserKubeconfigResponse
	GetBody() *DescribeClusterUserKubeconfigResponseBody
}

type DescribeClusterUserKubeconfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterUserKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterUserKubeconfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterUserKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterUserKubeconfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterUserKubeconfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterUserKubeconfigResponse) GetBody() *DescribeClusterUserKubeconfigResponseBody {
	return s.Body
}

func (s *DescribeClusterUserKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeClusterUserKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterUserKubeconfigResponse) SetStatusCode(v int32) *DescribeClusterUserKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterUserKubeconfigResponse) SetBody(v *DescribeClusterUserKubeconfigResponseBody) *DescribeClusterUserKubeconfigResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterUserKubeconfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
