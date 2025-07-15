// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CancelMuteGroupUserRequest
	GetAppId() *string
	SetBroadCastType(v int32) *CancelMuteGroupUserRequest
	GetBroadCastType() *int32
	SetCancelMuteUserList(v []*string) *CancelMuteGroupUserRequest
	GetCancelMuteUserList() []*string
	SetGroupId(v string) *CancelMuteGroupUserRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *CancelMuteGroupUserRequest
	GetOperatorUserId() *string
}

type CancelMuteGroupUserRequest struct {
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
	CancelMuteUserList []*string `json:"CancelMuteUserList,omitempty" xml:"CancelMuteUserList,omitempty" type:"Repeated"`
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

func (s CancelMuteGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteGroupUserRequest) GoString() string {
	return s.String()
}

func (s *CancelMuteGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *CancelMuteGroupUserRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *CancelMuteGroupUserRequest) GetCancelMuteUserList() []*string {
	return s.CancelMuteUserList
}

func (s *CancelMuteGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CancelMuteGroupUserRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *CancelMuteGroupUserRequest) SetAppId(v string) *CancelMuteGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *CancelMuteGroupUserRequest) SetBroadCastType(v int32) *CancelMuteGroupUserRequest {
	s.BroadCastType = &v
	return s
}

func (s *CancelMuteGroupUserRequest) SetCancelMuteUserList(v []*string) *CancelMuteGroupUserRequest {
	s.CancelMuteUserList = v
	return s
}

func (s *CancelMuteGroupUserRequest) SetGroupId(v string) *CancelMuteGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *CancelMuteGroupUserRequest) SetOperatorUserId(v string) *CancelMuteGroupUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *CancelMuteGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
