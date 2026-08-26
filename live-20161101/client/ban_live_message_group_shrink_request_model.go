// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBanLiveMessageGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BanLiveMessageGroupShrinkRequest
	GetAppId() *string
	SetDataCenter(v string) *BanLiveMessageGroupShrinkRequest
	GetDataCenter() *string
	SetExceptUsersShrink(v string) *BanLiveMessageGroupShrinkRequest
	GetExceptUsersShrink() *string
	SetGroupId(v string) *BanLiveMessageGroupShrinkRequest
	GetGroupId() *string
}

type BanLiveMessageGroupShrinkRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. This value must be the same as the data center that you specified when you called [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html).
	//
	// > The supported data centers are China (Shanghai) (cn-shanghai) and Singapore (ap-southeast-1).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The users who are exempt from the group-wide mute. Specify up to 30 users. Separate multiple user IDs with a comma (,).
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

func (s BanLiveMessageGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BanLiveMessageGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *BanLiveMessageGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *BanLiveMessageGroupShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *BanLiveMessageGroupShrinkRequest) GetExceptUsersShrink() *string {
	return s.ExceptUsersShrink
}

func (s *BanLiveMessageGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BanLiveMessageGroupShrinkRequest) SetAppId(v string) *BanLiveMessageGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *BanLiveMessageGroupShrinkRequest) SetDataCenter(v string) *BanLiveMessageGroupShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *BanLiveMessageGroupShrinkRequest) SetExceptUsersShrink(v string) *BanLiveMessageGroupShrinkRequest {
	s.ExceptUsersShrink = &v
	return s
}

func (s *BanLiveMessageGroupShrinkRequest) SetGroupId(v string) *BanLiveMessageGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *BanLiveMessageGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
