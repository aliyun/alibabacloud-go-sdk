// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChildGroups(v []interface{}) *GetUserGroupResponseBody
	GetChildGroups() []interface{}
	SetCode(v string) *GetUserGroupResponseBody
	GetCode() *string
	SetMembers(v []interface{}) *GetUserGroupResponseBody
	GetMembers() []interface{}
	SetMessage(v string) *GetUserGroupResponseBody
	GetMessage() *string
	SetParentGroup(v interface{}) *GetUserGroupResponseBody
	GetParentGroup() interface{}
	SetRequestId(v string) *GetUserGroupResponseBody
	GetRequestId() *string
	SetUserGroup(v interface{}) *GetUserGroupResponseBody
	GetUserGroup() interface{}
}

type GetUserGroupResponseBody struct {
	// **The list of direct child user groups.**
	ChildGroups []interface{} `json:"childGroups,omitempty" xml:"childGroups,omitempty" type:"Repeated"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// **The list of direct members in the current user group.**
	Members []interface{} `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// **The parent user group information. This is empty for the root node.**
	//
	// example:
	//
	// {"userGroupId":"b07fb0a4-0b7d-44a7-a3d5-a43a5964c8f0","userGroupName":"Sales Center","parentId":null,"level":1,"description":"Sales organization","childGroupCount":1,"directMemberCount":0,"sourceType":"internal","externalSyncStatus":null,"gmtCreate":"2026-08-27T08:00:00Z","gmtModified":"2026-08-27T08:00:00Z"}
	ParentGroup interface{} `json:"parentGroup,omitempty" xml:"parentGroup,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// **The target user group information.**
	//
	// example:
	//
	// {"userGroupId":"7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11","userGroupName":"South China Sales","parentId":"b07fb0a4-0b7d-44a7-a3d5-a43a5964c8f0","level":2,"description":"South China Sales organization","childGroupCount":0,"directMemberCount":2,"sourceType":"internal","externalSyncStatus":null,"gmtCreate":"2026-08-27T09:00:00Z","gmtModified":"2026-08-27T10:00:00Z"}
	UserGroup interface{} `json:"userGroup,omitempty" xml:"userGroup,omitempty"`
}

func (s GetUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) GetChildGroups() []interface{} {
	return s.ChildGroups
}

func (s *GetUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserGroupResponseBody) GetMembers() []interface{} {
	return s.Members
}

func (s *GetUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserGroupResponseBody) GetParentGroup() interface{} {
	return s.ParentGroup
}

func (s *GetUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserGroupResponseBody) GetUserGroup() interface{} {
	return s.UserGroup
}

func (s *GetUserGroupResponseBody) SetChildGroups(v []interface{}) *GetUserGroupResponseBody {
	s.ChildGroups = v
	return s
}

func (s *GetUserGroupResponseBody) SetCode(v string) *GetUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserGroupResponseBody) SetMembers(v []interface{}) *GetUserGroupResponseBody {
	s.Members = v
	return s
}

func (s *GetUserGroupResponseBody) SetMessage(v string) *GetUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserGroupResponseBody) SetParentGroup(v interface{}) *GetUserGroupResponseBody {
	s.ParentGroup = v
	return s
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroup(v interface{}) *GetUserGroupResponseBody {
	s.UserGroup = v
	return s
}

func (s *GetUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
