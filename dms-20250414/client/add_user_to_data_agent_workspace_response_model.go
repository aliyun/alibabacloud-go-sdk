// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDataAgentWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToDataAgentWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToDataAgentWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToDataAgentWorkspaceResponseBody) *AddUserToDataAgentWorkspaceResponse
	GetBody() *AddUserToDataAgentWorkspaceResponseBody
}

type AddUserToDataAgentWorkspaceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToDataAgentWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToDataAgentWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDataAgentWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *AddUserToDataAgentWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToDataAgentWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToDataAgentWorkspaceResponse) GetBody() *AddUserToDataAgentWorkspaceResponseBody {
	return s.Body
}

func (s *AddUserToDataAgentWorkspaceResponse) SetHeaders(v map[string]*string) *AddUserToDataAgentWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponse) SetStatusCode(v int32) *AddUserToDataAgentWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponse) SetBody(v *AddUserToDataAgentWorkspaceResponseBody) *AddUserToDataAgentWorkspaceResponse {
	s.Body = v
	return s
}

func (s *AddUserToDataAgentWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
