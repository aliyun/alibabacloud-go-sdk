// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMemberRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveMemberRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveMemberRoleResponse
	GetStatusCode() *int32
	SetBody(v *RemoveMemberRoleResponseBody) *RemoveMemberRoleResponse
	GetBody() *RemoveMemberRoleResponseBody
}

type RemoveMemberRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMemberRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *RemoveMemberRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveMemberRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveMemberRoleResponse) GetBody() *RemoveMemberRoleResponseBody {
	return s.Body
}

func (s *RemoveMemberRoleResponse) SetHeaders(v map[string]*string) *RemoveMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *RemoveMemberRoleResponse) SetStatusCode(v int32) *RemoveMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMemberRoleResponse) SetBody(v *RemoveMemberRoleResponseBody) *RemoveMemberRoleResponse {
	s.Body = v
	return s
}

func (s *RemoveMemberRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
