// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrafanaWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGrafanaWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGrafanaWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *GetGrafanaWorkspaceResponseBody) *GetGrafanaWorkspaceResponse
	GetBody() *GetGrafanaWorkspaceResponseBody
}

type GetGrafanaWorkspaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGrafanaWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGrafanaWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetGrafanaWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGrafanaWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGrafanaWorkspaceResponse) GetBody() *GetGrafanaWorkspaceResponseBody {
	return s.Body
}

func (s *GetGrafanaWorkspaceResponse) SetHeaders(v map[string]*string) *GetGrafanaWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetGrafanaWorkspaceResponse) SetStatusCode(v int32) *GetGrafanaWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGrafanaWorkspaceResponse) SetBody(v *GetGrafanaWorkspaceResponseBody) *GetGrafanaWorkspaceResponse {
	s.Body = v
	return s
}

func (s *GetGrafanaWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
