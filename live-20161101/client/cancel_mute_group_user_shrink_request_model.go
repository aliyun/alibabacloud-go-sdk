// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteGroupUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CancelMuteGroupUserShrinkRequest
	GetAppId() *string
	SetBroadCastType(v int32) *CancelMuteGroupUserShrinkRequest
	GetBroadCastType() *int32
	SetCancelMuteUserListShrink(v string) *CancelMuteGroupUserShrinkRequest
	GetCancelMuteUserListShrink() *string
	SetGroupId(v string) *CancelMuteGroupUserShrinkRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *CancelMuteGroupUserShrinkRequest
	GetOperatorUserId() *string
}

type CancelMuteGroupUserShrinkRequest struct {
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
	// The IDs of the users.
	//
	// This parameter is required.
	CancelMuteUserListShrink *string `json:"CancelMuteUserList,omitempty" xml:"CancelMuteUserList,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user who performs the operation.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
}

func (s CancelMuteGroupUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteGroupUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelMuteGroupUserShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CancelMuteGroupUserShrinkRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *CancelMuteGroupUserShrinkRequest) GetCancelMuteUserListShrink() *string {
	return s.CancelMuteUserListShrink
}

func (s *CancelMuteGroupUserShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CancelMuteGroupUserShrinkRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *CancelMuteGroupUserShrinkRequest) SetAppId(v string) *CancelMuteGroupUserShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CancelMuteGroupUserShrinkRequest) SetBroadCastType(v int32) *CancelMuteGroupUserShrinkRequest {
	s.BroadCastType = &v
	return s
}

func (s *CancelMuteGroupUserShrinkRequest) SetCancelMuteUserListShrink(v string) *CancelMuteGroupUserShrinkRequest {
	s.CancelMuteUserListShrink = &v
	return s
}

func (s *CancelMuteGroupUserShrinkRequest) SetGroupId(v string) *CancelMuteGroupUserShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CancelMuteGroupUserShrinkRequest) SetOperatorUserId(v string) *CancelMuteGroupUserShrinkRequest {
	s.OperatorUserId = &v
	return s
}

func (s *CancelMuteGroupUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
