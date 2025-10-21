// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Member) *UpdateMemberRequest
	GetBody() *Member
}

type UpdateMemberRequest struct {
	// The permission information about the member.
	Body *Member `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberRequest) GetBody() *Member {
	return s.Body
}

func (s *UpdateMemberRequest) SetBody(v *Member) *UpdateMemberRequest {
	s.Body = v
	return s
}

func (s *UpdateMemberRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
