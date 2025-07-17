// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLhMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberIdsShrink(v string) *DeleteLhMembersShrinkRequest
	GetMemberIdsShrink() *string
	SetObjectId(v int64) *DeleteLhMembersShrinkRequest
	GetObjectId() *int64
	SetObjectType(v int32) *DeleteLhMembersShrinkRequest
	GetObjectType() *int32
	SetTid(v int64) *DeleteLhMembersShrinkRequest
	GetTid() *int64
}

type DeleteLhMembersShrinkRequest struct {
	// The ID of the user to be removed. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain the user ID.
	//
	// This parameter is required.
	MemberIdsShrink *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
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

func (s DeleteLhMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLhMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteLhMembersShrinkRequest) GetMemberIdsShrink() *string {
	return s.MemberIdsShrink
}

func (s *DeleteLhMembersShrinkRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *DeleteLhMembersShrinkRequest) GetObjectType() *int32 {
	return s.ObjectType
}

func (s *DeleteLhMembersShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteLhMembersShrinkRequest) SetMemberIdsShrink(v string) *DeleteLhMembersShrinkRequest {
	s.MemberIdsShrink = &v
	return s
}

func (s *DeleteLhMembersShrinkRequest) SetObjectId(v int64) *DeleteLhMembersShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *DeleteLhMembersShrinkRequest) SetObjectType(v int32) *DeleteLhMembersShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *DeleteLhMembersShrinkRequest) SetTid(v int64) *DeleteLhMembersShrinkRequest {
	s.Tid = &v
	return s
}

func (s *DeleteLhMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
