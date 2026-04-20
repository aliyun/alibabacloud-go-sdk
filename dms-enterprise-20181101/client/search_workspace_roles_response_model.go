// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchWorkspaceRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchWorkspaceRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchWorkspaceRolesResponse
	GetStatusCode() *int32
	SetBody(v *SearchWorkspaceRolesResponseBody) *SearchWorkspaceRolesResponse
	GetBody() *SearchWorkspaceRolesResponseBody
}

type SearchWorkspaceRolesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchWorkspaceRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchWorkspaceRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchWorkspaceRolesResponse) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchWorkspaceRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchWorkspaceRolesResponse) GetBody() *SearchWorkspaceRolesResponseBody {
	return s.Body
}

func (s *SearchWorkspaceRolesResponse) SetHeaders(v map[string]*string) *SearchWorkspaceRolesResponse {
	s.Headers = v
	return s
}

func (s *SearchWorkspaceRolesResponse) SetStatusCode(v int32) *SearchWorkspaceRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchWorkspaceRolesResponse) SetBody(v *SearchWorkspaceRolesResponseBody) *SearchWorkspaceRolesResponse {
	s.Body = v
	return s
}

func (s *SearchWorkspaceRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
