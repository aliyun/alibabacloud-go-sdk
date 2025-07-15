// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminListShrink(v string) *ModifyLiveMessageGroupShrinkRequest
	GetAdminListShrink() *string
	SetAppId(v string) *ModifyLiveMessageGroupShrinkRequest
	GetAppId() *string
	SetDataCenter(v string) *ModifyLiveMessageGroupShrinkRequest
	GetDataCenter() *string
	SetGroupId(v string) *ModifyLiveMessageGroupShrinkRequest
	GetGroupId() *string
	SetGroupInfo(v string) *ModifyLiveMessageGroupShrinkRequest
	GetGroupInfo() *string
	SetModifyAdmin(v bool) *ModifyLiveMessageGroupShrinkRequest
	GetModifyAdmin() *bool
	SetModifyInfo(v bool) *ModifyLiveMessageGroupShrinkRequest
	GetModifyInfo() *bool
}

type ModifyLiveMessageGroupShrinkRequest struct {
	// The list of administrators after your change.
	AdminListShrink *string `json:"AdminList,omitempty" xml:"AdminList,omitempty"`
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
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The additional information about the group after the modification. The value can be up to 32 KB in length.
	//
	// example:
	//
	// newmeta
	GroupInfo *string `json:"GroupInfo,omitempty" xml:"GroupInfo,omitempty"`
	// Specifies whether to change the group administrators.
	//
	// example:
	//
	// true
	ModifyAdmin *bool `json:"ModifyAdmin,omitempty" xml:"ModifyAdmin,omitempty"`
	// Specifies whether to modify the additional information about the group.
	//
	// example:
	//
	// true
	ModifyInfo *bool `json:"ModifyInfo,omitempty" xml:"ModifyInfo,omitempty"`
}

func (s ModifyLiveMessageGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetAdminListShrink() *string {
	return s.AdminListShrink
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetModifyAdmin() *bool {
	return s.ModifyAdmin
}

func (s *ModifyLiveMessageGroupShrinkRequest) GetModifyInfo() *bool {
	return s.ModifyInfo
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetAdminListShrink(v string) *ModifyLiveMessageGroupShrinkRequest {
	s.AdminListShrink = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetAppId(v string) *ModifyLiveMessageGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetDataCenter(v string) *ModifyLiveMessageGroupShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetGroupId(v string) *ModifyLiveMessageGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetGroupInfo(v string) *ModifyLiveMessageGroupShrinkRequest {
	s.GroupInfo = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetModifyAdmin(v bool) *ModifyLiveMessageGroupShrinkRequest {
	s.ModifyAdmin = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) SetModifyInfo(v bool) *ModifyLiveMessageGroupShrinkRequest {
	s.ModifyInfo = &v
	return s
}

func (s *ModifyLiveMessageGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
