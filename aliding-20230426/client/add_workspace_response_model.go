// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkspaceResponseBody) *AddWorkspaceResponse
	GetBody() *AddWorkspaceResponseBody
}

type AddWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkspaceResponse) GetBody() *AddWorkspaceResponseBody {
	return s.Body
}

func (s *AddWorkspaceResponse) SetHeaders(v map[string]*string) *AddWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceResponse) SetStatusCode(v int32) *AddWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceResponse) SetBody(v *AddWorkspaceResponseBody) *AddWorkspaceResponse {
	s.Body = v
	return s
}

func (s *AddWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
