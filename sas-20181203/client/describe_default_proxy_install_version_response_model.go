// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultProxyInstallVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefaultProxyInstallVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefaultProxyInstallVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefaultProxyInstallVersionResponseBody) *DescribeDefaultProxyInstallVersionResponse
	GetBody() *DescribeDefaultProxyInstallVersionResponseBody
}

type DescribeDefaultProxyInstallVersionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefaultProxyInstallVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefaultProxyInstallVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultProxyInstallVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefaultProxyInstallVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefaultProxyInstallVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefaultProxyInstallVersionResponse) GetBody() *DescribeDefaultProxyInstallVersionResponseBody {
	return s.Body
}

func (s *DescribeDefaultProxyInstallVersionResponse) SetHeaders(v map[string]*string) *DescribeDefaultProxyInstallVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefaultProxyInstallVersionResponse) SetStatusCode(v int32) *DescribeDefaultProxyInstallVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefaultProxyInstallVersionResponse) SetBody(v *DescribeDefaultProxyInstallVersionResponseBody) *DescribeDefaultProxyInstallVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeDefaultProxyInstallVersionResponse) Validate() error {
	return dara.Validate(s)
}
