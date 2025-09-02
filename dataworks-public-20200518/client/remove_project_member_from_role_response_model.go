// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberFromRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveProjectMemberFromRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveProjectMemberFromRoleResponse
	GetStatusCode() *int32
	SetBody(v *RemoveProjectMemberFromRoleResponseBody) *RemoveProjectMemberFromRoleResponse
	GetBody() *RemoveProjectMemberFromRoleResponseBody
}

type RemoveProjectMemberFromRoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveProjectMemberFromRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveProjectMemberFromRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberFromRoleResponse) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberFromRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveProjectMemberFromRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveProjectMemberFromRoleResponse) GetBody() *RemoveProjectMemberFromRoleResponseBody {
	return s.Body
}

func (s *RemoveProjectMemberFromRoleResponse) SetHeaders(v map[string]*string) *RemoveProjectMemberFromRoleResponse {
	s.Headers = v
	return s
}

func (s *RemoveProjectMemberFromRoleResponse) SetStatusCode(v int32) *RemoveProjectMemberFromRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveProjectMemberFromRoleResponse) SetBody(v *RemoveProjectMemberFromRoleResponseBody) *RemoveProjectMemberFromRoleResponse {
	s.Body = v
	return s
}

func (s *RemoveProjectMemberFromRoleResponse) Validate() error {
	return dara.Validate(s)
}
