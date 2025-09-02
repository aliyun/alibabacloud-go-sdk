// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageRelationRegisterBulkVO interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimestamp(v int64) *LineageRelationRegisterBulkVO
	GetCreateTimestamp() *int64
	SetDestEntities(v []*LineageEntityVO) *LineageRelationRegisterBulkVO
	GetDestEntities() []*LineageEntityVO
	SetRelationship(v *RelationshipVO) *LineageRelationRegisterBulkVO
	GetRelationship() *RelationshipVO
	SetSrcEntities(v []*LineageEntityVO) *LineageRelationRegisterBulkVO
	GetSrcEntities() []*LineageEntityVO
}

type LineageRelationRegisterBulkVO struct {
	// example:
	//
	// 1684327487964
	CreateTimestamp *int64             `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DestEntities    []*LineageEntityVO `json:"DestEntities,omitempty" xml:"DestEntities,omitempty" type:"Repeated"`
	Relationship    *RelationshipVO    `json:"Relationship,omitempty" xml:"Relationship,omitempty"`
	SrcEntities     []*LineageEntityVO `json:"SrcEntities,omitempty" xml:"SrcEntities,omitempty" type:"Repeated"`
}

func (s LineageRelationRegisterBulkVO) String() string {
	return dara.Prettify(s)
}

func (s LineageRelationRegisterBulkVO) GoString() string {
	return s.String()
}

func (s *LineageRelationRegisterBulkVO) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *LineageRelationRegisterBulkVO) GetDestEntities() []*LineageEntityVO {
	return s.DestEntities
}

func (s *LineageRelationRegisterBulkVO) GetRelationship() *RelationshipVO {
	return s.Relationship
}

func (s *LineageRelationRegisterBulkVO) GetSrcEntities() []*LineageEntityVO {
	return s.SrcEntities
}

func (s *LineageRelationRegisterBulkVO) SetCreateTimestamp(v int64) *LineageRelationRegisterBulkVO {
	s.CreateTimestamp = &v
	return s
}

func (s *LineageRelationRegisterBulkVO) SetDestEntities(v []*LineageEntityVO) *LineageRelationRegisterBulkVO {
	s.DestEntities = v
	return s
}

func (s *LineageRelationRegisterBulkVO) SetRelationship(v *RelationshipVO) *LineageRelationRegisterBulkVO {
	s.Relationship = v
	return s
}

func (s *LineageRelationRegisterBulkVO) SetSrcEntities(v []*LineageEntityVO) *LineageRelationRegisterBulkVO {
	s.SrcEntities = v
	return s
}

func (s *LineageRelationRegisterBulkVO) Validate() error {
	return dara.Validate(s)
}
