// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanDelete(v bool) *GetRecallManagementTableResponseBody
	GetCanDelete() *bool
	SetConfig(v string) *GetRecallManagementTableResponseBody
	GetConfig() *string
	SetDataSource(v string) *GetRecallManagementTableResponseBody
	GetDataSource() *string
	SetDescription(v string) *GetRecallManagementTableResponseBody
	GetDescription() *string
	SetEnableDataSizeFluctuationThreshold(v bool) *GetRecallManagementTableResponseBody
	GetEnableDataSizeFluctuationThreshold() *bool
	SetEnableRowCountFluctuationThreshold(v bool) *GetRecallManagementTableResponseBody
	GetEnableRowCountFluctuationThreshold() *bool
	SetFields(v []*GetRecallManagementTableResponseBodyFields) *GetRecallManagementTableResponseBody
	GetFields() []*GetRecallManagementTableResponseBodyFields
	SetGmtCreateTime(v string) *GetRecallManagementTableResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetRecallManagementTableResponseBody
	GetGmtModifiedTime() *string
	SetIndexEffectiveTime(v string) *GetRecallManagementTableResponseBody
	GetIndexEffectiveTime() *string
	SetIndexVersionId(v string) *GetRecallManagementTableResponseBody
	GetIndexVersionId() *string
	SetMaxDataSizeFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody
	GetMaxDataSizeFluctuationThreshold() *int32
	SetMaxRowCountFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody
	GetMaxRowCountFluctuationThreshold() *int32
	SetMaxcomputeProjectName(v string) *GetRecallManagementTableResponseBody
	GetMaxcomputeProjectName() *string
	SetMaxcomputeSchema(v string) *GetRecallManagementTableResponseBody
	GetMaxcomputeSchema() *string
	SetMaxcomputeTableName(v string) *GetRecallManagementTableResponseBody
	GetMaxcomputeTableName() *string
	SetMinDataSizeFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody
	GetMinDataSizeFluctuationThreshold() *int32
	SetMinRowCountFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody
	GetMinRowCountFluctuationThreshold() *int32
	SetName(v string) *GetRecallManagementTableResponseBody
	GetName() *string
	SetPartitionFields(v string) *GetRecallManagementTableResponseBody
	GetPartitionFields() *string
	SetRecallManagementTableId(v string) *GetRecallManagementTableResponseBody
	GetRecallManagementTableId() *string
	SetRecallType(v string) *GetRecallManagementTableResponseBody
	GetRecallType() *string
	SetRequestId(v string) *GetRecallManagementTableResponseBody
	GetRequestId() *string
	SetType(v string) *GetRecallManagementTableResponseBody
	GetType() *string
}

type GetRecallManagementTableResponseBody struct {
	// example:
	//
	// true
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// example:
	//
	// {"item_id":""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// Api
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// example:
	//
	// this is a test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	EnableDataSizeFluctuationThreshold *bool `json:"EnableDataSizeFluctuationThreshold,omitempty" xml:"EnableDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// true
	EnableRowCountFluctuationThreshold *bool                                         `json:"EnableRowCountFluctuationThreshold,omitempty" xml:"EnableRowCountFluctuationThreshold,omitempty"`
	Fields                             []*GetRecallManagementTableResponseBodyFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	IndexEffectiveTime *string `json:"IndexEffectiveTime,omitempty" xml:"IndexEffectiveTime,omitempty"`
	// example:
	//
	// 20250701
	IndexVersionId *string `json:"IndexVersionId,omitempty" xml:"IndexVersionId,omitempty"`
	// example:
	//
	// 100
	MaxDataSizeFluctuationThreshold *int32 `json:"MaxDataSizeFluctuationThreshold,omitempty" xml:"MaxDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// 100
	MaxRowCountFluctuationThreshold *int32 `json:"MaxRowCountFluctuationThreshold,omitempty" xml:"MaxRowCountFluctuationThreshold,omitempty"`
	// example:
	//
	// test
	MaxcomputeProjectName *string `json:"MaxcomputeProjectName,omitempty" xml:"MaxcomputeProjectName,omitempty"`
	// maxcompute schema。
	//
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
	MinDataSizeFluctuationThreshold *int32 `json:"MinDataSizeFluctuationThreshold,omitempty" xml:"MinDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// 10
	MinRowCountFluctuationThreshold *int32 `json:"MinRowCountFluctuationThreshold,omitempty" xml:"MinRowCountFluctuationThreshold,omitempty"`
	// example:
	//
	// table-123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// dt
	PartitionFields *string `json:"PartitionFields,omitempty" xml:"PartitionFields,omitempty"`
	// example:
	//
	// 3
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
	// example:
	//
	// X2I
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Recall
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRecallManagementTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecallManagementTableResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetRecallManagementTableResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetRecallManagementTableResponseBody) GetDataSource() *string {
	return s.DataSource
}

func (s *GetRecallManagementTableResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetRecallManagementTableResponseBody) GetEnableDataSizeFluctuationThreshold() *bool {
	return s.EnableDataSizeFluctuationThreshold
}

func (s *GetRecallManagementTableResponseBody) GetEnableRowCountFluctuationThreshold() *bool {
	return s.EnableRowCountFluctuationThreshold
}

func (s *GetRecallManagementTableResponseBody) GetFields() []*GetRecallManagementTableResponseBodyFields {
	return s.Fields
}

func (s *GetRecallManagementTableResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetRecallManagementTableResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetRecallManagementTableResponseBody) GetIndexEffectiveTime() *string {
	return s.IndexEffectiveTime
}

func (s *GetRecallManagementTableResponseBody) GetIndexVersionId() *string {
	return s.IndexVersionId
}

func (s *GetRecallManagementTableResponseBody) GetMaxDataSizeFluctuationThreshold() *int32 {
	return s.MaxDataSizeFluctuationThreshold
}

func (s *GetRecallManagementTableResponseBody) GetMaxRowCountFluctuationThreshold() *int32 {
	return s.MaxRowCountFluctuationThreshold
}

func (s *GetRecallManagementTableResponseBody) GetMaxcomputeProjectName() *string {
	return s.MaxcomputeProjectName
}

func (s *GetRecallManagementTableResponseBody) GetMaxcomputeSchema() *string {
	return s.MaxcomputeSchema
}

func (s *GetRecallManagementTableResponseBody) GetMaxcomputeTableName() *string {
	return s.MaxcomputeTableName
}

func (s *GetRecallManagementTableResponseBody) GetMinDataSizeFluctuationThreshold() *int32 {
	return s.MinDataSizeFluctuationThreshold
}

func (s *GetRecallManagementTableResponseBody) GetMinRowCountFluctuationThreshold() *int32 {
	return s.MinRowCountFluctuationThreshold
}

func (s *GetRecallManagementTableResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementTableResponseBody) GetPartitionFields() *string {
	return s.PartitionFields
}

func (s *GetRecallManagementTableResponseBody) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *GetRecallManagementTableResponseBody) GetRecallType() *string {
	return s.RecallType
}

func (s *GetRecallManagementTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecallManagementTableResponseBody) GetType() *string {
	return s.Type
}

func (s *GetRecallManagementTableResponseBody) SetCanDelete(v bool) *GetRecallManagementTableResponseBody {
	s.CanDelete = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetConfig(v string) *GetRecallManagementTableResponseBody {
	s.Config = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetDataSource(v string) *GetRecallManagementTableResponseBody {
	s.DataSource = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetDescription(v string) *GetRecallManagementTableResponseBody {
	s.Description = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetEnableDataSizeFluctuationThreshold(v bool) *GetRecallManagementTableResponseBody {
	s.EnableDataSizeFluctuationThreshold = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetEnableRowCountFluctuationThreshold(v bool) *GetRecallManagementTableResponseBody {
	s.EnableRowCountFluctuationThreshold = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetFields(v []*GetRecallManagementTableResponseBodyFields) *GetRecallManagementTableResponseBody {
	s.Fields = v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetGmtCreateTime(v string) *GetRecallManagementTableResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetGmtModifiedTime(v string) *GetRecallManagementTableResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetIndexEffectiveTime(v string) *GetRecallManagementTableResponseBody {
	s.IndexEffectiveTime = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetIndexVersionId(v string) *GetRecallManagementTableResponseBody {
	s.IndexVersionId = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMaxDataSizeFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody {
	s.MaxDataSizeFluctuationThreshold = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMaxRowCountFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody {
	s.MaxRowCountFluctuationThreshold = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMaxcomputeProjectName(v string) *GetRecallManagementTableResponseBody {
	s.MaxcomputeProjectName = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMaxcomputeSchema(v string) *GetRecallManagementTableResponseBody {
	s.MaxcomputeSchema = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMaxcomputeTableName(v string) *GetRecallManagementTableResponseBody {
	s.MaxcomputeTableName = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMinDataSizeFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody {
	s.MinDataSizeFluctuationThreshold = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetMinRowCountFluctuationThreshold(v int32) *GetRecallManagementTableResponseBody {
	s.MinRowCountFluctuationThreshold = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetName(v string) *GetRecallManagementTableResponseBody {
	s.Name = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetPartitionFields(v string) *GetRecallManagementTableResponseBody {
	s.PartitionFields = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetRecallManagementTableId(v string) *GetRecallManagementTableResponseBody {
	s.RecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetRecallType(v string) *GetRecallManagementTableResponseBody {
	s.RecallType = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetRequestId(v string) *GetRecallManagementTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) SetType(v string) *GetRecallManagementTableResponseBody {
	s.Type = &v
	return s
}

func (s *GetRecallManagementTableResponseBody) Validate() error {
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

type GetRecallManagementTableResponseBodyFields struct {
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// STRING
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

func (s GetRecallManagementTableResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementTableResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetRecallManagementTableResponseBodyFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *GetRecallManagementTableResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementTableResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetRecallManagementTableResponseBodyFields) GetVectorDimension() *int32 {
	return s.VectorDimension
}

func (s *GetRecallManagementTableResponseBodyFields) GetVectorMetricType() *string {
	return s.VectorMetricType
}

func (s *GetRecallManagementTableResponseBodyFields) SetAttributes(v []*string) *GetRecallManagementTableResponseBodyFields {
	s.Attributes = v
	return s
}

func (s *GetRecallManagementTableResponseBodyFields) SetName(v string) *GetRecallManagementTableResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetRecallManagementTableResponseBodyFields) SetType(v string) *GetRecallManagementTableResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetRecallManagementTableResponseBodyFields) SetVectorDimension(v int32) *GetRecallManagementTableResponseBodyFields {
	s.VectorDimension = &v
	return s
}

func (s *GetRecallManagementTableResponseBodyFields) SetVectorMetricType(v string) *GetRecallManagementTableResponseBodyFields {
	s.VectorMetricType = &v
	return s
}

func (s *GetRecallManagementTableResponseBodyFields) Validate() error {
	return dara.Validate(s)
}
