// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallCloudAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallCloudAppResponse
	GetStatusCode() *int32
	SetBody(v *InstallCloudAppResponseBody) *InstallCloudAppResponse
	GetBody() *InstallCloudAppResponseBody
}

type InstallCloudAppResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCloudAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCloudAppResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallCloudAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallCloudAppResponse) GetBody() *InstallCloudAppResponseBody {
	return s.Body
}

func (s *InstallCloudAppResponse) SetHeaders(v map[string]*string) *InstallCloudAppResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudAppResponse) SetStatusCode(v int32) *InstallCloudAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudAppResponse) SetBody(v *InstallCloudAppResponseBody) *InstallCloudAppResponse {
	s.Body = v
	return s
}

func (s *InstallCloudAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
