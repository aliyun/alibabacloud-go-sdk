// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverLiveMessageDeletedGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RecoverLiveMessageDeletedGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *RecoverLiveMessageDeletedGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *RecoverLiveMessageDeletedGroupRequest
	GetGroupId() *string
}

type RecoverLiveMessageDeletedGroupRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the group that you want to restore.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s RecoverLiveMessageDeletedGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverLiveMessageDeletedGroupRequest) GoString() string {
	return s.String()
}

func (s *RecoverLiveMessageDeletedGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *RecoverLiveMessageDeletedGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *RecoverLiveMessageDeletedGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RecoverLiveMessageDeletedGroupRequest) SetAppId(v string) *RecoverLiveMessageDeletedGroupRequest {
	s.AppId = &v
	return s
}

func (s *RecoverLiveMessageDeletedGroupRequest) SetDataCenter(v string) *RecoverLiveMessageDeletedGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *RecoverLiveMessageDeletedGroupRequest) SetGroupId(v string) *RecoverLiveMessageDeletedGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RecoverLiveMessageDeletedGroupRequest) Validate() error {
	return dara.Validate(s)
}
