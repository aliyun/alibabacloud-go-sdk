// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLhMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*AddLhMembersRequestMembers) *AddLhMembersRequest
	GetMembers() []*AddLhMembersRequestMembers
	SetObjectId(v int64) *AddLhMembersRequest
	GetObjectId() *int64
	SetObjectType(v int32) *AddLhMembersRequest
	GetObjectType() *int32
	SetTid(v int64) *AddLhMembersRequest
	GetTid() *int64
}

type AddLhMembersRequest struct {
	// The information about the users to be added.
	//
	// This parameter is required.
	Members []*AddLhMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The ID of the object.
	//
	// 	- If the object is a workspace, you can call the [GetLhSpaceByName](https://help.aliyun.com/document_detail/424379.html) operation to obtain the workspace ID.
	//
	// 	- If the object is a task flow, you can call the [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ObjectId *int64 `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- **0**: workspace
	//
	// 	- **1**: task flow
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ObjectType *int32 `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3000
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddLhMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLhMembersRequest) GoString() string {
	return s.String()
}

func (s *AddLhMembersRequest) GetMembers() []*AddLhMembersRequestMembers {
	return s.Members
}

func (s *AddLhMembersRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *AddLhMembersRequest) GetObjectType() *int32 {
	return s.ObjectType
}

func (s *AddLhMembersRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddLhMembersRequest) SetMembers(v []*AddLhMembersRequestMembers) *AddLhMembersRequest {
	s.Members = v
	return s
}

func (s *AddLhMembersRequest) SetObjectId(v int64) *AddLhMembersRequest {
	s.ObjectId = &v
	return s
}

func (s *AddLhMembersRequest) SetObjectType(v int32) *AddLhMembersRequest {
	s.ObjectType = &v
	return s
}

func (s *AddLhMembersRequest) SetTid(v int64) *AddLhMembersRequest {
	s.Tid = &v
	return s
}

func (s *AddLhMembersRequest) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddLhMembersRequestMembers struct {
	// The role. Valid values:
	//
	// 	- **ADMIN**: workspace administrator. You can add a workspace administrator only as a DMS administrator or a DBA.
	//
	// 	- **MEMBER**: workspace member.
	//
	// 	- **DEVELOPER**: task flow developer. Only a workspace member can be added as a task flow developer.
	//
	// This parameter is required.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The ID of the user to be added. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddLhMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s AddLhMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddLhMembersRequestMembers) GetRoles() []*string {
	return s.Roles
}

func (s *AddLhMembersRequestMembers) GetUserId() *int64 {
	return s.UserId
}

func (s *AddLhMembersRequestMembers) SetRoles(v []*string) *AddLhMembersRequestMembers {
	s.Roles = v
	return s
}

func (s *AddLhMembersRequestMembers) SetUserId(v int64) *AddLhMembersRequestMembers {
	s.UserId = &v
	return s
}

func (s *AddLhMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}
