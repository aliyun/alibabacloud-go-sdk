// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Member) *CreateMemberRequest
	GetBody() *Member
}

type CreateMemberRequest struct {
	// The mappings between the ID and permissions of the member.
	Body *Member `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberRequest) GetBody() *Member {
	return s.Body
}

func (s *CreateMemberRequest) SetBody(v *Member) *CreateMemberRequest {
	s.Body = v
	return s
}

func (s *CreateMemberRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
