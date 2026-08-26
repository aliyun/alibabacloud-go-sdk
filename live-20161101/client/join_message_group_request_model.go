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
	// Interactive message application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a494caec-***-695ef345db77
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to broadcast statistics messages. When enabled, statistics information of the message group will be broadcast after joining the message group, and the client can receive and process this message. Valid values:
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
	// The ID of the message group to join. Make sure the GroupId you provide exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// User ID, which is customized by the user and must be unique under the AppId. It can contain lowercase letters, numbers, underscores (_), and periods (.). The maximum length is 32 characters. Different users must use different UserIds.
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
