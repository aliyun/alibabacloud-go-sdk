// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterBucUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RegisterBucUserCmd) *RegisterBucUserRequest
	GetBody() *RegisterBucUserCmd
}

type RegisterBucUserRequest struct {
	Body *RegisterBucUserCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterBucUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterBucUserRequest) GoString() string {
	return s.String()
}

func (s *RegisterBucUserRequest) GetBody() *RegisterBucUserCmd {
	return s.Body
}

func (s *RegisterBucUserRequest) SetBody(v *RegisterBucUserCmd) *RegisterBucUserRequest {
	s.Body = v
	return s
}

func (s *RegisterBucUserRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
