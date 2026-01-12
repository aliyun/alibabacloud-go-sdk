// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceType(v string) *RegisterUserRequest
	GetSourceType() *string
}

type RegisterUserRequest struct {
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s RegisterUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterUserRequest) GoString() string {
	return s.String()
}

func (s *RegisterUserRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *RegisterUserRequest) SetSourceType(v string) *RegisterUserRequest {
	s.SourceType = &v
	return s
}

func (s *RegisterUserRequest) Validate() error {
	return dara.Validate(s)
}
