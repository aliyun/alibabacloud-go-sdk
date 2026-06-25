// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelationship interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]interface{}) *Relationship
	GetAttributes() map[string]interface{}
	SetDataChannel(v string) *Relationship
	GetDataChannel() *string
	SetRelationshipGuid(v string) *Relationship
	GetRelationshipGuid() *string
	SetRelationshipType(v string) *Relationship
	GetRelationshipType() *string
}

type Relationship struct {
	// A collection of key-value pairs providing additional details about the relationship.
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The channel or system through which the relationship is established.
	DataChannel *string `json:"DataChannel,omitempty" xml:"DataChannel,omitempty"`
	// The globally unique identifier (GUID) for the relationship.
	RelationshipGuid *string `json:"RelationshipGuid,omitempty" xml:"RelationshipGuid,omitempty"`
	// Specifies the type of relationship.
	RelationshipType *string `json:"RelationshipType,omitempty" xml:"RelationshipType,omitempty"`
}

func (s Relationship) String() string {
	return dara.Prettify(s)
}

func (s Relationship) GoString() string {
	return s.String()
}

func (s *Relationship) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *Relationship) GetDataChannel() *string {
	return s.DataChannel
}

func (s *Relationship) GetRelationshipGuid() *string {
	return s.RelationshipGuid
}

func (s *Relationship) GetRelationshipType() *string {
	return s.RelationshipType
}

func (s *Relationship) SetAttributes(v map[string]interface{}) *Relationship {
	s.Attributes = v
	return s
}

func (s *Relationship) SetDataChannel(v string) *Relationship {
	s.DataChannel = &v
	return s
}

func (s *Relationship) SetRelationshipGuid(v string) *Relationship {
	s.RelationshipGuid = &v
	return s
}

func (s *Relationship) SetRelationshipType(v string) *Relationship {
	s.RelationshipType = &v
	return s
}

func (s *Relationship) Validate() error {
	return dara.Validate(s)
}
