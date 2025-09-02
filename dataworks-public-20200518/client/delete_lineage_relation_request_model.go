// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestEntityQualifiedName(v string) *DeleteLineageRelationRequest
	GetDestEntityQualifiedName() *string
	SetRelationshipGuid(v string) *DeleteLineageRelationRequest
	GetRelationshipGuid() *string
	SetRelationshipType(v string) *DeleteLineageRelationRequest
	GetRelationshipType() *string
	SetSrcEntityQualifiedName(v string) *DeleteLineageRelationRequest
	GetSrcEntityQualifiedName() *string
}

type DeleteLineageRelationRequest struct {
	// Destination entity unique identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// custom-report.report123
	DestEntityQualifiedName *string `json:"DestEntityQualifiedName,omitempty" xml:"DestEntityQualifiedName,omitempty"`
	// Lineage relationship unique identifier
	//
	// example:
	//
	// dfazcdfdfccdedd
	RelationshipGuid *string `json:"RelationshipGuid,omitempty" xml:"RelationshipGuid,omitempty"`
	// Relationship type
	//
	// example:
	//
	// sql
	RelationshipType *string `json:"RelationshipType,omitempty" xml:"RelationshipType,omitempty"`
	// Source entity unique identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.project.table
	SrcEntityQualifiedName *string `json:"SrcEntityQualifiedName,omitempty" xml:"SrcEntityQualifiedName,omitempty"`
}

func (s DeleteLineageRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationRequest) GetDestEntityQualifiedName() *string {
	return s.DestEntityQualifiedName
}

func (s *DeleteLineageRelationRequest) GetRelationshipGuid() *string {
	return s.RelationshipGuid
}

func (s *DeleteLineageRelationRequest) GetRelationshipType() *string {
	return s.RelationshipType
}

func (s *DeleteLineageRelationRequest) GetSrcEntityQualifiedName() *string {
	return s.SrcEntityQualifiedName
}

func (s *DeleteLineageRelationRequest) SetDestEntityQualifiedName(v string) *DeleteLineageRelationRequest {
	s.DestEntityQualifiedName = &v
	return s
}

func (s *DeleteLineageRelationRequest) SetRelationshipGuid(v string) *DeleteLineageRelationRequest {
	s.RelationshipGuid = &v
	return s
}

func (s *DeleteLineageRelationRequest) SetRelationshipType(v string) *DeleteLineageRelationRequest {
	s.RelationshipType = &v
	return s
}

func (s *DeleteLineageRelationRequest) SetSrcEntityQualifiedName(v string) *DeleteLineageRelationRequest {
	s.SrcEntityQualifiedName = &v
	return s
}

func (s *DeleteLineageRelationRequest) Validate() error {
	return dara.Validate(s)
}
