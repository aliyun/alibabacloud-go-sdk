// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGrafanaWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGrafanaWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGrafanaWorkspaceResponseBody) *UpdateGrafanaWorkspaceResponse
	GetBody() *UpdateGrafanaWorkspaceResponseBody
}

type UpdateGrafanaWorkspaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGrafanaWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGrafanaWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGrafanaWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGrafanaWorkspaceResponse) GetBody() *UpdateGrafanaWorkspaceResponseBody {
	return s.Body
}

func (s *UpdateGrafanaWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateGrafanaWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateGrafanaWorkspaceResponse) SetStatusCode(v int32) *UpdateGrafanaWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGrafanaWorkspaceResponse) SetBody(v *UpdateGrafanaWorkspaceResponseBody) *UpdateGrafanaWorkspaceResponse {
	s.Body = v
	return s
}

func (s *UpdateGrafanaWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
