// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserToDataAgentWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserToDataAgentWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserToDataAgentWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserToDataAgentWorkspaceResponseBody) *RemoveUserToDataAgentWorkspaceResponse
	GetBody() *RemoveUserToDataAgentWorkspaceResponseBody
}

type RemoveUserToDataAgentWorkspaceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserToDataAgentWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserToDataAgentWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserToDataAgentWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserToDataAgentWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserToDataAgentWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserToDataAgentWorkspaceResponse) GetBody() *RemoveUserToDataAgentWorkspaceResponseBody {
	return s.Body
}

func (s *RemoveUserToDataAgentWorkspaceResponse) SetHeaders(v map[string]*string) *RemoveUserToDataAgentWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponse) SetStatusCode(v int32) *RemoveUserToDataAgentWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponse) SetBody(v *RemoveUserToDataAgentWorkspaceResponseBody) *RemoveUserToDataAgentWorkspaceResponse {
	s.Body = v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
