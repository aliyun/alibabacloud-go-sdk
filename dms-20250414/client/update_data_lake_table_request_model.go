// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakeTableRequest
	GetCatalogName() *string
	SetDbName(v string) *UpdateDataLakeTableRequest
	GetDbName() *string
	SetTableInput(v *DLTableInput) *UpdateDataLakeTableRequest
	GetTableInput() *DLTableInput
	SetTableName(v string) *UpdateDataLakeTableRequest
	GetTableName() *string
	SetTid(v int64) *UpdateDataLakeTableRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakeTableRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakeTableRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The name of the database.
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
	TableInput *DLTableInput `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	// The name of the table to update. If you do not want to change the table name, set this parameter to the same value as the Name parameter in TableInput.
	//
	// example:
	//
	// 100g_customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The tenant ID.
	//
	// > Hover over your profile picture in the upper-right corner of the DMS console to obtain the tenant ID. For details, see [View tenant information](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataLakeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeTableRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakeTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakeTableRequest) GetTableInput() *DLTableInput {
	return s.TableInput
}

func (s *UpdateDataLakeTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateDataLakeTableRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakeTableRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakeTableRequest) SetCatalogName(v string) *UpdateDataLakeTableRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakeTableRequest) SetDbName(v string) *UpdateDataLakeTableRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakeTableRequest) SetTableInput(v *DLTableInput) *UpdateDataLakeTableRequest {
	s.TableInput = v
	return s
}

func (s *UpdateDataLakeTableRequest) SetTableName(v string) *UpdateDataLakeTableRequest {
	s.TableName = &v
	return s
}

func (s *UpdateDataLakeTableRequest) SetTid(v int64) *UpdateDataLakeTableRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakeTableRequest) SetWorkspaceId(v int64) *UpdateDataLakeTableRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakeTableRequest) Validate() error {
	if s.TableInput != nil {
		if err := s.TableInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}
