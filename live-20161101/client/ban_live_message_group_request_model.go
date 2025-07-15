// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBanLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BanLiveMessageGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *BanLiveMessageGroupRequest
	GetDataCenter() *string
	SetExceptUsers(v []*string) *BanLiveMessageGroupRequest
	GetExceptUsers() []*string
	SetGroupId(v string) *BanLiveMessageGroupRequest
	GetGroupId() *string
}

type BanLiveMessageGroupRequest struct {
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

func (s BanLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s BanLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *BanLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *BanLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *BanLiveMessageGroupRequest) GetExceptUsers() []*string {
	return s.ExceptUsers
}

func (s *BanLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BanLiveMessageGroupRequest) SetAppId(v string) *BanLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *BanLiveMessageGroupRequest) SetDataCenter(v string) *BanLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *BanLiveMessageGroupRequest) SetExceptUsers(v []*string) *BanLiveMessageGroupRequest {
	s.ExceptUsers = v
	return s
}

func (s *BanLiveMessageGroupRequest) SetGroupId(v string) *BanLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *BanLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
