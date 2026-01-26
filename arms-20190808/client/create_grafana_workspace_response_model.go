// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGrafanaWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGrafanaWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGrafanaWorkspaceResponseBody) *CreateGrafanaWorkspaceResponse
	GetBody() *CreateGrafanaWorkspaceResponseBody
}

type CreateGrafanaWorkspaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGrafanaWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGrafanaWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGrafanaWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGrafanaWorkspaceResponse) GetBody() *CreateGrafanaWorkspaceResponseBody {
	return s.Body
}

func (s *CreateGrafanaWorkspaceResponse) SetHeaders(v map[string]*string) *CreateGrafanaWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateGrafanaWorkspaceResponse) SetStatusCode(v int32) *CreateGrafanaWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGrafanaWorkspaceResponse) SetBody(v *CreateGrafanaWorkspaceResponseBody) *CreateGrafanaWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreateGrafanaWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
