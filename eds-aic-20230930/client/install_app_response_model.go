// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAppResponse
	GetStatusCode() *int32
	SetBody(v *InstallAppResponseBody) *InstallAppResponse
	GetBody() *InstallAppResponseBody
}

type InstallAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAppResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAppResponse) GoString() string {
	return s.String()
}

func (s *InstallAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAppResponse) GetBody() *InstallAppResponseBody {
	return s.Body
}

func (s *InstallAppResponse) SetHeaders(v map[string]*string) *InstallAppResponse {
	s.Headers = v
	return s
}

func (s *InstallAppResponse) SetStatusCode(v int32) *InstallAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAppResponse) SetBody(v *InstallAppResponseBody) *InstallAppResponse {
	s.Body = v
	return s
}

func (s *InstallAppResponse) Validate() error {
	return dara.Validate(s)
}
