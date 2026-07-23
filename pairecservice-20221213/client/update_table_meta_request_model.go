// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTableMetaRequest
	GetDescription() *string
	SetFields(v []*UpdateTableMetaRequestFields) *UpdateTableMetaRequest
	GetFields() []*UpdateTableMetaRequestFields
	SetInstanceId(v string) *UpdateTableMetaRequest
	GetInstanceId() *string
	SetModule(v string) *UpdateTableMetaRequest
	GetModule() *string
	SetName(v string) *UpdateTableMetaRequest
	GetName() *string
	SetResourceId(v string) *UpdateTableMetaRequest
	GetResourceId() *string
	SetTableName(v string) *UpdateTableMetaRequest
	GetTableName() *string
}

type UpdateTableMetaRequest struct {
	// The description of the data table.
	//
	// example:
	//
	// this is a test table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fields of the data table.
	//
	// This parameter is required.
	Fields []*UpdateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The instance ID. You can call the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module to which the data table belongs.
	//
	// - ABTest: a data table for A/B testing.
	//
	// - ExperimentTool: a data table for experiment tools.
	//
	// - DataDiagnosis: a data table for data diagnosis.
	//
	// This parameter is required.
	//
	// example:
	//
	// ABTest
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The name of the data table.
	//
	// This parameter is required.
	//
	// example:
	//
	// table_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource ID. You can call the [ListInstanceResource](https://help.aliyun.com/document_detail/2672886.html) operation to obtain the resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// reso-2s416t***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the table in the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// table_mysql
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s UpdateTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTableMetaRequest) GetFields() []*UpdateTableMetaRequestFields {
	return s.Fields
}

func (s *UpdateTableMetaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTableMetaRequest) GetModule() *string {
	return s.Module
}

func (s *UpdateTableMetaRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTableMetaRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateTableMetaRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateTableMetaRequest) SetDescription(v string) *UpdateTableMetaRequest {
	s.Description = &v
	return s
}

func (s *UpdateTableMetaRequest) SetFields(v []*UpdateTableMetaRequestFields) *UpdateTableMetaRequest {
	s.Fields = v
	return s
}

func (s *UpdateTableMetaRequest) SetInstanceId(v string) *UpdateTableMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTableMetaRequest) SetModule(v string) *UpdateTableMetaRequest {
	s.Module = &v
	return s
}

func (s *UpdateTableMetaRequest) SetName(v string) *UpdateTableMetaRequest {
	s.Name = &v
	return s
}

func (s *UpdateTableMetaRequest) SetResourceId(v string) *UpdateTableMetaRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateTableMetaRequest) SetTableName(v string) *UpdateTableMetaRequest {
	s.TableName = &v
	return s
}

func (s *UpdateTableMetaRequest) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTableMetaRequestFields struct {
	// The data type of the field.
	//
	// example:
	//
	// primary_key
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Indicates whether the field is a dimension field.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	IsDimensionField *bool `json:"IsDimensionField,omitempty" xml:"IsDimensionField,omitempty"`
	// Indicates whether the field is a partition field.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsPartitionField *string `json:"IsPartitionField,omitempty" xml:"IsPartitionField,omitempty"`
	// The business meaning of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// 名称字段
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// The field name.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// BIGINT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTableMetaRequestFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableMetaRequestFields) GoString() string {
	return s.String()
}

func (s *UpdateTableMetaRequestFields) GetDataType() *string {
	return s.DataType
}

func (s *UpdateTableMetaRequestFields) GetIsDimensionField() *bool {
	return s.IsDimensionField
}

func (s *UpdateTableMetaRequestFields) GetIsPartitionField() *string {
	return s.IsPartitionField
}

func (s *UpdateTableMetaRequestFields) GetMeaning() *string {
	return s.Meaning
}

func (s *UpdateTableMetaRequestFields) GetName() *string {
	return s.Name
}

func (s *UpdateTableMetaRequestFields) GetType() *string {
	return s.Type
}

func (s *UpdateTableMetaRequestFields) SetDataType(v string) *UpdateTableMetaRequestFields {
	s.DataType = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetIsDimensionField(v bool) *UpdateTableMetaRequestFields {
	s.IsDimensionField = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetIsPartitionField(v string) *UpdateTableMetaRequestFields {
	s.IsPartitionField = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetMeaning(v string) *UpdateTableMetaRequestFields {
	s.Meaning = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetName(v string) *UpdateTableMetaRequestFields {
	s.Name = &v
	return s
}

func (s *UpdateTableMetaRequestFields) SetType(v string) *UpdateTableMetaRequestFields {
	s.Type = &v
	return s
}

func (s *UpdateTableMetaRequestFields) Validate() error {
	return dara.Validate(s)
}
