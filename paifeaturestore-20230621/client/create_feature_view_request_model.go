// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateFeatureViewRequest
	GetConfig() *string
	SetFeatureEntityId(v string) *CreateFeatureViewRequest
	GetFeatureEntityId() *string
	SetFields(v []*CreateFeatureViewRequestFields) *CreateFeatureViewRequest
	GetFields() []*CreateFeatureViewRequestFields
	SetName(v string) *CreateFeatureViewRequest
	GetName() *string
	SetProjectId(v string) *CreateFeatureViewRequest
	GetProjectId() *string
	SetRegisterDatasourceId(v string) *CreateFeatureViewRequest
	GetRegisterDatasourceId() *string
	SetRegisterTable(v string) *CreateFeatureViewRequest
	GetRegisterTable() *string
	SetSyncOnlineTable(v bool) *CreateFeatureViewRequest
	GetSyncOnlineTable() *bool
	SetTTL(v int32) *CreateFeatureViewRequest
	GetTTL() *int32
	SetTags(v []*string) *CreateFeatureViewRequest
	GetTags() []*string
	SetType(v string) *CreateFeatureViewRequest
	GetType() *string
	SetWriteMethod(v string) *CreateFeatureViewRequest
	GetWriteMethod() *string
	SetWriteToFeatureDB(v bool) *CreateFeatureViewRequest
	GetWriteToFeatureDB() *bool
}

type CreateFeatureViewRequest struct {
	// The configurations of the feature view.
	//
	// example:
	//
	// {"save_original_field":true}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the feature entity.
	//
	// example:
	//
	// 4
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// The fields.
	Fields []*CreateFeatureViewRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The name of the feature view.
	//
	// This parameter is required.
	//
	// example:
	//
	// FeatureView1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the data source to which the table to be registered belongs. You can call the ListDatasources operation to obtain the data source ID.
	//
	// example:
	//
	// 5
	RegisterDatasourceId *string `json:"RegisterDatasourceId,omitempty" xml:"RegisterDatasourceId,omitempty"`
	// The name of the table to register.
	//
	// example:
	//
	// table1
	RegisterTable *string `json:"RegisterTable,omitempty" xml:"RegisterTable,omitempty"`
	// Specifies whether to synchronize the online feature table. Valid values:
	//
	// - `true`: Synchronizes the online feature table.
	//
	// - `false`: Does not synchronize the online feature table.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SyncOnlineTable *bool `json:"SyncOnlineTable,omitempty" xml:"SyncOnlineTable,omitempty"`
	// The time-to-live (TTL) of the feature view, in days.
	//
	// example:
	//
	// 90
	TTL *int32 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The tags of the feature view.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the feature view. Valid values:
	//
	// - `Batch`: Offline feature.
	//
	// - `Stream`: Real-time feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Batch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write method. Valid values:
	//
	// - `ByReadyMadeTable`: Registers the feature view by using an existing table.
	//
	// - `Custom`: Defines a custom table structure.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	WriteMethod *string `json:"WriteMethod,omitempty" xml:"WriteMethod,omitempty"`
	// Specifies whether to write data to the online feature store. Valid values:
	//
	// - `true`: Writes data to the online feature store.
	//
	// - `false`: Does not write data to the online feature store.
	//
	// example:
	//
	// false
	WriteToFeatureDB *bool `json:"WriteToFeatureDB,omitempty" xml:"WriteToFeatureDB,omitempty"`
}

func (s CreateFeatureViewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureViewRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateFeatureViewRequest) GetFeatureEntityId() *string {
	return s.FeatureEntityId
}

func (s *CreateFeatureViewRequest) GetFields() []*CreateFeatureViewRequestFields {
	return s.Fields
}

func (s *CreateFeatureViewRequest) GetName() *string {
	return s.Name
}

func (s *CreateFeatureViewRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateFeatureViewRequest) GetRegisterDatasourceId() *string {
	return s.RegisterDatasourceId
}

func (s *CreateFeatureViewRequest) GetRegisterTable() *string {
	return s.RegisterTable
}

func (s *CreateFeatureViewRequest) GetSyncOnlineTable() *bool {
	return s.SyncOnlineTable
}

func (s *CreateFeatureViewRequest) GetTTL() *int32 {
	return s.TTL
}

func (s *CreateFeatureViewRequest) GetTags() []*string {
	return s.Tags
}

func (s *CreateFeatureViewRequest) GetType() *string {
	return s.Type
}

func (s *CreateFeatureViewRequest) GetWriteMethod() *string {
	return s.WriteMethod
}

func (s *CreateFeatureViewRequest) GetWriteToFeatureDB() *bool {
	return s.WriteToFeatureDB
}

func (s *CreateFeatureViewRequest) SetConfig(v string) *CreateFeatureViewRequest {
	s.Config = &v
	return s
}

func (s *CreateFeatureViewRequest) SetFeatureEntityId(v string) *CreateFeatureViewRequest {
	s.FeatureEntityId = &v
	return s
}

func (s *CreateFeatureViewRequest) SetFields(v []*CreateFeatureViewRequestFields) *CreateFeatureViewRequest {
	s.Fields = v
	return s
}

func (s *CreateFeatureViewRequest) SetName(v string) *CreateFeatureViewRequest {
	s.Name = &v
	return s
}

func (s *CreateFeatureViewRequest) SetProjectId(v string) *CreateFeatureViewRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFeatureViewRequest) SetRegisterDatasourceId(v string) *CreateFeatureViewRequest {
	s.RegisterDatasourceId = &v
	return s
}

func (s *CreateFeatureViewRequest) SetRegisterTable(v string) *CreateFeatureViewRequest {
	s.RegisterTable = &v
	return s
}

func (s *CreateFeatureViewRequest) SetSyncOnlineTable(v bool) *CreateFeatureViewRequest {
	s.SyncOnlineTable = &v
	return s
}

func (s *CreateFeatureViewRequest) SetTTL(v int32) *CreateFeatureViewRequest {
	s.TTL = &v
	return s
}

func (s *CreateFeatureViewRequest) SetTags(v []*string) *CreateFeatureViewRequest {
	s.Tags = v
	return s
}

func (s *CreateFeatureViewRequest) SetType(v string) *CreateFeatureViewRequest {
	s.Type = &v
	return s
}

func (s *CreateFeatureViewRequest) SetWriteMethod(v string) *CreateFeatureViewRequest {
	s.WriteMethod = &v
	return s
}

func (s *CreateFeatureViewRequest) SetWriteToFeatureDB(v bool) *CreateFeatureViewRequest {
	s.WriteToFeatureDB = &v
	return s
}

func (s *CreateFeatureViewRequest) Validate() error {
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

type CreateFeatureViewRequestFields struct {
	// The attributes of the field. Valid values:
	//
	// - `Partition`: partition field.
	//
	// - `PrimaryKey`: primary key.
	//
	// - `EventTime`: event time.
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the field.
	//
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The feature generation configurations.
	Transform []*CreateFeatureViewRequestFieldsTransform `json:"Transform,omitempty" xml:"Transform,omitempty" type:"Repeated"`
	// The data type of the field. Valid values:
	//
	// - INT32
	//
	// - INT64
	//
	// - FLOAT
	//
	// - DOUBLE
	//
	// - STRING
	//
	// - BOOLEAN
	//
	// - TIMESTAMP
	//
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFeatureViewRequestFields) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureViewRequestFields) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewRequestFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *CreateFeatureViewRequestFields) GetName() *string {
	return s.Name
}

func (s *CreateFeatureViewRequestFields) GetTransform() []*CreateFeatureViewRequestFieldsTransform {
	return s.Transform
}

func (s *CreateFeatureViewRequestFields) GetType() *string {
	return s.Type
}

func (s *CreateFeatureViewRequestFields) SetAttributes(v []*string) *CreateFeatureViewRequestFields {
	s.Attributes = v
	return s
}

func (s *CreateFeatureViewRequestFields) SetName(v string) *CreateFeatureViewRequestFields {
	s.Name = &v
	return s
}

func (s *CreateFeatureViewRequestFields) SetTransform(v []*CreateFeatureViewRequestFieldsTransform) *CreateFeatureViewRequestFields {
	s.Transform = v
	return s
}

func (s *CreateFeatureViewRequestFields) SetType(v string) *CreateFeatureViewRequestFields {
	s.Type = &v
	return s
}

func (s *CreateFeatureViewRequestFields) Validate() error {
	if s.Transform != nil {
		for _, item := range s.Transform {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFeatureViewRequestFieldsTransform struct {
	// The input for feature generation.
	Input []*CreateFeatureViewRequestFieldsTransformInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// The ID of the large language model (LLM) configuration.
	//
	// example:
	//
	// 1
	LLMConfigId *int32 `json:"LLMConfigId,omitempty" xml:"LLMConfigId,omitempty"`
	// The type of the feature generation.
	//
	// example:
	//
	// LLMEmbedding
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFeatureViewRequestFieldsTransform) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureViewRequestFieldsTransform) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewRequestFieldsTransform) GetInput() []*CreateFeatureViewRequestFieldsTransformInput {
	return s.Input
}

func (s *CreateFeatureViewRequestFieldsTransform) GetLLMConfigId() *int32 {
	return s.LLMConfigId
}

func (s *CreateFeatureViewRequestFieldsTransform) GetType() *string {
	return s.Type
}

func (s *CreateFeatureViewRequestFieldsTransform) SetInput(v []*CreateFeatureViewRequestFieldsTransformInput) *CreateFeatureViewRequestFieldsTransform {
	s.Input = v
	return s
}

func (s *CreateFeatureViewRequestFieldsTransform) SetLLMConfigId(v int32) *CreateFeatureViewRequestFieldsTransform {
	s.LLMConfigId = &v
	return s
}

func (s *CreateFeatureViewRequestFieldsTransform) SetType(v string) *CreateFeatureViewRequestFieldsTransform {
	s.Type = &v
	return s
}

func (s *CreateFeatureViewRequestFieldsTransform) Validate() error {
	if s.Input != nil {
		for _, item := range s.Input {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFeatureViewRequestFieldsTransformInput struct {
	// The modality of the input, such as text or image.
	Modality *string `json:"Modality,omitempty" xml:"Modality,omitempty"`
	// The name of the input field.
	//
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the input field.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFeatureViewRequestFieldsTransformInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureViewRequestFieldsTransformInput) GoString() string {
	return s.String()
}

func (s *CreateFeatureViewRequestFieldsTransformInput) GetModality() *string {
	return s.Modality
}

func (s *CreateFeatureViewRequestFieldsTransformInput) GetName() *string {
	return s.Name
}

func (s *CreateFeatureViewRequestFieldsTransformInput) GetType() *string {
	return s.Type
}

func (s *CreateFeatureViewRequestFieldsTransformInput) SetModality(v string) *CreateFeatureViewRequestFieldsTransformInput {
	s.Modality = &v
	return s
}

func (s *CreateFeatureViewRequestFieldsTransformInput) SetName(v string) *CreateFeatureViewRequestFieldsTransformInput {
	s.Name = &v
	return s
}

func (s *CreateFeatureViewRequestFieldsTransformInput) SetType(v string) *CreateFeatureViewRequestFieldsTransformInput {
	s.Type = &v
	return s
}

func (s *CreateFeatureViewRequestFieldsTransformInput) Validate() error {
	return dara.Validate(s)
}
