// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveMessageGroupBandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddLiveMessageGroupBandRequest
	GetAppId() *string
	SetBannedUsers(v []*string) *AddLiveMessageGroupBandRequest
	GetBannedUsers() []*string
	SetDataCenter(v string) *AddLiveMessageGroupBandRequest
	GetDataCenter() *string
	SetGroupId(v string) *AddLiveMessageGroupBandRequest
	GetGroupId() *string
}

type AddLiveMessageGroupBandRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The users whom you want to mute.
	//
	// This parameter is required.
	BannedUsers []*string `json:"BannedUsers,omitempty" xml:"BannedUsers,omitempty" type:"Repeated"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application.
	//
	// >  Valid values: cn-shanghai and ap-southeast-1 (Singapore).
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

func (s AddLiveMessageGroupBandRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveMessageGroupBandRequest) GoString() string {
	return s.String()
}

func (s *AddLiveMessageGroupBandRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddLiveMessageGroupBandRequest) GetBannedUsers() []*string {
	return s.BannedUsers
}

func (s *AddLiveMessageGroupBandRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *AddLiveMessageGroupBandRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddLiveMessageGroupBandRequest) SetAppId(v string) *AddLiveMessageGroupBandRequest {
	s.AppId = &v
	return s
}

func (s *AddLiveMessageGroupBandRequest) SetBannedUsers(v []*string) *AddLiveMessageGroupBandRequest {
	s.BannedUsers = v
	return s
}

func (s *AddLiveMessageGroupBandRequest) SetDataCenter(v string) *AddLiveMessageGroupBandRequest {
	s.DataCenter = &v
	return s
}

func (s *AddLiveMessageGroupBandRequest) SetGroupId(v string) *AddLiveMessageGroupBandRequest {
	s.GroupId = &v
	return s
}

func (s *AddLiveMessageGroupBandRequest) Validate() error {
	return dara.Validate(s)
}
