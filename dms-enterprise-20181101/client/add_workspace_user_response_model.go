// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkspaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkspaceUserResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkspaceUserResponseBody) *AddWorkspaceUserResponse
	GetBody() *AddWorkspaceUserResponseBody
}

type AddWorkspaceUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkspaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkspaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUserResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkspaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkspaceUserResponse) GetBody() *AddWorkspaceUserResponseBody {
	return s.Body
}

func (s *AddWorkspaceUserResponse) SetHeaders(v map[string]*string) *AddWorkspaceUserResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceUserResponse) SetStatusCode(v int32) *AddWorkspaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceUserResponse) SetBody(v *AddWorkspaceUserResponseBody) *AddWorkspaceUserResponse {
	s.Body = v
	return s
}

func (s *AddWorkspaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
