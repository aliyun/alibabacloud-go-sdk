// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelationshipVO interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *RelationshipVO
	GetAttributes() map[string]*string
	SetRelationshipGuid(v string) *RelationshipVO
	GetRelationshipGuid() *string
	SetRelationshipType(v string) *RelationshipVO
	GetRelationshipType() *string
}

type RelationshipVO struct {
	Attributes       map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	RelationshipGuid *string            `json:"RelationshipGuid,omitempty" xml:"RelationshipGuid,omitempty"`
	// example:
	//
	// sql
	RelationshipType *string `json:"RelationshipType,omitempty" xml:"RelationshipType,omitempty"`
}

func (s RelationshipVO) String() string {
	return dara.Prettify(s)
}

func (s RelationshipVO) GoString() string {
	return s.String()
}

func (s *RelationshipVO) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *RelationshipVO) GetRelationshipGuid() *string {
	return s.RelationshipGuid
}

func (s *RelationshipVO) GetRelationshipType() *string {
	return s.RelationshipType
}

func (s *RelationshipVO) SetAttributes(v map[string]*string) *RelationshipVO {
	s.Attributes = v
	return s
}

func (s *RelationshipVO) SetRelationshipGuid(v string) *RelationshipVO {
	s.RelationshipGuid = &v
	return s
}

func (s *RelationshipVO) SetRelationshipType(v string) *RelationshipVO {
	s.RelationshipType = &v
	return s
}

func (s *RelationshipVO) Validate() error {
	return dara.Validate(s)
}
