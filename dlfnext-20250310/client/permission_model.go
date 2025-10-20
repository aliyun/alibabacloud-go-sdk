// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPermission interface {
	dara.Model
	String() string
	GoString() string
	SetAccess(v string) *Permission
	GetAccess() *string
	SetColumns(v *PermissionColumns) *Permission
	GetColumns() *PermissionColumns
	SetDatabase(v string) *Permission
	GetDatabase() *string
	SetFunction(v string) *Permission
	GetFunction() *string
	SetPrincipal(v string) *Permission
	GetPrincipal() *string
	SetResourceType(v string) *Permission
	GetResourceType() *string
	SetTable(v string) *Permission
	GetTable() *string
	SetView(v string) *Permission
	GetView() *string
}

type Permission struct {
	Access       *string            `json:"access,omitempty" xml:"access,omitempty"`
	Columns      *PermissionColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Struct"`
	Database     *string            `json:"database,omitempty" xml:"database,omitempty"`
	Function     *string            `json:"function,omitempty" xml:"function,omitempty"`
	Principal    *string            `json:"principal,omitempty" xml:"principal,omitempty"`
	ResourceType *string            `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Table        *string            `json:"table,omitempty" xml:"table,omitempty"`
	View         *string            `json:"view,omitempty" xml:"view,omitempty"`
}

func (s Permission) String() string {
	return dara.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) GetAccess() *string {
	return s.Access
}

func (s *Permission) GetColumns() *PermissionColumns {
	return s.Columns
}

func (s *Permission) GetDatabase() *string {
	return s.Database
}

func (s *Permission) GetFunction() *string {
	return s.Function
}

func (s *Permission) GetPrincipal() *string {
	return s.Principal
}

func (s *Permission) GetResourceType() *string {
	return s.ResourceType
}

func (s *Permission) GetTable() *string {
	return s.Table
}

func (s *Permission) GetView() *string {
	return s.View
}

func (s *Permission) SetAccess(v string) *Permission {
	s.Access = &v
	return s
}

func (s *Permission) SetColumns(v *PermissionColumns) *Permission {
	s.Columns = v
	return s
}

func (s *Permission) SetDatabase(v string) *Permission {
	s.Database = &v
	return s
}

func (s *Permission) SetFunction(v string) *Permission {
	s.Function = &v
	return s
}

func (s *Permission) SetPrincipal(v string) *Permission {
	s.Principal = &v
	return s
}

func (s *Permission) SetResourceType(v string) *Permission {
	s.ResourceType = &v
	return s
}

func (s *Permission) SetTable(v string) *Permission {
	s.Table = &v
	return s
}

func (s *Permission) SetView(v string) *Permission {
	s.View = &v
	return s
}

func (s *Permission) Validate() error {
	if s.Columns != nil {
		if err := s.Columns.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PermissionColumns struct {
	ColumnNames         []*string `json:"columnNames,omitempty" xml:"columnNames,omitempty" type:"Repeated"`
	ExcludedColumnNames []*string `json:"excludedColumnNames,omitempty" xml:"excludedColumnNames,omitempty" type:"Repeated"`
}

func (s PermissionColumns) String() string {
	return dara.Prettify(s)
}

func (s PermissionColumns) GoString() string {
	return s.String()
}

func (s *PermissionColumns) GetColumnNames() []*string {
	return s.ColumnNames
}

func (s *PermissionColumns) GetExcludedColumnNames() []*string {
	return s.ExcludedColumnNames
}

func (s *PermissionColumns) SetColumnNames(v []*string) *PermissionColumns {
	s.ColumnNames = v
	return s
}

func (s *PermissionColumns) SetExcludedColumnNames(v []*string) *PermissionColumns {
	s.ExcludedColumnNames = v
	return s
}

func (s *PermissionColumns) Validate() error {
	return dara.Validate(s)
}
