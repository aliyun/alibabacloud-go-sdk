// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshProxyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshProxyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshProxyStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshProxyStatusResponseBody) *DescribeServiceMeshProxyStatusResponse
	GetBody() *DescribeServiceMeshProxyStatusResponseBody
}

type DescribeServiceMeshProxyStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshProxyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshProxyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshProxyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshProxyStatusResponse) GetBody() *DescribeServiceMeshProxyStatusResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshProxyStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshProxyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponse) SetStatusCode(v int32) *DescribeServiceMeshProxyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponse) SetBody(v *DescribeServiceMeshProxyStatusResponseBody) *DescribeServiceMeshProxyStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshProxyStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
