// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelGroupDTO interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreate(v string) *ModelGroupDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ModelGroupDTO
	GetGmtModified() *string
	SetGroupId(v string) *ModelGroupDTO
	GetGroupId() *string
	SetModelCount(v int32) *ModelGroupDTO
	GetModelCount() *int32
	SetModelList(v []*int64) *ModelGroupDTO
	GetModelList() []*int64
	SetName(v string) *ModelGroupDTO
	GetName() *string
	SetType(v string) *ModelGroupDTO
	GetType() *string
}

type ModelGroupDTO struct {
	// example:
	//
	// 2026-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2026-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// mg_a1b2c3d4e5f6g7h8i9j0
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 3
	ModelCount *int32 `json:"modelCount,omitempty" xml:"modelCount,omitempty"`
	// example:
	//
	// [101, 102, 103]
	ModelList []*int64 `json:"modelList,omitempty" xml:"modelList,omitempty" type:"Repeated"`
	// example:
	//
	// Professional Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// manual
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelGroupDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelGroupDTO) GoString() string {
	return s.String()
}

func (s *ModelGroupDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelGroupDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelGroupDTO) GetGroupId() *string {
	return s.GroupId
}

func (s *ModelGroupDTO) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *ModelGroupDTO) GetModelList() []*int64 {
	return s.ModelList
}

func (s *ModelGroupDTO) GetName() *string {
	return s.Name
}

func (s *ModelGroupDTO) GetType() *string {
	return s.Type
}

func (s *ModelGroupDTO) SetGmtCreate(v string) *ModelGroupDTO {
	s.GmtCreate = &v
	return s
}

func (s *ModelGroupDTO) SetGmtModified(v string) *ModelGroupDTO {
	s.GmtModified = &v
	return s
}

func (s *ModelGroupDTO) SetGroupId(v string) *ModelGroupDTO {
	s.GroupId = &v
	return s
}

func (s *ModelGroupDTO) SetModelCount(v int32) *ModelGroupDTO {
	s.ModelCount = &v
	return s
}

func (s *ModelGroupDTO) SetModelList(v []*int64) *ModelGroupDTO {
	s.ModelList = v
	return s
}

func (s *ModelGroupDTO) SetName(v string) *ModelGroupDTO {
	s.Name = &v
	return s
}

func (s *ModelGroupDTO) SetType(v string) *ModelGroupDTO {
	s.Type = &v
	return s
}

func (s *ModelGroupDTO) Validate() error {
	return dara.Validate(s)
}
