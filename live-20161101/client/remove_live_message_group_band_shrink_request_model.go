// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveLiveMessageGroupBandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RemoveLiveMessageGroupBandShrinkRequest
	GetAppId() *string
	SetDataCenter(v string) *RemoveLiveMessageGroupBandShrinkRequest
	GetDataCenter() *string
	SetGroupId(v string) *RemoveLiveMessageGroupBandShrinkRequest
	GetGroupId() *string
	SetUnbannedUsersShrink(v string) *RemoveLiveMessageGroupBandShrinkRequest
	GetUnbannedUsersShrink() *string
}

type RemoveLiveMessageGroupBandShrinkRequest struct {
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
	UnbannedUsersShrink *string `json:"UnbannedUsers,omitempty" xml:"UnbannedUsers,omitempty"`
}

func (s RemoveLiveMessageGroupBandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveLiveMessageGroupBandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) GetUnbannedUsersShrink() *string {
	return s.UnbannedUsersShrink
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) SetAppId(v string) *RemoveLiveMessageGroupBandShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) SetDataCenter(v string) *RemoveLiveMessageGroupBandShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) SetGroupId(v string) *RemoveLiveMessageGroupBandShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) SetUnbannedUsersShrink(v string) *RemoveLiveMessageGroupBandShrinkRequest {
	s.UnbannedUsersShrink = &v
	return s
}

func (s *RemoveLiveMessageGroupBandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
