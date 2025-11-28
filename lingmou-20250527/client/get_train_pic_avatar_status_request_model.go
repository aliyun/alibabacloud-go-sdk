// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainPicAvatarStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarId(v string) *GetTrainPicAvatarStatusRequest
	GetAvatarId() *string
}

type GetTrainPicAvatarStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// M1YJTNTH2yoLmLnzKdYHeGBg
	AvatarId *string `json:"avatarId,omitempty" xml:"avatarId,omitempty"`
}

func (s GetTrainPicAvatarStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrainPicAvatarStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTrainPicAvatarStatusRequest) GetAvatarId() *string {
	return s.AvatarId
}

func (s *GetTrainPicAvatarStatusRequest) SetAvatarId(v string) *GetTrainPicAvatarStatusRequest {
	s.AvatarId = &v
	return s
}

func (s *GetTrainPicAvatarStatusRequest) Validate() error {
	return dara.Validate(s)
}
