// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *LeaveMessageGroupRequest
	GetAppId() *string
	SetBroadCastStatistics(v bool) *LeaveMessageGroupRequest
	GetBroadCastStatistics() *bool
	SetBroadCastType(v int32) *LeaveMessageGroupRequest
	GetBroadCastType() *int32
	SetGroupId(v string) *LeaveMessageGroupRequest
	GetGroupId() *string
	SetUserId(v string) *LeaveMessageGroupRequest
	GetUserId() *string
}

type LeaveMessageGroupRequest struct {
	// Interactive message application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to broadcast statistics messages. When enabled, the statistics information of the message group will be broadcast after joining the message group, and the client can receive and process this message. Valid values:
	//
	// - true: Broadcast statistics messages.
	//
	// - false: Do not broadcast statistics messages.
	//
	// example:
	//
	// true
	BroadCastStatistics *bool `json:"BroadCastStatistics,omitempty" xml:"BroadCastStatistics,omitempty"`
	// System message diffusion type. Valid values:
	//
	// - 0 (default): No diffusion.
	//
	// - 1: Diffusion to specified users.
	//
	// - 2: Diffusion to the group.
	//
	// example:
	//
	// 2
	BroadCastType *int32 `json:"BroadCastType,omitempty" xml:"BroadCastType,omitempty"`
	// Message group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// User ID, which is customized by the user and must be unique under the AppId.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LeaveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s LeaveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *LeaveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *LeaveMessageGroupRequest) GetBroadCastStatistics() *bool {
	return s.BroadCastStatistics
}

func (s *LeaveMessageGroupRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *LeaveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *LeaveMessageGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *LeaveMessageGroupRequest) SetAppId(v string) *LeaveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *LeaveMessageGroupRequest) SetBroadCastStatistics(v bool) *LeaveMessageGroupRequest {
	s.BroadCastStatistics = &v
	return s
}

func (s *LeaveMessageGroupRequest) SetBroadCastType(v int32) *LeaveMessageGroupRequest {
	s.BroadCastType = &v
	return s
}

func (s *LeaveMessageGroupRequest) SetGroupId(v string) *LeaveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *LeaveMessageGroupRequest) SetUserId(v string) *LeaveMessageGroupRequest {
	s.UserId = &v
	return s
}

func (s *LeaveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
