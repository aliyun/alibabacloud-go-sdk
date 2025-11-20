// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceDocMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkspaceDocMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkspaceDocMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkspaceDocMembersResponseBody) *AddWorkspaceDocMembersResponse
	GetBody() *AddWorkspaceDocMembersResponseBody
}

type AddWorkspaceDocMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkspaceDocMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkspaceDocMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkspaceDocMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkspaceDocMembersResponse) GetBody() *AddWorkspaceDocMembersResponseBody {
	return s.Body
}

func (s *AddWorkspaceDocMembersResponse) SetHeaders(v map[string]*string) *AddWorkspaceDocMembersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceDocMembersResponse) SetStatusCode(v int32) *AddWorkspaceDocMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceDocMembersResponse) SetBody(v *AddWorkspaceDocMembersResponseBody) *AddWorkspaceDocMembersResponse {
	s.Body = v
	return s
}

func (s *AddWorkspaceDocMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
