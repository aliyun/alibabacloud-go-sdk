// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLhMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMembersShrink(v string) *AddLhMembersShrinkRequest
	GetMembersShrink() *string
	SetObjectId(v int64) *AddLhMembersShrinkRequest
	GetObjectId() *int64
	SetObjectType(v int32) *AddLhMembersShrinkRequest
	GetObjectType() *int32
	SetTid(v int64) *AddLhMembersShrinkRequest
	GetTid() *int64
}

type AddLhMembersShrinkRequest struct {
	// The information about the users to be added.
	//
	// This parameter is required.
	MembersShrink *string `json:"Members,omitempty" xml:"Members,omitempty"`
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

func (s AddLhMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLhMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddLhMembersShrinkRequest) GetMembersShrink() *string {
	return s.MembersShrink
}

func (s *AddLhMembersShrinkRequest) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *AddLhMembersShrinkRequest) GetObjectType() *int32 {
	return s.ObjectType
}

func (s *AddLhMembersShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddLhMembersShrinkRequest) SetMembersShrink(v string) *AddLhMembersShrinkRequest {
	s.MembersShrink = &v
	return s
}

func (s *AddLhMembersShrinkRequest) SetObjectId(v int64) *AddLhMembersShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *AddLhMembersShrinkRequest) SetObjectType(v int32) *AddLhMembersShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *AddLhMembersShrinkRequest) SetTid(v int64) *AddLhMembersShrinkRequest {
	s.Tid = &v
	return s
}

func (s *AddLhMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
