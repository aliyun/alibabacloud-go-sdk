// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareResource interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *ShareResource
	GetCreatedAt() *int64
	SetCreatedBy(v string) *ShareResource
	GetCreatedBy() *string
	SetDatabaseName(v string) *ShareResource
	GetDatabaseName() *string
	SetShareType(v string) *ShareResource
	GetShareType() *string
	SetTableName(v string) *ShareResource
	GetTableName() *string
	SetUpdatedAt(v int64) *ShareResource
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *ShareResource
	GetUpdatedBy() *string
}

type ShareResource struct {
	CreatedAt    *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy    *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	ShareType    *string `json:"shareType,omitempty" xml:"shareType,omitempty"`
	TableName    *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	UpdatedAt    *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy    *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s ShareResource) String() string {
	return dara.Prettify(s)
}

func (s ShareResource) GoString() string {
	return s.String()
}

func (s *ShareResource) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ShareResource) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ShareResource) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ShareResource) GetShareType() *string {
	return s.ShareType
}

func (s *ShareResource) GetTableName() *string {
	return s.TableName
}

func (s *ShareResource) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ShareResource) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ShareResource) SetCreatedAt(v int64) *ShareResource {
	s.CreatedAt = &v
	return s
}

func (s *ShareResource) SetCreatedBy(v string) *ShareResource {
	s.CreatedBy = &v
	return s
}

func (s *ShareResource) SetDatabaseName(v string) *ShareResource {
	s.DatabaseName = &v
	return s
}

func (s *ShareResource) SetShareType(v string) *ShareResource {
	s.ShareType = &v
	return s
}

func (s *ShareResource) SetTableName(v string) *ShareResource {
	s.TableName = &v
	return s
}

func (s *ShareResource) SetUpdatedAt(v int64) *ShareResource {
	s.UpdatedAt = &v
	return s
}

func (s *ShareResource) SetUpdatedBy(v string) *ShareResource {
	s.UpdatedBy = &v
	return s
}

func (s *ShareResource) Validate() error {
	return dara.Validate(s)
}
