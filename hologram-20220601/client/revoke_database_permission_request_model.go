// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDatabasePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *RevokeDatabasePermissionRequest
	GetDatabaseName() *string
	SetPrivileges(v []*string) *RevokeDatabasePermissionRequest
	GetPrivileges() []*string
	SetUserName(v string) *RevokeDatabasePermissionRequest
	GetUserName() *string
}

type RevokeDatabasePermissionRequest struct {
	// example:
	//
	// test_db
	DatabaseName *string   `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	Privileges   []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// example:
	//
	// p4_1234xxxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s RevokeDatabasePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeDatabasePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeDatabasePermissionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *RevokeDatabasePermissionRequest) GetPrivileges() []*string {
	return s.Privileges
}

func (s *RevokeDatabasePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *RevokeDatabasePermissionRequest) SetDatabaseName(v string) *RevokeDatabasePermissionRequest {
	s.DatabaseName = &v
	return s
}

func (s *RevokeDatabasePermissionRequest) SetPrivileges(v []*string) *RevokeDatabasePermissionRequest {
	s.Privileges = v
	return s
}

func (s *RevokeDatabasePermissionRequest) SetUserName(v string) *RevokeDatabasePermissionRequest {
	s.UserName = &v
	return s
}

func (s *RevokeDatabasePermissionRequest) Validate() error {
	return dara.Validate(s)
}
