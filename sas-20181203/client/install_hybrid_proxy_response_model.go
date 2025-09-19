// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallHybridProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallHybridProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallHybridProxyResponse
	GetStatusCode() *int32
	SetBody(v *InstallHybridProxyResponseBody) *InstallHybridProxyResponse
	GetBody() *InstallHybridProxyResponseBody
}

type InstallHybridProxyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallHybridProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallHybridProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallHybridProxyResponse) GoString() string {
	return s.String()
}

func (s *InstallHybridProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallHybridProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallHybridProxyResponse) GetBody() *InstallHybridProxyResponseBody {
	return s.Body
}

func (s *InstallHybridProxyResponse) SetHeaders(v map[string]*string) *InstallHybridProxyResponse {
	s.Headers = v
	return s
}

func (s *InstallHybridProxyResponse) SetStatusCode(v int32) *InstallHybridProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallHybridProxyResponse) SetBody(v *InstallHybridProxyResponseBody) *InstallHybridProxyResponse {
	s.Body = v
	return s
}

func (s *InstallHybridProxyResponse) Validate() error {
	return dara.Validate(s)
}
