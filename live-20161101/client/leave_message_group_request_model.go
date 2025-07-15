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
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to broadcast statistical messages. If you set the value to true, statistical messages of the message group are broadcasted after the users join the message group. The client can receive and process these messages. Valid values:
	//
	// 	- true: broadcasts statistical messages.
	//
	// 	- false: does not broadcast statistical messages.
	//
	// example:
	//
	// true
	BroadCastStatistics *bool `json:"BroadCastStatistics,omitempty" xml:"BroadCastStatistics,omitempty"`
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
	// The ID of the user. Each user has a unique ID in the application.
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
