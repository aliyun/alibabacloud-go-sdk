// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddWorkspaceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddWorkspaceMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddWorkspaceMembersResponseBody) *AddWorkspaceMembersResponse
	GetBody() *AddWorkspaceMembersResponseBody
}

type AddWorkspaceMembersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkspaceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkspaceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddWorkspaceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddWorkspaceMembersResponse) GetBody() *AddWorkspaceMembersResponseBody {
	return s.Body
}

func (s *AddWorkspaceMembersResponse) SetHeaders(v map[string]*string) *AddWorkspaceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceMembersResponse) SetStatusCode(v int32) *AddWorkspaceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkspaceMembersResponse) SetBody(v *AddWorkspaceMembersResponseBody) *AddWorkspaceMembersResponse {
	s.Body = v
	return s
}

func (s *AddWorkspaceMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
