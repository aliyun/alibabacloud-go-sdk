// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaDatabaseObject interface {
	dara.Model
	String() string
	GoString() string
	SetObjectName(v string) *OneMetaDatabaseObject
	GetObjectName() *string
	SetObjectQualifiedName(v string) *OneMetaDatabaseObject
	GetObjectQualifiedName() *string
	SetObjectType(v string) *OneMetaDatabaseObject
	GetObjectType() *string
}

type OneMetaDatabaseObject struct {
	ObjectName          *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectQualifiedName *string `json:"ObjectQualifiedName,omitempty" xml:"ObjectQualifiedName,omitempty"`
	ObjectType          *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s OneMetaDatabaseObject) String() string {
	return dara.Prettify(s)
}

func (s OneMetaDatabaseObject) GoString() string {
	return s.String()
}

func (s *OneMetaDatabaseObject) GetObjectName() *string {
	return s.ObjectName
}

func (s *OneMetaDatabaseObject) GetObjectQualifiedName() *string {
	return s.ObjectQualifiedName
}

func (s *OneMetaDatabaseObject) GetObjectType() *string {
	return s.ObjectType
}

func (s *OneMetaDatabaseObject) SetObjectName(v string) *OneMetaDatabaseObject {
	s.ObjectName = &v
	return s
}

func (s *OneMetaDatabaseObject) SetObjectQualifiedName(v string) *OneMetaDatabaseObject {
	s.ObjectQualifiedName = &v
	return s
}

func (s *OneMetaDatabaseObject) SetObjectType(v string) *OneMetaDatabaseObject {
	s.ObjectType = &v
	return s
}

func (s *OneMetaDatabaseObject) Validate() error {
	return dara.Validate(s)
}
