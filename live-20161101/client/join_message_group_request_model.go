// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *JoinMessageGroupRequest
	GetAppId() *string
	SetBroadCastStatistics(v bool) *JoinMessageGroupRequest
	GetBroadCastStatistics() *bool
	SetBroadCastType(v int32) *JoinMessageGroupRequest
	GetBroadCastType() *int32
	SetGroupId(v string) *JoinMessageGroupRequest
	GetGroupId() *string
	SetUserId(v string) *JoinMessageGroupRequest
	GetUserId() *string
}

type JoinMessageGroupRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
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
	// The ID of the message group to join.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user. Each user has a unique ID in the application. The ID can be up to 32 characters in length and can contain lowercase letters, digits, underscores (_), and periods (.). You can specify multiple user IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// de1**a0
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s JoinMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *JoinMessageGroupRequest) GetBroadCastStatistics() *bool {
	return s.BroadCastStatistics
}

func (s *JoinMessageGroupRequest) GetBroadCastType() *int32 {
	return s.BroadCastType
}

func (s *JoinMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *JoinMessageGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *JoinMessageGroupRequest) SetAppId(v string) *JoinMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *JoinMessageGroupRequest) SetBroadCastStatistics(v bool) *JoinMessageGroupRequest {
	s.BroadCastStatistics = &v
	return s
}

func (s *JoinMessageGroupRequest) SetBroadCastType(v int32) *JoinMessageGroupRequest {
	s.BroadCastType = &v
	return s
}

func (s *JoinMessageGroupRequest) SetGroupId(v string) *JoinMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *JoinMessageGroupRequest) SetUserId(v string) *JoinMessageGroupRequest {
	s.UserId = &v
	return s
}

func (s *JoinMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
