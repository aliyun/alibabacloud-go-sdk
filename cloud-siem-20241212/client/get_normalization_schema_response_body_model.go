// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNormalizationSchema(v *GetNormalizationSchemaResponseBodyNormalizationSchema) *GetNormalizationSchemaResponseBody
	GetNormalizationSchema() *GetNormalizationSchemaResponseBodyNormalizationSchema
	SetRequestId(v string) *GetNormalizationSchemaResponseBody
	GetRequestId() *string
}

type GetNormalizationSchemaResponseBody struct {
	// The normalization schema.
	NormalizationSchema *GetNormalizationSchemaResponseBodyNormalizationSchema `json:"NormalizationSchema,omitempty" xml:"NormalizationSchema,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNormalizationSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBody) GetNormalizationSchema() *GetNormalizationSchemaResponseBodyNormalizationSchema {
	return s.NormalizationSchema
}

func (s *GetNormalizationSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNormalizationSchemaResponseBody) SetNormalizationSchema(v *GetNormalizationSchemaResponseBodyNormalizationSchema) *GetNormalizationSchemaResponseBody {
	s.NormalizationSchema = v
	return s
}

func (s *GetNormalizationSchemaResponseBody) SetRequestId(v string) *GetNormalizationSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBody) Validate() error {
	if s.NormalizationSchema != nil {
		if err := s.NormalizationSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNormalizationSchemaResponseBodyNormalizationSchema struct {
	// The creation time.
	//
	// example:
	//
	// 1733269771123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the normalization rule category.
	//
	// example:
	//
	// NETWORK_CATEGORY
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// native
	NormalizationFieldSource *string `json:"NormalizationFieldSource,omitempty" xml:"NormalizationFieldSource,omitempty"`
	// The list of normalization fields.
	NormalizationFields []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields `json:"NormalizationFields,omitempty" xml:"NormalizationFields,omitempty" type:"Repeated"`
	// The normalization schema description.
	//
	// example:
	//
	// Network 5-Tuple Logs
	NormalizationSchemaDescription *string `json:"NormalizationSchemaDescription,omitempty" xml:"NormalizationSchemaDescription,omitempty"`
	// The source of the normalization schema. Valid values:
	//
	// - preset: predefined.
	//
	// - custom: custom.
	//
	// example:
	//
	// preset
	NormalizationSchemaFrom *string `json:"NormalizationSchemaFrom,omitempty" xml:"NormalizationSchemaFrom,omitempty"`
	// The normalization schema ID.
	//
	// example:
	//
	// HTTP_ACTIVITY
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
	// The normalization schema name.
	//
	// example:
	//
	// normalization_rule_Z57np
	NormalizationSchemaName *string `json:"NormalizationSchemaName,omitempty" xml:"NormalizationSchemaName,omitempty"`
	// The list of normalization schema references.
	NormalizationSchemaReferences []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences `json:"NormalizationSchemaReferences,omitempty" xml:"NormalizationSchemaReferences,omitempty" type:"Repeated"`
	// The normalization schema type. Valid values:
	//
	// - log
	//
	// - entity
	//
	// - incident
	//
	// example:
	//
	// log
	NormalizationSchemaType *string `json:"NormalizationSchemaType,omitempty" xml:"NormalizationSchemaType,omitempty"`
	// example:
	//
	// DATA_SECURITY
	NormalizationSecurityDomainId *string `json:"NormalizationSecurityDomainId,omitempty" xml:"NormalizationSecurityDomainId,omitempty"`
	// example:
	//
	// alibaba_cloud_bastionhost
	ProductId         *string   `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RecommendEntities []*string `json:"RecommendEntities,omitempty" xml:"RecommendEntities,omitempty" type:"Repeated"`
	// The Simple Log Service LogStore.
	//
	// example:
	//
	// network-activity
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// The Simple Log Service StoreView.
	//
	// example:
	//
	// network_activity
	TargetStoreView *string `json:"TargetStoreView,omitempty" xml:"TargetStoreView,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1733269771123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// alibaba_cloud
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchema) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchema) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationFieldSource() *string {
	return s.NormalizationFieldSource
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationFields() []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	return s.NormalizationFields
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaDescription() *string {
	return s.NormalizationSchemaDescription
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaFrom() *string {
	return s.NormalizationSchemaFrom
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaName() *string {
	return s.NormalizationSchemaName
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaReferences() []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences {
	return s.NormalizationSchemaReferences
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSchemaType() *string {
	return s.NormalizationSchemaType
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetNormalizationSecurityDomainId() *string {
	return s.NormalizationSecurityDomainId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetProductId() *string {
	return s.ProductId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetRecommendEntities() []*string {
	return s.RecommendEntities
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetTargetLogStore() *string {
	return s.TargetLogStore
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetTargetStoreView() *string {
	return s.TargetStoreView
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) GetVendorId() *string {
	return s.VendorId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetCreateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.CreateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationCategoryId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationCategoryId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationFieldSource(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationFieldSource = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationFields(v []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationFields = v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaDescription(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaDescription = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaFrom(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaFrom = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaName(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaName = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaReferences(v []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaReferences = v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSchemaType(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSchemaType = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetNormalizationSecurityDomainId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.NormalizationSecurityDomainId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetProductId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.ProductId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetRecommendEntities(v []*string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.RecommendEntities = v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetTargetLogStore(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.TargetLogStore = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetTargetStoreView(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.TargetStoreView = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetUpdateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.UpdateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) SetVendorId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchema {
	s.VendorId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchema) Validate() error {
	if s.NormalizationFields != nil {
		for _, item := range s.NormalizationFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NormalizationSchemaReferences != nil {
		for _, item := range s.NormalizationSchemaReferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields struct {
	// The creation time.
	//
	// example:
	//
	// 1736386610000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The normalization field description.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// The normalization field example.
	//
	// example:
	//
	// 173326*******
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// The source of the standard field. Valid values:
	//
	// - preset: built-in.
	//
	// - custom: custom.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// Indicates whether indexes are created for all keys of the JSON-type standard field.
	NormalizationFieldJsonIndexAll *bool `json:"NormalizationFieldJsonIndexAll,omitempty" xml:"NormalizationFieldJsonIndexAll,omitempty"`
	// The key list of the JSON-type standard field.
	NormalizationFieldJsonKeys []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys `json:"NormalizationFieldJsonKeys,omitempty" xml:"NormalizationFieldJsonKeys,omitempty" type:"Repeated"`
	// The normalization field name.
	//
	// example:
	//
	// cloud_user
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Indicates whether the field is required.
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Indicates whether the normalization field is required.
	//
	// example:
	//
	// true
	NormalizationFieldRequirement *bool `json:"NormalizationFieldRequirement,omitempty" xml:"NormalizationFieldRequirement,omitempty"`
	// Indicates whether the normalization field is reserved.
	//
	// example:
	//
	// true
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// Indicates whether the standard field is tokenized.
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// The normalization field type. Valid values:
	//
	// - varchar
	//
	// - bigint
	//
	// - double
	//
	// example:
	//
	// varchar
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1736386610000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldJsonIndexAll() *bool {
	return s.NormalizationFieldJsonIndexAll
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldJsonKeys() []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	return s.NormalizationFieldJsonKeys
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldRequirement() *bool {
	return s.NormalizationFieldRequirement
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetCreateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.CreateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldDescription(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldExample(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldExample = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldFrom(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldJsonIndexAll(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldJsonIndexAll = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldJsonKeys(v []*GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldJsonKeys = v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldName(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldName = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldRequired(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldRequirement(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldRequirement = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldReserved(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldTokenize(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetNormalizationFieldType(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.NormalizationFieldType = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) SetUpdateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields {
	s.UpdateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFields) Validate() error {
	if s.NormalizationFieldJsonKeys != nil {
		for _, item := range s.NormalizationFieldJsonKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys struct {
	// The creation time.
	//
	// example:
	//
	// 1736386610000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The key description of the JSON-type standard field.
	//
	// example:
	//
	// Cloud Provider Code
	NormalizationFieldDescription *string `json:"NormalizationFieldDescription,omitempty" xml:"NormalizationFieldDescription,omitempty"`
	// The key example of the JSON-type standard field.
	//
	// example:
	//
	// alibaba_cloud
	NormalizationFieldExample *string `json:"NormalizationFieldExample,omitempty" xml:"NormalizationFieldExample,omitempty"`
	// The key source of the JSON-type standard field.
	//
	// example:
	//
	// preset
	NormalizationFieldFrom *string `json:"NormalizationFieldFrom,omitempty" xml:"NormalizationFieldFrom,omitempty"`
	// The key name of the JSON-type standard field.
	//
	// example:
	//
	// cloud_code
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// Indicates whether the key of the JSON-type standard field is required.
	NormalizationFieldRequired *bool `json:"NormalizationFieldRequired,omitempty" xml:"NormalizationFieldRequired,omitempty"`
	// Indicates whether the field is a system built-in standard field name.
	NormalizationFieldReserved *bool `json:"NormalizationFieldReserved,omitempty" xml:"NormalizationFieldReserved,omitempty"`
	// Indicates whether the key of the JSON-type standard field is tokenized.
	NormalizationFieldTokenize *bool `json:"NormalizationFieldTokenize,omitempty" xml:"NormalizationFieldTokenize,omitempty"`
	// The key type of the JSON-type standard field.
	//
	// example:
	//
	// varchar
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1736386610000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldDescription() *string {
	return s.NormalizationFieldDescription
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldExample() *string {
	return s.NormalizationFieldExample
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldFrom() *string {
	return s.NormalizationFieldFrom
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldRequired() *bool {
	return s.NormalizationFieldRequired
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldReserved() *bool {
	return s.NormalizationFieldReserved
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldTokenize() *bool {
	return s.NormalizationFieldTokenize
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetCreateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.CreateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldDescription(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldDescription = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldExample(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldExample = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldFrom(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldFrom = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldName(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldName = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldRequired(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldRequired = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldReserved(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldReserved = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldTokenize(v bool) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldTokenize = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetNormalizationFieldType(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.NormalizationFieldType = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) SetUpdateTime(v int64) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys {
	s.UpdateTime = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationFieldsNormalizationFieldJsonKeys) Validate() error {
	return dara.Validate(s)
}

type GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences struct {
	// The normalization rule ID.
	//
	// example:
	//
	// alibaba_cloud_cfw_flow_rule
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences) SetNormalizationRuleId(v string) *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences {
	s.NormalizationRuleId = &v
	return s
}

func (s *GetNormalizationSchemaResponseBodyNormalizationSchemaNormalizationSchemaReferences) Validate() error {
	return dara.Validate(s)
}
