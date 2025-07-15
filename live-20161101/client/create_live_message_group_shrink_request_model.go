// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveMessageGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministratorsShrink(v string) *CreateLiveMessageGroupShrinkRequest
	GetAdministratorsShrink() *string
	SetAppId(v string) *CreateLiveMessageGroupShrinkRequest
	GetAppId() *string
	SetCreatorId(v string) *CreateLiveMessageGroupShrinkRequest
	GetCreatorId() *string
	SetDataCenter(v string) *CreateLiveMessageGroupShrinkRequest
	GetDataCenter() *string
	SetGroupId(v string) *CreateLiveMessageGroupShrinkRequest
	GetGroupId() *string
	SetGroupInfo(v string) *CreateLiveMessageGroupShrinkRequest
	GetGroupInfo() *string
	SetGroupName(v string) *CreateLiveMessageGroupShrinkRequest
	GetGroupName() *string
}

type CreateLiveMessageGroupShrinkRequest struct {
	// The list of administrators.
	AdministratorsShrink *string `json:"Administrators,omitempty" xml:"Administrators,omitempty"`
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

func (s CreateLiveMessageGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveMessageGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveMessageGroupShrinkRequest) GetAdministratorsShrink() *string {
	return s.AdministratorsShrink
}

func (s *CreateLiveMessageGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateLiveMessageGroupShrinkRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateLiveMessageGroupShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateLiveMessageGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateLiveMessageGroupShrinkRequest) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *CreateLiveMessageGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLiveMessageGroupShrinkRequest) SetAdministratorsShrink(v string) *CreateLiveMessageGroupShrinkRequest {
	s.AdministratorsShrink = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) SetAppId(v string) *CreateLiveMessageGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) SetCreatorId(v string) *CreateLiveMessageGroupShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) SetDataCenter(v string) *CreateLiveMessageGroupShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) SetGroupId(v string) *CreateLiveMessageGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) SetGroupInfo(v string) *CreateLiveMessageGroupShrinkRequest {
	s.GroupInfo = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) SetGroupName(v string) *CreateLiveMessageGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLiveMessageGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
