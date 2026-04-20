// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveWorkspaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveWorkspaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveWorkspaceUserResponse
	GetStatusCode() *int32
	SetBody(v *RemoveWorkspaceUserResponseBody) *RemoveWorkspaceUserResponse
	GetBody() *RemoveWorkspaceUserResponseBody
}

type RemoveWorkspaceUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveWorkspaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveWorkspaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveWorkspaceUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveWorkspaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveWorkspaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveWorkspaceUserResponse) GetBody() *RemoveWorkspaceUserResponseBody {
	return s.Body
}

func (s *RemoveWorkspaceUserResponse) SetHeaders(v map[string]*string) *RemoveWorkspaceUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveWorkspaceUserResponse) SetStatusCode(v int32) *RemoveWorkspaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWorkspaceUserResponse) SetBody(v *RemoveWorkspaceUserResponseBody) *RemoveWorkspaceUserResponse {
	s.Body = v
	return s
}

func (s *RemoveWorkspaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
