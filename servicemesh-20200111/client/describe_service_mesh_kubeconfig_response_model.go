// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshKubeconfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshKubeconfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshKubeconfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshKubeconfigResponseBody) *DescribeServiceMeshKubeconfigResponse
	GetBody() *DescribeServiceMeshKubeconfigResponseBody
}

type DescribeServiceMeshKubeconfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshKubeconfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshKubeconfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshKubeconfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshKubeconfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshKubeconfigResponse) GetBody() *DescribeServiceMeshKubeconfigResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshKubeconfigResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshKubeconfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponse) SetStatusCode(v int32) *DescribeServiceMeshKubeconfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponse) SetBody(v *DescribeServiceMeshKubeconfigResponseBody) *DescribeServiceMeshKubeconfigResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshKubeconfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
