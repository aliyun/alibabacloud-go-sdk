// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteGroupUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MuteGroupUserShrinkRequest
	GetAppId() *string
	SetBroadCastType(v int32) *MuteGroupUserShrinkRequest
	GetBroadCastType() *int32
	SetGroupId(v string) *MuteGroupUserShrinkRequest
	GetGroupId() *string
	SetMuteTime(v int32) *MuteGroupUserShrinkRequest
	GetMuteTime() *int32
	SetMuteUserListShrink(v string) *MuteGroupUserShrinkRequest
	GetMuteUserListShrink() *string
	SetOperatorUserId(v string) *MuteGroupUserShrinkRequest
	GetOperatorUserId() *string
}

type MuteGroupUserShrinkRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The mode in which system messages are broadcasted. Valid values:
	//
	// 	- 0: specifies that system messages are not broadcasted. This is the default value.
	//
	// 	- 1: specifies that system messages are broadcasted to specified users.
	//
	// 	- 2: specifies that system messages are broadcasted to the message group.
	//
	// example:
	//
	// 2
	BroadCastType *int32 `json:"BroadCastType,omitempty" xml:"BroadCastType,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The duration of the mute. Unit: seconds.
	//
	// > If you do not specify this parameter or set the value to 0, the default duration of 86,400 seconds is used.
	//
	// example:
	//
	// 3600
	MuteTime *int32 `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	// Details about the mute.
	//
	// This parameter is required.
	MuteUserListShrink *string `json:"MuteUserList,omitempty" xml:"MuteUserList,omitempty"`
	// The ID of the user who performs the operation.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
}

func (s MuteGroupUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteGroupUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *MuteGroupUserShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *MuteGroupUserShrinkRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *MuteGroupUserShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *MuteGroupUserShrinkRequest) GetMuteTime() *int32 {
	return s.MuteTime
}

func (s *MuteGroupUserShrinkRequest) GetMuteUserListShrink() *string {
	return s.MuteUserListShrink
}

func (s *MuteGroupUserShrinkRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *MuteGroupUserShrinkRequest) SetAppId(v string) *MuteGroupUserShrinkRequest {
	s.AppId = &v
	return s
}

func (s *MuteGroupUserShrinkRequest) SetBroadCastType(v int32) *MuteGroupUserShrinkRequest {
	s.BroadCastType = &v
	return s
}

func (s *MuteGroupUserShrinkRequest) SetGroupId(v string) *MuteGroupUserShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *MuteGroupUserShrinkRequest) SetMuteTime(v int32) *MuteGroupUserShrinkRequest {
	s.MuteTime = &v
	return s
}

func (s *MuteGroupUserShrinkRequest) SetMuteUserListShrink(v string) *MuteGroupUserShrinkRequest {
	s.MuteUserListShrink = &v
	return s
}

func (s *MuteGroupUserShrinkRequest) SetOperatorUserId(v string) *MuteGroupUserShrinkRequest {
	s.OperatorUserId = &v
	return s
}

func (s *MuteGroupUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
