// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserIsGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckUserIsGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckUserIsGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *CheckUserIsGroupMemberResponseBody) *CheckUserIsGroupMemberResponse
	GetBody() *CheckUserIsGroupMemberResponseBody
}

type CheckUserIsGroupMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserIsGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserIsGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckUserIsGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckUserIsGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckUserIsGroupMemberResponse) GetBody() *CheckUserIsGroupMemberResponseBody {
	return s.Body
}

func (s *CheckUserIsGroupMemberResponse) SetHeaders(v map[string]*string) *CheckUserIsGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *CheckUserIsGroupMemberResponse) SetStatusCode(v int32) *CheckUserIsGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserIsGroupMemberResponse) SetBody(v *CheckUserIsGroupMemberResponseBody) *CheckUserIsGroupMemberResponse {
	s.Body = v
	return s
}

func (s *CheckUserIsGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
