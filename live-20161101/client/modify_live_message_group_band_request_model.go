// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupBandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyLiveMessageGroupBandRequest
	GetAppId() *string
	SetBannedAll(v bool) *ModifyLiveMessageGroupBandRequest
	GetBannedAll() *bool
	SetBannnedUsers(v []*string) *ModifyLiveMessageGroupBandRequest
	GetBannnedUsers() []*string
	SetDataCenter(v string) *ModifyLiveMessageGroupBandRequest
	GetDataCenter() *string
	SetExceptUsers(v []*string) *ModifyLiveMessageGroupBandRequest
	GetExceptUsers() []*string
	SetGroupId(v string) *ModifyLiveMessageGroupBandRequest
	GetGroupId() *string
}

type ModifyLiveMessageGroupBandRequest struct {
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
	BannnedUsers []*string `json:"BannnedUsers,omitempty" xml:"BannnedUsers,omitempty" type:"Repeated"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the user whom you do not want to mute when you set the BannedAll parameter to true. Separate multiple user IDs with commas (,). You can specify up to 30 users IDs.
	ExceptUsers []*string `json:"ExceptUsers,omitempty" xml:"ExceptUsers,omitempty" type:"Repeated"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ModifyLiveMessageGroupBandRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupBandRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupBandRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageGroupBandRequest) GetBannedAll() *bool {
	return s.BannedAll
}

func (s *ModifyLiveMessageGroupBandRequest) GetBannnedUsers() []*string {
	return s.BannnedUsers
}

func (s *ModifyLiveMessageGroupBandRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageGroupBandRequest) GetExceptUsers() []*string {
	return s.ExceptUsers
}

func (s *ModifyLiveMessageGroupBandRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLiveMessageGroupBandRequest) SetAppId(v string) *ModifyLiveMessageGroupBandRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageGroupBandRequest) SetBannedAll(v bool) *ModifyLiveMessageGroupBandRequest {
	s.BannedAll = &v
	return s
}

func (s *ModifyLiveMessageGroupBandRequest) SetBannnedUsers(v []*string) *ModifyLiveMessageGroupBandRequest {
	s.BannnedUsers = v
	return s
}

func (s *ModifyLiveMessageGroupBandRequest) SetDataCenter(v string) *ModifyLiveMessageGroupBandRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageGroupBandRequest) SetExceptUsers(v []*string) *ModifyLiveMessageGroupBandRequest {
	s.ExceptUsers = v
	return s
}

func (s *ModifyLiveMessageGroupBandRequest) SetGroupId(v string) *ModifyLiveMessageGroupBandRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyLiveMessageGroupBandRequest) Validate() error {
	return dara.Validate(s)
}
