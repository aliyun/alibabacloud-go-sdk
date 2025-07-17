// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLhMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberIds(v []*int32) *DeleteLhMembersRequest
	GetMemberIds() []*int32
	SetObjectId(v int64) *DeleteLhMembersRequest
	GetObjectId() *int64
	SetObjectType(v int32) *DeleteLhMembersRequest
	GetObjectType() *int32
	SetTid(v int64) *DeleteLhMembersRequest
	GetTid() *int64
}

type DeleteLhMembersRequest struct {
	// The ID of the user to be removed. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain the user ID.
	//
	// This parameter is required.
	MemberIds []*int32 `json:"MemberIds,omitempty" xml:"MemberIds,omitempty" type:"Repeated"`
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
	// 9***
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
	// 1
	ObjectType *int32 `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteLhMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLhMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteLhMembersRequest) GetMemberIds() []*int32 {
	return s.MemberIds
}

func (s *DeleteLhMembersRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *DeleteLhMembersRequest) GetObjectType() *int32 {
	return s.ObjectType
}

func (s *DeleteLhMembersRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteLhMembersRequest) SetMemberIds(v []*int32) *DeleteLhMembersRequest {
	s.MemberIds = v
	return s
}

func (s *DeleteLhMembersRequest) SetObjectId(v int64) *DeleteLhMembersRequest {
	s.ObjectId = &v
	return s
}

func (s *DeleteLhMembersRequest) SetObjectType(v int32) *DeleteLhMembersRequest {
	s.ObjectType = &v
	return s
}

func (s *DeleteLhMembersRequest) SetTid(v int64) *DeleteLhMembersRequest {
	s.Tid = &v
	return s
}

func (s *DeleteLhMembersRequest) Validate() error {
	return dara.Validate(s)
}
