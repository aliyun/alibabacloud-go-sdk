// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRbacRoleHierarchyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SetRbacRoleHierarchyRequest
	GetBizId() *string
	SetChildRoleId(v string) *SetRbacRoleHierarchyRequest
	GetChildRoleId() *string
	SetParentRoleId(v string) *SetRbacRoleHierarchyRequest
	GetParentRoleId() *string
}

type SetRbacRoleHierarchyRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The child role ID.
	//
	// example:
	//
	// fc94cc51-310f-4485-adb2-ed8c706aff3b
	ChildRoleId *string `json:"ChildRoleId,omitempty" xml:"ChildRoleId,omitempty"`
	// The parent role ID.
	//
	// example:
	//
	// 71e07711-9a17-49f4-9f83-387a60ee5b64
	ParentRoleId *string `json:"ParentRoleId,omitempty" xml:"ParentRoleId,omitempty"`
}

func (s SetRbacRoleHierarchyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRbacRoleHierarchyRequest) GoString() string {
	return s.String()
}

func (s *SetRbacRoleHierarchyRequest) GetBizId() *string {
	return s.BizId
}

func (s *SetRbacRoleHierarchyRequest) GetChildRoleId() *string {
	return s.ChildRoleId
}

func (s *SetRbacRoleHierarchyRequest) GetParentRoleId() *string {
	return s.ParentRoleId
}

func (s *SetRbacRoleHierarchyRequest) SetBizId(v string) *SetRbacRoleHierarchyRequest {
	s.BizId = &v
	return s
}

func (s *SetRbacRoleHierarchyRequest) SetChildRoleId(v string) *SetRbacRoleHierarchyRequest {
	s.ChildRoleId = &v
	return s
}

func (s *SetRbacRoleHierarchyRequest) SetParentRoleId(v string) *SetRbacRoleHierarchyRequest {
	s.ParentRoleId = &v
	return s
}

func (s *SetRbacRoleHierarchyRequest) Validate() error {
	return dara.Validate(s)
}
