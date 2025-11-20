// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkspaceDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkspaceDocResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkspaceDocResponseBody) *CreateWorkspaceDocResponse
	GetBody() *CreateWorkspaceDocResponseBody
}

type CreateWorkspaceDocResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkspaceDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkspaceDocResponse) GetBody() *CreateWorkspaceDocResponseBody {
	return s.Body
}

func (s *CreateWorkspaceDocResponse) SetHeaders(v map[string]*string) *CreateWorkspaceDocResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceDocResponse) SetStatusCode(v int32) *CreateWorkspaceDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceDocResponse) SetBody(v *CreateWorkspaceDocResponseBody) *CreateWorkspaceDocResponse {
	s.Body = v
	return s
}

func (s *CreateWorkspaceDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
