// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleAclOnObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectName(v string) *GetRoleAclOnObjectRequest
	GetObjectName() *string
	SetObjectType(v string) *GetRoleAclOnObjectRequest
	GetObjectType() *string
}

type GetRoleAclOnObjectRequest struct {
	// The name of the object.
	//
	// This parameter is required.
	//
	// example:
	//
	// tableA
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The type of the object.
	//
	// This parameter is required.
	//
	// example:
	//
	// table
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s GetRoleAclOnObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoleAclOnObjectRequest) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectRequest) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetRoleAclOnObjectRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetRoleAclOnObjectRequest) SetObjectName(v string) *GetRoleAclOnObjectRequest {
	s.ObjectName = &v
	return s
}

func (s *GetRoleAclOnObjectRequest) SetObjectType(v string) *GetRoleAclOnObjectRequest {
	s.ObjectType = &v
	return s
}

func (s *GetRoleAclOnObjectRequest) Validate() error {
	return dara.Validate(s)
}
