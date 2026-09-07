// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGrafanaWorkspaceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGrafanaWorkspaceAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGrafanaWorkspaceAccountResponseBody) *DeleteGrafanaWorkspaceAccountResponse
	GetBody() *DeleteGrafanaWorkspaceAccountResponseBody
}

type DeleteGrafanaWorkspaceAccountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGrafanaWorkspaceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGrafanaWorkspaceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGrafanaWorkspaceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGrafanaWorkspaceAccountResponse) GetBody() *DeleteGrafanaWorkspaceAccountResponseBody {
	return s.Body
}

func (s *DeleteGrafanaWorkspaceAccountResponse) SetHeaders(v map[string]*string) *DeleteGrafanaWorkspaceAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponse) SetStatusCode(v int32) *DeleteGrafanaWorkspaceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponse) SetBody(v *DeleteGrafanaWorkspaceAccountResponseBody) *DeleteGrafanaWorkspaceAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
