// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageRelationRegisterVO interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimestamp(v int64) *LineageRelationRegisterVO
	GetCreateTimestamp() *int64
	SetDestEntity(v *LineageEntityVO) *LineageRelationRegisterVO
	GetDestEntity() *LineageEntityVO
	SetRelationship(v *RelationshipVO) *LineageRelationRegisterVO
	GetRelationship() *RelationshipVO
	SetSrcEntity(v *LineageEntityVO) *LineageRelationRegisterVO
	GetSrcEntity() *LineageEntityVO
}

type LineageRelationRegisterVO struct {
	// example:
	//
	// 1684327487964
	CreateTimestamp *int64           `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DestEntity      *LineageEntityVO `json:"DestEntity,omitempty" xml:"DestEntity,omitempty"`
	Relationship    *RelationshipVO  `json:"Relationship,omitempty" xml:"Relationship,omitempty"`
	SrcEntity       *LineageEntityVO `json:"SrcEntity,omitempty" xml:"SrcEntity,omitempty"`
}

func (s LineageRelationRegisterVO) String() string {
	return dara.Prettify(s)
}

func (s LineageRelationRegisterVO) GoString() string {
	return s.String()
}

func (s *LineageRelationRegisterVO) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *LineageRelationRegisterVO) GetDestEntity() *LineageEntityVO {
	return s.DestEntity
}

func (s *LineageRelationRegisterVO) GetRelationship() *RelationshipVO {
	return s.Relationship
}

func (s *LineageRelationRegisterVO) GetSrcEntity() *LineageEntityVO {
	return s.SrcEntity
}

func (s *LineageRelationRegisterVO) SetCreateTimestamp(v int64) *LineageRelationRegisterVO {
	s.CreateTimestamp = &v
	return s
}

func (s *LineageRelationRegisterVO) SetDestEntity(v *LineageEntityVO) *LineageRelationRegisterVO {
	s.DestEntity = v
	return s
}

func (s *LineageRelationRegisterVO) SetRelationship(v *RelationshipVO) *LineageRelationRegisterVO {
	s.Relationship = v
	return s
}

func (s *LineageRelationRegisterVO) SetSrcEntity(v *LineageEntityVO) *LineageRelationRegisterVO {
	s.SrcEntity = v
	return s
}

func (s *LineageRelationRegisterVO) Validate() error {
	return dara.Validate(s)
}
