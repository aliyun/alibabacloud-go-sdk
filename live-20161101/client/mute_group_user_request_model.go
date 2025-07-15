// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MuteGroupUserRequest
	GetAppId() *string
	SetBroadCastType(v int32) *MuteGroupUserRequest
	GetBroadCastType() *int32
	SetGroupId(v string) *MuteGroupUserRequest
	GetGroupId() *string
	SetMuteTime(v int32) *MuteGroupUserRequest
	GetMuteTime() *int32
	SetMuteUserList(v []*string) *MuteGroupUserRequest
	GetMuteUserList() []*string
	SetOperatorUserId(v string) *MuteGroupUserRequest
	GetOperatorUserId() *string
}

type MuteGroupUserRequest struct {
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
	MuteUserList []*string `json:"MuteUserList,omitempty" xml:"MuteUserList,omitempty" type:"Repeated"`
	// The ID of the user who performs the operation.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
}

func (s MuteGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteGroupUserRequest) GoString() string {
	return s.String()
}

func (s *MuteGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *MuteGroupUserRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *MuteGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *MuteGroupUserRequest) GetMuteTime() *int32 {
	return s.MuteTime
}

func (s *MuteGroupUserRequest) GetMuteUserList() []*string {
	return s.MuteUserList
}

func (s *MuteGroupUserRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *MuteGroupUserRequest) SetAppId(v string) *MuteGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *MuteGroupUserRequest) SetBroadCastType(v int32) *MuteGroupUserRequest {
	s.BroadCastType = &v
	return s
}

func (s *MuteGroupUserRequest) SetGroupId(v string) *MuteGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *MuteGroupUserRequest) SetMuteTime(v int32) *MuteGroupUserRequest {
	s.MuteTime = &v
	return s
}

func (s *MuteGroupUserRequest) SetMuteUserList(v []*string) *MuteGroupUserRequest {
	s.MuteUserList = v
	return s
}

func (s *MuteGroupUserRequest) SetOperatorUserId(v string) *MuteGroupUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *MuteGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
