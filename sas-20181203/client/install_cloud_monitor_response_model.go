// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallCloudMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallCloudMonitorResponse
	GetStatusCode() *int32
	SetBody(v *InstallCloudMonitorResponseBody) *InstallCloudMonitorResponse
	GetBody() *InstallCloudMonitorResponseBody
}

type InstallCloudMonitorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCloudMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCloudMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudMonitorResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallCloudMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallCloudMonitorResponse) GetBody() *InstallCloudMonitorResponseBody {
	return s.Body
}

func (s *InstallCloudMonitorResponse) SetHeaders(v map[string]*string) *InstallCloudMonitorResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudMonitorResponse) SetStatusCode(v int32) *InstallCloudMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudMonitorResponse) SetBody(v *InstallCloudMonitorResponseBody) *InstallCloudMonitorResponse {
	s.Body = v
	return s
}

func (s *InstallCloudMonitorResponse) Validate() error {
	return dara.Validate(s)
}
