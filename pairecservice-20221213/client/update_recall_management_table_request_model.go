// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableDataSizeFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest
	GetEnableDataSizeFluctuationThreshold() *bool
	SetEnableRowCountFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest
	GetEnableRowCountFluctuationThreshold() *bool
	SetFields(v *UpdateRecallManagementTableRequestFields) *UpdateRecallManagementTableRequest
	GetFields() *UpdateRecallManagementTableRequestFields
	SetIndexVersionId(v string) *UpdateRecallManagementTableRequest
	GetIndexVersionId() *string
	SetInstanceId(v string) *UpdateRecallManagementTableRequest
	GetInstanceId() *string
	SetMaxDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMaxDataSizeFluctuationThreshold() *int32
	SetMaxRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMaxRowCountFluctuationThreshold() *int32
	SetMinDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMinDataSizeFluctuationThreshold() *int32
	SetMinRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMinRowCountFluctuationThreshold() *int32
}

type UpdateRecallManagementTableRequest struct {
	// example:
	//
	// true
	EnableDataSizeFluctuationThreshold *bool `json:"EnableDataSizeFluctuationThreshold,omitempty" xml:"EnableDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// true
	EnableRowCountFluctuationThreshold *bool                                     `json:"EnableRowCountFluctuationThreshold,omitempty" xml:"EnableRowCountFluctuationThreshold,omitempty"`
	Fields                             *UpdateRecallManagementTableRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Struct"`
	// example:
	//
	// 20250701
	IndexVersionId *string `json:"IndexVersionId,omitempty" xml:"IndexVersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 10
	MinDataSizeFluctuationThreshold *int32 `json:"MinDataSizeFluctuationThreshold,omitempty" xml:"MinDataSizeFluctuationThreshold,omitempty"`
	// example:
	//
	// 10
	MinRowCountFluctuationThreshold *int32 `json:"MinRowCountFluctuationThreshold,omitempty" xml:"MinRowCountFluctuationThreshold,omitempty"`
}

func (s UpdateRecallManagementTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementTableRequest) GetEnableDataSizeFluctuationThreshold() *bool {
	return s.EnableDataSizeFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetEnableRowCountFluctuationThreshold() *bool {
	return s.EnableRowCountFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetFields() *UpdateRecallManagementTableRequestFields {
	return s.Fields
}

func (s *UpdateRecallManagementTableRequest) GetIndexVersionId() *string {
	return s.IndexVersionId
}

func (s *UpdateRecallManagementTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRecallManagementTableRequest) GetMaxDataSizeFluctuationThreshold() *int32 {
	return s.MaxDataSizeFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetMaxRowCountFluctuationThreshold() *int32 {
	return s.MaxRowCountFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetMinDataSizeFluctuationThreshold() *int32 {
	return s.MinDataSizeFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetMinRowCountFluctuationThreshold() *int32 {
	return s.MinRowCountFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) SetEnableDataSizeFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest {
	s.EnableDataSizeFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetEnableRowCountFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest {
	s.EnableRowCountFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetFields(v *UpdateRecallManagementTableRequestFields) *UpdateRecallManagementTableRequest {
	s.Fields = v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetIndexVersionId(v string) *UpdateRecallManagementTableRequest {
	s.IndexVersionId = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetInstanceId(v string) *UpdateRecallManagementTableRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMaxDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MaxDataSizeFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMaxRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MaxRowCountFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMinDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MinDataSizeFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMinRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MinRowCountFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) Validate() error {
	if s.Fields != nil {
		if err := s.Fields.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRecallManagementTableRequestFields struct {
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

func (s UpdateRecallManagementTableRequestFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementTableRequestFields) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementTableRequestFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *UpdateRecallManagementTableRequestFields) GetName() *string {
	return s.Name
}

func (s *UpdateRecallManagementTableRequestFields) GetType() *string {
	return s.Type
}

func (s *UpdateRecallManagementTableRequestFields) GetVectorDimension() *int32 {
	return s.VectorDimension
}

func (s *UpdateRecallManagementTableRequestFields) GetVectorMetricType() *string {
	return s.VectorMetricType
}

func (s *UpdateRecallManagementTableRequestFields) SetAttributes(v []*string) *UpdateRecallManagementTableRequestFields {
	s.Attributes = v
	return s
}

func (s *UpdateRecallManagementTableRequestFields) SetName(v string) *UpdateRecallManagementTableRequestFields {
	s.Name = &v
	return s
}

func (s *UpdateRecallManagementTableRequestFields) SetType(v string) *UpdateRecallManagementTableRequestFields {
	s.Type = &v
	return s
}

func (s *UpdateRecallManagementTableRequestFields) SetVectorDimension(v int32) *UpdateRecallManagementTableRequestFields {
	s.VectorDimension = &v
	return s
}

func (s *UpdateRecallManagementTableRequestFields) SetVectorMetricType(v string) *UpdateRecallManagementTableRequestFields {
	s.VectorMetricType = &v
	return s
}

func (s *UpdateRecallManagementTableRequestFields) Validate() error {
	return dara.Validate(s)
}
