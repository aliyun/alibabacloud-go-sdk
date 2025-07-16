// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceMembersResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceMembersResponseBody) *UpdateWorkspaceMembersResponse
	GetBody() *UpdateWorkspaceMembersResponseBody
}

type UpdateWorkspaceMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceMembersResponse) GetBody() *UpdateWorkspaceMembersResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceMembersResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceMembersResponse) SetStatusCode(v int32) *UpdateWorkspaceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceMembersResponse) SetBody(v *UpdateWorkspaceMembersResponseBody) *UpdateWorkspaceMembersResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceMembersResponse) Validate() error {
	return dara.Validate(s)
}
