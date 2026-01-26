// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGrafanaWorkspaceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGrafanaWorkspaceVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGrafanaWorkspaceVersionResponseBody) *UpdateGrafanaWorkspaceVersionResponse
	GetBody() *UpdateGrafanaWorkspaceVersionResponseBody
}

type UpdateGrafanaWorkspaceVersionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGrafanaWorkspaceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGrafanaWorkspaceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGrafanaWorkspaceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGrafanaWorkspaceVersionResponse) GetBody() *UpdateGrafanaWorkspaceVersionResponseBody {
	return s.Body
}

func (s *UpdateGrafanaWorkspaceVersionResponse) SetHeaders(v map[string]*string) *UpdateGrafanaWorkspaceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponse) SetStatusCode(v int32) *UpdateGrafanaWorkspaceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponse) SetBody(v *UpdateGrafanaWorkspaceVersionResponseBody) *UpdateGrafanaWorkspaceVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
