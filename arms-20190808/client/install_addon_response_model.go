// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAddonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAddonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAddonResponse
	GetStatusCode() *int32
	SetBody(v *InstallAddonResponseBody) *InstallAddonResponse
	GetBody() *InstallAddonResponseBody
}

type InstallAddonResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAddonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAddonResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonResponse) GoString() string {
	return s.String()
}

func (s *InstallAddonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAddonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAddonResponse) GetBody() *InstallAddonResponseBody {
	return s.Body
}

func (s *InstallAddonResponse) SetHeaders(v map[string]*string) *InstallAddonResponse {
	s.Headers = v
	return s
}

func (s *InstallAddonResponse) SetStatusCode(v int32) *InstallAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAddonResponse) SetBody(v *InstallAddonResponseBody) *InstallAddonResponse {
	s.Body = v
	return s
}

func (s *InstallAddonResponse) Validate() error {
	return dara.Validate(s)
}
