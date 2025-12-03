// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGroupMemberResponseBody) *UpdateGroupMemberResponse
	GetBody() *UpdateGroupMemberResponseBody
}

type UpdateGroupMemberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGroupMemberResponse) GetBody() *UpdateGroupMemberResponseBody {
	return s.Body
}

func (s *UpdateGroupMemberResponse) SetHeaders(v map[string]*string) *UpdateGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupMemberResponse) SetStatusCode(v int32) *UpdateGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupMemberResponse) SetBody(v *UpdateGroupMemberResponseBody) *UpdateGroupMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
