// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallEnvironmentFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallEnvironmentFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallEnvironmentFeatureResponse
	GetStatusCode() *int32
	SetBody(v *InstallEnvironmentFeatureResponseBody) *InstallEnvironmentFeatureResponse
	GetBody() *InstallEnvironmentFeatureResponseBody
}

type InstallEnvironmentFeatureResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallEnvironmentFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallEnvironmentFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallEnvironmentFeatureResponse) GoString() string {
	return s.String()
}

func (s *InstallEnvironmentFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallEnvironmentFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallEnvironmentFeatureResponse) GetBody() *InstallEnvironmentFeatureResponseBody {
	return s.Body
}

func (s *InstallEnvironmentFeatureResponse) SetHeaders(v map[string]*string) *InstallEnvironmentFeatureResponse {
	s.Headers = v
	return s
}

func (s *InstallEnvironmentFeatureResponse) SetStatusCode(v int32) *InstallEnvironmentFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallEnvironmentFeatureResponse) SetBody(v *InstallEnvironmentFeatureResponseBody) *InstallEnvironmentFeatureResponse {
	s.Body = v
	return s
}

func (s *InstallEnvironmentFeatureResponse) Validate() error {
	return dara.Validate(s)
}
