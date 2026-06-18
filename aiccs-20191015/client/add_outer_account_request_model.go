// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOuterAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *AddOuterAccountRequest
	GetAvatar() *string
	SetExt(v string) *AddOuterAccountRequest
	GetExt() *string
	SetOuterAccountId(v string) *AddOuterAccountRequest
	GetOuterAccountId() *string
	SetOuterAccountName(v string) *AddOuterAccountRequest
	GetOuterAccountName() *string
	SetOuterAccountType(v string) *AddOuterAccountRequest
	GetOuterAccountType() *string
	SetOuterDepartmentId(v string) *AddOuterAccountRequest
	GetOuterDepartmentId() *string
	SetOuterDepartmentType(v string) *AddOuterAccountRequest
	GetOuterDepartmentType() *string
	SetOuterGroupIds(v string) *AddOuterAccountRequest
	GetOuterGroupIds() *string
	SetOuterGroupType(v string) *AddOuterAccountRequest
	GetOuterGroupType() *string
	SetRealName(v string) *AddOuterAccountRequest
	GetRealName() *string
}

type AddOuterAccountRequest struct {
	// Profile picture.
	//
	// example:
	//
	// http://****
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// Extension field, in JSON string format.
	//
	// example:
	//
	// {"备注":"临时技能组"}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// External Account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1212
	OuterAccountId *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	// External account name.
	//
	// example:
	//
	// 测试
	OuterAccountName *string `json:"OuterAccountName,omitempty" xml:"OuterAccountName,omitempty"`
	// External account type.
	//
	// This parameter is required.
	//
	// example:
	//
	// alipay
	OuterAccountType *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
	// External department ID.
	//
	// example:
	//
	// 3
	OuterDepartmentId *string `json:"OuterDepartmentId,omitempty" xml:"OuterDepartmentId,omitempty"`
	// External department type: Invalid department type.
	//
	// example:
	//
	// type_invalid
	OuterDepartmentType *string `json:"OuterDepartmentType,omitempty" xml:"OuterDepartmentType,omitempty"`
	// List of external skill group IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["222","333","444"]
	OuterGroupIds *string `json:"OuterGroupIds,omitempty" xml:"OuterGroupIds,omitempty"`
	// Skill group type.
	//
	// example:
	//
	// mybank
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
	// Real name of the account.
	//
	// example:
	//
	// 张三
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
}

func (s AddOuterAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOuterAccountRequest) GoString() string {
	return s.String()
}

func (s *AddOuterAccountRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *AddOuterAccountRequest) GetExt() *string {
	return s.Ext
}

func (s *AddOuterAccountRequest) GetOuterAccountId() *string {
	return s.OuterAccountId
}

func (s *AddOuterAccountRequest) GetOuterAccountName() *string {
	return s.OuterAccountName
}

func (s *AddOuterAccountRequest) GetOuterAccountType() *string {
	return s.OuterAccountType
}

func (s *AddOuterAccountRequest) GetOuterDepartmentId() *string {
	return s.OuterDepartmentId
}

func (s *AddOuterAccountRequest) GetOuterDepartmentType() *string {
	return s.OuterDepartmentType
}

func (s *AddOuterAccountRequest) GetOuterGroupIds() *string {
	return s.OuterGroupIds
}

func (s *AddOuterAccountRequest) GetOuterGroupType() *string {
	return s.OuterGroupType
}

func (s *AddOuterAccountRequest) GetRealName() *string {
	return s.RealName
}

func (s *AddOuterAccountRequest) SetAvatar(v string) *AddOuterAccountRequest {
	s.Avatar = &v
	return s
}

func (s *AddOuterAccountRequest) SetExt(v string) *AddOuterAccountRequest {
	s.Ext = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterAccountId(v string) *AddOuterAccountRequest {
	s.OuterAccountId = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterAccountName(v string) *AddOuterAccountRequest {
	s.OuterAccountName = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterAccountType(v string) *AddOuterAccountRequest {
	s.OuterAccountType = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterDepartmentId(v string) *AddOuterAccountRequest {
	s.OuterDepartmentId = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterDepartmentType(v string) *AddOuterAccountRequest {
	s.OuterDepartmentType = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterGroupIds(v string) *AddOuterAccountRequest {
	s.OuterGroupIds = &v
	return s
}

func (s *AddOuterAccountRequest) SetOuterGroupType(v string) *AddOuterAccountRequest {
	s.OuterGroupType = &v
	return s
}

func (s *AddOuterAccountRequest) SetRealName(v string) *AddOuterAccountRequest {
	s.RealName = &v
	return s
}

func (s *AddOuterAccountRequest) Validate() error {
	return dara.Validate(s)
}
