// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *WorkspaceResult) *UpdateWorkspaceResponse
	GetBody() *WorkspaceResult
}

type UpdateWorkspaceResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WorkspaceResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceResponse) GetBody() *WorkspaceResult {
	return s.Body
}

func (s *UpdateWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResponse) SetStatusCode(v int32) *UpdateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResponse) SetBody(v *WorkspaceResult) *UpdateWorkspaceResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
