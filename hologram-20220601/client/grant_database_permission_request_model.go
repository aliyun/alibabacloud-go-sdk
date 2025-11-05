// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDatabasePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *GrantDatabasePermissionRequest
	GetDatabaseName() *string
	SetPrivileges(v []*string) *GrantDatabasePermissionRequest
	GetPrivileges() []*string
	SetUserName(v string) *GrantDatabasePermissionRequest
	GetUserName() *string
}

type GrantDatabasePermissionRequest struct {
	// example:
	//
	// test_db
	DatabaseName *string   `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	Privileges   []*string `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	// example:
	//
	// p4_123444xxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GrantDatabasePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantDatabasePermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantDatabasePermissionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GrantDatabasePermissionRequest) GetPrivileges() []*string {
	return s.Privileges
}

func (s *GrantDatabasePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *GrantDatabasePermissionRequest) SetDatabaseName(v string) *GrantDatabasePermissionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GrantDatabasePermissionRequest) SetPrivileges(v []*string) *GrantDatabasePermissionRequest {
	s.Privileges = v
	return s
}

func (s *GrantDatabasePermissionRequest) SetUserName(v string) *GrantDatabasePermissionRequest {
	s.UserName = &v
	return s
}

func (s *GrantDatabasePermissionRequest) Validate() error {
	return dara.Validate(s)
}
