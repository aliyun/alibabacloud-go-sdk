// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGrafanaWorkspaceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGrafanaWorkspaceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGrafanaWorkspaceAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateGrafanaWorkspaceAccountResponseBody) *CreateGrafanaWorkspaceAccountResponse
	GetBody() *CreateGrafanaWorkspaceAccountResponseBody
}

type CreateGrafanaWorkspaceAccountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGrafanaWorkspaceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGrafanaWorkspaceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGrafanaWorkspaceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateGrafanaWorkspaceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGrafanaWorkspaceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGrafanaWorkspaceAccountResponse) GetBody() *CreateGrafanaWorkspaceAccountResponseBody {
	return s.Body
}

func (s *CreateGrafanaWorkspaceAccountResponse) SetHeaders(v map[string]*string) *CreateGrafanaWorkspaceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponse) SetStatusCode(v int32) *CreateGrafanaWorkspaceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponse) SetBody(v *CreateGrafanaWorkspaceAccountResponseBody) *CreateGrafanaWorkspaceAccountResponse {
	s.Body = v
	return s
}

func (s *CreateGrafanaWorkspaceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
