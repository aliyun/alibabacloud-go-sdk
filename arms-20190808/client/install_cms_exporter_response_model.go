// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCmsExporterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallCmsExporterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallCmsExporterResponse
	GetStatusCode() *int32
	SetBody(v *InstallCmsExporterResponseBody) *InstallCmsExporterResponse
	GetBody() *InstallCmsExporterResponseBody
}

type InstallCmsExporterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCmsExporterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCmsExporterResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallCmsExporterResponse) GoString() string {
	return s.String()
}

func (s *InstallCmsExporterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallCmsExporterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallCmsExporterResponse) GetBody() *InstallCmsExporterResponseBody {
	return s.Body
}

func (s *InstallCmsExporterResponse) SetHeaders(v map[string]*string) *InstallCmsExporterResponse {
	s.Headers = v
	return s
}

func (s *InstallCmsExporterResponse) SetStatusCode(v int32) *InstallCmsExporterResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCmsExporterResponse) SetBody(v *InstallCmsExporterResponseBody) *InstallCmsExporterResponse {
	s.Body = v
	return s
}

func (s *InstallCmsExporterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
