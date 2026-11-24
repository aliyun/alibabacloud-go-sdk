// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageRelation interface {
	dara.Model
	String() string
	GoString() string
	SetDestEntityQualifiedName(v string) *LineageRelation
	GetDestEntityQualifiedName() *string
	SetRelationshipGuid(v string) *LineageRelation
	GetRelationshipGuid() *string
	SetSrcEntityQualifiedName(v string) *LineageRelation
	GetSrcEntityQualifiedName() *string
}

type LineageRelation struct {
	// The unique qualified name of the destination lineage entity.
	//
	// example:
	//
	// maxcompute-table.demo_project.training_output
	DestEntityQualifiedName *string `json:"DestEntityQualifiedName,omitempty" xml:"DestEntityQualifiedName,omitempty"`
	// The globally unique identifier of the lineage relationship.
	//
	// example:
	//
	// 0f94c587-e926-48c8-9efe-bf9f8d9aa5c2
	RelationshipGuid *string `json:"RelationshipGuid,omitempty" xml:"RelationshipGuid,omitempty"`
	// The unique qualified name of the source lineage entity.
	//
	// example:
	//
	// maxcompute-table.demo_project.training_data
	SrcEntityQualifiedName *string `json:"SrcEntityQualifiedName,omitempty" xml:"SrcEntityQualifiedName,omitempty"`
}

func (s LineageRelation) String() string {
	return dara.Prettify(s)
}

func (s LineageRelation) GoString() string {
	return s.String()
}

func (s *LineageRelation) GetDestEntityQualifiedName() *string {
	return s.DestEntityQualifiedName
}

func (s *LineageRelation) GetRelationshipGuid() *string {
	return s.RelationshipGuid
}

func (s *LineageRelation) GetSrcEntityQualifiedName() *string {
	return s.SrcEntityQualifiedName
}

func (s *LineageRelation) SetDestEntityQualifiedName(v string) *LineageRelation {
	s.DestEntityQualifiedName = &v
	return s
}

func (s *LineageRelation) SetRelationshipGuid(v string) *LineageRelation {
	s.RelationshipGuid = &v
	return s
}

func (s *LineageRelation) SetSrcEntityQualifiedName(v string) *LineageRelation {
	s.SrcEntityQualifiedName = &v
	return s
}

func (s *LineageRelation) Validate() error {
	return dara.Validate(s)
}
