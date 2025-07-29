// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallClusterAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallClusterAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallClusterAddonsResponse
	GetStatusCode() *int32
	SetBody(v *InstallClusterAddonsResponseBody) *InstallClusterAddonsResponse
	GetBody() *InstallClusterAddonsResponseBody
}

type InstallClusterAddonsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallClusterAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallClusterAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsResponse) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallClusterAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallClusterAddonsResponse) GetBody() *InstallClusterAddonsResponseBody {
	return s.Body
}

func (s *InstallClusterAddonsResponse) SetHeaders(v map[string]*string) *InstallClusterAddonsResponse {
	s.Headers = v
	return s
}

func (s *InstallClusterAddonsResponse) SetStatusCode(v int32) *InstallClusterAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallClusterAddonsResponse) SetBody(v *InstallClusterAddonsResponseBody) *InstallClusterAddonsResponse {
	s.Body = v
	return s
}

func (s *InstallClusterAddonsResponse) Validate() error {
	return dara.Validate(s)
}
