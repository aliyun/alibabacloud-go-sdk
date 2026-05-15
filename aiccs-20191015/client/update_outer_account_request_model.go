// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOuterAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *UpdateOuterAccountRequest
	GetAvatar() *string
	SetExt(v string) *UpdateOuterAccountRequest
	GetExt() *string
	SetOuterAccountId(v string) *UpdateOuterAccountRequest
	GetOuterAccountId() *string
	SetOuterAccountName(v string) *UpdateOuterAccountRequest
	GetOuterAccountName() *string
	SetOuterAccountType(v string) *UpdateOuterAccountRequest
	GetOuterAccountType() *string
	SetOuterDepartmentId(v string) *UpdateOuterAccountRequest
	GetOuterDepartmentId() *string
	SetOuterDepartmentType(v string) *UpdateOuterAccountRequest
	GetOuterDepartmentType() *string
	SetOuterGroupIds(v string) *UpdateOuterAccountRequest
	GetOuterGroupIds() *string
	SetOuterGroupType(v string) *UpdateOuterAccountRequest
	GetOuterGroupType() *string
	SetRealName(v string) *UpdateOuterAccountRequest
	GetRealName() *string
}

type UpdateOuterAccountRequest struct {
	// example:
	//
	// http://****
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Ext    *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OuterAccountId   *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	OuterAccountName *string `json:"OuterAccountName,omitempty" xml:"OuterAccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alipay
	OuterAccountType *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
	// example:
	//
	// 3
	OuterDepartmentId *string `json:"OuterDepartmentId,omitempty" xml:"OuterDepartmentId,omitempty"`
	// example:
	//
	// type_invalid
	OuterDepartmentType *string `json:"OuterDepartmentType,omitempty" xml:"OuterDepartmentType,omitempty"`
	// example:
	//
	// ["222","333","444"]
	OuterGroupIds *string `json:"OuterGroupIds,omitempty" xml:"OuterGroupIds,omitempty"`
	// example:
	//
	// mybank
	OuterGroupType *string `json:"OuterGroupType,omitempty" xml:"OuterGroupType,omitempty"`
	RealName       *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
}

func (s UpdateOuterAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOuterAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateOuterAccountRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateOuterAccountRequest) GetExt() *string {
	return s.Ext
}

func (s *UpdateOuterAccountRequest) GetOuterAccountId() *string {
	return s.OuterAccountId
}

func (s *UpdateOuterAccountRequest) GetOuterAccountName() *string {
	return s.OuterAccountName
}

func (s *UpdateOuterAccountRequest) GetOuterAccountType() *string {
	return s.OuterAccountType
}

func (s *UpdateOuterAccountRequest) GetOuterDepartmentId() *string {
	return s.OuterDepartmentId
}

func (s *UpdateOuterAccountRequest) GetOuterDepartmentType() *string {
	return s.OuterDepartmentType
}

func (s *UpdateOuterAccountRequest) GetOuterGroupIds() *string {
	return s.OuterGroupIds
}

func (s *UpdateOuterAccountRequest) GetOuterGroupType() *string {
	return s.OuterGroupType
}

func (s *UpdateOuterAccountRequest) GetRealName() *string {
	return s.RealName
}

func (s *UpdateOuterAccountRequest) SetAvatar(v string) *UpdateOuterAccountRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetExt(v string) *UpdateOuterAccountRequest {
	s.Ext = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterAccountId(v string) *UpdateOuterAccountRequest {
	s.OuterAccountId = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterAccountName(v string) *UpdateOuterAccountRequest {
	s.OuterAccountName = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterAccountType(v string) *UpdateOuterAccountRequest {
	s.OuterAccountType = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterDepartmentId(v string) *UpdateOuterAccountRequest {
	s.OuterDepartmentId = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterDepartmentType(v string) *UpdateOuterAccountRequest {
	s.OuterDepartmentType = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterGroupIds(v string) *UpdateOuterAccountRequest {
	s.OuterGroupIds = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetOuterGroupType(v string) *UpdateOuterAccountRequest {
	s.OuterGroupType = &v
	return s
}

func (s *UpdateOuterAccountRequest) SetRealName(v string) *UpdateOuterAccountRequest {
	s.RealName = &v
	return s
}

func (s *UpdateOuterAccountRequest) Validate() error {
	return dara.Validate(s)
}
