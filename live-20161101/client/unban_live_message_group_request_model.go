// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbanLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UnbanLiveMessageGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *UnbanLiveMessageGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *UnbanLiveMessageGroupRequest
	GetGroupId() *string
}

type UnbanLiveMessageGroupRequest struct {
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
}

func (s UnbanLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbanLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *UnbanLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *UnbanLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *UnbanLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UnbanLiveMessageGroupRequest) SetAppId(v string) *UnbanLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *UnbanLiveMessageGroupRequest) SetDataCenter(v string) *UnbanLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *UnbanLiveMessageGroupRequest) SetGroupId(v string) *UnbanLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UnbanLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
