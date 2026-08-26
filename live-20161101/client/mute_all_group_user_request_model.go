// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MuteAllGroupUserRequest
	GetAppId() *string
	SetBroadCastType(v int32) *MuteAllGroupUserRequest
	GetBroadCastType() *int32
	SetGroupId(v string) *MuteAllGroupUserRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *MuteAllGroupUserRequest
	GetOperatorUserId() *string
}

type MuteAllGroupUserRequest struct {
	// The application ID for interactive messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The system message diffusion type. Valid values:
	//
	// - 0: No diffusion.
	//
	// - 1: Diffusion to specified users.
	//
	// - 2: Diffusion to the group.
	//
	// example:
	//
	// 2
	BroadCastType *int32 `json:"BroadCastType,omitempty" xml:"BroadCastType,omitempty"`
	// The message group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user ID of the operator. This user must be the creator of the group.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
}

func (s MuteAllGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteAllGroupUserRequest) GoString() string {
	return s.String()
}

func (s *MuteAllGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *MuteAllGroupUserRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *MuteAllGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *MuteAllGroupUserRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *MuteAllGroupUserRequest) SetAppId(v string) *MuteAllGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *MuteAllGroupUserRequest) SetBroadCastType(v int32) *MuteAllGroupUserRequest {
	s.BroadCastType = &v
	return s
}

func (s *MuteAllGroupUserRequest) SetGroupId(v string) *MuteAllGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *MuteAllGroupUserRequest) SetOperatorUserId(v string) *MuteAllGroupUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *MuteAllGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
