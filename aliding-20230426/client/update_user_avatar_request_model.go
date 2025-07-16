// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAvatarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarMediaId(v string) *UpdateUserAvatarRequest
	GetAvatarMediaId() *string
}

type UpdateUserAvatarRequest struct {
	// example:
	//
	// @lADPDfYH3A-xxxx
	AvatarMediaId *string `json:"AvatarMediaId,omitempty" xml:"AvatarMediaId,omitempty"`
}

func (s UpdateUserAvatarRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAvatarRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserAvatarRequest) GetAvatarMediaId() *string {
	return s.AvatarMediaId
}

func (s *UpdateUserAvatarRequest) SetAvatarMediaId(v string) *UpdateUserAvatarRequest {
	s.AvatarMediaId = &v
	return s
}

func (s *UpdateUserAvatarRequest) Validate() error {
	return dara.Validate(s)
}
