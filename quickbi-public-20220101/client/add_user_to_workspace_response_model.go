// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToWorkspaceResponseBody) *AddUserToWorkspaceResponse
	GetBody() *AddUserToWorkspaceResponseBody
}

type AddUserToWorkspaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *AddUserToWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToWorkspaceResponse) GetBody() *AddUserToWorkspaceResponseBody {
	return s.Body
}

func (s *AddUserToWorkspaceResponse) SetHeaders(v map[string]*string) *AddUserToWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *AddUserToWorkspaceResponse) SetStatusCode(v int32) *AddUserToWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToWorkspaceResponse) SetBody(v *AddUserToWorkspaceResponseBody) *AddUserToWorkspaceResponse {
	s.Body = v
	return s
}

func (s *AddUserToWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
