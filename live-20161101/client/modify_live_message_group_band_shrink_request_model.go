// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupBandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageGroupBandShrinkRequest
	GetAppId() *string
	SetBannedAll(v bool) *ModifyLiveMessageGroupBandShrinkRequest
	GetBannedAll() *bool
	SetBannnedUsersShrink(v string) *ModifyLiveMessageGroupBandShrinkRequest
	GetBannnedUsersShrink() *string
	SetDataCenter(v string) *ModifyLiveMessageGroupBandShrinkRequest
	GetDataCenter() *string
	SetExceptUsersShrink(v string) *ModifyLiveMessageGroupBandShrinkRequest
	GetExceptUsersShrink() *string
	SetGroupId(v string) *ModifyLiveMessageGroupBandShrinkRequest
	GetGroupId() *string
}

type ModifyLiveMessageGroupBandShrinkRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to mute all users.
	//
	// example:
	//
	// false
	BannedAll *bool `json:"BannedAll,omitempty" xml:"BannedAll,omitempty"`
	// The ID of the user whom you want to mute. Separate multiple user IDs with commas (,). You can specify up to 30 users IDs.
	BannnedUsersShrink *string `json:"BannnedUsers,omitempty" xml:"BannnedUsers,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the user whom you do not want to mute when you set the BannedAll parameter to true. Separate multiple user IDs with commas (,). You can specify up to 30 users IDs.
	ExceptUsersShrink *string `json:"ExceptUsers,omitempty" xml:"ExceptUsers,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ModifyLiveMessageGroupBandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupBandShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) GetBannedAll() *bool {
	return s.BannedAll
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) GetBannnedUsersShrink() *string {
	return s.BannnedUsersShrink
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) GetExceptUsersShrink() *string {
	return s.ExceptUsersShrink
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) SetAppId(v string) *ModifyLiveMessageGroupBandShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) SetBannedAll(v bool) *ModifyLiveMessageGroupBandShrinkRequest {
	s.BannedAll = &v
	return s
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) SetBannnedUsersShrink(v string) *ModifyLiveMessageGroupBandShrinkRequest {
	s.BannnedUsersShrink = &v
	return s
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) SetDataCenter(v string) *ModifyLiveMessageGroupBandShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) SetExceptUsersShrink(v string) *ModifyLiveMessageGroupBandShrinkRequest {
	s.ExceptUsersShrink = &v
	return s
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) SetGroupId(v string) *ModifyLiveMessageGroupBandShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyLiveMessageGroupBandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
