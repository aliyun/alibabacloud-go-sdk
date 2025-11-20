// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceMembersResponseBody) *DeleteWorkspaceMembersResponse
	GetBody() *DeleteWorkspaceMembersResponseBody
}

type DeleteWorkspaceMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceMembersResponse) GetBody() *DeleteWorkspaceMembersResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceMembersResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceMembersResponse) SetStatusCode(v int32) *DeleteWorkspaceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceMembersResponse) SetBody(v *DeleteWorkspaceMembersResponseBody) *DeleteWorkspaceMembersResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
