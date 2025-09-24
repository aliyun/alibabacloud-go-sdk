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
	DestEntityQualifiedName *string `json:"DestEntityQualifiedName,omitempty" xml:"DestEntityQualifiedName,omitempty"`
	RelationshipGuid        *string `json:"RelationshipGuid,omitempty" xml:"RelationshipGuid,omitempty"`
	SrcEntityQualifiedName  *string `json:"SrcEntityQualifiedName,omitempty" xml:"SrcEntityQualifiedName,omitempty"`
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
