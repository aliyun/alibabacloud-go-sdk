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
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application.
	//
	// >  Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The users whom you do not want to mute when the group is muted.
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
