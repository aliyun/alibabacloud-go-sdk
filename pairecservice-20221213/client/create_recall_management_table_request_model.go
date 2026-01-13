// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateRecallManagementTableRequest
	GetConfig() *string
	SetDataSource(v string) *CreateRecallManagementTableRequest
	GetDataSource() *string
	SetDescription(v string) *CreateRecallManagementTableRequest
	GetDescription() *string
	SetEnableDataSizeFluctuationThreshold(v bool) *CreateRecallManagementTableRequest
	GetEnableDataSizeFluctuationThreshold() *bool
	SetEnableRowCountFluctuationThreshold(v bool) *CreateRecallManagementTableRequest
	GetEnableRowCountFluctuationThreshold() *bool
	SetFields(v []*CreateRecallManagementTableRequestFields) *CreateRecallManagementTableRequest
	GetFields() []*CreateRecallManagementTableRequestFields
	SetInstanceId(v string) *CreateRecallManagementTableRequest
	GetInstanceId() *string
	SetMaxDataSizeFluctuationThreshold(v int64) *CreateRecallManagementTableRequest
	GetMaxDataSizeFluctuationThreshold() *int64
	SetMaxRowCountFluctuationThreshold(v int64) *CreateRecallManagementTableRequest
	GetMaxRowCountFluctuationThreshold() *int64
	SetMaxcomputeProjectName(v string) *CreateRecallManagementTableRequest
	GetMaxcomputeProjectName() *string
	SetMaxcomputeSchema(v string) *CreateRecallManagementTableRequest
	GetMaxcomputeSchema() *string
	SetMaxcomputeTableName(v string) *CreateRecallManagementTableRequest
	GetMaxcomputeTableName() *string
	SetMinDataSizeFluctuationThreshold(v int64) *CreateRecallManagementTableRequest
	GetMinDataSizeFluctuationThreshold() *int64
	SetMinRowCountFluctuationThreshold(v int64) *CreateRecallManagementTableRequest
	GetMinRowCountFluctuationThreshold() *int64
	SetName(v string) *CreateRecallManagementTableRequest
	GetName() *string
	SetRecallType(v string) *CreateRecallManagementTableRequest
	GetRecallType() *string
	SetType(v string) *CreateRecallManagementTableRequest
	GetType() *string
}

type CreateRecallManagementTableRequest struct {
	// example:
	//
	// {"":""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// MaxcomputeAndApiApi
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// example:
	//
	// this is a test table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	EnableDataSizeFluctuationThreshold *bool `json:"EnableDataSizeFluctuationThreshold,omitempty" xml:"EnableDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// true
	EnableRowCountFluctuationThreshold *bool                                       `json:"EnableRowCountFluctuationThreshold,omitempty" xml:"EnableRowCountFluctuationThreshold,omitempty"`
	Fields                             []*CreateRecallManagementTableRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 100
	MaxDataSizeFluctuationThreshold *int64 `json:"MaxDataSizeFluctuationThreshold,omitempty" xml:"MaxDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// 100
	MaxRowCountFluctuationThreshold *int64 `json:"MaxRowCountFluctuationThreshold,omitempty" xml:"MaxRowCountFluctuationThreshold,omitempty"`
	// example:
	//
	// test
	MaxcomputeProjectName *string `json:"MaxcomputeProjectName,omitempty" xml:"MaxcomputeProjectName,omitempty"`
	// example:
	//
	// default
	MaxcomputeSchema *string `json:"MaxcomputeSchema,omitempty" xml:"MaxcomputeSchema,omitempty"`
	// example:
	//
	// table-1
	MaxcomputeTableName *string `json:"MaxcomputeTableName,omitempty" xml:"MaxcomputeTableName,omitempty"`
	// example:
	//
	// 10
	MinDataSizeFluctuationThreshold *int64 `json:"MinDataSizeFluctuationThreshold,omitempty" xml:"MinDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// 10
	MinRowCountFluctuationThreshold *int64 `json:"MinRowCountFluctuationThreshold,omitempty" xml:"MinRowCountFluctuationThreshold,omitempty"`
	// example:
	//
	// table-123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// X2I
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// example:
	//
	// Recall
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecallManagementTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementTableRequest) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementTableRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateRecallManagementTableRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *CreateRecallManagementTableRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRecallManagementTableRequest) GetEnableDataSizeFluctuationThreshold() *bool {
	return s.EnableDataSizeFluctuationThreshold
}

func (s *CreateRecallManagementTableRequest) GetEnableRowCountFluctuationThreshold() *bool {
	return s.EnableRowCountFluctuationThreshold
}

func (s *CreateRecallManagementTableRequest) GetFields() []*CreateRecallManagementTableRequestFields {
	return s.Fields
}

func (s *CreateRecallManagementTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRecallManagementTableRequest) GetMaxDataSizeFluctuationThreshold() *int64 {
	return s.MaxDataSizeFluctuationThreshold
}

func (s *CreateRecallManagementTableRequest) GetMaxRowCountFluctuationThreshold() *int64 {
	return s.MaxRowCountFluctuationThreshold
}

func (s *CreateRecallManagementTableRequest) GetMaxcomputeProjectName() *string {
	return s.MaxcomputeProjectName
}

func (s *CreateRecallManagementTableRequest) GetMaxcomputeSchema() *string {
	return s.MaxcomputeSchema
}

func (s *CreateRecallManagementTableRequest) GetMaxcomputeTableName() *string {
	return s.MaxcomputeTableName
}

func (s *CreateRecallManagementTableRequest) GetMinDataSizeFluctuationThreshold() *int64 {
	return s.MinDataSizeFluctuationThreshold
}

func (s *CreateRecallManagementTableRequest) GetMinRowCountFluctuationThreshold() *int64 {
	return s.MinRowCountFluctuationThreshold
}

func (s *CreateRecallManagementTableRequest) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementTableRequest) GetRecallType() *string {
	return s.RecallType
}

func (s *CreateRecallManagementTableRequest) GetType() *string {
	return s.Type
}

func (s *CreateRecallManagementTableRequest) SetConfig(v string) *CreateRecallManagementTableRequest {
	s.Config = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetDataSource(v string) *CreateRecallManagementTableRequest {
	s.DataSource = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetDescription(v string) *CreateRecallManagementTableRequest {
	s.Description = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetEnableDataSizeFluctuationThreshold(v bool) *CreateRecallManagementTableRequest {
	s.EnableDataSizeFluctuationThreshold = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetEnableRowCountFluctuationThreshold(v bool) *CreateRecallManagementTableRequest {
	s.EnableRowCountFluctuationThreshold = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetFields(v []*CreateRecallManagementTableRequestFields) *CreateRecallManagementTableRequest {
	s.Fields = v
	return s
}

func (s *CreateRecallManagementTableRequest) SetInstanceId(v string) *CreateRecallManagementTableRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMaxDataSizeFluctuationThreshold(v int64) *CreateRecallManagementTableRequest {
	s.MaxDataSizeFluctuationThreshold = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMaxRowCountFluctuationThreshold(v int64) *CreateRecallManagementTableRequest {
	s.MaxRowCountFluctuationThreshold = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMaxcomputeProjectName(v string) *CreateRecallManagementTableRequest {
	s.MaxcomputeProjectName = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMaxcomputeSchema(v string) *CreateRecallManagementTableRequest {
	s.MaxcomputeSchema = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMaxcomputeTableName(v string) *CreateRecallManagementTableRequest {
	s.MaxcomputeTableName = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMinDataSizeFluctuationThreshold(v int64) *CreateRecallManagementTableRequest {
	s.MinDataSizeFluctuationThreshold = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetMinRowCountFluctuationThreshold(v int64) *CreateRecallManagementTableRequest {
	s.MinRowCountFluctuationThreshold = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetName(v string) *CreateRecallManagementTableRequest {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetRecallType(v string) *CreateRecallManagementTableRequest {
	s.RecallType = &v
	return s
}

func (s *CreateRecallManagementTableRequest) SetType(v string) *CreateRecallManagementTableRequest {
	s.Type = &v
	return s
}

func (s *CreateRecallManagementTableRequest) Validate() error {
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

type CreateRecallManagementTableRequestFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 32
	VectorDimension *int32 `json:"VectorDimension,omitempty" xml:"VectorDimension,omitempty"`
	// example:
	//
	// L2
	VectorMetricType *string `json:"VectorMetricType,omitempty" xml:"VectorMetricType,omitempty"`
}

func (s CreateRecallManagementTableRequestFields) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementTableRequestFields) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementTableRequestFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *CreateRecallManagementTableRequestFields) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementTableRequestFields) GetType() *string {
	return s.Type
}

func (s *CreateRecallManagementTableRequestFields) GetVectorDimension() *int32 {
	return s.VectorDimension
}

func (s *CreateRecallManagementTableRequestFields) GetVectorMetricType() *string {
	return s.VectorMetricType
}

func (s *CreateRecallManagementTableRequestFields) SetAttributes(v []*string) *CreateRecallManagementTableRequestFields {
	s.Attributes = v
	return s
}

func (s *CreateRecallManagementTableRequestFields) SetName(v string) *CreateRecallManagementTableRequestFields {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementTableRequestFields) SetType(v string) *CreateRecallManagementTableRequestFields {
	s.Type = &v
	return s
}

func (s *CreateRecallManagementTableRequestFields) SetVectorDimension(v int32) *CreateRecallManagementTableRequestFields {
	s.VectorDimension = &v
	return s
}

func (s *CreateRecallManagementTableRequestFields) SetVectorMetricType(v string) *CreateRecallManagementTableRequestFields {
	s.VectorMetricType = &v
	return s
}

func (s *CreateRecallManagementTableRequestFields) Validate() error {
	return dara.Validate(s)
}
