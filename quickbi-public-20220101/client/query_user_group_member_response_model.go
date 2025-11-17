// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserGroupMemberResponseBody) *QueryUserGroupMemberResponse
	GetBody() *QueryUserGroupMemberResponseBody
}

type QueryUserGroupMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryUserGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserGroupMemberResponse) GetBody() *QueryUserGroupMemberResponseBody {
	return s.Body
}

func (s *QueryUserGroupMemberResponse) SetHeaders(v map[string]*string) *QueryUserGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryUserGroupMemberResponse) SetStatusCode(v int32) *QueryUserGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserGroupMemberResponse) SetBody(v *QueryUserGroupMemberResponseBody) *QueryUserGroupMemberResponse {
	s.Body = v
	return s
}

func (s *QueryUserGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
