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
  // This parameter is required.
  // 
  // example:
  // 
  // 1000001
  BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
  // example:
  // 
  // 2eb6f9****
  FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
  // example:
  // 
  // 10000
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // example:
  // 
  // dw
  ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
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

