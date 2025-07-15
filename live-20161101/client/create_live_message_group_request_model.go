// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministrators(v []*string) *CreateLiveMessageGroupRequest
	GetAdministrators() []*string
	SetAppId(v string) *CreateLiveMessageGroupRequest
	GetAppId() *string
	SetCreatorId(v string) *CreateLiveMessageGroupRequest
	GetCreatorId() *string
	SetDataCenter(v string) *CreateLiveMessageGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *CreateLiveMessageGroupRequest
	GetGroupId() *string
	SetGroupInfo(v string) *CreateLiveMessageGroupRequest
	GetGroupInfo() *string
	SetGroupName(v string) *CreateLiveMessageGroupRequest
	GetGroupName() *string
}

type CreateLiveMessageGroupRequest struct {
	// The list of administrators.
	Administrators []*string `json:"Administrators,omitempty" xml:"Administrators,omitempty" type:"Repeated"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the group creator. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// example:
	//
	// uid1
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2593195.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the group that you want to create. The group ID must be unique within your business. The ID can be up to 64 bytes in length and can contain letters and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The additional information about the group. The value can be up to 32 KB in length.
	//
	// example:
	//
	// testgroupinfo
	GroupInfo *string `json:"GroupInfo,omitempty" xml:"GroupInfo,omitempty"`
	// The name of the group. The name can be up to 128 bytes in length.
	//
	// example:
	//
	// mytestgroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageGroupRequest) GetAdministrators() []*string {
	return s.Administrators
}

func (s *CreateLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateLiveMessageGroupRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateLiveMessageGroupRequest) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *CreateLiveMessageGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLiveMessageGroupRequest) SetAdministrators(v []*string) *CreateLiveMessageGroupRequest {
	s.Administrators = v
	return s
}

func (s *CreateLiveMessageGroupRequest) SetAppId(v string) *CreateLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveMessageGroupRequest) SetCreatorId(v string) *CreateLiveMessageGroupRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateLiveMessageGroupRequest) SetDataCenter(v string) *CreateLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateLiveMessageGroupRequest) SetGroupId(v string) *CreateLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateLiveMessageGroupRequest) SetGroupInfo(v string) *CreateLiveMessageGroupRequest {
	s.GroupInfo = &v
	return s
}

func (s *CreateLiveMessageGroupRequest) SetGroupName(v string) *CreateLiveMessageGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
