// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceDocMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceDocMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceDocMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceDocMembersResponseBody) *DeleteWorkspaceDocMembersResponse
	GetBody() *DeleteWorkspaceDocMembersResponseBody
}

type DeleteWorkspaceDocMembersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceDocMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceDocMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceDocMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceDocMembersResponse) GetBody() *DeleteWorkspaceDocMembersResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceDocMembersResponse) SetStatusCode(v int32) *DeleteWorkspaceDocMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceDocMembersResponse) SetBody(v *DeleteWorkspaceDocMembersResponseBody) *DeleteWorkspaceDocMembersResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceDocMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
