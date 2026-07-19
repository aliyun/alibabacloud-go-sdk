// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRbacRoleHierarchyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RemoveRbacRoleHierarchyRequest
	GetBizId() *string
	SetChildRoleId(v string) *RemoveRbacRoleHierarchyRequest
	GetChildRoleId() *string
	SetParentRoleId(v string) *RemoveRbacRoleHierarchyRequest
	GetParentRoleId() *string
}

type RemoveRbacRoleHierarchyRequest struct {
	// example:
	//
	// WS20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// fc94cc51-310f-4485-adb2-ed8c706aff3b
	ChildRoleId *string `json:"ChildRoleId,omitempty" xml:"ChildRoleId,omitempty"`
	// example:
	//
	// 71e07711-9a17-49f4-9f83-387a60ee5b64
	ParentRoleId *string `json:"ParentRoleId,omitempty" xml:"ParentRoleId,omitempty"`
}

func (s RemoveRbacRoleHierarchyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRbacRoleHierarchyRequest) GoString() string {
	return s.String()
}

func (s *RemoveRbacRoleHierarchyRequest) GetBizId() *string {
	return s.BizId
}

func (s *RemoveRbacRoleHierarchyRequest) GetChildRoleId() *string {
	return s.ChildRoleId
}

func (s *RemoveRbacRoleHierarchyRequest) GetParentRoleId() *string {
	return s.ParentRoleId
}

func (s *RemoveRbacRoleHierarchyRequest) SetBizId(v string) *RemoveRbacRoleHierarchyRequest {
	s.BizId = &v
	return s
}

func (s *RemoveRbacRoleHierarchyRequest) SetChildRoleId(v string) *RemoveRbacRoleHierarchyRequest {
	s.ChildRoleId = &v
	return s
}

func (s *RemoveRbacRoleHierarchyRequest) SetParentRoleId(v string) *RemoveRbacRoleHierarchyRequest {
	s.ParentRoleId = &v
	return s
}

func (s *RemoveRbacRoleHierarchyRequest) Validate() error {
	return dara.Validate(s)
}
