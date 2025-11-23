// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakeTableRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakeTableRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakeTableRequest
	GetDbName() *string
	SetTableInput(v *OpenStructDLTableInput) *CreateDataLakeTableRequest
	GetTableInput() *OpenStructDLTableInput
	SetTid(v int64) *CreateDataLakeTableRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakeTableRequest
	GetWorkspaceId() *int64
}

type CreateDataLakeTableRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The region where the data lake resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The information about the table.
	//
	// This parameter is required.
	TableInput *OpenStructDLTableInput `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	// The ID of the tenant.
	//
	// > You can move the pointer over the profile picture in the upper-right corner of the DMS console to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataLakeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeTableRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakeTableRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakeTableRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakeTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakeTableRequest) GetTableInput() *OpenStructDLTableInput {
	return s.TableInput
}

func (s *CreateDataLakeTableRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakeTableRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakeTableRequest) SetCatalogName(v string) *CreateDataLakeTableRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakeTableRequest) SetDataRegion(v string) *CreateDataLakeTableRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakeTableRequest) SetDbName(v string) *CreateDataLakeTableRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakeTableRequest) SetTableInput(v *OpenStructDLTableInput) *CreateDataLakeTableRequest {
	s.TableInput = v
	return s
}

func (s *CreateDataLakeTableRequest) SetTid(v int64) *CreateDataLakeTableRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakeTableRequest) SetWorkspaceId(v int64) *CreateDataLakeTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakeTableRequest) Validate() error {
	if s.TableInput != nil {
		if err := s.TableInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}
