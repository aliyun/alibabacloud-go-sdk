// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceResponseBody) *UpdateWorkspaceResponse
	GetBody() *UpdateWorkspaceResponseBody
}

type UpdateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceResponse) GetBody() *UpdateWorkspaceResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResponse) SetStatusCode(v int32) *UpdateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResponse) SetBody(v *UpdateWorkspaceResponseBody) *UpdateWorkspaceResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
