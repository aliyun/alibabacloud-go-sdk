// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteLiveMessageGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *DeleteLiveMessageGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *DeleteLiveMessageGroupRequest
	GetGroupId() *string
	SetOperatorId(v string) *DeleteLiveMessageGroupRequest
	GetOperatorId() *string
}

type DeleteLiveMessageGroupRequest struct {
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
	// The ID of the group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the user who performs the deletion operation. The ID can be up to 64 bytes in length and can contain only letters and digits.
	//
	// example:
	//
	// uid1
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
}

func (s DeleteLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *DeleteLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteLiveMessageGroupRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *DeleteLiveMessageGroupRequest) SetAppId(v string) *DeleteLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *DeleteLiveMessageGroupRequest) SetDataCenter(v string) *DeleteLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *DeleteLiveMessageGroupRequest) SetGroupId(v string) *DeleteLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteLiveMessageGroupRequest) SetOperatorId(v string) *DeleteLiveMessageGroupRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
