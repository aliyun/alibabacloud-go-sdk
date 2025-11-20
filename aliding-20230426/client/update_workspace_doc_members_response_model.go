// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDocMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceDocMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceDocMembersResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceDocMembersResponseBody) *UpdateWorkspaceDocMembersResponse
	GetBody() *UpdateWorkspaceDocMembersResponseBody
}

type UpdateWorkspaceDocMembersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceDocMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceDocMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceDocMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceDocMembersResponse) GetBody() *UpdateWorkspaceDocMembersResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceDocMembersResponse) SetStatusCode(v int32) *UpdateWorkspaceDocMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceDocMembersResponse) SetBody(v *UpdateWorkspaceDocMembersResponseBody) *UpdateWorkspaceDocMembersResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceDocMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
