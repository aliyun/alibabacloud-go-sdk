// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceUserResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceUserResponseBody) *UpdateWorkspaceUserResponse
	GetBody() *UpdateWorkspaceUserResponseBody
}

type UpdateWorkspaceUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceUserResponse) GetBody() *UpdateWorkspaceUserResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceUserResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceUserResponse) SetStatusCode(v int32) *UpdateWorkspaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceUserResponse) SetBody(v *UpdateWorkspaceUserResponseBody) *UpdateWorkspaceUserResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
