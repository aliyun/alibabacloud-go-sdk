// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGrafanaWorkspaceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGrafanaWorkspaceAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListGrafanaWorkspaceAccountResponseBody) *ListGrafanaWorkspaceAccountResponse
	GetBody() *ListGrafanaWorkspaceAccountResponseBody
}

type ListGrafanaWorkspaceAccountResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGrafanaWorkspaceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGrafanaWorkspaceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceAccountResponse) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGrafanaWorkspaceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGrafanaWorkspaceAccountResponse) GetBody() *ListGrafanaWorkspaceAccountResponseBody {
	return s.Body
}

func (s *ListGrafanaWorkspaceAccountResponse) SetHeaders(v map[string]*string) *ListGrafanaWorkspaceAccountResponse {
	s.Headers = v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponse) SetStatusCode(v int32) *ListGrafanaWorkspaceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponse) SetBody(v *ListGrafanaWorkspaceAccountResponseBody) *ListGrafanaWorkspaceAccountResponse {
	s.Body = v
	return s
}

func (s *ListGrafanaWorkspaceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
