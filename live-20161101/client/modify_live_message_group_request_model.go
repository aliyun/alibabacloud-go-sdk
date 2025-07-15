// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminList(v []*string) *ModifyLiveMessageGroupRequest
	GetAdminList() []*string
	SetAppId(v string) *ModifyLiveMessageGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *ModifyLiveMessageGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *ModifyLiveMessageGroupRequest
	GetGroupId() *string
	SetGroupInfo(v string) *ModifyLiveMessageGroupRequest
	GetGroupInfo() *string
	SetModifyAdmin(v bool) *ModifyLiveMessageGroupRequest
	GetModifyAdmin() *bool
	SetModifyInfo(v bool) *ModifyLiveMessageGroupRequest
	GetModifyInfo() *bool
}

type ModifyLiveMessageGroupRequest struct {
	// The list of administrators after your change.
	AdminList []*string `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
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

func (s ModifyLiveMessageGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupRequest) GetAdminList() []*string {
	return s.AdminList
}

func (s *ModifyLiveMessageGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyLiveMessageGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ModifyLiveMessageGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLiveMessageGroupRequest) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *ModifyLiveMessageGroupRequest) GetModifyAdmin() *bool {
	return s.ModifyAdmin
}

func (s *ModifyLiveMessageGroupRequest) GetModifyInfo() *bool {
	return s.ModifyInfo
}

func (s *ModifyLiveMessageGroupRequest) SetAdminList(v []*string) *ModifyLiveMessageGroupRequest {
	s.AdminList = v
	return s
}

func (s *ModifyLiveMessageGroupRequest) SetAppId(v string) *ModifyLiveMessageGroupRequest {
	s.AppId = &v
	return s
}

func (s *ModifyLiveMessageGroupRequest) SetDataCenter(v string) *ModifyLiveMessageGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *ModifyLiveMessageGroupRequest) SetGroupId(v string) *ModifyLiveMessageGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyLiveMessageGroupRequest) SetGroupInfo(v string) *ModifyLiveMessageGroupRequest {
	s.GroupInfo = &v
	return s
}

func (s *ModifyLiveMessageGroupRequest) SetModifyAdmin(v bool) *ModifyLiveMessageGroupRequest {
	s.ModifyAdmin = &v
	return s
}

func (s *ModifyLiveMessageGroupRequest) SetModifyInfo(v bool) *ModifyLiveMessageGroupRequest {
	s.ModifyInfo = &v
	return s
}

func (s *ModifyLiveMessageGroupRequest) Validate() error {
	return dara.Validate(s)
}
