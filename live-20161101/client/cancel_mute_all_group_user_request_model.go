// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteAllGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CancelMuteAllGroupUserRequest
	GetAppId() *string
	SetBroadCastType(v int32) *CancelMuteAllGroupUserRequest
	GetBroadCastType() *int32
	SetGroupId(v string) *CancelMuteAllGroupUserRequest
	GetGroupId() *string
	SetOperatorUserId(v string) *CancelMuteAllGroupUserRequest
	GetOperatorUserId() *string
}

type CancelMuteAllGroupUserRequest struct {
	// Interactive Messages application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// System message diffusion type. Valid values:
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
	// Group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// User ID of the operator.
	//
	// > This parameter is required. The user must be the creator of the group.
	//
	// example:
	//
	// de1**a0
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
}

func (s CancelMuteAllGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteAllGroupUserRequest) GoString() string {
	return s.String()
}

func (s *CancelMuteAllGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *CancelMuteAllGroupUserRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *CancelMuteAllGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CancelMuteAllGroupUserRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *CancelMuteAllGroupUserRequest) SetAppId(v string) *CancelMuteAllGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *CancelMuteAllGroupUserRequest) SetBroadCastType(v int32) *CancelMuteAllGroupUserRequest {
	s.BroadCastType = &v
	return s
}

func (s *CancelMuteAllGroupUserRequest) SetGroupId(v string) *CancelMuteAllGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *CancelMuteAllGroupUserRequest) SetOperatorUserId(v string) *CancelMuteAllGroupUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *CancelMuteAllGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
