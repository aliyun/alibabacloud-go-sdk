// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkspaceUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkspaceUsersResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkspaceUsersResponseBody) *AddWorkspaceUsersResponse
	GetBody() *AddWorkspaceUsersResponseBody
}

type AddWorkspaceUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkspaceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkspaceUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUsersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkspaceUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkspaceUsersResponse) GetBody() *AddWorkspaceUsersResponseBody {
	return s.Body
}

func (s *AddWorkspaceUsersResponse) SetHeaders(v map[string]*string) *AddWorkspaceUsersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceUsersResponse) SetStatusCode(v int32) *AddWorkspaceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceUsersResponse) SetBody(v *AddWorkspaceUsersResponseBody) *AddWorkspaceUsersResponse {
	s.Body = v
	return s
}

func (s *AddWorkspaceUsersResponse) Validate() error {
	return dara.Validate(s)
}
