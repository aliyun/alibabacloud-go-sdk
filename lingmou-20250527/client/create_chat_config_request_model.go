// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarId(v string) *CreateChatConfigRequest
	GetAvatarId() *string
	SetBackgroundId(v string) *CreateChatConfigRequest
	GetBackgroundId() *string
}

type CreateChatConfigRequest struct {
	// example:
	//
	// M1ONzwuILu-nPT7pvr6maKvQ
	AvatarId *string `json:"avatarId,omitempty" xml:"avatarId,omitempty"`
	// example:
	//
	// M1ONzwuILu-nPT7pvr6maKvQ
	BackgroundId *string `json:"backgroundId,omitempty" xml:"backgroundId,omitempty"`
}

func (s CreateChatConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateChatConfigRequest) GetAvatarId() *string {
	return s.AvatarId
}

func (s *CreateChatConfigRequest) GetBackgroundId() *string {
	return s.BackgroundId
}

func (s *CreateChatConfigRequest) SetAvatarId(v string) *CreateChatConfigRequest {
	s.AvatarId = &v
	return s
}

func (s *CreateChatConfigRequest) SetBackgroundId(v string) *CreateChatConfigRequest {
	s.BackgroundId = &v
	return s
}

func (s *CreateChatConfigRequest) Validate() error {
	return dara.Validate(s)
}
