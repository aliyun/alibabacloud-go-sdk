// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddGroupMemberResponseBody) *AddGroupMemberResponse
	GetBody() *AddGroupMemberResponseBody
}

type AddGroupMemberResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGroupMemberResponse) GetBody() *AddGroupMemberResponseBody {
	return s.Body
}

func (s *AddGroupMemberResponse) SetHeaders(v map[string]*string) *AddGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMemberResponse) SetStatusCode(v int32) *AddGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupMemberResponse) SetBody(v *AddGroupMemberResponseBody) *AddGroupMemberResponse {
	s.Body = v
	return s
}

func (s *AddGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
