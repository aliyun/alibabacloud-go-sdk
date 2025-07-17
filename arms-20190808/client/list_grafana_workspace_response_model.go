// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGrafanaWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGrafanaWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *ListGrafanaWorkspaceResponseBody) *ListGrafanaWorkspaceResponse
	GetBody() *ListGrafanaWorkspaceResponseBody
}

type ListGrafanaWorkspaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGrafanaWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGrafanaWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGrafanaWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGrafanaWorkspaceResponse) GetBody() *ListGrafanaWorkspaceResponseBody {
	return s.Body
}

func (s *ListGrafanaWorkspaceResponse) SetHeaders(v map[string]*string) *ListGrafanaWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *ListGrafanaWorkspaceResponse) SetStatusCode(v int32) *ListGrafanaWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGrafanaWorkspaceResponse) SetBody(v *ListGrafanaWorkspaceResponseBody) *ListGrafanaWorkspaceResponse {
	s.Body = v
	return s
}

func (s *ListGrafanaWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
