// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstablishRelationTableToBusinessRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBusinessId(v string) *EstablishRelationTableToBusinessRequest
  GetBusinessId() *string 
  SetFolderId(v string) *EstablishRelationTableToBusinessRequest
  GetFolderId() *string 
  SetProjectId(v int64) *EstablishRelationTableToBusinessRequest
  GetProjectId() *int64 
  SetProjectIdentifier(v string) *EstablishRelationTableToBusinessRequest
  GetProjectIdentifier() *string 
  SetTableGuid(v string) *EstablishRelationTableToBusinessRequest
  GetTableGuid() *string 
}

type EstablishRelationTableToBusinessRequest struct {
  // The ID of the workflow. You can call the [ListBusiness](https://help.aliyun.com/document_detail/173945.html) operation to query the ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1000001
  BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
  // The ID of the folder. You can call the [GetFolder](https://help.aliyun.com/document_detail/173952.html) or [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to query the ID.
  // 
  // example:
  // 
  // 2eb6f9****
  FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
  // The ID of the DataWorks workspace. You can click the Workspace Manage icon in the upper-right corner of the DataStudio page to go to the Workspace Management page and view the workspace ID.
  // 
  // example:
  // 
  // 10000
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The unique identifier of the DataWorks workspace. You can click the identifier in the upper-left corner of the DataStudio page to switch to another workspace.
  // 
  // You must specify either this parameter or ProjectId to determine the DataWorks workspace to which the operation is applied.
  // 
  // example:
  // 
  // dw
  ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
  // The universally unique identifier (UUID) of the table. You can call the [SearchMetaTables](https://help.aliyun.com/document_detail/173919.html) operation to query the UUID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // odps.dw_project.tb1
  TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s EstablishRelationTableToBusinessRequest) String() string {
  return dara.Prettify(s)
}

func (s EstablishRelationTableToBusinessRequest) GoString() string {
  return s.String()
}

func (s *EstablishRelationTableToBusinessRequest) GetBusinessId() *string  {
  return s.BusinessId
}

func (s *EstablishRelationTableToBusinessRequest) GetFolderId() *string  {
  return s.FolderId
}

func (s *EstablishRelationTableToBusinessRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *EstablishRelationTableToBusinessRequest) GetProjectIdentifier() *string  {
  return s.ProjectIdentifier
}

func (s *EstablishRelationTableToBusinessRequest) GetTableGuid() *string  {
  return s.TableGuid
}

func (s *EstablishRelationTableToBusinessRequest) SetBusinessId(v string) *EstablishRelationTableToBusinessRequest {
  s.BusinessId = &v
  return s
}

func (s *EstablishRelationTableToBusinessRequest) SetFolderId(v string) *EstablishRelationTableToBusinessRequest {
  s.FolderId = &v
  return s
}

func (s *EstablishRelationTableToBusinessRequest) SetProjectId(v int64) *EstablishRelationTableToBusinessRequest {
  s.ProjectId = &v
  return s
}

func (s *EstablishRelationTableToBusinessRequest) SetProjectIdentifier(v string) *EstablishRelationTableToBusinessRequest {
  s.ProjectIdentifier = &v
  return s
}

func (s *EstablishRelationTableToBusinessRequest) SetTableGuid(v string) *EstablishRelationTableToBusinessRequest {
  s.TableGuid = &v
  return s
}

func (s *EstablishRelationTableToBusinessRequest) Validate() error {
  return dara.Validate(s)
}

