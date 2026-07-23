// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTableMetaRequest
	GetDescription() *string
	SetFields(v []*CreateTableMetaRequestFields) *CreateTableMetaRequest
	GetFields() []*CreateTableMetaRequestFields
	SetInstanceId(v string) *CreateTableMetaRequest
	GetInstanceId() *string
	SetModule(v string) *CreateTableMetaRequest
	GetModule() *string
	SetName(v string) *CreateTableMetaRequest
	GetName() *string
	SetResourceId(v string) *CreateTableMetaRequest
	GetResourceId() *string
	SetTableName(v string) *CreateTableMetaRequest
	GetTableName() *string
}

type CreateTableMetaRequest struct {
	// The description of the data table.
	//
	// This parameter is required.
	//
	// example:
	//
	// this is a test table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of fields.
	//
	// This parameter is required.
	Fields []*CreateTableMetaRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The instance ID. To get this ID, call the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The module to which the data table belongs.
	//
	// - ABTest: The data table for A/B testing.
	//
	// - ExperimentTool: The data table for the experiment tool.
	//
	// - DataDiagnosis: The data table for data diagnosis.
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
	// The resource ID. To get this ID, call the ListInstanceResources operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// reso-2s416t146ffjc3yefx
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

func (s CreateTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateTableMetaRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTableMetaRequest) GetFields() []*CreateTableMetaRequestFields {
	return s.Fields
}

func (s *CreateTableMetaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTableMetaRequest) GetModule() *string {
	return s.Module
}

func (s *CreateTableMetaRequest) GetName() *string {
	return s.Name
}

func (s *CreateTableMetaRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateTableMetaRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateTableMetaRequest) SetDescription(v string) *CreateTableMetaRequest {
	s.Description = &v
	return s
}

func (s *CreateTableMetaRequest) SetFields(v []*CreateTableMetaRequestFields) *CreateTableMetaRequest {
	s.Fields = v
	return s
}

func (s *CreateTableMetaRequest) SetInstanceId(v string) *CreateTableMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTableMetaRequest) SetModule(v string) *CreateTableMetaRequest {
	s.Module = &v
	return s
}

func (s *CreateTableMetaRequest) SetName(v string) *CreateTableMetaRequest {
	s.Name = &v
	return s
}

func (s *CreateTableMetaRequest) SetResourceId(v string) *CreateTableMetaRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateTableMetaRequest) SetTableName(v string) *CreateTableMetaRequest {
	s.TableName = &v
	return s
}

func (s *CreateTableMetaRequest) Validate() error {
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

type CreateTableMetaRequestFields struct {
	// The logical data type of the field.
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
	// The business definition of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// this is gender of people
	Meaning *string `json:"Meaning,omitempty" xml:"Meaning,omitempty"`
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// gender
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

func (s CreateTableMetaRequestFields) String() string {
	return dara.Prettify(s)
}

func (s CreateTableMetaRequestFields) GoString() string {
	return s.String()
}

func (s *CreateTableMetaRequestFields) GetDataType() *string {
	return s.DataType
}

func (s *CreateTableMetaRequestFields) GetIsDimensionField() *bool {
	return s.IsDimensionField
}

func (s *CreateTableMetaRequestFields) GetIsPartitionField() *string {
	return s.IsPartitionField
}

func (s *CreateTableMetaRequestFields) GetMeaning() *string {
	return s.Meaning
}

func (s *CreateTableMetaRequestFields) GetName() *string {
	return s.Name
}

func (s *CreateTableMetaRequestFields) GetType() *string {
	return s.Type
}

func (s *CreateTableMetaRequestFields) SetDataType(v string) *CreateTableMetaRequestFields {
	s.DataType = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetIsDimensionField(v bool) *CreateTableMetaRequestFields {
	s.IsDimensionField = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetIsPartitionField(v string) *CreateTableMetaRequestFields {
	s.IsPartitionField = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetMeaning(v string) *CreateTableMetaRequestFields {
	s.Meaning = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetName(v string) *CreateTableMetaRequestFields {
	s.Name = &v
	return s
}

func (s *CreateTableMetaRequestFields) SetType(v string) *CreateTableMetaRequestFields {
	s.Type = &v
	return s
}

func (s *CreateTableMetaRequestFields) Validate() error {
	return dara.Validate(s)
}
