// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGrafanaWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGrafanaWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGrafanaWorkspaceResponseBody) *DeleteGrafanaWorkspaceResponse
	GetBody() *DeleteGrafanaWorkspaceResponseBody
}

type DeleteGrafanaWorkspaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGrafanaWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGrafanaWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGrafanaWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGrafanaWorkspaceResponse) GetBody() *DeleteGrafanaWorkspaceResponseBody {
	return s.Body
}

func (s *DeleteGrafanaWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteGrafanaWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGrafanaWorkspaceResponse) SetStatusCode(v int32) *DeleteGrafanaWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGrafanaWorkspaceResponse) SetBody(v *DeleteGrafanaWorkspaceResponseBody) *DeleteGrafanaWorkspaceResponse {
	s.Body = v
	return s
}

func (s *DeleteGrafanaWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
