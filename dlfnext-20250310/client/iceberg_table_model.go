// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIcebergTable interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *IcebergTable
	GetCreatedAt() *int64
	SetCreatedBy(v string) *IcebergTable
	GetCreatedBy() *string
	SetIcebergTableMetadata(v *IcebergTableMetadata) *IcebergTable
	GetIcebergTableMetadata() *IcebergTableMetadata
	SetId(v string) *IcebergTable
	GetId() *string
	SetName(v string) *IcebergTable
	GetName() *string
	SetOwner(v string) *IcebergTable
	GetOwner() *string
	SetPath(v string) *IcebergTable
	GetPath() *string
	SetUpdatedAt(v int64) *IcebergTable
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *IcebergTable
	GetUpdatedBy() *string
	SetVersion(v int64) *IcebergTable
	GetVersion() *int64
}

type IcebergTable struct {
	CreatedAt            *int64                `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy            *string               `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	IcebergTableMetadata *IcebergTableMetadata `json:"icebergTableMetadata,omitempty" xml:"icebergTableMetadata,omitempty"`
	Id                   *string               `json:"id,omitempty" xml:"id,omitempty"`
	Name                 *string               `json:"name,omitempty" xml:"name,omitempty"`
	Owner                *string               `json:"owner,omitempty" xml:"owner,omitempty"`
	Path                 *string               `json:"path,omitempty" xml:"path,omitempty"`
	UpdatedAt            *int64                `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy            *string               `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
	Version              *int64                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s IcebergTable) String() string {
	return dara.Prettify(s)
}

func (s IcebergTable) GoString() string {
	return s.String()
}

func (s *IcebergTable) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *IcebergTable) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *IcebergTable) GetIcebergTableMetadata() *IcebergTableMetadata {
	return s.IcebergTableMetadata
}

func (s *IcebergTable) GetId() *string {
	return s.Id
}

func (s *IcebergTable) GetName() *string {
	return s.Name
}

func (s *IcebergTable) GetOwner() *string {
	return s.Owner
}

func (s *IcebergTable) GetPath() *string {
	return s.Path
}

func (s *IcebergTable) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *IcebergTable) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *IcebergTable) GetVersion() *int64 {
	return s.Version
}

func (s *IcebergTable) SetCreatedAt(v int64) *IcebergTable {
	s.CreatedAt = &v
	return s
}

func (s *IcebergTable) SetCreatedBy(v string) *IcebergTable {
	s.CreatedBy = &v
	return s
}

func (s *IcebergTable) SetIcebergTableMetadata(v *IcebergTableMetadata) *IcebergTable {
	s.IcebergTableMetadata = v
	return s
}

func (s *IcebergTable) SetId(v string) *IcebergTable {
	s.Id = &v
	return s
}

func (s *IcebergTable) SetName(v string) *IcebergTable {
	s.Name = &v
	return s
}

func (s *IcebergTable) SetOwner(v string) *IcebergTable {
	s.Owner = &v
	return s
}

func (s *IcebergTable) SetPath(v string) *IcebergTable {
	s.Path = &v
	return s
}

func (s *IcebergTable) SetUpdatedAt(v int64) *IcebergTable {
	s.UpdatedAt = &v
	return s
}

func (s *IcebergTable) SetUpdatedBy(v string) *IcebergTable {
	s.UpdatedBy = &v
	return s
}

func (s *IcebergTable) SetVersion(v int64) *IcebergTable {
	s.Version = &v
	return s
}

func (s *IcebergTable) Validate() error {
	if s.IcebergTableMetadata != nil {
		if err := s.IcebergTableMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
