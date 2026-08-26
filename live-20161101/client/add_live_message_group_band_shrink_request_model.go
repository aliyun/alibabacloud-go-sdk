// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveMessageGroupBandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddLiveMessageGroupBandShrinkRequest
	GetAppId() *string
	SetBannedUsersShrink(v string) *AddLiveMessageGroupBandShrinkRequest
	GetBannedUsersShrink() *string
	SetDataCenter(v string) *AddLiveMessageGroupBandShrinkRequest
	GetDataCenter() *string
	SetGroupId(v string) *AddLiveMessageGroupBandShrinkRequest
	GetGroupId() *string
}

type AddLiveMessageGroupBandShrinkRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// A list of users to mute. Separate multiple user IDs with a comma (,). You can specify a maximum of 30 users.
	//
	// This parameter is required.
	BannedUsersShrink *string `json:"BannedUsers,omitempty" xml:"BannedUsers,omitempty"`
	// The data center. This must be the same data center that you specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation.
	//
	// > Currently, this operation is supported in Shanghai (value: cn-shanghai) and Singapore (value: ap-southeast-1).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s AddLiveMessageGroupBandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveMessageGroupBandShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddLiveMessageGroupBandShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddLiveMessageGroupBandShrinkRequest) GetBannedUsersShrink() *string {
	return s.BannedUsersShrink
}

func (s *AddLiveMessageGroupBandShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *AddLiveMessageGroupBandShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddLiveMessageGroupBandShrinkRequest) SetAppId(v string) *AddLiveMessageGroupBandShrinkRequest {
	s.AppId = &v
	return s
}

func (s *AddLiveMessageGroupBandShrinkRequest) SetBannedUsersShrink(v string) *AddLiveMessageGroupBandShrinkRequest {
	s.BannedUsersShrink = &v
	return s
}

func (s *AddLiveMessageGroupBandShrinkRequest) SetDataCenter(v string) *AddLiveMessageGroupBandShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *AddLiveMessageGroupBandShrinkRequest) SetGroupId(v string) *AddLiveMessageGroupBandShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *AddLiveMessageGroupBandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
