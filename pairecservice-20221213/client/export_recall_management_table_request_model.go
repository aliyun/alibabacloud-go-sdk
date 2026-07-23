// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecallManagementTableRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *ExportRecallManagementTableRequest
  GetInstanceId() *string 
  SetMaxcomputeProjectName(v string) *ExportRecallManagementTableRequest
  GetMaxcomputeProjectName() *string 
  SetMaxcomputeSchema(v string) *ExportRecallManagementTableRequest
  GetMaxcomputeSchema() *string 
  SetMaxcomputeTableName(v string) *ExportRecallManagementTableRequest
  GetMaxcomputeTableName() *string 
  SetPartitions(v map[string]*string) *ExportRecallManagementTableRequest
  GetPartitions() map[string]*string 
  SetRecallManagementTableVersionId(v string) *ExportRecallManagementTableRequest
  GetRecallManagementTableVersionId() *string 
}

type ExportRecallManagementTableRequest struct {
  // The instance ID.
  // 
  // example:
  // 
  // pairec-test1
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The MaxCompute project name.
  // 
  // example:
  // 
  // test
  MaxcomputeProjectName *string `json:"MaxcomputeProjectName,omitempty" xml:"MaxcomputeProjectName,omitempty"`
  // The MaxCompute project schema.
  // 
  // example:
  // 
  // default
  MaxcomputeSchema *string `json:"MaxcomputeSchema,omitempty" xml:"MaxcomputeSchema,omitempty"`
  // The name of the destination table in MaxCompute.
  // 
  // example:
  // 
  // table-1
  MaxcomputeTableName *string `json:"MaxcomputeTableName,omitempty" xml:"MaxcomputeTableName,omitempty"`
  // The table partitions to export. The value must be a JSON object where each key is a partition field and its value is the corresponding partition value.
  Partitions map[string]*string `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
  // The version ID of the recall management table. This parameter defaults to the current published version.
  // 
  // example:
  // 
  // 20250213000000
  RecallManagementTableVersionId *string `json:"RecallManagementTableVersionId,omitempty" xml:"RecallManagementTableVersionId,omitempty"`
}

func (s ExportRecallManagementTableRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportRecallManagementTableRequest) GoString() string {
  return s.String()
}

func (s *ExportRecallManagementTableRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportRecallManagementTableRequest) GetMaxcomputeProjectName() *string  {
  return s.MaxcomputeProjectName
}

func (s *ExportRecallManagementTableRequest) GetMaxcomputeSchema() *string  {
  return s.MaxcomputeSchema
}

func (s *ExportRecallManagementTableRequest) GetMaxcomputeTableName() *string  {
  return s.MaxcomputeTableName
}

func (s *ExportRecallManagementTableRequest) GetPartitions() map[string]*string  {
  return s.Partitions
}

func (s *ExportRecallManagementTableRequest) GetRecallManagementTableVersionId() *string  {
  return s.RecallManagementTableVersionId
}

func (s *ExportRecallManagementTableRequest) SetInstanceId(v string) *ExportRecallManagementTableRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportRecallManagementTableRequest) SetMaxcomputeProjectName(v string) *ExportRecallManagementTableRequest {
  s.MaxcomputeProjectName = &v
  return s
}

func (s *ExportRecallManagementTableRequest) SetMaxcomputeSchema(v string) *ExportRecallManagementTableRequest {
  s.MaxcomputeSchema = &v
  return s
}

func (s *ExportRecallManagementTableRequest) SetMaxcomputeTableName(v string) *ExportRecallManagementTableRequest {
  s.MaxcomputeTableName = &v
  return s
}

func (s *ExportRecallManagementTableRequest) SetPartitions(v map[string]*string) *ExportRecallManagementTableRequest {
  s.Partitions = v
  return s
}

func (s *ExportRecallManagementTableRequest) SetRecallManagementTableVersionId(v string) *ExportRecallManagementTableRequest {
  s.RecallManagementTableVersionId = &v
  return s
}

func (s *ExportRecallManagementTableRequest) Validate() error {
  return dara.Validate(s)
}

