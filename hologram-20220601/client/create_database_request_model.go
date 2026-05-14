// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *CreateDatabaseRequest
	GetDatabaseName() *string
	SetPermissionModel(v string) *CreateDatabaseRequest
	GetPermissionModel() *string
}

type CreateDatabaseRequest struct {
	// example:
	//
	// my_db
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// example:
	//
	// SPM
	PermissionModel *string `json:"permissionModel,omitempty" xml:"permissionModel,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CreateDatabaseRequest) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *CreateDatabaseRequest) SetDatabaseName(v string) *CreateDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseRequest) SetPermissionModel(v string) *CreateDatabaseRequest {
	s.PermissionModel = &v
	return s
}

func (s *CreateDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
