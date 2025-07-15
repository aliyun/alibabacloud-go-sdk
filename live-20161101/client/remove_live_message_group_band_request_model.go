// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveLiveMessageGroupBandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RemoveLiveMessageGroupBandRequest
	GetAppId() *string
	SetDataCenter(v string) *RemoveLiveMessageGroupBandRequest
	GetDataCenter() *string
	SetGroupId(v string) *RemoveLiveMessageGroupBandRequest
	GetGroupId() *string
	SetUnbannedUsers(v []*string) *RemoveLiveMessageGroupBandRequest
	GetUnbannedUsers() []*string
}

type RemoveLiveMessageGroupBandRequest struct {
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
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The users whom you want to unmute.
	//
	// This parameter is required.
	UnbannedUsers []*string `json:"UnbannedUsers,omitempty" xml:"UnbannedUsers,omitempty" type:"Repeated"`
}

func (s RemoveLiveMessageGroupBandRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveLiveMessageGroupBandRequest) GoString() string {
	return s.String()
}

func (s *RemoveLiveMessageGroupBandRequest) GetAppId() *string {
	return s.AppId
}

func (s *RemoveLiveMessageGroupBandRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *RemoveLiveMessageGroupBandRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveLiveMessageGroupBandRequest) GetUnbannedUsers() []*string {
	return s.UnbannedUsers
}

func (s *RemoveLiveMessageGroupBandRequest) SetAppId(v string) *RemoveLiveMessageGroupBandRequest {
	s.AppId = &v
	return s
}

func (s *RemoveLiveMessageGroupBandRequest) SetDataCenter(v string) *RemoveLiveMessageGroupBandRequest {
	s.DataCenter = &v
	return s
}

func (s *RemoveLiveMessageGroupBandRequest) SetGroupId(v string) *RemoveLiveMessageGroupBandRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveLiveMessageGroupBandRequest) SetUnbannedUsers(v []*string) *RemoveLiveMessageGroupBandRequest {
	s.UnbannedUsers = v
	return s
}

func (s *RemoveLiveMessageGroupBandRequest) Validate() error {
	return dara.Validate(s)
}
