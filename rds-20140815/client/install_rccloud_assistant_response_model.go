// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallRCCloudAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallRCCloudAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallRCCloudAssistantResponse
	GetStatusCode() *int32
	SetBody(v *InstallRCCloudAssistantResponseBody) *InstallRCCloudAssistantResponse
	GetBody() *InstallRCCloudAssistantResponseBody
}

type InstallRCCloudAssistantResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallRCCloudAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallRCCloudAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallRCCloudAssistantResponse) GoString() string {
	return s.String()
}

func (s *InstallRCCloudAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallRCCloudAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallRCCloudAssistantResponse) GetBody() *InstallRCCloudAssistantResponseBody {
	return s.Body
}

func (s *InstallRCCloudAssistantResponse) SetHeaders(v map[string]*string) *InstallRCCloudAssistantResponse {
	s.Headers = v
	return s
}

func (s *InstallRCCloudAssistantResponse) SetStatusCode(v int32) *InstallRCCloudAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallRCCloudAssistantResponse) SetBody(v *InstallRCCloudAssistantResponseBody) *InstallRCCloudAssistantResponse {
	s.Body = v
	return s
}

func (s *InstallRCCloudAssistantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
