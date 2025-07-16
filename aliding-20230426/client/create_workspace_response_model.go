// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse
	GetBody() *CreateWorkspaceResponseBody
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkspaceResponse) GetBody() *CreateWorkspaceResponseBody {
	return s.Body
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreateWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
