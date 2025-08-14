// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvatarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarId(v string) *GetAvatarRequest
	GetAvatarId() *string
}

type GetAvatarRequest struct {
	// 	- The ID of the digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// Avatar-XXXX
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
}

func (s GetAvatarRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAvatarRequest) GoString() string {
	return s.String()
}

func (s *GetAvatarRequest) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetAvatarRequest) SetAvatarId(v string) *GetAvatarRequest {
	s.AvatarId = &v
	return s
}

func (s *GetAvatarRequest) Validate() error {
	return dara.Validate(s)
}
